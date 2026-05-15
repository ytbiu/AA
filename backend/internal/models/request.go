package models

type Authorize7702Request struct {
	UserAddress       string `json:"user_address" binding:"required"`
	AuthorizationData string `json:"authorization_data" binding:"required"`
	Signature         string `json:"signature" binding:"required"`
}

type Clear7702Request struct {
	UserAddress       string `json:"user_address" binding:"required"`
	AuthorizationData string `json:"authorization_data" binding:"required"`
	Signature         string `json:"signature" binding:"required"`
}

type TransferUSDTRequest struct {
	UserAddress   string `json:"user_address" binding:"required"`
	TargetAddress string `json:"target_address" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
	Signature     string `json:"signature" binding:"required"`
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
