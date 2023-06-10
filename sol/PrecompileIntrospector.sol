// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./PrecompileRegistry.sol";
import "forge-std/Test.sol";

contract PrecompileIntrospector {
    string private precompileName;

    constructor(string memory name) {
        precompileName = name;
    }

    function _getPrecompileAddress() internal view returns (address) {
        bytes4 selector = bytes4(keccak256("getPrecompileByName(string)"));
        bytes memory data = abi.encodeWithSelector(selector, precompileName);
        (bool success, bytes memory result) = address(
            0xCc00000000000000000000000000000000000000
        ).staticcall(data);
        if (success && result.length == 32) {
            return abi.decode(result, (address));
        } else {
            return address(0);
        }
    }
}
