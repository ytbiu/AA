# AA Wallet 后端 API 文档

## 用户 API

### GET /api/user-status/:address

查询用户状态

**响应:**
```json
{
  "address": "0x...",
  "is_7702_bound": true,
  "bound_contract": "0x...",
  "usdt_balance": "100000000000000000000"
}
```

### GET /api/faucet-info

获取水龙头信息

**响应:**
```json
{
  "faucet_amount": "100",
  "usdt_address": "0x..."
}
```

### POST /api/authorize-7702

提交 7702 授权

**请求:**
```json
{
  "user_address": "0x...",
  "authorization_data": "...",
  "signature": "0x..."
}
```

**响应:**
```json
{
  "tx_hash": "0x...",
  "status": "pending",
  "bound_contract": "0x..."
}
```

### POST /api/clear-7702

清除 7702 绑定

**请求:**
```json
{
  "user_address": "0x...",
  "authorization_data": "...",
  "signature": "0x..."
}
```

**响应:**
```json
{
  "tx_hash": "0x...",
  "status": "pending"
}
```

### POST /api/transfer-usdt

USDT 转账

**请求:**
```json
{
  "user_address": "0x...",
  "target_address": "0x...",
  "amount": "1000000000000000000",
  "signature": "0x..."
}
```

**响应:**
```json
{
  "tx_hash": "0x...",
  "status": "pending",
  "compensation": "0",
  "gas_used": 0
}
```

## 管理 API

### GET /api/admin/relayers

列出所有 Relayer 状态

**响应:**
```json
{
  "relayers": [
    {
      "address": "0x...",
      "pending_tx": 0
    }
  ]
}
```

### POST /api/admin/add-relayer

添加 Relayer 到白名单

**请求:**
```json
{
  "relayer_address": "0x..."
}
```

### POST /api/admin/remove-relayer

移除 Relayer

**请求:**
```json
{
  "relayer_address": "0x..."
}
```

### POST /api/admin/set-fee-rate

设置手续费率

**请求:**
```json
{
  "fee_rate": 100
}
```

### POST /api/admin/set-oracle

更新 Oracle 地址

**请求:**
```json
{
  "oracle_address": "0x..."
}
```

## 健康检查

### GET /health

**响应:**
```json
{
  "status": "ok"
}
```