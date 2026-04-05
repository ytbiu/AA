# 01. 账户抽象基础

## 什么是账户抽象（AA）

传统 EVM 里，账户分两类：

1. EOA（外部账户）：由私钥控制，发交易时必须是 `tx.from`。
2. 合约账户：有代码，但不能主动发起交易。

账户抽象的目标是让“账户行为可编程”，例如：

- 自定义签名规则（多签、社交恢复）
- 自定义支付 gas 的方式（用 ERC20 代付）
- 批处理、会话密钥、自动化策略

## EIP-4337 解决了什么

4337 没改共识层交易格式，而是在应用层引入：

- `UserOperation`
- `EntryPoint`
- `Bundler`
- `Paymaster`

它把“用户意图”先交给 bundler，再由 bundler 上链调用 EntryPoint 执行。

## EIP-7702 解决了什么

7702 让 EOA 在不完全迁移钱包地址的情况下，临时/可控地“挂载执行逻辑”。

在本项目中，7702 的作用是：

1. 让 EOA 地址表现为可执行 `validateUserOp/execute` 的账户逻辑。
2. 与 4337 组合，完成 AA 路径（签 UserOp、被 EntryPoint 调用）。

## 4337 + 7702 在本项目中的关系

1. 先通过 7702 把用户地址挂到 `EIP7702DelegateAccount`。
2. 再通过 4337 提交 `UserOperation`，由 EntryPoint 调用这个账户逻辑。
3. Paymaster 用 BNB 预付 gas，最后在 `postOp` 中收取 USDT。

## 术语速查

- `tx.from`: 外层链上交易发送者（通常是 bundler）。
- `UserOperation.sender`: 用户账户地址（本项目是已挂 7702 的 EOA）。
- `beneficiary`: EntryPoint 给打包方结算 BNB 的地址。
- `paymasterAndData`: paymaster 的报价、授权、签名等打包数据。
