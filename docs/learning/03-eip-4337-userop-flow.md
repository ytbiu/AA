# 03. EIP-4337 UserOperation 流程

## 先记住一句话

4337 里上链发交易的人是 bundler，但真正被执行的“账户动作”属于 `UserOperation.sender`。

## 关键角色

1. `UserOperation.sender`：用户账户（本项目里是已 7702 挂载后的 EOA 地址）。
2. `EntryPoint`：统一入口，负责校验和执行。
3. `Bundler`：接收并打包 UserOp，发送链上交易调用 `handleOps`。
4. `Paymaster`：给 gas 垫资，后续再按约定收费（本项目收 Mock USDT）。

## 本项目的一次 UserOp 生命周期

1. 前端构造调用：`execute(USDT.transfer(to, amount))`。
2. 前端请求后端 quote：`/api/v1/paymaster/quote`。
3. 前端请求 sponsor：`/api/v1/paymaster/sponsor`，拿到 `paymasterAndData`。
4. 前端调用 `EntryPoint.getUserOpHash`，用户签名该 hash。
5. 前端把完整 UserOp 发到后端 `/api/v1/userop/send`。
6. 后端转发 `eth_sendUserOperation` 到 bundler（有 bundler URL 时）。
7. bundler 模拟成功后上链 `EntryPoint.handleOps`。
8. EntryPoint 先调账户 `validateUserOp`，再调 paymaster `validatePaymasterUserOp`，最后执行 `callData`。
9. 执行结束后，EntryPoint 调 paymaster `postOp` 做最终结算。

## 为什么你会看到“内部 BNB 转账”

浏览器里常见一条内部交易：`EntryPoint -> bundler beneficiary` 的 BNB 转账。

这代表 EntryPoint 在结算打包方 gas 成本，不是用户主动转 BNB。

## 本项目里的签名

1. UserOp 签名：给账户 `validateUserOp` 验证 UserOp 合法性。
2. `approve` 放在 UserOp 的批处理里执行，不需要额外 Permit 签名。

## 与 AA33 / AA50 报错的关系

1. `AA33 reverted (or OOG)`：常见于 paymaster 验证阶段报价不足、签名不匹配、验证 gas 不合理。
2. `AA50 postOp revert`：常见于 `postOp` 扣款失败（余额不足、allowance 不足、transferFrom 失败）。

## 代码锚点

1. `EntryPoint` 提交流程后端入口：`backend/internal/httpapi/handlers.go` 的 `handleSendUserOp`。
2. 报价和 `paymasterAndData` 组装：`backend/internal/service/paymaster.go`。
3. paymaster 验证与结算：`contracts/src/USDTTokenPaymaster.sol`。
4. 账户验签与执行：`contracts/src/EIP7702DelegateAccount.sol`。
