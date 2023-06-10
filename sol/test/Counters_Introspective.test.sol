// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../Counters_Introspective.sol";
import "forge-std/Test.sol";

contract Counters_IntrospectiveTest is Test {

    ICounters public counters;

    function setUp() public {
        counters = new Counters_Introspective();
    }

    function test_Counters() public {
        assertEq(counters.get(0), 0);
        assertEq(counters.increment(0), 1);
        assertEq(counters.get(0), 1);
    }
}

