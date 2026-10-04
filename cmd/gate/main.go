package main

import (
	"context"
	"fmt"
	"log"
	"onchain-payment-gate-backend/internal/chain"
)

func main() {
	chainClient, err := chain.NewClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
		return
	}

	balance, err := chainClient.GetBalance(context.Background(), "0x5FbDB2315678afecb367f032d93F642f64180aa3")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Balance of Vault: ", balance)
}
