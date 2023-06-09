package pcs

import (
	_ "embed"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete/api"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

func TestCounters(t *testing.T) {
	var (
		addr       = common.HexToAddress("0x1234")
		db         = state.NewDatabase(rawdb.NewMemoryDatabase())
		statedb, _ = state.New(common.Hash{}, db, nil)
		evm        = vm.NewEVM(vm.BlockContext{}, vm.TxContext{}, statedb, params.TestChainConfig, vm.Config{})
		concrete   = api.New(evm.NewConcreteEVM(), addr)
	)
	increment := func(id int64) int64 {
		input, err := CountersABI.Pack("increment", big.NewInt(id))
		if err != nil {
			t.Fatal(err)
		}
		output, err := CountersPrecompile.Run(concrete, input)
		if err != nil {
			t.Fatal(err)
		}
		returns, err := CountersABI.Unpack("increment", output)
		if err != nil {
			t.Fatal(err)
		}
		return returns[0].(*big.Int).Int64()
	}
	get := func(id int64) int64 {
		input, err := CountersABI.Pack("get", big.NewInt(id))
		if err != nil {
			t.Fatal(err)
		}
		output, err := CountersPrecompile.Run(concrete, input)
		if err != nil {
			t.Fatal(err)
		}
		returns, err := CountersABI.Unpack("get", output)
		if err != nil {
			t.Fatal(err)
		}
		return returns[0].(*big.Int).Int64()
	}

	for id := int64(0); id < 3; id++ {
		if get(id) != 0 {
			t.Fatal("expected initial value to be 0")
		}
		for ii := 0; ii < 3; ii++ {
			if increment(id) != int64(ii+1) {
				t.Fatal("expected increment to return", ii+1)
			}
			if get(id) != int64(ii+1) {
				t.Fatal("expected get to return", ii+1)
			}
		}
	}
}
