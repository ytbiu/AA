'use client'

import EIP7702AuthPanel from '@/components/EIP7702AuthPanel'

export default function AuthorizePage() {
  return (
    <EIP7702AuthPanel
      mode="authorize"
      title="EIP-7702 Authorization"
      description="EIP-7702 allows accounts to temporarily delegate control to a smart contract. This enables traditional EOAs to gain smart contract functionality such as batched transactions, meta-transactions, and more efficient gas usage."
      buttonText="Authorize EIP-7702 Delegation"
      buttonColor="blue"
    />
  )
}