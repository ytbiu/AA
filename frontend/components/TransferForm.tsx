'use client';

import { useState } from 'react';

interface TransferFormProps {
  onTransfer: (to: string, amount: string) => Promise<void>;
  loading: boolean;
  disabled?: boolean;
  submitLabel?: string;
  loadingLabel?: string;
}

export function TransferForm({
  onTransfer,
  loading,
  disabled,
  submitLabel = 'Transfer USDT',
  loadingLabel = 'Processing...',
}: TransferFormProps) {
  const [targetAddress, setTargetAddress] = useState('');
  const [amount, setAmount] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onTransfer(targetAddress, amount);
  };

  const isDisabled = disabled || loading;

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label htmlFor="targetAddress" className="block text-sm font-medium mb-1">
          Target Address
        </label>
        <input
          id="targetAddress"
          type="text"
          value={targetAddress}
          onChange={(e) => setTargetAddress(e.target.value)}
          placeholder="Enter recipient address"
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
          required
          disabled={isDisabled}
        />
      </div>

      <div>
        <label htmlFor="amount" className="block text-sm font-medium mb-1">
          Amount (USDT)
        </label>
        <input
          id="amount"
          type="number"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          placeholder="Enter amount to transfer"
          step="0.01"
          min="0.01"
          max="1000"
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
          required
          disabled={isDisabled}
        />
      </div>

      <button
        type="submit"
        disabled={isDisabled}
        className={`w-full px-4 py-2 rounded-md ${
          isDisabled
            ? 'bg-gray-400 cursor-not-allowed'
            : 'bg-green-500 hover:bg-green-600 text-white'
        }`}
      >
        {loading ? loadingLabel : disabled ? 'Connect Wallet First' : submitLabel}
      </button>
    </form>
  );
}
