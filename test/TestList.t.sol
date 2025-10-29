// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.20;

import "../src/lib/List.sol";
import {Test, console} from "forge-std/Test.sol";

contract TestList is Test {
    using List for List.U256;

    List.U256 private sell_list;
    List.U256 private buy_list;

    function setUp() public {}

    function test_poc_asc() external {
        insert(20, 1);
        insert(30, 1);
        insert(25, 1);
        uint256[] memory values = sell_list.values();
        assertEq(values.length, 3);
        assertEq(values[0], 20);
        assertEq(values[1], 25);
        assertEq(values[2], 30);
    }

    function test_poc_desc() external {
        insertDESC(20, 1);
        insertDESC(30, 1);
        insertDESC(25, 1);
        uint256[] memory values = buy_list.values();
        assertEq(values.length, 3);
        assertEq(values[0], 30);
        assertEq(values[1], 25);
        assertEq(values[2], 20);
    }

    function testFuzz_poc_sell(uint256[] calldata values) external {
        uint256 length = values.length;
        for (uint256 i = 0; i < length; i++) {
            if (values[i] == 0) continue;
            insert(values[i], length);
        }
        uint256 len = List.length(sell_list);
        console.log("sell_list.length", len);

        uint256 before = 0;
        for (uint256 i = 0; i < len; i++) {
            uint256 value = sell_list.at(i);
            assertGt(value, before);
            before = value;
        }
    }

    function testFuzz_poc_buy(uint256[] memory values) external {
        uint256 length = values.length;
        for (uint256 i = 0; i < length; i++) {
            if (values[i] == 0) continue;
            if (values[i] == type(uint256).max) --values[i];
            insertDESC(values[i], length);
        }
        uint256 len = List.length(buy_list);
        console.log("buy_list.length", len);

        uint256 before = type(uint256).max;
        for (uint256 i = 0; i < len; i++) {
            uint256 value = buy_list.at(i);
            assertLt(value, before);
            before = value;
        }
    }

    function insert(uint256 value, uint256 maxMatchCount) public {
        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prevPrice = sell_list.findASCPrev(value, adjacent, maxMatchCount);
        require(value >= prevPrice, "value < prevPrice");
        sell_list.insert(value, prevPrice);
    }

    function insertDESC(uint256 value, uint256 maxMatchCount) public {
        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prevPrice = buy_list.findDESCPrev(value, adjacent, maxMatchCount);
        require(prevPrice == 0 || prevPrice >= value, "prevPrice < value");
        buy_list.insert(value, prevPrice);
    }
}
