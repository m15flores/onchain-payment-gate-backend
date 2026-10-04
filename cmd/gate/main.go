package main

import (
	"context"
	"fmt"
	"log"
	"onchain-payment-gate-backend/internal/chain"
	"onchain-payment-gate-backend/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file loaded: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := chain.NewClient(cfg.RPCEndpoint)
	if err != nil {
		log.Fatal(err)
		return
	}

	balance, err := client.GetBalance(context.Background(), cfg.VaultAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Balance of Vault: ", balance)
}
