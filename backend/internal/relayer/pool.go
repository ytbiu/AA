package relayer

import (
	"crypto/ecdsa"
	"sync"
	"time"

	"aa-wallet-backend/pkg/eth"

	"github.com/ethereum/go-ethereum/common"
)

type Relayer struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
	PendingTx  int
	LastUsed   int64
}

type Pool struct {
	relayers []*Relayer
	mu       sync.RWMutex
}

func NewPool(privateKeys []string) (*Pool, error) {
	relayers := make([]*Relayer, 0)
	for _, key := range privateKeys {
		address, privKey, err := eth.PrivateKeyToAddress(key)
		if err != nil {
			return nil, err
		}
		relayers = append(relayers, &Relayer{
			Address:    address,
			PrivateKey: privKey,
			PendingTx:  0,
		})
	}
	return &Pool{relayers: relayers}, nil
}

type RelayerInfo struct {
	Address   string `json:"address"`
	PendingTx int    `json:"pending_tx"`
}

func (p *Pool) GetAll() []RelayerInfo {
	p.mu.RLock()
	defer p.mu.RUnlock()

	infos := make([]RelayerInfo, 0)
	for _, r := range p.relayers {
		infos = append(infos, RelayerInfo{
			Address:   r.Address.Hex(),
			PendingTx: r.PendingTx,
		})
	}
	return infos
}

func (p *Pool) GetCount() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.relayers)
}

func (p *Pool) MarkPending(address common.Address) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, r := range p.relayers {
		if r.Address == address {
			r.PendingTx++
			r.LastUsed = time.Now().Unix()
		}
	}
}

func (p *Pool) MarkComplete(address common.Address) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, r := range p.relayers {
		if r.Address == address {
			r.PendingTx--
			if r.PendingTx < 0 {
				r.PendingTx = 0
			}
		}
	}
}
