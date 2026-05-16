# AA Wallet - Claude Code Instructions

## Project Overview

EIP-7702 gasless account abstraction wallet on BSC testnet.

## Build Commands

### Smart Contracts
```bash
forge build          # Compile contracts
forge test           # Run tests
forge test -vvv      # Verbose tests
```

### Backend (Go)
```bash
cd backend && go build ./...   # Build
cd backend && go test ./...    # Test
cd backend && go run cmd/main.go  # Run server
```

### Frontend (Next.js)
```bash
cd frontend && npm run build   # Build
cd frontend && npm run dev     # Dev server
```

## Key Files

### Smart Contracts
- `src/base/PaymasterBase.sol` - Abstract base for Paymaster
- `src/USDTPaymaster.sol` - UUPS proxy version
- `src/USDTPaymasterNonProxy.sol` - Non-proxy version (recommended)
- `src/Simple7702Account.sol` - EIP-7702 account logic

### Backend
- `backend/internal/api/handlers.go` - Main API handlers
- `backend/internal/api/handlers_7702.go` - EIP-7702 handlers
- `backend/internal/api/handlers_admin.go` - Admin API
- `backend/internal/relayer/pool.go` - Relayer management

### Frontend
- `frontend/app/authorize/page.tsx` - EIP-7702 authorization
- `frontend/app/transfer/page.tsx` - USDT transfer
- `frontend/lib/signer.ts` - EIP-7702 signing
- `frontend/lib/batchSigner.ts` - Batch operation signing

## Environment Variables

Required for deployment scripts:
```bash
export DEPLOYER_PRIVATE_KEY=0x...
```

## Contract Addresses (BSC Testnet)

```
MockUSDT:           0x0cF1130E64744860cbA5f992008527485C88F3C8
PaymasterNonProxy:  0x6CEFb6Cf2773565B69FBb23afAAb89292bDb73f2
Simple7702Account:  0x4057a9Fe6196a636506B9b65456e92CaB39Cb256
```

## Architecture

```
User → Frontend → Backend API → Relayer Pool → Paymaster Contract → USDT Transfer
```

## Notes

- Use `USDTPaymasterNonProxy` for production (avoids UUPS + EIP-7702 conflict)
- Private keys must be loaded from environment variables
- Gasless flow: user signs, relayer executes, USDT pays gas