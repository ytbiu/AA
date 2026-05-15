package eth

import (
	"testing"
)

func TestPrivateKeyToAddress(t *testing.T) {
	testKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	address, privKey, err := PrivateKeyToAddress(testKey)

	if err != nil {
		t.Errorf("failed to convert private key: %v", err)
	}

	if address.Hex() != "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" {
		t.Errorf("expected address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266, got %s", address.Hex())
	}

	if privKey == nil {
		t.Error("expected private key to be returned")
	}
}

func TestNewSimulatedClient(t *testing.T) {
	simClient, err := NewSimulatedClient()
	if err != nil {
		t.Errorf("failed to create simulated client: %v", err)
	}

	if simClient == nil {
		t.Error("expected simulated client")
	}

	if simClient.chainID.Int64() != 1337 {
		t.Errorf("expected chain ID 1337, got %d", simClient.chainID.Int64())
	}
}
