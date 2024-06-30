const Client = require('bitcoin-core');

// Configure the Bitcoin Core client
const client = new Client({
  network: 'regtest', // Use 'testnet' or 'mainnet' for other networks
  username: 'yourusername',
  password: 'yourpassword',
  host: 'localhost',
  port: 18443, // Default port for regtest, adjust if necessary
  wallet: 'dada' // Specify the wallet filename
});

async function importDescriptor(privateKey) {
  try {
    const descriptor = `wpkh([your_fingerprint/84h/1h/0h/0/0]${privateKey})`;
    const importData = [
      {
        desc: descriptor,
        active: true,
        timestamp: 'now'
      }
    ];
    
    const result = await client.command('importdescriptors', importData);
    console.log(`Descriptor imported successfully`, result);
  } catch (error) {
    console.error('Error importing descriptor:', error);
  }
}

// Example private key
const privateKey = '1000000000000000000000000000000000000000000000000000000000000001'; // Replace with your actual private key

importDescriptor(privateKey);
