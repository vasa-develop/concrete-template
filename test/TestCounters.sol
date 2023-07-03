// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../sol/Counters.sol";

contract TestCounters {
    Counters internal counters;

    function setUp() public {
        counters = Counters(COUNTERS_ADDRESS);
    }

    function testCounter() public {
        for (uint256 id = 0; id < 3; id++) {
            require(counters.get(id) == 0, "expected initial value to be 0");
            for (uint256 ii = 0; ii < 3; ii++) {
                require(
                    counters.increment(id) == ii + 1,
                    "unexpected increment return value"
                );
                require(
                    counters.get(id) == ii + 1,
                    "unexpected get return value"
                );
            }
        }
    }
}
