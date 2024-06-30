const bitcoin = require('bitcoinjs-lib');

// Redeem script provided
const redeemScriptHex = '522102347d79020cf8914031ed69aae2dd7f6e6ce7e036d2976e50c3e3c412165df74621037aaa7852ba48c949c6e7a98263999c60ab0f2dde7031eaa080245b8bb250e2852102e9a4f0b9434ad0f8f87735aba91b852346558c7c8aed6bcd50f3ef5986c347bc53ae';
const redeemScript = Buffer.from(redeemScriptHex, 'hex');

// Compute the hash160 of the redeem script
const hash160 = bitcoin.crypto.hash160(redeemScript);

console.log('Redeem Script Hash:', hash160.toString('hex'));

// Build the scriptPubKey for the P2SH address
const scriptPubKey = bitcoin.payments.p2sh({ redeem: { output: redeemScript } }).output;
console.log('scriptPubKey:', scriptPubKey.toString('hex'));

// Get the P2SH address
const address = bitcoin.payments.p2sh({ redeem: { output: redeemScript }, network: bitcoin.networks.regtest }).address;
console.log('P2SH Address:', address);
