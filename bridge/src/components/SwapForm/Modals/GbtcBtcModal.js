import React, { useState, useEffect } from 'react';
import Modal from 'react-modal';
import './GbtcBtcModal.css';
import Loader from '../../Loader'; // Ensure you have the Loader component

const GbtcBtcModal = ({ modalIsOpen, setModalIsOpen, amount, btcAddress, setBtcAddress, handleBridge, setAmount }) => {
  const [btcPrice, setBtcPrice] = useState(0);
  const [usdValue, setUsdValue] = useState(0);
  const [fees, setFees] = useState(0);
  const [loading, setLoading] = useState(false); // Add loading state

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
    setFees(amount * 0.003); // 0.3% fee
  }, [amount, btcPrice]);

  const formatUsdValue = (value) => {
    if (value >= 1000000) {
      return `${(value / 1000000).toFixed(2)}M USD`;
    }
    return `${Math.round(value).toLocaleString()} USD`;
  };

  const handleBridgeClick = async () => {
    setLoading(true);
    setTimeout(() => {
      setLoading(false);
    }, 5000);
    await handleBridge(); // Call handleBridge if needed
  };

  return (
    <Modal
      isOpen={modalIsOpen}
      onRequestClose={() => setModalIsOpen(false)}
      contentLabel="Bridge Modal"
      style={{
        content: {
          borderRadius: '16px',
          padding: '40px',
          maxWidth: '500px',
          margin: '0 auto',
          boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
          transition: 'transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out',
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
        <h2>Bridge GBTC to BTC</h2>
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
            <span className="detail-value">{fees.toFixed(8)} BTC</span>
          </div>
          <div className="bridge-section">
            <div className="bridge-section-header">BTC Address</div>
            <div className="bridge-section-content">
              <input
                type="text"
                value={btcAddress}
                onChange={(e) => setBtcAddress(e.target.value)}
                placeholder="Enter your BTC address"
                className="amount-input2"
              />
            </div>
            
          </div>
          <div className="bridge-section">
            <div className="bridge-section-header">Amount</div>
            <div className="bridge-section-content">
              <input
                type="number"
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                placeholder="Enter the amount you want to transfer"
                className="amount-input2"
              />
            </div>
            
          </div>
          {loading ? (
null          ) : (
            <button onClick={handleBridgeClick} className="connect-wallet-button">Bridge</button>
          )}
          <button onClick={() => setModalIsOpen(false)} className="connect-wallet-button cancel-button">Cancel</button>
        </div>
      </div>
    </Modal>
  );
};

export default GbtcBtcModal;
