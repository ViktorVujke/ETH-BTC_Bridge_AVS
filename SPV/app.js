const express = require("express")
const { txs, saveTxs } = require("./spv");
const { checkSignature } = require("./util/checkSignature");
const bitcoin = require("bitcoinjs-lib");
const Client = require("bitcoin-core");
const { ECPairFactory } = require("ecpair");
const tinysecp = require("tiny-secp256k1");
const ECPair = ECPairFactory(tinysecp);

const app = express();

app.use(express.json())
app.post("/verify-payment", async (req, res) => {
    const { txid, eth_address, signed_eth_address } = req.body;

    console.log("Zahtevana verifikacija");
    console.log(req.body);

    const tx = txs[txid];
    if (!tx)
        return res.json({ ok: false })

    const valid = checkSignature(tx.payer, eth_address, signed_eth_address);
    if (!valid)
        return res.json({ ok: false })

    delete txs[txid];
    saveTxs();

    res.json({ ok: true, amount: Math.floor(tx.amount * 0.995), confirmations: tx.confirmations })
})

const PRIVATE_KEY = "1000000000100000000000000000000000000000000000000000000000000001";
const MULTISIG_ADDRESS = "2N4ZJY7R6FKZsc7SqKw4yZ5eHZo9BP3C7Go";
const REDEEM_SCRIPT = "5121033aa9e95a07bb841d61ab66dd4ee0ccc2f74baead5a6e0230f208f7794792671a2102b3c9eee67b10bb80410263902eefdf9a321488206dde15fda69f51ead395bb5652ae"

const client = new Client({
    network: "regtest",
    username: "yourusername",
    password: "yourpassword",
    host: "127.0.0.1",
    port: 18443,
});

const NETWORK = bitcoin.networks.regtest;
app.post("/sign-tx", async (req, res) => {
    try {
        // AMOUNT JE U SATOSHIJIMA
        const { destination_address, amount } = req.body;

        const fee = 10000;
        const keyPair = ECPair.fromPrivateKey(Buffer.from(PRIVATE_KEY, "hex"), {
            network: NETWORK,
        });
        const psbt = new bitcoin.Psbt({ network: NETWORK });
        const descriptor = `addr(${MULTISIG_ADDRESS})`;
        const scanResult = await client.command("scantxoutset", "start", [
            { desc: descriptor },
        ]);
        let total = 0;
        for (let utxo of scanResult.unspents) {
            const rawTx = await client.getRawTransaction(utxo.txid, false);
            total += utxo.amount;
            psbt.addInput({
                hash: utxo.txid,
                index: utxo.vout,
                nonWitnessUtxo: Buffer.from(rawTx, "hex"),
                redeemScript: Buffer.from(REDEEM_SCRIPT, "hex"),
            });
        }
        total = Math.floor(total * 1e8)
        psbt.addOutput({
            address: destination_address,
            value: Math.floor(amount * 0.995), // Amount to send in satoshis
        });
        psbt.addOutput({
            address: MULTISIG_ADDRESS,
            value: total - Math.floor(amount * 0.995) - fee, // Change amount in satoshis
        });

        for (let j = 0; j < scanResult.unspents.length; j++)
            psbt.signInput(j, keyPair);
        const rawSignedTx = psbt.toBase64();

        res.json({ ok: true, signed_tx: rawSignedTx })
    }
    catch (e) {
        res.json({ ok: false })
    }
})

function combineAndFinalizeMultisigPsbt(signedPsbtBase64List) {
    const psbts = signedPsbtBase64List.map((psbtBase64) =>
        bitcoin.Psbt.fromBase64(psbtBase64, {
            network: bitcoin.networks.regtest,
        })
    );

    const combinedPsbt = psbts[0];
    for (let i = 1; i < psbts.length; i++) {
        combinedPsbt.combine(psbts[i]);
    }

    combinedPsbt.finalizeAllInputs();

    const txHex = combinedPsbt.extractTransaction().toHex();
    console.log("Final transaction hex:", txHex);
    return txHex;
}

async function broadcastTransaction(txHex) {
    try {
        const txid = await client.sendRawTransaction(txHex);
        console.log("Transaction broadcasted, txid:", txid);
        return true;
    } catch (error) {
        console.error("Error broadcasting transaction:", error);
        return false;
    }
}


app.post("/broadcast", async (req, res) => {
    try {
        const { signed_txs } = req.body;
        const finalTxHex = combineAndFinalizeMultisigPsbt(signed_txs);
        const ok = await broadcastTransaction(finalTxHex);
        res.json({ ok })
    }
    catch (e) {
        res.json({ ok: false })
    }

})

app.listen("48594")