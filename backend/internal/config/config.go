package config

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	ListenAddr            string
	ChainID               *big.Int
	PaymasterAddress      common.Address
	DelegateLogicAddress  common.Address
	QuoteSignerPrivateKey string
	QuoteValidSeconds     int64
	TokenPerNativeByToken map[common.Address]*big.Int
	SettlementMarkupBps   uint64
	BundlerRPCURL         string
	RPCURL                string
	RelayerPrivateKey     string
}

func Load() (*Config, error) {
	cfg := &Config{
		ListenAddr:            envOr("LISTEN_ADDR", ":8080"),
		QuoteSignerPrivateKey: os.Getenv("QUOTE_SIGNER_PRIVATE_KEY"),
		BundlerRPCURL:         os.Getenv("BUNDLER_RPC_URL"),
		RPCURL:                envOr("RPC_URL", envOr("BSC_TESTNET_RPC_URL", "")),
		RelayerPrivateKey:     strings.TrimSpace(os.Getenv("RELAYER_PRIVATE_KEY")),
	}

	if cfg.QuoteSignerPrivateKey == "" {
		return nil, fmt.Errorf("QUOTE_SIGNER_PRIVATE_KEY is required")
	}

	chainID, ok := parseBigInt(envOr("CHAIN_ID", "97"))
	if !ok {
		return nil, fmt.Errorf("invalid CHAIN_ID")
	}
	cfg.ChainID = chainID

	paymasterAddrHex := os.Getenv("PAYMASTER_ADDRESS")
	if !common.IsHexAddress(paymasterAddrHex) {
		return nil, fmt.Errorf("invalid PAYMASTER_ADDRESS")
	}
	cfg.PaymasterAddress = common.HexToAddress(paymasterAddrHex)

	delegateLogicAddrHex := strings.TrimSpace(envOr("DELEGATE_LOGIC_ADDRESS", os.Getenv("NEXT_PUBLIC_DELEGATE_ACCOUNT_ADDRESS")))
	if !common.IsHexAddress(delegateLogicAddrHex) {
		return nil, fmt.Errorf("invalid DELEGATE_LOGIC_ADDRESS")
	}
	cfg.DelegateLogicAddress = common.HexToAddress(delegateLogicAddrHex)

	validSeconds, err := strconv.ParseInt(envOr("QUOTE_VALID_SECONDS", "300"), 10, 64)
	if err != nil || validSeconds <= 0 {
		return nil, fmt.Errorf("invalid QUOTE_VALID_SECONDS")
	}
	cfg.QuoteValidSeconds = validSeconds

	rates, err := parseTokenRates(envOr("TOKEN_RATES", ""))
	if err != nil {
		return nil, err
	}
	if len(rates) == 0 {
		return nil, fmt.Errorf("TOKEN_RATES is required, format: 0xTokenA:600000000,0xTokenB:600000000")
	}
	cfg.TokenPerNativeByToken = rates

	markup, err := strconv.ParseUint(envOr("SETTLEMENT_MARKUP_BPS", "500"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid SETTLEMENT_MARKUP_BPS")
	}
	cfg.SettlementMarkupBps = markup

	if cfg.BundlerRPCURL == "" {
		if cfg.RPCURL == "" {
			return nil, fmt.Errorf("RPC_URL is required when BUNDLER_RPC_URL is empty")
		}
		if cfg.RelayerPrivateKey == "" {
			return nil, fmt.Errorf("RELAYER_PRIVATE_KEY is required when BUNDLER_RPC_URL is empty")
		}
	}

	return cfg, nil
}

func envOr(key, fallback string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	return v
}

func parseBigInt(raw string) (*big.Int, bool) {
	raw = strings.TrimSpace(raw)
	base := 10
	if strings.HasPrefix(raw, "0x") || strings.HasPrefix(raw, "0X") {
		raw = raw[2:]
		base = 16
	}
	n := new(big.Int)
	n, ok := n.SetString(raw, base)
	if !ok {
		return nil, false
	}
	return n, true
}

func parseTokenRates(raw string) (map[common.Address]*big.Int, error) {
	out := make(map[common.Address]*big.Int)
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return out, nil
	}

	pairs := strings.Split(raw, ",")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid TOKEN_RATES pair: %s", pair)
		}
		tokenRaw := strings.TrimSpace(parts[0])
		if !common.IsHexAddress(tokenRaw) {
			return nil, fmt.Errorf("invalid token address in TOKEN_RATES: %s", tokenRaw)
		}
		rate, ok := parseBigInt(strings.TrimSpace(parts[1]))
		if !ok || rate.Sign() <= 0 {
			return nil, fmt.Errorf("invalid token rate in TOKEN_RATES: %s", pair)
		}
		out[common.HexToAddress(tokenRaw)] = rate
	}
	return out, nil
}
