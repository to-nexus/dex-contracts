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
        vm.warp(86400);
        TICK_SIZE_SETTER.allUpdates();
    }

    function test_ticksize_findPriceIndex() external view {
        uint256 index;
        uint256 price;

        // gte
        price = 10 ** (quoteDecimals - 1);
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 1);

        price = 10 ** (quoteDecimals);
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 2);

        price = 10 ** (quoteDecimals + 1);
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 3);

        price = 10 ** (quoteDecimals + 1);
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 5 * price);
        assertEq(index, 4);

        price = 10 ** (quoteDecimals + 2);
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 5);

        price = 10 ** (quoteDecimals + 2);
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 5 * price);
        assertEq(index, 6);

        price = 10 ** (quoteDecimals + 3);
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 7);

        // lt
        price = 10 ** (quoteDecimals - 1) - 1;
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 0);

        price = 10 ** (quoteDecimals) - 1;
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 1);

        price = 10 ** (quoteDecimals + 1) - 1;
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 2);

        price = 10 ** (quoteDecimals + 1) - 1;
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 5 * price);
        assertEq(index, 3);

        price = 10 ** (quoteDecimals + 2) - 1;
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 4);

        price = 10 ** (quoteDecimals + 2) - 1;
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 5 * price);
        assertEq(index, 5);

        price = 10 ** (quoteDecimals + 3) - 1;
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
        assertEq(index, 6);

        // check 0 and max to ensure no overflow
        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, 0);
        assertEq(index, 0);

        (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, type(uint256).max);
        assertEq(index, 7);
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

    address[] private _allMarkets;
    mapping(address market => address[]) private _allPairs;

    function test_ticksize_all_updates_max_count_case1() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 41 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(1, 41, 10 ** (quoteDecimals - 1));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_max_count_case2() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 40 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(2, 20, 10 ** (quoteDecimals));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_max_count_case3() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 40 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(4, 10, 10 ** (quoteDecimals + 1));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_max_count_case4() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 40 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(10, 4, 5 * (10 ** (quoteDecimals + 1)));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_max_count_case5() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 38 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(19, 2, (10 ** (quoteDecimals + 2)));

        vm.roll(block.number + 1);
        vm.warp(block.timestamp + TICK_SIZE_SETTER.updateInterval());
        TICK_SIZE_SETTER.allUpdates{gas: GAS_LIMIT}();

        (address market_, uint256 startIndex) = TICK_SIZE_SETTER.updatable();
        assertTrue((market_ == address(0) && startIndex == 0));
    }

    function test_ticksize_all_updates_max_count_case6() external {
        uint256 GAS_LIMIT = 4e6;
        // With a gas limit of 4 million (4e6),
        // up to 38 pairs can be updated in a single transaction.

        _allMarkets.push(address(MARKET));
        _allPairs[address(MARKET)].push(address(PAIR));

        _make_market_pair(38, 1, 5 * (10 ** (quoteDecimals + 2)));

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

        _make_market_pair(2, 40, 10 ** (quoteDecimals - 1));

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

        _make_market_pair(3, 40, 10 ** (quoteDecimals - 1));

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
        // gte ~ lt price | ticSize | lotSize
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
            uint256 price;
            uint256 index;
            // call findPriceIndex to check if the size formats are correct
            // gte
            price = 10 ** (quoteDecimals - 1);
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 1);

            price = 10 ** (quoteDecimals);
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 2);

            price = 10 ** (quoteDecimals + 1);
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 3);

            price = 10 ** (quoteDecimals + 2);
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 4);

            price = 10 ** (quoteDecimals + 3);
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 5);

            // lt
            price = 10 ** (quoteDecimals - 1) - 1;
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 0);

            price = 10 ** (quoteDecimals) - 1;
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 1);

            price = 10 ** (quoteDecimals + 1) - 1;
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 2);

            price = 10 ** (quoteDecimals + 2) - 1;
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 3);

            price = 10 ** (quoteDecimals + 3) - 1;
            (index,) = TICK_SIZE_SETTER.findPriceIndex(quoteDecimals, price);
            assertEq(index, 4);
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
        // gte ~ lt price | ticSize | lotSize
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
