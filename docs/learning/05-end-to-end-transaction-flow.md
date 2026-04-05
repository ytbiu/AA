# 05. 完整交易流程（本项目）

## 总览

可配合以下图阅读：

1. 时序图：`docs/diagrams/tx-flow-sequence.mmd`
2. 架构图：`docs/diagrams/tx-flow-architecture.mmd`

版本一致性检查（很重要）：

1. 前后端配置的 `PAYMASTER_ADDRESS` 必须一致。
2. 链上 `PAYMASTER_DATA_LENGTH` 必须与当前编码方案一致（approve 方案为 `123`）。

## A. 预备阶段（一次性）

1. 部署 `MockUSDT`、`EIP7702DelegateAccount`、`USDTTokenPaymaster`。
2. 给 paymaster 在 EntryPoint 充值 BNB deposit。
3. 必要时给 paymaster 增加 stake（满足 bundler stake/unstake delay 约束）。
4. 给用户地址发一些 MockUSDT。

## B. 7702 升级阶段

1. 前端调用 `/api/v1/7702/upgrade`。
2. 后端构造 `SetCodeTx` 并发送上链。
3. 升级完成后，用户地址具备 delegation code（指向 delegate account 逻辑）。

## C. 4337 交易阶段

1. 前端构造 `callData`：`executeBatch([approve, transfer])`。
2. 后端 quote：基于 gas 参数给出 tokenAmount 与 quote 签名。
3. 后端 sponsor：返回 `paymasterAndData`。
4. 用户签 `getUserOpHash`。
5. 后端调用 bundler `eth_sendUserOperation`。
6. bundler 打包上链调用 `EntryPoint.handleOps`。

## D. EntryPoint 内部执行阶段

1. 校验账户：`sender.validateUserOp(...)`。
2. 校验 paymaster：`paymaster.validatePaymasterUserOp(...)`。
3. 执行业务调用：`sender.callData`。
4. 结算阶段：`paymaster.postOp(...)`。

## E. 结算阶段（USDT 扣款）

1. 按 `actualGasCost` 换算 USDT + markup。
2. `transferFrom(user -> paymaster)` 收取结算金额。
4. EntryPoint 向 bundler beneficiary 支付 BNB。

## 你在浏览器看到的三类记录

1. 主交易：bundler 发往 EntryPoint。
2. 内部 BNB 转账：EntryPoint 支付 bundler。
3. token 转账：用户给收款地址的业务转账 + 用户给 paymaster 的 gas 结算转账。

## 为什么“返还 gas”看起来可能偏大

1. 前置 quote 是上限估算（含安全系数与 markup）。
2. 结算按 `actualGasCost`，再受最大扣款上限约束。
3. 浏览器展示的“交易费”和“内部结算转账”口径可能不同步。

关键点：以链上 `postOp` 里实际发生的 token 扣款与事件为准。
