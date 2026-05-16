# AA Wallet 技术设计文档

> 基于 EIP-7702 的 AA 钱包，实现无感授权和 USDT 支付 gas

---

## 1. 系统架构图

```mermaid
flowchart TB
    subgraph Frontend["前端 (Next.js)"]
        Faucet["水龙头页面"]
        Auth["7702授权页面"]
        Transfer["USDT转账页面"]
        Clear["清除7702页面"]
        Status["状态页面"]
        Admin["管理页面"]
    end

    subgraph Backend["后端 (Go)"]
        Pool["Relayer Pool<br/>私钥管理"]
        Handler["API Handlers<br/>交易构建"]
        Contract["合约交互层"]
    end

    subgraph Contracts["BSC测试网合约"]
        USDT["MockUSDT<br/>ERC20 + faucet"]
        Oracle["PriceOracle<br/>BNB/USDT汇率"]
        Paymaster["USDTPaymaster<br/>UUPS Proxy"]
        Account["Simple7702Account<br/>UUPS Proxy"]
    end

    subgraph Relayers["Relayer池"]
        R1["Relayer 1"]
        R2["Relayer 2"]
        R3["Relayer 3"]
        RN["Relayer N"]
    end

    Frontend -->|"HTTP API"| Backend
    Backend -->|"RPC调用"| Contracts
    Backend -->|"选择空闲"| Relayers
    Relayers -->|"代付gas"| Contracts

    USDT -->|"价格查询"| Oracle
    Account -->|"batch执行"| Paymaster
    Paymaster -->|"USDT补偿"| Relayers
    Paymaster -->|"汇率"| Oracle

    style Frontend fill:#e1f5fe
    style Backend fill:#fff3e0
    style Contracts fill:#e8f5e9
    style Relayers fill:#fce4ec
```

---

## 2. 合约部署依赖图

```mermaid
flowchart LR
    subgraph Deploy["部署顺序"]
        D1["1. MockUSDT"]
        D2["2. PriceOracle"]
        D3["3. USDTPaymaster"]
        D4["4. Simple7702Account"]
    end

    D1 -->|"USDT地址"| D2
    D1 -->|"USDT地址"| D3
    D2 -->|"Oracle地址"| D3
    D3 -->|"Paymaster地址"| D4

    subgraph Config["初始化配置"]
        Router["PancakeRouter"]
        WBNB["WBNB地址"]
    end

    Router --> D2
    WBNB --> D2
```

---

## 3. EIP-7702 授权流程

```mermaid
sequenceDiagram
    participant U as 用户EOA
    participant F as 前端
    participant B as 后端
    participant R as Relayer
    participant C as 合约(链上)

    Note over U,F: 用户首次使用流程

    U->>F: 输入私钥(本地)
    F->>F: 构造authorization<br/>绑定Simple7702Account
    F->>F: 私钥签名authorization
    F->>B: POST /api/authorize-7702<br/>{signedAuth}
    B->>B: 选择空闲Relayer
    B->>R: 构建setCode交易
    R->>C: 提交7702授权交易
    C->>C: EOA.code = Simple7702Account
    C-->>R: 交易确认
    R-->>B: txHash
    B-->>F: 返回交易哈希
    F-->>U: 显示授权成功

    Note over U,C: 用户EOA现在拥有合约代码
```

---

## 4. USDT转账流程 (无感授权 + Gas补偿)

