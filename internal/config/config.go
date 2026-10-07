package config

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Config struct {
	RPCEndpoint       string
	ChainID           int64
	VaultAddress      common.Address
	ControllerAddress common.Address
	PrivateKey        *ecdsa.PrivateKey
}

func Load() (*Config, error) {
	chainIDStr := os.Getenv("CHAIN_ID")
	privateKeyStr := os.Getenv("BACKEND_PRIVATE_KEY")

	rpcURL, err := requireEnv("RPC_URL")
	if err != nil {
		return nil, err
	}

	chainID, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("CHAIN_ID must be an integer: %w", err)
	}

	vaultAddress, err := parseAddress("VAULT_ADDRESS")
	if err != nil {
		return nil, err
	}
	controllerAddress, err := parseAddress("CONTROLLER_ADDRESS")
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyStr, "0x"))
	if err != nil {
		return nil, fmt.Errorf("BACKEND_PRIVATE_KEY is not a valid private key: %w", err)
	}

	return &Config{
		RPCEndpoint:       rpcURL,
		ChainID:           chainID,
		VaultAddress:      vaultAddress,
		ControllerAddress: controllerAddress,
		PrivateKey:        privateKey,
	}, nil
}

func parseAddress(key string) (common.Address, error) {
	value, err := requireEnv(key)
	if err != nil {
		return common.Address{}, err
	}
	if !common.IsHexAddress(value) {
		return common.Address{}, fmt.Errorf(key, " is not a valid address: %q", key)
	}
	return common.HexToAddress(value), nil
}

func requireEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s is required but not set", key)
	}
	return value, nil
}
