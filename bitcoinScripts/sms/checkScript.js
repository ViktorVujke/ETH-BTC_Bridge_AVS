const bitcoin = require('bitcoinjs-lib');

// Provided hex script
//const hexScript = '01021976a91457bc4527db05de09595a35e487a03488fb49585288ac1976a9146e2977a471d2751e3aa35f10c17322ece3380da788ac1976a914e731a232c77d79c0e02e68f99d67a25ac364087888ac0103ae';
const hexScript = '522102347d79020cf8914031ed69aae2dd7f6e6ce7e036d2976e50c3e3c412165df74621037aaa7852ba48c949c6e7a98263999c60ab0f2dde7031eaa080245b8bb250e2852102e9a4f0b9434ad0f8f87735aba91b852346558c7c8aed6bcd50f3ef5986c347bc53ae'
//const hexScript = '01022102347d79020cf8914031ed69aae2dd7f6e6ce7e036d2976e50c3e3c412165df74621037aaa7852ba48c949c6e7a98263999c60ab0f2dde7031eaa080245b8bb250e2852102e9a4f0b9434ad0f8f87735aba91b852346558c7c8aed6bcd50f3ef5986c347bc0103ae'
// Decode the hex script
const script = bitcoin.script.decompile(Buffer.from(hexScript, 'hex'));

// Expected script structure
const expectedScript = [
  bitcoin.opcodes.OP_2,
  Buffer.from('02347d79020cf8914031ed69aae2dd7f6e6ce7e036d2976e50c3e3c412165df746', 'hex'),
  Buffer.from('037aaa7852ba48c949c6e7a98263999c60ab0f2dde7031eaa080245b8bb250e285', 'hex'),
  Buffer.from('02e9a4f0b9434ad0f8f87735aba91b852346558c7c8aed6bcd50f3ef5986c347bc', 'hex'),
  bitcoin.opcodes.OP_3,
  bitcoin.opcodes.OP_CHECKMULTISIG
];

// Function to check if two scripts are equal
function scriptsEqual(script1, script2) {
  if (script1.length !== script2.length) {
    return false;
  }
  for (let i = 0; i < script1.length; i++) {
    if (Buffer.isBuffer(script1[i]) && Buffer.isBuffer(script2[i])) {
      if (!script1[i].equals(script2[i])) {
        console.log(script1[i].toString('hex'))
        return false;
      }
    } else if (script1[i] !== script2[i]) {
      return false;
    }
  }
  return true;
}

// Check if the decoded script matches the expected script
const isMatch = scriptsEqual(script, expectedScript);

console.log('Decoded Script:', script);
console.log('Expected Script:', expectedScript);
console.log('Does the provided hex correspond to the expected script?', isMatch);

