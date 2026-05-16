'use client'

import { useState, useEffect } from 'react'
import WalletConnect from '@/components/WalletConnect'
import { TransferForm } from '@/components/TransferForm'
import { transferUSDT, getFeeEstimate } from '@/lib/api'
import { createUSDTTransferBatch, signBatchOperationWithWallet } from '@/lib/batchSigner'
import { CONTRACTS } from '@/lib/contracts'
import { ethers } from 'ethers'

export default function TransferPage() {
  const [connectedAddress, setConnectedAddress] = useState<string | null>(null)
  const [signer, setSigner] = useState<ethers.JsonRpcSigner | null>(null)
  const [loading, setLoading] = useState(false)
  const [result, setResult] = useState<any>(null)
  const [feeInfo, setFeeInfo] = useState<any>(null)
  const [showConfirm, setShowConfirm] = useState(false)
  const [pendingTransfer, setPendingTransfer] = useState<{to: string, amount: string} | null>(null)

  useEffect(() => {
    // 获取费用预估
    getFeeEstimate().then(setFeeInfo).catch(console.error)
  }, [])

  const handleConnect = (address: string, walletSigner: ethers.JsonRpcSigner) => {
    setConnectedAddress(address)
    setSigner(walletSigner)
  }

  const handleDisconnect = () => {
    setConnectedAddress(null)
    setSigner(null)
    setResult(null)
  }

  const handleTransfer = async (to: string, amount: string) => {
    if (!connectedAddress || !signer) {
      alert('Please connect your wallet first')
      return
    }

    // 验证金额
    const amountNum = parseFloat(amount)
    if (isNaN(amountNum) || amountNum <= 0) {
      alert('Please enter a valid positive amount')
      return
    }
    if (amountNum > 1000) {
      alert('Amount too large. Maximum is 1000 USDT')
      return
    }

    // 显示确认对话框
    setPendingTransfer({ to, amount })
    setShowConfirm(true)
  }

  const confirmTransfer = async () => {
    if (!pendingTransfer || !connectedAddress || !signer) return

    setShowConfirm(false)
    setLoading(true)
    setResult(null)

    try {
      const { to, amount } = pendingTransfer
      const amountInSmallestUnit = ethers.parseUnits(amount, 18).toString()

      const batch = createUSDTTransferBatch(
        connectedAddress,
        CONTRACTS.USDT,
        CONTRACTS.PAYMASTER,
        to,
        amount
      )

      const signature = await signBatchOperationWithWallet(signer, batch)

      const response = await transferUSDT(connectedAddress, to, amountInSmallestUnit, signature)
      setResult(response)

      // 更新费用预估
      getFeeEstimate().then(setFeeInfo).catch(console.error)
    } catch (error: any) {
      console.error('Error transferring USDT:', error)
      if (error.code === 4001) {
        setResult({ error: 'Signature rejected by user' })
      } else {
        setResult({ error: error.message || 'Failed to transfer USDT' })
      }
    } finally {
      setLoading(false)
      setPendingTransfer(null)
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
        {feeInfo && (
          <p className="mt-2 text-sm text-orange-600">
            <strong>Fee:</strong> ~{feeInfo.total_fee_display} USDT will be deducted as gas compensation
          </p>
        )}
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
          />
        </div>
      </div>

      {/* 确认对话框 */}
      {showConfirm && pendingTransfer && feeInfo && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white p-6 rounded-lg max-w-md">
            <h3 className="text-xl font-bold mb-4">Confirm Transfer</h3>
            <div className="mb-4 space-y-2">
              <p><strong>To:</strong> {pendingTransfer.to}</p>
              <p><strong>Amount:</strong> {pendingTransfer.amount} USDT</p>
              <p><strong>Gas Fee:</strong> ~{feeInfo.total_fee_display} USDT</p>
              <p className="text-sm text-gray-600">
                Total deduction: ~{(parseFloat(pendingTransfer.amount) + parseFloat(feeInfo.total_fee_display || '0')).toFixed(4)} USDT
              </p>
            </div>
            <div className="flex gap-4">
              <button
                onClick={confirmTransfer}
                className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
              >
                Confirm
              </button>
              <button
                onClick={() => { setShowConfirm(false); setPendingTransfer(null); }}
                className="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
              >
                Cancel
              </button>
            </div>
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