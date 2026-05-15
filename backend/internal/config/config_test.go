package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("BSC_RPC_URL", "https://test.com")
	os.Setenv("RELAYER_PRIVATE_KEYS", "key1,key2")
	os.Setenv("CONTRACT_USDT", "0x123")

	cfg := Load()

	if cfg.BSCRpcURL != "https://test.com" {
		t.Errorf("expected BSCRpcURL https://test.com, got %s", cfg.BSCRpcURL)
	}
	if len(cfg.RelayerPrivateKeys) != 2 {
		t.Errorf("expected 2 relayer keys, got %d", len(cfg.RelayerPrivateKeys))
	}
	if cfg.ContractUSDT != "0x123" {
		t.Errorf("expected ContractUSDT 0x123, got %s", cfg.ContractUSDT)
	}
}

func TestValidate(t *testing.T) {
	cfg := &Config{
		BSCRpcURL:          "https://test.com",
		RelayerPrivateKeys: []string{"key1"},
		ContractUSDT:       "0x123",
		ContractPaymaster:  "0x456",
	}

	if err := cfg.Validate(); err != nil {
		t.Errorf("expected valid config, got error %v", err)
	}

	invalidCfg := &Config{BSCRpcURL: ""}
	if err := invalidCfg.Validate(); err == nil {
		t.Error("expected validation error for missing BSC_RPC_URL")
	}
}
