const bitcoin = require('bitcoinjs-lib');
const Client = require('bitcoin-core');
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');

const ECPair = ECPairFactory(tinysecp);

const testId = 1;
// 0 - treshold-1, fail
// 1 - treshold  , success
// 2 - treshold+1, fail

const client = new Client({
    network: 'regtest',
    username: 'yourusername',
    password: 'yourpassword',
    host: 'localhost',
    port: 18443
});

const NETWORK = bitcoin.networks.regtest;

const { address, redeemScriptHex, privateKeys, required } = require("./multisig.json");

async function sendTransaction() {
    const destinationAddress = 'mqYT9upmDU7WGVXWk3DKcMxGZCYiMGEhGg'; // Zitomir, hardcodovano
    const amountToSend = 0.6;
    const fee = 0.0001;

    const keyPairs = privateKeys.map(key => ECPair.fromPrivateKey(Buffer.from(key, 'hex'), { network: NETWORK }));

    const psbt = new bitcoin.Psbt({ network: NETWORK });

    const descriptor = `addr(${address})`;
    const scanResult = await client.command('scantxoutset', 'start', [{ desc: descriptor }]);
    let total = 0;
    for (let utxo of scanResult.unspents) {
        const rawTx = await client.getRawTransaction(utxo.txid, false);
        console.log(utxo)
        total += utxo.amount;
        psbt.addInput({
            hash: utxo.txid,
            index: utxo.vout,
            nonWitnessUtxo: Buffer.from(rawTx, 'hex'),
            redeemScript: Buffer.from(redeemScriptHex, 'hex')
        });
    }

    psbt.addOutput({
        address: destinationAddress,
        value: Math.floor(amountToSend * 1e8) // Amount to send in satoshis
    });
    psbt.addOutput({
        address,
        value: Math.floor((total - amountToSend - fee) * 1e8) // Change amount in satoshis
    });

    const signerNumber = required - 1 + testId;
    console.log(psbt);
    for (let j = 0; j < scanResult.unspents.length; j++)
        for (let i = 0; i < 1; i++) {
            const keyPair = keyPairs[i];
            console.log(keyPair.publicKey)
            psbt.signInput(j, keyPair);
        }
    psbt.finalizeAllInputs();

    const txHex = psbt.extractTransaction().toHex();

    try {
        const txid = await client.sendRawTransaction(txHex);
        console.log(`Transaction broadcasted with txid: ${txid}`);
        if (testId == 0) {
            console.log(`If everything is ok, it should not be visible/fail after mining 10 blocks`);
        }
        else {
            console.log(`If everything is ok, it should be visible/succeed after mining 10 blocks`);
        }
    } catch (error) {
        console.error('Error broadcasting transaction:', error);
    }
}

sendTransaction();
