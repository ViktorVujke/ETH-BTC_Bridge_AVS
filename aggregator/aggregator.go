package aggregator

import (
    "bytes"
    "encoding/json"
	"net/http"
	"context"
	"math/big"
	"sync"
	"time"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
)

const (
	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
	avsName                  = "incredible-squaring"
)

// Aggregator sends tasks (numbers to square) onchain, then listens for operator signed TaskResponses.
// It aggregates responses signatures, and if any of the TaskResponses reaches the QuorumThresholdPercentage for each quorum
// (currently we only use a single quorum of the ERC20Mock token), it sends the aggregated TaskResponse and signature onchain.
//
// The signature is checked in the BLSSignatureChecker.sol contract, which expects a
//
//	struct NonSignerStakesAndSignature {
//		uint32[] nonSignerQuorumBitmapIndices;
//		BN254.G1Point[] nonSignerPubkeys;
//		BN254.G1Point[] quorumApks;
//		BN254.G2Point apkG2;
//		BN254.G1Point sigma;
//		uint32[] quorumApkIndices;
//		uint32[] totalStakeIndices;
//		uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
//	}
//
// A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThresholdPercentage.
// In order to verify this onchain, the Registry contracts store the history of stakes and aggregate pubkeys (apks) for each operators and each quorum. These are
// updated everytime an operator registers or deregisters with the BLSRegistryCoordinatorWithIndices.sol contract, or calls UpdateStakes() on the StakeRegistry.sol contract,
// after having received new delegated shares or having delegated shares removed by stakers queuing withdrawals. Each of these pushes to their respective datatype array a new entry.
//
// This is true for quorumBitmaps (represent the quorums each operator is opted into), quorumApks (apks per quorum), totalStakes (total stake per quorum), and nonSignerStakes (stake per quorum per operator).
// The 4 "indices" in NonSignerStakesAndSignature basically represent the index at which to fetch their respective data, given a blockNumber at which the task was created.
// Note that different data types might have different indices, since for eg QuorumBitmaps are updated for operators registering/deregistering, but not for UpdateStakes.
// Thankfully, we have deployed a helper contract BLSOperatorStateRetriever.sol whose function getCheckSignaturesIndices() can be used to fetch the indices given a block number.
//
// The 4 other fields nonSignerPubkeys, quorumApks, apkG2, and sigma, however, must be computed individually.
// apkG2 and sigma are just the aggregated signature and pubkeys of the operators who signed the task response (aggregated over all quorums, so individual signatures might be duplicated).
// quorumApks are the G1 aggregated pubkeys of the operators who signed the task response, but one per quorum, as opposed to apkG2 which is summed over all quorums.
// nonSignerPubkeys are the G1 pubkeys of the operators who did not sign the task response, but were opted into the quorum at the blocknumber at which the task was created.
// Upon sending a task onchain (or receiving a NewTaskCreated Event if the tasks were sent by an external task generator), the aggregator can get the list of all operators opted into each quorum at that
// block number by calling the getOperatorState() function of the BLSOperatorStateRetriever.sol contract.
type Aggregator struct {
	logger           logging.Logger
	serverIpPortAddr string
	avsWriter        chainio.AvsWriterer
	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	tasks                 map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask
	tasksMu               sync.RWMutex
	taskResponses         map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	txSignatures          map[uint32][]string
	taskResponsesMu       sync.RWMutex
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c *config.Config) (*Aggregator, error) {

	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.IncredibleSquaringRegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    avsName,
		PromMetricsIpPortAddress:   ":9090",
	}
	clients, err := clients.BuildAll(chainioConfig, c.EcdsaPrivateKey, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create sdk clients", "err", err)
		return nil, err
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(context.Background(), clients.AvsRegistryChainSubscriber, clients.AvsRegistryChainReader, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.Logger)

	return &Aggregator{
		logger:                c.Logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		avsWriter:             avsWriter,
		blsAggregationService: blsAggregationService,
		tasks:                 make(map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses:         make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse),
		txSignatures: 		   make(map[uint32][]string),
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Infof("Starting aggregator.")
	agg.logger.Infof("Starting aggregator rpc server.")
	go agg.startServer(ctx)

	// TODO(soubhik): refactor task generation/sending into a separate function that we can run as goroutine
	ticker := time.NewTicker(15 * time.Second)
	agg.logger.Infof("Aggregator set to recive tasks")
	defer ticker.Stop()
	//taskNum := int64(0)

	// ticker doesn't tick immediately, so we send the first task here
	// see https://github.com/golang/go/issues/17601

	// Use the random hex string in your function

	//agg.sendNewTask("0x58a2551789add523319ba3fc996904ad6e9d3115444f9a321b6656cc88aa03dd", "58b82b0b3281257ee1e9c6ea71f87362ecbb82f97a66fea12fe9ceca8c543539", gethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))
	for {
		select {
		case <-ctx.Done():
			return nil
		case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
			agg.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			agg.sendAggregatedResponseToContract(blsAggServiceResp)
			/*	case <-ticker.C:
				err := agg.sendNewTask("0x58a2551789add523319ba3fc996904ad6e9d3115444f9a321b6656cc88aa03dd", "58b82b0b3281257ee1e9c6ea71f87362ecbb82f97a66fea12fe9ceca8c543539", gethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))

				taskNum++
				if err != nil {
					// we log the errors inside sendNewTask() so here we just continue to the next task
					continue
				} */
		}
	}
}

func (agg *Aggregator) sendAggregatedResponseToContract(blsAggServiceResp blsagg.BlsAggregationServiceResponse) {
	// Check if blsAggServiceResp contains an error
	if blsAggServiceResp.Err != nil {
		agg.logger.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		// Panic to help with debugging (fail fast), but avoid panicking in production
		panic(blsAggServiceResp.Err)
	}

	nonSignerPubkeys := make([]cstaskmanager.BN254G1Point, len(blsAggServiceResp.NonSignersPubkeysG1))
	for i, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys[i] = core.ConvertToBN254G1Point(nonSignerPubkey)
	}
	quorumApks := make([]cstaskmanager.BN254G1Point, len(blsAggServiceResp.QuorumApksG1))
	for i, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks[i] = core.ConvertToBN254G1Point(quorumApk)
	}
	nonSignerStakesAndSignature := cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	agg.logger.Info("Threshold reached. Sending aggregated response onchain.",
		"taskIndex", blsAggServiceResp.TaskIndex,
	)

	// Retrieve task and taskResponse
	agg.tasksMu.RLock()
	task := agg.tasks[blsAggServiceResp.TaskIndex]
	agg.tasksMu.RUnlock()
	agg.taskResponsesMu.RLock()
	taskResponse := agg.taskResponses[blsAggServiceResp.TaskIndex][blsAggServiceResp.TaskResponseDigest]
	agg.taskResponsesMu.RUnlock()

	// Retrieve txSignatures2
	txSignatures2 := agg.txSignatures[blsAggServiceResp.TaskIndex]
	agg.logger.Info("Vreme za agregaciju", "txSignatures", txSignatures2)

	// Prepare JSON body for POST request
	requestBody, err := json.Marshal(map[string]interface{}{
		"signed_txs": txSignatures2,
	})
	if err != nil {
		agg.logger.Error("Failed to marshal request body", "error", err)
		return // Skip further execution if marshaling fails
	}

	// Make POST request to broadcast endpoint
	resp, err := http.Post("http://172.21.144.1:48594/broadcast", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		agg.logger.Error("Failed to make POST request to broadcast endpoint", "error", err)
		return // Skip further execution if POST request fails
	}
	defer resp.Body.Close()

	// Log response
	var responseObj map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseObj); err != nil {
		agg.logger.Error("Failed to decode response body", "error", err)
		return // Skip further execution if decoding response fails
	}

	// Log the response
	agg.logger.Info("Response to broadcast", "object", responseObj)

	// Proceed with the function, skipping errors and logging them
	_, err = agg.avsWriter.SendAggregatedResponse(context.Background(), task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		agg.logger.Error("Aggregator failed to respond to task", "err", err)
	}
}


