// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Counters_EVM.sol";
import "./PrecompileIntrospector.sol";

import "forge-std/Test.sol";

contract Counters_Introspective is Counters_EVM, PrecompileIntrospector {
    constructor() PrecompileIntrospector("SimpleCounters") {}

    function increment(uint256 id) public override returns (uint256) {
        address pc = _getPrecompileAddress();
        if (pc == address(0)) {
            return super.increment(id);
        } else {
            return ICounters(pc).increment(id);
        }
    }

    function get(uint256 id) public view override returns (uint256) {
        address pc = _getPrecompileAddress();
        if (pc == address(0)) {
            return super.get(id);
        } else {
            return ICounters(pc).get(id);
        }
    }
}
