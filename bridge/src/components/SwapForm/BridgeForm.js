import React, { useState } from 'react';
import Web3Modal from 'web3modal';
import { Web3Provider } from '@ethersproject/providers';
import GbtcBtcModal from './Modals/GbtcBtcModal';
import BtcGbtcModal from './Modals/BtcGbtcModal';
import bitcoin from "bitcoinjs-lib";
import { ECPairFactory } from "ecpair";
import tinysecp from "tiny-secp256k1";
import './BridgeForm.css';
import '../BridgeButton.css'
import { ethers } from 'ethers';

const BridgeForm = () => {
  const [fromToken, setFromToken] = useState('BTC');
  const [toToken, setToToken] = useState('GBTC');
  const [amount, setAmount] = useState('');
  const [fees, setFees] = useState(0);
  const [estimatedReceived, setEstimatedReceived] = useState(0);
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [account, setAccount] = useState('');
  const [modalIsOpen, setModalIsOpen] = useState(false);
  const [btcAddress, setBtcAddress] = useState('mqYT9upmDU7WGVXWk3DKcMxGZCYiMGEhGg');
  const [ethAddress, setEthAddress] = useState('0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266');
  const [btcTxHash, setBtcTxHash] = useState('0x58a2551789add523319ba3fc996904ad6e9d3115444f9a321b6656cc88aa03dd');
  const [signedMessage, setSignedMessage] = useState('0x58a2551789add523319ba3fc996904ad6e9d3115444f9a321b6656cc88aa03dd');
  const [modalType, setModalType] = useState('');

  const avsTaskManagerCreateTaskAbi = [
    {
      "inputs": [
        { "internalType": "string", "name": "btcTxHash", "type": "string" },
        { "internalType": "string", "name": "signedMessage", "type": "string" },
        { "internalType": "address", "name": "mintTo", "type": "address" },
        { "internalType": "uint256", "name": "burnAmount", "type": "uint256" },
        { "internalType": "string", "name": "btcDestinationAddress", "type": "string" },
        { "internalType": "bool", "name": "isBurnTask", "type": "bool" },
        { "internalType": "uint32", "name": "quorumThresholdPercentage", "type": "uint32" },
        { "internalType": "bytes", "name": "quorumNumbers", "type": "bytes" }
      ],
      "name": "createNewTask",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ];

  const handleSwap = () => {
    setFromToken(fromToken === 'BTC' ? 'GBTC' : 'BTC');
    setToToken(toToken === 'BTC' ? 'GBTC' : 'BTC');
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (fromToken === 'GBTC' && toToken === 'BTC') {
      setModalType('GbtcBtc');
      setModalIsOpen(true);
    } else if (fromToken === 'BTC' && toToken === 'GBTC') {
      setModalType('BtcGbtc');
      setModalIsOpen(true);
    } else {
      // Add bridging logic here
    }
  };

  const connectWallet = async () => {
    const web3Modal = new Web3Modal();
    const instance = await web3Modal.connect();
    const newProvider = new Web3Provider(instance);
    const newSigner = newProvider.getSigner();
    const newAccount = await newSigner.getAddress();

    setProvider(newProvider);
    setSigner(newSigner);
    setAccount(newAccount);
    setEthAddress(newAccount); // Set ETH address to MetaMask address by default
  };

  const privateKey = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";

  const initiateMintOnAvs = async () => {
    try {
      if (!window.ethereum) {
        throw new Error("MetaMask is not installed");
      }

      // Set up a provider (MetaMask in this case)
      const provider = new ethers.BrowserProvider(window.ethereum);

      // Prompt user to connect their wallet
      await provider.send("eth_requestAccounts", []);

      // Get the signer
      const signer = await provider.getSigner();
      const contractAddress = "0x9E545E3C0baAB3E08CdfD552C960A1050f373042";
      const contract = new ethers.Contract(contractAddress, avsTaskManagerCreateTaskAbi, signer);

      // Convert the amount to a string before parsing it to units
      const amountString = amount.toString();
      // Define the types and values you want to encode
      const quorumNumbers = '0x00'; // Hardcoded value as seen in the QuorumTask event

      const tx = await contract.createNewTask(btcTxHash, signedMessage, ethAddress, ethers.parseUnits(amountString, 8), btcAddress, false, 100, quorumNumbers);
      await tx.wait();
      console.log('Transaction successful:', tx);
    } catch (error) {
      console.error('Transaction failed:', error);
    }
  };

  const initiateBurnOnAvs = async () => {
    try {
      if (!window.ethereum) {
        throw new Error("MetaMask is not installed");
      }

      // Set up a provider (MetaMask in this case)
      const provider = new ethers.BrowserProvider(window.ethereum);

      // Prompt user to connect their wallet
      await provider.send("eth_requestAccounts", []);

      // Get the signer
      const signer = await provider.getSigner();
      const contractAddress = "0x9E545E3C0baAB3E08CdfD552C960A1050f373042";

      const contract = new ethers.Contract(contractAddress, avsTaskManagerCreateTaskAbi, signer);

      // Convert the amount to a string before parsing it to units
      const amountString = amount.toString();
      // Define the types and values you want to encode
      const quorumNumbers = '0x00'; // Hardcoded value as seen in the QuorumTask event

      const tx = await contract.createNewTask(btcTxHash, signedMessage, ethAddress, ethers.parseUnits(amountString, 8), btcAddress, true, 100, quorumNumbers);
      await tx.wait();
      console.log('Transaction successful:', tx);
    } catch (error) {
      console.error('Transaction failed:', error);
    }
  };

  return (
    <div className="bridge-form-container" >
      <div className="bridge-form-header">
        <h2>Send</h2>
        <div className="network-select">
  <div className="network-pill">
    <input
      type="radio"
      id="btc"
      name="network"
      value="BTC"
      checked={fromToken === 'BTC'}
      onChange={handleSwap}
      style={{margin:10}}
    />
    <label htmlFor="btc">BTC</label>
    <input
      type="radio"
      id="gbtc"
      name="network"
      value="GBTC"
      checked={fromToken === 'GBTC'}
      onChange={handleSwap}
      style={{margin:10}}

    />
    <label htmlFor="gbtc">GBTC</label>
  </div>
</div>

      </div>
      <form className="bridge-form" onSubmit={handleSubmit}>
        <div className="bridge-section">
          <div className="bridge-section-header">From</div>
          <div className="bridge-section-content">
            <input
              type="number"
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
              placeholder="0.0"
              className="amount-input3"
              step="0.01"
            />
            <span className="network-currency">{fromToken}</span>
          </div>
        </div>
        <div className="bridge-arrow" onClick={handleSwap}>⇅</div>
        <div className="bridge-section">
          <div className="bridge-section-header">To (estimated)</div>
          <div className="bridge-section-content">
            <input type="number" placeholder="0.0" className="amount-input3" readOnly value={amount * 0.995} />
            <span className="network-currency">{toToken}</span>
          </div>
        </div>
        <div className="bridge-details">
          {/* Additional details can be added here */}
        </div>
        {account ? (
          <>
            <button type="submit" className="bridge-button">Bridge</button>
            <div className="connected-account" style={{ marginTop: '10px' }}>Wallet: {account}</div>
          </>
        ) : (
          <button type="button" className="bridge-button" onClick={connectWallet}>Connect A Wallet</button>
        )}
      </form>
      {modalType === 'GbtcBtc' && (
        <GbtcBtcModal
          modalIsOpen={modalIsOpen}
          setModalIsOpen={setModalIsOpen}
          amount={amount}
          fees={fees}
          btcAddress={btcAddress}
          setBtcAddress={setBtcAddress}
          handleBridge={initiateBurnOnAvs}
          setAmount={setAmount}
        />
      )}
      {modalType === 'BtcGbtc' && (
        <BtcGbtcModal
          modalIsOpen={modalIsOpen}
          setModalIsOpen={setModalIsOpen}
          amount={amount}
          fees={fees}
          ethAddress={ethAddress}
          setEthAddress={setEthAddress}
          txHash={btcTxHash}
          setTxHash={setBtcTxHash}
          signedTxHash={signedMessage}
          setSignedTxHash={setSignedMessage}
          handleBridge={initiateMintOnAvs}
        />
      )}
    </div>
  );
};

export default BridgeForm;
