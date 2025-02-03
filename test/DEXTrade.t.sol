// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./DEXBase.t.sol";

contract DEXTradeTest is DEXBaseTest {
    function setUp() external {
        _deploy(18, 18, 1e2, 1e6);
    }

    function test_trade_create_limit_sell() external {
        address user = address(0x1);
        vm.label(user, "user");

        vm.prank(OWNER);
        BASE.transfer(user, _toBase(100));

        vm.startPrank(user);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        IPair.Order memory callOrder = PAIR.orderById(orderId);

        assertNotEq(address(0), callOrder.owner);
        assertNotEq(0, callOrder.price);
        assertNotEq(0, callOrder.amount);

        assertEq(0, BASE.balanceOf(user));
        assertEq(_toBase(100), BASE.balanceOf(address(PAIR)));
    }

    function test_trade_create_limit_buy() external {
        address user = address(0x1);
        vm.label(user, "user");

        vm.prank(OWNER);
        QUOTE.transfer(user, _toQuote(100));

        vm.startPrank(user);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        IPair.Order memory callOrder = PAIR.orderById(orderId);

        assertNotEq(address(0), callOrder.owner);
        assertNotEq(0, callOrder.price);
        assertNotEq(0, callOrder.amount);

        assertEq(0, QUOTE.balanceOf(user));
        assertEq(_toQuote(100), QUOTE.balanceOf(address(PAIR)));
    }

    function test_trade_limit_sell_limit_buy_all_match() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(100));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(100));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);

        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(100), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(100), "BASE user2");
    }

    function test_trade_limit_buy_limit_sell_all_match() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        QUOTE.transfer(user1, _toQuote(100));
        vm.prank(OWNER);
        BASE.transfer(user2, _toBase(100));

        vm.startPrank(user1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), _toBase(100), "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), _toQuote(100), "QUOTE user2");
    }

    function test_trade_limit_sell_less_limit_buy() external {
        // 100개를 판매하고, 50개를 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(100));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(50));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(50), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertNotEq(address(0), order1.owner, "order1, deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        // PAIR 에는 BASE 토큰이 50 개 남아있는지 확인
        // QUOTE 는 없어야 한다.
        assertEq(BASE.balanceOf(address(PAIR)), _toBase(50), "BASE PAIR");
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(50), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(50), "BASE user2");
    }

    function test_trade_limit_buy_less_limit_sell() external {
        // 100개 구매를 요청하고 50개를 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        QUOTE.transfer(user1, _toQuote(100));
        vm.prank(OWNER);
        BASE.transfer(user2, _toBase(50));

        vm.startPrank(user1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(50), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertNotEq(address(0), order1.owner, "order1, deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        // PAIR 50 QUOTE 가 남아 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), _toQuote(50), "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), _toBase(50), "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), _toQuote(50), "QUOTE user2");
    }

    function test_trade_limit_sell_over_limit_buy() external {
        // 100개를 판매하고, 200개를 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(100));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(200));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertNotEq(address(0), order2.owner, "order2, deleted");
        // PAIR 에는 QUOTE 토큰이 100 개 남아있는지 확인
        // BASE 는 없어야 한다.
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        assertEq(QUOTE.balanceOf(address(PAIR)), _toQuote(100), "QUOTE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(100), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2"); // 컨트랙트에 남아있다.
        assertEq(BASE.balanceOf(user2), _toBase(100), "BASE user2");
    }

    function test_trade_limit_buy_over_limit_sell() external {
        // 100개 구매를 요청하고 200개를 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        QUOTE.transfer(user1, _toQuote(100));
        vm.prank(OWNER);
        BASE.transfer(user2, _toBase(200));

        vm.startPrank(user1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertNotEq(address(0), order2.owner, "order2, deleted");
        // PAIR 100 BASE 가 남아 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), _toBase(100), "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), _toBase(100), "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2"); // 100 개는 PAIR 에서 보유
        assertEq(QUOTE.balanceOf(user2), _toQuote(100), "QUOTE user2"); // 100개를 판 만큼의 수익이 들어와 있어야 한다.
    }

    function test_trade_multi_limit_sell_limit_buy_all_match() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, _toBase(100));
        BASE.transfer(user2, _toBase(100));
        QUOTE.transfer(user3, _toQuote(200));
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        assertEq(address(0), order3.owner, "order3, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(100), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), _toQuote(100), "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), _toBase(200), "BASE user3");
    }

    function test_trade_multi_limit_buy_limit_sell_all_match() external {
        // 동일 가격으로 2명의 유저가 구매하고 한 유저가 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, _toQuote(100));
        QUOTE.transfer(user2, _toQuote(100));
        BASE.transfer(user3, _toBase(200));
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        assertEq(address(0), order3.owner, "order3, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), _toBase(100), "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(100), "BASE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), _toQuote(200), "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
    }

    function test_trade_multi_limit_sell_less_limit_buy() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, _toBase(100));
        BASE.transfer(user2, _toBase(100));
        QUOTE.transfer(user3, _toQuote(150));
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(150), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertNotEq(address(0), order2.owner, "order2, deleted");
        assertEq(address(0), order3.owner, "order3, not deleted");
        // PAIR 에는 50 개의 BASE 가 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), _toBase(50), "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(100), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), _toQuote(50), "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), _toBase(150), "BASE user3");
    }

    function test_trade_multi_limit_buy_less_limit_sell() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, _toQuote(100));
        QUOTE.transfer(user2, _toQuote(100));
        BASE.transfer(user3, _toBase(150));
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(150), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertNotEq(address(0), order2.owner, "order2, deleted");
        assertEq(address(0), order3.owner, "order3, not deleted");
        // PAIR 에는 50 개의 QUOTE 가 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), _toQuote(50), "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), _toBase(100), "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(50), "BASE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), _toQuote(150), "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
    }

    function test_trade_multi_limit_sell_over_limit_buy() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, _toBase(100));
        BASE.transfer(user2, _toBase(100));
        QUOTE.transfer(user3, _toQuote(300));
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(300), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        assertNotEq(address(0), order3.owner, "order3, deleted");
        // PAIR 에는 100 개의 QUOTE 가 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), _toQuote(100), "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(100), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), _toQuote(100), "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), _toBase(200), "BASE user3");
    }

    function test_trade_multi_limit_buy_over_limit_sell() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, _toQuote(100));
        QUOTE.transfer(user2, _toQuote(100));
        BASE.transfer(user3, _toBase(300));
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(300), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        assertNotEq(address(0), order3.owner, "order3, deleted");
        // PAIR 에는 100 개의 BASE 가 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), _toBase(100), "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), _toBase(100), "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(100), "BASE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), _toQuote(200), "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
    }

    function test_trade_limit_sell_multi_limit_buy_all_match() external {
        // 동일 가격으로 1명의 유저가 판매하고 2명의 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, _toBase(200));
        QUOTE.transfer(user2, _toQuote(100));
        QUOTE.transfer(user3, _toQuote(100));
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(200), 0, 0);

        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        assertEq(address(0), order3.owner, "order3, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(200), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), _toBase(100), "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user3), _toBase(100), "BASE user3");
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
    }

    function test_trade_limit_buy_multi_limit_sell_all_match() external {
        // 동일 가격으로 1명의 유저가 구매하고 2명의 유저가 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, _toQuote(200));
        BASE.transfer(user2, _toBase(100));
        BASE.transfer(user3, _toBase(100));
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        assertEq(address(0), order3.owner, "order3, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), _toBase(200), "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), _toQuote(100), "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
        assertEq(QUOTE.balanceOf(user3), _toQuote(100), "QUOTE user3");
    }

    function test_trade_limit_sell_less_multi_limit_buy() external {
        // 동일 가격으로 1명의 유저가 판매하고 두명의 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, _toBase(200));
        QUOTE.transfer(user2, _toQuote(100));
        QUOTE.transfer(user3, _toQuote(50));
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(50), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertNotEq(address(0), order1.owner, "order1, deleted");
        assertEq(address(0), order2.owner, "order2, deleted");
        assertEq(address(0), order3.owner, "order3, not deleted");
        // PAIR 에는 50 개의 BASE 가 있어야 한다.
        assertEq(BASE.balanceOf(address(PAIR)), _toBase(50), "BASE PAIR");
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(150), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), _toBase(100), "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user3), _toBase(50), "BASE user3");
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
    }

    function test_trade_limit_buy_less_multi_limit_sell() external {
        // 동일 가격으로 1명의 유저가 구매하고 두명의 유저가 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, _toQuote(200));
        BASE.transfer(user2, _toBase(100));
        BASE.transfer(user3, _toBase(50));
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitBuy(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();
        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(50), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertNotEq(address(0), order1.owner, "order1, deleted");
        assertEq(address(0), order2.owner, "order2, deleted");
        assertEq(address(0), order3.owner, "order3, not deleted");
        // PAIR 에는 50 개의 QUOTE 가 있어야 한다.
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        assertEq(QUOTE.balanceOf(address(PAIR)), _toQuote(50), "QUOTE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), _toBase(150), "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), _toQuote(100), "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
        assertEq(QUOTE.balanceOf(user3), _toQuote(50), "QUOTE user3");
    }

    function test_trade_ticks() external {
        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.limitSell(address(PAIR), 1.5 ether, 200 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 1.4 ether, 200 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 1 ether, 200 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 1.3 ether, 200 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 1.1 ether, 200 ether, 0, 0);
        ROUTER.limitSell(address(PAIR), 1.2 ether, 200 ether, 0, 0);

        vm.roll(block.number + 1);

        ROUTER.limitBuy(address(PAIR), 0.8 ether, 200 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 0.9 ether, 200 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 0.4 ether, 200 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 0.5 ether, 200 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 0.7 ether, 200 ether, 0, 0);
        ROUTER.limitBuy(address(PAIR), 0.6 ether, 200 ether, 0, 0);

        vm.roll(block.number + 1);
        vm.stopPrank();

        (uint256[] memory sellPrices, uint256[] memory buyPrices) = PAIR.ticks();
        {
            uint256 length = sellPrices.length;
            uint256 startPrices = 1 ether;
            for (uint256 i = 0; i < length; i++) {
                assertEq(startPrices + (i * 0.1 ether), sellPrices[i]);
            }
        }
        {
            uint256 length = buyPrices.length;
            uint256 startPrices = 0.9 ether;
            for (uint256 i = 0; i < length; i++) {
                assertEq(startPrices - (i * 0.1 ether), buyPrices[i]);
            }
        }
    }

    uint256 public constant fuzz_limit_length = 1000;

    function testFuzz_trade_limit_order(
        uint8[fuzz_limit_length] memory prices,
        uint8[fuzz_limit_length] memory amounts,
        bool[fuzz_limit_length] memory isSell
    ) external {
        address[fuzz_limit_length] memory users;
        for (uint256 i = 0; i < fuzz_limit_length; i++) {
            users[i] = address(uint160(i + 1));
        }
        (uint256 quoteTickSize, uint256 baseTickSize) = (PAIR.quoteTickSize(), PAIR.baseTickSize());
        uint256 length = fuzz_limit_length / 2;
        uint256 latestOrderId;

        for (uint256 i = 0; i < length; i++) {
            if (i % 50 == 0) vm.roll(block.number + 1);

            address user = users[i];
            uint256 price = prices[i] * quoteTickSize;
            uint256 amount = amounts[i] * baseTickSize;
            if (price == 0 || amount == 0) continue;
            vm.prank(OWNER);
            if (isSell[i]) BASE.transfer(user, amount);
            else QUOTE.transfer(user, Math.mulDiv(price, amount, BASE_DECIMALS));

            vm.startPrank(user);
            BASE.approve(address(ROUTER), type(uint256).max);
            QUOTE.approve(address(ROUTER), type(uint256).max);
            latestOrderId = isSell[i]
                ? ROUTER.limitSell(address(PAIR), price, amount, 0, 0)
                : ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);

            vm.stopPrank();
        }

        uint256 pairBaseBalance = BASE.balanceOf(address(PAIR));
        uint256 pairQuoteBalance = QUOTE.balanceOf(address(PAIR));

        uint256 baseReserve = PAIR.baseReserve();
        uint256 quoteReserve = PAIR.quoteReserve();

        assertEq(baseReserve, pairBaseBalance);
        assertEq(quoteReserve, pairQuoteBalance);

        uint256 checkBaseBalance = 0;
        uint256 checkQuoteBalance = 0;
        for (uint256 i = 1; i <= latestOrderId; i++) {
            IPair.Order memory order = PAIR.orderById(i);
            if (order.side == IPair.OrderSide.SELL) checkBaseBalance += order.amount;
            if (order.side == IPair.OrderSide.BUY) {
                checkQuoteBalance += Math.mulDiv(order.price, order.amount, BASE_DECIMALS);
            }
        }

        assertEq(pairBaseBalance, checkBaseBalance, "not matched Pair BASE Balance and order amount");
        assertEq(pairQuoteBalance, checkQuoteBalance, "not matched Pair QUOTE Balance and order amount");

        uint256 usedBaseAmount = type(uint256).max - BASE.balanceOf(OWNER);
        uint256 usedQuoteAmount = type(uint256).max - QUOTE.balanceOf(OWNER);
        for (uint256 i = 0; i < fuzz_limit_length; i++) {
            checkBaseBalance += BASE.balanceOf(users[i]);
            checkQuoteBalance += QUOTE.balanceOf(users[i]);
        }

        assertEq(usedBaseAmount, checkBaseBalance);
        assertEq(usedQuoteAmount, checkQuoteBalance);
    }

    function test_trade_single_limit_sell_market_buy_all_match() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(100));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(100));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketBuy(address(PAIR), _toQuote(100), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(2);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "market data is must not init");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(100), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(100), "BASE user2");
    }

    function test_trade_multi_limit_sell_market_buy_all_match() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(200));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(300));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(2), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketBuy(address(PAIR), _toQuote(300), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        assertEq(address(0), order3.owner, "market data is must not init");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(300), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(200), "BASE user2");
    }

    function test_trade_single_limit_sell_less_market_buy() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(100));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(50));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketBuy(address(PAIR), _toQuote(50), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(2);
        assertNotEq(address(0), order1.owner, "order1, deleted");
        assertEq(address(0), order2.owner, "market data is must not init");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), _toBase(50), "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(50), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(50), "BASE user2");
    }

    function test_trade_multi_limit_sell_less_market_buy() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(200));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(200));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(2), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketBuy(address(PAIR), _toQuote(200), 0);

        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertNotEq(address(0), order2.owner, "order2, deleted");
        assertEq(address(0), order3.owner, "market data is must not init");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), _toBase(50), "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(200), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(150), "BASE user2");
    }

    function test_trade_single_limit_sell_over_market_buy() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(100));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(200));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketBuy(address(PAIR), _toQuote(200), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(2);
        assertEq(address(0), order1.owner, "order1, deleted");
        assertEq(address(0), order2.owner, "market data is must not init");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(100), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), _toQuote(100), "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(100), "BASE user2");
    }

    function test_trade_multi_limit_sell_over_market_buy() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(200));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(400));

        vm.startPrank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId1 = ROUTER.limitSell(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        uint256 orderId2 = ROUTER.limitSell(address(PAIR), _toQuote(2), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.marketBuy(address(PAIR), _toQuote(400), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(3);
        assertEq(address(0), order1.owner, "order1, not deleted");
        assertEq(address(0), order2.owner, "order2, not deleted");
        assertEq(address(0), order3.owner, "market data is must not init");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), _toQuote(300), "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), _toQuote(100), "QUOTE user2");
        assertEq(BASE.balanceOf(user2), _toBase(200), "BASE user2");
    }

    uint256 private constant fuzz_market_length = 1000;

    function testFuzz_trade_market_order(
        uint8[fuzz_market_length] memory prices,
        uint8[fuzz_market_length] memory amounts,
        uint8[fuzz_market_length] memory orderTypes
    ) external {
        address[fuzz_market_length] memory users;
        for (uint256 i = 0; i < fuzz_market_length; i++) {
            users[i] = address(uint160(i + 1));
        }

        (uint256 quoteTickSize, uint256 baseTickSize) = (PAIR.quoteTickSize(), PAIR.baseTickSize());
        uint256 length = fuzz_market_length / 2;
        uint256 latestOrderId;

        for (uint256 i = 0; i < length; i++) {
            if (i % 50 == 0) vm.roll(block.number + 1);

            address user = users[i];
            uint256 price = prices[i] * quoteTickSize;
            uint256 amount = amounts[i] * baseTickSize;
            if (price == 0 || amount == 0) continue;

            uint256 _type = uint256(orderTypes[i]) % 4;
            if (_type == 0) {
                vm.prank(OWNER);
                BASE.transfer(user, amount);
                vm.prank(user);
                BASE.approve(address(ROUTER), type(uint256).max);
                vm.prank(user);
                latestOrderId = ROUTER.limitSell(address(PAIR), price, amount, 0, 0);
            } else if (_type == 1) {
                vm.prank(OWNER);
                QUOTE.transfer(user, Math.mulDiv(price, amount, BASE_DECIMALS));
                vm.prank(user);
                QUOTE.approve(address(ROUTER), type(uint256).max);
                vm.prank(user);
                latestOrderId = ROUTER.limitBuy(address(PAIR), price, amount, 0, 0);
            } else if (_type == 2) {
                ++latestOrderId;

                vm.prank(OWNER);
                BASE.transfer(user, amount);
                vm.prank(user);
                BASE.approve(address(ROUTER), type(uint256).max);
                vm.prank(user);
                ROUTER.marketSell(address(PAIR), amount, 0);
            } else {
                ++latestOrderId;
                uint256 _value = Math.mulDiv(price, amount, BASE_DECIMALS);
                vm.prank(OWNER);
                QUOTE.transfer(user, _value);
                vm.prank(user);
                QUOTE.approve(address(ROUTER), type(uint256).max);
                vm.prank(user);
                ROUTER.marketBuy(address(PAIR), _value, 0);
            }
        }

        uint256 pairBaseBalance = BASE.balanceOf(address(PAIR));
        uint256 pairQuoteBalance = QUOTE.balanceOf(address(PAIR));

        uint256 baseReserve = PAIR.baseReserve();
        uint256 quoteReserve = PAIR.quoteReserve();

        assertEq(baseReserve, pairBaseBalance);
        assertEq(quoteReserve, pairQuoteBalance);

        uint256 checkBaseBalance = 0;
        uint256 checkQuoteBalance = 0;
        for (uint256 i = 1; i <= latestOrderId; i++) {
            IPair.Order memory order = PAIR.orderById(i);
            if (order.side == IPair.OrderSide.SELL) checkBaseBalance += order.amount;
            if (order.side == IPair.OrderSide.BUY) {
                checkQuoteBalance += Math.mulDiv(order.amount, order.price, BASE_DECIMALS);
            }
        }

        assertEq(pairBaseBalance, checkBaseBalance, "not matched Pair BASE Balance and order amount");
        assertEq(pairQuoteBalance, checkQuoteBalance, "not matched Pair QUOTE Balance and order amount");

        uint256 usedBaseAmount = type(uint256).max - BASE.balanceOf(OWNER);
        uint256 usedQuoteAmount = type(uint256).max - QUOTE.balanceOf(OWNER);
        for (uint256 i = 0; i < fuzz_market_length; i++) {
            checkBaseBalance += BASE.balanceOf(users[i]);
            checkQuoteBalance += QUOTE.balanceOf(users[i]);
        }

        assertEq(usedBaseAmount, checkBaseBalance);
        assertEq(usedQuoteAmount, checkQuoteBalance);
    }
}
