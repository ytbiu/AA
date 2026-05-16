'use client'

import { useState } from 'react'
import PrivateKeyInput from './PrivateKeyInput'
import { authorize7702, clear7702 } from '@/lib/api'
import { signAuthorization, getPrivateKeyAddress } from '@/lib/signer'
import { CONTRACTS } from '@/lib/contracts'
import { BSC_RPC_URL, BSC_CHAIN_ID } from '@/lib/config'
import { JsonRpcProvider } from 'ethers'
import { ApiResult } from '@/lib/utils'

export interface EIP7702AuthPanelProps {
  mode: 'authorize' | 'clear'
  title: string
  description: string
  buttonText: string
  buttonColor: 'blue' | 'red'
}

export default function EIP7702AuthPanel({
  mode,
  title,
  description,
  buttonText,
  buttonColor,
}: EIP7702AuthPanelProps) {
  const [privateKey, setPrivateKey] = useState('')
  const [loading, setLoading] = useState(false)
  const [result, setResult] = useState<ApiResult | null>(null)

  const handleSubmit = async () => {
    if (!privateKey) {
      alert('Please enter your private key')
      return
    }

    setLoading(true)
    setResult(null)

    try {
      const userAddress = getPrivateKeyAddress(privateKey)
      const provider = new JsonRpcProvider(BSC_RPC_URL)
      const accountNonce = await provider.getTransactionCount(userAddress)

      const targetAddress =
        mode === 'authorize' ? CONTRACTS['7702_ACCOUNT'] : '0x0000000000000000000000000000000000000000'

      const authorizationData = {
        chainId: BSC_CHAIN_ID,
        address: targetAddress,
        nonce: accountNonce,
      }

      const { v, r, s } = await signAuthorization(privateKey, authorizationData)

      const response =
        mode === 'authorize'
          ? await authorize7702(userAddress, authorizationData.chainId, authorizationData.nonce, v, r, s)
          : await clear7702(userAddress, authorizationData.chainId, authorizationData.nonce, v, r, s)

      setResult(response)
    } catch (error) {
      setResult({ error: error instanceof Error ? error.message : 'Unknown error' })
    } finally {
      setLoading(false)
    }
  }

  const bgColorClass = buttonColor === 'blue' ? 'bg-blue-100' : 'bg-red-100 border border-red-400'
  const btnBgClass = loading
    ? 'bg-gray-400 cursor-not-allowed'
    : buttonColor === 'blue'
      ? 'bg-blue-500 hover:bg-blue-600 text-white'
      : 'bg-red-500 hover:bg-red-600 text-white'

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">{title}</h1>

      <div className={`mb-6 p-4 ${bgColorClass} rounded-lg`}>
        <h2 className="text-xl font-semibold mb-2">{mode === 'authorize' ? 'About EIP-7702' : 'Warning'}</h2>
        <p>{description}</p>
        <p className="mt-2 text-sm text-gray-600">
          <strong>Note:</strong> Relayer will send the setCode transaction on your behalf. You do not need to hold BNB
          for gas.
        </p>
      </div>

      <div className="mb-6">
        <PrivateKeyInput onSubmit={(pk) => setPrivateKey(pk)} />
      </div>

      <button onClick={handleSubmit} disabled={loading} className={`px-4 py-2 rounded-md ${btnBgClass}`}>
        {loading ? `${buttonText.split(' ')[0]}...` : buttonText}
      </button>

      {result && (
        <div className="mt-6 p-4 bg-gray-100 rounded-lg">
          <h3 className="font-semibold mb-2">Result:</h3>
          <pre className="whitespace-pre-wrap">{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}
    </div>
  )
}