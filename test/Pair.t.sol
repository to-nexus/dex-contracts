// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {ERC20, IERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/ERC20.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";
import {Pair, PairImpl} from "../src/Pair.sol";

contract T20 is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(_msgSender(), type(uint256).max);
    }
}

contract PairTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    PairImpl public PAIR;
    IERC20 public QUOTE;
    IERC20 public BASE;

    function setUp() external {
        vm.label(OWNER, "owner");

        vm.startPrank(OWNER);
        PAIR = PairImpl(address(new Pair(address(new PairImpl()))));
        QUOTE = IERC20(address(new T20("QUOTE", "Q")));
        BASE = IERC20(address(new T20("BASE", "B")));

        vm.label(address(PAIR), "PAIR");
        vm.label(address(QUOTE), "QUOTE");
        vm.label(address(BASE), "BASE");

        PAIR.initialize(address(QUOTE), address(BASE), 1e18, 0.001e18, 0.00001e18, 0, 0);
        vm.stopPrank();
    }

    function test_create_limit_sell() external {
        address user = address(0x1);
        vm.label(user, "user");

        vm.prank(OWNER);
        BASE.transfer(user, 100 ether);

        vm.startPrank(user);
        PairImpl.Order memory order =
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user, price: 1 ether, amount: 100 ether});
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId = PAIR.limit(order);
        vm.stopPrank();

        PairImpl.Order memory callOrder = PAIR.orderById(orderId);

        assertEq(uint8(order._type), uint8(callOrder._type));
        assertEq(order.owner, callOrder.owner);
        assertEq(order.price, callOrder.price);
        assertEq(order.amount, callOrder.amount);

        assertEq(0, BASE.balanceOf(user));
        assertEq(100 ether, BASE.balanceOf(address(PAIR)));
    }

    function test_create_limit_buy() external {
        address user = address(0x1);
        vm.label(user, "user");

        vm.prank(OWNER);
        QUOTE.transfer(user, 100 ether);

        vm.startPrank(user);
        PairImpl.Order memory order =
            PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user, price: 1 ether, amount: 100 ether});
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId = PAIR.limit(order);
        vm.stopPrank();

        PairImpl.Order memory callOrder = PAIR.orderById(orderId);

        assertEq(uint8(order._type), uint8(callOrder._type));
        assertEq(order.owner, callOrder.owner);
        assertEq(order.price, callOrder.price);
        assertEq(order.amount, callOrder.amount);

        assertEq(0, QUOTE.balanceOf(user));
        assertEq(100 ether, QUOTE.balanceOf(address(PAIR)));
    }

    function test_limit_sell_limit_buy_all_match() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 100 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 100 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 100 ether, "BASE user2");
    }

    function test_limit_buy_limit_sell_all_match() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        QUOTE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        BASE.transfer(user2, 100 ether);

        vm.startPrank(user1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user1, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user2, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), 100 ether, "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 100 ether, "QUOTE user2");
    }

    function test_limit_sell_less_limit_buy() external {
        // 100개를 판매하고, 50개를 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 50 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 50 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        // PAIR 에는 BASE 토큰이 50 개 남아있는지 확인
        // QUOTE 는 없어야 한다.
        assertEq(BASE.balanceOf(address(PAIR)), 50 ether, "BASE PAIR");
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 50 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 50 ether, "BASE user2");
    }

    function test_limit_buy_less_limit_sell() external {
        // 100개 구매를 요청하고 50개를 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        QUOTE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        BASE.transfer(user2, 50 ether);

        vm.startPrank(user1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user1, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user2, price: 1 ether, amount: 50 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        // PAIR 50 QUOTE 가 남아 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 50 ether, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), 50 ether, "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 50 ether, "QUOTE user2");
    }

    function test_limit_sell_over_limit_buy() external {
        // 100개를 판매하고, 200개를 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 200 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 200 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
        // PAIR 에는 QUOTE 토큰이 100 개 남아있는지 확인
        // BASE 는 없어야 한다.
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        assertEq(QUOTE.balanceOf(address(PAIR)), 100 ether, "QUOTE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 100 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2"); // 컨트랙트에 남아있다.
        assertEq(BASE.balanceOf(user2), 100 ether, "BASE user2");
    }

    function test_limit_buy_over_limit_sell() external {
        // 100개 구매를 요청하고 200개를 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        QUOTE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        BASE.transfer(user2, 200 ether);

        vm.startPrank(user1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user1, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user2, price: 1 ether, amount: 200 ether})
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
        // PAIR 100 BASE 가 남아 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 100 ether, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), 100 ether, "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2"); // 100 개는 PAIR 에서 보유
        assertEq(QUOTE.balanceOf(user2), 100 ether, "QUOTE user2"); // 100개를 판 만큼의 수익이 들어와 있어야 한다.
    }

    function test_multi_limit_sell_limit_buy_all_match() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, 100 ether);
        BASE.transfer(user2, 100 ether);
        QUOTE.transfer(user3, 200 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user2, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user3, price: 1 ether, amount: 200 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 100 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 100 ether, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 200 ether, "BASE user3");
    }

    function test_multi_limit_buy_limit_sell_all_match() external {
        // 동일 가격으로 2명의 유저가 구매하고 한 유저가 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, 100 ether);
        QUOTE.transfer(user2, 100 ether);
        BASE.transfer(user3, 200 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user1, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user3, price: 1 ether, amount: 200 ether})
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), 100 ether, "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 100 ether, "BASE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 200 ether, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
    }

    function test_multi_limit_sell_less_limit_buy() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, 100 ether);
        BASE.transfer(user2, 100 ether);
        QUOTE.transfer(user3, 150 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user2, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user3, price: 1 ether, amount: 150 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
        // PAIR 에는 50 개의 BASE 가 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 50 ether, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 100 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 50 ether, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 150 ether, "BASE user3");
    }

    function test_multi_limit_buy_less_limit_sell() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, 100 ether);
        QUOTE.transfer(user2, 100 ether);
        BASE.transfer(user3, 150 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user1, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user3, price: 1 ether, amount: 150 ether})
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
        // PAIR 에는 50 개의 QUOTE 가 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 50 ether, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), 100 ether, "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 50 ether, "BASE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 150 ether, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
    }

    function test_multi_limit_sell_over_limit_buy() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, 100 ether);
        BASE.transfer(user2, 100 ether);
        QUOTE.transfer(user3, 300 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user2, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user3, price: 1 ether, amount: 300 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertNotEq(0, uint8(order3._type), "order3, deleted");
        // PAIR 에는 100 개의 QUOTE 가 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 100 ether, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 100 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 100 ether, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 200 ether, "BASE user3");
    }

    function test_multi_limit_buy_over_limit_sell() external {
        // 동일 가격으로 2명의 유저가 판매하고 한 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, 100 ether);
        QUOTE.transfer(user2, 100 ether);
        BASE.transfer(user3, 300 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user1, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user3, price: 1 ether, amount: 300 ether})
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertNotEq(0, uint8(order3._type), "order3, deleted");
        // PAIR 에는 100 개의 BASE 가 있어야 한다.
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 100 ether, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        assertEq(BASE.balanceOf(user1), 100 ether, "BASE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 100 ether, "BASE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user3), 200 ether, "QUOTE user3");
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
    }

    function test_limit_sell_multi_limit_buy_all_match() external {
        // 동일 가격으로 1명의 유저가 판매하고 2명의 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, 200 ether);
        QUOTE.transfer(user2, 100 ether);
        QUOTE.transfer(user3, 100 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 200 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        vm.startPrank(user3);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user3, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 200 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 100 ether, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user3), 100 ether, "BASE user3");
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
    }

    function test_limit_buy_multi_limit_sell_all_match() external {
        // 동일 가격으로 1명의 유저가 구매하고 2명의 유저가 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, 200 ether);
        BASE.transfer(user2, 100 ether);
        BASE.transfer(user3, 100 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user1, price: 1 ether, amount: 200 ether}));
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user2, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        vm.startPrank(user3);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user3, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 200 ether, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 100 ether, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
        assertEq(QUOTE.balanceOf(user3), 100 ether, "QUOTE user3");
    }

    function test_limit_sell_less_multi_limit_buy() external {
        // 동일 가격으로 1명의 유저가 판매하고 두명의 유저가 구매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        BASE.transfer(user1, 200 ether);
        QUOTE.transfer(user2, 100 ether);
        QUOTE.transfer(user3, 50 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 200 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}));
        vm.stopPrank();

        vm.startPrank(user3);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user3, price: 1 ether, amount: 50 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
        // PAIR 에는 50 개의 BASE 가 있어야 한다.
        assertEq(BASE.balanceOf(address(PAIR)), 50 ether, "BASE PAIR");
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 150 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 100 ether, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user3), 50 ether, "BASE user3");
        assertEq(QUOTE.balanceOf(user3), 0, "QUOTE user3");
    }

    function test_limit_buy_less_multi_limit_sell() external {
        // 동일 가격으로 1명의 유저가 구매하고 두명의 유저가 판매한다.
        address user1 = address(0x1);
        address user2 = address(0x2);
        address user3 = address(0x3);
        vm.label(user1, "user1");
        vm.label(user2, "user2");
        vm.label(user3, "user3");

        vm.startPrank(OWNER);
        QUOTE.transfer(user1, 200 ether);
        BASE.transfer(user2, 100 ether);
        BASE.transfer(user3, 50 ether);
        vm.stopPrank();

        vm.startPrank(user1);
        QUOTE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user1, price: 1 ether, amount: 200 ether}));
        vm.stopPrank();
        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user2, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        vm.startPrank(user3);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId3 =
            PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user3, price: 1 ether, amount: 50 ether}));
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(orderId3);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
        // PAIR 에는 50 개의 QUOTE 가 있어야 한다.
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        assertEq(QUOTE.balanceOf(address(PAIR)), 50 ether, "QUOTE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 150 ether, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user2), 0, "BASE user2");
        assertEq(QUOTE.balanceOf(user2), 100 ether, "QUOTE user2");
        // [user3] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user3), 0, "BASE user3");
        assertEq(QUOTE.balanceOf(user3), 50 ether, "QUOTE user3");
    }

    function test_ticks() external {
        vm.startPrank(OWNER);
        BASE.approve(address(PAIR), type(uint256).max);
        QUOTE.approve(address(PAIR), type(uint256).max);

        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: OWNER, price: 1.5 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: OWNER, price: 1.4 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: OWNER, price: 1 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: OWNER, price: 1.3 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: OWNER, price: 1.1 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: OWNER, price: 1.2 ether, amount: 200 ether}));

        vm.roll(block.number + 1);

        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: OWNER, price: 0.8 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: OWNER, price: 0.9 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: OWNER, price: 0.4 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: OWNER, price: 0.5 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: OWNER, price: 0.7 ether, amount: 200 ether}));
        PAIR.limit(PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: OWNER, price: 0.6 ether, amount: 200 ether}));

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

    uint256 public constant fuzz_length = 1000;

    function testFuzz_limit_order(
        uint8[fuzz_length] memory prices,
        uint8[fuzz_length] memory amounts,
        bool[fuzz_length] memory isSell
    ) external {
        address[fuzz_length] memory users;
        for (uint256 i = 0; i < fuzz_length; i++) {
            users[i] = address(uint160(i + 1));
        }

        vm.startPrank(OWNER);
        BASE.approve(address(PAIR), type(uint256).max);
        QUOTE.approve(address(PAIR), type(uint256).max);

        (uint256 quoteTickSize, uint256 baseTickSize) = (PAIR.quoteTickSize(), PAIR.baseTickSize());
        uint256 denominator = 1e18;
        uint256 length = fuzz_length / 2;
        uint256 latestOrderId;

        for (uint256 i = 0; i < length; i++) {
            if (i % 50 == 0) vm.roll(block.number + 1);

            address user = users[i];
            uint256 price = prices[i] * quoteTickSize;
            uint256 amount = amounts[i] * baseTickSize;
            if (price == 0 || amount == 0) continue;
            PairImpl.OrderType _type = isSell[i] ? PairImpl.OrderType.SELL : PairImpl.OrderType.BUY;

            latestOrderId = PAIR.limit(PairImpl.Order({_type: _type, owner: user, price: price, amount: amount}));
        }

        uint256 pairBaseBalance = BASE.balanceOf(address(PAIR));
        uint256 pairQuoteBalance = QUOTE.balanceOf(address(PAIR));

        uint256 checkBaseBalance = 0;
        uint256 checkQuoteBalance = 0;
        for (uint256 i = 1; i <= latestOrderId; i++) {
            PairImpl.Order memory order = PAIR.orderById(i);
            if (order._type == PairImpl.OrderType.SELL) checkBaseBalance += order.amount;
            if (order._type == PairImpl.OrderType.BUY) {
                checkQuoteBalance += Math.mulDiv(order.amount, order.price, denominator);
            }
        }

        assertEq(pairBaseBalance, checkBaseBalance, "not matched Pair BASE Balance and order amount");
        assertEq(pairQuoteBalance, checkQuoteBalance, "not matched Pair QUOTE Balance and order amount");

        uint256 usedBaseAmount = type(uint256).max - BASE.balanceOf(OWNER);
        uint256 usedQuoteAmount = type(uint256).max - QUOTE.balanceOf(OWNER);
        for (uint256 i = 0; i < fuzz_length; i++) {
            checkBaseBalance += BASE.balanceOf(users[i]);
            checkQuoteBalance += QUOTE.balanceOf(users[i]);
        }

        assertEq(usedBaseAmount, checkBaseBalance);
        assertEq(usedQuoteAmount, checkQuoteBalance);
    }

    function test_single_limit_sell_market_buy_all_match() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 100 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        PAIR.market(
            PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}), 100 ether
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "market data is must not inited");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 100 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 100 ether, "BASE user2");
    }

    function test_multi_limit_sell_market_buy_all_match() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 200 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 300 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 2 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        PAIR.market(
            PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}), 300 ether
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "market data is must not inited");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 300 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 200 ether, "BASE user2");
    }

    function test_single_limit_sell_less_market_buy() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 50 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        PAIR.market(
            PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}), 50 ether
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(2);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "market data is must not inited");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 50 ether, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 50 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 50 ether, "BASE user2");
    }

    function test_multi_limit_sell_less_market_buy() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 200 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 200 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 2 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        PAIR.market(
            PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}), 200 ether
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "market data is must not inited");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 50 ether, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 200 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 0, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 150 ether, "BASE user2");
    }

    function test_single_limit_sell_over_market_buy() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 100 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 200 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        PAIR.market(
            PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}), 200 ether
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(2);
        assertEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "market data is must not inited");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 100 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 100 ether, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 100 ether, "BASE user2");
    }

    function test_multi_limit_sell_over_market_buy() external {
        address user1 = address(0x1);
        address user2 = address(0x2);
        vm.label(user1, "user1");
        vm.label(user2, "user2");

        vm.prank(OWNER);
        BASE.transfer(user1, 200 ether);
        vm.prank(OWNER);
        QUOTE.transfer(user2, 400 ether);

        vm.startPrank(user1);
        BASE.approve(address(PAIR), type(uint256).max);
        uint256 orderId1 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 1 ether, amount: 100 ether})
        );
        uint256 orderId2 = PAIR.limit(
            PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user1, price: 2 ether, amount: 100 ether})
        );
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(PAIR), type(uint256).max);
        PAIR.market(
            PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user2, price: 1 ether, amount: 100 ether}), 400 ether
        );
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        PairImpl.Order memory order1 = PAIR.orderById(orderId1);
        PairImpl.Order memory order2 = PAIR.orderById(orderId2);
        PairImpl.Order memory order3 = PAIR.orderById(3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "market data is must not inited");
        // PAIR 에는 잔액이 없는지 확인
        assertEq(QUOTE.balanceOf(address(PAIR)), 0, "QUOTE PAIR");
        assertEq(BASE.balanceOf(address(PAIR)), 0, "BASE PAIR");
        // [user1] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(BASE.balanceOf(user1), 0, "BASE user1");
        assertEq(QUOTE.balanceOf(user1), 300 ether, "QUOTE user1");
        // [user2] 거래 잔액이 잘 전송 되었는지 확인
        assertEq(QUOTE.balanceOf(user2), 100 ether, "QUOTE user2");
        assertEq(BASE.balanceOf(user2), 200 ether, "BASE user2");
    }

    uint256 private constant fuzz_market_length = 1000;

    function testFuzz_market_order(
        uint8[fuzz_market_length] memory prices,
        uint8[fuzz_market_length] memory amounts,
        uint8[fuzz_market_length] memory orderTypes
    ) external {
        address[fuzz_market_length] memory users;
        for (uint256 i = 0; i < fuzz_market_length; i++) {
            users[i] = address(uint160(i + 1));
        }

        vm.startPrank(OWNER);
        BASE.approve(address(PAIR), type(uint256).max);
        QUOTE.approve(address(PAIR), type(uint256).max);

        (uint256 quoteTickSize, uint256 baseTickSize) = (PAIR.quoteTickSize(), PAIR.baseTickSize());
        uint256 denominator = 1e18;
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
                latestOrderId = PAIR.limit(
                    PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user, price: price, amount: amount})
                );
            } else if (_type == 1) {
                latestOrderId = PAIR.limit(
                    PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user, price: price, amount: amount})
                );
            } else if (_type == 2) {
                ++latestOrderId;
                PAIR.market(PairImpl.Order({_type: PairImpl.OrderType.SELL, owner: user, price: 0, amount: 0}), amount);
            } else {
                ++latestOrderId;
                PAIR.market(
                    PairImpl.Order({_type: PairImpl.OrderType.BUY, owner: user, price: 0, amount: 0}),
                    Math.mulDiv(price, amount, denominator)
                );
            }
        }

        uint256 pairBaseBalance = BASE.balanceOf(address(PAIR));
        uint256 pairQuoteBalance = QUOTE.balanceOf(address(PAIR));

        uint256 checkBaseBalance = 0;
        uint256 checkQuoteBalance = 0;
        for (uint256 i = 1; i <= latestOrderId; i++) {
            PairImpl.Order memory order = PAIR.orderById(i);
            if (order._type == PairImpl.OrderType.SELL) checkBaseBalance += order.amount;
            if (order._type == PairImpl.OrderType.BUY) {
                checkQuoteBalance += Math.mulDiv(order.amount, order.price, denominator);
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
