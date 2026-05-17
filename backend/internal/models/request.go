package models

// EIP-7702 Authorization 结构
type EIP7702Authorization struct {
	ChainID uint64 `json:"chain_id"`
	Address string `json:"address"` // 要绑定的实现合约地址
	Nonce   uint64 `json:"nonce"`
	YParity uint8  `json:"y_parity"` // v 值 (0 或 1)
	R       string `json:"r"`        // 签名 r (hex)
	S       string `json:"s"`        // 签名 s (hex)
}

// Authorize7702Request - 7702 授权请求
// Relayer 会发送 setCode 交易，将用户 EOA 绑定到 Simple7702Account
type Authorize7702Request struct {
	UserAddress string `json:"user_address" binding:"required"`
	ChainID     uint64 `json:"chain_id" binding:"required"`
	Nonce       uint64 `json:"nonce"` // nonce=0 是有效值
	V           uint8  `json:"v" binding:"required"`
	R           string `json:"r" binding:"required"`
	S           string `json:"s" binding:"required"`
	Signature   string `json:"signature"` // 完整 65 字节签名 (可选，后端可从 v/r/s 恢复)
}

// Clear7702Request - 清除 7702 授权请求
// Relayer 发送 setCode 交易，implementation 设置为空地址
type Clear7702Request struct {
	UserAddress string `json:"user_address" binding:"required"`
	ChainID     uint64 `json:"chain_id" binding:"required"`
	Nonce       uint64 `json:"nonce"` // nonce=0 是有效值
	V           uint8  `json:"v" binding:"required"`
	R           string `json:"r" binding:"required"`
	S           string `json:"s" binding:"required"`
	Signature   string `json:"signature"`
}

type TransferUSDTRequest struct {
	UserAddress          string `json:"user_address" binding:"required"`
	TargetAddress        string `json:"target_address" binding:"required"`
	Amount               string `json:"amount" binding:"required"`
	Signature            string `json:"signature" binding:"required"`
	QuotedRelayerAddress string `json:"quoted_relayer_address"`
	// EIP-7702 authorization (可选，用户已授权则不需要)
	AuthChainID uint64 `json:"auth_chain_id"`
	AuthNonce   uint64 `json:"auth_nonce"`
	AuthV       uint8  `json:"auth_v"`
	AuthR       string `json:"auth_r"`
	AuthS       string `json:"auth_s"`
}

type AddRelayerRequest struct {
	RelayerAddress string `json:"relayer_address" binding:"required"`
}

type RemoveRelayerRequest struct {
	RelayerAddress string `json:"relayer_address" binding:"required"`
}

type SetFeeRateRequest struct {
	FeeRate uint64 `json:"fee_rate" binding:"required"`
}

type SetOracleRequest struct {
	OracleAddress string `json:"oracle_address" binding:"required"`
}
