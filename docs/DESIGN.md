# AA Wallet 方案设计文档

## 1. 项目概述

### 1.1 项目背景

AA Wallet 是一个基于 EIP-7702 的账户抽象钱包解决方案，运行在 BSC 测试网上。用户无需持有原生代币（BNB）即可进行 USDT 转账操作，Gas 费用由 Relayer 代付，并以 USDT 形式从用户账户扣除。

### 1.2 核心特性

- **Gasless 交易**: 用户无需持有 BNB 即可发起交易
- **EIP-7702 授权**: 将 EOA 账户临时转换为智能合约账户
- **USDT 支付 Gas**: 使用 USDT 支付交易费用
- **批量操作**: 支持单次签名执行多个操作
- **Relayer 网络**: 分布式 Relayer 池代付 Gas

### 1.3 技术栈

| 层级 | 技术 |
|------|------|
| 区块链 | BSC 测试网 (Chain ID: 97) |
| 智能合约 | Solidity ^0.8.20, Foundry |
| 后端 | Go 1.22+, Gin, ethers-go |
| 前端 | Next.js 14, TypeScript, Tailwind CSS |
| RPC | Alchemy BSC Testnet API |

---

## 2. 系统架构

### 2.1 架构概览

```
┌─────────────────────────────────────────────────────────────────┐
│                         用户层                                   │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐ │
│  │   Web 前端      │  │   MetaMask      │  │   移动端 App    │ │
│  │  (Next.js)      │  │   (钱包插件)    │  │   (未来扩展)    │ │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │ HTTP REST API
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                         后端服务层                               │
│  ┌───────────────────────────────────────────────────────────┐ │
│  │                    Go Backend (Gin)                        │ │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐       │ │
│  │  │ API Handlers│  │ Relayer Pool│  │ Eth Client  │       │ │
│  │  └─────────────┘  └─────────────┘  └─────────────┘       │ │
│  └───────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │ JSON-RPC / Contract ABI
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                       智能合约层                                 │
│  ┌────────────────┐  ┌────────────────┐  ┌────────────────┐   │
│  │ MockUSDT       │  │ USDTPaymaster  │  │ Simple7702     │   │
│  │ (ERC20 + Faucet)│  │ (Relayer管理) │  │ (Account逻辑) │   │
│  └────────────────┘  └────────────────┘  └────────────────┘   │
│  ┌────────────────┐                                             │
│  │ PriceOracle    │  (BNB/USDT 价格查询)                        │
│  └────────────────┘                                             │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │ BSC Testnet RPC
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                       区块链层                                   │
│                    BSC 测试网 (Chain ID: 97)                     │
└─────────────────────────────────────────────────────────────────┘
```

### 2.2 合约地址 (BSC 测试网)

| 合约 | 地址 | 功能 |
|------|------|------|
| MockUSDT | `0x0cF1130E64744860cbA5f992008527485C88F3C8` | 测试 USDT + 水龙头 |
| PriceOracle | `0x18CC7E9CF8f40dd32Aa0fafD5FfE938B88E455a4` | BNB/USDT 价格 |
| USDTPaymaster | `0xf516c9C8D1f824Cae05Dfe8b6573E9079189E08B` | Relayer + 签名验证 |
| Simple7702Account | `0x9e4e06F875464EEB3aE0AA7993243f910f119Bee` | EIP-7702 账户逻辑 |

---

## 3. 智能合约设计

### 3.1 MockUSDT

**功能**: ERC20 代币 + 水龙头功能

```solidity
contract MockUSDT is ERC20 {
    uint256 public constant FAUCET_AMOUNT = 100 * 10**18;
    
    function faucet() external {
        _mint(msg.sender, FAUCET_AMOUNT);  // 每次领取 100 USDT
    }
    
    function mint(address to, uint256 amount) external {
        _mint(to, amount);  // 管理员可铸造任意数量
    }
}
```

**要点**:
- 18 decimals (与真实 USDT 不同，测试用)
- `faucet()` 无限制领取（仅测试网）
- `mint()` 无权限限制（仅测试网）

### 3.2 USDTPaymaster

**功能**: Relayer 管理、签名验证、Gas 补偿计算

```solidity
contract USDTPaymaster is UUPSUpgradeable {
    IERC20 private _usdtToken;
    IPriceOracle private _oracle;
    uint256 public feeRate;  // 10000 = 100%
    mapping(address => bool) private _relayers;
    
    function executeBatch(UserOperation calldata userOp, bytes calldata signature)
        external onlyRelayer nonReentrant {
        // 1. 验证用户签名
        // 2. 执行批量操作
        // 3. 计算 Gas 补偿 (BNB -> USDT)
        // 4. 从用户账户扣除 USDT
    }
}
```

