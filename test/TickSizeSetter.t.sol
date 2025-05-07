// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {TickSizeSetter} from "../src/TickSizeSetter.sol";
import "./DEXBase.t.sol";

contract TickSizeSetterTest is DEXBaseTest {
    TickSizeSetter private TICK_SIZE_SETTER;
    uint8 quoteDecimals;
    uint8 baseDecimals;

    function setUp() external {
        _deploy(18, 6, 1e2, 1e4);
        vm.startPrank(OWNER);
        TICK_SIZE_SETTER = new TickSizeSetter(OWNER, address(CROSS_DEX));
        CROSS_DEX.setTickSizeSetter(address(TICK_SIZE_SETTER));
        QUOTE.approve(address(ROUTER), type(uint256).max);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.stopPrank();
        quoteDecimals = IERC20Metadata(address(QUOTE)).decimals();
        baseDecimals = IERC20Metadata(address(BASE)).decimals();
        vm.warp(86400);
        TICK_SIZE_SETTER.allUpdates();
    }

    function test_ticksize_findPriceIndex() external view {
        uint256 index;
        uint256 price;

        // gte
        price = 10 ** (quoteDecimals - 1);
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 1);

        price = 10 ** (quoteDecimals);
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 2);

        price = 10 ** (quoteDecimals + 1);
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 3);

        price = 10 ** (quoteDecimals + 1);
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 5 * price);
        assertEq(index, 4);

        price = 10 ** (quoteDecimals + 2);
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 5);

        price = 10 ** (quoteDecimals + 2);
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 5 * price);
        assertEq(index, 6);

        price = 10 ** (quoteDecimals + 3);
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 7);

        // lt
        price = 10 ** (quoteDecimals - 1) - 1;
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 0);

        price = 10 ** (quoteDecimals) - 1;
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 1);

        price = 10 ** (quoteDecimals + 1) - 1;
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 2);

        price = 10 ** (quoteDecimals + 1) - 1;
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 5 * price);
        assertEq(index, 3);

        price = 10 ** (quoteDecimals + 2) - 1;
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 4);

        price = 10 ** (quoteDecimals + 2) - 1;
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 5 * price);
        assertEq(index, 5);

        price = 10 ** (quoteDecimals + 3) - 1;
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 6);

        // check 0 and max to ensure no overflow
        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 0);
        assertEq(index, 0);

        index = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, type(uint256).max);
        assertEq(index, 7);
    }

    function test_ticksize_all_updates() external {
        uint256 index;
        uint256 price;

        // gte
        price = 10 ** (quoteDecimals - 1);
        index = 1;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = 10 ** (quoteDecimals);
        index = 2;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = 10 ** (quoteDecimals + 1);
        index = 3;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = 5 * (10 ** (quoteDecimals + 1));
        index = 4;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = 10 ** (quoteDecimals + 2);
        index = 5;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = 5 * (10 ** (quoteDecimals + 2));
        index = 6;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = 10 ** (quoteDecimals + 3);
        index = 7;
        _check_ticksize_sinario(address(PAIR), price, index);

        // lt
        price = (10 ** (quoteDecimals + 3)) - PAIR.tickSize();
        index = 6;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = (5 * (10 ** (quoteDecimals + 2))) - PAIR.tickSize();
        index = 5;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = (10 ** (quoteDecimals + 2)) - PAIR.tickSize();
        index = 4;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = (5 * (10 ** (quoteDecimals + 1))) - PAIR.tickSize();
        index = 3;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = (10 ** (quoteDecimals + 1)) - PAIR.tickSize();
        index = 2;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = (10 ** (quoteDecimals)) - PAIR.tickSize();
        index = 1;
        _check_ticksize_sinario(address(PAIR), price, index);

        price = (10 ** (quoteDecimals - 1)) - PAIR.tickSize();
        index = 0;
        _check_ticksize_sinario(address(PAIR), price, index);
    }

    function _check_ticksize_sinario(address pair, uint256 price, uint256 expectIndex) private {
        vm.startPrank(OWNER);

        (, uint256 tickSize,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, uint256 lotSize) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        uint256 updateInterval = TICK_SIZE_SETTER.updateInterval();
        _ensurePrice(pair, price);
        vm.warp(block.timestamp + updateInterval);

        TICK_SIZE_SETTER.allUpdates();

        (uint256 tick, uint256 lot) = PAIR.tickSizes();
        assertEq(tick, tickSize, "tick size");
        assertEq(lot, lotSize, "lot size");
        vm.stopPrank();
    }

    function _ensurePrice(address pair, uint256 price) private {
        vm.startPrank(OWNER);
        ROUTER.submitSellLimit(pair, price, 10 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        ROUTER.submitBuyLimit(pair, price, 10 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        vm.stopPrank();
    }
}
