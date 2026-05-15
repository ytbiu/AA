'use client';

import { useState } from 'react';
import { PrivateKeyInput } from '@/components/PrivateKeyInput';
import { TransferForm } from '@/components/TransferForm';
import { transferUSDT, createUSDTTransferBatch, signBatchOperation } from '@/lib/api';
import { getPrivateKeyAddress } from '@/lib/signer';

export default function TransferPage() {
  const [privateKey, setPrivateKey] = useState('');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  const handleTransfer = async (to: string, amount: string) => {
    if (!privateKey) {
      alert('Please enter your private key');
      return;
    }

    setLoading(true);
    setResult(null);

    try {
      const userAddress = getPrivateKeyAddress(privateKey);

      // Create a USDT transfer batch
      const batchData = await createUSDTTransferBatch(userAddress, to, amount);

      // Sign the batch operation
      const signature = await signBatchOperation(privateKey, batchData);

      // Execute the transfer
      const response = await transferUSDT(userAddress, to, amount, signature);
      setResult(response);
    } catch (error) {
      console.error('Error transferring USDT:', error);
      setResult({ error: 'Failed to transfer USDT' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">USDT Transfer</h1>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div>
          <h2 className="text-xl font-semibold mb-4">Private Key</h2>
          <div className="mb-6">
            <PrivateKeyInput onPrivateKeyChange={setPrivateKey} />
          </div>
        </div>

        <div>
          <h2 className="text-xl font-semibold mb-4">Transfer Details</h2>
          <TransferForm onTransfer={handleTransfer} loading={loading} />
        </div>
      </div>

      {result && (
        <div className="mt-6 p-4 bg-gray-100 rounded-lg">
          <h3 className="font-semibold mb-2">Result:</h3>
          <pre className="whitespace-pre-wrap">{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}