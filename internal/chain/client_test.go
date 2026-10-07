package chain

import (
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Anvil account (9): nothing in the project uses it, so it keeps its initial 10000 ETH.
const (
	anvilRPC      = "http://127.0.0.1:8545"
	untouchedAddr = "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"
)

var wantBalance = new(big.Int).Mul(big.NewInt(10000), big.NewInt(1e18))

func testContext(t *testing.T) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	t.Cleanup(cancel)
	return ctx
}

func newTestClient(t *testing.T) *Client {
	t.Helper()
	client, err := NewClient(anvilRPC)
	if err != nil {
		t.Fatalf("creating the client: %v", err)
	}

	_, err = client.GetBalance(testContext(t), common.HexToAddress(untouchedAddr))
	if err != nil {
		t.Skip("no Anvil node available")
	}

	return client
}

func TestGetBalance(t *testing.T) {
	client := newTestClient(t)

	balance, err := client.GetBalance(testContext(t), common.HexToAddress(untouchedAddr))
	if err != nil {
		t.Fatalf("getting the balance: %v", err)
	}

	if balance.Cmp(wantBalance) != 0 {
		t.Errorf("balance of %s = %s, want %s", untouchedAddr, balance, wantBalance)
	}
}

func TestEnsureContract_NoContract(t *testing.T) {
	client := newTestClient(t)

	// A regular account has no bytecode, so there is no contract at this address.
	err := client.EnsureContract(testContext(t), common.HexToAddress(untouchedAddr))
	if err == nil {
		t.Fatal("expected an error for an address with no contract, got nil")
	}
	if !strings.Contains(err.Error(), "no contract found") {
		t.Errorf("error %q does not say that no contract was found", err)
	}
}
