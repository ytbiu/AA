'use client';

import { useState } from 'react';
import PrivateKeyInput from '@/components/PrivateKeyInput';
import { TransferForm } from '@/components/TransferForm';
import { transferUSDT } from '@/lib/api';
import { getPrivateKeyAddress } from '@/lib/signer';
import { createUSDTTransferBatch, signBatchOperation } from '@/lib/batchSigner';
import { CONTRACTS } from '@/lib/contracts';
import { ethers } from 'ethers';

export default function TransferPage() {
  const [privateKey, setPrivateKey] = useState('');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  const handleTransfer = async (to: string, amount: string) => {
    if (!privateKey) {
      alert('Please enter your private key');
      return;
    }

    // 验证金额格式和范围
    const amountNum = parseFloat(amount);
    if (isNaN(amountNum) || amountNum <= 0) {
      alert('Please enter a valid positive amount');
      return;
    }
    if (amountNum > 1000) {
      alert('Amount too large. Maximum is 1000 USDT');
      return;
    }

    setLoading(true);
    setResult(null);

    try {
      const userAddress = getPrivateKeyAddress(privateKey);

      // 转换金额到最小单位 (USDT 18 decimals)
      const amountInSmallestUnit = ethers.parseUnits(amount, 18).toString();

      // 创建批量操作 (approve + transfer)
      const batch = createUSDTTransferBatch(
        userAddress,
        CONTRACTS.USDT,
        CONTRACTS.PAYMASTER,
        to,
        amount
      );

      // 签名批量操作
      const signature = await signBatchOperation(privateKey, batch);

      // 执行转账 - 传递最小单位的金额
      const response = await transferUSDT(userAddress, to, amountInSmallestUnit, signature);
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

      <div className="mb-6 p-4 bg-blue-100 rounded-lg">
        <h2 className="text-xl font-semibold mb-2">Gasless Transfer</h2>
        <p>
          This page allows you to transfer USDT without paying gas fees.
          The Relayer will execute the transaction on your behalf using the Paymaster.
        </p>
        <p className="mt-2 text-sm text-gray-600">
          <strong>Note:</strong> You need to have USDT balance and have authorized EIP-7702 first.
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div>
          <h2 className="text-xl font-semibold mb-4">Private Key</h2>
          <div className="mb-6">
            <PrivateKeyInput onSubmit={(pk) => setPrivateKey(pk)} />
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