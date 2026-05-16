# AA Wallet - EIP-7702 Gasless Account Abstraction

基于 EIP-7702 的账户抽象钱包，运行在 BSC 测试网。用户无需持有 BNB 即可进行 USDT 转账。

## 核心特性

- **Gasless 交易** - 用户无需持有 BNB
- **EIP-7702 授权** - EOA 转换为智能合约账户
- **USDT 支付 Gas** - 使用 USDT 支付交易费用
- **批量操作** - 单次签名执行多个操作

## 项目结构

```
src/           # 智能合约
backend/       # Go 后端服务
frontend/      # Next.js 前端
script/        # 部署脚本
test/          # 合约测试
```

## 快速开始

### 智能合约

```bash
forge build
forge test
forge script script/DeployNonProxy.s.sol --rpc-url $RPC_URL --broadcast
```

### 后端服务

```bash
cd backend
go run cmd/main.go
```

### 前端应用

```bash
cd frontend
npm install && npm run dev
```

## 配置

### 后端 .env

```env
BSC_RPC_URL=https://bnb-testnet.g.alchemy.com/v2/YOUR_KEY
RELAYER_PRIVATE_KEYS=your_key
CONTRACT_USDT=0x0cF1130E64744860cbA5f992008527485C88F3C8
CONTRACT_PAYMASTER=0xa61d461af55029b58d4846c9ea818de9cbc711d3
```

### 前端 .env.local

```env
NEXT_PUBLIC_BACKEND_URL=http://localhost:8080
NEXT_PUBLIC_USDT_ADDRESS=0x0cF1130E64744860cbA5f992008527485C88F3C8
NEXT_PUBLIC_PAYMASTER_ADDRESS=0xa61d461af55029b58d4846c9ea818de9cbc711d3
```

### 部署脚本

```bash
export DEPLOYER_PRIVATE_KEY=your_key
```

## 合约地址 (BSC 测试网)

| 合约 | 地址 |
|------|------|
| MockUSDT | `0x0cF1130E64744860cbA5f992008527485C88F3C8` |
| USDTPaymasterNonProxy | `0x6CEFb6Cf2773565B69FBb23afAAb89292bDb73f2` |
| Simple7702Account | `0x4057a9Fe6196a636506B9b65456e92CaB39Cb256` |

## 使用流程

1. `/faucet` - 领取测试 USDT
2. `/authorize` - EIP-7702 授权
3. `/transfer` - Gasless 转账
4. `/clear` - 清除授权

## 技术栈

- Solidity ^0.8.24, Foundry
- Go 1.22+, Gin
- Next.js 14, TypeScript, ethers.js

## License

MIT