package aggregator

import (
	"context"
	"errors"
	"fmt"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"net/http"
	"net/rpc"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	SignatureVerificationFailed400           = errors.New("400. Signature verification failed")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (agg *Aggregator) startServer(ctx context.Context) error {

	err := rpc.Register(agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.serverIpPortAddr, nil)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

type SignedTaskResponse struct {
	TaskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	Task         cstaskmanager.IIncredibleSquaringTaskManagerTask
	BlsSignature bls.Signature
	OperatorId   types.OperatorId
	SignedTx string
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator) ProcessSignedTaskResponse(signedTaskResponse *SignedTaskResponse, reply *bool) error {

	agg.logger.Infof("Received signed task, addint task to bls agg service: %#v", agg.tasks[signedTaskResponse.TaskResponse.ReferenceTaskIndex])
	if task, ok := agg.tasks[signedTaskResponse.TaskResponse.ReferenceTaskIndex]; !ok || isTaskEmpty(task) {

		agg.tasksMu.Lock()
		agg.tasks[signedTaskResponse.TaskResponse.ReferenceTaskIndex] = signedTaskResponse.Task
		agg.tasksMu.Unlock()
		// The entry does not exist or is empty (nil)
		quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(signedTaskResponse.Task.QuorumNumbers))
		for i := range signedTaskResponse.Task.QuorumNumbers {
			quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(signedTaskResponse.Task.QuorumThresholdPercentage)
		}
		// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
		// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
		taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
		var quorumNums sdktypes.QuorumNums
		for _, quorumNum := range signedTaskResponse.Task.QuorumNumbers {
			quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
		}
		agg.txSignatures[signedTaskResponse.TaskResponse.ReferenceTaskIndex] = []string{}
		agg.blsAggregationService.InitializeNewTask(signedTaskResponse.TaskResponse.ReferenceTaskIndex, signedTaskResponse.Task.TaskCreatedBlock, quorumNums, quorumThresholdPercentages, taskTimeToExpiry)
	} else {
		// The entry exists and is not empty
		fmt.Println("The task entry exists and is not empty:", task)
	}

	// Ovde dodajem u niz
	agg.txSignatures[signedTaskResponse.TaskResponse.ReferenceTaskIndex] = append(agg.txSignatures[signedTaskResponse.TaskResponse.ReferenceTaskIndex], signedTaskResponse.SignedTx)

	agg.logger.Infof("Received signed task response: %#v", signedTaskResponse)
	taskIndex := signedTaskResponse.TaskResponse.ReferenceTaskIndex
	taskResponseDigest, err := core.GetMintTaskResponseDigest(&signedTaskResponse.TaskResponse)
	if err != nil {
		agg.logger.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}

	agg.logger.Infof("Received signed task, addint task to bls agg service: %#v", agg.tasks[signedTaskResponse.TaskResponse.ReferenceTaskIndex])

	agg.taskResponsesMu.Lock()
	if _, ok := agg.taskResponses[taskIndex]; !ok {
		agg.taskResponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse)
	}
	if _, ok := agg.taskResponses[taskIndex][taskResponseDigest]; !ok {
		agg.taskResponses[taskIndex][taskResponseDigest] = signedTaskResponse.TaskResponse
	}
	agg.taskResponsesMu.Unlock()

	err = agg.blsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, taskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId,
	)
	return err
}

//helper functions

// Function to check if the task is empty
func isTaskEmpty(task cstaskmanager.IIncredibleSquaringTaskManagerTask) bool {
	return task.BtcTxHash == "" &&
		task.SignedMessage == "" &&
		task.MintTo == gethcommon.Address{} &&
		task.TaskCreatedBlock == 0 &&
		task.BurnAmount == nil &&
		!task.IsBurnTask &&
		len(task.QuorumNumbers) == 0 &&
		task.QuorumThresholdPercentage == 0
}

// In your main logic
