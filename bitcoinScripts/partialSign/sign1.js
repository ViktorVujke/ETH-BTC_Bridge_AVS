const bitcoin = require('bitcoinjs-lib');
const Client = require('bitcoin-core');
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');

const ECPair = ECPairFactory(tinysecp);

const client = new Client({
    network: 'regtest',
    username: 'yourusername',
    password: 'yourpassword',
    host: 'localhost',
    port: 18443
});

const NETWORK = bitcoin.networks.regtest;

const { address, redeemScriptHex } = require("../dobar/multisig.json");
const privateKey = '1000000000000000000000000000000000000000000000000000000000000001'

async function sendTransaction() {
    const destinationAddress = 'mqYT9upmDU7WGVXWk3DKcMxGZCYiMGEhGg'; // Zitomir, hardcodovano
    const amountToSend = 0.1;
    const fee = 0.0001;

    const keyPair = ECPair.fromPrivateKey(Buffer.from(privateKey, 'hex'), { network: NETWORK });

    const psbt = new bitcoin.Psbt({ network: NETWORK });

    const descriptor = `addr(${address})`;
    const scanResult = await client.command('scantxoutset', 'start', [{ desc: descriptor }]);
    let total = 0;
    for (let utxo of scanResult.unspents) {
        const rawTx = await client.getRawTransaction(utxo.txid, false);
        total += utxo.amount;
        psbt.addInput({
            hash: utxo.txid,
            index: utxo.vout,
            nonWitnessUtxo: Buffer.from(rawTx, 'hex'),
            redeemScript: Buffer.from(redeemScriptHex, 'hex')
        });
    }

    console.log({
        address: destinationAddress,
        value: Math.floor(amountToSend * 1e8) // Amount to send in satoshis
    })
    psbt.addOutput({
        address: destinationAddress,
        value: Math.floor(amountToSend * 1e8) // Amount to send in satoshis
    });
    console.log({
        address,
        value: Math.floor((total - amountToSend - fee) * 1e8) // Change amount in satoshis
    })
    psbt.addOutput({
        address,
        value: Math.floor((total - amountToSend - fee) * 1e8) // Change amount in satoshis
    });
    console.log("DASd")

    for (let j = 0; j < scanResult.unspents.length; j++)
        psbt.signInput(j, keyPair);
    const rawSignedTx = psbt.toBase64();
    console.log("Potpisan base64")
    console.log(rawSignedTx)

}

sendTransaction();
