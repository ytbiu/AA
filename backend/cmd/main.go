package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"aa-wallet-backend/internal/api"
	"aa-wallet-backend/internal/config"
	"aa-wallet-backend/internal/contract"
	"aa-wallet-backend/internal/relayer"
	"aa-wallet-backend/pkg/eth"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	ethClient, err := eth.NewClient(cfg.BSCRpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum: %v", err)
	}

	relayerPool, err := relayer.NewPool(cfg.RelayerPrivateKeys)
	if err != nil {
		log.Fatalf("Failed to create relayer pool: %v", err)
	}

	paymaster := contract.NewPaymaster(cfg.ContractPaymaster, ethClient.GetEthClient())
	usdt := contract.NewUSDT(cfg.ContractUSDT, ethClient.GetEthClient())

	handlers := api.NewHandlers(relayerPool, paymaster, usdt)

	r := gin.Default()

	r.GET("/api/user-status/:address", handlers.GetUserStatus)
	r.GET("/api/faucet-info", handlers.GetFaucetInfo)
	r.POST("/api/authorize-7702", handlers.Authorize7702)
	r.POST("/api/clear-7702", handlers.Clear7702)
	r.POST("/api/transfer-usdt", handlers.TransferUSDT)

	r.GET("/api/admin/relayers", handlers.GetRelayers)
	r.POST("/api/admin/add-relayer", handlers.AddRelayer)
	r.POST("/api/admin/remove-relayer", handlers.RemoveRelayer)
	r.POST("/api/admin/set-fee-rate", handlers.SetFeeRate)
	r.POST("/api/admin/set-oracle", handlers.SetOracle)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("Server starting on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
