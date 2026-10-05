package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"onchain-payment-gate-backend/internal/bindings/controller"
	"onchain-payment-gate-backend/internal/chain"
	"onchain-payment-gate-backend/internal/config"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joho/godotenv"
)

const (
	demoRecipient = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC" // Anvil account (2)
	demoAmountWei = 1e18                                         // 1 ETH
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file loaded: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	client, err := chain.NewClient(cfg.RPCEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.EnsureContract(ctx, cfg.VaultAddress); err != nil {
		log.Fatal(err)
	}

	if err := client.EnsureContract(ctx, cfg.ControllerAddress); err != nil {
		log.Fatal(err)
	}

	balanceBefore, err := client.GetBalance(context.Background(), cfg.VaultAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance of Vault: ", balanceBefore)

	auth, err := bind.NewKeyedTransactorWithChainID(cfg.PrivateKey, big.NewInt(cfg.ChainID))
	if err != nil {
		log.Fatal(err)
	}

	controllerContract, err := controller.NewController(cfg.ControllerAddress, client.Backend())
	if err != nil {
		log.Fatal(err)
	}

	amount := big.NewInt(demoAmountWei)
	recipient := common.HexToAddress(demoRecipient)

	tx, err := controllerContract.RequestRelease(auth, recipient, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Printing tx.Hash().Hex(): ", tx.Hash().Hex())

	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatalf("transaction %s reverted", tx.Hash().Hex())
	}

	for _, l := range receipt.Logs {
		if l.Address != cfg.ControllerAddress {
			continue
		}
		controllerReleaseRequested, err := controllerContract.ParseReleaseRequested(*l)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Recipient: %s\n", controllerReleaseRequested.To)
		fmt.Printf("Amount: %d\n", controllerReleaseRequested.Amount)
	}

	balanceAfter, err := client.GetBalance(context.Background(), cfg.VaultAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	released := new(big.Int).Sub(balanceBefore, balanceAfter)
	if released.Cmp(amount) != 0 {
		log.Fatalf("vault balance dropped by %s, expected %s", released, amount)
	}
}