```mermaid
sequenceDiagram
    participant U as 用户EOA(7702)
    participant F as 前端
    participant B as 后端
    participant R as Relayer
    participant P as Paymaster
    participant O as Oracle
    participant T as MockUSDT

    Note over U,F: 用户想转账USDT给他人

    U->>F: 输入目标地址+金额<br/>+私钥签名
    F->>F: 构造batch操作:<br/>[approve, transfer]
    F->>F: 用户私钥签名batch
    F->>B: POST /api/transfer-usdt<br/>{calls, signature}
    B->>B: 选择空闲Relayer
    B->>R: 构建executeBatch交易
    R->>P: executeBatch(userOp, sig)

    rect rgb(240, 248, 255)
        Note over P,T: 合约执行流程
        P->>P: 1.验证Relayer白名单
        P->>P: 2.验证用户签名
        P->>T: 3.执行approve(Paymaster)
        P->>T: 4.执行transfer(target)
        P->>P: 5.计算gasUsed
        P->>O: 6.获取BNB/USDT价格
        O-->>P: 返回汇率
        P->>P: 7.计算USDT补偿金额
        P->>T: 8.transferFrom(user→Relayer)
        P->>T: 9.transferFrom(user→feeRecipient)
    end

    P-->>R: 交易完成
    R-->>B: txHash + gasUsed
    B-->>F: 返回结果+补偿金额
    F-->>U: 显示转账成功

    Note over U,T: 用户无需持有BNB支付gas
```

---

## 5. Gas补偿计算流程

```mermaid
flowchart TB
    subgraph Input["输入"]
        GasUsed["gasUsed<br/>(交易消耗)"]
        GasPrice["gasPrice<br/>(当前gas价格)"]
        BNBPrice["BNB/USDT汇率<br/>(Oracle查询)"]
        FeeRate["feeRate<br/>(手续费率)"]
    end

    subgraph Calc["计算"]
        BNBCost["BNB成本<br/>= gasUsed × gasPrice"]
        USDTComp["USDT补偿<br/>= BNB成本 × BNB价格 ÷ 1e18"]
        Fee["手续费<br/>= USDT补偿 × feeRate ÷ 10000"]
        Total["总补偿<br/>= USDT补偿 + 手续费"]
    end

    subgraph Output["转账"]
        ToRelayer["→ Relayer<br/>(USDT补偿)"]
        ToFeeRecipient["→ feeRecipient<br/>(手续费)"]
    end

    GasUsed --> BNBCost
    GasPrice --> BNBCost
    BNBCost --> USDTComp
    BNBPrice --> USDTComp
    USDTComp --> Fee
    FeeRate --> Fee
    USDTComp --> ToRelayer
    Fee --> ToFeeRecipient

    style Input fill:#e3f2fd
    style Calc fill:#fff8e1
    style Output fill:#e8f5e9
```

---

## 6. Relayer选择流程

```mermaid
flowchart TB
    Request["交易请求"]

    Check["查询所有Relayer<br/>pending状态"]

    subgraph Pool["Relayer池"]
        R1["Relayer1<br/>pending: 0"]
        R2["Relayer2<br/>pending: 2"]
        R3["Relayer3<br/>pending: 1"]
    end

    Select["选择pending最少<br/>的Relayer"]

    Mark["标记为pending<br/>pending += 1"]

    Submit["提交交易"]

    Complete["交易完成<br/>pending -= 1"]

    Request --> Check
    Check --> Pool
    Pool -->|"pending=0"| Select
    Select --> Mark --> Submit --> Complete

    style Pool fill:#fce4ec
    style Select fill:#c8e6c9
```

---

## 7. 数据流图

```mermaid
flowchart LR
    subgraph UserInput["用户输入"]
        PrivKey["私钥(本地)"]
        Target["目标地址"]
        Amount["转账金额"]
    end

    subgraph FrontendProc["前端处理"]
        Sign1["签名7702 authorization"]
        Sign2["签名batch内容"]
        Build["构建calls数组"]
    end

    subgraph BackendProc["后端处理"]
        Select["选择Relayer"]
        BuildTx["构建交易"]
        Monitor["监控状态"]
    end

    subgraph OnChain["链上数据"]
        Auth7702["7702授权状态"]
        USDTBal["USDT余额"]
        Pending["待处理交易"]
    end

    subgraph Output["输出"]
        TxHash["交易哈希"]
        GasUsed["gas消耗"]
        Comp["USDT补偿"]
    end

    PrivKey --> Sign1 --> BackendProc
    PrivKey --> Sign2 --> BackendProc
    Target --> Build --> BackendProc
    Amount --> Build --> BackendProc

    BackendProc --> OnChain
    OnChain --> Output

    FrontendProc -->|"HTTP"| BackendProc
    BackendProc -->|"RPC"| OnChain

    style UserInput fill:#e1f5fe
    style FrontendProc fill:#fff3e0
    style BackendProc fill:#f3e5f5
    style OnChain fill:#e8f5e9
    style Output fill:#c8e6c9
```

