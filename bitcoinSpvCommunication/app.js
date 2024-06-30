const bitcoin = require("bitcoinjs-lib");
const Client = require("bitcoin-core");
const { ECPairFactory } = require("ecpair");
const tinysecp = require("tiny-secp256k1");
const ECPair = ECPairFactory(tinysecp);

// START COMMAND LINE ARGUMENTS
const privateKey = process.argv[2];
const destinationAddress = process.argv[3];
const multisigAddress = process.argv[4];
const amountToSend = parseFloat(process.argv[5]) / 1000;
const redeemScript = process.argv[6];
// END COMMAND LINE ARGUMENTS

const client = new Client({
    network: "regtest",
    username: "yourusername",
    password: "yourpassword",
    host: "192.168.2.156",
    port: 18443,
});

const NETWORK = bitcoin.networks.regtest;
async function sendTransaction() {
    const fee = 0.0001;
    const keyPair = ECPair.fromPrivateKey(Buffer.from(privateKey, "hex"), {
        network: NETWORK,
    });
    const psbt = new bitcoin.Psbt({ network: NETWORK });
    const descriptor = `addr(${multisigAddress})`;
    const scanResult = await client.command("scantxoutset", "start", [
        { desc: descriptor },
    ]);
    let total = 0;
    for (let utxo of scanResult.unspents) {
        const rawTx = await client.getRawTransaction(utxo.txid, false);
        total += utxo.amount;
        psbt.addInput({
            hash: utxo.txid,
            index: utxo.vout,
            nonWitnessUtxo: Buffer.from(rawTx, "hex"),
            redeemScript: Buffer.from(redeemScript, "hex"),
        });
    }
    psbt.addOutput({
        address: destinationAddress,
        value: Math.floor(amountToSend * 1e8), // Amount to send in satoshis
    });
    psbt.addOutput({
        address: multisigAddress,
        value: Math.floor((total - amountToSend - fee) * 1e8), // Change amount in satoshis
    });

    for (let j = 0; j < scanResult.unspents.length; j++)
        psbt.signInput(j, keyPair);
    const rawSignedTx = psbt.toBase64();
    console.log(rawSignedTx);
}
sendTransaction();
