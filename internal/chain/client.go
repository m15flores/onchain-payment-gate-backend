package chain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	eth *ethclient.Client
}

func NewClient(rpcURL string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &Client{eth: client}, nil
}

func (client *Client) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	return client.eth.BalanceAt(ctx, address, nil)
}

func (client *Client) Backend() bind.ContractBackend {
	return client.eth
}

func (client *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(ctx, client.eth, tx)
}

func (client *Client) EnsureContract(ctx context.Context, address common.Address) error {
	code, err := client.eth.CodeAt(ctx, address, nil)
	if err != nil {
		return fmt.Errorf("Error while calling RPC: %w", err)
	}

	if len(code) == 0 {
		return fmt.Errorf("checking code at %s: %w", address, err)
	}

	return nil
}