---

## 8. 合约调用关系图

```mermaid
flowchart TB
    subgraph Paymaster["USDTPaymaster"]
        ExecuteBatch["executeBatch()"]
        AddRelayer["addRelayer()"]
        SetFee["setFeeRate()"]
        SetOracle["setOracle()"]
        IsRelayer["isRelayer()"]
    end

    subgraph Account["Simple7702Account"]
        BatchExec["executeBatch()"]
        ValidateSig["isValidSignature()"]
        SetPaymaster["setPaymaster()"]
    end

    subgraph Oracle["PriceOracle"]
        GetPrice["getBNBPriceInUSDT()"]
        SetRouter["setRouter()"]
    end

    subgraph USDT["MockUSDT"]
        Faucet["faucet()"]
        Mint["mint()"]
        Transfer["transfer()"]
        Approve["approve()"]
        TransferFrom["transferFrom()"]
    end

    ExecuteBatch -->|"验证签名"| ValidateSig
    ExecuteBatch -->|"执行操作"| USDT
    ExecuteBatch -->|"获取汇率"| GetPrice
    ExecuteBatch -->|"补偿"| TransferFrom

    Account -->|"onlyPaymaster"| Paymaster

    style Paymaster fill:#e8f5e9
    style Account fill:#fff3e0
    style Oracle fill:#e3f2fd
    style USDT fill:#fce4ec
```

---

## 9. API接口映射图

```mermaid
flowchart LR
    subgraph UserAPI["用户API"]
        U1["POST /api/authorize-7702"]
        U2["POST /api/clear-7702"]
        U3["POST /api/transfer-usdt"]
        U4["GET /api/user-status/:addr"]
        U5["GET /api/faucet-info"]
        U6["POST /api/faucet/:addr"]
    end

    subgraph AdminAPI["管理API"]
        A1["GET /api/admin/relayers"]
        A2["POST /api/admin/add-relayer"]
        A3["POST /api/admin/remove-relayer"]
        A4["POST /api/admin/set-fee-rate"]
        A5["POST /api/admin/set-oracle"]
    end

    subgraph ContractOp["合约操作"]
        C1["setCode(7702)"]
        C2["setCode(清空)"]
        C3["executeBatch()"]
        C4["查询余额/状态"]
        C5["mint()"]
        C6["addRelayer()"]
        C7["removeRelayer()"]
        C8["setFeeRate()"]
        C9["setOracle()"]
    end

    U1 --> C1
    U2 --> C2
    U3 --> C3
    U4 --> C4
    U5 -->|"返回地址"| U6
    U6 --> C5

    A1 -->|"查询状态"| A1
    A2 --> C6
    A3 --> C7
    A4 --> C8
    A5 --> C9

    style UserAPI fill:#e1f5fe
    style AdminAPI fill:#fff3e0
    style ContractOp fill:#e8f5e9
```

---

## 10. 安全验证流程