**关键流程**:
1. Relayer 调用 `executeBatch()`
2. 验证 EIP-7701 签名
3. 执行用户操作
4. 计算 Gas 费用并转换为 USDT
5. 从用户 USDT 余额扣除

### 3.3 Simple7702Account

**功能**: EIP-7702 账户逻辑实现

```solidity
contract Simple7702Account is UUPSUpgradeable, EIP1271 {
    IUSDTPaymaster public paymaster;
    
    function executeBatch(Call[] calldata calls) external onlyPaymaster {
        // 执行批量调用 (最多 5 个)
    }
    
    function isValidSignature(bytes32 hash, bytes calldata signature) 
        external view returns (bytes4) {
        // EIP-1271 签名验证
    }
}
```

**要点**:
- 仅 Paymaster 可调用 `executeBatch()`
- 支持 EIP-1271 签名验证
- UUPS 可升级模式

### 3.4 PriceOracle

**功能**: BNB/USDT 价格查询

```solidity
contract PriceOracle {
    function getBNBPriceInUSDT() external view returns (uint256) {
        // 通过 PancakeRouter 获取实时价格
    }
}
```

---

## 4. 用户交互流程图

### 4.1 新用户完整流程

```
┌─────────────────────────────────────────────────────────────────┐
│                      新用户使用流程                              │
└─────────────────────────────────────────────────────────────────┘

┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│  1. 连接钱包  │────▶│ 2. 查询状态  │────▶│ 3. 领取测试币 │
│  MetaMask    │     │ 用户信息页面 │     │  水龙头页面   │
└──────────────┘     └──────────────┘     └──────────────┘
      │                    │                    │
      │ 输入钱包地址       │ 显示:              │ 输入地址
      │ 或连接插件         │ - USDT 余额        │ 点击领取
      │                    │ - 7702 授权状态    │
      │                    │                    │
      ▼                    ▼                    ▼
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│ 前端验证地址 │     │ 后端查询合约 │     │ Relayer 调用 │
│ 格式正确性   │     │ 余额和状态   │     │ mint() 函数  │
└──────────────┘     └──────────────┘     └──────────────┘
                                                │
                                                │ 返回交易哈希
                                                ▼
                                         ┌──────────────┐
                                         │ 用户获得     │
                                         │ 100 USDT    │
                                         └──────────────┘
                                                │
                                                │ USDT 余额 > 0
                                                ▼
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│ 4. 7702授权  │────▶│ 5. 签名授权  │────▶│ 6. 执行授权  │
│  授权页面    │     │ MetaMask弹窗 │     │ Relayer提交 │
└──────────────┘     └──────────────┘     └──────────────┘
      │                    │                    │
      │ 显示授权信息       │ 用户确认签名       │ 设置授权代码
      │                    │                    │
      ▼                    ▼                    ▼
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│ 构造授权数据 │     │ 生成 EIP-7701│     │ EOA 临时变成 │
│ authorization│     │ 签名消息     │     │ 智能合约账户 │
└──────────────┘     └──────────────┘     └──────────────┘
                                                │
                                                │ 7702 已绑定
                                                ▼
                                         ┌──────────────┐
                                         │ 用户可以进行 │
                                         │ Gasless 转账 │
                                         └──────────────┘
                                                │
                                                ▼
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│ 7. USDT转账  │────▶│ 8. 签名交易  │────▶│ 9. 执行转账  │
│  转账页面    │     │ MetaMask弹窗 │     │ Relayer提交  │
└──────────────┘     └──────────────┘     └──────────────┘
      │                    │                    │
      │ 输入:              │ 用户签名:          │ Paymaster:
      │ - 目标地址         │ - 目标地址         │ - 验证签名
      │ - 转账金额         │ - 金额             │ - 执行转账
      │                    │ - nonce            │ - 计算 Gas
      │                    │                    │ - 扣除 USDT
      ▼                    ▼                    ▼
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│ 前端验证     │     │ 构造批量操作 │     │ 用户 USDT   │
│ 地址和金额   │     │ 数据结构     │     │ 转账成功    │
└──────────────┘     └──────────────┘     └──────────────┘
                                                │
                                                ▼
                                         ┌──────────────┐
                                         │ 交易完成     │
                                         │ Gas 用 USDT │
                                         │ 支付        │
                                         └──────────────┘
```

### 4.2 USDT 转账流程（详细）

