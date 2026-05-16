# AA Wallet 技术设计文档

> 基于 EIP-7702 的 AA 钱包，实现 Gasless 交易和 USDT 支付 Gas

**版本**: v1.0 | **日期**: 2026-05-15 | **网络**: BSC 测试网

---

## 1. 项目概述

### 1.1 核心特性

| 特性 | 描述 |
|------|------|
| **Gasless 交易** | 用户无需持有 BNB 即可发起交易 |
| **EIP-7702 授权** | EOA 账户临时转换为智能合约账户 |
| **USDT 支付 Gas** | 用 USDT 支付交易费用，无需原生代币 |
| **批量操作** | 单次签名执行多个操作（approve + transfer） |
| **Relayer 网络** | 分布式 Relayer 池代付 Gas，负载均衡 |

### 1.2 技术栈

```mermaid
flowchart LR
    subgraph Stack["技术栈"]
        Chain["BSC 测试网<br/>Chain ID: 97"]
        Contract["Solidity ^0.8.20<br/>Foundry"]
        Backend["Go 1.22+<br/>Gin"]
        Frontend["Next.js 14<br/>TypeScript"]
        RPC["Alchemy<br/>BSC API"]
    end

    Chain --> Contract --> Backend --> Frontend
    Backend --> RPC --> Chain

    style Stack fill:#f5f5f5
```

---

## 2. 系统架构

```mermaid
flowchart TB
    subgraph UserLayer["用户层"]
        Web["Web 前端<br/>(Next.js)"]
        MM["MetaMask<br/>(钱包插件)"]
        Mobile["移动端 App<br/>(未来扩展)"]
    end

    subgraph BackendLayer["后端服务层"]
        API["API Handlers<br/>(Gin)"]
        Pool["Relayer Pool<br/>(私钥管理)"]
        Eth["Eth Client<br/>(RPC交互)"]
    end

    subgraph ContractLayer["智能合约层"]
        USDT["MockUSDT<br/>ERC20 + Faucet"]
        Paymaster["USDTPaymaster<br/>Relayer管理 + 签名验证"]
        Account["Simple7702Account<br/>EIP-7702 账户逻辑"]
        Oracle["PriceOracle<br/>BNB/USDT 价格"]
    end

    subgraph ChainLayer["区块链层"]
        BSC["BSC 测试网<br/>Chain ID: 97"]
    end

    UserLayer -->|"HTTP REST"| BackendLayer
    BackendLayer -->|"JSON-RPC"| ContractLayer
    ContractLayer -->|"RPC"| BSC

    style UserLayer fill:#e1f5fe
    style BackendLayer fill:#fff3e0
    style ContractLayer fill:#e8f5e9
    style ChainLayer fill:#f3e5f5
```

### 2.1 合约地址 (BSC 测试网)

```mermaid
flowchart LR
    subgraph Deployed["已部署合约"]
        U["MockUSDT<br/><code>0x0cF1130E64744860<br/>cbA5f99200852748<br/>5C88F3C8</code>"]
        O["PriceOracle<br/><code>0x18CC7E9CF8f40dd3<br/>2Aa0fafD5FfE938B<br/>88E455a4</code>"]
        P["USDTPaymaster<br/><code>0xf516c9C8D1f824C<br/>ae05Dfe8b6573E90<br/>79189E08B</code>"]
        A["Simple7702Account<br/><code>0x9e4e06F875464EE<br/>B3aE0AA7993243f<br/>910f119Bee</code>"]
    end

    style Deployed fill:#c8e6c9
```

---

## 3. 合约设计

### 3.1 合约依赖关系

```mermaid
flowchart TB
    subgraph Order["部署顺序"]
        D1["1. MockUSDT<br/>(无依赖)"]
        D2["2. PriceOracle<br/>(依赖 Router)"]
        D3["3. USDTPaymaster<br/>(依赖 USDT + Oracle)"]
        D4["4. Simple7702Account<br/>(依赖 Paymaster)"]
    end

    Router["PancakeRouter"] --> D2
    WBNB["WBNB地址"] --> D2
    D1 -->|"USDT地址"| D2
    D1 -->|"USDT地址"| D3
    D2 -->|"Oracle地址"| D3
    D3 -->|"Paymaster地址"| D4

    style Order fill:#fff8e1
```

