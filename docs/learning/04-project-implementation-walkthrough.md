# 04. 项目实现拆解

## 项目目标

在 BSC 测试网实现：

1. EOA 升级到 7702 委托执行模式。
2. 通过 4337 执行 token 转账。
3. gas 由 paymaster 垫付，最终用 Mock USDT 结算。
4. 在单笔 UserOp 中执行 `approve + transfer`，并由 Paymaster 代付 gas。

## 合约层（Foundry）

### `MockUSDT.sol`

1. 6 位精度。
2. 支持 `permit`（EIP-2612）。
3. 作为“gas 代币 + 业务转账代币”。

### `EIP7702DelegateAccount.sol`

1. 作为 7702 delegation target。
2. `validateUserOp` 验证 `eth_sign` 风格签名并维护 nonce。
3. `execute/executeBatch` 由 EntryPoint 驱动执行目标调用。

### `USDTTokenPaymaster.sol`

1. `validatePaymasterUserOp` 做长度、报价签名、最低 token 覆盖校验。
2. `postOp` 执行 `transferFrom`，按实际 gas 结算并加 markup。
4. 维护 EntryPoint deposit/stake 与提币管理方法。

## 后端层（Go）

### Paymaster API

1. `/api/v1/paymaster/quote`：计算 `maxCost`、token 报价、有效期并签名 quote。
2. `/api/v1/paymaster/sponsor`：返回 `paymasterAndData`。

### UserOp 转发 API

1. `/api/v1/userop/send`：优先走 `eth_sendUserOperation` 到 bundler。
2. `/api/v1/userop/receipt`：查询 `eth_getUserOperationReceipt`。
3. 无 `BUNDLER_RPC_URL` 时会 fallback 到直发 `EntryPoint.handleOps`。

### 7702 升级 API

1. `/api/v1/7702/upgrade`：后端用私钥构造并发送 `SetCodeTx`。
2. 用于替代当前钱包对 `eip7702Auth` 能力不足的场景。

## 前端层（Next.js + viem）

### 页面三步

1. 连接 MetaMask（尽量避开 Safe 注入干扰）。
2. 调用后端执行 7702 升级。
3. 构建 UserOp（内含 approve+transfer）并发送，轮询回执。

### 前端核心逻辑

1. 构造 `execute(USDT.transfer)` 的 `callData`。
2. 调 quote -> 调 sponsor。
3. 调 `getUserOpHash`，让用户签 UserOp。
4. 发给后端 `/userop/send`，后端转发 bundler。

## 为什么这是“学习版架构”

1. 可观察：每个步骤都有明确 API 和日志。
2. 可替换：quote/sponsor、bundler、paymaster 策略都能独立替换。
3. 可迁移：MockUSDT 可替换成真实稳定币，流程不变。

## 从 Demo 到生产还差什么

1. 更严格的风控和限额（地址白名单、每日上限、黑名单）。
2. 更稳健的报价系统（链上预言机+滑点+熔断）。
3. 更完善的监控与审计（UserOp 生命周期追踪、失败分类、告警）。
4. 安全治理（多签 owner、密钥托管、应急暂停）。
