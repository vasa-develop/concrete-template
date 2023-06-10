// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Counters.sol";

contract Counters_EVM is ICounters {
    mapping(uint256 => uint256) internal _counters;

    function increment(uint256 id) public virtual returns (uint256) {
        _counters[id]++;
        return _counters[id];
    }

    function get(uint256 id) public view virtual returns (uint256) {
        return _counters[id];
    }
}
