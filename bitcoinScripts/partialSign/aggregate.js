const bitcoin = require('bitcoinjs-lib');



// Function to combine and finalize multiple PSBTs
function combineAndFinalizeMultisigPsbt(signedPsbtBase64List) {
    const psbts = signedPsbtBase64List.map(psbtBase64 => bitcoin.Psbt.fromBase64(psbtBase64, { network: bitcoin.networks.regtest }));
    
    // Combine partially signed PSBTs
    const combinedPsbt = psbts[0];
    for (let i = 1; i < psbts.length; i++) {
        combinedPsbt.combine(psbts[i]);
    }

    // Finalize all inputs
    combinedPsbt.finalizeAllInputs();

    // Extract the raw transaction hex
    const txHex = combinedPsbt.extractTransaction().toHex();
    console.log('Final transaction hex:', txHex);
    
    return txHex;
}

// Example usage
let signedPsbtBase64List = [
    'cHNidP8BAHUCAAAAATwdx4+gRoo0jKGWvvwxL8TqV+hzfaC/fr0I6Zgl/LgLAQAAAAD/////AoCWmAAAAAAAGXapFG35w7w/OUip7qujn8rTS9bHcD+wiKy/kDABAAAAABepFHx7WNqjcfsD1cOgobMu+Cf7RL4xhwAAAAAAAQD9UAECAAAAAUUj6tT9tqFi83MYMIg6yIbTbe0V6qDxuAARB7bbWOGYAQAAANsASDBFAiEA+3uC3Icio8AFtCn3zFUdA8iQ02nKNQGKrTz0n1UR+hUCIE3o/NVjHfPogNniQv8JlLkLWfTHQDhSlaGYblApxvgLAUgwRQIhAN+L/1DpuwPP+Do1qIsbv6Lx3IpfK58LuAY+Id+oLiH9AiAmF2h3UUM1t/j2Wa2ymfQQjjKTNuhlT58yrnfGeoHMeQFHUiECNH15Agz4kUAx7Wmq4t1/bmzn4DbSl25Qw+PEEhZd90YhA3qqeFK6SMlJxuepgmOZnGCrDy3ecDHqoIAkW4uyUOKFUq7/////AoDw+gIAAAAAGXapFG35w7w/OUip7qujn8rTS9bHcD+wiKxPTskBAAAAABepFHx7WNqjcfsD1cOgobMu+Cf7RL4xhwAAAAAiAgI0fXkCDPiRQDHtaari3X9ubOfgNtKXblDD48QSFl33RkgwRQIhANyDB6xDGNE76g/IgGu/lIhZePHMGzRC+QkZaLOl+x8hAiByAj/vM6cUmdImgBW9r9TXWUzM++j5Gd7RK/oqeRLPzwEBBEdSIQI0fXkCDPiRQDHtaari3X9ubOfgNtKXblDD48QSFl33RiEDeqp4UrpIyUnG56mCY5mcYKsPLd5wMeqggCRbi7JQ4oVSrgAAAA==',
    //'cHNidP8BAHUCAAAAATwdx4+gRoo0jKGWvvwxL8TqV+hzfaC/fr0I6Zgl/LgLAQAAAAD/////AoCWmAAAAAAAGXapFG35w7w/OUip7qujn8rTS9bHcD+wiKy/kDABAAAAABepFHx7WNqjcfsD1cOgobMu+Cf7RL4xhwAAAAAAAQD9UAECAAAAAUUj6tT9tqFi83MYMIg6yIbTbe0V6qDxuAARB7bbWOGYAQAAANsASDBFAiEA+3uC3Icio8AFtCn3zFUdA8iQ02nKNQGKrTz0n1UR+hUCIE3o/NVjHfPogNniQv8JlLkLWfTHQDhSlaGYblApxvgLAUgwRQIhAN+L/1DpuwPP+Do1qIsbv6Lx3IpfK58LuAY+Id+oLiH9AiAmF2h3UUM1t/j2Wa2ymfQQjjKTNuhlT58yrnfGeoHMeQFHUiECNH15Agz4kUAx7Wmq4t1/bmzn4DbSl25Qw+PEEhZd90YhA3qqeFK6SMlJxuepgmOZnGCrDy3ecDHqoIAkW4uyUOKFUq7/////AoDw+gIAAAAAGXapFG35w7w/OUip7qujn8rTS9bHcD+wiKxPTskBAAAAABepFHx7WNqjcfsD1cOgobMu+Cf7RL4xhwAAAAAiAgN6qnhSukjJScbnqYJjmZxgqw8t3nAx6qCAJFuLslDihUgwRQIhAOIg657uwDhhPfd2MI4oJmgstd8pZz8XQcr0JuoXttrUAiBidwDcN7WwLJrFDzT9Bk8J/rS9cjo+0nU1hB2CvfZI/AEBBEdSIQI0fXkCDPiRQDHtaari3X9ubOfgNtKXblDD48QSFl33RiEDeqp4UrpIyUnG56mCY5mcYKsPLd5wMeqggCRbi7JQ4oVSrgAAAA=='
];

/*signedPsbtBase64List=[
    'AQAAAAE8HcePoEaKNIyhlr78MS/E6lfoc32gv369COmYJfy4CwEAAABrSDBFAiEA81DFx2xj403bYoelFkHw+IKZSqmiTuAWTMrmc1eaUEYCID07oAoe/0Q70fIPjI7EjBpNzLDVgQZ+gv2GVSiVklN1ASECNH15Agz4kUAx7Wmq4t1/bmzn4DbSl25Qw+PEEhZd90b/////AoCWmAAAAAAAGXapFG35w7w/OUip7qujn8rTS9bHcD+wiKy/kDABAAAAABepFHx7WNqjcfsD1cOgobMu+Cf7RL4xhwAAAAA=',
    'AQAAAAE8HcePoEaKNIyhlr78MS/E6lfoc32gv369COmYJfy4CwEAAABqRzBEAiAW+h/rj7IMfbzeTj1BHtE5OpuEXaZ4udubZulJaWKQ2gIgFNbLzIVXBMLbgelThPtzGk/URwCyYlOBSjwkE7hNsi0BIQN6qnhSukjJScbnqYJjmZxgqw8t3nAx6qCAJFuLslDihf////8CgJaYAAAAAAAZdqkUbfnDvD85SKnuq6OfytNL1sdwP7CIrL+QMAEAAAAAF6kUfHtY2qNx+wPVw6Chsy74J/tEvjGHAAAAAA=='
]*/

const finalTxHex = combineAndFinalizeMultisigPsbt(signedPsbtBase64List);
//console.log(finalTxHex);

async function broadcastTransaction(txHex) {
    const Client = require('bitcoin-core');

    const client = new Client({
        network: 'regtest',
        username: 'yourusername',
        password: 'yourpassword',
        host: 'localhost',
        port: 18443
    });

    try {
        const txid = await client.sendRawTransaction(txHex);
        console.log('Transaction broadcasted, txid:', txid);
    } catch (error) {
        console.error('Error broadcasting transaction:', error);
    }
}

// Example usage
broadcastTransaction(finalTxHex);

