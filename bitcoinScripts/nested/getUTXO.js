const Client = require('bitcoin-core');
const fs = require('fs');

const client = new Client({
    network: 'regtest',
    username: 'yourusername',
    password: 'yourpassword',
    host: 'localhost',
    port: 18443,
});

const { address } = require("./multisig.json");

async function getAddressBalance(address) {
    try {
        const descriptor = `addr(${address})`;
        const scanResult = await client.command('scantxoutset', 'start', [{ desc: descriptor }]);

        if (scanResult.success) {
            console.log(scanResult)
            const UTXO = {
                txid: scanResult.unspents[0].txid,
                vout: scanResult.unspents[0].vout,
                scriptPubKey: scanResult.unspents[0].scriptPubKey,
                desc: scanResult.unspents[0].desc,
                amount: scanResult.unspents[0].amount,
                coinbase: scanResult.unspents[0].coinbase,
                height: scanResult.unspents[0].height,
                rawTx: "INSERT rawTx here"
            }
            fs.writeFile('UTXO.json', JSON.stringify(UTXO, null, 2), (err) => {
                if (err) throw err;
                console.log('Result saved to UTXO.json');
            });
            const balance = scanResult.total_amount;
            console.log(`Balance for address ${address}: ${balance} BTC`);
        } else {
            console.log(`No UTXOs found for address ${address}`);
        }
    } catch (error) {
        console.error('Error fetching address balance:', error);
    }
}

getAddressBalance(address);
