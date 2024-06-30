const bitcoin = require('bitcoinjs-lib');
const ECPairFactory = require('ecpair').default;
const tinysecp = require('tiny-secp256k1');

const ECPair = ECPairFactory(tinysecp);

function generateLegacyAddress(privateKeyHex) {
    // Decode the private key
    const keyPair = ECPair.fromPrivateKey(Buffer.from(privateKeyHex, 'hex'), { network: bitcoin.networks.regtest });

    // Get the public key
    const { address } = bitcoin.payments.p2pkh({
        pubkey: keyPair.publicKey,
        network: bitcoin.networks.regtest,
    });

    // Get the address script
    const { output } = bitcoin.payments.p2pkh({
        pubkey: keyPair.publicKey,
        network: bitcoin.networks.regtest,
    });

    console.log('Legacy Address:', address);
    console.log('Address Script Hex:', output.toString('hex'));

    return { address, addressScriptHex: output.toString('hex') };
}

// Example private key (make sure to use a valid private key)
const privateKeyHex = '1000000000000000000000000000000000000000000000000000000000000001';

generateLegacyAddress(privateKeyHex);
