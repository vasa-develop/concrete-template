package main

import (
	"github.com/therealbytes/concrete-template/engine/pcs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete"
	"github.com/ethereum/go-ethereum/concrete/precompiles"
)

func setup(engine concrete.ConcreteApp) {
	address := common.HexToAddress("0x80")
	metadata := precompiles.PrecompileMetadata{
		Name:        "SimpleCounters",
		Version:     precompiles.NewVersion(0, 1, 0),
		Author:      "The concrete-geth authors",
		Description: "Increment and retrieve counters",
		Source:      "https://github.com/therealbytes/concrete-template/blob/main/engine/pcs/counters.go",
	}
	engine.AddPrecompile(address, pcs.CountersPrecompile, metadata)
}

func main() {
	engine := concrete.ConcreteGeth
	setup(engine)
	engine.Run()
}
