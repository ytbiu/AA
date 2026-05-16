'use client'

import EIP7702AuthPanel from '@/components/EIP7702AuthPanel'

export default function ClearPage() {
  return (
    <EIP7702AuthPanel
      mode="clear"
      title="Clear EIP-7702 Authorization"
      description="Clearing the EIP-7702 authorization will remove the delegation to the smart contract validator. This means you will lose smart contract wallet features like batched transactions and meta-transactions, and your account will revert to a standard EOA."
      buttonText="Clear EIP-7702 Delegation"
      buttonColor="red"
    />
  )
}