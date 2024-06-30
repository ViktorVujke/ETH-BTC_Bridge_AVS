const Client = require("bitcoin-core");

// Initialize the Bitcoin RPC client
const client = new Client({
    network: "regtest", // or 'mainnet', 'testnet'
    host: "192.168.2.156",
    port: 18443,
    username: "yourusername",
    password: "yourpassword",
    timeout: 30000, // optional timeout, in milliseconds
});

async function fetchUTXOs(address) {
    const desc = `addr(${address})`;
    const request = [{ desc: desc }];

    try {
        console.log("Sending scantxoutset RPC call");
        const rawResponse = await client.command(
            "scantxoutset",
            "start",
            request
        );

        if (!rawResponse.success) {
            throw new Error("scantxoutset failed");
        }

        console.log("Fetched UTXOs:", rawResponse.unspents);
        return rawResponse.unspents;
    } catch (error) {
        console.error("Failed to perform scantxoutset RPC call:", error);
        throw error;
    }
}

// Example usage
(async () => {
    try {
        const address = "2N4bRYNoXtmhj1A1r4d9gz6jB3CdBQ1KjCv";
        const utxos = await fetchUTXOs(address);
        console.log("UTXOs:", utxos);
    } catch (error) {
        console.error("Error fetching UTXOs:", error);
    }
})();
