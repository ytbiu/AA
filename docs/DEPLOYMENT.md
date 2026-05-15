# AA Wallet 合约部署文档

## 部署信息

**网络:** BSC 测试网 (Chain ID: 97)

**部署时间:** 2026-05-15

**部署者:** `0x84D98c4faa590cD7cA746E18AcF3bcE8AD61E1b2`

---

## 合约地址

| 合约 | 地址 | 说明 |
|------|------|------|
| MockUSDT | `0x0cF1130E64744860cbA5f992008527485C88F3C8` | 测试 USDT 代币，含水龙头功能 |
| PriceOracle | `0x18CC7E9CF8f40dd32Aa0fafD5FfE938B88E455a4` | BNB/USDT 价格查询 |
| USDTPaymaster | `0xf516c9C8D1f824Cae05Dfe8b6573E9079189E08B` | Relayer 管理、签名验证、gas 补偿 |
| Simple7702Account | `0x9e4e06F875464EEB3aE0AA7993243f910f119Bee` | EIP-7702 账户逻辑 |

---

## 合约功能

### MockUSDT
- ERC20 代币 (18 decimals)
- `faucet()` - 每次领取 100 USDT
- `mint(address, uint256)` - 管理员铸造

### PriceOracle
- `getBNBPriceInUSDT()` - 获取 BNB 价格
- `setRouter(address)` - 更新 PancakeRouter 地址 (仅 owner)

### USDTPaymaster (UUPS Proxy)
- `executeBatch(UserOperation, signature)` - 执行批量操作
- `addRelayer(address)` / `removeRelayer(address)` - 管理 Relayer 白名单
- `setFeeRate(uint256)` - 设置手续费率 (10000 = 100%)
- `setFeeRecipient(address)` - 设置手续费接收地址
- `setOracle(address)` - 更新 Oracle 地址

### Simple7702Account (UUPS Proxy)
- `executeBatch(Call[])` - 执行批量调用 (最多 5 个)
- `isValidSignature(hash, signature)` - EIP-1271 签名验证
- `setPaymaster(address)` - 更新 Paymaster 地址

---

## 验证合约

```bash
# 验证 MockUSDT
forge verify-contract 0x0cF1130E64744860cbA5f992008527485C88F3C8 src/MockUSDT.sol --chain 97

# 验证 PriceOracle
forge verify-contract 0x18CC7E9CF8f40dd32Aa0fafD5FfE938B88E455a4 src/PriceOracle.sol --chain 97

# 验证 USDTPaymaster
forge verify-contract 0xf516c9C8D1f824Cae05Dfe8b6573E9079189E08B src/USDTPaymaster.sol --chain 97

# 验证 Simple7702Account
forge verify-contract 0x9e4e06F875464EEB3aE0AA7993243f910f119Bee src/Simple7702Account.sol --chain 97
```

---

## 后端配置

将以下地址配置到后端 `.env` 文件：

```
CONTRACT_USDT=0x0cF1130E64744860cbA5f992008527485C88F3C8
CONTRACT_PAYMASTER=0xf516c9C8D1f824Cae05Dfe8b6573E9079189E08B
CONTRACT_ORACLE=0x18CC7E9CF8f40dd32Aa0fafD5FfE938B88E455a4
CONTRACT_7702_ACCOUNT=0x9e4e06F875464EEB3aE0AA7993243f910f119Bee
```

---

## 前端配置

将以下地址配置到前端 `.env.local` 文件：

```
NEXT_PUBLIC_USDT_ADDRESS=0x0cF1130E64744860cbA5f992008527485C88F3C8
NEXT_PUBLIC_PAYMASTER_ADDRESS=0xf516c9C8D1f824Cae05Dfe8b6573E9079189E08B
NEXT_PUBLIC_7702_ACCOUNT_ADDRESS=0x9e4e06F875464EEB3aE0AA7993243f910f119Bee
```