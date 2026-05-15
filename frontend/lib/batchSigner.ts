import { ethers } from 'ethers'

export interface Call {
  to: string
  data: string
}

export interface BatchOperation {
  user: string
  calls: Call[]
}

export function signBatchOperation(
  privateKey: string,
  batch: BatchOperation
): string {
  const wallet = new ethers.Wallet(privateKey)

  // 按照合约定义的结构编码
  const encoded = ethers.AbiCoder.defaultAbiCoder().encode(
    ['tuple(address user, tuple(address to, bytes data)[] calls)'],
    [batch]
  )

  const hash = ethers.keccak256(encoded)
  const signature = wallet.signMessage(ethers.getBytes(hash))

  return signature
}

export function encodeApproveCall(
  tokenAddress: string,
  spender: string,
  amount: string
): Call {
  const iface = new ethers.Interface([
    'function approve(address spender, uint256 amount) returns (bool)'
  ])
  const data = iface.encodeFunctionData('approve', [spender, amount])

  return {
    to: tokenAddress,
    data: data,
  }
}

export function encodeTransferCall(
  tokenAddress: string,
  to: string,
  amount: string
): Call {
  const iface = new ethers.Interface([
    'function transfer(address to, uint256 amount) returns (bool)'
  ])
  const data = iface.encodeFunctionData('transfer', [to, amount])

  return {
    to: tokenAddress,
    data: data,
  }
}

export function createUSDTTransferBatch(
  userAddress: string,
  usdtAddress: string,
  paymasterAddress: string,
  targetAddress: string,
  amount: string
): BatchOperation {
  const calls: Call[] = [
    encodeApproveCall(usdtAddress, paymasterAddress, amount),
    encodeTransferCall(usdtAddress, targetAddress, amount),
  ]

  return {
    user: userAddress,
    calls: calls,
  }
}