import { ethers } from 'ethers'

// EIP-7702 authorization 签名
export interface Authorization {
  chainId: number
  address: string // 要绑定的合约地址
  nonce: number
}

export function signAuthorization(
  privateKey: string,
  authorization: Authorization
): string {
  const wallet = new ethers.Wallet(privateKey)

  // EIP-7702 签名格式（简化版本，实际需要按照 EIP 规范）
  const message = ethers.solidityPackedKeccak256(
    ['uint256', 'address', 'uint256'],
    [authorization.chainId, authorization.address, authorization.nonce]
  )

  const signature = wallet.signMessage(ethers.getBytes(message))

  // 注意：签名后应立即清除私钥引用
  return signature
}

export function getPrivateKeyAddress(privateKey: string): string {
  const wallet = new ethers.Wallet(privateKey)
  return wallet.address
}

// 清除私钥（从内存）
export function clearPrivateKey(privateKey: string): void {
  // 在 JavaScript 中无法真正清除内存中的字符串
  // 但可以提醒用户不要保存私钥
  // 实际安全实现需要使用更底层的方式
  console.log('Private key usage complete. Do not save this key.')
}