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
