const Client = require('bitcoin-core');

// Configure the Bitcoin Core client
const client = new Client({
  network: 'regtest', // Use 'testnet' or 'mainnet' for other networks
  username: 'yourusername',
  password: 'yourpassword',
  host: 'localhost',
  port: 18443, // Default port for regtest, adjust if necessary
});

async function getAddressBalance(address) {
  try {
    // Use the addr() descriptor for the address
    const descriptor = `addr(${address})`;
    const scanResult = await client.command('scantxoutset', 'start', [{ desc: descriptor }]);

    if (scanResult.success) {
        console.log(scanResult);
      const balance = scanResult.total_amount;
      console.log(`Balance for address ${address}: ${balance} BTC`);
    } else {
      console.log(`No UTXOs found for address ${address}`);
    }
  } catch (error) {
    console.error('Error fetching address balance:', error);
  }
}

// Example address
const address = '2MueWd516A6pSDHiVz4jjm3aTEN63Rvb7x7'; // Replace with the address you want to check

getAddressBalance(address);
