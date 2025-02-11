// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./DEXBase.t.sol";

contract DEXCheckTest is DEXBaseTest {
    function setUp() external {
        FEE_PERMIL = 50;

        _deploy(18, 18, 1e2, 1e4);
    }

    // 주문을 취소하면 데이터가 삭제 되고, 토큰을 돌려준다.
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
            uint256 sellOrderId = ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
            assertEq(0, BASE.balanceOf(seller));

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = sellOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);

            vm.startPrank(seller);
            ROUTER.cancel(address(PAIR), cancelOrderIds);
            assertEq(amount, BASE.balanceOf(seller));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
        {
            // cancel buy
            vm.startPrank(buyer);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
            assertEq(0, QUOTE.balanceOf(buyer));

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = buyOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);

            vm.startPrank(buyer);
            ROUTER.cancel(address(PAIR), cancelOrderIds);
            assertEq(volume, QUOTE.balanceOf(buyer));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
    }

    // 거래가 일부분 이루어진 주문또한 취소하면 데이터가 삭제 되고, 남은 토큰을 돌려준다.
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
            uint256 sellOrderId = ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
            assertEq(0, BASE.balanceOf(seller));
            vm.startPrank(OWNER);
            ROUTER.limitBuy(address(PAIR), price, amount / 2, 0, 0);

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = sellOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);

            vm.startPrank(seller);
            ROUTER.cancel(address(PAIR), cancelOrderIds);
            assertEq(amount / 2, BASE.balanceOf(seller));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
        {
            // cancel buy
            vm.startPrank(buyer);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
            assertEq(0, QUOTE.balanceOf(buyer));

            vm.startPrank(OWNER);
            ROUTER.marketSell(address(PAIR), amount / 2, 0);

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = buyOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);

            vm.startPrank(buyer);
            ROUTER.cancel(address(PAIR), cancelOrderIds);
            assertEq(volume / 2, QUOTE.balanceOf(buyer));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
    }

    // 거래 취소는 해당 주문의 OWNER 만 가능하다.
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
            uint256 sellOrderId = ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
            assertEq(0, BASE.balanceOf(seller));

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = sellOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);

            vm.startPrank(hacker);
            vm.expectRevert(abi.encodeWithSignature("PairNotOwner(uint256,address)", sellOrderId, hacker));
            ROUTER.cancel(address(PAIR), cancelOrderIds);
            assertEq(0, BASE.balanceOf(seller));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(1, sellPrices.length);
            assertEq(0, buyPrices.length);

            vm.startPrank(seller);
            ROUTER.cancel(address(PAIR), cancelOrderIds);

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(0, buyPrices.length);
        }
        {
            // cancel buy
            vm.startPrank(buyer);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            uint256 buyOrderId = ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
            assertEq(0, QUOTE.balanceOf(buyer));

            uint256[] memory cancelOrderIds = new uint256[](1);
            cancelOrderIds[0] = buyOrderId;

            (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);

            vm.startPrank(hacker);
            vm.expectRevert(abi.encodeWithSignature("PairNotOwner(uint256,address)", buyOrderId, hacker));
            ROUTER.cancel(address(PAIR), cancelOrderIds);
            assertEq(0, QUOTE.balanceOf(buyer));

            (sellPrices, buyPrices) = PAIR.ticks();
            assertEq(0, sellPrices.length);
            assertEq(1, buyPrices.length);
        }
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
        ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 fee = Math.mulDiv(volume, FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid quote fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // 잔액 확인
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
        ROUTER.limitSell(address(PAIR), price, sellAmount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitBuy(address(PAIR), price, buyAmount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 fee = Math.mulDiv(matchVolume, FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // 잔액 확인
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
        ROUTER.limitSell(address(PAIR), price, sellAmount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitBuy(address(PAIR), price, buyAmount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        uint256 fee = Math.mulDiv(matchVolume, FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // 잔액 확인
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
        ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        uint256 fee = Math.mulDiv(volume, FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // 잔액 확인
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
        ROUTER.limitBuy(address(PAIR), price, buyAmount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitSell(address(PAIR), price, sellAmount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        uint256 fee = Math.mulDiv(matchVolume, FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // 잔액 확인
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
        ROUTER.limitBuy(address(PAIR), price, buyAmount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitSell(address(PAIR), price, sellAmount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        uint256 fee = Math.mulDiv(matchVolume, FEE_PERMIL, 1000);

        // 수수료 확인
        assertEq(fee, QUOTE.balanceOf(FEE_COLLECTOR), "invalid fee");
        assertEq(0, BASE.balanceOf(FEE_COLLECTOR), "invalid base fee");

        // 잔액 확인
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
        ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.marketBuy(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // 정확이 매칭 되었는지 확인

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
        ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller1));
        vm.startPrank(seller2);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(1, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.marketBuy(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // 정확이 매칭 되었는지 확인

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
        ROUTER.limitSell(address(PAIR), price1, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller1));
        vm.startPrank(seller2);
        BASE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitSell(address(PAIR), price2, amount, 0, 0);
        assertEq(0, BASE.balanceOf(seller2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(2, sellPrices.length);
        assertEq(0, buyPrices.length);

        vm.startPrank(buyer);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.marketBuy(address(PAIR), volume, 0);
        assertEq(0, QUOTE.balanceOf(buyer)); // 정확이 매칭 되었는지 확인

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
        ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);

        ROUTER.marketSell(address(PAIR), amount, 0);
        assertEq(0, BASE.balanceOf(seller)); // 정확이 매칭 되었는지 확인

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
        ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer1));

        vm.startPrank(buyer2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(1, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);

        ROUTER.marketSell(address(PAIR), amount * 2, 0);
        assertEq(0, BASE.balanceOf(seller)); // 정확이 매칭 되었는지 확인

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
        ROUTER.limitBuy(address(PAIR), price1, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer1));

        vm.startPrank(buyer2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.limitBuy(address(PAIR), price2, amount, 0, 0);
        assertEq(0, QUOTE.balanceOf(buyer2));

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        assertEq(0, sellPrices.length);
        assertEq(2, buyPrices.length);

        vm.startPrank(seller);
        BASE.approve(address(ROUTER), type(uint256).max);

        ROUTER.marketSell(address(PAIR), amount * 2, 0);
        assertEq(0, BASE.balanceOf(seller)); // 정확이 매칭 되었는지 확인

        (sellPrices, buyPrices) = PAIR.ticks();
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
        uint256 orderId = ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
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
        uint256 orderId = ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
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

    function test_check_ordersByPrices() external {
        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.limitSell(address(PAIR), 12 ether, 1 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 12 ether, 1 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 12 ether, 1 ether, 0, 0);

        ROUTER.limitSell(address(PAIR), 11 ether, 1 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 11 ether, 1 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 11 ether, 1 ether, 0, 0);

        ROUTER.limitSell(address(PAIR), 10 ether, 1 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 10 ether, 1 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 10 ether, 1 ether, 0, 0);

        ROUTER.limitBuy(address(PAIR), 9 ether, 1 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 9 ether, 1 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 9 ether, 1 ether, 0, 0);

        ROUTER.limitBuy(address(PAIR), 8 ether, 1 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 8 ether, 1 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 8 ether, 1 ether, 0, 0);

        ROUTER.limitBuy(address(PAIR), 7 ether, 1 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 7 ether, 1 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 7 ether, 1 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 7 ether, 1 ether, 0, 0);
        vm.stopPrank();

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        uint256[][] memory sellOrders = PAIR.ordersByPrices(IPair.OrderSide.SELL, sellPrices);
        uint256[][] memory buyOrders = PAIR.ordersByPrices(IPair.OrderSide.BUY, buyPrices);

        uint256 l1 = sellOrders.length;
        for (uint256 i = 0; i < l1; i++) {
            uint256 l2 = sellOrders[i].length;
            for (uint256 j = 0; j < l2; j++) {
                console.log(sellOrders[i][j]);
            }
        }

        l1 = buyOrders.length;
        for (uint256 i = 0; i < l1; i++) {
            uint256 l2 = buyOrders[i].length;
            for (uint256 j = 0; j < l2; j++) {
                console.log(buyOrders[i][j]);
            }
        }
    }
}
