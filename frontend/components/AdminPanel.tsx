'use client';

import { useState } from 'react';

interface AdminPanelProps {
  onAction: (action: string, params: any) => Promise<void>;
  loading: boolean;
}

export function AdminPanel({ onAction, loading }: AdminPanelProps) {
  const [relayerAddress, setRelayerAddress] = useState('');
  const [feeRate, setFeeRate] = useState('');
  const [oracleAddress, setOracleAddress] = useState('');

  const handleAddRelayer = async (e: React.FormEvent) => {
    e.preventDefault();
    await onAction('addRelayer', { address: relayerAddress });
    setRelayerAddress('');
  };

  const handleRemoveRelayer = async (e: React.FormEvent) => {
    e.preventDefault();
    await onAction('removeRelayer', { address: relayerAddress });
    setRelayerAddress('');
  };

  const handleChangeFeeRate = async (e: React.FormEvent) => {
    e.preventDefault();
    await onAction('setFeeRate', { rate: parseFloat(feeRate) });
    setFeeRate('');
  };

  const handleChangeOracle = async (e: React.FormEvent) => {
    e.preventDefault();
    await onAction('setOracle', { address: oracleAddress });
    setOracleAddress('');
  };

  return (
    <div className="space-y-6">
      {/* Relayer Management */}
      <div className="border border-gray-200 rounded-lg p-4">
        <h3 className="text-lg font-semibold mb-3">Relayer Management</h3>

        <form onSubmit={handleAddRelayer} className="mb-4 flex gap-2">
          <input
            type="text"
            value={relayerAddress}
            onChange={(e) => setRelayerAddress(e.target.value)}
            placeholder="Relayer address"
            className="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <button
            type="submit"
            disabled={loading}
            className={`px-4 py-2 rounded-md ${
              loading
                ? 'bg-gray-400 cursor-not-allowed'
                : 'bg-blue-500 hover:bg-blue-600 text-white'
            }`}
          >
            Add Relayer
          </button>
        </form>

        <form onSubmit={handleRemoveRelayer} className="flex gap-2">
          <input
            type="text"
            value={relayerAddress}
            onChange={(e) => setRelayerAddress(e.target.value)}
            placeholder="Relayer address"
            className="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <button
            type="submit"
            disabled={loading}
            className={`px-4 py-2 rounded-md ${
              loading
                ? 'bg-gray-400 cursor-not-allowed'
                : 'bg-red-500 hover:bg-red-600 text-white'
            }`}
          >
            Remove Relayer
          </button>
        </form>
      </div>

      {/* Fee Rate Configuration */}
      <div className="border border-gray-200 rounded-lg p-4">
        <h3 className="text-lg font-semibold mb-3">Fee Rate Configuration</h3>

        <form onSubmit={handleChangeFeeRate} className="flex gap-2">
          <input
            type="number"
            value={feeRate}
            onChange={(e) => setFeeRate(e.target.value)}
            placeholder="New fee rate (e.g., 0.01 for 1%)"
            step="any"
            min="0"
            className="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <button
            type="submit"
            disabled={loading}
            className={`px-4 py-2 rounded-md ${
              loading
                ? 'bg-gray-400 cursor-not-allowed'
                : 'bg-purple-500 hover:bg-purple-600 text-white'
            }`}
          >
            Set Fee Rate
          </button>
        </form>
      </div>

      {/* Oracle Configuration */}
      <div className="border border-gray-200 rounded-lg p-4">
        <h3 className="text-lg font-semibold mb-3">Oracle Configuration</h3>

        <form onSubmit={handleChangeOracle} className="flex gap-2">
          <input
            type="text"
            value={oracleAddress}
            onChange={(e) => setOracleAddress(e.target.value)}
            placeholder="New oracle address"
            className="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <button
            type="submit"
            disabled={loading}
            className={`px-4 py-2 rounded-md ${
              loading
                ? 'bg-gray-400 cursor-not-allowed'
                : 'bg-orange-500 hover:bg-orange-600 text-white'
            }`}
          >
            Set Oracle
          </button>
        </form>
      </div>
    </div>
  );
}