### 3.2 合约调用关系

```mermaid
flowchart TB
    subgraph PaymasterFunc["USDTPaymaster"]
        EB["executeBatch()<br/>核心方法"]
        AR["addRelayer()<br/>管理员"]
        RR["removeRelayer()<br/>管理员"]
        SF["setFeeRate()<br/>管理员"]
        SO["setOracle()<br/>管理员"]
        IR["isRelayer()<br/>查询"]
    end

    subgraph AccountFunc["Simple7702Account"]
        Batch["executeBatch()<br/>批量执行"]
        Valid["isValidSignature()<br/>EIP-1271"]
        SetP["setPaymaster()<br/>管理员"]
    end

    subgraph OracleFunc["PriceOracle"]
        Price["getBNBPriceInUSDT()<br/>价格查询"]
        SetR["setRouter()<br/>管理员"]
    end

    subgraph USDTFunc["MockUSDT"]
        Faucet["faucet()<br/>领取"]
        Mint["mint()<br/>铸造"]
        Trans["transfer()<br/>转账"]
        Approve["approve()<br/>授权"]
        TF["transferFrom()<br/>代扣"]
    end

    EB -->|"验证签名"| Valid
    EB -->|"执行操作"| USDTFunc
    EB -->|"获取汇率"| Price
    EB -->|"补偿"| TF

    AccountFunc -->|"onlyPaymaster"| PaymasterFunc

    style PaymasterFunc fill:#e8f5e9
    style AccountFunc fill:#fff3e0
    style OracleFunc fill:#e3f2fd
    style USDTFunc fill:#fce4ec
```

### 3.3 MockUSDT 合约结构

```mermaid
classDiagram
    class MockUSDT {
        +uint256 FAUCET_AMOUNT = 100e18
        +string name = "Mock USDT"
        +string symbol = "USDT"
        +uint8 decimals = 18
        +faucet() void
        +mint(address, uint256) void
        +transfer(address, uint256) bool
        +approve(address, uint256) bool
        +transferFrom(address, address, uint256) bool
    }
    
    class ERC20 {
        +totalSupply() uint256
        +balanceOf(address) uint256
        +allowance(address, address) uint256
    }
    
    MockUSDT --|> ERC20 : inherits
```

### 3.4 USDTPaymaster 合约结构

```mermaid
classDiagram
    class USDTPaymaster {
        +IERC20 _usdtToken
        +IPriceOracle _oracle
        +uint256 feeRate
        +address feeRecipient
        +mapping relayers
        +executeBatch(UserOp, sig) void
        +addRelayer(address) void
        +removeRelayer(address) void
        +setFeeRate(uint256) void
        +setFeeRecipient(address) void
        +setOracle(address) void
        +isRelayer(address) bool
    }
    
    class UUPSUpgradeable {
        +upgradeTo(address) void
        +_authorizeUpgrade(address) void
    }
    
    class ReentrancyGuard {
        +nonReentrant modifier
    }
    
    class OwnableUpgradeable {
        +onlyOwner modifier
    }
    
    USDTPaymaster --|> UUPSUpgradeable
    USDTPaymaster --|> ReentrancyGuard  
    USDTPaymaster --|> OwnableUpgradeable
```

### 3.5 Simple7702Account 合约结构

```mermaid
classDiagram
    class Simple7702Account {
        +IUSDTPaymaster paymaster
        +uint256 MAX_BATCH_SIZE = 5
        +initialize(address, address) void
        +executeBatch(Call[]) void
        +isValidSignature(hash, sig) bytes4
        +setPaymaster(address) void
        +recoverSigner(hash, sig) address
    }
    
    class UUPSUpgradeable {
        +upgradeTo(address) void
    }
    
    class OwnableUpgradeable {
        +onlyOwner modifier
    }
    
    class EIP1271 {
        +isValidSignature(hash, sig) bytes4
    }
    
    Simple7702Account --|> UUPSUpgradeable
    Simple7702Account --|> OwnableUpgradeable
    Simple7702Account ..|> EIP1271 : implements
```

---

## 4. 用户交互流程

### 4.1 新用户完整流程

