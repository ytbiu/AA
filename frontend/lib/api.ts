const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL || 'http://localhost:8080'

export interface UserStatus {
  address: string
  is_7702_bound: boolean
  bound_contract: string
  usdt_balance: string
}

export interface FaucetInfo {
  faucet_amount: string
  usdt_address: string
}

export interface Authorize7702Response {
  tx_hash: string
  status: string
  bound_contract: string
}

export interface TransferUSDTResponse {
  tx_hash: string
  status: string
  compensation: string
  gas_used: number
}

export interface RelayerInfo {
  address: string
  pending_tx: number
}

export async function getUserStatus(address: string): Promise<UserStatus> {
  const res = await fetch(`${BACKEND_URL}/api/user-status/${address}`)
  if (!res.ok) {
    throw new Error('Failed to get user status')
  }
  return res.json()
}

export async function getFaucetInfo(): Promise<FaucetInfo> {
  const res = await fetch(`${BACKEND_URL}/api/faucet-info`)
  if (!res.ok) {
    throw new Error('Failed to get faucet info')
  }
  return res.json()
}

// EIP-7702 授权 API
export async function authorize7702(
  userAddress: string,
  chainId: number,
  nonce: number,
  v: number,
  r: string,
  s: string
): Promise<Authorize7702Response> {
  const res = await fetch(`${BACKEND_URL}/api/authorize-7702`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      user_address: userAddress,
      chain_id: chainId,
      nonce: nonce,
      v: v,
      r: r,
      s: s,
      signature: '',
    }),
  })
  if (!res.ok) {
    throw new Error('Failed to authorize 7702')
  }
  return res.json()
}

// 清除 7702 授权 API - 新格式
export async function clear7702(
  userAddress: string,
  chainId: number,
  nonce: number,
  v: number,
  r: string,
  s: string
): Promise<{ tx_hash: string; status: string }> {
  const res = await fetch(`${BACKEND_URL}/api/clear-7702`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      user_address: userAddress,
      chain_id: chainId,
      nonce: nonce,
      v: v,
      r: r,
      s: s,
      signature: '',
    }),
  })
  if (!res.ok) {
    throw new Error('Failed to clear 7702')
  }
  return res.json()
}

export async function transferUSDT(
  userAddress: string,
  targetAddress: string,
  amount: string,
  signature: string
): Promise<TransferUSDTResponse> {
  const res = await fetch(`${BACKEND_URL}/api/transfer-usdt`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      user_address: userAddress,
      target_address: targetAddress,
      amount: amount,
      signature: signature,
    }),
  })
  if (!res.ok) {
    throw new Error('Failed to transfer USDT')
  }
  return res.json()
}

export async function getRelayers(): Promise<{ relayers: RelayerInfo[] }> {
  const res = await fetch(`${BACKEND_URL}/api/admin/relayers`)
  if (!res.ok) {
    throw new Error('Failed to get relayers')
  }
  return res.json()
}

export async function addRelayer(relayerAddress: string): Promise<void> {
  const res = await fetch(`${BACKEND_URL}/api/admin/add-relayer`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ relayer_address: relayerAddress }),
  })
  if (!res.ok) {
    throw new Error('Failed to add relayer')
  }
}

export async function removeRelayer(relayerAddress: string): Promise<void> {
  const res = await fetch(`${BACKEND_URL}/api/admin/remove-relayer`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ relayer_address: relayerAddress }),
  })
  if (!res.ok) {
    throw new Error('Failed to remove relayer')
  }
}

export async function setFeeRate(feeRate: number): Promise<void> {
  const res = await fetch(`${BACKEND_URL}/api/admin/set-fee-rate`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ fee_rate: feeRate }),
  })
  if (!res.ok) {
    throw new Error('Failed to set fee rate')
  }
}

export async function setOracle(oracleAddress: string): Promise<void> {
  const res = await fetch(`${BACKEND_URL}/api/admin/set-oracle`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ oracle_address: oracleAddress }),
  })
  if (!res.ok) {
    throw new Error('Failed to set oracle')
  }
}