package main

import (
	"log"
	"net/http"

	"aa-bsc-7702-demo/backend/internal/config"
	appcrypto "aa-bsc-7702-demo/backend/internal/crypto"
	"aa-bsc-7702-demo/backend/internal/httpapi"
	"aa-bsc-7702-demo/backend/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	signer, err := appcrypto.NewSigner(cfg.QuoteSignerPrivateKey)
	if err != nil {
		log.Fatalf("init signer failed: %v", err)
	}

	paymasterSvc, err := service.NewPaymasterService(cfg, signer)
	if err != nil {
		log.Fatalf("init paymaster service failed: %v", err)
	}

	server := httpapi.NewServer(cfg, paymasterSvc)
	log.Printf("paymaster signer: %s", signer.Address().Hex())
	log.Printf("listen on %s", cfg.ListenAddr)
	if err := http.ListenAndServe(cfg.ListenAddr, server.Routes()); err != nil {
		log.Fatalf("http server stopped: %v", err)
	}
}
