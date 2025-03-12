// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./DEXBase.t.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";

contract DexWrapBaseTest is DEXBaseTest {
    using Address for address payable;

    function setUp() external {}

    // [WRAP BASE] create limit sell -> buy market
    function test_wrap_trade_base_case1() external wrapBase {
        address seller = address(0x1);
        address buyer = address(0x2);
        vm.label(seller, "seller");
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        vm.deal(seller, amount);
        vm.prank(seller);
        ROUTER.limitSell{value: amount}(
            address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, WCROSSx.balanceOf(seller));
        vm.assertEq(0, WCROSSx.balanceOf(buyer));
        vm.assertEq(0, WCROSSx.balanceOf(address(ROUTER)));
        vm.assertEq(amount, WCROSSx.balanceOf(address(PAIR)));

        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);
        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketBuy(address(PAIR), volume, 0);

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(amount, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, WCROSSx.balanceOf(seller));
        vm.assertEq(0, WCROSSx.balanceOf(buyer));
        vm.assertEq(0, WCROSSx.balanceOf(address(ROUTER)));
        vm.assertEq(0, WCROSSx.balanceOf(address(PAIR)));
    }

    // [WRAP BASE] create limit buy -> sell market
    function test_wrap_trade_base_case2() external wrapBase {
        address seller = address(0x1);
        address buyer = address(0x2);
        vm.label(seller, "seller");
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);
        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitBuy(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        vm.stopPrank();

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, WCROSSx.balanceOf(seller));
        vm.assertEq(0, WCROSSx.balanceOf(buyer));
        vm.assertEq(0, WCROSSx.balanceOf(address(ROUTER)));
        vm.assertEq(0, WCROSSx.balanceOf(address(PAIR)));

        vm.deal(seller, amount);
        vm.prank(seller);
        ROUTER.marketSell{value: amount}(address(PAIR), amount, 0);

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(amount, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, WCROSSx.balanceOf(seller));
        vm.assertEq(0, WCROSSx.balanceOf(buyer));
        vm.assertEq(0, WCROSSx.balanceOf(address(ROUTER)));
        vm.assertEq(0, WCROSSx.balanceOf(address(PAIR)));
    }

    // [WRAP QUOTE] create limit sell -> buy market
    function test_wrap_trade_quote_case1() external wrapQuote {
        address seller = address(0x1);
        address buyer = address(0x2);
        vm.label(seller, "seller");
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitSell(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        vm.stopPrank();

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, WCROSSx.balanceOf(seller));
        vm.assertEq(0, WCROSSx.balanceOf(buyer));
        vm.assertEq(0, WCROSSx.balanceOf(address(ROUTER)));
        vm.assertEq(0, WCROSSx.balanceOf(address(PAIR)));

        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.deal(buyer, volume);
        vm.prank(buyer);
        ROUTER.marketBuy{value: volume}(address(PAIR), volume, 0);

        vm.assertEq(volume, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, WCROSSx.balanceOf(seller));
        vm.assertEq(0, WCROSSx.balanceOf(buyer));
        vm.assertEq(0, WCROSSx.balanceOf(address(ROUTER)));
        vm.assertEq(0, WCROSSx.balanceOf(address(PAIR)));
    }

    // [WRAP QUOTE] create limit buy -> sell market
    function test_wrap_trade_quote_case2() external wrapQuote {
        address seller = address(0x1);
        address buyer = address(0x2);
        vm.label(seller, "seller");
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.deal(buyer, volume);
        vm.prank(buyer);
        ROUTER.limitBuy{value: volume}(
            address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.assertEq(0, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, WCROSSx.balanceOf(seller));
        vm.assertEq(0, WCROSSx.balanceOf(buyer));
        vm.assertEq(0, WCROSSx.balanceOf(address(ROUTER)));
        vm.assertEq(volume, WCROSSx.balanceOf(address(PAIR)));

        vm.prank(OWNER);
        BASE.transfer(seller, amount);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketSell(address(PAIR), amount, 0);
        vm.stopPrank();

        vm.assertEq(volume, payable(seller).balance);
        vm.assertEq(0, payable(buyer).balance);
        vm.assertEq(0, payable(address(ROUTER)).balance);
        vm.assertEq(0, payable(address(PAIR)).balance);

        vm.assertEq(0, WCROSSx.balanceOf(seller));
        vm.assertEq(0, WCROSSx.balanceOf(buyer));
        vm.assertEq(0, WCROSSx.balanceOf(address(ROUTER)));
        vm.assertEq(0, WCROSSx.balanceOf(address(PAIR)));
    }

    // An EOA cannot hold WETH.sol in token form.
    function test_wrap_check() external wrapBase {
        vm.startPrank(OWNER);
        vm.deal(OWNER, 100);
        uint256 beforeBalance = OWNER.balance;
        assertNotEq(0, beforeBalance);

        payable(address(WCROSSx)).sendValue(10);
        assertEq(0, WCROSSx.balanceOf(OWNER));
        assertEq(beforeBalance, OWNER.balance);
    }

    // Returns an error if msg.value is incorrect.
    function test_wrap_exception_case1() external wrapBase {
        address seller = address(0x1);
        vm.label(seller, "seller");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        vm.deal(seller, amount + 1);
        vm.prank(seller);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        ROUTER.limitSell{value: amount + 1}(
            address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(seller);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        ROUTER.marketSell{value: amount + 1}(address(PAIR), amount, 0);
    }

    // Returns an error if msg.value is incorrect.
    function test_wrap_exception_case2() external wrapQuote {
        address buyer = address(0x1);
        vm.label(buyer, "buyer");

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.deal(buyer, volume + 1);
        vm.prank(buyer);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        ROUTER.limitBuy{value: volume + 1}(
            address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.prank(buyer);
        vm.expectRevert(abi.encodeWithSignature("RouterInvalidValue()"));
        ROUTER.marketBuy{value: volume + 1}(address(PAIR), volume, 0);
    }

    modifier wrapBase() {
        _setWrapBase();
        _;
    }

    modifier wrapQuote() {
        _setWrapQuote();
        _;
    }

    function _setWrapBase() private {
        uint256 quote_tick_size = 1e4;
        uint256 base_tick_size = 1e2;

        _deploy(18, 18, quote_tick_size, base_tick_size);
        BASE = IERC20(address(WCROSSx));

        vm.startPrank(OWNER);
        address pair =
            MARKET.createPair(address(BASE), QUOTE_DECIMALS / quote_tick_size, BASE_DECIMALS / base_tick_size, FEE_BPS);

        PAIR = PairImpl(pair);
        vm.stopPrank();
    }

    function _setWrapQuote() private {
        uint256 quote_tick_size = 1e4;
        uint256 base_tick_size = 1e2;

        _deploy(18, 18, quote_tick_size, base_tick_size);
        QUOTE = IERC20(address(WCROSSx));

        vm.startPrank(OWNER);
        address market = CROSS_DEX.createMarket(OWNER, FEE_COLLECTOR, address(QUOTE));
        MARKET = MarketImpl(market);
        address pair =
            MARKET.createPair(address(BASE), QUOTE_DECIMALS / quote_tick_size, BASE_DECIMALS / base_tick_size, FEE_BPS);
        PAIR = PairImpl(pair);
        vm.stopPrank();
    }
}