```
┌─────────────────────────────────────────────────────────────────┐
│                    USDT Gasless 转账流程                         │
└─────────────────────────────────────────────────────────────────┘

用户                     前端                    后端                区块链
 │                        │                       │                    │
 │  输入转账信息          │                       │                    │
 │  (地址,金额)           │                       │                    │
 │───────────────────────▶│                       │                    │
 │                        │                       │                    │
 │                        │  构造 UserOperation   │                    │
 │                        │  ┌────────────────┐   │                    │
 │                        │  │ user: 用户地址 │   │                    │
 │                        │  │ calls: [{     │   │                    │
 │                        │  │   to: USDT合约 │   │                    │
 │                        │  │   data: transfer│   │                    │
 │                        │  │ }]            │   │                    │
 │                        │  └────────────────┘   │                    │
 │                        │                       │                    │
 │  MetaMask 弹窗签名     │                       │                    │
 │◀───────────────────────│                       │                    │
 │                        │                       │                    │
 │  用户确认签名          │                       │                    │
 │───────────────────────▶│                       │                    │
 │                        │                       │                    │
 │                        │  发送签名请求         │                    │
 │                        │──────────────────────▶│                    │
 │                        │                       │                    │
 │                        │                       │  选择空闲 Relayer  │
 │                        │                       │  ┌────────────┐    │
 │                        │                       │  │ 获取私钥   │    │
 │                        │                       │  │ 构造交易   │    │
 │                        │                       │  └────────────┘    │
 │                        │                       │                    │
 │                        │                       │  调用 Paymaster    │
 │                        │                       │  .executeBatch()   │
 │                        │                       │───────────────────▶│
 │                        │                       │                    │
 │                        │                       │                    │  验证签名
 │                        │                       │                    │  ┌────────┐
 │                        │                       │                    │  │ECDSA   │
 │                        │                       │                    │  │recover │
 │                        │                       │                    │  └────────┘
 │                        │                       │                    │
 │                        │                       │                    │  执行 transfer
 │                        │                       │                    │  ┌────────────┐
 │                        │                       │                    │  │USDT.transfer│
 │                        │                       │                    │  │(to, amount)│
 │                        │                       │                    │  └────────────┘
 │                        │                       │                    │
 │                        │                       │                    │  计算 Gas 补偿
 │                        │                       │                    │  ┌────────────┐
 │                        │                       │                    │  │BNB Gas →   │
 │                        │                       │                    │  │USDT Amount │
 │                        │                       │                    │  └────────────┘
 │                        │                       │                    │
 │                        │                       │                    │  扣除用户 USDT
 │                        │                       │                    │  ┌────────────┐
 │                        │                       │                    │  │USDT.transfer│
 │                        │                       │                    │  │From(user)  │
 │                        │                       │                    │  └────────────┘
 │                        │                       │                    │
 │                        │                       │◀───────────────────│
 │                        │                       │  返回交易哈希       │
 │                        │◀──────────────────────│                    │
 │                        │  返回结果             │                    │
 │◀───────────────────────│                       │                    │
 │  显示交易结果          │                       │                    │
 │                        │                       │                    │
```

### 4.3 EIP-7702 授权流程

```
┌─────────────────────────────────────────────────────────────────┐
│                    EIP-7702 授权流程                             │
└─────────────────────────────────────────────────────────────────┘

用户                     前端                    后端                区块链
 │                        │                       │                    │
 │  点击"7702授权"        │                       │                    │
 │───────────────────────▶│                       │                    │
 │                        │                       │                    │
 │                        │  构造授权数据         │                    │
 │                        │  ┌────────────────┐   │                    │
 │                        │  │ chainId: 97    │   │                    │
 │                        │  │ address: 账户 │   │                    │
 │                        │  │ nonce: 当前值  │   │                    │
 │                        │  │ implementation:│   │                    │
 │                        │  │   Simple7702   │   │                    │
 │                        │  └────────────────┘   │                    │
 │                        │                       │                    │
 │  MetaMask 签名弹窗     │                       │                    │
 │◀───────────────────────│                       │                    │
 │                        │                       │                    │
 │  用户签名确认          │                       │                    │
 │───────────────────────▶│                       │                    │
 │                        │                       │                    │
 │                        │  发送授权请求         │                    │
 │                        │──────────────────────▶│                    │
 │                        │                       │                    │
 │                        │                       │  Relayer 提交      │
 │                        │                       │  setCode 交易      │
 │                        │                       │───────────────────▶│
 │                        │                       │                    │
 │                        │                       │                    │  EOA → 智能合约
 │                        │                       │                    │  ┌────────────┐
 │                        │                       │                    │  │用户地址代码│
 │                        │                       │                    │  │= Simple7702│
 │                        │                       │                    │  └────────────┘
 │                        │                       │                    │
 │                        │                       │◀───────────────────│
 │                        │                       │  返回交易哈希       │
 │                        │◀──────────────────────│                    │
 │                        │  显示授权成功         │                    │
 │◀───────────────────────│                       │                    │
 │  7702 已绑定           │                       │                    │
 │                        │                       │                    │
```

