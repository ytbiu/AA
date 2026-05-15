package models

type Authorize7702Response struct {
	TxHash        string `json:"tx_hash"`
	Status        string `json:"status"`
	BoundContract string `json:"bound_contract"`
}

type Clear7702Response struct {
	TxHash string `json:"tx_hash"`
	Status string `json:"status"`
}

type TransferUSDTResponse struct {
	TxHash       string `json:"tx_hash"`
	Status       string `json:"status"`
	Compensation string `json:"compensation"`
	GasUsed      uint64 `json:"gas_used"`
}

type UserStatusResponse struct {
	Address       string `json:"address"`
	Is7702Bound   bool   `json:"is_7702_bound"`
	BoundContract string `json:"bound_contract"`
	USDTBalance   string `json:"usdt_balance"`
}

type FaucetInfoResponse struct {
	FaucetAmount string `json:"faucet_amount"`
	UsdtAddress  string `json:"usdt_address"`
}

type RelayerStatusResponse struct {
	Relayers []RelayerInfo `json:"relayers"`
}

type RelayerInfo struct {
	Address   string `json:"address"`
	PendingTx int    `json:"pending_tx"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