```mermaid
flowchart TB
    subgraph Step1["步骤1: 准备"]
        Connect["连接钱包<br/>或输入地址"]
        Query["查询状态<br/>余额 + 7702状态"]
        Faucet["领取测试币<br/>100 USDT"]
    end

    subgraph Step2["步骤2: 授权"]
        Auth["7702授权页面"]
        SignAuth["MetaMask签名<br/>authorization"]
        ExecAuth["Relayer提交<br/>setCode交易"]
        Bound["EOA绑定合约代码"]
    end

    subgraph Step3["步骤3: 使用"]
        Transfer["USDT转账页面"]
        SignTx["签名batch操作<br/>approve + transfer"]
        ExecTx["Relayer提交<br/>executeBatch"]
        Success["转账成功<br/>USDT支付Gas"]
    end

    subgraph Step4["步骤4: 清除(可选)"]
        Clear["清除7702页面"]
        SignClear["签名清除authorization"]
        ExecClear["Relayer提交<br/>setCode(空)"]
        Restore["恢复为普通EOA"]
    end

    Connect --> Query --> Faucet
    Faucet -->|"余额>0"| Auth
    Auth --> SignAuth --> ExecAuth --> Bound
    Bound --> Transfer
    Transfer --> SignTx --> ExecTx --> Success
    Success -->|"可选"| Clear
    Clear --> SignClear --> ExecClear --> Restore

    style Step1 fill:#e1f5fe
    style Step2 fill:#fff3e0
    style Step3 fill:#e8f5e9
    style Step4 fill:#fce4ec
```

### 4.2 EIP-7702 授权流程

```mermaid
sequenceDiagram
    participant U as 用户
    participant F as 前端
    participant B as 后端
    participant R as Relayer
    participant C as 区块链

    U->>F: 点击"7702授权"
    
    rect rgb(230, 245, 255)
        Note over F: 构造authorization
        F->>F: chainId = 97
        F->>F: address = 用户EOA
        F->>F: nonce = 当前值
        F->>F: implementation = Simple7702Account
    end
    
    F->>U: MetaMask签名弹窗
    U->>F: 确认签名
    
    F->>B: POST /api/authorize-7702<br/>{signedAuthorization}
    
    B->>B: 选择空闲Relayer
    B->>R: 构建setCode交易
    
    R->>C: 提交7702授权交易
    
    rect rgb(200, 230, 200)
        Note over C: EOA → 智能合约账户
        C->>C: 用户地址.code = Simple7702Account
    end
    
    C-->>R: 交易确认
    R-->>B: txHash
    B-->>F: {txHash, status}
    F-->>U: 显示授权成功
```

### 4.3 USDT Gasless 转账流程

```mermaid
sequenceDiagram
    participant U as 用户EOA(7702)
    participant F as 前端
    participant B as 后端
    participant R as Relayer
    participant P as Paymaster
    participant O as Oracle
    participant T as USDT合约

    U->>F: 输入目标地址 + 金额
    
    rect rgb(230, 245, 255)
        Note over F: 构造UserOperation
        F->>F: user = 用户地址
        F->>F: calls = [<br/>  approve(Paymaster, amount),<br/>  transfer(target, amount)<br/>]
        F->>F: nonce = 当前值
    end
    
    F->>U: MetaMask签名batch内容
    U->>F: 返回签名
    
    F->>B: POST /api/transfer-usdt<br/>{userOp, signature}
    
    B->>B: 选择空闲Relayer (pending最少)
    B->>R: 构建executeBatch交易
    
    R->>P: executeBatch(userOp, sig)
    
    rect rgb(255, 243, 224)
        Note over P: 合约验证层
        P->>P: 1. 检查Relayer白名单 ✓
        P->>P: 2. ECDSA恢复signer
        P->>P: 3. 验证signer == user ✓
    end
    
    rect rgb(200, 230, 200)
        Note over P: 执行层
        P->>T: 4. approve(Paymaster, amount)
        P->>T: 5. transfer(target, amount)
        P->>P: 6. 记录gasUsed
    end
    
    rect rgb(252, 228, 236)
        Note over P: 补偿层
        P->>O: 7. getBNBPriceInUSDT()
        O-->>P: 返回汇率
        P->>P: 8. 计算USDT补偿金额
        P->>T: 9. transferFrom(user → Relayer)
        P->>T: 10. transferFrom(user → feeRecipient)
    end
    
    P-->>R: 交易完成
    R-->>B: {txHash, gasUsed, compensation}
    B-->>F: 返回结果
    F-->>U: 显示转账成功<br/>Gas用USDT支付
```