### 4.4 清除 7702 授权流程

```
┌─────────────────────────────────────────────────────────────────┐
│                    清除 7702 授权流程                            │
└─────────────────────────────────────────────────────────────────┘

用户                     前端                    后端                区块链
 │                        │                       │                    │
 │  点击"清除授权"        │                       │                    │
 │───────────────────────▶│                       │                    │
 │                        │                       │                    │
 │                        │  构造清除数据         │                    │
 │                        │  ┌────────────────┐   │                    │
 │                        │  │ chainId: 97    │   │                    │
 │                        │  │ address: 账户 │   │                    │
 │                        │  │ nonce: 当前值  │   │                    │
 │                        │  │ implementation:│   │                    │
 │                        │  │   0x0000...    │   │                    │
 │                        │  └────────────────┘   │                    │
 │                        │                       │                    │
 │  MetaMask 签名弹窗     │                       │                    │
 │◀───────────────────────│                       │                    │
 │                        │                       │                    │
 │  用户签名确认          │                       │                    │
 │───────────────────────▶│                       │                    │
 │                        │                       │                    │
 │                        │  发送清除请求         │                    │
 │                        │──────────────────────▶│                    │
 │                        │                       │                    │
 │                        │                       │  Relayer 提交      │
 │                        │                       │  setCode(空)       │
 │                        │                       │───────────────────▶│
 │                        │                       │                    │
 │                        │                       │                    │  恢复为 EOA
 │                        │                       │                    │  ┌────────────┐
 │                        │                       │                    │  │用户地址代码│
 │                        │                       │                    │  │= 空        │
 │                        │                       │                    │  └────────────┘
 │                        │                       │                    │
 │                        │                       │◀───────────────────│
 │                        │                       │  返回交易哈希       │
 │                        │◀──────────────────────│                    │
 │                        │  显示清除成功         │                    │
 │◀───────────────────────│                       │                    │
 │  7702 已清除           │                       │                    │
 │  恢复为普通 EOA        │                       │                    │
 │                        │                       │                    │
```

---

## 5. API 设计

### 5.1 API 端点列表

| 端点 | 方法 | 功能 | 是否需要 7702 |
|------|------|------|---------------|
| `/api/user-status/:address` | GET | 查询用户状态 | 否 |
| `/api/faucet-info` | GET | 水龙头信息 | 否 |
| `/api/faucet/:address` | POST | 领取测试 USDT | 否 |
| `/api/authorize-7702` | POST | 7702 授权 | 否 (设置后需要) |
| `/api/clear-7702` | POST | 清除 7702 | 需要 7702 |
| `/api/transfer-usdt` | POST | USDT 转账 | 需要 7702 |
| `/api/admin/relayers` | GET | 查询 Relayers | 否 |
| `/api/admin/add-relayer` | POST | 添加 Relayer | 否 |
| `/api/admin/remove-relayer` | POST | 移除 Relayer | 否 |
| `/api/admin/set-fee-rate` | POST | 设置手续费率 | 否 |
| `/api/admin/set-oracle` | POST | 设置 Oracle | 否 |

### 5.2 数据结构

#### UserOperation (批量操作)

```go
type UserOperation struct {
    User    common.Address  // 用户地址
    Calls   []Call          // 批量调用列表
    Nonce   uint64          // 防重放 nonce
}

type Call struct {
    To   common.Address  // 目标合约
    Data []byte          // 调用数据
}
```

#### 请求/响应示例

**USDT 转账请求**:
```json
{
  "user_address": "0x84D98c4faa590cD7cA746E18AcF3bcE8AD61E1b2",
  "target_address": "0x1234567890abcdef1234567890abcdef12345678",
  "amount": "100000000000000000000",
  "signature": "0x..."
}
```

**转账响应**:
```json
{
  "tx_hash": "0x4225989b4eceddc429d69b1e24d5b30e4e591bb5f27de86c15e42db2aeb3af7b",
  "status": "success",
  "compensation": "5000000",
  "gas_used": 85000
}
```

---

## 6. Relayer 池设计

### 6.1 Relayer 选择策略

