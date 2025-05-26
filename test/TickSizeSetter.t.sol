// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {TickSizeSetter} from "../src/TickSizeSetter.sol";
import "./DEXBase.t.sol";

contract TickSizeSetterTest is DEXBaseTest {
    TickSizeSetter private TICK_SIZE_SETTER;
    uint8 quoteDecimals;
    uint8 baseDecimals;

    function setUp() external {
        _deploy(18, 6, 1e4, 1e4);
        vm.startPrank(OWNER);
        TICK_SIZE_SETTER = new TickSizeSetter(OWNER, address(CROSS_DEX));
        CROSS_DEX.setTickSizeSetter(address(TICK_SIZE_SETTER));
        QUOTE.approve(address(ROUTER), type(uint256).max);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.stopPrank();
        quoteDecimals = IERC20Metadata(address(QUOTE)).decimals();
        baseDecimals = IERC20Metadata(address(BASE)).decimals();
        vm.warp(86400 * 2);
        TICK_SIZE_SETTER.allUpdates();
    }

    function test_ticksize_findPriceIndex() external view {
        uint256 tickSize;
        uint256 lotSize;
        uint256 callTick;
        uint256 callLot;
        uint256 price;
        uint256 expectIndex;

        // gte
        price = 10 ** (quoteDecimals - 1);
        expectIndex = 1;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 10 ** (quoteDecimals);
        expectIndex = 2;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 10 ** (quoteDecimals + 1);
        expectIndex = 3;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 5 * (10 ** (quoteDecimals + 1));
        expectIndex = 4;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 10 ** (quoteDecimals + 2);
        expectIndex = 5;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 5 * (10 ** (quoteDecimals + 2));
        expectIndex = 6;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 10 ** (quoteDecimals + 3);
        expectIndex = 7;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        // lt
        price = 10 ** (quoteDecimals - 1) - 1;
        expectIndex = 0;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 10 ** (quoteDecimals) - 1;
        expectIndex = 1;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 10 ** (quoteDecimals + 1) - 1;
        expectIndex = 2;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 5 * (10 ** (quoteDecimals + 1)) - 1;
        expectIndex = 3;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 10 ** (quoteDecimals + 2) - 1;
        expectIndex = 4;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 5 * (10 ** (quoteDecimals + 2)) - 1;
        expectIndex = 5;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = 10 ** (quoteDecimals + 3) - 1;
        expectIndex = 6;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        // check 0 and max to ensure no overflow
        price = 0;
        expectIndex = 0;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);

        price = type(uint256).max;
        expectIndex = 7;
        (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        assertEq(callTick, tickSize);
        assertEq(callLot, lotSize);
    }

    function test_ticksize_all_updates() external {
        uint256 index;
        uint256 price;

        // gte
        price = 10 ** (quoteDecimals - 1);
        index = 1;
        _all_updates_execute(address(PAIR), price, index);

        price = 10 ** (quoteDecimals);
        index = 2;
        _all_updates_execute(address(PAIR), price, index);

        price = 10 ** (quoteDecimals + 1);
        index = 3;
        _all_updates_execute(address(PAIR), price, index);

        price = 5 * (10 ** (quoteDecimals + 1));
        index = 4;
        _all_updates_execute(address(PAIR), price, index);

        price = 10 ** (quoteDecimals + 2);
        index = 5;
        _all_updates_execute(address(PAIR), price, index);

        price = 5 * (10 ** (quoteDecimals + 2));
        index = 6;
        _all_updates_execute(address(PAIR), price, index);

        price = 10 ** (quoteDecimals + 3);
        index = 7;
        _all_updates_execute(address(PAIR), price, index);

        // lt
        price = (10 ** (quoteDecimals + 3)) - PAIR.tickSize();
        index = 6;
        _all_updates_execute(address(PAIR), price, index);

        price = (5 * (10 ** (quoteDecimals + 2))) - PAIR.tickSize();
        index = 5;
        _all_updates_execute(address(PAIR), price, index);

        price = (10 ** (quoteDecimals + 2)) - PAIR.tickSize();
        index = 4;
        _all_updates_execute(address(PAIR), price, index);

        price = (5 * (10 ** (quoteDecimals + 1))) - PAIR.tickSize();
        index = 3;
        _all_updates_execute(address(PAIR), price, index);

        price = (10 ** (quoteDecimals + 1)) - PAIR.tickSize();
        index = 2;
        _all_updates_execute(address(PAIR), price, index);

        price = (10 ** (quoteDecimals)) - PAIR.tickSize();
        index = 1;
        _all_updates_execute(address(PAIR), price, index);

        price = (10 ** (quoteDecimals - 1)) - PAIR.tickSize();
        index = 0;
        _all_updates_execute(address(PAIR), price, index);
    }

    function test_ticksize_matched_price_at(uint32 _warp) external {
        vm.warp(block.timestamp + _warp);

        vm.startPrank(OWNER);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        BASE.approve(address(ROUTER), type(uint256).max);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(1);
        uint256 priceStep = price / 10;

        // sell orders
        for (uint256 i = 0; i < 10; i++) {
            ROUTER.submitSellLimit(
                address(PAIR),
                price + (priceStep * i),
                amount,
                IPair.LimitConstraints.GOOD_TILL_CANCEL,
                _searchPrices,
                0
            );
            // The latest price is not updated because the orders are not matched.
            assertEq(0, PAIR.matchedPrice());
        }

        // buy orders

        uint256 priceTime = TICK_SIZE_SETTER.calcTimestamp(block.timestamp);
        for (uint256 i = 0; i < 10; i++) {
            uint256 _price = price + (priceStep * i);
            ROUTER.submitBuyLimit(
                address(PAIR), _price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            // The latest price is updated because the orders are matched.
            assertEq(_price, PAIR.matchedPrice());
            assertEq(_price, ROUTER.matchedPriceAt(address(PAIR), priceTime));
        }
    }

    address[] private _allMarkets;
    mapping(address market => address[]) private _allPairs;

    function test_force_update() external {
        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        address[] storage pairs = _allPairs[address(MARKET)];
        address targetPair = address(PAIR);
        _make_market_pair(1, 2, 10 ** (quoteDecimals - 1));

        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates();

        uint256 expectUpdateTime = block.timestamp - (block.timestamp % TICK_SIZE_SETTER.updateInterval());
        for (uint256 i = 0; i < pairs.length; i++) {
            address pair = pairs[i];
            assertEq(expectUpdateTime, TICK_SIZE_SETTER.lastUpdateTimestamp(pair));
        }

        vm.prank(OWNER);
        TICK_SIZE_SETTER.forceUpdate(targetPair, 100e18, 200e18);
        // check set tick size
        (uint256 tickSize, uint256 lotSize) = PAIR.tickSizes();
        assertEq(tickSize, 200e18);
        assertEq(lotSize, 100e18);
        // check added skipPairs
        address[] memory skipPairs = TICK_SIZE_SETTER.allSkipPairs();
        assertEq(skipPairs.length, 1);
        assertEq(skipPairs[0], targetPair);
        // all pairs update
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates();

        // check skiped target pairs
        (tickSize, lotSize) = PAIR.tickSizes();
        assertEq(tickSize, 200e18);
        assertEq(lotSize, 100e18);

        uint256 expectUpdateTime2 = block.timestamp - (block.timestamp % TICK_SIZE_SETTER.updateInterval());
        assertNotEq(expectUpdateTime, expectUpdateTime2);
        for (uint256 i = 0; i < pairs.length; i++) {
            address pair = pairs[i];
            if (pair == targetPair) {
                // target pair not updated
                assertEq(expectUpdateTime, TICK_SIZE_SETTER.lastUpdateTimestamp(pair));
            } else {
                // other pairs updated
                assertEq(expectUpdateTime2, TICK_SIZE_SETTER.lastUpdateTimestamp(pair));
            }
        }

        // remove skip pair
        vm.prank(OWNER);
        TICK_SIZE_SETTER.removeSkipPair(targetPair);
        // update all pairs
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates();
        // check updated target pair
        (tickSize, lotSize) = PAIR.tickSizes();
        assertNotEq(tickSize, 200e18);
        assertNotEq(lotSize, 100e18);
        // check all pairs updated
        uint256 expectUpdateTime3 = block.timestamp - (block.timestamp % TICK_SIZE_SETTER.updateInterval());
        assertNotEq(expectUpdateTime2, expectUpdateTime3);
        for (uint256 i = 0; i < pairs.length; i++) {
            address pair = pairs[i];
            assertEq(expectUpdateTime3, TICK_SIZE_SETTER.lastUpdateTimestamp(pair));
        }
    }

    function test_ticksize_all_updates_max_count_case1() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 36 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(1, 36, 10 ** (quoteDecimals - 1));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_max_count_case2() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 34 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(2, 17, 10 ** (quoteDecimals));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_max_count_case3() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 34 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(17, 2, (10 ** (quoteDecimals + 2)));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_max_count_case4() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 33 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(33, 1, 5 * (10 ** (quoteDecimals + 2)));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_update() external {
        uint256 GAS_LIMIT = 4e6;
        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(2, 36, 10 ** (quoteDecimals - 1));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertFalse(market_ == address(0));

        TICK_SIZE_SETTER.update(market_, startIndex, 0);
        (market_,) = TICK_SIZE_SETTER.updatable();
        assertTrue(market_ == address(0));
    }

    function test_ticksize_all_updates_update_repeat() external {
        uint256 GAS_LIMIT = 4e6;
        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(3, 36, 10 ** (quoteDecimals - 1));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertFalse(market_ == address(0));

        TICK_SIZE_SETTER.update(market_, startIndex, 0);
        (market_, startIndex) = TICK_SIZE_SETTER.updatable();
        assertFalse(market_ == address(0));

        TICK_SIZE_SETTER.update(market_, startIndex, 0);
        (market_,) = TICK_SIZE_SETTER.updatable();
        assertTrue(market_ == address(0));
    }

    function test_ticksize_update_count() external {
        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(1, 60, 10 ** (quoteDecimals - 1));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());

        uint256 updateTime = block.timestamp - (block.timestamp % TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.update(address(MARKET), 0, 30);
        address pair = _allPairs[address(MARKET)][29];
        assertEq(updateTime, TICK_SIZE_SETTER.lastUpdateTimestamp(pair));

        pair = _allPairs[address(MARKET)][30];
        assertNotEq(updateTime, TICK_SIZE_SETTER.lastUpdateTimestamp(pair));

        TICK_SIZE_SETTER.update(address(MARKET), 30, 30);
        (address market,) = TICK_SIZE_SETTER.updatable();
        assertEq(market, address(0));
    }

    function test_ticksize_set_size_formats() external {
        _all_updates_execute(address(PAIR), 10 ** (quoteDecimals - 1), 1);
        // gte ~ lt price | tickSize | lotSize
        // -----------------------------------
        // [0]      ~ 0.1     | 0.0001  | 100
        // [1] 0.1  ~ 1       | 0.001   | 10
        // [2] 1    ~ 10      | 0.01    | 1
        // [3] 10   ~ 100     | 0.1     | 0.1
        // [4] 100  ~ 1000    | 1       | 0.01
        // [5] 1000 ~         | 10      | 0.001
        {
            // update size formats
            TickSizeSetter.SizeFormat[] memory formats = new TickSizeSetter.SizeFormat[](6);
            formats[0] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: -1,
                tickSizeUnit: 1,
                tickSizeScale: -4,
                lotSizeUnit: 1,
                lotSizeScale: 2
            });
            formats[1] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 0,
                tickSizeUnit: 1,
                tickSizeScale: -3,
                lotSizeUnit: 1,
                lotSizeScale: 1
            });
            formats[2] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 1,
                tickSizeUnit: 1,
                tickSizeScale: -2,
                lotSizeUnit: 1,
                lotSizeScale: 0
            });
            formats[3] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 2,
                tickSizeUnit: 1,
                tickSizeScale: -1,
                lotSizeUnit: 1,
                lotSizeScale: -1
            });
            formats[4] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 3,
                tickSizeUnit: 1,
                tickSizeScale: 0,
                lotSizeUnit: 1,
                lotSizeScale: -2
            });
            formats[5] = TickSizeSetter.SizeFormat({
                minPriceUnit: 0,
                minPriceScale: 0,
                tickSizeUnit: 1,
                tickSizeScale: 1,
                lotSizeUnit: 1,
                lotSizeScale: -3
            });
            vm.startPrank(OWNER);
            TICK_SIZE_SETTER.setSizeFormats(formats);
            vm.stopPrank();
        }
        {
            uint256 tickSize;
            uint256 lotSize;
            uint256 callTick;
            uint256 callLot;
            uint256 price;
            uint256 expectIndex;
            // call findPriceIndex to check if the size formats are correct
            // gte
            price = 10 ** (quoteDecimals - 1);
            expectIndex = 1;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            price = 10 ** (quoteDecimals);
            expectIndex = 2;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            price = 10 ** (quoteDecimals + 1);
            expectIndex = 3;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            price = 10 ** (quoteDecimals + 2);
            expectIndex = 4;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            price = 10 ** (quoteDecimals + 3);
            expectIndex = 5;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            // lt
            price = 10 ** (quoteDecimals - 1) - 1;
            expectIndex = 0;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            price = 10 ** (quoteDecimals) - 1;
            expectIndex = 1;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            price = 10 ** (quoteDecimals + 1) - 1;
            expectIndex = 2;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            price = 10 ** (quoteDecimals + 2) - 1;
            expectIndex = 3;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);

            price = 10 ** (quoteDecimals + 3) - 1;
            expectIndex = 4;
            (tickSize, lotSize) = TICK_SIZE_SETTER.tickSizeByPrice(address(PAIR), price);
            (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
            (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
            assertEq(callTick, tickSize);
            assertEq(callLot, lotSize);
        }
        {
            uint256 price;
            uint256 index;

            // gte
            price = 10 ** (quoteDecimals - 1);
            index = 1;
            _all_updates_execute(address(PAIR), price, index);

            price = 10 ** (quoteDecimals);
            index = 2;
            _all_updates_execute(address(PAIR), price, index);

            price = 10 ** (quoteDecimals + 1);
            index = 3;
            _all_updates_execute(address(PAIR), price, index);

            price = 10 ** (quoteDecimals + 2);
            index = 4;
            _all_updates_execute(address(PAIR), price, index);

            price = 10 ** (quoteDecimals + 3);
            index = 5;
            _all_updates_execute(address(PAIR), price, index);

            price = 10 ** (quoteDecimals + 4);
            index = 5;
            _all_updates_execute(address(PAIR), price, index);

            // lt
            price = (10 ** (quoteDecimals + 3)) - PAIR.tickSize();
            index = 4;
            _all_updates_execute(address(PAIR), price, index);

            price = (10 ** (quoteDecimals + 2)) - PAIR.tickSize();
            index = 3;
            _all_updates_execute(address(PAIR), price, index);

            price = (10 ** (quoteDecimals + 1)) - PAIR.tickSize();
            index = 2;
            _all_updates_execute(address(PAIR), price, index);

            price = (10 ** (quoteDecimals)) - PAIR.tickSize();
            index = 1;
            _all_updates_execute(address(PAIR), price, index);

            price = (10 ** (quoteDecimals - 1)) - PAIR.tickSize();
            index = 0;
            _all_updates_execute(address(PAIR), price, index);
        }
    }

    function test_ticksize_set_size_formats_invalid() external {
        // gte ~ lt price | tickSize | lotSize
        // -----------------------------------
        // [0]      ~ 0.1     | 0.0001  | 100
        // [1] 0.1  ~ 1       | 0.001   | 10
        // [2] 1    ~ 10      | 0.01    | 1
        // [3] 10   ~ 100     | 0.1     | 0.1
        // [4] 100  ~ 1000    | 1       | 0.01
        // [5] 1000 ~         | 10      | 0.001
        {
            // update size formats
            TickSizeSetter.SizeFormat[] memory formats = new TickSizeSetter.SizeFormat[](6);
            formats[0] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: -1,
                tickSizeUnit: 1,
                tickSizeScale: -4,
                lotSizeUnit: 1,
                lotSizeScale: 2
            });
            formats[1] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 0,
                tickSizeUnit: 1,
                tickSizeScale: -3,
                lotSizeUnit: 1,
                lotSizeScale: 1
            });
            formats[3] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 1,
                tickSizeUnit: 1,
                tickSizeScale: -2,
                lotSizeUnit: 1,
                lotSizeScale: 0
            });
            formats[2] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 2,
                tickSizeUnit: 1,
                tickSizeScale: -1,
                lotSizeUnit: 1,
                lotSizeScale: -1
            });
            formats[4] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 3,
                tickSizeUnit: 1,
                tickSizeScale: 0,
                lotSizeUnit: 1,
                lotSizeScale: -2
            });
            formats[5] = TickSizeSetter.SizeFormat({
                minPriceUnit: 0,
                minPriceScale: 0,
                tickSizeUnit: 1,
                tickSizeScale: 1,
                lotSizeUnit: 1,
                lotSizeScale: -3
            });
            vm.startPrank(OWNER);
            vm.expectRevert(abi.encodeWithSignature("TickSizeSetterInvalidPrice(uint256)", 3));
            TICK_SIZE_SETTER.setSizeFormats(formats);

            TickSizeSetter.SizeFormat[] memory emptyFormats;
            bytes32 field = bytes32("formats");
            vm.expectRevert(abi.encodeWithSignature("TickSizeSetterZeroInput(bytes32)", field));
            TICK_SIZE_SETTER.setSizeFormats(emptyFormats);
            vm.stopPrank();
        }
    }

    function test_ticksize_set_size_formats_limit() external {
        address market = address(MARKET);
        address[] storage pairs = _allPairs[market];
        // When there are 6 decimal places in total,
        // with a gas limit of 4e6, formatting succeeds for up to 11 entries.

        vm.startPrank(OWNER);
        T20 erc20 = new T20("BASE", "BASE", 18);
        erc20.approve(address(ROUTER), type(uint256).max);
        address pair = MarketImpl(market).createPair(address(erc20), QUOTE_DECIMALS / 1e4, 10 ** 18 / 1e4);
        pairs.push(pair);
        erc20 = new T20("BASE", "BASE", 5);
        erc20.approve(address(ROUTER), type(uint256).max);
        pair = MarketImpl(market).createPair(address(erc20), QUOTE_DECIMALS / 1e4, 10 ** 5 / 1e4);
        pairs.push(pair);
        erc20 = new T20("BASE", "BASE", 6);
        erc20.approve(address(ROUTER), type(uint256).max);
        pair = MarketImpl(market).createPair(address(erc20), QUOTE_DECIMALS / 1e4, 10 ** 6 / 1e4);
        pairs.push(pair);
        erc20 = new T20("BASE", "BASE", 8);
        erc20.approve(address(ROUTER), type(uint256).max);
        pair = MarketImpl(market).createPair(address(erc20), QUOTE_DECIMALS / 1e4, 10 ** 8 / 1e4);
        pairs.push(pair);
        erc20 = new T20("BASE", "BASE", 12);
        erc20.approve(address(ROUTER), type(uint256).max);
        pair = MarketImpl(market).createPair(address(erc20), QUOTE_DECIMALS / 1e4, 10 ** 12 / 1e4);
        pairs.push(pair);
        erc20 = new T20("BASE", "BASE", 20);
        erc20.approve(address(ROUTER), type(uint256).max);
        pair = MarketImpl(market).createPair(address(erc20), QUOTE_DECIMALS / 1e4, 10 ** 20 / 1e4);
        pairs.push(pair);

        TICK_SIZE_SETTER.allUpdates();
        uint256[] memory allDecimals = TICK_SIZE_SETTER.allDecimals();
        assertEq(allDecimals.length, 6);

        TickSizeSetter.SizeFormat[] memory formats = new TickSizeSetter.SizeFormat[](11);
        for (uint256 i = 0; i < formats.length; i++) {
            formats[i] = TickSizeSetter.SizeFormat({
                minPriceUnit: 1 + uint8(i),
                minPriceScale: -1,
                tickSizeUnit: 1,
                tickSizeScale: -4,
                lotSizeUnit: 1,
                lotSizeScale: 2
            });
        }
        TICK_SIZE_SETTER.setSizeFormats{gas: 4e6}(formats);

        vm.stopPrank();
    }

    function test_ticksize_set_resolved_sizes_by_pair() external {
        _make_market_pair(1, 2, 10 ** (quoteDecimals - 1));
        address targetPair = _allPairs[_allMarkets[0]][1];
        assertNotEq(address(PAIR), targetPair);
        vm.label(targetPair, "targetPair");

        vm.startPrank(OWNER);
        TickSizeSetter.ResolvedSize[] memory sizes;
        {
            // check revert
            bytes32 field = bytes32("sizes");
            vm.expectRevert(abi.encodeWithSignature("TickSizeSetterZeroInput(bytes32)", field));
            TICK_SIZE_SETTER.setResolvedSizesByPair(targetPair, sizes);
        }
        {
            // check revert
            sizes = new TickSizeSetter.ResolvedSize[](4);
            sizes[1] = TickSizeSetter.ResolvedSize({
                minPrice: 10 ** quoteDecimals,
                tickSize: 7 * (10 ** quoteDecimals),
                lotSize: 7 * (10 ** baseDecimals)
            });
            sizes[0] = TickSizeSetter.ResolvedSize({
                minPrice: 10 * (10 ** (quoteDecimals)),
                tickSize: 70 * (10 ** quoteDecimals),
                lotSize: 70 * (10 ** baseDecimals)
            });
            sizes[2] = TickSizeSetter.ResolvedSize({
                minPrice: 100 * (10 ** (quoteDecimals)),
                tickSize: 700 * (10 ** quoteDecimals),
                lotSize: 700 * (10 ** baseDecimals)
            });
            vm.expectRevert(abi.encodeWithSignature("TickSizeSetterInvalidPrice(uint256)", 1));
            TICK_SIZE_SETTER.setResolvedSizesByPair(targetPair, sizes);
        }

        // gte ~ lt price | tickSize | lotSize
        // -----------------------------------
        // [0]       ~ 100   | 4   | 4
        // [1] 100   ~ 1000  | 40  | 40
        // [2] 1000  ~ 10000 | 400 | 400
        sizes = new TickSizeSetter.ResolvedSize[](4);
        sizes[0] = TickSizeSetter.ResolvedSize({
            minPrice: 100 * (10 ** quoteDecimals),
            tickSize: 4 * (10 ** quoteDecimals),
            lotSize: 4 * (10 ** baseDecimals)
        });
        sizes[1] = TickSizeSetter.ResolvedSize({
            minPrice: 1000 * (10 ** (quoteDecimals)),
            tickSize: 40 * (10 ** quoteDecimals),
            lotSize: 40 * (10 ** baseDecimals)
        });
        sizes[2] = TickSizeSetter.ResolvedSize({
            minPrice: 10000 * (10 ** (quoteDecimals)),
            tickSize: 400 * (10 ** quoteDecimals),
            lotSize: 400 * (10 ** baseDecimals)
        });
        TICK_SIZE_SETTER.setResolvedSizesByPair(targetPair, sizes);
        vm.stopPrank();

        uint256 tickSize;
        uint256 lotSize;
        uint256 minPrice;
        uint256 callTick;
        uint256 callLot;

        // case [0]
        (tickSize,) = PAIR.tickSizes();
        (minPrice,,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, 0);
        _ensurePrice(address(PAIR), minPrice - tickSize);

        (tickSize,) = PairImpl(targetPair).tickSizes();
        _ensurePrice(targetPair, sizes[0].minPrice - tickSize);

        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: 4e6}();
        assertEq(TICK_SIZE_SETTER.lastUpdateTimestamp(address(PAIR)), TICK_SIZE_SETTER.lastUpdateTimestamp(targetPair));

        (tickSize, lotSize) = PAIR.tickSizes();
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, 0);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, 0);
        assertEq(callTick, tickSize, "[0] address(PAIR) tick size");
        assertEq(callLot, lotSize, "[0] address(PAIR) lot size");

        (tickSize, lotSize) = PairImpl(targetPair).tickSizes();
        assertEq(tickSize, sizes[0].tickSize, "[0] targetPair tick size");
        assertEq(lotSize, sizes[0].lotSize, "[0] targetPair lot size");

        assertNotEq(callTick, tickSize);
        assertNotEq(callLot, lotSize);

        // case [1]
        _ensurePrice(address(PAIR), minPrice);
        _ensurePrice(targetPair, sizes[0].minPrice);

        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: 4e6}();
        assertEq(TICK_SIZE_SETTER.lastUpdateTimestamp(address(PAIR)), TICK_SIZE_SETTER.lastUpdateTimestamp(targetPair));

        (tickSize, lotSize) = PAIR.tickSizes();
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, 1);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, 1);
        assertEq(callTick, tickSize, "[1] address(PAIR) tick size");
        assertEq(callLot, lotSize, "[1] address(PAIR) lot size");

        (tickSize, lotSize) = PairImpl(targetPair).tickSizes();
        assertEq(tickSize, sizes[1].tickSize, "[1] targetPair tick size");
        assertEq(lotSize, sizes[1].lotSize, "[1] targetPair lot size");

        assertNotEq(callTick, tickSize);
        assertNotEq(callLot, lotSize);

        // case [2]
        (minPrice,,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, 1);
        _ensurePrice(address(PAIR), minPrice);
        _ensurePrice(targetPair, sizes[1].minPrice);

        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: 4e6}();
        assertEq(TICK_SIZE_SETTER.lastUpdateTimestamp(address(PAIR)), TICK_SIZE_SETTER.lastUpdateTimestamp(targetPair));

        (tickSize, lotSize) = PAIR.tickSizes();
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, 2);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, 2);
        assertEq(callTick, tickSize, "[2] address(PAIR) tick size");
        assertEq(callLot, lotSize, "[2] address(PAIR) lot size");

        (tickSize, lotSize) = PairImpl(targetPair).tickSizes();
        assertEq(tickSize, sizes[2].tickSize, "[2] targetPair tick size");
        assertEq(lotSize, sizes[2].lotSize, "[2] targetPair lot size");

        assertNotEq(callTick, tickSize);
        assertNotEq(callLot, lotSize);

        // remove
        vm.prank(OWNER);
        TICK_SIZE_SETTER.removeResolvedSizesByPair(targetPair);

        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: 4e6}();
        assertEq(TICK_SIZE_SETTER.lastUpdateTimestamp(address(PAIR)), TICK_SIZE_SETTER.lastUpdateTimestamp(targetPair));

        (tickSize, lotSize) = PAIR.tickSizes();
        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, 2);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, 2);
        assertEq(callTick, tickSize, "[remove] address(PAIR) tick size");
        assertEq(callLot, lotSize, "[remove] address(PAIR) lot size");

        (tickSize, lotSize) = PairImpl(targetPair).tickSizes();
        assertNotEq(tickSize, sizes[2].tickSize, "[remove] 1 targetPair tick size");
        assertNotEq(lotSize, sizes[2].lotSize, "[remove] 1 targetPair lot size");

        (, callTick,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, 7);
        (,, callLot) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, 7);
        assertEq(callTick, tickSize, "[remove] address(PAIR) tick size");
        assertEq(callLot, lotSize, "[remove] address(PAIR) lot size");
    }

    function _make_market_pair(uint256 marketCount, uint256 pairCount, uint256 price) private {
        vm.startPrank(OWNER);
        for (; _allMarkets.length < marketCount;) {
            T20 erc20 = new T20("QUOTE", "QUOTE", quoteDecimals);
            erc20.approve(address(ROUTER), type(uint256).max);
            address market = CROSS_DEX.createMarket(OWNER, address(erc20), FEE_COLLECTOR, FEE_BPS);
            _allMarkets.push(market);
        }
        vm.roll(block.number + 1);
        for (uint256 i = 0; i < _allMarkets.length; i++) {
            address market = _allMarkets[i];
            address[] storage pairs = _allPairs[market];
            for (; pairs.length < pairCount;) {
                T20 erc20 = new T20("BASE", "BASE", baseDecimals);
                erc20.approve(address(ROUTER), type(uint256).max);
                address pair = MarketImpl(market).createPair(address(erc20), QUOTE_DECIMALS / 1e4, BASE_DECIMALS / 1e4);
                pairs.push(pair);
            }
            vm.roll(block.number + 1);
        }
        for (uint256 i = 0; i < _allMarkets.length; i++) {
            address market = _allMarkets[i];
            address[] storage pairs = _allPairs[market];
            for (uint256 j = 0; j < pairs.length; j++) {
                _ensurePrice(pairs[j], price);
            }
        }
        vm.stopPrank();
    }

    function _all_updates_execute(address pair, uint256 price, uint256 expectIndex) private {
        vm.startPrank(OWNER);

        (, uint256 tickSize,) = TICK_SIZE_SETTER.resolvedSizes(quoteDecimals, expectIndex);
        (,, uint256 lotSize) = TICK_SIZE_SETTER.resolvedSizes(baseDecimals, expectIndex);
        uint256 updateInterval = TICK_SIZE_SETTER.updateInterval();
        _ensurePrice(pair, price);
        vm.warp(block.timestamp + updateInterval);

        vm.startBroadcast();
        TICK_SIZE_SETTER.allUpdates();
        vm.stopBroadcast();

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
