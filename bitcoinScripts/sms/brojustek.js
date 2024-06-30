const bitcoin = require('bitcoinjs-lib');

// Number to push onto the stack
const number = 2;
const hexNumber = number.toString(16); // Convert number to hexadecimal

// Ensure the hexadecimal string has even length
const paddedHexNumber = hexNumber.length % 2 === 0 ? hexNumber : '0' + hexNumber;

// Create the script to push the number onto the stack
const script = Buffer.concat([
  Buffer.from([(paddedHexNumber.length / 2)]), // Length of the data
  Buffer.from(paddedHexNumber, 'hex') // The data itself
]);

console.log('Hex Script:', script.toString('hex'));

// Decode the script using bitcoinjs-lib
const decodedScript = bitcoin.script.decompile(script);

console.log('Decoded Script:', decodedScript);
