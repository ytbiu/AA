package api

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"

	"aa-wallet-backend/internal/models"
)

type adminTxResult struct {
	Status string `json:"status"`
	TxHash string `json:"tx_hash"`
}

func (h *Handlers) executeAdminTx(c *gin.Context, contractMethod func(*bind.TransactOpts) (*types.Transaction, error)) {
	relayer, err := h.relayerPool.SelectIdle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "no_relayer",
			Message: "no available relayer",
		})
		return
	}

	transactor, err := h.createTransactor(relayer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "transactor_failed",
			Message: err.Error(),
		})
		return
	}

	tx, err := contractMethod(transactor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "tx_failed",
			Message: err.Error(),
		})
		return
	}

	receipt, err := h.waitForTx(tx, 30*time.Second)
	if err != nil {
		c.JSON(http.StatusOK, adminTxResult{Status: "pending", TxHash: tx.Hash().Hex()})
		return
	}

	h.relayerPool.MarkComplete(relayer.Address)

	status := "success"
	if receipt.Status == 0 {
		status = "failed"
	}

	c.JSON(http.StatusOK, adminTxResult{Status: status, TxHash: tx.Hash().Hex()})
}

func respondBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, models.ErrorResponse{
		Error:   "invalid_request",
		Message: err.Error(),
	})
}

func respondInternalError(c *gin.Context, errorType, message string) {
	c.JSON(http.StatusInternalServerError, models.ErrorResponse{
		Error:   errorType,
		Message: message,
	})
}

func parseAddress(hex string) common.Address {
	return common.HexToAddress(hex)
}
