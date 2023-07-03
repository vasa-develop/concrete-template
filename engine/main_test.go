package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/concrete"
	cc_testing "github.com/ethereum/go-ethereum/concrete/testing"
)

func TestE2E(t *testing.T) {
	engine := concrete.ConcreteGeth
	setup(engine)
	_, fails := cc_testing.Test(cc_testing.TestConfig{
		TestDir: "../test",
		OutDir:  "../out",
	})
	if fails > 0 {
		t.Errorf("TestE2E failed with %d failures", fails)
	}
}
