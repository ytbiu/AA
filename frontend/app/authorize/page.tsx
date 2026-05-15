'use client';

import { useState } from 'react';
import { PrivateKeyInput } from '@/components/PrivateKeyInput';
import { authorize7702 } from '@/lib/api';
import { signAuthorization, getPrivateKeyAddress } from '@/lib/signer';
import { CONTRACTS } from '@/lib/contracts';

export default function AuthorizePage() {
  const [privateKey, setPrivateKey] = useState('');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  const handleAuthorize = async () => {
    if (!privateKey) {
      alert('Please enter your private key');
      return;
    }

    setLoading(true);
    setResult(null);

    try {
      const userAddress = getPrivateKeyAddress(privateKey);

      // Sign authorization for 7702 delegation
      const authorization = await signAuthorization(
        privateKey,
        CONTRACTS.ECDSAValidator,
        Math.floor(Date.now() / 1000) + 3600 // Valid for 1 hour
      );

      // Submit to the API
      const response = await authorize7702(userAddress, CONTRACTS.ECDSAValidator, authorization.signature);
      setResult(response);
    } catch (error) {
      console.error('Error authorizing 7702:', error);
      setResult({ error: 'Failed to authorize 7702 delegation' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">EIP-7702 Authorization</h1>

      <div className="mb-6 p-4 bg-blue-100 rounded-lg">
        <h2 className="text-xl font-semibold mb-2">About EIP-7702</h2>
        <p>
          EIP-7702 allows accounts to temporarily delegate control to a smart contract.
          This enables traditional EOAs to gain smart contract functionality such as
          batched transactions, meta-transactions, and more efficient gas usage.
        </p>
      </div>

      <div className="mb-6">
        <PrivateKeyInput onPrivateKeyChange={setPrivateKey} />
      </div>

      <button
        onClick={handleAuthorize}
        disabled={loading}
        className={`px-4 py-2 rounded-md ${
          loading
            ? 'bg-gray-400 cursor-not-allowed'
            : 'bg-blue-500 hover:bg-blue-600 text-white'
        }`}
      >
        {loading ? 'Authorizing...' : 'Authorize EIP-7702 Delegation'}
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