### 4.4 清除 7702 授权流程

```mermaid
sequenceDiagram
    participant U as 用户
    participant F as 前端
    participant B as 后端
    participant R as Relayer
    participant C as 区块链

    U->>F: 点击"清除授权"
    
    rect rgb(230, 245, 255)
        Note over F: 构造清除authorization
        F->>F: chainId = 97
        F->>F: address = 用户EOA
        F->>F: implementation = 0x0000...<br/>(空地址)
    end
    
    F->>U: MetaMask签名弹窗
    U->>F: 确认签名
    
    F->>B: POST /api/clear-7702<br/>{signedAuthorization}
    
    B->>R: 选择Relayer
    R->>C: 提交setCode(空)交易
    
    rect rgb(200, 230, 200)
        Note over C: 恢复为EOA
        C->>C: 用户地址.code = 空
    end
    
    C-->>R: 交易确认
    R-->>B: txHash
    B-->>F: {txHash}
    F-->>U: 7702已清除<br/>恢复为普通EOA
```

---

## 5. Gas补偿计算

### 5.1 补偿计算流程

```mermaid
flowchart TB
    subgraph Input["输入参数"]
        GasUsed["gasUsed<br/>(交易消耗gas)"]
        GasPrice["tx.gasprice<br/>(当前gas价格)"]
        BNBPrice["BNB/USDT汇率<br/>(Oracle查询)"]
        FeeRate["feeRate<br/>(手续费率 10000=100%)"]
    end

    subgraph Calc["计算过程"]
        BNBCost["BNB成本<br/>= gasUsed × gasPrice"]
        USDTComp["USDT补偿<br/>= BNB成本 × BNB价格 ÷ 1e18"]
        Fee["手续费<br/>= USDT补偿 × feeRate ÷ 10000"]
        Total["总扣除<br/>= USDT补偿 + 手续费"]
    end

    subgraph Output["USDT转账"]
        ToRelayer["→ Relayer<br/>(USDT补偿)"]
        ToFee["→ feeRecipient<br/>(手续费)"]
    end

    GasUsed --> BNBCost
    GasPrice --> BNBCost
    BNBCost --> USDTComp
    BNBPrice --> USDTComp
    USDTComp --> Fee
    FeeRate --> Fee
    USDTComp --> ToRelayer
    Fee --> ToFee

    style Input fill:#e3f2fd
    style Calc fill:#fff8e1
    style Output fill:#e8f5e9
```

### 5.2 计算公式

```mermaid
block-beta
    columns 3
    
    block:Input["输入"]
        A["gasUsed"]
        B["gasPrice"]
        C["BNB/USDT"]
    end
    
    block:Calc["计算"]
        D["BNB成本<br/>= A × B"]
        E["USDT补偿<br/>= D × C ÷ 1e18"]
        F["手续费<br/>= E × feeRate ÷ 10000"]
    end
    
    block:Output["输出"]
        G["→ Relayer"]
        H["→ feeRecipient"]
    end
    
    Input --> Calc --> Output
```

---

## 6. Relayer池设计

### 6.1 Relayer选择策略

```mermaid
flowchart TB
    Request["交易请求到达"]
    
    Check["查询所有Relayer<br/>pending状态"]
    
    subgraph Pool["Relayer池状态"]
        R1["Relayer1<br/>pending: 0<br/>✓ 空闲"]
        R2["Relayer2<br/>pending: 2<br/>✗ 繁忙"]
        R3["Relayer3<br/>pending: 1<br/>✗ 繁忙"]
    end
    
    Select["选择pending最少<br/>的Relayer"]
    
    Mark["标记pending += 1"]
    
    Submit["提交交易"]
    
    Wait["等待交易确认"]
    
    Complete["交易完成<br/>pending -= 1"]
    
    Request --> Check --> Pool
    Pool -->|"pending=0"| Select
    Select --> Mark --> Submit --> Wait --> Complete

    style Pool fill:#fce4ec
    style Select fill:#c8e6c9
```

### 6.2 Relayer状态