```mermaid
flowchart TB
    subgraph Input["executeBatch输入"]
        UserOp["UserOperation<br/>{user, calls}"]
        Sig["用户签名"]
        Caller["调用者(Relayer)"]
    end

    subgraph Validate["验证层"]
        V1["检查Relayer白名单"]
        V2["恢复签名获取signer"]
        V3["验证signer == user"]
    end

    subgraph Execute["执行层"]
        E1["遍历执行calls"]
        E2["记录gasUsed"]
    end

    subgraph Compensate["补偿层"]
        C1["查询Oracle汇率"]
        C2["计算USDT补偿"]
        C3["执行USDT转账"]
    end

    Input --> Validate
    V1 -->|"失败"| Reject["NotRelayer错误"]
    V2 --> V3
    V3 -->|"失败"| Reject2["InvalidSignature错误"]
    V3 -->|"成功"| Execute
    Execute --> Compensate
    E1 -->|"失败"| Reject3["CallFailed错误"]
    C3 -->|"失败"| Reject4["TransferFailed错误"]
    C3 -->|"成功"| Success["交易完成"]

    style Input fill:#e3f2fd
    style Validate fill:#fff8e1
    style Execute fill:#e8f5e9
    style Compensate fill:#fce4ec
    style Reject fill:#ffcdd2
    style Reject2 fill:#ffcdd2
    style Reject3 fill:#ffcdd2
    style Reject4 fill:#ffcdd2
    style Success fill:#c8e6c9
```

---

## 11. 前端页面流程

```mermaid
flowchart TB
    subgraph Pages["页面"]
        Home["首页<br/>用户状态"]
        Faucet["水龙头"]
        Auth["7702授权"]
        Clear["清除7702"]
        Transfer["USDT转账"]
        Admin["管理页面"]
    end

    Home -->|"查看状态"| Home
    Home -->|"领取USDT"| Faucet
    Home -->|"授权钱包"| Auth
    Home -->|"清除授权"| Clear
    Home -->|"转账USDT"| Transfer
    Home -->|"管理员"| Admin

    subgraph HomeState["首页状态"]
        Addr["地址输入"]
        Bound["7702绑定状态"]
        Bal["USDT余额"]
        Quick["快捷操作"]
    end

    subgraph TransferState["转账页面"]
        Priv["私钥输入"]
        Target["目标地址"]
        Amnt["金额"]
        EstGas["预估gas"]
        EstComp["预估补偿"]
    end

    style Pages fill:#e1f5fe
    style HomeState fill:#fff3e0
    style TransferState fill:#f3e5f5
```

---

## 12. 状态机：用户7702生命周期

```mermaid
stateDiagram-v2
    [*] --> EOA: 用户创建EOA

    EOA --> AuthPending: 提交7702授权
    AuthPending --> Auth7702: 交易确认
    AuthPending --> EOA: 交易失败

    Auth7702 --> TransferReady: 可执行batch
    TransferReady --> Transfering: 发起转账
    Transfering --> TransferReady: 转账完成
    Transfering --> TransferReady: 转账失败

    Auth7702 --> ClearPending: 提交清除
    ClearPending --> EOA: 清除成功
    ClearPending --> Auth7702: 清除失败

    state Auth7702 {
        [*] --> Bound
        Bound --> Executing: 执行batch
        Executing --> Bound: 完成
    }

    note right of EOA: 纯EOA状态<br/>需要BNB支付gas
    note right of Auth7702: 已绑定合约<br/>可用USDT支付gas
```

---

## 附录：合约地址配置

| 合约 | 描述 | 依赖 |
|------|------|------|
| MockUSDT | ERC20代币，带faucet | 无 |
| PriceOracle | BNB/USDT汇率 | PancakeRouter, WBNB, USDT |
| USDTPaymaster | UUPS代理，执行batch | USDT, Oracle |
| Simple7702Account | UUPS代理，用户合约 | Paymaster |

---

## 附录：关键参数

| 参数 | 值 | 说明 |
|------|------|------|
| MAX_BATCH_SIZE | 5 | 单次batch最多操作数 |
| FAUCET_AMOUNT | 100 USDT | 每次领取金额 |
| feeRate | 0 (默认) | 手续费率 (10000=100%) |
| 汇率精度 | 1e18 | BNB 18位精度 |

---

**文档版本**: v1.0
**生成日期**: 2026-05-15
**适用网络**: BSC测试网