'use client';

import React, { useState } from 'react';
import StatusCard from '@/components/StatusCard';
import { getUserStatus, UserStatus } from '@/lib/api';

const HomePage: React.FC = () => {
  const [address, setAddress] = useState<string>('');
  const [status, setStatus] = useState<UserStatus | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const handleQuery = async () => {
    if (!address.trim()) {
      setError('Please enter a valid address');
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const userStatus = await getUserStatus(address);
      setStatus(userStatus);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch user status');
      setStatus(null);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-3xl mx-auto">
        <div className="text-center mb-10">
          <h1 className="text-3xl font-bold text-gray-900">Account Status Dashboard</h1>
          <p className="mt-2 text-gray-600">Query user account information and status</p>
        </div>

        <div className="bg-white rounded-lg shadow p-6 mb-8">
          <div className="flex flex-col sm:flex-row gap-4">
            <input
              type="text"
              value={address}
              onChange={(e) => setAddress(e.target.value)}
              placeholder="Enter wallet address"
              className="flex-grow px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              disabled={loading}
            />
            <button
              onClick={handleQuery}
              disabled={loading}
              className={`px-6 py-2 rounded-md text-white font-medium ${
                loading
                  ? 'bg-blue-400 cursor-not-allowed'
                  : 'bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2'
              }`}
            >
              {loading ? 'Loading...' : 'Query Status'}
            </button>
          </div>

          {error && (
            <div className="mt-4 p-3 bg-red-50 text-red-700 rounded-md">
              {error}
            </div>
          )}
        </div>

        {status && (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <StatusCard
              title="Wallet Address"
              value={status.address}
              type="address"
            />
            <StatusCard
              title="7702 Bound"
              value={status.is_7702_bound}
              type="boolean"
            />
            <StatusCard
              title="Bound Contract"
              value={status.bound_contract}
              type="address"
            />
            <StatusCard
              title="USDT Balance"
              value={parseFloat(status.usdt_balance)}
              type="balance"
            />
          </div>
        )}
      </div>
    </div>
  );
};

export default HomePage;