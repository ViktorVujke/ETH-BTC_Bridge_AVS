const Client = require('bitcoin-core');

// Configure the Bitcoin Core client
const client = new Client({
    network: 'regtest',
    username: 'yourusername',
    password: 'yourpassword',
    host: 'localhost',
    port: 18443
});

// Function to get raw transaction info
async function getTransactionInfo(txhash) {
    try {
        // Get the raw transaction
        const rawTx = await client.getRawTransaction(txhash, true); // Set 'true' to decode the transaction
        console.log(rawTx.vin)
        // Loop through each input (vin) to find the previous transactions
        for (const input of rawTx.vin) {
            const prevTx = await client.getRawTransaction(input.txid, true);
            const payerOutput = prevTx.vout[input.vout];

            // Extract the addresses from the previous output
            const payerAddresses = payerOutput.scriptPubKey.address || payerOutput.scriptPubKey.addresses;

            console.log(`Payer Addresses for TXID ${input.txid}, VOUT ${input.vout}:`, payerAddresses);
        }
    } catch (error) {
        console.error('Error fetching transaction info:', error);
    }
}

// Replace 'your-txhash' with the actual transaction hash you want to query
const txhash = 'f9c63d65cedd76056318f9a47a1bc48f60bca9c783a165af4a507fb9bfd54329';
getTransactionInfo(txhash);
