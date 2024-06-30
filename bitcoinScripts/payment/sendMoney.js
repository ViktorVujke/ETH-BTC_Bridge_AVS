const Client = require('bitcoin-core');
const bitcoin = require('bitcoinjs-lib');
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');

const ECPair = ECPairFactory(tinysecp);

const client = new Client({
    network: 'regtest',
    username: 'yourusername',
    password: 'yourpassword',
    host: '127.0.0.1',
    port: 18443
});

const { address, privateKey } = require("./wallet.json");

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
                height: scanResult.unspents[0].height
            }


            const NETWORK = bitcoin.networks.regtest;

            const destinationAddress = '2N4ZJY7R6FKZsc7SqKw4yZ5eHZo9BP3C7Go'; // Zitomir, hardcodovano
            const amountToSend = 1;
            const fee = 0.0001;

            const psbt = new bitcoin.Psbt({ network: NETWORK });
            
            const rawTx = await client.getRawTransaction(UTXO.txid, false);

            psbt.addInput({
                hash: UTXO.txid,
                index: UTXO.vout,
                nonWitnessUtxo: Buffer.from(rawTx, 'hex') // I have txid in UTXO.txid. I want to get it (using async await before this)
            });

            psbt.addOutput({
                address: destinationAddress,
                value: Math.floor(amountToSend * 1e8) // Amount to send in satoshis
            });
            psbt.addOutput({
                address,
                value: Math.floor((UTXO.amount - amountToSend - fee) * 1e8) // Change amount in satoshis
            });

            psbt.signInput(0, ECPair.fromPrivateKey(Buffer.from(privateKey, 'hex'), { network: NETWORK }));

            psbt.finalizeAllInputs();

            const txHex = psbt.extractTransaction().toHex();

            try {
                const txid = await client.sendRawTransaction(txHex);
                console.log(`Transaction broadcasted with txid: ${txid}`);
            } catch (error) {
                console.error('Error broadcasting transaction:', error);
            }

            //const balance = scanResult.total_amount;
            //console.log(`Balance for address ${address}: ${balance} BTC`);
        } else {
            console.log(`No UTXOs found for address ${address}`);
        }
    } catch (error) {
        console.error('Error fetching address balance:', error);
    }
}

getAddressBalance(address);
