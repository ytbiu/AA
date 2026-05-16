export function normalizePrivateKey(privateKey: string): string {
  if (!privateKey.startsWith('0x')) {
    return '0x' + privateKey
  }
  return privateKey
}

export function handleApiError(error: unknown): { error: string } {
  console.error('API Error:', error)
  const message = error instanceof Error ? error.message : 'Unknown error'
  return { error: message }
}

export interface ApiResult {
  tx_hash?: string
  status?: string
  bound_contract?: string
  compensation?: string
  gas_used?: number
  error?: string
}