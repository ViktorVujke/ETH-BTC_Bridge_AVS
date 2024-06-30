// Generise skriptu i njen hes

const bitcoin = require('bitcoinjs-lib');
const network = bitcoin.networks.regtest;
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');
const fs = require('fs');

const ECPair = ECPairFactory(tinysecp);

const { privateKeys, required } = require("./config1of2.json")
const total = privateKeys.length;

const pushNumberToStack = (number) => {
    const hexNumber = number.toString(16);

    const paddedHexNumber = hexNumber.length % 2 === 0 ? hexNumber : '0' + hexNumber;

    const script = Buffer.concat([
        Buffer.from([(paddedHexNumber.length / 2)]),
        Buffer.from(paddedHexNumber, 'hex')
    ]);

    const decodedScript = bitcoin.script.decompile(script);

    return decodedScript[0]
}

let script = []
script.push(pushNumberToStack(required));
for (const privateKey of privateKeys) {
    const keyPair = ECPair.fromPrivateKey(Buffer.from(privateKey, 'hex'), { network });
    const pubkey = keyPair.publicKey;
    script.push(pubkey);
}
script.push(pushNumberToStack(total))
script.push(bitcoin.opcodes.OP_CHECKMULTISIG)

console.log(script);

const redeemScript = bitcoin.script.compile(script);
// const redeemScriptHash = bitcoin.crypto.hash160(redeemScript);
const scriptPubKey = bitcoin.payments.p2sh({ redeem: { output: redeemScript } }).output;
const address = bitcoin.address.fromOutputScript(scriptPubKey, network);
const result = {
    redeemScriptHex: redeemScript.toString('hex'),
    scriptPubKey: scriptPubKey.toString('hex'),
    address: address,
    privateKeys,
    required
}

console.log(`Send a transaction to ${address}, then mine 10 blocks`)

fs.writeFile('multisig.json', JSON.stringify(result, null, 2), (err) => {
    if (err) throw err;
    console.log('Result saved to multisig.json');
});