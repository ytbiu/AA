import React, { useState } from 'react';

interface PrivateKeyInputProps {
  onSubmit: (privateKey: string) => void;
  disabled?: boolean;
}

const PrivateKeyInput: React.FC<PrivateKeyInputProps> = ({ onSubmit, disabled }) => {
  const [showWarning, setShowWarning] = useState(true);
  const [privateKey, setPrivateKey] = useState('');
  const [error, setError] = useState('');

  // Function to validate private key format
  const isValidPrivateKey = (key: string): boolean => {
    // Remove 0x prefix if present
    const cleanedKey = key.startsWith('0x') ? key.slice(2) : key;

    // Private key should be 64 characters hex string
    const privateKeyRegex = /^[0-9a-fA-F]{64}$/;
    return privateKeyRegex.test(cleanedKey);
  };

  const handleSubmit = () => {
    setError('');

    if (!isValidPrivateKey(privateKey)) {
      setError('Invalid private key format. Must be a 64-character hex string.');
      return;
    }

    // If valid, call the onSubmit callback and clear the input
    onSubmit(privateKey);
    setPrivateKey(''); // Clear input after submission
  };

  if (showWarning) {
    return (
      <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-4">
        <div className="flex items-start">
          <div className="flex-shrink-0">
            <svg className="h-5 w-5 text-yellow-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
              <path fillRule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
            </svg>
          </div>
          <div className="ml-3">
            <p className="text-sm text-yellow-700">
              <strong>Security Warning:</strong> Please be aware of the following security considerations:
            </p>
            <ul className="mt-2 text-xs text-yellow-600 list-disc pl-5 space-y-1">
              <li>Your private key is only used for local signing operations and is not transmitted to any server</li>
              <li>Please clear this page after completing your transaction</li>
              <li>Do not use this on public computers or devices you do not trust</li>
              <li>We do not recommend saving your private key in browser storage</li>
            </ul>
            <div className="mt-3">
              <button
                onClick={() => setShowWarning(false)}
                className="inline-flex items-center px-3 py-1 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-yellow-600 hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-yellow-500"
              >
                I Understand and Accept Risks
              </button>
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div>
        <label htmlFor="privateKey" className="block text-sm font-medium text-gray-700 mb-1">
          Private Key
        </label>
        <input
          id="privateKey"
          type="password"
          value={privateKey}
          onChange={(e) => {
            setPrivateKey(e.target.value);
            setError(''); // Clear error when user types
          }}
          placeholder="Enter your private key (64-character hex string)"
          className={`w-full px-3 py-2 border ${
            error ? 'border-red-500' : 'border-gray-300'
          } rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500`}
          disabled={disabled}
        />
        {error && (
          <p className="mt-1 text-sm text-red-600">{error}</p>
        )}
      </div>

      <div className="flex space-x-3">
        <button
          onClick={handleSubmit}
          disabled={disabled || !privateKey.trim()}
          className={`inline-flex items-center px-3 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white ${
            disabled || !privateKey.trim()
              ? 'bg-gray-400 cursor-not-allowed'
              : 'bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500'
          }`}
        >
          Submit Private Key
        </button>

        <button
          onClick={() => {
            setPrivateKey(''); // Clear input
            setError(''); // Clear error
          }}
          disabled={disabled}
          className={`inline-flex items-center px-3 py-2 border border-gray-300 text-sm font-medium rounded-md shadow-sm ${
            disabled
              ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
              : 'bg-white text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500'
          }`}
        >
          Clear
        </button>
      </div>
    </div>
  );
};

export default PrivateKeyInput;