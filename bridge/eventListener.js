const { ethers } = require('ethers');

// Replace with your provider URL (ensure it supports WebSockets)
const providerUrl = 'ws://localhost:8545'; // Ensure WebSocket is enabled on this URL

// Create a WebSocket provider
const provider = new ethers.WebSocketProvider(providerUrl);

// Replace with your contract address and ABI
const contractAddress = '0x9E545E3C0baAB3E08CdfD552C960A1050f373042';
const contractABI = [
    {
        "anonymous": false,
        "inputs": [
            { "indexed": true, "internalType": "uint32", "name": "taskIndex", "type": "uint32" },
            { "indexed": false, "internalType": "struct IIncredibleSquaringTaskManager.Task", "name": "task", "type": "tuple", "components": [
                { "internalType": "string", "name": "txHash", "type": "string" },
                { "internalType": "string", "name": "signedMessage", "type": "string" },
                { "internalType": "address", "name": "mintTo", "type": "address" },
                { "internalType": "uint32", "name": "taskCreatedBlock", "type": "uint32" },
                { "internalType": "uint256", "name": "amount", "type": "uint256" },
                { "internalType": "bool", "name": "isBurnTask", "type": "bool" },
                { "internalType": "bytes", "name": "quorumNumbers", "type": "bytes" },
                { "internalType": "uint32", "name": "quorumThresholdPercentage", "type": "uint32" }
            ]}
        ],
        "name": "NewMintTaskCreated",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            { "internalType": "bytes", "name": "quorumNumbers", "type": "bytes" }
        ],
        "name": "QuorumTask",
        "type": "event"
    }
];

// Create a contract instance
const contract = new ethers.Contract(contractAddress, contractABI, provider);

console.log('Listening for NewMintTaskCreated and QuorumTask events...');

// Listen for the NewMintTaskCreated event
contract.on('NewMintTaskCreated', (taskIndex, task, event) => {
    console.log('NewMintTaskCreated event detected:');
    console.log('Task Index:', taskIndex);
    console.log('Event:', event);
});

// Listen for the QuorumTask event
contract.on('QuorumTask', (quorumNumbers, event) => {
    console.log('QuorumTask event detected:');
    console.log('Quorum Numbers:', quorumNumbers);
    console.log('Event:', event);
});

// Keep the process running
process.stdin.resume();
