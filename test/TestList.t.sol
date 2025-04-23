// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "../src/lib/List.sol";
import {Test, console} from "forge-std/Test.sol";

contract TestList is Test {
    using List for List.U256;

    List.U256 private sell_list;

    function setUp() public {}

    function test_poc() external {
        insert(20, 1);
        insert(30, 1);
        insert(25, 1);
        uint256[] memory values = sell_list.values();
        assertEq(values.length, 3);
        assertEq(values[0], 20);
        assertEq(values[1], 25);
        assertEq(values[2], 30);
    }

    function insert(uint256 value, uint256 maxMatchCount) public {
        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prevPrice = sell_list.findASCPrev(value, adjacent, maxMatchCount);
        require(value >= prevPrice);
        sell_list.insert(value, prevPrice);
    }
}
