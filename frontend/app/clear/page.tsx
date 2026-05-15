'use client';

import { useState } from 'react';
import { PrivateKeyInput } from '@/components/PrivateKeyInput';
import { clear7702 } from '@/lib/api';
import { getPrivateKeyAddress } from '@/lib/signer';

export default function ClearPage() {
  const [privateKey, setPrivateKey] = useState('');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  const handleClear = async () => {
    if (!privateKey) {
      alert('Please enter your private key');
      return;
    }

    setLoading(true);
    setResult(null);

    try {
      const userAddress = getPrivateKeyAddress(privateKey);

      // Clear the 7702 delegation
      const response = await clear7702(userAddress);
      setResult(response);
    } catch (error) {
      console.error('Error clearing 7702:', error);
      setResult({ error: 'Failed to clear 7702 delegation' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">Clear EIP-7702 Authorization</h1>

      <div className="mb-6 p-4 bg-red-100 border border-red-400 rounded-lg">
        <h2 className="text-xl font-semibold mb-2">Warning</h2>
        <p>
          Clearing the EIP-7702 authorization will remove the delegation to the smart contract validator.
          This means you will lose smart contract wallet features like batched transactions and
          meta-transactions, and your account will revert to a standard EOA (Externally Owned Account).
        </p>
      </div>

      <div className="mb-6">
        <PrivateKeyInput onPrivateKeyChange={setPrivateKey} />
      </div>

      <button
        onClick={handleClear}
        disabled={loading}
        className={`px-4 py-2 rounded-md ${
          loading
            ? 'bg-gray-400 cursor-not-allowed'
            : 'bg-red-500 hover:bg-red-600 text-white'
        }`}
      >
        {loading ? 'Clearing...' : 'Clear EIP-7702 Delegation'}
      </button>

      {result && (
        <div className="mt-6 p-4 bg-gray-100 rounded-lg">
          <h3 className="font-semibold mb-2">Result:</h3>
          <pre className="whitespace-pre-wrap">{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}