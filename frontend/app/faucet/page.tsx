'use client';

import { useState, useEffect } from 'react';
import { getFaucetInfo } from '@/lib/api';
import { CONTRACTS } from '@/lib/contracts';

export default function FaucetPage() {
  const [faucetInfo, setFaucetInfo] = useState<any>(null);
  const [address, setAddress] = useState('');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  useEffect(() => {
    fetchFaucetInfo();
  }, []);

  const fetchFaucetInfo = async () => {
    try {
      const info = await getFaucetInfo();
      setFaucetInfo(info);
    } catch (error) {
      console.error('Error fetching faucet info:', error);
    }
  };

  const handleClaim = async () => {
    if (!address) {
      alert('Please enter an address');
      return;
    }

    setLoading(true);
    setResult(null);

    try {
      // Faucet requires user to call contract directly (need BNB for gas)
      // Implementation would go here using the smart wallet infrastructure
      const response = await fetch('/api/faucet', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ address }),
      });

      const data = await response.json();
      setResult(data);
    } catch (error) {
      console.error('Error claiming faucet:', error);
      setResult({ error: 'Failed to claim faucet' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">Faucet</h1>

      {faucetInfo && (
        <div className="mb-6 p-4 bg-gray-100 rounded-lg">
          <h2 className="text-xl font-semibold mb-2">Faucet Info</h2>
          <p>Amount: {faucetInfo.amount} BNB</p>
          <p>Available: {faucetInfo.available ? 'Yes' : 'No'}</p>
        </div>
      )}

      <div className="mb-6">
        <label htmlFor="address" className="block text-sm font-medium mb-2">
          Address to Receive Funds
        </label>
        <input
          id="address"
          type="text"
          value={address}
          onChange={(e) => setAddress(e.target.value)}
          placeholder="Enter wallet address"
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>

      <button
        onClick={handleClaim}
        disabled={loading}
        className={`px-4 py-2 rounded-md ${
          loading
            ? 'bg-gray-400 cursor-not-allowed'
            : 'bg-blue-500 hover:bg-blue-600 text-white'
        }`}
      >
        {loading ? 'Claiming...' : 'Claim Faucet'}
      </button>

      {result && (
        <div className="mt-6 p-4 bg-gray-100 rounded-lg">
          <h3 className="font-semibold mb-2">Result:</h3>
          <pre className="whitespace-pre-wrap">{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}

      <div className="mt-6 p-4 bg-yellow-100 border border-yellow-400 rounded-lg">
        <p className="text-sm">
          <strong>Note:</strong> Faucet requires user to call contract directly (need BNB for gas).
          Make sure you have sufficient gas fees to cover the transaction.
        </p>
      </div>
    </div>
  );
}