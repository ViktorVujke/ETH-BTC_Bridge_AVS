const bitcoin = require('bitcoinjs-lib');
const Client = require('bitcoin-core');
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');

const ECPair = ECPairFactory(tinysecp);

const testId = 0;
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

const { txid, vout, scriptPubKey, amount, rawTx } = require("./UTXO.json");
const { address, redeemScriptHex, privateKeys } = require("./multisig.json");

async function sendTransaction() {
    const utxo = {
        txid, vout, scriptPubKey, amount, redeemScript: Buffer.from(redeemScriptHex, 'hex')
    };

    const destinationAddress = 'mqYT9upmDU7WGVXWk3DKcMxGZCYiMGEhGg'; // Zitomir, hardcodovano
    const amountToSend = 1;
    const fee = 0.0001;

    const keyPairs = privateKeys.map(key => ECPair.fromPrivateKey(Buffer.from(key.hex, 'hex'), { network: NETWORK }));

    const psbt = new bitcoin.Psbt({ network: NETWORK });

    console.log({
        hash: utxo.txid,
        index: utxo.vout,
        nonWitnessUtxo: Buffer.from(rawTx, 'hex'),
        redeemScript: utxo.redeemScript
    })
    psbt.addInput({
        hash: utxo.txid,
        index: utxo.vout,
        nonWitnessUtxo: Buffer.from(rawTx, 'hex'),
        redeemScript: utxo.redeemScript
    });

    psbt.addOutput({
        address: destinationAddress,
        value: Math.floor(amountToSend * 1e8) // Amount to send in satoshis
    });
    psbt.addOutput({
        address,
        value: Math.floor((utxo.amount - amountToSend - fee) * 1e8) // Change amount in satoshis
    });

    const signerNumber = 4;
    for (let i = 0; i < 4; i++) {
        const keyPair = keyPairs[i];
        psbt.signInput(0, keyPair);
    }
    console.log(psbt.data.inputs)
    psbt.finalizeAllInputs();

    const txHex = psbt.extractTransaction().toHex();

    try {
        console.log(txHex);
        //const txid = await client.sendRawTransaction(txHex);
        console.log(`Transaction broadcasted with txid: ${txid}`);
        if (testId == 0) {
            console.log(`If everything is ok, it should not be visible/fail after mining 10 blocks`);
        }
        else{
            console.log(`If everything is ok, it should be visible/succeed after mining 10 blocks`);
        }
    } catch (error) {
        console.error('Error broadcasting transaction:', error);
    }
}

sendTransaction();
