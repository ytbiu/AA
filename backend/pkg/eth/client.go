package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client  *ethclient.Client
	chainID *big.Int
}

func NewClient(rpcURL string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	return &Client{client: client, chainID: chainID}, nil
}

func (c *Client) GetEthClient() *ethclient.Client {
	return c.client
}

func (c *Client) GetChainID() *big.Int {
	return c.chainID
}

func (c *Client) GetBalance(address common.Address) (*big.Int, error) {
	return c.client.BalanceAt(context.Background(), address, nil)
}

func (c *Client) GetPendingNonce(address common.Address) (uint64, error) {
	return c.client.PendingNonceAt(context.Background(), address)
}

func (c *Client) GetNonceAt(address common.Address) (uint64, error) {
	return c.client.NonceAt(context.Background(), address, nil)
}

func (c *Client) SuggestGasPrice() (*big.Int, error) {
	return c.client.SuggestGasPrice(context.Background())
}

func PrivateKeyToAddress(privateKey string) (common.Address, *ecdsa.PrivateKey, error) {
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("invalid private key: %v", err)
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, nil, fmt.Errorf("cannot get public key")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, privKey, nil
}

func CreateTransactor(privKey *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(privKey, chainID)
}

type SimulatedClient struct {
	*backends.SimulatedBackend
	chainID *big.Int
}

func NewSimulatedClient() (*SimulatedClient, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	genesisAlloc := core.GenesisAlloc{
		address: {Balance: big.NewInt(1000000000000000000)},
	}

	backend := backends.NewSimulatedBackend(genesisAlloc, 10000000)
	chainID := big.NewInt(1337)

	return &SimulatedClient{SimulatedBackend: backend, chainID: chainID}, nil
}
