'use client'

import { useState } from 'react'
import WalletConnect from '@/components/WalletConnect'
import { TransferForm } from '@/components/TransferForm'
import { quoteTransferUSDT, transferUSDT, type TransferUSDTQuoteResponse } from '@/lib/api'
import { createUSDTTransferBatch, signBatchOperationWithWallet } from '@/lib/batchSigner'
import { CONTRACTS } from '@/lib/contracts'
import { ethers } from 'ethers'

interface PendingTransfer {
  userAddress: string
  targetAddress: string
  amount: string
  displayAmount: string
  signature: string
  quote: TransferUSDTQuoteResponse
}

export default function TransferPage() {
  const [connectedAddress, setConnectedAddress] = useState<string | null>(null)
  const [signer, setSigner] = useState<ethers.JsonRpcSigner | null>(null)
  const [loading, setLoading] = useState(false)
  const [confirming, setConfirming] = useState(false)
  const [result, setResult] = useState<any>(null)
  const [pendingTransfer, setPendingTransfer] = useState<PendingTransfer | null>(null)

  const handleConnect = (address: string, walletSigner: ethers.JsonRpcSigner) => {
    setConnectedAddress(address)
    setSigner(walletSigner)
  }

  const handleDisconnect = () => {
    setConnectedAddress(null)
    setSigner(null)
    setResult(null)
    setPendingTransfer(null)
  }

  const handleTransfer = async (to: string, amount: string) => {
    if (!connectedAddress || !signer) {
      alert('Please connect your wallet first')
      return
    }

    // 验证金额格式和范围
    const amountNum = parseFloat(amount)
    if (isNaN(amountNum) || amountNum <= 0) {
      alert('Please enter a valid positive amount')
      return
    }
    if (amountNum > 1000) {
      alert('Amount too large. Maximum is 1000 USDT')
      return
    }

    setLoading(true)
    setResult(null)
    setPendingTransfer(null)

    try {
      // 转换金额到最小单位 (USDT 18 decimals)
      const amountInSmallestUnit = ethers.parseUnits(amount, 18).toString()

      // 创建批量操作 (approve + transfer)
      const batch = createUSDTTransferBatch(
        connectedAddress,
        CONTRACTS.USDT,
        CONTRACTS.PAYMASTER,
        to,
        amount
      )

      // 使用钱包签名批量操作
      const signature = await signBatchOperationWithWallet(signer, batch)

      const quote = await quoteTransferUSDT(connectedAddress, to, amountInSmallestUnit, signature)
      setPendingTransfer({
        userAddress: connectedAddress,
        targetAddress: to,
        amount: amountInSmallestUnit,
        displayAmount: amount,
        signature,
        quote,
      })
    } catch (error: any) {
      console.error('Error transferring USDT:', error)
      if (error.code === 4001) {
        setResult({ error: 'Signature rejected by user' })
      } else {
        setResult({ error: error.message || 'Failed to estimate transfer USDT gas' })
      }
    } finally {
      setLoading(false)
    }
  }

  const handleConfirmTransfer = async () => {
    if (!pendingTransfer) {
      return
    }

    setConfirming(true)
    setResult(null)

    try {
      const response = await transferUSDT(
        pendingTransfer.userAddress,
        pendingTransfer.targetAddress,
        pendingTransfer.amount,
        pendingTransfer.signature,
        pendingTransfer.quote.relayer_address
      )
      setResult(response)
      setPendingTransfer(null)
    } catch (error: any) {
      console.error('Error sending USDT transfer:', error)
      setResult({ error: error.message || 'Failed to transfer USDT' })
    } finally {
      setConfirming(false)
    }
  }

  const formatUsdt = (value: string) => {
    try {
      return Number(ethers.formatUnits(value, 18)).toFixed(6)
    } catch {
      return value
    }
  }

  const formatGwei = (value: string) => {
    try {
      return Number(ethers.formatUnits(value, 'gwei')).toFixed(3)
    } catch {
      return value
    }
  }

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
          <h2 className="text-xl font-semibold mb-4">Connect Wallet</h2>
          <WalletConnect
            onConnect={handleConnect}
            onDisconnect={handleDisconnect}
            connectedAddress={connectedAddress}
          />
        </div>

        <div>
          <h2 className="text-xl font-semibold mb-4">Transfer Details</h2>
          <TransferForm
            onTransfer={handleTransfer}
            loading={loading}
            disabled={!connectedAddress}
            submitLabel="Sign & Estimate Gas"
            loadingLabel="Signing & Estimating..."
          />
        </div>
      </div>

      {pendingTransfer && (
        <div className="mt-6 rounded-lg border border-amber-300 bg-amber-50 p-5">
          <h3 className="text-lg font-semibold text-amber-900">Confirm Gas Quote</h3>
          <p className="mt-2 text-sm text-amber-800">
            The backend has estimated this transfer&apos;s gas cost. The final USDT deduction may vary slightly if on-chain gas price changes before submission.
          </p>
          <div className="mt-4 grid grid-cols-1 gap-3 text-sm md:grid-cols-2">
            <div className="rounded-md bg-white p-3">
              <div className="text-gray-500">Transfer Amount</div>
              <div className="mt-1 font-medium">{pendingTransfer.displayAmount} USDT</div>
            </div>
            <div className="rounded-md bg-white p-3">
              <div className="text-gray-500">Estimated Gas Used</div>
              <div className="mt-1 font-medium">{pendingTransfer.quote.gas_estimate.toLocaleString()}</div>
            </div>
            <div className="rounded-md bg-white p-3">
              <div className="text-gray-500">Gas Price</div>
              <div className="mt-1 font-medium">{formatGwei(pendingTransfer.quote.gas_price)} Gwei</div>
            </div>
            <div className="rounded-md bg-white p-3">
              <div className="text-gray-500">Estimated Gas Cost</div>
              <div className="mt-1 font-medium">{formatUsdt(pendingTransfer.quote.estimated_total_gas_cost)} USDT</div>
            </div>
          </div>
          <div className="mt-3 text-sm text-gray-700">
            Includes paymaster fee: {formatUsdt(pendingTransfer.quote.estimated_paymaster_fee)} USDT
          </div>
          <div className="mt-3 rounded-md bg-white p-3 text-sm">
            <div className="text-gray-500">executeBatch Calldata Hash</div>
            <div className="mt-1 break-all font-mono text-xs text-gray-900">{pendingTransfer.quote.calldata_hash}</div>
          </div>
          <div className="mt-1 text-sm font-medium text-gray-900">
            Estimated total spend: {(Number(pendingTransfer.displayAmount) + Number(formatUsdt(pendingTransfer.quote.estimated_total_gas_cost))).toFixed(6)} USDT
          </div>
          <div className="mt-4 flex flex-col gap-3 sm:flex-row">
            <button
              type="button"
              onClick={handleConfirmTransfer}
              disabled={confirming}
              className={`rounded-md px-4 py-2 text-white ${
                confirming ? 'bg-gray-400 cursor-not-allowed' : 'bg-emerald-600 hover:bg-emerald-700'
              }`}
            >
              {confirming ? 'Sending...' : 'Confirm & Send'}
            </button>
            <button
              type="button"
              onClick={() => setPendingTransfer(null)}
              disabled={confirming}
              className="rounded-md border border-gray-300 bg-white px-4 py-2 text-gray-700 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-60"
            >
              Cancel
            </button>
          </div>
        </div>
      )}

      {result && (
        <div className="mt-6 p-4 bg-gray-100 rounded-lg">
          <h3 className="font-semibold mb-2">Result:</h3>
          <pre className="whitespace-pre-wrap">{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}
    </div>
  )
}
