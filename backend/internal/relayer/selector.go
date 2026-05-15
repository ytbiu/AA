package relayer

import (
	"sort"
)

type NoRelayerError struct{}

func (e *NoRelayerError) Error() string {
	return "no available relayer"
}

func (p *Pool) SelectIdle() (*Relayer, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if len(p.relayers) == 0 {
		return nil, &NoRelayerError{}
	}

	sorted := make([]*Relayer, len(p.relayers))
	copy(sorted, p.relayers)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].PendingTx < sorted[j].PendingTx
	})

	return sorted[0], nil
}
