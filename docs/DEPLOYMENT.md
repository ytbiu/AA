# AA Wallet 合约部署文档

## 部署信息

**网络:** BSC 测试网 (Chain ID: 97)

**部署时间:** 2026-05-16 (最新更新)

**部署者:** `0x84D98c4faa590cD7cA746E18AcF3bcE8AD61E1b2`

---

## 合约地址

| 合约 | 地址 | 说明 |
|------|------|------|
| MockUSDT | `0x0cF1130E64744860cbA5f992008527485C88F3C8` | 测试 USDT 代币，含水龙头功能 |
| SimplePriceOracle | `0x961E014B856299b1D4B1083F453dE226944f108E` | 固定价格 Oracle (600 USDT/BNB) |
| USDTPaymasterNonProxy | `0x6CEFb6Cf2773565B69FBb23afAAb89292bDb73f2` | **非代理版本** - 避免 EIP-7702 context 问题 |
| Simple7702Account | `0x4057a9Fe6196a636506B9b65456e92CaB39Cb256` | EIP-7702 账户实现合约 |

### 用户账户 (EIP-7702 绑定)

| 用户地址 | 绑定状态 |
|----------|----------|
| `0x8C26fc69D2E5C79b244E1779c9317990f958d152` | 已绑定到 `0x4057a9Fe...` |

---

## 重要发现

### EIP-7702 + UUPS 代理冲突

**问题**: 当 Paymaster 使用 UUPS 代理时，`delegatecall` 会破坏 EIP-7702 的 `address(this)` context，导致：
- `approve` 的 owner 变成 Paymaster 地址（而不是用户账户）
- 用户无法正确授权给 Paymaster

**解决方案**: 部署非代理版本的 Paymaster (`USDTPaymasterNonProxy`)，避免 delegatecall 问题。

---

## 部署脚本

### 1. 部署非代理 Paymaster 和 7702 Account

```bash
forge script script/DeployNonProxy.s.sol:DeployNonProxyScript \
  --rpc-url "https://bnb-testnet.g.alchemy.com/v2/YOUR_API_KEY" \
  --broadcast \
  --private-key "YOUR_PRIVATE_KEY" \
  --via-ir \
  -vv
```

### 2. 部署 SimplePriceOracle

```bash
forge script script/UpdateOracle.s.sol:UpdateOracleScript \
  --rpc-url "https://bnb-testnet.g.alchemy.com/v2/YOUR_API_KEY" \
  --broadcast \
  --private-key "YOUR_PRIVATE_KEY" \
  --via-ir \
  -vv
```

---

## 合约功能

### MockUSDT
- ERC20 代币 (18 decimals)
- `faucet()` - 每次领取 100 USDT
- `mint(address, uint256)` - 管理员铸造

### SimplePriceOracle
- `getBNBPriceInUSDT()` - 返回固定价格 600 USDT/BNB
- 无需 PancakeRouter 依赖，简化部署

### USDTPaymasterNonProxy (非代理)
- `executeBatch(UserOperation, signature)` - 执行批量操作
- `addRelayer(address)` / `removeRelayer(address)` - 管理 Relayer 白名单
- `setFeeRate(uint256)` - 设置手续费率 (10000 = 100%)
- `setFeeRecipient(address)` - 设置手续费接收地址
- `setOracle(address)` - 更新 Oracle 地址

### Simple7702Account (实现合约)
- `executeBatch(Call[])` - 执行批量调用 (最多 5 个)
- `isValidSignature(hash, signature)` - EIP-1271 签名验证
- **注意**: EIP-7702 只绑定代码，不绑定存储

---

## Gasless Transfer 流程

1. **用户签名 UserOperation**
   - 编码: `keccak256(abi.encode(userOp))`
   - 签名: EIP-191 格式 (`signMessage(hash)`)

2. **Paymaster 验证签名**
   - 恢复签名者地址
   - 确认签名者 == userOp.user

3. **执行 EIP-7702 批量操作**
   - 调用 `user.executeBatch(calls)`
   - 内部执行 `approve(Paymaster, amount)`
   - 内部执行 `transfer(target, amount)`

4. **补偿 Relayer**
   - 计算: `gasUsed * gasPrice * BNB价格`
   - 执行: `USDT.transferFrom(user, relayer, compensation)`

---

## 后端配置

将以下地址配置到后端 `.env` 文件：

```bash
BSC_RPC_URL=https://bnb-testnet.g.alchemy.com/v2/YOUR_API_KEY
RELAYER_PRIVATE_KEYS=YOUR_RELAYER_PRIVATE_KEY
CONTRACT_USDT=0x0cF1130E64744860cbA5f992008527485C88F3C8
CONTRACT_PAYMASTER=0x6CEFb6Cf2773565B69FBb23afAAb89292bDb73f2
CONTRACT_ORACLE=0x961E014B856299b1D4B1083F453dE226944f108E
CONTRACT_7702_ACCOUNT=0x4057a9Fe6196a636506B9b65456e92CaB39Cb256
PORT=8080
```

---

## 前端配置

将以下地址配置到前端 `.env.local` 文件：

```bash
NEXT_PUBLIC_BACKEND_URL=http://localhost:8080
NEXT_PUBLIC_USDT_ADDRESS=0x0cF1130E64744860cbA5f992008527485C88F3C8
NEXT_PUBLIC_PAYMASTER_ADDRESS=0x6CEFb6Cf2773565B69FBb23afAAb89292bDb73f2
NEXT_PUBLIC_7702_ACCOUNT_ADDRESS=0x4057a9Fe6196a636506B9b65456e92CaB39Cb256
```

---

## 成功交易记录

| 交易类型 | TX Hash | 状态 |
|----------|---------|------|
| Gasless USDT Transfer | `0x4c9d91e041c62df1f41d40e9f5384638702a58e39d04a3ef0f8b54906d915859` | ✅ 成功 |

**交易详情**:
- Gas Used: 104,819
- 转账金额: 1 USDT
- 补偿金额: 0.03 USDT (给 Relayer)

---

## 验证合约

```bash
# 验证 MockUSDT
forge verify-contract 0x0cF1130E64744860cbA5f992008527485C88F3C8 src/MockUSDT.sol --chain 97

# 验证 SimplePriceOracle
forge verify-contract 0x961E014B856299b1D4B1083F453dE226944f108E src/SimplePriceOracle.sol --chain 97

# 验证 USDTPaymasterNonProxy
forge verify-contract 0x6CEFb6Cf2773565B69FBb23afAAb89292bDb73f2 src/USDTPaymasterNonProxy.sol --chain 97

# 验证 Simple7702Account
forge verify-contract 0x4057a9Fe6196a636506B9b65456e92CaB39Cb256 src/Simple7702Account.sol --chain 97
```

---

## 已废弃的合约地址

以下合约地址已废弃，请勿使用：

| 合约 | 地址 | 原因 |
|------|------|------|
| USDTPaymaster (UUPS) | `0xf516c9C8D1f824Cae05Dfe8b6573E9079189E08B` | EIP-7702 context 问题 |
| PriceOracle | `0x18CC7E9CF8f40dd32Aa0fafD5FfE938B88E455a4` | PancakeRouter 地址无效 |
| Simple7702Account (旧) | `0xf28DF5555822B7FB87d3FABcbA2Af6cf99F48e9D` | 已重新部署 |