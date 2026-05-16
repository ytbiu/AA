package api

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"

	"aa-wallet-backend/internal/models"
)

func (h *Handlers) AddRelayer(c *gin.Context) {
	var req models.AddRelayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondBadRequest(c, err)
		return
	}

	h.executeAdminTx(c, func(t *bind.TransactOpts) (*types.Transaction, error) {
		return h.paymaster.AddRelayerWithTransactor(t, parseAddress(req.RelayerAddress))
	})
}

func (h *Handlers) RemoveRelayer(c *gin.Context) {
	var req models.RemoveRelayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondBadRequest(c, err)
		return
	}

	h.executeAdminTx(c, func(t *bind.TransactOpts) (*types.Transaction, error) {
		return h.paymaster.RemoveRelayerWithTransactor(t, parseAddress(req.RelayerAddress))
	})
}

func (h *Handlers) SetFeeRate(c *gin.Context) {
	var req models.SetFeeRateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondBadRequest(c, err)
		return
	}

	h.executeAdminTx(c, func(t *bind.TransactOpts) (*types.Transaction, error) {
		return h.paymaster.SetFeeRateWithTransactor(t, big.NewInt(int64(req.FeeRate)))
	})
}

func (h *Handlers) SetOracle(c *gin.Context) {
	var req models.SetOracleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondBadRequest(c, err)
		return
	}

	h.executeAdminTx(c, func(t *bind.TransactOpts) (*types.Transaction, error) {
		return h.paymaster.SetOracleWithTransactor(t, parseAddress(req.OracleAddress))
	})
}
