package relayer

import (
	"testing"
)

func TestNewPool(t *testing.T) {
	testKeys := []string{
		"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"59c6995e99842d4f3bfeecdd8c3e9a7b3d5df8a1e0f3f8c9b3e9a7b3d5df8a1e",
	}

	pool, err := NewPool(testKeys)
	if err != nil {
		t.Errorf("failed to create pool: %v", err)
	}

	if pool.GetCount() != 2 {
		t.Errorf("expected 2 relayers, got %d", pool.GetCount())
	}
}

func TestSelectIdle(t *testing.T) {
	testKeys := []string{"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"}
	pool, _ := NewPool(testKeys)

	relayer, err := pool.SelectIdle()
	if err != nil {
		t.Errorf("failed to select idle relayer: %v", err)
	}

	if relayer == nil {
		t.Error("expected relayer to be selected")
	}
}

func TestMarkPendingAndComplete(t *testing.T) {
	testKeys := []string{"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"}
	pool, _ := NewPool(testKeys)

	relayer, _ := pool.SelectIdle()
	pool.MarkPending(relayer.Address)

	infos := pool.GetAll()
	if infos[0].PendingTx != 1 {
		t.Errorf("expected pending tx 1, got %d", infos[0].PendingTx)
	}

	pool.MarkComplete(relayer.Address)
	infos = pool.GetAll()
	if infos[0].PendingTx != 0 {
		t.Errorf("expected pending tx 0, got %d", infos[0].PendingTx)
	}
}
