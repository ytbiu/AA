package config

import (
	"os"
	"strings"
)

type Config struct {
	BSCRpcURL           string
	RelayerPrivateKeys  []string
	ContractUSDT        string
	ContractPaymaster   string
	ContractOracle      string
	Contract7702Account string
	Port                string
}

func Load() *Config {
	relayerKeys := strings.Split(os.Getenv("RELAYER_PRIVATE_KEYS"), ",")
	validKeys := make([]string, 0)
	for _, key := range relayerKeys {
		if strings.TrimSpace(key) != "" {
			validKeys = append(validKeys, strings.TrimSpace(key))
		}
	}

	return &Config{
		BSCRpcURL:           getEnvOrDefault("BSC_RPC_URL", "https://data-seed-prebsc-1-s1.binance.org:8545"),
		RelayerPrivateKeys:  validKeys,
		ContractUSDT:        os.Getenv("CONTRACT_USDT"),
		ContractPaymaster:   os.Getenv("CONTRACT_PAYMASTER"),
		ContractOracle:      os.Getenv("CONTRACT_ORACLE"),
		Contract7702Account: os.Getenv("CONTRACT_7702_ACCOUNT"),
		Port:                getEnvOrDefault("PORT", "8080"),
	}
}

func getEnvOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func (c *Config) Validate() error {
	if c.BSCRpcURL == "" {
		return &ValidationError{Field: "BSC_RPC_URL", Msg: "required"}
	}
	if len(c.RelayerPrivateKeys) == 0 {
		return &ValidationError{Field: "RELAYER_PRIVATE_KEYS", Msg: "at least one key required"}
	}
	if c.ContractUSDT == "" {
		return &ValidationError{Field: "CONTRACT_USDT", Msg: "required"}
	}
	if c.ContractPaymaster == "" {
		return &ValidationError{Field: "CONTRACT_PAYMASTER", Msg: "required"}
	}
	return nil
}

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Msg
}
