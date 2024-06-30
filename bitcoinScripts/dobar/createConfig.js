const fs = require('fs');

const generatePrivateKey = (index) => {
    const hexIndex = index.toString(16).padStart(64, '0');
    return hexIndex;
};

const main = (total, required) => {
    if (isNaN(total) || isNaN(required) || total < 1 || required < 1 || required > total) {
        console.error('Invalid arguments. Make sure both arguments are positive integers and required <= total.');
        process.exit(1);
    }

    const privateKeys = [];
    for (let i = 1; i <= total; i++) {
        privateKeys.push(generatePrivateKey(i));
    }

    const config = {
        privateKeys,
        required
    };

    const fileName = `config${required}of${total}.json`;
    fs.writeFile(fileName, JSON.stringify(config, null, 2), (err) => {
        if (err) throw err;
        console.log(`Config file saved as ${fileName}`);
    });
};

main(14, 5);
