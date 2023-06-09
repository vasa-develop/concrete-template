// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// The address the precompile will be mapped to.
address constant COUNTERS_ADDRESS = address(0x80);

interface Counters {
    // Increment the counter for the given id, and return the new value.
    function increment(uint256 id) external returns (uint256);

    // Get the current value of the counter for the given id.
    function get(uint256 id) external view returns (uint256);
}