```mermaid
stateDiagram-v2
    [*] --> Idle: Relayer初始化
    
    Idle --> Busy: 接收交易<br/>pending += 1
    Busy --> Idle: 交易完成<br/>pending -= 1
    
    Busy --> Busy: 接收更多交易<br/>pending持续增加
    
    note right of Idle: 可接收新交易
    note right of Busy: 正在处理交易
```

### 6.3 Relayer数据结构

```mermaid
classDiagram
    class Relayer {
        +Address address
        +PrivateKey privateKey
        +int PendingTx
        +int64 LastUsed
    }
    
    class Pool {
        +Relayer[] relayers
        +sync.RWMutex mu
        +SelectIdle() Relayer
        +MarkPending(address) void
        +MarkComplete(address) void
        +GetAll() RelayerInfo[]
        +GetCount() int
    }
    
    Pool "1" --> "*" Relayer : manages
```

---

## 7. API设计

### 7.1 API端点总览

```mermaid
flowchart LR
    subgraph UserAPI["用户API"]
        U1["GET /api/user-status/:addr"]
        U2["GET /api/faucet-info"]
        U3["POST /api/faucet/:addr"]
        U4["POST /api/authorize-7702"]
        U5["POST /api/clear-7702"]
        U6["POST /api/transfer-usdt"]
    end

    subgraph AdminAPI["管理API"]
        A1["GET /api/admin/relayers"]
        A2["POST /api/admin/add-relayer"]
        A3["POST /api/admin/remove-relayer"]
        A4["POST /api/admin/set-fee-rate"]
        A5["POST /api/admin/set-oracle"]
    end

    style UserAPI fill:#e1f5fe
    style AdminAPI fill:#fff3e0
```

### 7.2 API与合约映射

```mermaid
flowchart TB
    subgraph APIs["API端点"]
        FaucetAPI["POST /api/faucet/:addr"]
        AuthAPI["POST /api/authorize-7702"]
        ClearAPI["POST /api/clear-7702"]
        TransferAPI["POST /api/transfer-usdt"]
        StatusAPI["GET /api/user-status/:addr"]
    end

    subgraph Operations["合约操作"]
        Mint["USDT.mint()"]
        SetCodeAuth["setCode(绑定Simple7702)"]
        SetCodeClear["setCode(空地址)"]
        ExecBatch["Paymaster.executeBatch()"]
        Query["查询余额+7702状态"]
    end

    FaucetAPI --> Mint
    AuthAPI --> SetCodeAuth
    ClearAPI --> SetCodeClear
    TransferAPI --> ExecBatch
    StatusAPI --> Query

    style APIs fill:#e1f5fe
    style Operations fill:#e8f5e9
```

### 7.3 数据结构

```mermaid
classDiagram
    class UserOperation {
        +Address user
        +Call[] calls
        +uint64 nonce
    }
    
    class Call {
        +Address to
        +bytes data
    }
    
    class Authorize7702Request {
        +string authorization_data
        +bytes signature
    }
    
    class TransferUSDTRequest {
        +string user_address
        +string target_address
        +string amount
        +bytes signature
    }
    
    class TransferUSDTResponse {
        +string tx_hash
        +string status
        +string compensation
        +int gas_used
    }
    
    class UserStatusResponse {
        +string address
        +bool is7702_bound
        +string bound_contract
        +string usdt_balance
    }
    
    class RelayerInfo {
        +string address
        +int pending_tx
    }

    UserOperation "1" --> "*" Call : contains
```

### 7.4 请求响应示例

**USDT转账请求:**
```json
{
  "user_address": "0x84D98c4faa590cD7cA746E18AcF3bcE8AD61E1b2",
  "target_address": "0x1234567890abcdef1234567890abcdef12345678",
  "amount": "100000000000000000000",
  "signature": "0x..."
}
```

**转账响应:**
```json
{
  "tx_hash": "0x4225989b4eceddc429d69b1e24d5b30e4e591bb5f27de86c15e42db2aeb3af7b",
  "status": "success",
  "compensation": "5000000",
  "gas_used": 85000
}
```

---

## 8. 安全设计

### 8.1 多层验证机制

