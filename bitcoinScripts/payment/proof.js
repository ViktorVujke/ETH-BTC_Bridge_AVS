const bitcoin = require('bitcoinjs-lib');
const { payments } = bitcoin;
const bs58check = require('bs58check');
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');
const bitcoinMessage = require('bitcoinjs-message');

const ECPair = ECPairFactory(tinysecp);

// Your private key in hex
const privateKeyHex = '1010101010101010101010101010101010101010101010106010101000000011';

// Create a key pair from the private key
const keyPair = ECPair.fromPrivateKey(Buffer.from(privateKeyHex, 'hex'), { network: bitcoin.networks.regtest });

// The message to sign
const messageToSign = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';

// Generate the address from the key pair
const { address } = payments.p2pkh({ pubkey: keyPair.publicKey, network: bitcoin.networks.regtest });

const privateKeyBuffer = keyPair.privateKey;
// Sign the message
const signature = bitcoinMessage.sign(messageToSign, privateKeyBuffer, keyPair.compressed);
console.log('Signed Message:', signature);

// Convert signature to base64
const signatureBase64 = signature.toString('base64');
const signatureHex = signature.toString('hex');

console.log('Signed Message:', signatureBase64);
console.log('Signed Message:', signatureHex);


//const isValid = bitcoinMessage.verify(messageToSign, "mnJH8EX8efT3ci2myZf9shKbBUaiFaMxmW", signatureBase64);
//console.log('Message:', messageToSign);
//console.log('Address:', address);
//console.log('Signature is valid:', isValid);


