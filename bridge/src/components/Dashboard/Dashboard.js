import React from 'react';
import './Dashboard.css';
import BridgeForm from '../SwapForm/BridgeForm';
import logo from '../../assets/logo.png';
import { ethers } from 'ethers';

const Dashboard = () => {


  const handleMint = async (destAddress, txSignature, txHash) => {
    try {
      // Ensure the user is connected to MetaMask and has allowed access
      if (!window.ethereum) {
        throw new Error("MetaMask is not installed");
      }
  
      // Set up a provider (MetaMask in this case)
      const provider = new ethers.BrowserProvider(window.ethereum);
  
      // Prompt user to connect their wallet
      await provider.send("eth_requestAccounts", []);
  
      // Get the signer
      const signer = await provider.getSigner();
  
      // Contract address and ABI
      const contractAddress = "0x5f3f1dBD7B74C6B46e8c44f98792A1dAf8d69154"; // Replace with your contract address
      const contractABI = [
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "account",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "amount",
              "type": "uint256"
            },
            {
              "internalType": "string",
              "name": "btcTxHash",
              "type": "string"
            }
          ],
          "name": "mint",
          "outputs": [],
          "stateMutability": "nonpayable",
          "type": "function"
        }
      ];
  
      // Create a contract instance
      const contract = new ethers.Contract(contractAddress, contractABI, signer);
      const amount = 100;
  
      // Convert the amount to a string before parsing it to units
      const amountString = amount.toString();
  
      // Call the mint function
      const tx = await contract.mint(destAddress, ethers.parseUnits(amountString, 8), txHash);
  
      // Wait for the transaction to be mined
      await tx.wait();
  
      console.log('Mint transaction successful:', tx);
    } catch (error) {
      console.error('Mint transaction failed:', error);
  
      // Log the specific error message
      if (error.data && error.data.message) {
        console.error('Error message:', error.data.message);
      }
  
      // Log the full error object
      console.error('Full error object:', error);
    }
  };
  
  
 
  return (
    <div className="dashboard-container">
      <div className="dashboard-text">
        <img src={logo} alt="Logo" className="logo" />
        <h1>Gazelle, Trustless Bitcoin-Ethereum Bridge</h1>
        <p>
          This project aims to provide a trustless, decentralized bridge between Bitcoin and Ethereum networks. By utilizing Eigen Layer and wrapped tokens, it allows for seamless, secure swaps between BTC and GBTC, ensuring that users can leverage the advantages of both blockchains without relying on a central authority.
        </p>
        <p>
          <a href="https://github.com/your-repo-link" target="_blank" rel="noopener noreferrer">
            View the Code Repository
          </a>
        </p>
        <p>
          <a href="https://eigenlayer.xyz" target="_blank" rel="noopener noreferrer" className="eigen-layer-button">
            Find out more about Eigen Layer
          </a>
        </p>
      </div>
      <div className="dashboard-swap">
        

        <BridgeForm />
      </div>
    </div>
  );
};

export default Dashboard;