// sendNewTask sends a new task to the task manager contract, and updates the Task dict struct
// with the information of operators opted into quorum 0 at the block of task creation.
func (agg *Aggregator) sendNewTask(btcTxHash string, signedMessage string, mintTo gethcommon.Address) error {
	agg.logger.Info("Aggregator sending new Mint task", "createNewTask", btcTxHash)
	// Send number to square to the task manager contract
	newTask, taskIndex, err := agg.avsWriter.SendNewTaskCheckTransaction(context.Background(), btcTxHash, signedMessage, mintTo, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS)
	if err != nil {
		agg.logger.Error("Aggregator failed to send check transaction task", "err", err)
		return err
	}
	agg.logger.Info("Uspelo", "createNewTask", btcTxHash)

	//agg.tasksMu.Lock()
	//agg.tasks[taskIndex] = newTask
	//agg.tasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range newTask.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}

	agg.blsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, quorumNums, quorumThresholdPercentages, taskTimeToExpiry)

	return nil
}

func (agg *Aggregator) sendNewTaskBurn(btcDestinationAddress string, burnAmount *big.Int) error {
	agg.logger.Info("Aggregator sending new Mint task", "createNewTask", btcDestinationAddress)
	// Send number to square to the task manager contract
	newTask, taskIndex, err := agg.avsWriter.SendNewTaskBurnBtc(context.Background(), btcDestinationAddress, burnAmount, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS)
	if err != nil {
		agg.logger.Error("Aggregator failed to send check transaction task", "err", err)
		return err
	}
	agg.logger.Info("Uspelo", "createNewTask", btcDestinationAddress)

	agg.tasksMu.Lock()
	agg.tasks[taskIndex] = newTask
	agg.tasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range newTask.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}
	agg.blsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, quorumNums, quorumThresholdPercentages, taskTimeToExpiry)
	return nil
}

func (agg *Aggregator) sendTaskHandler(task *cstaskmanager.IIncredibleSquaringTaskManagerTask) {

}

//helper functions
