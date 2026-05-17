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
	CalldataHash string `json:"calldata_hash"`
}

type TransferUSDTQuoteResponse struct {
	RelayerAddress        string `json:"relayer_address"`
	GasEstimate           uint64 `json:"gas_estimate"`
	GasPrice              string `json:"gas_price"`
	BNBPriceInUSDT        string `json:"bnb_price_in_usdt"`
	FeeRate               uint64 `json:"fee_rate"`
	EstimatedGasCost      string `json:"estimated_gas_cost"`
	EstimatedPaymasterFee string `json:"estimated_paymaster_fee"`
	EstimatedTotalGasCost string `json:"estimated_total_gas_cost"`
	CalldataHash          string `json:"calldata_hash"`
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

type FaucetClaimResponse struct {
	TxHash string `json:"tx_hash"`
	Amount string `json:"amount"`
	Status string `json:"status"`
}

type RelayerStatusResponse struct {
	Relayers []RelayerInfo `json:"relayers"`
}

type RelayerInfo struct {
	Address   string `json:"address"`
	PendingTx int    `json:"pending_tx"`
	IsActive  bool   `json:"is_active"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
