// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Factory, FactoryImpl} from "../src/Factory.sol";
import {Pair, PairImpl} from "../src/Pair.sol";
import {Router, RouterImpl} from "../src/Router.sol";
import {IPair} from "../src/interfaces/IPair.sol";
import {ERC20, IERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/ERC20.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.1.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";

import {Test, console} from "forge-std/Test.sol";

contract T20 is ERC20 {
    uint8 private immutable _decimals;

    constructor(string memory name, string memory symbol, uint8 decimals_) ERC20(name, symbol) {
        _mint(_msgSender(), type(uint256).max);
        _decimals = decimals_;
    }

    function decimals() public view override returns (uint8) {
        return _decimals;
    }
}

contract DEXTradeTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    RouterImpl public ROUTER;
    PairImpl public PAIR;
    IERC20 public QUOTE;
    uint256 public QUOTE_DECIMALS;
    IERC20 public BASE;
    uint256 public BASE_DECIMALS;

    function setUp() external {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");

        vm.startPrank(OWNER);

        QUOTE = IERC20(address(new T20("QUOTE", "Q", 6)));
        BASE = IERC20(address(new T20("BASE", "B", 18)));
        QUOTE_DECIMALS = 10 ** IERC20Metadata(address(QUOTE)).decimals();
        BASE_DECIMALS = 10 ** IERC20Metadata(address(BASE)).decimals();

        address routerImpl = address(new RouterImpl());
        address payable router = payable(address(new Router(routerImpl, type(uint256).max)));
        ROUTER = RouterImpl(router);

        address pairImpl = address(new PairImpl());
        address factoryImpl = address(new FactoryImpl());
        address factory = address(new Factory(factoryImpl, address(ROUTER), FEE_COLLECTOR, address(QUOTE), pairImpl));
        FactoryImpl F = FactoryImpl(factory);

        PAIR = PairImpl(F.createPair(address(BASE), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e4, 0, 0));

        ROUTER.addPair(address(PAIR));

        vm.label(address(ROUTER), "ROUTER");
        vm.label(address(PAIR), "PAIR");
        vm.label(address(QUOTE), "QUOTE");
        vm.label(address(BASE), "BASE");

        vm.stopPrank();
    }

    function _toBase(uint256 x) private view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) private view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    function test_create_limit_sell() external {
        address user = address(0x1);
        vm.label(user, "user");

        vm.prank(OWNER);
        BASE.transfer(user, _toBase(100));

        vm.startPrank(user);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        IPair.Order memory callOrder = PAIR.orderById(orderId);

        assertNotEq(uint8(0), uint8(callOrder._type));
        assertNotEq(address(0), callOrder.owner);
        assertNotEq(0, callOrder.price);
        assertNotEq(0, callOrder.amount);

        assertEq(0, BASE.balanceOf(user));
        assertEq(_toBase(100), BASE.balanceOf(address(PAIR)));
    }

    function test_create_limit_buy() external {
        address user = address(0x1);
        vm.label(user, "user");

        vm.prank(OWNER);
        QUOTE.transfer(user, _toQuote(100));

        vm.startPrank(user);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        IPair.Order memory callOrder = PAIR.orderById(orderId);

        assertNotEq(uint8(0), uint8(callOrder._type));
        assertNotEq(address(0), callOrder.owner);
        assertNotEq(0, callOrder.price);
        assertNotEq(0, callOrder.amount);

        assertEq(0, QUOTE.balanceOf(user));
        assertEq(_toQuote(100), QUOTE.balanceOf(address(PAIR)));
    }

    function test_limit_sell_limit_buy_all_match() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);

        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
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

    function test_limit_buy_limit_sell_all_match() external {
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
        uint256 orderId1 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
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

    function test_limit_sell_less_limit_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(50), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
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

    function test_limit_buy_less_limit_sell() external {
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
        uint256 orderId1 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(50), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
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

    function test_limit_sell_over_limit_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
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

    function test_limit_buy_over_limit_sell() external {
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
        uint256 orderId1 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
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

    function test_multi_limit_sell_limit_buy_all_match() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
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

    function test_multi_limit_buy_limit_sell_all_match() external {
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
        uint256 orderId1 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
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

    function test_multi_limit_sell_less_limit_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(150), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
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

    function test_multi_limit_buy_less_limit_sell() external {
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
        uint256 orderId1 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(150), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
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

    function test_multi_limit_sell_over_limit_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(300), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertNotEq(0, uint8(order3._type), "order3, deleted");
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

    function test_multi_limit_buy_over_limit_sell() external {
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
        uint256 orderId1 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");
        assertEq(BASE.balanceOf(user2), 0, "before BASE user2");
        assertEq(QUOTE.balanceOf(user2), 0, "before QUOTE user2");

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(300), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertNotEq(0, uint8(order3._type), "order3, deleted");
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

    function test_limit_sell_multi_limit_buy_all_match() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(200), 0, 0);

        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
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

    function test_limit_buy_multi_limit_sell_all_match() external {
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
        uint256 orderId1 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
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

    function test_limit_sell_less_multi_limit_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user3);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(50), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
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

    function test_limit_buy_less_multi_limit_sell() external {
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
        uint256 orderId1 = ROUTER.buyLimitOrder(address(PAIR), _toQuote(1), _toBase(200), 0, 0);
        vm.stopPrank();
        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        vm.startPrank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        uint256 orderId3 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(50), 0, 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(orderId3);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "order3, not deleted");
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

    function test_ticks() external {
        vm.startPrank(OWNER);
        BASE.approve(address(ROUTER), type(uint256).max);
        QUOTE.approve(address(ROUTER), type(uint256).max);

        ROUTER.sellLimitOrder(address(PAIR), 1.5 ether, 200 ether, 0, 0);
        ROUTER.sellLimitOrder(address(PAIR), 1.4 ether, 200 ether, 0, 0);
        ROUTER.sellLimitOrder(address(PAIR), 1 ether, 200 ether, 0, 0);
        ROUTER.sellLimitOrder(address(PAIR), 1.3 ether, 200 ether, 0, 0);
        ROUTER.sellLimitOrder(address(PAIR), 1.1 ether, 200 ether, 0, 0);
        ROUTER.sellLimitOrder(address(PAIR), 1.2 ether, 200 ether, 0, 0);

        vm.roll(block.number + 1);

        ROUTER.buyLimitOrder(address(PAIR), 0.8 ether, 200 ether, 0, 0);
        ROUTER.buyLimitOrder(address(PAIR), 0.9 ether, 200 ether, 0, 0);
        ROUTER.buyLimitOrder(address(PAIR), 0.4 ether, 200 ether, 0, 0);
        ROUTER.buyLimitOrder(address(PAIR), 0.5 ether, 200 ether, 0, 0);
        ROUTER.buyLimitOrder(address(PAIR), 0.7 ether, 200 ether, 0, 0);
        ROUTER.buyLimitOrder(address(PAIR), 0.6 ether, 200 ether, 0, 0);

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

    function testFuzz_limit_order(
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
                ? ROUTER.sellLimitOrder(address(PAIR), price, amount, 0, 0)
                : ROUTER.buyLimitOrder(address(PAIR), price, amount, 0, 0);

            vm.stopPrank();
        }

        uint256 pairBaseBalance = BASE.balanceOf(address(PAIR));
        uint256 pairQuoteBalance = QUOTE.balanceOf(address(PAIR));

        uint256 checkBaseBalance = 0;
        uint256 checkQuoteBalance = 0;
        for (uint256 i = 1; i <= latestOrderId; i++) {
            IPair.Order memory order = PAIR.orderById(i);
            if (order._type == IPair.OrderType.SELL) checkBaseBalance += order.amount;
            if (order._type == IPair.OrderType.BUY) {
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

    function test_single_limit_sell_market_buy_all_match() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyMarketOrder(address(PAIR), _toQuote(100), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(2);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "market data is must not inited");
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

    function test_multi_limit_sell_market_buy_all_match() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(2), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyMarketOrder(address(PAIR), _toQuote(300), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "market data is must not inited");
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

    function test_single_limit_sell_less_market_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyMarketOrder(address(PAIR), _toQuote(50), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(2);
        assertNotEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "market data is must not inited");
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

    function test_multi_limit_sell_less_market_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(2), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyMarketOrder(address(PAIR), _toQuote(200), 0);

        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertNotEq(0, uint8(order2._type), "order2, deleted");
        assertEq(0, uint8(order3._type), "market data is must not inited");
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

    function test_single_limit_sell_over_market_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyMarketOrder(address(PAIR), _toQuote(200), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(2);
        assertEq(0, uint8(order1._type), "order1, deleted");
        assertEq(0, uint8(order2._type), "market data is must not inited");
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

    function test_multi_limit_sell_over_market_buy() external {
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
        uint256 orderId1 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(1), _toBase(100), 0, 0);
        uint256 orderId2 = ROUTER.sellLimitOrder(address(PAIR), _toQuote(2), _toBase(100), 0, 0);
        vm.stopPrank();

        assertEq(BASE.balanceOf(user1), 0, "before BASE user1");
        assertEq(QUOTE.balanceOf(user1), 0, "before QUOTE user1");

        vm.startPrank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        ROUTER.buyMarketOrder(address(PAIR), _toQuote(400), 0);
        vm.stopPrank();

        // order 데이터가 삭제 되었는지 확인
        IPair.Order memory order1 = PAIR.orderById(orderId1);
        IPair.Order memory order2 = PAIR.orderById(orderId2);
        IPair.Order memory order3 = PAIR.orderById(3);
        assertEq(0, uint8(order1._type), "order1, not deleted");
        assertEq(0, uint8(order2._type), "order2, not deleted");
        assertEq(0, uint8(order3._type), "market data is must not inited");
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

    function testFuzz_market_order(
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
                latestOrderId = ROUTER.sellLimitOrder(address(PAIR), price, amount, 0, 0);
            } else if (_type == 1) {
                vm.prank(OWNER);
                QUOTE.transfer(user, Math.mulDiv(price, amount, BASE_DECIMALS));
                vm.prank(user);
                QUOTE.approve(address(ROUTER), type(uint256).max);
                vm.prank(user);
                latestOrderId = ROUTER.buyLimitOrder(address(PAIR), price, amount, 0, 0);
            } else if (_type == 2) {
                ++latestOrderId;

                vm.prank(OWNER);
                BASE.transfer(user, amount);
                vm.prank(user);
                BASE.approve(address(ROUTER), type(uint256).max);
                vm.prank(user);
                ROUTER.sellMarketOrder(address(PAIR), amount, 0);
            } else {
                ++latestOrderId;
                uint256 _value = Math.mulDiv(price, amount, BASE_DECIMALS);
                vm.prank(OWNER);
                QUOTE.transfer(user, _value);
                vm.prank(user);
                QUOTE.approve(address(ROUTER), type(uint256).max);
                vm.prank(user);
                ROUTER.buyMarketOrder(address(PAIR), _value, 0);
            }
        }

        uint256 pairBaseBalance = BASE.balanceOf(address(PAIR));
        uint256 pairQuoteBalance = QUOTE.balanceOf(address(PAIR));

        uint256 checkBaseBalance = 0;
        uint256 checkQuoteBalance = 0;
        for (uint256 i = 1; i <= latestOrderId; i++) {
            IPair.Order memory order = PAIR.orderById(i);
            if (order._type == IPair.OrderType.SELL) checkBaseBalance += order.amount;
            if (order._type == IPair.OrderType.BUY) {
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