```mermaid
flowchart TB
    subgraph Input["executeBatch输入"]
        UserOp["UserOperation"]
        Sig["用户签名"]
        Caller["调用者(Relayer)"]
    end

    subgraph Layer1["第一层: Relayer验证"]
        V1["检查Relayer白名单"]
        V1Fail["❌ NotRelayer错误"]
    end

    subgraph Layer2["第二层: 签名验证"]
        V2["ECDSA恢复signer"]
        V3["验证signer == user"]
        V2Fail["❌ InvalidSignature错误"]
    end

    subgraph Layer3["第三层: 执行验证"]
        E1["遍历执行calls"]
        E2["记录gasUsed"]
        E1Fail["❌ CallFailed错误"]
    end

    subgraph Layer4["第四层: 补偿验证"]
        C1["计算补偿金额"]
        C2["执行USDT转账"]
        C2Fail["❌ TransferFailed错误"]
    end

    Success["✓ 交易完成"]

    Input --> V1
    V1 -->|"失败"| V1Fail
    V1 -->|"成功"| V2
    V2 --> V3
    V3 -->|"失败"| V2Fail
    V3 -->|"成功"| E1
    E1 -->|"失败"| E1Fail
    E1 -->|"成功"| E2 --> C1 --> C2
    C2 -->|"失败"| C2Fail
    C2 -->|"成功"| Success

    style Input fill:#e3f2fd
    style Layer1 fill:#fff8e1
    style Layer2 fill:#fff8e1
    style Layer3 fill:#fff8e1
    style Layer4 fill:#fff8e1
    style V1Fail fill:#ffcdd2
    style V2Fail fill:#ffcdd2
    style E1Fail fill:#ffcdd2
    style C2Fail fill:#ffcdd2
    style Success fill:#c8e6c9
```

### 8.2 权限控制矩阵

```mermaid
flowchart TB
    subgraph Functions["合约方法"]
        EB["executeBatch()"]
        AR["addRelayer()"]
        RR["removeRelayer()"]
        SF["setFeeRate()"]
        SO["setOracle()"]
        SFR["setFeeRecipient()"]
        UP["upgradeTo()"]
    end

    subgraph Roles["权限角色"]
        Relayer["Relayer白名单"]
        Owner["合约Owner"]
    end

    EB -->|"onlyRelayer"| Relayer
    AR -->|"onlyOwner"| Owner
    RR -->|"onlyOwner"| Owner
    SF -->|"onlyOwner"| Owner
    SO -->|"onlyOwner"| Owner
    SFR -->|"onlyOwner"| Owner
    UP -->|"onlyOwner"| Owner

    style Functions fill:#e8f5e9
    style Roles fill:#fff3e0
```

### 8.3 防护措施

```mermaid
mindmap
  root((安全防护))
    签名验证
      EIP-7701 authorization
      EIP-1271 智能合约签名
      ECDSA 恢复验证
    重放攻击
      chainId 包含在hash
      nonce 防止重复
      签名唯一性
    权限控制
      Relayer白名单
      Owner管理权限
      UUPS升级保护
    重入攻击
      ReentrancyGuard
      nonReentrant修饰符
    升级安全
      _authorizeUpgrade
      仅Owner可升级
```

---

## 9. 前端页面设计

### 9.1 页面结构

```mermaid
flowchart TB
    subgraph Pages["页面路由"]
        Home["/ 首页<br/>用户状态"]
        Faucet["/faucet<br/>水龙头"]
        Auth["/authorize<br/>7702授权"]
        Clear["/clear<br/>清除授权"]
        Transfer["/transfer<br/>USDT转账"]
        Admin["/admin<br/>管理页面"]
    end

    Home -->|"领取USDT"| Faucet
    Home -->|"授权钱包"| Auth
    Home -->|"清除授权"| Clear
    Home -->|"转账USDT"| Transfer
    Home -->|"管理员"| Admin

    style Pages fill:#e1f5fe
```

### 9.2 首页状态展示

```mermaid
flowchart LR
    subgraph HomePage["首页组件"]
        AddrInput["地址输入框"]
        QueryBtn["查询按钮"]
        
        subgraph Status["状态展示"]
            BoundBadge["7702绑定状态<br/>已绑定/未绑定"]
            USDTBal["USDT余额<br/>显示数值"]
        end
        
        subgraph Actions["快捷操作"]
            FaucetBtn["水龙头"]
            AuthBtn["授权"]
            ClearBtn["清除"]
            TransferBtn["转账"]
        end
    end

    AddrInput --> QueryBtn --> Status --> Actions

    style HomePage fill:#e1f5fe
    style Status fill:#c8e6c9
    style Actions fill:#fff3e0
```

