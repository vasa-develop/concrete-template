# Concrete-template

This is a simple template for projects using the Concrete framework.

## Directory structure

```bash
├── engine
│   ├── main.go                 # Chain setup
│   ├── main_test.go            # Minimal setup test
│   ├── pcs
│   │   ├── abi
│   │   │   └── Counters.json   # Precompile ABI created by `yarn build:abi`
│   │   ├── counters.go         # Example precompile
│   │   └── counters_test.go    # Precompile test
└── sol
    └── Counters.sol            # Precompile interface in Solidity
```

## Running a devnet

Run a normal Optimism Bedrock devnet replacing op-geth for `github.com/therealbytes/concrete-template-geth:latest` in `ops-bedrock/Dockerfile.l2`.

Alternatively, clone and run an already modified version of Bedrock:

```bash
# Clone repo
git clone -b concrete-template https://github.com/therealbytes/optimism.git
cd optimism
# Start devnet
make devnet-up
# Stop and clean devnet
make devnet-down && make devnet-clean
```
