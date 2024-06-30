const bitcoinMessage = require('bitcoinjs-message');

const checkSignature = (publicKey, messageToSign, signature) => {
    try {
        const isValidRaw = bitcoinMessage.verify(messageToSign, publicKey, signature);
        if (isValidRaw)
            return true;

        const buffer = Buffer.from(signature, 'hex');
        const base64String = buffer.toString('base64');
        return bitcoinMessage.verify(messageToSign, publicKey, base64String);
    }
    catch (e) {
        try {
            const buffer = Buffer.from(signature, 'hex');
            const base64String = buffer.toString('base64');
            return bitcoinMessage.verify(messageToSign, publicKey, base64String);
        }
        catch (e) {
            return false;
        }
    }

}

module.exports = { checkSignature }