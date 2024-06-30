const bitcoin = require('bitcoinjs-lib');
const Client = require('bitcoin-core');

// Configure the Bitcoin Core client
const client = new Client({
  network: 'regtest', // Use 'testnet' or 'mainnet' for other networks
  username: 'yourusername',
  password: 'yourpassword',
  host: 'localhost',
  port: 18443, // Default port for regtest, adjust if necessary
  wallet: 'Zitomir' // Specify the wallet filename, if necessary
});

const NETWORK = bitcoin.networks.regtest; // Use 'bitcoin.networks.testnet' for testnet

async function sendTransaction() {
  const utxo = {
    txid: 'cf855801396de8e8df291365ba4f826b39cf52e430ea770572f2f615f3ed9c2d',
    vout: 0,
    scriptPubKey: 'a914224fb288716c9a91dce44157deaa58adf80c95f687',
    amount: 4,
    redeemScript: Buffer.from('01021976a91457bc4527db05de09595a35e487a03488fb49585288ac1976a9146e2977a471d2751e3aa35f10c17322ece3380da788ac1976a914e731a232c77d79c0e02e68f99d67a25ac364087888ac0103ae', 'hex')
  };

  const privateKeys = [
    'your_private_key_hex_1',
    'your_private_key_hex_2'
  ];

  const destinationAddress = 'your_destination_address';
  const amountToSend = 1; // Amount in BTC
  const fee = 0.0001; // Fee in BTC

  const keyPairs = privateKeys.map(key => bitcoin.ECPair.fromPrivateKey(Buffer.from(key, 'hex'), { network: NETWORK }));

  // Create a transaction builder
  const txb = new bitcoin.TransactionBuilder(NETWORK);
  txb.addInput(utxo.txid, utxo.vout, null, Buffer.from(utxo.scriptPubKey, 'hex'));
  txb.addOutput(destinationAddress, Math.floor(amountToSend * 1e8));
  txb.addOutput('your_change_address', Math.floor((utxo.amount - amountToSend - fee) * 1e8));

  // Sign the transaction
  for (let i = 0; i < keyPairs.length; i++) {
    txb.sign({
      prevOutScriptType: 'p2sh',
      vin: 0,
      keyPair: keyPairs[i],
      redeemScript: utxo.redeemScript
    });
  }

  // Build the transaction
  const tx = txb.build();
  const txHex = tx.toHex();

  // Broadcast the transaction
  try {
    const txid = await client.sendRawTransaction(txHex);
    console.log(`Transaction broadcasted with txid: ${txid}`);
  } catch (error) {
    console.error('Error broadcasting transaction:', error);
  }
}

sendTransaction();
