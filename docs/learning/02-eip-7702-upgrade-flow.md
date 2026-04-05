# 02. EIP-7702 升级流程（本项目）

## 目标

把用户 EOA 地址升级为“可执行 AA 账户逻辑”的状态，供 4337 后续使用。

## 本项目采用的升级方式

由于钱包对 `eip7702Auth` 支持不稳定，本项目第 2 步采用后端发链上交易：

1. 前端调用 `POST /api/v1/7702/upgrade`
2. 后端使用私钥构造并发送 `SetCodeTx`
3. 账户地址挂载到 `EIP7702DelegateAccount`

## 升级后的关键结果

升级完成后，用户地址的 code 会体现 delegation 前缀（`0xef0100...`）。

你可以用命令检查：

```bash
cast code <USER_ADDRESS> --rpc-url <RPC_URL>
```

## 升级交易里两个 nonce 的含义

后端日志中通常有：

- `txNonce`: 发送 `SetCodeTx` 的交易 nonce
- `authNonce`: 7702 授权元组使用的 nonce

这两者不同是正常现象，属于不同层面的 nonce。

## 升级后参与的阶段

1. 4337 验证阶段：EntryPoint 调 `validateUserOp`
2. 执行阶段：EntryPoint 调 `execute`

也就是说，7702 合约不是“只参与升级一次”，而是后续每次 UserOp 都会被调用。
