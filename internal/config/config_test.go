package config

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// Public Anvil test values, safe to use in tests.
const (
	validRPC        = "http://127.0.0.1:8545"
	validChainID    = "31337"
	validVault      = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	validController = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
	validKey        = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

	wantChainID = 31337
)

const (
	envRPCURL     = "RPC_URL"
	envChainID    = "CHAIN_ID"
	envVault      = "VAULT_ADDRESS"
	envController = "CONTROLLER_ADDRESS"
	envPrivateKey = "BACKEND_PRIVATE_KEY"
)

var requiredVars = []string{envRPCURL, envChainID, envVault, envController, envPrivateKey}

func setValidEnv(t *testing.T) {
	t.Helper()
	t.Setenv(envRPCURL, validRPC)
	t.Setenv(envChainID, validChainID)
	t.Setenv(envVault, validVault)
	t.Setenv(envController, validController)
	t.Setenv(envPrivateKey, validKey)

}

func TestLoad_Success(t *testing.T) {
	setValidEnv(t)
	cfg, err := Load()
	if err != nil {
		t.Fatalf("error while loading config: %q", err)
	}

	if cfg.RPCEndpoint != validRPC {
		t.Errorf("RPC_URL = %s, want %s", cfg.RPCEndpoint, validRPC)
	}
	if cfg.ChainID != wantChainID {
		t.Errorf("CHAIN_ID = %d, want %d", cfg.ChainID, wantChainID)
	}
	if cfg.VaultAddress != common.HexToAddress(validVault) {
		t.Errorf("VAULT_ADDRESS = %s, want %s", cfg.VaultAddress, validVault)
	}
	if cfg.ControllerAddress != common.HexToAddress(validController) {
		t.Errorf("CONTROLLER_ADDRESS = %s, want %s", cfg.ControllerAddress, validController)
	}
	if cfg.PrivateKey == nil {
		t.Errorf("Error while parsing BACKEND_PRIVATE_KEY")
	}
}

func TestLoad_MissingVariable(t *testing.T) {
	for _, name := range requiredVars {
		t.Run(name, func(t *testing.T) {
			setValidEnv(t)
			t.Setenv(name, "")

			_, err := Load()
			if err == nil {
				t.Fatalf("expected an error when %s is missing, got nil", name)
			}
			if !strings.Contains(err.Error(), name) {
				t.Errorf("error %q does not mention %s", err, name)
			}
		})
	}
}

func TestLoad_InvalidValues(t *testing.T) {
	var invalidValues = []struct {
		name     string
		variable string
		value    string
	}{
		{"chain id is not a number", envChainID, "abc"},
		{"vault address bad length", envVault, "0x123"},
		{"controller address bad format", envController, "not-an-address"},
		{"private key bad format", envPrivateKey, "0xzz"},
	}

	for _, tc := range invalidValues {
		t.Run(tc.name, func(t *testing.T) {
			setValidEnv(t)
			t.Setenv(tc.variable, tc.value)

			_, err := Load()
			if err == nil {
				t.Fatalf("expected an error when %s has bad format, got nil", tc.variable)
			}
			if !strings.Contains(err.Error(), tc.variable) {
				t.Errorf("error %q does not mention %s", err, tc.variable)
			}
		})
	}
}

func TestLoad_InvalidPrivateKeyDoesNotLeakValue(t *testing.T) {
	setValidEnv(t)
	t.Setenv(envPrivateKey, "0xnot-a-key-SECRETMARKER")

	_, err := Load()
	if err == nil {
		t.Fatalf("expected an error when %s has bad format, got nil", envPrivateKey)
	}
	if strings.Contains(err.Error(), "SECRETMARKER") {
		t.Errorf("error %q leaks the key %s", err, "SECRETMARKER")
	}
}
