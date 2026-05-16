import { ethers } from 'ethers'
import * as rlp from 'rlp'

// EIP-7702 authorization 签名
export interface EIP7702Authorization {
  chainId: number
  address: string // 要绑定的实现合约地址
  nonce: number
}

// EIP-7702 签名结果
export interface EIP7702Signature {
  v: number
  r: string
  s: string
}

// 确保私钥有 0x 前缀
function normalizePrivateKey(privateKey: string): string {
  if (!privateKey.startsWith('0x')) {
    return '0x' + privateKey
  }
  return privateKey
}

// 签名 EIP-7702 authorization
// EIP-7702 签名格式: hash = keccak256(0x05 || rlp([chainId, address, nonce]))
export async function signAuthorization(
  privateKey: string,
  authorization: EIP7702Authorization
): Promise<EIP7702Signature> {
  const normalizedKey = normalizePrivateKey(privateKey)
  const wallet = new ethers.Wallet(normalizedKey)

  // EIP-7702 签名格式: keccak256(0x05 || rlp([chainId, address, nonce]))
  // chainId 和 nonce 是整数，address 是 20 字节
  // RLP 编码: [chainId (int), address (bytes), nonce (int)]
  const rlpEncoded = rlp.encode([
    authorization.chainId,
    ethers.getBytes(authorization.address),
    authorization.nonce,
  ])

  // 添加 0x05 magic byte 前缀
  const magicByte = new Uint8Array([0x05])
  const authData = ethers.concat([magicByte, rlpEncoded])

  const hash = ethers.keccak256(authData)

  // EIP-7702 需要直接签名哈希，不使用 EIP-191 前缀
  // 使用 SigningKey.sign() 而非 wallet.signMessage()
  const signingKey = new ethers.SigningKey(normalizedKey)
  const sig = signingKey.sign(ethers.getBytes(hash))

  // ethers v6 的 Signature 已经包含 r, s, yParity (v)
  // yParity: 0 或 1，对应 v: 27 或 28
  const v = sig.yParity === 0 ? 27 : 28

  return {
    v: v,
    r: sig.r,
    s: sig.s,
  }
}

export function getPrivateKeyAddress(privateKey: string): string {
  const normalizedKey = normalizePrivateKey(privateKey)
  const wallet = new ethers.Wallet(normalizedKey)
  return wallet.address
}