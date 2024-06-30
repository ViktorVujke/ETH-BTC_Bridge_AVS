const bitcoin = require('bitcoinjs-lib');
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');
const fs = require("fs")

const ECPair = ECPairFactory(tinysecp);

const privateKeyHex = '1010101010101010101010101010101010101010101010106010101000000011';

const privateKeyBuffer = Buffer.from(privateKeyHex, 'hex');

const keyPair = ECPair.fromPrivateKey(privateKeyBuffer, { network: bitcoin.networks.regtest });

const { address } = bitcoin.payments.p2pkh({ pubkey: keyPair.publicKey, network: bitcoin.networks.regtest });

console.log('Send transaction to:', address);


fs.writeFile('wallet.json', JSON.stringify({
    privateKey: privateKeyHex,
    address: address
}, null, 2), (err) => {
    if (err) throw err;
    console.log('Result saved to wallet.json');
});