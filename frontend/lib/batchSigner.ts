import { ethers } from 'ethers'

export interface Call {
  to: string
  data: string
}

export interface BatchOperation {
  user: string
  calls: Call[]
}

// 确保私钥有 0x 前缀
function normalizePrivateKey(privateKey: string): string {
  if (!privateKey.startsWith('0x')) {
    return '0x' + privateKey
  }
  return privateKey
}

export async function signBatchOperation(
  privateKey: string,
  batch: BatchOperation
): Promise<string> {
  const normalizedKey = normalizePrivateKey(privateKey)
  const wallet = new ethers.Wallet(normalizedKey)

  // 按照合约定义的结构编码 UserOperation
  // UserOperation 是一个 struct: (address user, Call[] calls)
  // 需要作为 tuple 编码，而不是分开的 address 和 array
  //
  // Solidity abi.encode((address, Call[])) 会产生包含偏移量的编码
  // ethers.js 需要 encode 为单个 tuple 类型

  // 正确的方式：将 UserOperation 作为单个 struct/tuple 编码
  // 类型: tuple(address user, tuple(address to, bytes data)[] calls)
  const encoded = ethers.AbiCoder.defaultAbiCoder().encode(
    ['tuple(address user, tuple(address to, bytes data)[] calls)'],
    [{
      user: batch.user,
      calls: batch.calls.map(c => ({ to: c.to, data: c.data }))
    }]
  )

  const hash = ethers.keccak256(encoded)
  // signMessage 会自动添加 EIP-191 前缀，与合约的 toEthSignedMessageHash() 一致
  const signature = await wallet.signMessage(ethers.getBytes(hash))

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
  // Approve 需要包含转账金额 + 补偿金额的缓冲区
  // 补偿约为 gasUsed * gasPrice * BNB价格, 通常 0.02-0.1 USDT
  // 为了安全，多 approve 1 USDT 作为缓冲区
  // USDT 使用 18 decimals
  const transferAmount = ethers.parseUnits(amount, 18)
  const approveAmount = transferAmount + ethers.parseUnits('1', 18)

  const calls: Call[] = [
    encodeApproveCall(usdtAddress, paymasterAddress, approveAmount.toString()),
    encodeTransferCall(usdtAddress, targetAddress, transferAmount.toString()),
  ]

  return {
    user: userAddress,
    calls: calls,
  }
}