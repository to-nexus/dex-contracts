// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import "./DEXBase.t.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";

contract DEXCheckTest is DEXBaseTest {
    function setUp() external {
        FEE_BPS = 99; // 0.99%

        _deploy(18, 18, 1e2, 1e4);
    }

    // When an order is canceled, the data is deleted and the tokens are refunded.
    function test_check_cancel_case1() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        {
            // cancel sell
            vm.startPrank(seller);
            BASE.approve(address(ROUTER), type(uint256).max);
            uint256 sellOrderId = ROUTER.submitSellLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, BASE.balanceOf(seller));

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = sellOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);

            vm.startPrank(seller);
            ROUTER.cancelOrder(address(PAIR), cancelOrderIds);
            assertEq(amount, BASE.balanceOf(seller));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
        {
            // cancel buy
            vm.startPrank(buyer);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.submitBuyLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, QUOTE.balanceOf(buyer));

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = buyOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);

            vm.startPrank(buyer);
            ROUTER.cancelOrder(address(PAIR), cancelOrderIds);
            assertEq(volume, QUOTE.balanceOf(buyer));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
    }

    // If a partially filled order is canceled, the data is deleted, and the remaining tokens are refunded.
    function test_check_cancel_case2() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);
        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.stopPrank();

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        {
            // cancel sell
            vm.startPrank(seller);
            BASE.approve(address(ROUTER), type(uint256).max);
            uint256 sellOrderId = ROUTER.submitSellLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, BASE.balanceOf(seller));
            vm.startPrank(OWNER);
            ROUTER.submitBuyLimit(
                address(PAIR), price, amount / 2, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = sellOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);

            vm.startPrank(seller);
            ROUTER.cancelOrder(address(PAIR), cancelOrderIds);
            assertEq(amount / 2, BASE.balanceOf(seller));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
        {
            // cancel buy
            vm.startPrank(buyer);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.submitBuyLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, QUOTE.balanceOf(buyer));

            vm.startPrank(OWNER);
            ROUTER.submitSellMarket(address(PAIR), amount / 2, 0);

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = buyOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);

            vm.startPrank(buyer);
            ROUTER.cancelOrder(address(PAIR), cancelOrderIds);
            assertEq(volume / 2, QUOTE.balanceOf(buyer));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
    }

    // Order cancellation is only allowed by the OWNER of the order.
    function test_check_cancel_case3() external {
        address seller = address(0x1);
        address buyer = address(0x2);
        address hacker = address(0x3);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        {
            // cancel sell
            vm.startPrank(seller);
            BASE.approve(address(ROUTER), type(uint256).max);
            uint256 sellOrderId = ROUTER.submitSellLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, BASE.balanceOf(seller));

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = sellOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);

            vm.startPrank(hacker);
            vm.expectRevert(abi.encodeWithSignature("PairNotOwner(uint256,address)", sellOrderId, hacker));
            ROUTER.cancelOrder(address(PAIR), cancelOrderIds);
            assertEq(0, BASE.balanceOf(seller));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);

            vm.startPrank(seller);
            ROUTER.cancelOrder(address(PAIR), cancelOrderIds);

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
        {
            // cancel buy
            vm.startPrank(buyer);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.submitBuyLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, QUOTE.balanceOf(buyer));

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = buyOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);

            vm.startPrank(hacker);
            vm.expectRevert(abi.encodeWithSignature("PairNotOwner(uint256,address)", buyOrderId, hacker));
            ROUTER.cancelOrder(address(PAIR), cancelOrderIds);
            assertEq(0, QUOTE.balanceOf(buyer));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);
        }
    }

    // When the PAIR is in a PAUSE state, the contract owner can forcibly cancel orders.
    function test_check_emergency_cancel_case1() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price1 = _toQuote(3);
        uint256 price2 = _toQuote(2);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price2, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        uint256[] memory _orderIds = new uint256[](2);
        {
            // cancel sell
            vm.startPrank(seller);
            BASE.approve(address(ROUTER), type(uint256).max);
            uint256 sellOrderId = ROUTER.submitSellLimit(
                address(PAIR), price1, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, BASE.balanceOf(seller));
            _orderIds[0] = sellOrderId;
        }
        {
            // cancel buy
            vm.startPrank(buyer);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.submitBuyLimit(
                address(PAIR), price2, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, QUOTE.balanceOf(buyer));
            _orderIds[1] = buyOrderId;
        }
        vm.roll(block.number + 1);

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.startPrank(OWNER);
        PAIR.setPause(true);
        PAIR.emergencyCancelOrder(_orderIds);

        // Check if the order information has been removed.
        for (uint256 i = 0; i < 2; i++) {
            IPair.Order memory order = PAIR.orderById(_orderIds[i]);
            assertEq(address(0), order.owner);
        }
        // check reserve
        assertEq(0, PAIR.baseReserve());
        assertEq(0, PAIR.quoteReserve());
        // check ticks
        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
        // check balance (PAIR)
        assertEq(0, BASE.balanceOf(address(PAIR)));
        assertEq(0, QUOTE.balanceOf(address(PAIR)));
        // check balance (USER)
        assertEq(amount, BASE.balanceOf(seller));
        assertEq(volume, QUOTE.balanceOf(buyer));
    }

    // If the PAIR is not in a PAUSE state, even the contract owner cannot forcibly cancel orders.
    function test_check_emergency_cancel_case2() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price1 = _toQuote(3);
        uint256 price2 = _toQuote(2);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price2, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        uint256[] memory _orderIds = new uint256[](2);
        {
            // cancel sell
            vm.startPrank(seller);
            BASE.approve(address(ROUTER), type(uint256).max);
            uint256 sellOrderId = ROUTER.submitSellLimit(
                address(PAIR), price1, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, BASE.balanceOf(seller));
            _orderIds[0] = sellOrderId;
        }
        {
            // cancel buy
            vm.startPrank(buyer);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.submitBuyLimit(
                address(PAIR), price2, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, QUOTE.balanceOf(buyer));
            _orderIds[1] = buyOrderId;
        }
        vm.roll(block.number + 1);

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.startPrank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("ExpectedPause()"));
        PAIR.emergencyCancelOrder(_orderIds);
    }

    // When the PAIR is in a PAUSE state, the order owner can cancel the order using 'cancel'.
    function test_check_emergency_cancel_case3() external {
        address user = address(0x1);

        uint256 price1 = _toQuote(3);
        uint256 price2 = _toQuote(2);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price2, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(user, amount);
        vm.prank(OWNER);
        QUOTE.transfer(user, volume);

        uint256[] memory _orderIds = new uint256[](2);
        vm.startPrank(user);
        {
            // sell
            BASE.approve(address(ROUTER), type(uint256).max);
            uint256 sellOrderId = ROUTER.submitSellLimit(
                address(PAIR), price1, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, BASE.balanceOf(user));
            _orderIds[0] = sellOrderId;
        }
        {
            // buy
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.submitBuyLimit(
                address(PAIR), price2, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            assertEq(0, QUOTE.balanceOf(user));
            _orderIds[1] = buyOrderId;
        }
        vm.stopPrank();
        vm.roll(block.number + 1);

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.prank(OWNER);
        PAIR.setPause(true);
        vm.prank(user);
        ROUTER.cancelOrder(address(PAIR), _orderIds);

        // Check if the order information has been removed.
        for (uint256 i = 0; i < 2; i++) {
            IPair.Order memory order = PAIR.orderById(_orderIds[i]);
            assertEq(address(0), order.owner);
        }
        // check reserve
        assertEq(0, PAIR.baseReserve());
        assertEq(0, PAIR.quoteReserve());
        // check ticks
        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
        // check balance (PAIR)
        assertEq(0, BASE.balanceOf(address(PAIR)));
        assertEq(0, QUOTE.balanceOf(address(PAIR)));
        // check balance (USER)
        assertEq(amount, BASE.balanceOf(user));
        assertEq(volume, QUOTE.balanceOf(user));
    }

    // [AMOUNT] sell == buy, seller == maker
    function test_check_fee_case1() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 fee = Math.mulDiv(volume, FEE_BPS, 10000);

        // check fee
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid quote fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // check balance
        assertEq(volume - fee, QUOTE.balanceOf(seller));
        assertEq(amount, BASE.balanceOf(buyer));
    }

    // [AMOUNT] sell < buy, seller == maker
    function test_check_fee_case2() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 sellAmount = _toBase(700);
        uint256 buyAmount = _toBase(900);
        uint256 buyerVolume = Math.mulDiv(price, buyAmount, BASE_DECIMALS);

        uint256 matchAmount = Math.min(sellAmount, buyAmount);
        uint256 matchVolume = Math.mulDiv(price, matchAmount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, sellAmount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, buyerVolume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(
            address(PAIR), price, sellAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(
            address(PAIR), price, buyAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 fee = Math.mulDiv(matchVolume, FEE_BPS, 10000);

        // check fee
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // check balance
        assertEq(matchVolume - fee, QUOTE.balanceOf(seller));
        assertEq(matchAmount, BASE.balanceOf(buyer));
    }

    // [AMOUNT] sell > buy, seller == maker
    function test_check_fee_case3() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 sellAmount = _toBase(700);
        uint256 buyAmount = _toBase(300);
        uint256 buyerVolume = Math.mulDiv(price, buyAmount, BASE_DECIMALS);

        uint256 matchAmount = Math.min(sellAmount, buyAmount);
        uint256 matchVolume = Math.mulDiv(price, matchAmount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, sellAmount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, buyerVolume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(
            address(PAIR), price, sellAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(
            address(PAIR), price, buyAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 fee = Math.mulDiv(matchVolume, FEE_BPS, 10000);

        // check fee
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // check balance
        assertEq(matchVolume - fee, QUOTE.balanceOf(seller));
        assertEq(matchAmount, BASE.balanceOf(buyer));
    }

    // [AMOUNT] buy == sell, seller == taker
    function test_check_fee_case4() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(7);
        uint256 amount = _toBase(300);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, BASE.balanceOf(seller));

        uint256 fee = Math.mulDiv(volume, FEE_BPS, 10000);

        // check fee
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // check balance
        assertEq(volume - fee, QUOTE.balanceOf(seller));
        assertEq(amount, BASE.balanceOf(buyer));
    }

    // [AMOUNT] buy < sell, seller == taker
    function test_check_fee_case5() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 sellAmount = _toBase(700);
        uint256 buyAmount = _toBase(500);
        uint256 buyerVolume = Math.mulDiv(price, buyAmount, BASE_DECIMALS);

        uint256 matchAmount = Math.min(sellAmount, buyAmount);
        uint256 matchVolume = Math.mulDiv(price, matchAmount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, sellAmount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, buyerVolume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(
            address(PAIR), price, buyAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(
            address(PAIR), price, sellAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(0, BASE.balanceOf(seller));

        uint256 fee = Math.mulDiv(matchVolume, FEE_BPS, 10000);

        // check fee
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // check balance
        assertEq(matchVolume - fee, QUOTE.balanceOf(seller));
        assertEq(matchAmount, BASE.balanceOf(buyer));
    }

    // [AMOUNT] buy > sell, seller == taker
    function test_check_fee_case6() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 sellAmount = _toBase(700);
        uint256 buyAmount = _toBase(900);
        uint256 buyerVolume = Math.mulDiv(price, buyAmount, BASE_DECIMALS);

        uint256 matchAmount = Math.min(sellAmount, buyAmount);
        uint256 matchVolume = Math.mulDiv(price, matchAmount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, sellAmount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, buyerVolume);

        assertEq(0, QUOTE.balanceOf(FEE_COLLECTOR));
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(
            address(PAIR), price, buyAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(
            address(PAIR), price, sellAmount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertEq(0, BASE.balanceOf(seller));

        uint256 fee = Math.mulDiv(matchVolume, FEE_BPS, 10000);

        // check fee
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // check balance
        assertEq(matchVolume - fee, QUOTE.balanceOf(seller));
        assertEq(matchAmount, BASE.balanceOf(buyer));
    }

    function test_check_gas_case1() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, BASE.balanceOf(seller));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.submitBuyMarket(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // Verify if the match was exact.

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    function test_check_gas_case2() external {
        address seller1 = address(0x1);
        address seller2 = address(0x2);
        address buyer = address(0x3);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount * 2, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller1, amount);
        vm.prank(OWNER);
        BASE.transfer(seller2, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(seller1);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, BASE.balanceOf(seller1));
        vm.startPrank(seller2);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, BASE.balanceOf(seller2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.submitBuyMarket(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // Verify if the match was exact.

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    function test_check_gas_case3() external {
        address seller1 = address(0x1);
        address seller2 = address(0x2);
        address buyer = address(0x3);

        uint256 price1 = _toQuote(3);
        uint256 price2 = _toQuote(5);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price1, amount, BASE_DECIMALS) + Math.mulDiv(price2, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller1, amount);
        vm.prank(OWNER);
        BASE.transfer(seller2, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(seller1);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(address(PAIR), price1, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, BASE.balanceOf(seller1));
        vm.startPrank(seller2);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(address(PAIR), price2, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, BASE.balanceOf(seller2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(2, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.submitBuyMarket(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // Verify if the match was exact.

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    function test_check_gas_case4() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);

        ROUTER.submitSellMarket(address(PAIR), amount, 0);
        assertEq(0, BASE.balanceOf(seller)); // Verify if the match was exact.

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    function test_check_gas_case5() external {
        address seller = address(0x1);
        address buyer1 = address(0x2);
        address buyer2 = address(0x3);

        uint256 price = _toQuote(3);
        uint256 amount = _toBase(700);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount * 2);
        vm.prank(OWNER);
        QUOTE.transfer(buyer1, volume);
        vm.prank(OWNER);
        QUOTE.transfer(buyer2, volume);

        vm.startPrank(buyer1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, QUOTE.balanceOf(buyer1));

        vm.startPrank(buyer2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, QUOTE.balanceOf(buyer2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);

        ROUTER.submitSellMarket(address(PAIR), amount * 2, 0);
        assertEq(0, BASE.balanceOf(seller)); // Verify if the match was exact.

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    function test_check_gas_case6() external {
        address seller = address(0x1);
        address buyer1 = address(0x2);
        address buyer2 = address(0x3);

        uint256 price1 = _toQuote(3);
        uint256 price2 = _toQuote(1);
        uint256 amount = _toBase(700);
        uint256 volume1 = Math.mulDiv(price1, amount, BASE_DECIMALS);
        uint256 volume2 = Math.mulDiv(price2, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        BASE.transfer(seller, amount * 2);
        vm.prank(OWNER);
        QUOTE.transfer(buyer1, volume1);
        vm.prank(OWNER);
        QUOTE.transfer(buyer2, volume2);

        vm.startPrank(buyer1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(address(PAIR), price1, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, QUOTE.balanceOf(buyer1));

        vm.startPrank(buyer2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(address(PAIR), price2, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0);
        assertEq(0, QUOTE.balanceOf(buyer2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(2, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);

        ROUTER.submitSellMarket(address(PAIR), amount * 2, 0);
        assertEq(0, BASE.balanceOf(seller)); // Verify if the match was exact.

        (sellPrices, buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    function test_check_gas_case7() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(1);
        uint256 loopCount = 50;

        vm.prank(OWNER);
        BASE.transfer(seller, amount * loopCount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, price * loopCount);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        for (uint256 i = 0; i < loopCount; i++) {
            ROUTER.submitSellLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }
        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyMarket(address(PAIR), price * loopCount, 0);

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    function test_check_gas_case8() external {
        address seller = address(0x1);
        address buyer = address(0x2);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(1);
        uint256 loopCount = 50;

        vm.prank(OWNER);
        BASE.transfer(seller, amount * loopCount);
        vm.prank(OWNER);
        QUOTE.transfer(buyer, price * loopCount);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        for (uint256 i = 0; i < loopCount; i++) {
            ROUTER.submitBuyLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }
        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellMarket(address(PAIR), amount * loopCount, 0);

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(0, buyPrices.length);
    }

    function test_check_skim_case1() external {
        address seller = address(0x1);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);

        vm.prank(OWNER);
        BASE.transfer(seller, amount);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId = ROUTER.submitSellLimit(
            address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertNotEq(0, orderId);
        vm.stopPrank();

        assertEq(BASE.balanceOf(address(PAIR)), amount);

        vm.startPrank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidReserve(address)", address(BASE)));
        PAIR.skim(BASE, OWNER, amount);

        BASE.transfer(address(PAIR), amount);
        address receiver = address(bytes20("RECEIVER"));

        PAIR.skim(BASE, receiver, amount);
        assertEq(BASE.balanceOf(receiver), amount);
        vm.stopPrank();
    }

    function test_check_skim_case2() external {
        address buyer = address(0x1);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(100);
        uint256 volume = Math.mulDiv(price, amount, BASE_DECIMALS);

        vm.prank(OWNER);
        QUOTE.transfer(buyer, volume);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId = ROUTER.submitBuyLimit(
            address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        assertNotEq(0, orderId);
        vm.stopPrank();
        assertEq(QUOTE.balanceOf(address(PAIR)), volume);

        vm.startPrank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("PairInvalidReserve(address)", address(QUOTE)));
        PAIR.skim(QUOTE, OWNER, volume - 1);
        QUOTE.transfer(address(PAIR), amount);
        address receiver = address(bytes20("RECEIVER"));

        PAIR.skim(QUOTE, receiver, amount);
        assertEq(QUOTE.balanceOf(receiver), amount);
        vm.stopPrank();
    }

    function test_check_skim_case3() external {
        vm.startPrank(OWNER);

        IERC20 erc20 = IERC20(address(new T20("ERC20", "E20", 18)));
        uint256 amount = 100 ether;
        erc20.transfer(address(PAIR), amount);

        address receiver = address(bytes20("RECEIVER"));
        PAIR.skim(erc20, receiver, amount);

        vm.stopPrank();
    }

    function test_check_orders_prices() external {
        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        // enroll sell prices

        ROUTER.submitSellLimit(
            address(PAIR), 12 ether, 1 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        ROUTER.submitSellLimit(
            address(PAIR), 10 ether, 1 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        ROUTER.submitSellLimit(
            address(PAIR), 11 ether, 1 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        // enroll buy prices

        ROUTER.submitBuyLimit(
            address(PAIR), 7 ether, 1 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        ROUTER.submitBuyLimit(
            address(PAIR), 9 ether, 1 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        ROUTER.submitBuyLimit(
            address(PAIR), 8 ether, 1 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );

        vm.stopPrank();

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();

        {
            uint256 price = sellPrices[0];
            for (uint256 i = 1; i < sellPrices.length; i++) {
                assertGt(sellPrices[i], price);
                price = sellPrices[i];
            }
        }
        {
            uint256 price = buyPrices[0];
            for (uint256 i = 1; i < buyPrices.length; i++) {
                assertLt(buyPrices[i], price);
            }
        }
    }

    mapping(uint8 => bool) _testFuzz_check_orders_prices_set;

    function testFuzz_check_orders_prices(uint8[30] calldata _prices) external {
        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        uint8 min = type(uint8).max;
        uint8 max = 0;
        uint256 length;
        for (uint256 i = 0; i < _prices.length; i++) {
            uint8 price = _prices[i];
            if (price == 0 || _testFuzz_check_orders_prices_set[price]) continue;
            _testFuzz_check_orders_prices_set[price] = true;
            ++length;
            if (price < min) min = price;
            if (price > max) max = price;
        }
        vm.assume(length > 10);

        uint256 mid = uint256(min) + uint256(max) / 2;
        for (uint256 i = 0; i < _prices.length; i++) {
            if (_prices[i] == 0) continue;
            uint8 price = _prices[i];
            if (price > mid) {
                ROUTER.submitSellLimit(
                    address(PAIR),
                    price * QUOTE_DECIMALS,
                    1 ether,
                    IPair.LimitConstraints.GOOD_TILL_CANCEL,
                    _searchPrices,
                    0
                );
            } else {
                ROUTER.submitBuyLimit(
                    address(PAIR),
                    price * QUOTE_DECIMALS,
                    1 ether,
                    IPair.LimitConstraints.GOOD_TILL_CANCEL,
                    _searchPrices,
                    0
                );
            }
        }

        vm.stopPrank();

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();

        {
            uint256 price = sellPrices[0];
            for (uint256 i = 1; i < sellPrices.length; i++) {
                assertGt(sellPrices[i], price);
                price = sellPrices[i];
            }
        }
        {
            uint256 price = buyPrices[0];
            for (uint256 i = 1; i < buyPrices.length; i++) {
                assertLt(buyPrices[i], price);
            }
        }
    }

    function test_check_setTickSizes_case() external {
        (uint256 tick, uint256 lot) = PAIR.tickSizes();
        assertNotEq(tick, 1e20);
        assertNotEq(lot, 1e20);

        address setter = address(bytes20("SETTER"));

        // check roles
        vm.prank(setter);
        vm.expectRevert(abi.encodeWithSignature("CrossDexUnauthorizedChangeTickSizes(address)", setter));
        PAIR.setTickSize(1e20, 1e20);
        // grant roles
        assertNotEq(setter, CROSS_DEX.tickSizeSetter());
        vm.prank(OWNER);
        CROSS_DEX.setTickSizeSetter(setter);
        // check success
        vm.prank(setter);
        PAIR.setTickSize(1e20, 1e20);
        (tick, lot) = PAIR.tickSizes();
        assertEq(tick, 1e20);
        assertEq(lot, 1e20);
    }

    function test_check_isMarket() external {
        assertTrue(CROSS_DEX.isMarket(address(MARKET)));
        vm.expectRevert();
        CROSS_DEX.isMarket(address(1));

        address pairImpl = address(new PairImpl());
        address marketImpl = address(new MarketImpl());
        address proxy = address(new ERC1967Proxy(marketImpl, hex""));

        MarketImpl market = MarketImpl(proxy);
        market.initialize(OWNER, address(ROUTER), address(QUOTE), pairImpl, FEE_COLLECTOR, FEE_BPS);

        assertFalse(CROSS_DEX.isMarket(address(market)));
    }

    function test_check_max_match_count_case1() external {
        address seller = address(1);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(1);
        uint256 maxMatchCount = 10;

        vm.prank(OWNER);
        BASE.transfer(seller, amount * (maxMatchCount + 1));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        for (uint256 i = 0; i < maxMatchCount + 1; i++) {
            ROUTER.submitSellLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }
        vm.stopPrank();

        vm.startPrank(OWNER);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyMarket(address(PAIR), QUOTE.balanceOf(OWNER), maxMatchCount);
        vm.stopPrank();

        uint256 baseReserve = PAIR.baseReserve();
        assertEq(amount, baseReserve);
    }

    function test_check_max_match_count_case2() external {
        address buyer = address(1);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(1);
        uint256 oneVolume = Math.mulDiv(price, amount, BASE_DECIMALS);
        uint256 maxMatchCount = 10;

        vm.prank(OWNER);
        QUOTE.transfer(buyer, oneVolume * (maxMatchCount + 1));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        for (uint256 i = 0; i < maxMatchCount + 1; i++) {
            ROUTER.submitBuyLimit(
                address(PAIR), price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
        }
        vm.stopPrank();

        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 baseBalance = BASE.balanceOf(OWNER);
        uint256 baseTickSize = PAIR.lotSize();
        uint256 sellAmount = baseBalance - (baseBalance % baseTickSize);
        ROUTER.submitSellMarket(address(PAIR), sellAmount, maxMatchCount);
        vm.stopPrank();

        uint256 quoteReserve = PAIR.quoteReserve();
        assertEq(oneVolume, quoteReserve);
    }

    function test_check_slot_matchedprice() external pure {
        bytes32 _matchedprice =
            keccak256(abi.encode(uint256(keccak256("crossdex.pair.matchedprice")) - 1)) & ~bytes32(uint256(0xff));

        console.logBytes32(_matchedprice);
    }

    function test_check_slot_feecollector() external pure {
        bytes32 _feecollector =
            keccak256(abi.encode(uint256(keccak256("crossdex.pair.feecollector")) - 1)) & ~bytes32(uint256(0xff));

        console.logBytes32(_feecollector);
    }

    function test_check_slot_feebps() external pure {
        bytes32 _feebps =
            keccak256(abi.encode(uint256(keccak256("crossdex.pair.feebps")) - 1)) & ~bytes32(uint256(0xff));

        console.logBytes32(_feebps);
    }

    function test_check_set_matchedprice_case1() external {
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
        for (uint256 i = 0; i < 10; i++) {
            uint256 _price = price + (priceStep * i);
            ROUTER.submitBuyLimit(
                address(PAIR), _price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            // The latest price is updated because the orders are matched.
            assertEq(_price, PAIR.matchedPrice());
        }
    }

    function test_check_set_matchedprice_case2() external {
        vm.startPrank(OWNER);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        BASE.approve(address(ROUTER), type(uint256).max);

        uint256 price = _toQuote(1);
        uint256 amount = _toBase(1);
        uint256 priceStep = price / 10;

        // buy orders
        for (uint256 i = 0; i < 10; i++) {
            ROUTER.submitBuyLimit(
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
        for (uint256 i = 10; i != 0;) {
            uint256 _price = price + (priceStep * --i);
            ROUTER.submitSellLimit(
                address(PAIR), _price, amount, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
            );
            // The latest price is updated because the orders are matched.
            assertEq(_price, PAIR.matchedPrice());
        }
    }

    function test_check_account_reserve_case1() external {
        address seller1 = address(0x1);
        address seller2 = address(0x2);
        address seller3 = address(0x3);

        vm.startPrank(OWNER);
        MARKET.setFeeBps(0);

        QUOTE.approve(address(ROUTER), type(uint256).max);
        BASE.transfer(seller1, 100 ether);
        BASE.transfer(seller2, 100 ether);
        BASE.transfer(seller3, 100 ether);
        vm.stopPrank();
        vm.startPrank(seller1);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(
            address(PAIR), 1 ether, 100 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        {
            (uint256 baseReserve,) = PAIR.accountReserves(seller1);
            assertEq(100 ether, baseReserve);
        }
        vm.startPrank(seller2);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(
            address(PAIR), 1 ether, 100 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        {
            (uint256 baseReserve,) = PAIR.accountReserves(seller2);
            assertEq(100 ether, baseReserve);
        }
        vm.startPrank(seller3);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellLimit(
            address(PAIR), 1 ether, 100 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        {
            (uint256 baseReserve,) = PAIR.accountReserves(seller3);
            assertEq(100 ether, baseReserve);
        }

        // check reduce by trade
        vm.startPrank(OWNER);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyMarket(address(PAIR), 50 ether, 0);

        {
            (uint256 baseReserve,) = PAIR.accountReserves(seller1);
            assertEq(50 ether, baseReserve);
            assertEq(50 ether, QUOTE.balanceOf(seller1));
        }
        {
            (uint256 baseReserve,) = PAIR.accountReserves(seller2);
            assertEq(100 ether, baseReserve);
        }
        {
            (uint256 baseReserve,) = PAIR.accountReserves(seller3);
            assertEq(100 ether, baseReserve);
        }
        assertEq(250 ether, PAIR.baseReserve());

        // check reduce by cancel
        {
            uint256[] memory orderIds = new uint256[](1);
            orderIds[0] = 2;
            vm.startPrank(seller2);
            ROUTER.cancelOrder(address(PAIR), orderIds);
            {
                (uint256 baseReserve,) = PAIR.accountReserves(seller1);
                assertEq(50 ether, baseReserve);
            }
            {
                (uint256 baseReserve,) = PAIR.accountReserves(seller2);
                assertEq(0, baseReserve);
            }
            {
                (uint256 baseReserve,) = PAIR.accountReserves(seller3);
                assertEq(100 ether, baseReserve);
            }
            assertEq(150 ether, PAIR.baseReserve());
        }

        // check reduce by emergencyCancel
        {
            vm.startPrank(OWNER);
            PAIR.setPause(true);
            uint256[] memory orderIds = new uint256[](2);
            orderIds[0] = 1;
            orderIds[1] = 3;
            PAIR.emergencyCancelOrder(orderIds);
            {
                (uint256 baseReserve,) = PAIR.accountReserves(seller1);
                assertEq(0, baseReserve);
            }
            {
                (uint256 baseReserve,) = PAIR.accountReserves(seller2);
                assertEq(0, baseReserve);
            }
            {
                (uint256 baseReserve,) = PAIR.accountReserves(seller3);
                assertEq(0, baseReserve);
            }
            assertEq(0, PAIR.baseReserve());
        }
    }

    function test_check_account_reserve_case2() external {
        address buyer1 = address(0x1);
        address buyer2 = address(0x2);
        address buyer3 = address(0x3);

        vm.startPrank(OWNER);
        MARKET.setFeeBps(0);

        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.transfer(buyer1, 100 ether);
        QUOTE.transfer(buyer2, 100 ether);
        QUOTE.transfer(buyer3, 100 ether);
        vm.stopPrank();
        vm.startPrank(buyer1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(
            address(PAIR), 1 ether, 100 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        {
            (, uint256 quoteReserve) = PAIR.accountReserves(buyer1);
            assertEq(100 ether, quoteReserve);
        }
        vm.startPrank(buyer2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(
            address(PAIR), 1 ether, 100 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        {
            (, uint256 quoteReserve) = PAIR.accountReserves(buyer2);
            assertEq(100 ether, quoteReserve);
        }
        vm.startPrank(buyer3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitBuyLimit(
            address(PAIR), 1 ether, 100 ether, IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        {
            (, uint256 quoteReserve) = PAIR.accountReserves(buyer3);
            assertEq(100 ether, quoteReserve);
        }

        // check reduce by trade
        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.submitSellMarket(address(PAIR), 50 ether, 0);

        {
            (, uint256 quoteReserve) = PAIR.accountReserves(buyer1);
            assertEq(50 ether, quoteReserve);
            assertEq(50 ether, BASE.balanceOf(buyer1));
        }
        {
            (, uint256 quoteReserve) = PAIR.accountReserves(buyer2);
            assertEq(100 ether, quoteReserve);
        }
        {
            (, uint256 quoteReserve) = PAIR.accountReserves(buyer3);
            assertEq(100 ether, quoteReserve);
        }
        assertEq(250 ether, PAIR.quoteReserve());
        // check reduce by cancel
        {
            uint256[] memory orderIds = new uint256[](1);
            orderIds[0] = 2;
            vm.startPrank(buyer2);
            ROUTER.cancelOrder(address(PAIR), orderIds);
            {
                (, uint256 quoteReserve) = PAIR.accountReserves(buyer1);
                assertEq(50 ether, quoteReserve);
            }
            {
                (, uint256 quoteReserve) = PAIR.accountReserves(buyer2);
                assertEq(0, quoteReserve);
            }
            {
                (, uint256 quoteReserve) = PAIR.accountReserves(buyer3);
                assertEq(100 ether, quoteReserve);
            }
            assertEq(150 ether, PAIR.quoteReserve());
        }
        // check reduce by emergencyCancel
        {
            vm.startPrank(OWNER);
            PAIR.setPause(true);
            uint256[] memory orderIds = new uint256[](2);
            orderIds[0] = 1;
            orderIds[1] = 3;
            PAIR.emergencyCancelOrder(orderIds);
            {
                (, uint256 quoteReserve) = PAIR.accountReserves(buyer1);
                assertEq(0, quoteReserve);
            }
            {
                (, uint256 quoteReserve) = PAIR.accountReserves(buyer2);
                assertEq(0, quoteReserve);
            }
            {
                (, uint256 quoteReserve) = PAIR.accountReserves(buyer3);
                assertEq(0, quoteReserve);
            }
            assertEq(0, PAIR.quoteReserve());
        }
    }
}
