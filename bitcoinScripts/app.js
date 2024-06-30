const Client = require('bitcoin-core');
const client = new Client({
  network: 'regtest',
  host: '127.0.0.1',
  username: 'yourusername',
  password: 'yourpassword',
  port: 18443,
  wallet: 'w2'
});

async function mineBlocks(numberOfBlocks) {
    try {
      await client.generate(numberOfBlocks);
      console.log(`Successfully mined ${numberOfBlocks} blocks.`);
    } catch (error) {
      console.error('Error mining blocks:', error.message);
    }
  }
  
  // Mine 100 blocks to mature the coinbase transactions
  mineBlocks(100);
  