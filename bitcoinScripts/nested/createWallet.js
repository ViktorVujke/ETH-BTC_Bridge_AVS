// Generise skriptu i njen hes

const bitcoin = require('bitcoinjs-lib');
const network = bitcoin.networks.regtest;
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');
const fs = require('fs');

const ECPair = ECPairFactory(tinysecp);

const { privateKeys, required } = require("./config3of4.json")
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
// script.push(pushNumberToStack(required));
for (const { hex, weight } of privateKeys) {
    const keyPair = ECPair.fromPrivateKey(Buffer.from(hex, 'hex'), { network });
    const pubkey = keyPair.publicKey;
    script.push(pubkey);
    script.push(bitcoin.opcodes.OP_CHECKSIGVERIFY);
    script.push(bitcoin.opcodes.OP_IF);
    script.push(pushNumberToStack(weight));
    script.push(bitcoin.opcodes.OP_ELSE);
    script.push(bitcoin.opcodes.OP_0);
    script.push(bitcoin.opcodes.OP_ENDIF);
    script.push(bitcoin.opcodes.OP_TOALTSTACK);
}

for (const _ of privateKeys) {
    script.push(bitcoin.opcodes.OP_FROMALTSTACK);
}

for (const _ of privateKeys) {
    script.push(bitcoin.opcodes.OP_ADD);
}

script.push(pushNumberToStack(required))
script.push(bitcoin.opcodes.OP_GREATERTHANOREQUAL)
script.push(bitcoin.opcodes.OP_VERIFY)

console.log(script);
/*script = [
    bitcoin.opcodes.OP_1
]*/

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