const bitcoin = require('bitcoinjs-lib');
const Client = require('bitcoin-core');
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');

const ECPair = ECPairFactory(tinysecp);
// Configure the Bitcoin Core client
const client = new Client({
    network: 'regtest', // Use 'testnet' or 'mainnet' for other networks
    username: 'yourusername',
    password: 'yourpassword',
    host: 'localhost',
    port: 18443 // Default port for regtest, adjust if necessary
});

const NETWORK = bitcoin.networks.regtest; // Use 'bitcoin.networks.testnet' for testnet

async function sendTransaction() {
    const utxo = {
        txid: '72748ab9bd2be6bc1e0f2bd1d5087d6104e9286467d0b850778f67ef165a3cd4',
        vout: 0,
        scriptPubKey: 'a9141a57b449dd1fb3b4221e5ff7ac9663d175c4173d87',
        amount: 5,
        redeemScript: Buffer.from('522102347d79020cf8914031ed69aae2dd7f6e6ce7e036d2976e50c3e3c412165df74621037aaa7852ba48c949c6e7a98263999c60ab0f2dde7031eaa080245b8bb250e2852102e9a4f0b9434ad0f8f87735aba91b852346558c7c8aed6bcd50f3ef5986c347bc53ae', 'hex')
    };

    const privateKeys = [
        '1000000000000000000000000000000000000000000000000000000000000001',
        '1000000000000000000000000000000000000000000000000000000000000002'
    ];

    const destinationAddress = 'mwejmQWu3TPscBMrsu6p8iDpgfUrADWfLg';
    const amountToSend = 1;
    const fee = 0.0001;

    const keyPairs = privateKeys.map(key => ECPair.fromPrivateKey(Buffer.from(key, 'hex'), { network: NETWORK }));

    console.log("K1")
    console.log(NETWORK);
    console.log("K2")
    const rawTx = "02000000013b3e2afab9e7e84dd8265826b6d650bc420008dbdd8607fde1759e709c7dd21c000000006a47304402204522876cfbb60817750cf438eef0ccc1c811d6c5ea24381e72fb38a8badd5d7002206f4299a1730470a84d422f7b073fae15c296e1d9da2ee5581b4d290e86e94e3f012103aedf1243e02e9a399106c4d36346f5e36202825e4ae7cb26977a86e4d281eedcfdffffff020065cd1d0000000017a9141a57b449dd1fb3b4221e5ff7ac9663d175c4173d8710b1e60e000000001976a9146e2c16ffdeb10ddcad50d32a28d08ca25fd0d81888aca6030000";
    const prevTx = bitcoin.Transaction.fromHex("02000000013b3e2afab9e7e84dd8265826b6d650bc420008dbdd8607fde1759e709c7dd21c000000006a47304402204522876cfbb60817750cf438eef0ccc1c811d6c5ea24381e72fb38a8badd5d7002206f4299a1730470a84d422f7b073fae15c296e1d9da2ee5581b4d290e86e94e3f012103aedf1243e02e9a399106c4d36346f5e36202825e4ae7cb26977a86e4d281eedcfdffffff020065cd1d0000000017a9141a57b449dd1fb3b4221e5ff7ac9663d175c4173d8710b1e60e000000001976a9146e2c16ffdeb10ddcad50d32a28d08ca25fd0d81888aca6030000");
    console.log(prevTx);
    const psbt = new bitcoin.Psbt({ network: NETWORK });

    // Add input (the UTXO to spend)
    psbt.addInput({
        hash: utxo.txid,
        index: utxo.vout,
        nonWitnessUtxo: Buffer.from(rawTx, 'hex'),
        redeemScript: utxo.redeemScript
    });

    // Add output (destination address and change address)
    psbt.addOutput({
        address: destinationAddress,
        value: Math.floor(amountToSend * 1e8) // Amount to send in satoshis
    });
    psbt.addOutput({
        address: "2MueWd516A6pSDHiVz4jjm3aTEN63Rvb7x7",
        value: Math.floor((utxo.amount - amountToSend - fee) * 1e8) // Change amount in satoshis
    });

    // Sign the transaction with each private key
    keyPairs.forEach((keyPair, index) => {
        console.log("KURAONJA")
        console.log(index)
        console.log(keyPair)
        psbt.signInput(0, keyPair);
    });

    // Validate all signatures and finalize the transaction
    //psbt.validateSignaturesOfAllInputs();
    psbt.finalizeAllInputs();

    // Get the raw transaction hex
    const txHex = psbt.extractTransaction().toHex();
    console.log(txHex)

    // Broadcast the transaction
    try {
        const txid = await client.sendRawTransaction(txHex);
        console.log(`Transaction broadcasted with txid: ${txid}`);
    } catch (error) {
        console.error('Error broadcasting transaction:', error);
    }
}


sendTransaction();
