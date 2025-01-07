// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.1.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.1.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.1.0/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.1.0/utils/PausableUpgradeable.sol";

import {IPair} from "./interfaces/IPair.sol";
import {ASCList} from "./lib/ASCList.sol";
import {DESCList} from "./lib/DESCList.sol";
import {DoubleLinkedList} from "./lib/DoubleLinkedList.sol";

contract Pair is ERC1967Proxy {
    constructor(
        address implementation,
        address owner,
        address router,
        address quote,
        address base,
        uint256 denominator,
        uint256 quoteTickSize,
        uint256 baseTickSize,
        uint256 quoteFeePermile,
        uint256 baseFeePermile
    )
        ERC1967Proxy(
            implementation,
            abi.encodeWithSelector(
                PairImpl.initialize.selector,
                owner,
                router,
                quote,
                base,
                denominator,
                quoteTickSize,
                baseTickSize,
                quoteFeePermile,
                baseFeePermile
            )
        )
    {}
}

contract PairImpl is IPair, UUPSUpgradeable, OwnableUpgradeable, PausableUpgradeable {
    using SafeERC20 for IERC20;
    using Math for uint256;
    using DoubleLinkedList for DoubleLinkedList.U256;
    using DESCList for DESCList.U256;
    using ASCList for ASCList.U256;

    error PairInvalidRouter(address);
    error PairInvalidOrderType(OrderType);
    error PairInvalidPrice(uint256);
    error PairInvalidAmount(uint256);
    error PairInvalidOrderId(uint256);
    error PairUnknownPrices(OrderType, uint256);

    event OrderCreated(
        address indexed owner,
        uint256 indexed orderId,
        OrderType indexed _type,
        uint256 price,
        uint256 amount,
        uint256 timestamp
    );

    event OrderMatched(
        uint256 indexed sellId, uint256 indexed buyId, uint256 indexed price, uint256 amount, uint256 timestamp
    );

    event OrderClosed(uint256 indexed orderId, CloseType indexed _type, uint256 timestamp);

    address public ROUTER;
    IERC20 public override QUOTE; // immutable
    IERC20 public override BASE; // immutable
    uint256 public override DENOMINATOR; // immutable

    uint256 public quoteTickSize; // quote 거래 단위
    uint256 public baseTickSize; // base 거래 단위

    uint256 public quoteFeePermile; // quote 수수료
    uint256 public baseFeePermile; // base 수수료

    uint256 private _orderIdCounter; // order id 생성기

    ASCList.U256 private _sellPrices; // sell order 의 가격 목록
    DESCList.U256 private _buyPrices; // buy order 의 가격 목록
    mapping(uint256 price => DoubleLinkedList.U256) private _sellOrders; // price => 판매 order id list
    mapping(uint256 price => DoubleLinkedList.U256) private _buyOrders; //  price => 구매 order id list
    mapping(uint256 orderId => Order) private _allOrders; // 모든 주문 정보

    modifier onlyRouter() {
        if (_msgSender() != ROUTER) revert PairInvalidRouter(_msgSender());
        _;
    }

    function initialize(
        address owner,
        address router,
        address quote,
        address base,
        uint256 denominator,
        uint256 _quoteTickSize,
        uint256 _baseTickSize,
        uint256 _quoteFeePermile,
        uint256 _baseFeePermile
    ) external initializer {
        ROUTER = router;
        QUOTE = IERC20(quote);
        BASE = IERC20(base);
        DENOMINATOR = denominator;

        quoteTickSize = _quoteTickSize;
        baseTickSize = _baseTickSize;

        quoteFeePermile = _quoteFeePermile;
        baseFeePermile = _baseFeePermile;

        __Ownable_init(owner);
        __Pausable_init();
    }

    function orderById(uint256 id) external view returns (Order memory) {
        return _allOrders[id];
    }

    function ticks() external view returns (uint256[] memory sellPrices, uint256[] memory buyPrices) {
        sellPrices = _sellPrices.values();
        buyPrices = _buyPrices.values();
    }

    function limit(Order memory order) external override whenNotPaused onlyRouter returns (uint256 orderId) {
        if (order._type == OrderType.NONE) revert PairInvalidOrderType(OrderType.NONE);

        // 입력된 수량의 조건을 확인한다.
        if (order.price == 0 || order.price % quoteTickSize != 0) revert PairInvalidPrice(order.price);
        if (order.amount == 0 || order.amount % baseTickSize != 0) revert PairInvalidAmount(order.amount);

        orderId = ++_orderIdCounter;
        if (order._type == OrderType.SELL) {
            // 거래 주문 생성
            order = _createSellOrder(orderId, order);
            // 남은 잔액에 따른 분기 처리
            if (order.amount == 0) {
                emit OrderClosed(orderId, CloseType.ALL_MATCH, block.timestamp);
            } else {
                _allOrders[orderId] = order;
                _sellPrices.push(order.price);
                _sellOrders[order.price].push(orderId);
            }
        } else {
            uint256 matchedBaseAmount;
            uint256 useQuoteAmount;
            (order, matchedBaseAmount, useQuoteAmount) = _createBuyOrder(orderId, order, 0);

            // 주문과 동시에 채결되었다면, 설정된 가격보다 싼 가격으로 샀을수 있기때문에 해당 잔액을 돌려준다.
            if (matchedBaseAmount != 0) {
                uint256 measuredQuoteAmount = Math.mulDiv(order.price, matchedBaseAmount, DENOMINATOR);
                uint256 diffQuoteAmount = measuredQuoteAmount - useQuoteAmount;
                if (diffQuoteAmount != 0) QUOTE.safeTransfer(order.owner, diffQuoteAmount);
            }

            if (order.amount == 0) {
                emit OrderClosed(orderId, CloseType.ALL_MATCH, block.timestamp);
            } else {
                _allOrders[orderId] = order;
                _buyPrices.push(order.price);
                _buyOrders[order.price].push(orderId);
            }
        }
    }

    function market(Order memory order, uint256 spendAmount) external override whenNotPaused onlyRouter {
        if (order._type == OrderType.NONE) revert PairInvalidOrderType(OrderType.NONE);

        uint256 orderId = ++_orderIdCounter;
        if (order._type == OrderType.SELL) {
            if (spendAmount == 0 || spendAmount % baseTickSize != 0) revert PairInvalidAmount(spendAmount);
            order.price = 0; // 판매는 spendAmount 을 다 팔때까지 가격을 내린다.
            order.amount = spendAmount;

            order = _createSellOrder(orderId, order);
            if (order.amount != 0) BASE.transfer(order.owner, order.amount);
        } else {
            // if (spendAmount < 0 || spendAmount % quoteTickSize != 0) revert(); todo
            order.price = type(uint256).max; // 구매는 spendAmount 을 모두 소진할때까지 가격을 올린다.

            uint256 useQuoteAmount;
            (order,, useQuoteAmount) = _createBuyOrder(orderId, order, spendAmount);
            uint256 remainQuoteAmount = spendAmount - useQuoteAmount;
            if (remainQuoteAmount != 0) QUOTE.transfer(order.owner, remainQuoteAmount);
        }
        emit OrderClosed(orderId, CloseType.MARKET, block.timestamp);
    }

    function close(uint256 orderId) external {
        Order memory order = _allOrders[orderId];
        if (order._type == OrderType.NONE) revert PairInvalidOrderId(orderId);

        DoubleLinkedList.U256 storage _orders;
        bool isSellOrder = order._type == OrderType.SELL;

        // 컨트랙트에 보유하고 있던 토큰을 돌려준다.
        if (isSellOrder) {
            _orders = _sellOrders[order.price];
            BASE.safeTransfer(order.owner, order.amount);
        } else {
            _orders = _buyOrders[order.price];
            QUOTE.safeTransfer(order.owner, Math.mulDiv(order.price, order.amount, DENOMINATOR));
        }

        // 해당 데이터를 제거한다.
        _removeOrder(orderId, CloseType.CANCEL, _orders);
        // 만약 해당 가격의 마지막 정보였다면 price 가격을 제거한다.
        if (_orders.empty()) {
            if (isSellOrder) _sellPrices.remove(order.price);
            else _buyPrices.remove(order.price);
        }
    }

    function _createSellOrder(uint256 orderId, Order memory order) private returns (Order memory) {
        if (order._type != OrderType.SELL) revert PairInvalidOrderType(order._type);

        // 1. 주문에 필요한 토큰을 가져온다.
        BASE.safeTransferFrom(_msgSender(), address(this), order.amount);
        emit OrderCreated(order.owner, orderId, order._type, order.price, order.amount, block.timestamp);

        // 2. 즉시 거래가 가능한 매물이 있다면 거래를 시킨다.
        // SELL 주문일 경우는 BUY 목록에서 비싼것 부터 찾으며, 입력된 price 보다 높거나 같은 구매 거래만 성사 시킨다.
        uint256 earnQuoteAmount;
        (order, earnQuoteAmount) = _tradeSellOrder(orderId, order);

        // 3. 즉시 거래된 수익은 판매자에게 전송한다. todo 수수료 처리
        if (earnQuoteAmount != 0) QUOTE.safeTransfer(order.owner, earnQuoteAmount);

        return order;
    }

    function _createBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 spendQuoteAmount // Market Order 인 경우 이 값을 설정한다.
    ) private returns (Order memory, uint256, uint256) {
        if (order._type != OrderType.BUY) revert PairInvalidOrderType(order._type);

        // 입력된 수량의 조건을 확인한다.
        uint256 quoteAmount;
        if (spendQuoteAmount != 0) quoteAmount = spendQuoteAmount;
        else quoteAmount = Math.mulDiv(order.price, order.amount, DENOMINATOR);

        // 1. 주문에 필요한 토큰을 가져온다.
        QUOTE.safeTransferFrom(_msgSender(), address(this), quoteAmount);
        emit OrderCreated(order.owner, orderId, order._type, order.price, order.amount, block.timestamp);

        // 2. 즉시 거래가 가능한 매물이 있다면 거래를 시킨다.
        // BUY 주문일 경우에는 SELL 목록에서 싼것부터 찾으며, 입력된 price 보다 낮거나 같은 판매 거래만 성사 시킨다.
        (Order memory tradedOrder, uint256 matchedBaseAmount, uint256 useQuoteAmount) =
            _tradeBuyOrder(orderId, order, spendQuoteAmount);

        // 3. 즉시 채결된 BASE 토큰을 전송한다.
        if (matchedBaseAmount != 0) BASE.safeTransfer(order.owner, matchedBaseAmount);

        return (tradedOrder, matchedBaseAmount, useQuoteAmount);
    }

    // SELL 주문일 경우는 BUY 목록에서 비싼것 부터 찾으며, order.price 보다 높거나 같은 구매 거래만 성사 시킨다.

    function _tradeSellOrder(uint256 orderId, Order memory order)
        private
        returns (Order memory cOrder, uint256 earnQuoteAmount)
    {
        // (earnQuoteAmount) 판매로 인해 벌어들인 Quote 수량
        cOrder = _copyOrder(order);

        while (!_buyPrices.empty()) {
            uint256 price = _buyPrices.at(0);
            if (price < cOrder.price) break;
            DoubleLinkedList.U256 storage _orders = _buyOrders[price];

            uint256 length = DoubleLinkedList.length(_orders);
            while (length != 0) {
                unchecked {
                    --length;
                }

                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // 채결 수량 업데이트
                (address targetOwner, uint256 tradeAmount) =
                    _matchOrderAmount(orderId, cOrder, targetId, target, price, _orders);

                // 거래 성사 todo target 수수료 처리
                BASE.safeTransfer(targetOwner, tradeAmount);
                earnQuoteAmount += Math.mulDiv(price, tradeAmount, DENOMINATOR);

                if (cOrder.amount == 0) return (cOrder, earnQuoteAmount);
            }
            if (!_buyPrices.remove(price)) revert PairUnknownPrices(OrderType.BUY, price);
        }
        return (cOrder, earnQuoteAmount);
    }

    // BUY 주문일 경우에는 SELL 목록에서 싼것부터 찾으며, 입력된 price 보다 낮거나 같은 판매 거래만 성사 시킨다.
    function _tradeBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 quoteAmount // Market 거래시 사용 할 Quote 수량
    ) private returns (Order memory cOrder, uint256 matchedBaseAmount, uint256 useQuoteAmount) {
        cOrder = _copyOrder(order);

        while (!_sellPrices.empty()) {
            uint256 price = _sellPrices.at(0);
            if (price > order.price) break;
            DoubleLinkedList.U256 storage _orders = _sellOrders[price];

            // Market 거래일 경우 해당 가격으로 최대 구매할 수 있는 수량을 계산한다.
            if (quoteAmount != 0) {
                uint256 tradableAmount = Math.mulDiv(quoteAmount - useQuoteAmount, DENOMINATOR, price);
                tradableAmount -= tradableAmount % baseTickSize;
                cOrder.amount = tradableAmount;
                if (cOrder.amount == 0) return (cOrder, matchedBaseAmount, useQuoteAmount);
            }

            // 해당 가격의 모든 SellOrder 를 순회한다.
            uint256 length = DoubleLinkedList.length(_orders);
            while (length != 0) {
                unchecked {
                    --length;
                }

                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // 거래 수량 업데이트
                (address targetOwner, uint256 tradeAmount) =
                    _matchOrderAmount(orderId, cOrder, targetId, target, price, _orders);
                // 거래 가치 계산
                uint256 tradeQuoteAmount = Math.mulDiv(price, tradeAmount, DENOMINATOR);
                // 거래 성사 todo target 수수료 처리
                QUOTE.safeTransfer(targetOwner, tradeQuoteAmount);

                // 정보 업데이트
                matchedBaseAmount += tradeAmount;
                useQuoteAmount += tradeQuoteAmount;
                if (cOrder.amount == 0) return (cOrder, matchedBaseAmount, useQuoteAmount);
            }
            if (!_sellPrices.remove(price)) revert PairUnknownPrices(OrderType.SELL, price);
        }
        return (cOrder, matchedBaseAmount, useQuoteAmount);
    }

    function _matchOrderAmount(
        uint256 orderId,
        Order memory order,
        uint256 targetId,
        Order storage target,
        uint256 price,
        DoubleLinkedList.U256 storage _orders
    ) private returns (address targetOwner, uint256 tradeAmount) {
        (targetOwner, tradeAmount) = (target.owner, Math.min(order.amount, target.amount));

        (uint256 sellId, uint256 buyId) = (order._type == OrderType.SELL ? (orderId, targetId) : (targetId, orderId));
        emit OrderMatched(sellId, buyId, price, tradeAmount, block.timestamp);

        //target 의 수량이 모두 거래되면 데이터를 지운다.
        if (tradeAmount == target.amount) {
            _removeOrder(targetId, CloseType.ALL_MATCH, _orders);
        } else {
            unchecked {
                target.amount -= tradeAmount;
            }
        }

        // order 의 모든 수량이 거래되면 종료를 유도한다.
        if (tradeAmount == order.amount) {
            order.amount = 0;
        } else {
            unchecked {
                order.amount -= tradeAmount;
            }
        }
    }

    function _copyOrder(Order memory order) private pure returns (Order memory) {
        return Order({_type: order._type, owner: order.owner, price: order.price, amount: order.amount});
    }

    function _removeOrder(uint256 orderId, CloseType _type, DoubleLinkedList.U256 storage _orders) private {
        if (!_orders.remove(orderId)) revert PairInvalidOrderId(orderId);
        delete _allOrders[orderId];
        emit OrderClosed(orderId, _type, block.timestamp);
    }

    function setPause(bool pause) external onlyOwner {
        if (pause) _pause();
        else _unpause();
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}

    uint256[36] __gap;
}
