import React, { useState, useEffect } from 'react';
import Modal from 'react-modal';
import './BtcGbtcModal.css';
import Loader from '../../Loader';

const BtcGbtcModal = ({
  modalIsOpen,
  setModalIsOpen,
  amount,
  fees,
  ethAddress,
  setEthAddress,
  txHash,
  setTxHash,
  signedTxHash,
  setSignedTxHash,
  handleBridge,
}) => {
  const [btcPrice, setBtcPrice] = useState(0);
  const [usdValue, setUsdValue] = useState(0);
  const [loading, setLoading] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");

  useEffect(() => {
    const fetchBtcPrice = async () => {
      try {
        const response = await fetch('https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd');
        const data = await response.json();
        setBtcPrice(data.bitcoin.usd);
      } catch (error) {
        console.error('Error fetching BTC price:', error);
      }
    };

    fetchBtcPrice();
  }, []);

  useEffect(() => {
    setUsdValue(amount * btcPrice);
  }, [amount, btcPrice]);

  async function handleMintSetLoading() {
    await handleBridge();
    setLoading(true);
    setTimeout(() => {
      setLoading(false);
      setSuccessMessage("Bridge request processed successfully");
    }, 5000);
  }

  const formatUsdValue = (value) => {
    if (value >= 1000000) {
      return `${(value / 1000000).toFixed(2)}M USD`;
    }
    return `${Math.round(value).toLocaleString()} USD`;
  };

  return (
    <Modal
      isOpen={modalIsOpen}
      onRequestClose={() => setModalIsOpen(false)}
      contentLabel="Bridge Modal"
      style={{
        content: {
          borderRadius: '16px',
          padding: '20px',
          maxWidth: '600px',
          width: '90%',
          margin: '0 auto',
          boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
          transition: 'transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out',
          zIndex:10
        },
        overlay: {
          backgroundColor: 'rgba(0, 0, 0, 0.5)',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }
      }}
    >
      <div className="modal-form-container">
        <h2>Bridge BTC to GBTC</h2>
        <p className="instructions">
          Send your Bitcoin to the following address before filling out these fields: <strong>2N4ZJY7R6FKZsc7SqKw4yZ5eHZo9BP3C7Go</strong>
        </p>
        {loading && <Loader />}
        <div className="bridge-details">
          <div className="detail-row">
            <span className="detail-title">Amount in BTC:</span>
            <span className="detail-value">{amount}</span>
          </div>
          <div className="detail-row">
            <span className="detail-title">USD Value:</span>
            <span className="detail-value">{formatUsdValue(usdValue)}</span>
          </div>
          <div className="detail-row">
            <span className="detail-title">Fee:</span>
            <span className="detail-value">{(amount * 0.005).toFixed(6)} BTC | {formatUsdValue(usdValue * 0.005)}</span>
          </div>
          <div style={{width:"100%", height:"7px", borderRadius:'50px', backgroundColor:"black", marginTop:"30px", marginBottom:"30px"}}></div>

          <div className="bridge-section">
            <div className="bridge-section-header">
              ETH Destination Address
              <span className="info-icon" data-tooltip="Enter your Ethereum address where the GBTC will be sent.">ℹ️</span>
            </div>
            <div className="bridge-section-content">
              <input
                type="text"
                value={ethAddress}
                onChange={(e) => setEthAddress(e.target.value)}
                placeholder="Enter your ETH address"
                className="amount-input"
              />
            </div>
          </div>
          <div className="bridge-section">
            <div className="bridge-section-header">
              Hash of the accepted Bitcoin transaction
              <span className="info-icon" data-tooltip="Enter the transaction hash of the Bitcoin you sent to the specified address.">ℹ️</span>
            </div>
            <div className="bridge-section-content">
              <input
                type="text"
                value={txHash}
                onChange={(e) => setTxHash(e.target.value)}
                placeholder="Enter the transaction hash"
                className="amount-input"
              />
            </div>
          </div>
          <div className="bridge-section">
            <div className="bridge-section-header">
              ETH Destination Address signed by BTC private key
              <span className="info-icon" data-tooltip="Enter the Ethereum address signed with your Bitcoin private key that was used to send the Bitcoin transaction">ℹ️</span>
            </div>
            <div className="bridge-section-content">
              <input
                type="text"
                value={signedTxHash}
                onChange={(e) => setSignedTxHash(e.target.value)}
                placeholder="Enter the signed transaction hash"
                className="amount-input"
              />
            </div>
          </div>
          {successMessage.length === 0 && <button onClick={handleMintSetLoading} className="bridge-button" >Bridge BTC to GBTC</button>}
          <button onClick={() => setModalIsOpen(false)} className="connect-wallet-button">{successMessage.length > 0 ? "Done" : "Cancel"}</button>

          {successMessage.length > 0 && <a className="success-message">{successMessage}</a>}
        </div>
      </div>
    </Modal>
  );
};

export default BtcGbtcModal;
