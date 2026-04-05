package crypto

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	gcrypto "github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
	privateKeyHex string
	privateKey    *ecdsa.PrivateKey
}

func NewSigner(privateKeyHex string) (*Signer, error) {
	normalized := strings.TrimPrefix(strings.TrimSpace(privateKeyHex), "0x")
	pk, err := gcrypto.HexToECDSA(normalized)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}
	return &Signer{privateKeyHex: normalized, privateKey: pk}, nil
}

func (s *Signer) Address() common.Address {
	return gcrypto.PubkeyToAddress(s.privateKey.PublicKey)
}

func (s *Signer) SignEthMessageHash(hash []byte) ([]byte, error) {
	signedHash := accounts.TextHash(hash)
	sig, err := gcrypto.Sign(signedHash, s.privateKey)
	if err != nil {
		return nil, err
	}
	// go-ethereum 返回的 v 是 0/1，这里转成 27/28 以兼容 solidity ECDSA.recover
	sig[64] += 27
	return sig, nil
}