### 9.3 转账页面流程

```mermaid
flowchart TB
    subgraph TransferPage["转账页面"]
        PrivKey["私钥输入<br/>(本地使用)"]
        Target["目标地址"]
        Amount["转账金额"]
        
        subgraph Preview["预估信息"]
            EstGas["预估Gas"]
            EstComp["预估USDT补偿"]
            EstFee["预估手续费"]
        end
        
        SignBtn["签名按钮"]
        SubmitBtn["提交按钮"]
        Result["交易结果"]
    end

    PrivKey --> Target --> Amount --> Preview --> SignBtn --> SubmitBtn --> Result

    style TransferPage fill:#e1f5fe
    style Preview fill:#fff8e1
```

---

## 10. 数据流图

```mermaid
flowchart LR
    subgraph UserInput["用户输入"]
        Priv["私钥(本地)"]
        Addr["地址"]
        Amnt["金额"]
    end

    subgraph Frontend["前端处理"]
        BuildOp["构建UserOperation"]
        SignOp["签名操作"]
        SignAuth["签名authorization"]
    end

    subgraph Backend["后端处理"]
        SelectR["选择Relayer"]
        BuildTx["构建交易"]
        Monitor["监控状态"]
    end

    subgraph Blockchain["链上执行"]
        Verify["验证签名"]
        Exec["执行calls"]
        Comp["计算补偿"]
        Transfer["USDT转账"]
    end

    subgraph Output["输出"]
        TxHash["交易哈希"]
        GasUsed["Gas消耗"]
        USDTComp["USDT补偿"]
    end

    Priv --> SignOp --> Backend
    Priv --> SignAuth --> Backend
    Addr --> BuildOp --> Backend
    Amnt --> BuildOp --> Backend
    
    Frontend -->|"HTTP"| Backend
    Backend -->|"RPC"| Blockchain
    Blockchain --> Output

    style UserInput fill:#e1f5fe
    style Frontend fill:#fff3e0
    style Backend fill:#f3e5f5
    style Blockchain fill:#e8f5e9
    style Output fill:#c8e6c9
```

---

## 11. 用户7702生命周期

```mermaid
stateDiagram-v2
    [*] --> EOA: 用户创建EOA

    EOA --> AuthPending: 提交7702授权
    AuthPending --> Auth7702: 交易确认
    AuthPending --> EOA: 交易失败

    state Auth7702 {
        [*] --> Bound
        Bound --> Executing: 执行batch操作
        Executing --> Bound: 操作完成
    }
    
    Auth7702 --> TransferReady: 可执行Gasless转账
    TransferReady --> TransferReady: 转账成功
    TransferReady --> TransferReady: 转账失败(Retry)

    Auth7702 --> ClearPending: 提交清除
    ClearPending --> EOA: 清除成功
    ClearPending --> Auth7702: 清除失败

    note right of EOA: 纯EOA状态<br/>需要BNB支付gas
    note right of Auth7702: 已绑定合约<br/>可用USDT支付gas
    note right of TransferReady: 无需BNB<br/>USDT支付Gas
```

---

## 12. 后端架构

```mermaid
flowchart TB
    subgraph Backend["Go后端"]
        subgraph API["API层"]
            Handlers["Handlers<br/>(Gin)"]
            Models["Models<br/>(请求/响应)"]
        end
        
        subgraph Business["业务层"]
            Pool["Relayer Pool<br/>(选择策略)"]
            Selector["Selector<br/>(负载均衡)"]
        end
        
        subgraph Contract["合约层"]
            PaymasterC["Paymaster<br/>(合约交互)"]
            USDTC["USDT<br/>(合约交互)"]
            OracleC["Oracle<br/>(合约交互)"]
        end
        
        subgraph Eth["以太坊层"]
            Client["Eth Client<br/>(RPC)"]
            TxBuilder["Tx Builder<br/>(交易构建)"]
        end
    end

    API --> Business --> Contract --> Eth

    style Backend fill:#fff3e0
    style API fill:#e1f5fe
    style Business fill:#f3e5f5
    style Contract fill:#e8f5e9
    style Eth fill:#c8e6c9
```

---

## 13. 扩展计划

### 13.1 开发阶段

