const bitcoin = require('bitcoinjs-lib');
const tinysecp = require('tiny-secp256k1');
const { ECPairFactory } = require('ecpair');
const ECPair = ECPairFactory(tinysecp);

const NETWORK = bitcoin.networks.regtest;

const rawPrivateKeys = [
    '2000000000000000000000000000000000000000000000000000000000000001',
    '2000000000000000000000000000000000000000000000000000000000000002',
    '2000000000000000000000000000000000000000000000000000000000000003',
    '2000000000000000000000000000000000000000000000000000000000000004'
];

const wifKeys = rawPrivateKeys.map(hex => {
    const keyPair = ECPair.fromPrivateKey(Buffer.from(hex, 'hex'));
    return keyPair.toWIF();
});

console.log(wifKeys);
