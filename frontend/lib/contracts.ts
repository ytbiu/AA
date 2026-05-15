export const CONTRACTS = {
  USDT: process.env.NEXT_PUBLIC_USDT_ADDRESS || '',
  PAYMASTER: process.env.NEXT_PUBLIC_PAYMASTER_ADDRESS || '',
  7702_ACCOUNT: process.env.NEXT_PUBLIC_7702_ACCOUNT_ADDRESS || '',
}

export function validateContracts() {
  if (!CONTRACTS.USDT || !CONTRACTS.PAYMASTER || !CONTRACTS.7702_ACCOUNT) {
    throw new Error('Contract addresses not configured. Check .env.local')
  }
}