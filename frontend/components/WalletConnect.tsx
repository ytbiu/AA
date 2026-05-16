'use client'

import { useState, useEffect } from 'react'
import { ethers } from 'ethers'

declare global {
  interface Window {
    ethereum?: any
  }
}

interface WalletConnectProps {
  onConnect: (address: string, signer: ethers.JsonRpcSigner) => void
  onDisconnect: () => void
  connectedAddress: string | null
}

export default function WalletConnect({ onConnect, onDisconnect, connectedAddress }: WalletConnectProps) {
  const [showInfo, setShowInfo] = useState(true)
  const [connecting, setConnecting] = useState(false)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    checkConnection()
  }, [])

  const checkConnection = async () => {
    if (typeof window !== 'undefined' && window.ethereum) {
      try {
        const provider = new ethers.BrowserProvider(window.ethereum)
        const accounts = await provider.listAccounts()
        if (accounts.length > 0) {
          const signer = await provider.getSigner()
          onConnect(accounts[0].address, signer)
          setShowInfo(false)
        }
      } catch (err) {
        console.error('Check connection error:', err)
      }
    }
  }

  const handleConnect = async () => {
    setConnecting(true)
    setError(null)

    try {
      if (!window.ethereum) {
        setError('MetaMask not detected. Please install MetaMask extension.')
        return
      }

      const provider = new ethers.BrowserProvider(window.ethereum)

      // Request network switch to BSC Testnet
      try {
        await provider.send('wallet_switchEthereumChain', [{ chainId: '0x61' }])
      } catch (switchError: any) {
        if (switchError.code === 4902) {
          // Chain not added, add it
          await provider.send('wallet_addEthereumChain', [{
            chainId: '0x61',
            chainName: 'BSC Testnet',
            nativeCurrency: { name: 'BNB', symbol: 'BNB', decimals: 18 },
            rpcUrls: ['https://data-seed-prebsc-1-s1.binance.org:8545/'],
            blockExplorerUrls: ['https://testnet.bscscan.com/']
          }])
        }
      }

      const accounts = await provider.send('eth_requestAccounts', [])
      if (accounts.length > 0) {
        const signer = await provider.getSigner()
        onConnect(accounts[0], signer)
        setShowInfo(false)
      }
    } catch (err: any) {
      console.error('Connection error:', err)
      if (err.code === 4001) {
        setError('Connection rejected by user')
      } else {
        setError('Failed to connect wallet: ' + (err.message || 'Unknown error'))
      }
    } finally {
      setConnecting(false)
    }
  }

  // Connected state
  if (connectedAddress) {
    return (
      <div className="bg-green-50 border-l-4 border-green-400 p-4">
        <div className="flex items-start justify-between">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <svg className="h-5 w-5 text-green-400" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
              </svg>
            </div>
            <div className="ml-3">
              <p className="text-sm font-medium text-green-700">Wallet Connected</p>
              <p className="text-xs text-green-600 mt-1">{connectedAddress}</p>
            </div>
          </div>
          <button
            onClick={onDisconnect}
            className="inline-flex items-center px-3 py-1 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-red-500 hover:bg-red-600 focus:outline-none"
          >
            Disconnect
          </button>
        </div>
      </div>
    )
  }

  // Info/warning state
  if (showInfo) {
    return (
      <div className="bg-blue-50 border-l-4 border-blue-400 p-4">
        <div className="flex items-start">
          <div className="flex-shrink-0">
            <svg className="h-5 w-5 text-blue-400" viewBox="0 0 20 20" fill="currentColor">
              <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clipRule="evenodd" />
            </svg>
          </div>
          <div className="ml-3">
            <p className="text-sm text-blue-700">
              <strong>Connect Your Wallet:</strong>
            </p>
            <ul className="mt-2 text-xs text-blue-600 list-disc pl-5 space-y-1">
              <li>Click to connect your MetaMask or other Web3 wallet</li>
              <li>Your wallet will be prompted to sign transactions locally</li>
              <li>Make sure you are on BSC Testnet (Chain ID: 97)</li>
              <li>No private keys are stored or transmitted</li>
            </ul>
            <div className="mt-3">
              <button
                onClick={() => setShowInfo(false)}
                className="inline-flex items-center px-3 py-1 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                Continue to Connect
              </button>
            </div>
          </div>
        </div>
      </div>
    )
  }

  // Connect button state
  return (
    <div className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Connect Wallet
        </label>
        <button
          onClick={handleConnect}
          disabled={connecting}
          className={`w-full px-4 py-3 border ${
            connecting
              ? 'border-gray-300 bg-gray-100 cursor-not-allowed'
              : 'border-transparent bg-indigo-600 hover:bg-indigo-700'
          } text-white font-medium rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500`}
        >
          {connecting ? 'Connecting...' : 'Connect MetaMask'}
        </button>
        {error && (
          <p className="mt-2 text-sm text-red-600">{error}</p>
        )}
        <p className="mt-2 text-xs text-gray-500">
          Requires MetaMask or other Web3 wallet extension
        </p>
      </div>
    </div>
  )
}