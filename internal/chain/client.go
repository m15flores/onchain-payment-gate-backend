package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

func (client *Client) GetBalance(ctx context.Context, addressStr string) (*big.Int, error) {
	address := common.HexToAddress(addressStr)
	return client.eth.BalanceAt(ctx, address, nil)
}