```mermaid
timeline
    title AA Wallet 开发计划
    
    section Phase 1 (当前)
        基础合约部署 : ✅ 完成
        后端API实现 : ✅ 完成
        前端页面 : ✅ 完成
        水龙头功能 : ✅ 完成
        7702授权 : 🔄 待完善
        USDT转账 : 🔄 待完善
    
    section Phase 2 (未来)
        多签钱包 : ⬜ 计划中
        批量转账 : ⬜ 计划中
        交易历史 : ⬜ 计划中
        移动端App : ⬜ 计划中
        多链支持 : ⬜ 计划中
    
    section Phase 3 (生产)
        主网部署 : ⬜ 计划中
        安全审计 : ⬜ 计划中
        Relayer网络 : ⬜ 计划中
        监控告警 : ⬜ 计划中
```

### 13.2 功能扩展路线

```mermaid
flowchart LR
    subgraph Current["Phase 1: MVP"]
        C1["单账户7702"]
        C2["USDT支付Gas"]
        C3["基础水龙头"]
    end
    
    subgraph Future["Phase 2: 增强"]
        F1["批量转账"]
        F2["交易历史"]
        F3["多签支持"]
    end
    
    subgraph Production["Phase 3: 生产"]
        P1["主网部署"]
        P2["安全审计"]
        P3["分布式Relayer"]
    end
    
    Current --> Future --> Production

    style Current fill:#c8e6c9
    style Future fill:#fff8e1
    style Production fill:#e3f2fd
```

---

## 14. 配置文件

### 14.1 后端配置

```mermaid
flowchart LR
    subgraph Env["环境变量"]
        RPC["BSC_RPC_URL"]
        Keys["RELAYER_PRIVATE_KEYS<br/>(逗号分隔)"]
        USDT["CONTRACT_USDT"]
        Paymaster["CONTRACT_PAYMASTER"]
        Oracle["CONTRACT_ORACLE"]
        Port["PORT=8080"]
    end

    style Env fill:#fff8e1
```

**.env 示例:**
```env
BSC_RPC_URL=https://bnb-testnet.g.alchemy.com/v2/YOUR_KEY
RELAYER_PRIVATE_KEYS=key1,key2,key3
CONTRACT_USDT=0x0cF1130E64744860cbA5f992008527485C88F3C8
CONTRACT_PAYMASTER=0xA61D461AF55029B58d4846C9EA818De9cBC711D3
CONTRACT_ORACLE=0x18CC7E9CF8f40dd32Aa0fafD5FfE938B88E455a4
PORT=8080
```

### 14.2 前端配置

**.env.local 示例:**
```env
NEXT_PUBLIC_BACKEND_URL=http://localhost:8080
NEXT_PUBLIC_USDT_ADDRESS=0x0cF1130E64744860cbA5f992008527485C88F3C8
NEXT_PUBLIC_PAYMASTER_ADDRESS=0xA61D461AF55029B58d4846C9EA818De9cBC711D3
```

---

## 15. 测试计划

### 15.1 测试层次

```mermaid
flowchart TB
    subgraph Tests["测试层次"]
        ContractTest["合约测试<br/>forge test"]
        APITest["API测试<br/>curl/Postman"]
        IntegrationTest["集成测试<br/>完整流程"]
        E2ETest["端到端测试<br/>前端+后端+链"]
    end

    ContractTest --> APITest --> IntegrationTest --> E2ETest

    style Tests fill:#fff8e1
```

### 15.2 测试场景

```mermaid
flowchart LR
    subgraph Scenarios["测试场景"]
        S1["新用户领取USDT"]
        S2["用户7702授权"]
        S3["Gasless转账"]
        S4["清除7702"]
        S5["Relayer切换"]
        S6["手续费设置"]
    end

    style Scenarios fill:#e8f5e9
```

---

## 附录：关键参数汇总

| 参数 | 值 | 描述 |
|------|------|------|
| MAX_BATCH_SIZE | 5 | 单次batch最多操作数 |
| FAUCET_AMOUNT | 100 USDT | 水龙头每次领取金额 |
| feeRate默认值 | 0 | 手续费率 (10000=100%) |
| decimals | 18 | USDT精度 |
| Chain ID | 97 | BSC测试网 |
| Gas补偿精度 | 1e18 | BNB 18位精度转换 |

---

**文档结束**