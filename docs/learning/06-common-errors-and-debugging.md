# 06. 常见错误与排障

## 1. `Unsupported non-optional capabilities: eip7702Auth`

含义：钱包不支持前端直接走 7702 capability。

处理：本项目已改为后端 `/api/v1/7702/upgrade` 发 `SetCodeTx`。

## 2. `Cannot set property ethereum of #<Window> which has only a getter`

含义：尝试覆盖 `window.ethereum` 导致冲突。

处理：不要强改全局 provider，优先用 EIP-6963 选择 MetaMask provider。

## 3. `AA33 reverted (or OOG)`

常见原因：

1. paymaster 报价偏低。
2. quote 签名无效。
3. verification gas 参数不合理。
4. validate 阶段调用了不被 bundler 接受的逻辑。

排查顺序：

1. 先看 `paymasterAndData` 地址是否匹配后端配置。
2. 对比 quote 的 gas 参数与最终 userOp 是否一致。
3. 适度提高 `verificationGasLimit`，但避免效率过低。
4. 检查 paymaster validate 阶段是否触发禁用 opcode。

本项目实战经验：

1. 如果你刚从 Permit 方案切到 approve 方案，务必确认链上 paymaster 也是新版。
2. 可直接检查 `PAYMASTER_DATA_LENGTH`：
   - approve 方案应为 `123`
   - 旧 Permit 方案为 `220`
3. 若前后端编码是 approve（短结构）但链上还是旧 paymaster，几乎必然 AA33。

## 4. `paymaster uses banned opcode: TIMESTAMP`

含义：bundler 在模拟 paymaster 校验时发现禁用 opcode。

说明：当前 approve 方案中，paymaster 不再调用 `permit`，这个问题一般不会再出现。

## 5. `entity stake/unstake delay too low`

含义：paymaster stake 或 unstake delay 不达 bundler 要求。

处理：调用 paymaster 的 `addStakeToEntryPoint(unstakeDelaySec)` 并补足 stake 金额。

## 6. `Verification gas limit efficiency too low`

含义：verificationGasLimit 相对实际使用偏高，被 bundler 拒绝。

处理：降低 `verificationGasLimit` 到合理范围，同时保证不触发 OOG。

本项目当前推荐起步参数：

1. `callGasLimit = 420000`
2. `verificationGasLimit = 110000`
3. `preVerificationGas = 90000`

不同 bundler 会有差异，建议在这个基线附近微调。

## 7. `已提交 hash 但一直轮询不到回执`

现象：

1. `eth_getUserOperationByHash` 能查到对象
2. 但 `blockNumber/transactionHash` 为 `null`
3. `eth_getUserOperationReceipt` 持续 `null`

含义：UserOp 在 bundler mempool，尚未被打包上链。

常见原因与处理：

1. `maxFeePerGas` 太低（最常见）。
2. 用更高 gas 重新发送同 nonce 的 UserOp 让其被优先打包。
3. 前端可设置 gas 下限（例如至少 1 gwei）并适度上浮（例如 +20%）。

## 8. `AA50 postOp revert`

常见原因：

1. 用户余额不足。
2. allowance 不足。
3. `transferFrom` 失败。

排查：

1. 检查用户 USDT 余额是否覆盖“转账金额 + gas 预留”。
2. 检查 `approve` 是否在同一笔 UserOp 内执行成功。
3. 检查 allowance 是否覆盖结算金额上限。

## 实用排障命令

```bash
# 看账户是否已 7702 挂载
cast code <USER_ADDRESS> --rpc-url <RPC_URL>

# 看 paymaster 在 EntryPoint 的 deposit
cast call <ENTRYPOINT> "balanceOf(address)(uint256)" <PAYMASTER> --rpc-url <RPC_URL>

# 看 paymaster 的 quoteSigner
cast call <PAYMASTER> "quoteSigner()(address)" --rpc-url <RPC_URL>

# 看 paymaster 当前编码版本（123=approve版，220=旧permit版）
cast call <PAYMASTER> "PAYMASTER_DATA_LENGTH()(uint256)" --rpc-url <RPC_URL>
```
