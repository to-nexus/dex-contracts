// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.1.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.1.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.1.0/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.1.0/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.1.0/utils/PausableUpgradeable.sol";

import {IPair} from "./interfaces/IPair.sol";
import {ASCList} from "./lib/ASCList.sol";
import {DESCList} from "./lib/DESCList.sol";
import {List} from "./lib/List.sol";

contract Pair is ERC1967Proxy {
    constructor(
        address implementation,
        address owner,
        address router,
        address quote,
        address base,
        uint256 quoteTickSize,
        uint256 baseTickSize,
        address feeCollector,
        uint256 makerFeePermil,
        uint256 takerFeePermil
    )
        ERC1967Proxy(
            implementation,
            abi.encodeWithSelector(
                PairImpl.initialize.selector,
                owner,
                router,
                quote,
                base,
                quoteTickSize,
                baseTickSize,
                feeCollector,
                makerFeePermil,
                takerFeePermil
            )
        )
    {}
}

contract PairImpl is IPair, UUPSUpgradeable, OwnableUpgradeable, PausableUpgradeable {
    using SafeERC20 for IERC20;
    using Math for uint256;
    using List for List.U256;
    using DESCList for DESCList.U256;
    using ASCList for ASCList.U256;

    error PairInvalidInitializeData(bytes32);
    error PairInvalidRouter(address);
    error PairInvalidOrderType(OrderType);
    error PairInvalidPrice(uint256);
    error PairInvalidAmount(uint256);
    error PairInsufficientTradeVolume(uint256, uint256);
    error PairInvalidOrderId(uint256);
    error PairUnknownPrices(OrderType, uint256);
    error PairNotOwner(uint256, address);
    error PairInvalidTickSize(uint256, uint256, uint256);

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

    event TickSizeUpdated(
        uint256 beforeBaseTickSize, uint256 newBaseTickSize, uint256 beforeQuoteTickSize, uint256 newQuoteTickSize
    );

    event FeeCollectorUpdated(address before, address current);

    event FeeUpdated(
        uint256 beforeMakerFeePermil, uint256 newTakerFeePermil, uint256 beforeTakerFeePermil, uint256 newMakerFeePermil
    );

    event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount);

    address public ROUTER;
    IERC20 public override BASE; // immutable
    IERC20 public override QUOTE; // immutable
    uint256 public override DENOMINATOR; // immutable

    // 거래단위
    uint256 public baseTickSize;
    uint256 public quoteTickSize;
    uint256 public minTradeVolume; // (QUOTE) 거래당 최소 가치

    // 수수료
    address public feeCollector;
    uint32 public makerFeePermil; // 천(1000)분율
    uint32 public takerFeePermil; // 천(1000)분율

    // 주문
    uint256 private _orderIdCounter; // order id 생성기
    ASCList.U256 private _sellPrices; // sell order 의 가격 목록
    DESCList.U256 private _buyPrices; // buy order 의 가격 목록
    mapping(uint256 price => List.U256) private _sellOrders; // price => 판매 order id list
    mapping(uint256 price => List.U256) private _buyOrders; //  price => 구매 order id list
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
        uint256 _quoteTickSize,
        uint256 _baseTickSize,
        address _feeCollector,
        uint256 _makerFeePermil,
        uint256 _takerFeePermil
    ) external initializer {
        if (owner == address(0)) revert PairInvalidInitializeData("owner");
        if (router == address(0)) revert PairInvalidInitializeData("router");
        if (quote == address(0)) revert PairInvalidInitializeData("quote");
        if (base == address(0)) revert PairInvalidInitializeData("base");
        if (_quoteTickSize == 0) revert PairInvalidInitializeData("quoteTickSize");
        if (_baseTickSize == 0) revert PairInvalidInitializeData("baseTickSize");
        if (_feeCollector == address(0)) revert PairInvalidInitializeData("feeCollector");
        if (_makerFeePermil > 1000) revert PairInvalidInitializeData("makerFeePermil");
        if (_takerFeePermil > 1000) revert PairInvalidInitializeData("takerFeePermil");

        ROUTER = router;
        QUOTE = IERC20(quote);
        BASE = IERC20(base);
        DENOMINATOR = 10 ** IERC20Metadata(base).decimals();

        if (_quoteTickSize * _baseTickSize % DENOMINATOR != 0) {
            revert PairInvalidTickSize(_quoteTickSize, _baseTickSize, DENOMINATOR);
        }
        quoteTickSize = _quoteTickSize;
        baseTickSize = _baseTickSize;
        minTradeVolume = Math.mulDiv(_quoteTickSize, _baseTickSize, DENOMINATOR);

        feeCollector = _feeCollector;
        makerFeePermil = uint32(_makerFeePermil);
        takerFeePermil = uint32(_takerFeePermil);

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

    function limit(Order memory order, uint256 searchPrice, uint256 maxMatchCount)
        external
        override
        whenNotPaused
        onlyRouter
        returns (uint256 orderId)
    {
        if (order._type == OrderType.NONE) revert PairInvalidOrderType(OrderType.NONE);

        // 입력된 수량의 조건을 확인한다.
        if (order.price == 0 || order.price % quoteTickSize != 0) revert PairInvalidPrice(order.price);
        if (order.amount == 0 || order.amount % baseTickSize != 0) revert PairInvalidAmount(order.amount);

        orderId = ++_orderIdCounter;
        if (order._type == OrderType.SELL) {
            // 거래 주문 생성
            order = _createSellOrder(orderId, order, maxMatchCount);

            // 남은 잔액에 따른 분기 처리
            if (order.amount == 0) {
                emit OrderClosed(orderId, CloseType.ALL_MATCH, block.timestamp);
            } else {
                order.feePermil = makerFeePermil;
                _allOrders[orderId] = order;
                ASCList.push(_sellPrices, order.price, searchPrice);
                _sellOrders[order.price].push(orderId);
            }
        } else {
            // 거래 주문 생성
            uint256 matchedBaseAmount;
            uint256 useQuoteAmount;
            (order, matchedBaseAmount, useQuoteAmount) = _createBuyOrder(orderId, order, 0, maxMatchCount);

            // 구매일 경우, 주문과 동시에 채결되었다면, 설정된 가격보다 싼 가격으로 샀을수 있기때문에 해당 잔액을 돌려준다.
            if (matchedBaseAmount != 0) {
                uint256 measuredQuoteAmount = Math.mulDiv(order.price, matchedBaseAmount, DENOMINATOR);
                uint256 diffQuoteAmount = measuredQuoteAmount - useQuoteAmount;
                if (diffQuoteAmount != 0) QUOTE.safeTransfer(order.owner, diffQuoteAmount);
            }

            // 남은 잔액에 따른 분기 처리
            if (order.amount == 0) {
                emit OrderClosed(orderId, CloseType.ALL_MATCH, block.timestamp);
            } else {
                order.feePermil = makerFeePermil;
                _allOrders[orderId] = order;
                DESCList.push(_buyPrices, order.price, searchPrice);
                _buyOrders[order.price].push(orderId);
            }
        }
    }

    function market(Order memory order, uint256 spendAmount, uint256 maxMatchCount)
        external
        override
        whenNotPaused
        onlyRouter
    {
        if (order._type == OrderType.NONE) revert PairInvalidOrderType(OrderType.NONE);

        uint256 orderId = ++_orderIdCounter;
        if (order._type == OrderType.SELL) {
            if (spendAmount == 0 || spendAmount % baseTickSize != 0) revert PairInvalidAmount(spendAmount);
            order.price = 0; // 판매는 spendAmount 을 다 팔때까지 가격을 내린다.
            order.amount = spendAmount;

            order = _createSellOrder(orderId, order, maxMatchCount);
            if (order.amount != 0) BASE.transfer(order.owner, order.amount);
        } else {
            if (spendAmount < minTradeVolume) revert PairInsufficientTradeVolume(spendAmount, minTradeVolume);
            order.price = type(uint256).max; // 구매는 spendAmount 을 모두 소진할때까지 가격을 올린다.

            uint256 useQuoteAmount;
            (order,, useQuoteAmount) = _createBuyOrder(orderId, order, spendAmount, maxMatchCount);
            uint256 remainQuoteAmount = spendAmount - useQuoteAmount;
            if (remainQuoteAmount != 0) QUOTE.transfer(order.owner, remainQuoteAmount);
        }
        emit OrderClosed(orderId, CloseType.MARKET, block.timestamp);
    }

    function cancel(address caller, uint256[] memory orderIds) external override onlyRouter {
        uint256 length = orderIds.length;
        for (uint256 i = 0; i < length;) {
            uint256 orderId = orderIds[i];
            Order memory order = _allOrders[orderId];
            if (order._type == OrderType.NONE) revert PairInvalidOrderId(orderId);
            if (order.owner != caller) revert PairNotOwner(orderId, caller);

            List.U256 storage _orders;
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

            unchecked {
                ++i;
            }
        }
    }

    function _createSellOrder(uint256 orderId, Order memory order, uint256 maxMatchCount)
        private
        returns (Order memory)
    {
        if (order._type != OrderType.SELL) revert PairInvalidOrderType(order._type);

        // 1. 주문에 필요한 토큰을 가져온다.
        BASE.safeTransferFrom(_msgSender(), address(this), order.amount);
        emit OrderCreated(order.owner, orderId, order._type, order.price, order.amount, block.timestamp);

        // 2. 즉시 거래가 가능한 매물이 있다면 거래를 시킨다.
        // SELL 주문일 경우는 BUY 목록에서 비싼것 부터 찾으며, 입력된 price 보다 높거나 같은 구매 거래만 성사 시킨다.
        uint256 earnQuoteAmount;
        (order, earnQuoteAmount) = _tradeSellOrder(orderId, order, maxMatchCount);

        // 3. 즉시 거래된 수익은 판매자에게 전송한다.
        if (earnQuoteAmount != 0) {
            if (takerFeePermil == 0) {
                QUOTE.safeTransfer(order.owner, earnQuoteAmount);
            } else {
                uint256 fee = Math.mulDiv(earnQuoteAmount, takerFeePermil, 1000);
                QUOTE.safeTransfer(feeCollector, fee);
                QUOTE.safeTransfer(order.owner, earnQuoteAmount - fee);
            }
        }

        return order;
    }

    function _createBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 spendQuoteAmount, // Market Order 인 경우 이 값을 설정한다.
        uint256 maxMatchCount
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
            _tradeBuyOrder(orderId, order, spendQuoteAmount, maxMatchCount);

        // 3. 즉시 채결된 BASE 토큰을 전송한다.
        if (matchedBaseAmount != 0) {
            if (takerFeePermil == 0) {
                BASE.safeTransfer(order.owner, matchedBaseAmount);
            } else {
                uint256 fee = Math.mulDiv(matchedBaseAmount, takerFeePermil, 1000);
                BASE.safeTransfer(feeCollector, fee);
                BASE.safeTransfer(order.owner, matchedBaseAmount - fee);
            }
        }

        return (tradedOrder, matchedBaseAmount, useQuoteAmount);
    }

    // SELL 주문일 경우는 BUY 목록에서 비싼것 부터 찾으며, order.price 보다 높거나 같은 구매 거래만 성사 시킨다.
    function _tradeSellOrder(uint256 orderId, Order memory order, uint256 maxMatchCount)
        private
        returns (Order memory cOrder, uint256 earnQuoteAmount)
    {
        // (earnQuoteAmount) 판매로 인해 벌어들인 Quote 수량
        cOrder = _copyOrder(order);

        while (!_buyPrices.empty()) {
            uint256 price = _buyPrices.at(0);
            if (price < cOrder.price) break;
            List.U256 storage _orders = _buyOrders[price];

            uint256 length = List.length(_orders);
            while (length != 0) {
                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // 채결 수량 업데이트
                (address targetOwner, uint256 tradeAmount, uint256 targetFeePermil) =
                    _matchOrderAmount(orderId, cOrder, targetId, target, price, _orders);

                // 거래 성사
                if (targetFeePermil == 0) {
                    BASE.safeTransfer(targetOwner, tradeAmount);
                } else {
                    uint256 fee = Math.mulDiv(tradeAmount, targetFeePermil, 1000);
                    BASE.safeTransfer(feeCollector, fee);
                    BASE.safeTransfer(targetOwner, tradeAmount - fee);
                }
                earnQuoteAmount += Math.mulDiv(price, tradeAmount, DENOMINATOR);

                if (cOrder.amount == 0 || maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // while 이 끝나기 전이지만,
                        // cOrder 와 orders 의 마지막 target.amount 가 같았다면,
                        // _orders 가 비어있을 수 있다.
                        if (!_buyPrices.remove(price)) revert PairUnknownPrices(OrderType.BUY, price);
                    }
                    return (cOrder, earnQuoteAmount);
                }
                unchecked {
                    --length;
                    --maxMatchCount;
                }
            }
            // 여기에 왔다는 것은 price 에 해당하는 모든 주문을 매칭했다는 것을 의미,
            // _buyPrices 에서 price 를 지운다.
            if (!_buyPrices.remove(price)) revert PairUnknownPrices(OrderType.BUY, price);
        }
        return (cOrder, earnQuoteAmount);
    }

    // BUY 주문일 경우에는 SELL 목록에서 싼것부터 찾으며, 입력된 price 보다 낮거나 같은 판매 거래만 성사 시킨다.
    function _tradeBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 quoteAmount, // Market 거래시 사용 할 Quote 수량
        uint256 maxMatchCount
    ) private returns (Order memory cOrder, uint256 matchedBaseAmount, uint256 useQuoteAmount) {
        cOrder = _copyOrder(order);

        while (!_sellPrices.empty()) {
            uint256 price = _sellPrices.at(0);
            if (price > order.price) break;
            List.U256 storage _orders = _sellOrders[price];

            // Market 거래일 경우 해당 가격으로 최대 구매할 수 있는 수량을 계산한다.
            if (quoteAmount != 0) {
                uint256 buyAmount = Math.mulDiv(quoteAmount - useQuoteAmount, DENOMINATOR, price);
                buyAmount -= buyAmount % baseTickSize;
                cOrder.amount = buyAmount;
                if (buyAmount == 0) return (cOrder, matchedBaseAmount, useQuoteAmount);
            }

            // 해당 가격의 모든 SellOrder 를 순회한다.
            uint256 length = List.length(_orders);
            while (length != 0) {
                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // 거래 수량 업데이트
                (address targetOwner, uint256 tradeAmount, uint256 targetFeePermil) =
                    _matchOrderAmount(orderId, cOrder, targetId, target, price, _orders);
                // 거래 가치 계산
                uint256 tradeQuoteAmount = Math.mulDiv(price, tradeAmount, DENOMINATOR);

                // 거래 성사
                if (targetFeePermil == 0) {
                    QUOTE.safeTransfer(targetOwner, tradeQuoteAmount);
                } else {
                    uint256 fee = Math.mulDiv(tradeQuoteAmount, targetFeePermil, 1000);
                    QUOTE.safeTransfer(feeCollector, fee);
                    QUOTE.safeTransfer(targetOwner, tradeQuoteAmount - fee);
                }

                // 정보 업데이트
                matchedBaseAmount += tradeAmount;
                useQuoteAmount += tradeQuoteAmount;
                if (cOrder.amount == 0 || maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // while 이 끝나기 전이지만,
                        // cOrder 와 orders 의 마지막 target.amount 가 같았다면,
                        // _orders 가 비어있을 수 있다.
                        if (!_sellPrices.remove(price)) revert PairUnknownPrices(OrderType.SELL, price);
                    }
                    return (cOrder, matchedBaseAmount, useQuoteAmount);
                }
                unchecked {
                    --length;
                    --maxMatchCount;
                }
            }
            // 여기에 왔다는 것은 price 에 해당하는 모든 주문을 매칭했다는 것을 의미,
            // _sellPrices 에서 price 를 지운다.
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
        List.U256 storage _orders
    ) private returns (address targetOwner, uint256 tradeAmount, uint256 targetFeePermil) {
        (targetOwner, tradeAmount, targetFeePermil) =
            (target.owner, Math.min(order.amount, target.amount), target.feePermil);

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
        return Order({
            _type: order._type,
            owner: order.owner,
            feePermil: order.feePermil,
            price: order.price,
            amount: order.amount
        });
    }

    function _removeOrder(uint256 orderId, CloseType _type, List.U256 storage _orders) private {
        if (!_orders.remove(orderId)) revert PairInvalidOrderId(orderId);
        delete _allOrders[orderId];
        emit OrderClosed(orderId, _type, block.timestamp);
    }

    function setTickSize(uint256 _baseTickSize, uint256 _quoteTickSize) external onlyOwner {
        if (_quoteTickSize == 0 || _baseTickSize == 0) revert();
        if (_quoteTickSize * _baseTickSize % DENOMINATOR != 0) {
            revert PairInvalidTickSize(_quoteTickSize, _baseTickSize, DENOMINATOR);
        }

        emit TickSizeUpdated(baseTickSize, _baseTickSize, quoteTickSize, _quoteTickSize);

        baseTickSize = _baseTickSize;
        quoteTickSize = _quoteTickSize;
    }

    function setFeeCollector(address _feeCollector) external onlyOwner {
        if (_feeCollector == address(0)) revert();
        emit FeeCollectorUpdated(feeCollector, _feeCollector);

        feeCollector = _feeCollector;
    }

    function setFee(uint256 _makerFeePermil, uint256 _takerFeePermil) external onlyOwner {
        if (_makerFeePermil > 1000 || _takerFeePermil > 1000) revert();
        emit FeeUpdated(makerFeePermil, _makerFeePermil, takerFeePermil, _takerFeePermil);

        makerFeePermil = uint32(_makerFeePermil);
        takerFeePermil = uint32(_takerFeePermil);
    }

    function skim(IERC20 erc20, address to, uint256 amount) external onlyOwner whenPaused {
        erc20.safeTransfer(to, amount);
        emit Skim(_msgSender(), address(erc20), to, amount);
    }

    function setPause(bool pause) external onlyOwner {
        if (pause) _pause();
        else _unpause();
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}

    uint256[36] private __gap;
}
