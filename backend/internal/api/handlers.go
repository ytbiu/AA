package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"aa-wallet-backend/internal/contract"
	"aa-wallet-backend/internal/models"
	"aa-wallet-backend/internal/relayer"
)

type Handlers struct {
	relayerPool *relayer.Pool
	paymaster   *contract.PaymasterContract
	usdt        *contract.USDTContract
}

func NewHandlers(pool *relayer.Pool, paymaster *contract.PaymasterContract, usdt *contract.USDTContract) *Handlers {
	return &Handlers{
		relayerPool: pool,
		paymaster:   paymaster,
		usdt:        usdt,
	}
}

func (h *Handlers) GetUserStatus(c *gin.Context) {
	address := c.Param("address")

	c.JSON(http.StatusOK, models.UserStatusResponse{
		Address:       address,
		Is7702Bound:   false,
		BoundContract: "",
		USDTBalance:   "0",
	})
}

func (h *Handlers) GetFaucetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, models.FaucetInfoResponse{
		FaucetAmount: "100",
		UsdtAddress:  h.usdt.GetAddress().Hex(),
	})
}

func (h *Handlers) GetRelayers(c *gin.Context) {
	infos := h.relayerPool.GetAll()
	relayerInfos := make([]models.RelayerInfo, len(infos))
	for i, info := range infos {
		relayerInfos[i] = models.RelayerInfo{
			Address:   info.Address,
			PendingTx: info.PendingTx,
		}
	}
	c.JSON(http.StatusOK, models.RelayerStatusResponse{
		Relayers: relayerInfos,
	})
}

func (h *Handlers) Authorize7702(c *gin.Context) {
	var req models.Authorize7702Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	relayer, err := h.relayerPool.SelectIdle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "no_relayer",
			Message: "no available relayer",
		})
		return
	}

	h.relayerPool.MarkPending(relayer.Address)

	c.JSON(http.StatusOK, models.Authorize7702Response{
		TxHash:        "0x...",
		Status:        "pending",
		BoundContract: req.AuthorizationData,
	})
}

func (h *Handlers) Clear7702(c *gin.Context) {
	var req models.Clear7702Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Clear7702Response{
		TxHash: "0x...",
		Status: "pending",
	})
}

func (h *Handlers) TransferUSDT(c *gin.Context) {
	var req models.TransferUSDTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	relayer, err := h.relayerPool.SelectIdle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "no_relayer",
			Message: "no available relayer",
		})
		return
	}

	h.relayerPool.MarkPending(relayer.Address)

	c.JSON(http.StatusOK, models.TransferUSDTResponse{
		TxHash:       "0x...",
		Status:       "pending",
		Compensation: "0",
		GasUsed:      0,
	})
}

func (h *Handlers) AddRelayer(c *gin.Context) {
	var req models.AddRelayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handlers) RemoveRelayer(c *gin.Context) {
	var req models.RemoveRelayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handlers) SetFeeRate(c *gin.Context) {
	var req models.SetFeeRateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handlers) SetOracle(c *gin.Context) {
	var req models.SetOracleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
