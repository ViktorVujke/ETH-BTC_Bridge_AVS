const bitcoin = require('bitcoinjs-lib');
const bitcoinRPC = require('bitcoin-rpc-promise');

const rpcConfig = {
    protocol: 'http',
    user: 'yourusername',
    pass: 'yourpassword',
    host: 'localhost',
    port: 18443,
};

const multisigScriptHex = '01024104347d79020cf8914031ed69aae2dd7f6e6ce7e036d2976e50c3e3c412165df746aef6d931803bdf2ed2549a87102cb12d0e4159355490b2cdec5dcea6e9cb1a2e41047aaa7852ba48c949c6e7a98263999c60ab0f2dde7031eaa080245b8bb250e2856b011c79811b5f9d356060560792918f48ed6861da1994ec6c2e5022a83deeef4104e9a4f0b9434ad0f8f87735aba91b852346558c7c8aed6bcd50f3ef5986c347bcffbe6dad3f638180bff5125615f451ee29369274a791235690122d386b23ade00103ae';
const multisigAddress = '2MzTHHunEeoiqPAsSe7yH2uJFcMW1EhDr9h';
const receivingAddress = 'bcrt1qq0vmxznvjlvaw2x394sxrcwnx4v5se95ns7q8u';
const privateKeys = [
    '1000000000000000000000000000000000000000000000000000000000000001',
    '1000000000000000000000000000000000000000000000000000000000000002',
    '1000000000000000000000000000000000000000000000000000000000000003'
];

(async function () {
    try {
        const client = new bitcoinRPC.Client(rpcConfig);

        // Step 1: Get UTXOs
        const utxos = await client.listUnspent(1, 9999999, [multisigAddress]);
        
        // Step 2: Create transaction builder
        const txb = new bitcoin.TransactionBuilder(bitcoin.networks.regtest);

        // Step 3: Add inputs from UTXOs
        for (const utxo of utxos) {
            txb.addInput(utxo.txid, utxo.vout);
        }

        // Step 4: Add output for destination address
        txb.addOutput(receivingAddress, 100000); // Amount in satoshis

        // Step 5: Sign transaction with private keys
        for (let i = 0; i < privateKeys.length; i++) {
            const keyPair = bitcoin.ECPair.fromPrivateKey(Buffer.from(privateKeys[i], 'hex'));
            for (let j = 0; j < utxos.length; j++) {
                if (utxos[j].address === multisigAddress) {
                    txb.sign({
                        prevOutScriptType: 'p2sh-p2wsh',
                        vin: j,
                        keyPair,
                        redeemScript: Buffer.from(multisigScriptHex, 'hex'),
                    });
                }
            }
        }

        // Step 6: Build and broadcast the transaction
        const tx = txb.build();
        const txHex = tx.toHex();
        const txid = await client.sendRawTransaction(txHex);
        console.log('Transaction ID:', txid);
    } catch (err) {
        console.error('Error:', err);
    }
})();
