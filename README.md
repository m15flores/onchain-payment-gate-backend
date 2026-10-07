# onchain-payment-gate-backend

Go backend for `onchain-payment-gate`, a backend-controlled vault for on-chain payment settlement. Personal portfolio project targeting a backend + blockchain profile (on-chain payments / stablecoin infrastructure), built alongside the Blockchain Accelerator program.

> Smart contracts: [`onchain-payment-gate-contracts`](https://github.com/m15flores/onchain-payment-gate-contracts)

## What this is

The off-chain half of the system. This service holds the key that the `Controller` contract recognizes as the authorized backend, requests fund releases through it, and verifies on-chain that the Vault state matches what it expects. Custody and access control live in the contracts repo; this repo is the client that drives them.

Current stage: **Phase 0** of the project's phased roadmap. A minimal client that reads the Vault state and requests a release through the Controller over RPC. Persistence and event processing are planned for later phases.

## How it works

`cmd/gate` runs a single release end to end:

1. Loads the configuration from the environment and fails fast if a variable is missing or malformed.
2. Checks that both contracts have code deployed at their configured addresses.
3. Reads the Vault balance.
4. Calls `Controller.requestRelease`, signing with the backend key.
5. Waits for the transaction to be mined and checks that it succeeded.
6. Decodes the `ReleaseRequested` event from the receipt.
7. Reads the Vault balance again and checks that it dropped by exactly the released amount.

## Design notes

- Configuration is parsed and validated once, at startup. Errors name the offending variable and never include the private key.
- `chain.Client` wraps go-ethereum's client, so logging or retries can be added in a single place.
- The contract check uses `eth_getCode`. A node restarted without redeploying fails with a clear error instead of reporting an empty balance.
- ABIs are copied by hand from the contracts repo and the generated bindings are committed, so the backend builds without the contracts repo or `abigen`.

## Project structure

```
cmd/gate/        entry point
internal/
  chain/         wrapper around go-ethereum's ethclient
  config/        environment-based configuration
  bindings/      abigen bindings for Vault and Controller
abi/             contract ABIs copied from the contracts repo
```

## Getting started

### Prerequisites

Go 1.25 or later, and [Foundry](https://book.getfoundry.sh/) (`anvil`, `forge`, `cast`). The contracts repo cloned next to this one, with its submodules:

```bash
git clone --recurse-submodules <contracts-repo-url>
```

### 1. Start a local node

```bash
anvil
```

Leave it running in its own terminal.

### 2. Deploy the contracts

The Vault takes the Controller address in its constructor, and the Controller takes the Vault address in its own. To break the cycle, the Vault is deployed first and points to the address the Controller will have. Contract addresses are derived from the deployer and its nonce, so on a fresh Anvil node the Controller deployed second from account (0) always lands at the same address.

From the contracts repo:

```bash
RPC=http://127.0.0.1:8545
DEPLOYER_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80   # Anvil account (0)

# --constructor-args must be the last flag
forge create src/Vault.sol:Vault --rpc-url $RPC --private-key $DEPLOYER_KEY --broadcast --constructor-args 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512

forge create src/Controller.sol:Controller --rpc-url $RPC --private-key $DEPLOYER_KEY --broadcast --constructor-args 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 0x5FbDB2315678afecb367f032d93F642f64180aa3

cast send 0x5FbDB2315678afecb367f032d93F642f64180aa3 "deposit()" --value 10ether --rpc-url $RPC --private-key $DEPLOYER_KEY
```

The Controller is deployed with Anvil account (1) as its authorized backend, and the Vault ends up holding 10 ETH.

### 3. Configure

```bash
cp .env.example .env
```

For the deployment above, `.env` looks like this:

```dotenv
RPC_URL=http://127.0.0.1:8545
CHAIN_ID=31337
VAULT_ADDRESS=0x5FbDB2315678afecb367f032d93F642f64180aa3
CONTROLLER_ADDRESS=0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
BACKEND_PRIVATE_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
```

The private key belongs to Anvil account (1), the one the Controller recognizes as its backend. These are publicly known development keys: never use them outside a local node. `.env` is git-ignored, and `.env.example` documents every variable the service needs.

### 4. Run

```bash
go run ./cmd/gate
```

It prints the Vault balance, the transaction hash and the decoded `ReleaseRequested` event, and exits with an error if the Vault balance did not drop by exactly the released amount.

Each run releases 1 ETH from the Vault to Anvil account (2), so the 10 ETH deposited above cover ten runs. Anvil does not persist state: after a restart, deploy again. The startup check reports that no contract is found until you do.

## Regenerating the bindings

Only needed when the contracts change. Run these steps from the directory that contains both repos, with `jq` and `abigen` installed:

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

### 1. Extract the ABIs (contracts repo)

```bash
cd onchain-payment-gate-contracts
forge build
jq '.abi' out/Vault.sol/Vault.json > ../onchain-payment-gate-backend/abi/Vault.abi.json
jq '.abi' out/Controller.sol/Controller.json > ../onchain-payment-gate-backend/abi/Controller.abi.json
```

### 2. Generate the bindings (backend repo)

```bash
cd ../onchain-payment-gate-backend
abigen --abi abi/Vault.abi.json --pkg vault --type Vault --out internal/bindings/vault/vault.go
abigen --abi abi/Controller.abi.json --pkg controller --type Controller --out internal/bindings/controller/controller.go
```

## Testing

```bash
go test ./...
```

The `config` tests are plain unit tests: valid configuration, each required variable missing, each malformed value, and a check that an invalid private key is never echoed back in the error. The `chain` tests run against a local Anvil node and skip themselves when none is listening, so the suite also passes on a machine without Foundry.