```go
// 选择策略: 负载均衡
func (p *Pool) SelectIdle() (*Relayer, error) {
    // 1. 按 PendingTx 排序
    // 2. 选择 PendingTx 最少的 Relayer
    // 3. 返回可用的 Relayer
}
```

### 6.2 Relayer 状态管理

| 状态 | PendingTx | 描述 |
|------|-----------|------|
| Idle | 0 | 可用，无待处理交易 |
| Busy | >0 | 有正在处理的交易 |
| Offline | - | 未激活 |

### 6.3 Gas 补偿计算

```solidity
// Gas 费用计算
uint256 bnbCost = gasUsed * tx.gasprice;
uint256 bnbPriceInUsdt = oracle.getBNBPriceInUSDT();
uint256 compensation = (bnbCost * bnbPriceInUsdt) / 1e18;

// 加上手续费
uint256 fee = (compensation * feeRate) / 10000;
uint256 totalCompensation = compensation + fee;
```

---

## 7. 安全设计

### 7.1 签名验证

- **EIP-7701**: 用于 7702 授权签名
- **EIP-1271**: 用于智能合约账户签名验证
- **ECDSA**: 标准 Ethereum 签名算法

### 7.2 重放攻击防护

```solidity
// 使用 nonce 防止重放
bytes32 hash = keccak256(abi.encode(userOp));
// 包含 chainId 和 nonce
```

### 7.3 权限控制

| 操作 | 权限 |
|------|------|
| `executeBatch()` | 仅白名单 Relayer |
| `addRelayer()` | 仅 Owner |
| `setFeeRate()` | 仅 Owner |
| `setOracle()` | 仅 Owner |

### 7.4 升级安全

- 使用 UUPS 代理模式
- `_authorizeUpgrade()` 仅 Owner 可调用
- 合约升级需多签确认（生产环境）

---

## 8. 前端页面设计

### 8.1 页面结构

```
/                   - 首页 (状态查询)
/faucet             - 水龙头 (领取测试币)
/authorize          - 7702 授权
/clear              - 清除授权
/transfer           - USDT 转账
/admin              - 管理页面
```

### 8.2 用户体验流程

1. **首页**: 用户输入地址，查询当前状态
2. **水龙头**: 无余额用户领取测试币
3. **授权**: 用户签名授权 7702
4. **转账**: Gasless USDT 转账
5. **清除**: 恢复为普通 EOA（可选）

---

## 9. 扩展计划

### 9.1 Phase 1 (当前)
- ✅ 基础合约部署
- ✅ 后端 API 实现
- ✅ 前端页面实现
- ✅ 水龙头功能
- 🔄 7702 授权 (待完善)
- 🔄 USDT 转账 (待完善)

### 9.2 Phase 2 (未来)
- ⬜ 多签钱包支持
- ⬜ 批量转账
- ⬜ 交易历史查询
- ⬜ 移动端 App
- ⬜ 多链支持

### 9.3 Phase 3 (生产)
- ⬜ 主网部署
- ⬜ 安全审计
- ⬜ Relayer 分布式网络
- ⬜ 监控和告警

---

## 10. 测试计划

### 10.1 合约测试

```bash
forge test -vvv
```

### 10.2 API 测试

```bash
# 水龙头测试
curl -X POST http://localhost:8080/api/faucet/0x...

# 用户状态查询
curl http://localhost:8080/api/user-status/0x...
```

### 10.3 集成测试

1. 新用户领取测试币
2. 执行 7702 授权
3. 进行 Gasless 转账
4. 清除 7702 授权

---

## 附录

### A. 交易示例

**水龙头领取成功**:
```
Tx Hash: 0x4225989b4eceddc429d69b1e24d5b30e4e591bb5f27de86c15e42db2aeb3af7b
Network: BSC Testnet
Explorer: https://testnet.bscscan.com/tx/0x4225989b4eceddc429d69b1e24d5b30e4e591bb5f27de86c15e42db2aeb3af7b
```

### B. 配置文件

**后端 .env**:
```env
BSC_RPC_URL=https://bnb-testnet.g.alchemy.com/v2/YOUR_KEY
RELAYER_PRIVATE_KEYS=your_private_key
CONTRACT_USDT=0x0cF1130E64744860cbA5f992008527485C88F3C8
CONTRACT_PAYMASTER=0xf516c9C8D1f824Cae05Dfe8b6573E9079189E08B
PORT=8080
```

**前端 .env.local**:
```env
NEXT_PUBLIC_BACKEND_URL=http://localhost:8080
NEXT_PUBLIC_USDT_ADDRESS=0x0cF1130E64744860cbA5f992008527485C88F3C8
```