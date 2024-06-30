// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/src/libraries/BN254.sol";

interface IIncredibleSquaringTaskManager {
    // EVENTS
    event NewMintTaskCreated(uint32 indexed taskIndex, Task task);
    event NewBurnTaskCreated(uint32 indexed taskIndex, Task task);

    event MintTaskResponded(
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );

    event BurnTaskResponded(
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );

    event MintTaskCompleted(uint32 indexed taskIndex);
    event BurnTaskCompleted(uint32 indexed taskIndex);

    event TaskChallengedSuccessfully(
        uint32 indexed taskIndex,
        address indexed challenger
    );

    event TaskChallengedUnsuccessfully(
        uint32 indexed taskIndex,
        address indexed challenger
    );

    // STRUCTS
    struct Task {
        string btcTxHash;
        string signedMessage; //this is the signed hash of the mintTo address (address to which user wants their GBTC to be minted to)
        address mintTo;
        address burnFrom;
        string btcDestinationAddress;
        uint32 taskCreatedBlock;
        uint256 burnAmount;
        bool isBurnTask;
        // task submitter decides on the criteria for a task to be completed
        // note that this does not mean the task was "correctly" answered (i.e. the number was squared correctly)
        //      this is for the challenge logic to verify
        // task is completed (and contract will accept its TaskResponse) when each quorumNumbers specified here
        // are signed by at least quorumThresholdPercentage of the operators
        // note that we set the quorumThresholdPercentage to be the same for all quorumNumbers, but this could be changed
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
    }


    // Task response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct TaskResponse {
        // Can be obtained by the operator from the event NewTaskCreated.
        uint32 referenceTaskIndex;
        // This is just the response that the operator has to compute by itself.

        uint256 amount;

    }

    // Extra information related to taskResponse, which is filled inside the contract.
    // It thus cannot be signed by operators, so we keep it in a separate struct than TaskResponse
    // This metadata is needed by the challenger, so we emit it in the TaskResponded event
    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    // FUNCTIONS
    // NOTE: this function creates new task.
    function createNewTask(
        string calldata btcTxHash,
        string calldata signedMessage,
        address mintTo,
        uint256 burnAmount,
        string calldata btcDestinationAddress,
        bool isBurnTask,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external;

    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view returns (uint32);

    // // NOTE: this function raises challenge to existing tasks.
    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external;

    /// @notice Returns the TASK_RESPONSE_WINDOW_BLOCK
    function getTaskResponseWindowBlock() external view returns (uint32);
}
