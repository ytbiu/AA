# BSC Testnet AA Demo (EIP-7702 + EIP-4337 + Multi-ERC20 Gas)

这个项目演示以下完整链路：

1. EOA 钱包通过 EIP-7702 升级为智能钱包执行模式
2. 使用 EIP-4337 UserOperation 执行 ERC20 转账
3. Gas 由 Paymaster 先垫付，再用 tUSDT / tUSDC 结算
4. 在单笔 UserOperation 中执行 `approve + transfer`
5. 内置 faucet（任意地址可领取任意数量测试 token）

## 技术栈

- 合约: Foundry + Solidity
- 后端: Go
- 前端: Next.js + wagmi + viem

## 目录结构

- `contracts/`: tUSDT/tUSDC、7702 委托账户逻辑、多 token Paymaster、部署脚本、测试
- `backend/`: Paymaster 报价/签名 API、UserOp 代理提交 API
- `frontend/`: 三步操作页面（连接钱包、7702 升级、4337 转账）
- `scripts/up.sh`: 一键启动后端 + 前端
- `scripts/deploy-bsc.sh`: BSC 测试网部署脚本
- `scripts/preflight-bsc.sh`: BSC 实链参数与合约状态预检脚本

## 合约说明

- `MockUSDT.sol`
  - 可配置名称和符号（用于部署 tUSDT/tUSDC）
  - 6 位精度
  - 开放 faucet mint（测试网用）
- `EIP7702DelegateAccount.sol`
  - 7702 delegation target
  - `validateUserOp` 按 4337 验签
  - `execute` 执行具体调用（例如 USDT.transfer）
- `USDTTokenPaymaster.sol`
  - 验证后端 quote 签名
  - 单 Paymaster 支持多 ERC20 作为 gas token
  - `validatePaymasterUserOp` 做报价与参数校验
  - 在 `postOp` 中执行 `transferFrom` 完成 USDT 结算

## 运行前准备

### 1. 填写环境变量

- 复制并填写：
  - `contracts/.env.example -> contracts/.env`
  - `backend/.env.example -> backend/.env`
  - `frontend/.env.example -> frontend/.env.local`

### 2. 安装依赖

```bash
make install
```

## 本地启动（你要的一键）

```bash
make up
```

启动后：

- 前端: `http://localhost:3000`
- 后端: `http://localhost:8080`

## 部署到 BSC 测试网

1. 配置 `contracts/.env`
2. 执行：

```bash
make deploy-bsc
```

部署后把地址写入：

- `backend/.env`
  - `PAYMASTER_ADDRESS`
- `frontend/.env.local`
  - `NEXT_PUBLIC_PAYMASTER_ADDRESS`
  - `NEXT_PUBLIC_USDT_ADDRESS`
  - `NEXT_PUBLIC_DELEGATE_ACCOUNT_ADDRESS`
  - `NEXT_PUBLIC_ENTRYPOINT_ADDRESS`

## 实链预检（推荐先跑）

```bash
make preflight-bsc
```

会自动检查：

1. 链 ID 是否为 `97`
2. EntryPoint / Paymaster / USDT / DelegateLogic 是否已部署（有代码）
3. 后端 `QUOTE_SIGNER_PRIVATE_KEY` 推导地址是否等于链上 `paymaster.quoteSigner`
4. Paymaster 在 EntryPoint 的 deposit 是否大于 0（若为 0 会警告）

## 页面流程

### 1) 连接 MetaMask

前端通过 wagmi injected connector 连接钱包。

### 2) 升级为智能钱包（EIP-7702）

考虑到部分钱包尚未开放 `eip7702Auth` capability，当前 Demo 由后端使用 `RELAYER_PRIVATE_KEY` 发送 `SetCodeTx` 上链完成升级。

### 3) 转账（4337 + approve + ERC20 支付 gas）

1. 前端构造 UserOperation（`execute(USDT.transfer)`）
2. 调用后端 `/api/v1/paymaster/quote` 获取 token 报价
3. 调用后端 `/api/v1/paymaster/sponsor` 拿到 `paymasterAndData`
4. 钱包签名 `getUserOpHash`
5. 调用后端 `/api/v1/userop/send` 提交到 bundler

注意：

1. 当前方案是单笔 UserOp 内执行 `approve + transfer`。
2. 为避免 UserOp 在 bundler 长时间 pending，前端对 gas 出价采用“至少 1 gwei，再上浮 20%”。
3. 可在前端选择转账 token 和 gas token（tUSDT/tUSDC）。

## API

- `POST /api/v1/paymaster/quote`
- `POST /api/v1/paymaster/sponsor`
- `POST /api/v1/faucet/mint`
- `POST /api/v1/7702/upgrade`
- `POST /api/v1/userop/send`
- `GET /api/v1/userop/receipt?hash=0x...`
- `GET /healthz`

## 无 Bundler 模式

如果你没有 `BUNDLER_RPC_URL`，后端会自动切到直发模式：

1. 用 `RELAYER_PRIVATE_KEY` 签名交易
2. 直接调用 `EntryPoint.handleOps`
3. `userop/receipt` 接口改为查询 `eth_getTransactionReceipt`

## 测试和构建

```bash
make contracts-test
make build
```

## 参考文档

- [EIP-7702](https://eips.ethereum.org/EIPS/eip-7702)
- [EIP-4337](https://eips.ethereum.org/EIPS/eip-4337)
- [Circle Paymaster Overview](https://developers.circle.com/stablecoins/paymaster-overview)
- [Circle Paymaster Quickstart](https://developers.circle.com/paymaster/pay-gas-fees-usdc)

## 流程图

- 时序图: `docs/diagrams/tx-flow-sequence.mmd`
- 架构图: `docs/diagrams/tx-flow-architecture.mmd`

## 学习文档

- 学习手册索引: `docs/learning/README.md`
