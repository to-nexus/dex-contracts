// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";

import {PausableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/utils/PausableUpgradeable.sol";

import {IOwnable} from "./interfaces/IOwnable.sol";
import {IPair} from "./interfaces/IPair.sol";
import {ASCList} from "./lib/ASCList.sol";
import {DESCList} from "./lib/DESCList.sol";
import {List} from "./lib/List.sol";

contract PairImpl is IPair, UUPSUpgradeable, PausableUpgradeable {
    using SafeERC20 for IERC20;
    using Math for uint256;
    using List for List.U256;
    using DESCList for DESCList.U256;
    using ASCList for ASCList.U256;

    error PairInvalidReserve(address);
    error PairInvalidInitializeData(bytes32);
    error PairInvalidRouter(address);
    error PairInvalidOrderSide(OrderSide);
    error PairInvalidPrice(uint256);
    error PairInvalidAmount(uint256);
    error PairInsufficientTradeVolume(uint256, uint256);
    error PairInvalidOrderId(uint256);
    error PairUnknownPrices(OrderSide, uint256);
    error PairNotOwner(uint256, address);
    error PairInvalidTickSize(uint256, uint256, uint256);
    error PairFillOrKill(address);

    event OrderCreated(
        address indexed owner,
        uint256 indexed orderId,
        OrderSide indexed side,
        uint256 price,
        uint256 amount,
        uint256 timestamp
    );
    event OrderMatched(
        uint256 indexed sellId, uint256 indexed buyId, uint256 indexed price, uint256 amount, uint256 timestamp
    );
    event OrderClosed(uint256 indexed orderId, CloseType indexed closeType, uint256 timestamp);
    event TickSizeUpdated(
        uint256 beforeBaseTickSize, uint256 newBaseTickSize, uint256 beforeQuoteTickSize, uint256 newQuoteTickSize
    );
    event RouterUpdated(address before, address current);
    event FeeCollectorUpdated(address before, address current);
    event FeeUpdated(uint256 before, uint256 current);
    event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount);

    address public MARKET; // immutable
    address public ROUTER; // immutable
    IERC20 public BASE; // immutable
    IERC20 public QUOTE; // immutable
    uint256 public DENOMINATOR; // immutable ( == 10 ** BASE.decimals())

    // reserves
    uint256 public baseReserve;
    uint256 public quoteReserve;

    // tick size
    uint256 public baseTickSize;
    uint256 public quoteTickSize;
    uint256 public minTradeVolume; // [QUOTE] Math.mulDiv(_quoteTickSize, _baseTickSize, DENOMINATOR);

    // fee
    address public feeCollector;
    uint32 public feePermil;

    // orders
    uint256 private _orderIdCounter;
    ASCList.U256 private _sellPrices; // This is the price list for sell orders. It will be stored in ascending order.
    DESCList.U256 private _buyPrices; // This is the price list for buy orders. It will be stored in descending order.
    mapping(uint256 price => List.U256) private _sellOrders; // price => sell order id list (For the same price, orders will be stored in chronological order.)
    mapping(uint256 price => List.U256) private _buyOrders; //  price => buy order id list (For the same price, orders will be stored in chronological order.)
    mapping(uint256 orderId => Order) private _allOrders;

    uint256[31] private __gap;

    modifier onlyOwner() {
        // The Pair is the same as the Owner of the Router.
        if (_msgSender() != IOwnable(MARKET).owner()) revert IOwnable.OwnableUnauthorizedAccount(_msgSender());
        _;
    }

    modifier onlyRouter() {
        if (_msgSender() != ROUTER) revert PairInvalidRouter(_msgSender());
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address router,
        address quote,
        address base,
        uint256 _quoteTickSize, // tick size for quote token
        uint256 _baseTickSize, // lot size for base token
        address _feeCollector,
        uint256 _feePermil
    ) external initializer {
        if (router == address(0)) revert PairInvalidInitializeData("router");
        if (quote == address(0)) revert PairInvalidInitializeData("quote");
        if (base == address(0)) revert PairInvalidInitializeData("base");
        if (_quoteTickSize == 0) revert PairInvalidInitializeData("quoteTickSize");
        if (_baseTickSize == 0) revert PairInvalidInitializeData("baseTickSize");
        if (_feeCollector == address(0)) revert PairInvalidInitializeData("feeCollector");
        if (feePermil > 1000) revert PairInvalidInitializeData("feePermil");

        MARKET = _msgSender();
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
        feePermil = uint32(_feePermil);

        __Pausable_init();
    }

    //  #    # # ###### #    #  ####
    //  #    # # #      #    # #
    //  #    # # #####  #    #  ####
    //  #    # # #      # ## #      #
    //   #  #  # #      ##  ## #    #
    //    ##   # ###### #    #  ####

    function getTokenConfig() external view returns (TokenConfig memory) {
        return TokenConfig({QUOTE: QUOTE, BASE: BASE, DENOMINATOR: DENOMINATOR});
    }

    function orderById(uint256 id) external view returns (Order memory) {
        return _allOrders[id];
    }

    function ticks() external view returns (uint256[] memory sellPrices, uint256[] memory buyPrices) {
        sellPrices = _sellPrices.values();
        buyPrices = _buyPrices.values();
    }

    function ordersByPrices(OrderSide side, uint256[] memory prices) external view returns (uint256[][] memory) {
        mapping(uint256 price => List.U256) storage orders = side == OrderSide.SELL ? _sellOrders : _buyOrders;
        uint256 length = prices.length;

        uint256[][] memory orderIds = new uint256[][](length);
        for (uint256 i = 0; i < length;) {
            orderIds[i] = orders[prices[i]].values();
            unchecked {
                ++i;
            }
        }
        return orderIds;
    }

    //  ###### #    # ######  ####  #    # ##### ######  ####
    //  #       #  #  #      #    # #    #   #   #      #
    //  #####    ##   #####  #      #    #   #   #####   ####
    //  #        ##   #      #      #    #   #   #           #
    //  #       #  #  #      #    # #    #   #   #      #    #
    //  ###### #    # ######  ####   ####    #   ######  ####

    function limit(Order memory order, LimitConstraints constraints, uint256 searchPrice, uint256 maxMatchCount)
        external
        override
        whenNotPaused
        onlyRouter
        returns (uint256 orderId)
    {
        // Check the conditions of the entered quantity.
        if (order.price == 0 || order.price % quoteTickSize != 0) revert PairInvalidPrice(order.price);
        if (order.amount == 0 || order.amount % baseTickSize != 0) revert PairInvalidAmount(order.amount);

        orderId = ++_orderIdCounter;
        uint256 mustRemainQuoteAmount;
        (bool isSellOrder) = order.side == OrderSide.SELL;
        (order, mustRemainQuoteAmount) = isSellOrder
            ? _executeSellOrder(orderId, order, maxMatchCount)
            : _executeBuyOrder(orderId, order, 0, maxMatchCount);

        if (order.amount == 0) {
            emit OrderClosed(orderId, CloseType.ALL_MATCH, block.timestamp);
        } else {
            if (constraints == LimitConstraints.IMMEDIATE_OR_CANCEL) {
                emit OrderClosed(orderId, CloseType.IMMEDIATE_OR_CANCEL, block.timestamp);
                // Return the remaining balance.
                if (isSellOrder) BASE.safeTransfer(order.owner, order.amount);
                // The quantity returned as a quote can occur under all conditions except FILL_OR_KILL,
                // so it is refunded collectively at the end of the function before termination.
            } else if (constraints == LimitConstraints.FILL_OR_KILL) {
                revert PairFillOrKill(order.owner);
            } else {
                if (isSellOrder) {
                    baseReserve += order.amount;

                    order.feePermil = feePermil; // For sell orders, a fee is charged when acting as a maker.
                    _allOrders[orderId] = order;
                    ASCList.push(_sellPrices, order.price, searchPrice);
                    _sellOrders[order.price].push(orderId);
                } else {
                    quoteReserve += Math.mulDiv(order.price, order.amount, DENOMINATOR);

                    order.feePermil = 0; // Buy orders have no fees.
                    _allOrders[orderId] = order;
                    DESCList.push(_buyPrices, order.price, searchPrice);
                    _buyOrders[order.price].push(orderId);
                }
            }
        }
        if (!isSellOrder) _returnRemainQuote(order.owner, mustRemainQuoteAmount);
    }

    function market(Order memory order, uint256 spendAmount, uint256 maxMatchCount)
        external
        override
        whenNotPaused
        onlyRouter
    {
        uint256 orderId = ++_orderIdCounter;
        uint256 mustRemainQuoteAmount;
        if (order.side == OrderSide.SELL) {
            if (spendAmount == 0 || spendAmount % baseTickSize != 0) revert PairInvalidAmount(spendAmount);
            order.price = 0; // For selling, the price is lowered until the entire spendAmount is sold.
            order.amount = spendAmount;

            // Return the remaining balance.
            (order, mustRemainQuoteAmount) = _executeSellOrder(orderId, order, maxMatchCount);
            if (order.amount != 0) BASE.safeTransfer(order.owner, order.amount);
        } else {
            if (spendAmount < minTradeVolume) revert PairInsufficientTradeVolume(spendAmount, minTradeVolume);
            order.price = type(uint256).max; // For buying, the price is increased until the entire spendAmount is exhausted.

            // Return the remaining balance.
            (order, mustRemainQuoteAmount) = _executeBuyOrder(orderId, order, spendAmount, maxMatchCount);
            _returnRemainQuote(order.owner, mustRemainQuoteAmount);
        }

        emit OrderClosed(orderId, CloseType.MARKET, block.timestamp);
    }

    function cancel(address caller, uint256[] memory orderIds) external override onlyRouter {
        uint256 length = orderIds.length;
        for (uint256 i = 0; i < length;) {
            uint256 orderId = orderIds[i];
            Order memory order = _allOrders[orderId];
            if (order.owner == address(0)) continue;
            if (order.owner != caller) revert PairNotOwner(orderId, caller);

            _cancelOrder(orderId, order, CloseType.CANCEL);
            unchecked {
                ++i;
            }
        }
    }

    //  #####  #####  # #    #   ##   ##### ######  ####
    //  #    # #    # # #    #  #  #    #   #      #
    //  #    # #    # # #    # #    #   #   #####   ####
    //  #####  #####  # #    # ######   #   #           #
    //  #      #   #  #  #  #  #    #   #   #      #    #
    //  #      #    # #   ##   #    #   #   ######  ####

    function _executeSellOrder(uint256 orderId, Order memory order, uint256 maxMatchCount)
        private
        returns (Order memory, uint256)
    {
        if (order.side != OrderSide.SELL) revert PairInvalidOrderSide(order.side);

        // 1. Verify that the required tokens for the order have been deposited.
        if (BASE.balanceOf(address(this)) < baseReserve + order.amount) revert PairInvalidReserve(address(BASE));
        emit OrderCreated(order.owner, orderId, order.side, order.price, order.amount, block.timestamp);

        // 2. If there are immediately tradable orders, execute the trade.
        //    For a SELL order, search from the most expensive price in the BUY list
        //    and only match with buy orders that have a price equal to or higher than the input price.
        uint256 earnQuoteAmount;
        (order, earnQuoteAmount) = _matchSellOrder(orderId, order, maxMatchCount);

        // 3. Immediately transfer the proceeds from the trade to the seller.
        if (earnQuoteAmount != 0) {
            quoteReserve -= earnQuoteAmount;
            if (feePermil == 0) {
                QUOTE.safeTransfer(order.owner, earnQuoteAmount);
            } else {
                uint256 fee = Math.mulDiv(earnQuoteAmount, feePermil, 1000);
                QUOTE.safeTransfer(feeCollector, fee);
                QUOTE.safeTransfer(order.owner, earnQuoteAmount - fee);
            }
        }

        return (order, 0);
    }

    function _executeBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 spendQuoteAmount, // Set this value if it is a Market Order.
        uint256 maxMatchCount
    ) private returns (Order memory, uint256) {
        if (order.side != OrderSide.BUY) revert PairInvalidOrderSide(order.side);

        // Verify the conditions of the entered quantity.
        uint256 quoteAmount;
        if (spendQuoteAmount != 0) quoteAmount = spendQuoteAmount;
        else quoteAmount = Math.mulDiv(order.price, order.amount, DENOMINATOR);

        // 1. Verify that the required tokens for the order have been deposited.
        (bool ok, uint256 skimQuoteAmount) = Math.trySub(QUOTE.balanceOf(address(this)), quoteReserve + quoteAmount);
        if (!ok) revert PairInvalidReserve(address(QUOTE));
        emit OrderCreated(order.owner, orderId, order.side, order.price, order.amount, block.timestamp);

        // 2. If there are immediately tradable orders, execute the trade.
        //    For a BUY order, search from the cheapest price in the SELL list
        //    and only match with sell orders that have a price equal to or lower than the input price.
        (Order memory tradedOrder, uint256 buyBaseAmount) =
            _matchBuyOrder(orderId, order, spendQuoteAmount, maxMatchCount);

        // 3. Transfer the immediately settled BASE tokens.
        if (buyBaseAmount != 0) {
            baseReserve -= buyBaseAmount;
            // Purchased tokens are not subject to fees.
            BASE.safeTransfer(order.owner, buyBaseAmount);
        }

        return (tradedOrder, skimQuoteAmount);
    }

    // For a SELL order, search from the most expensive price in the BUY list
    // and only execute trades where order.price is equal to or lower than the buy order price.
    function _matchSellOrder(uint256 orderId, Order memory order, uint256 maxMatchCount)
        private
        returns (Order memory cOrder, uint256 earnQuoteAmount)
    {
        // (earnQuoteAmount) The amount of Quote earned from the sale.
        cOrder = _copyOrder(order);

        while (!_buyPrices.empty()) {
            uint256 price = _buyPrices.at(0);
            if (price < cOrder.price) break;
            List.U256 storage _orders = _buyOrders[price];

            uint256 length = List.length(_orders);
            while (length != 0) {
                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // Update the settled quantity.
                (address targetOwner, uint256 tradeAmount,) =
                    _matchOrderAmount(orderId, cOrder, targetId, target, price, _orders);

                // Trade executed.
                BASE.safeTransfer(targetOwner, tradeAmount);

                // Update information.
                earnQuoteAmount += Math.mulDiv(price, tradeAmount, DENOMINATOR);
                if (cOrder.amount == 0 || maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // Although the `while` loop has not yet ended,
                        // if `cOrder` and the last `target.amount` in `orders` are the same,
                        // `_orders` may be empty.
                        if (!_buyPrices.remove(price)) revert PairUnknownPrices(OrderSide.BUY, price);
                    }
                    return (cOrder, earnQuoteAmount);
                }
                unchecked {
                    --length;
                    --maxMatchCount;
                }
            }
            // Reaching this point means that all orders at the given `price` have been matched,
            // so remove `price` from `_buyPrices`.
            if (!_buyPrices.remove(price)) revert PairUnknownPrices(OrderSide.BUY, price);
        }
        return (cOrder, earnQuoteAmount);
    }

    // For a BUY order, search from the cheapest price in the SELL list
    // and only execute trades where the sell order price is equal to or lower than the input price.
    function _matchBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 quoteAmount, // Quote amount to be used for Market trades.
        uint256 maxMatchCount
    ) private returns (Order memory cOrder, uint256 matchedBaseAmount) {
        cOrder = _copyOrder(order);

        uint256 useQuoteAmount;
        while (!_sellPrices.empty()) {
            uint256 price = _sellPrices.at(0);
            if (price > order.price) break;
            List.U256 storage _orders = _sellOrders[price];

            // If it is a Market trade, calculate the maximum quantity that can be purchased at the given price.
            if (quoteAmount != 0) {
                uint256 buyAmount = Math.mulDiv(quoteAmount - useQuoteAmount, DENOMINATOR, price);
                buyAmount -= buyAmount % baseTickSize;
                cOrder.amount = buyAmount;
                if (buyAmount == 0) return (cOrder, matchedBaseAmount);
            }

            uint256 length = List.length(_orders);
            while (length != 0) {
                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // Update the settled quantity.
                (address targetOwner, uint256 tradeAmount, uint256 targetFeePermil) =
                    _matchOrderAmount(orderId, cOrder, targetId, target, price, _orders);
                uint256 tradeQuoteAmount = Math.mulDiv(price, tradeAmount, DENOMINATOR);

                // Trade executed. ( Calculate using the fee rate at the time the seller registered the sale.)
                if (targetFeePermil == 0) {
                    QUOTE.safeTransfer(targetOwner, tradeQuoteAmount);
                } else {
                    // The seller pays the fee.
                    uint256 fee = Math.mulDiv(tradeQuoteAmount, targetFeePermil, 1000);
                    QUOTE.safeTransfer(feeCollector, fee);
                    QUOTE.safeTransfer(targetOwner, tradeQuoteAmount - fee);
                }

                // Update information.
                matchedBaseAmount += tradeAmount;
                useQuoteAmount += tradeQuoteAmount;
                if (cOrder.amount == 0 || maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // Although the `while` loop has not yet ended,
                        // if `cOrder` and the last `target.amount` in `orders` are the same,
                        // `_orders` may be empty.
                        if (!_sellPrices.remove(price)) revert PairUnknownPrices(OrderSide.SELL, price);
                    }
                    return (cOrder, matchedBaseAmount);
                }
                unchecked {
                    --length;
                    --maxMatchCount;
                }
            }
            // Reaching this point means that all orders at the given `price` have been matched,
            // so remove `price` from `_sellPrices`.
            if (!_sellPrices.remove(price)) revert PairUnknownPrices(OrderSide.SELL, price);
        }
        return (cOrder, matchedBaseAmount);
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

        (uint256 sellId, uint256 buyId) = (order.side == OrderSide.SELL ? (orderId, targetId) : (targetId, orderId));
        emit OrderMatched(sellId, buyId, price, tradeAmount, block.timestamp);

        // If the entire quantity of target is traded, remove the data.
        if (tradeAmount == target.amount) {
            _removeOrder(targetId, CloseType.ALL_MATCH, _orders);
        } else {
            unchecked {
                target.amount -= tradeAmount;
            }
        }

        // If the entire quantity of order is traded, trigger termination.
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
            side: order.side,
            owner: order.owner,
            feePermil: order.feePermil,
            price: order.price,
            amount: order.amount
        });
    }

    function _cancelOrder(uint256 orderId, Order memory order, CloseType _type) private {
        List.U256 storage _orders;
        bool isSellOrder = order.side == OrderSide.SELL;

        // Return the tokens held by the contract.
        if (isSellOrder) {
            _orders = _sellOrders[order.price];
            if (order.amount != 0) {
                BASE.safeTransfer(order.owner, order.amount);
                baseReserve -= order.amount;
            }
        } else {
            _orders = _buyOrders[order.price];
            uint256 returnQuoteAmount = Math.mulDiv(order.price, order.amount, DENOMINATOR);
            if (returnQuoteAmount != 0) {
                QUOTE.safeTransfer(order.owner, returnQuoteAmount);
                quoteReserve -= returnQuoteAmount;
            }
        }

        // Remove the corresponding data.
        _removeOrder(orderId, _type, _orders);

        // If this was the last entry for the given price, remove the price.
        if (_orders.empty()) {
            if (isSellOrder) _sellPrices.remove(order.price);
            else _buyPrices.remove(order.price);
        }
    }

    function _removeOrder(uint256 orderId, CloseType closeType, List.U256 storage _orders) private {
        if (!_orders.remove(orderId)) revert PairInvalidOrderId(orderId);
        delete _allOrders[orderId];
        emit OrderClosed(orderId, closeType, block.timestamp);
    }

    function _returnRemainQuote(address to, uint256 mustRemainQuoteAmount) private {
        uint256 remainQuoteAmount = QUOTE.balanceOf(address(this)) - (mustRemainQuoteAmount + quoteReserve);
        if (remainQuoteAmount != 0) QUOTE.safeTransfer(to, remainQuoteAmount);
    }

    //    ##   #    # ##### #    #  ####  #####  # ######   ##   ##### #  ####  #    #
    //   #  #  #    #   #   #    # #    # #    # #     #   #  #    #   # #    # ##   #
    //  #    # #    #   #   ###### #    # #    # #    #   #    #   #   # #    # # #  #
    //  ###### #    #   #   #    # #    # #####  #   #    ######   #   # #    # #  # #
    //  #    # #    #   #   #    # #    # #   #  #  #     #    #   #   # #    # #   ##
    //  #    #  ####    #   #    #  ####  #    # # ###### #    #   #   #  ####  #    #

    function setTickSize(uint256 _baseTickSize, uint256 _quoteTickSize) external onlyOwner {
        if (_quoteTickSize == 0) revert PairInvalidInitializeData("quoteTickSize");
        if (_baseTickSize == 0) revert PairInvalidInitializeData("baseTickSize");
        if (_quoteTickSize * _baseTickSize % DENOMINATOR != 0) {
            revert PairInvalidTickSize(_quoteTickSize, _baseTickSize, DENOMINATOR);
        }
        emit TickSizeUpdated(baseTickSize, _baseTickSize, quoteTickSize, _quoteTickSize);

        baseTickSize = _baseTickSize;
        quoteTickSize = _quoteTickSize;
    }

    function setFeeCollector(address _feeCollector) external onlyOwner {
        if (_feeCollector == address(0)) revert PairInvalidInitializeData("feeCollector");
        emit FeeCollectorUpdated(feeCollector, _feeCollector);

        feeCollector = _feeCollector;
    }

    function setFee(uint256 _feePermil) external onlyOwner {
        if (_feePermil > 1000) revert PairInvalidInitializeData("feePermil");
        emit FeeUpdated(feePermil, _feePermil);

        feePermil = uint32(_feePermil);
    }

    function skim(IERC20 erc20, address to, uint256 amount) external onlyOwner {
        if (amount == 0) return;

        if (erc20 == BASE && BASE.balanceOf(address(this)) - amount < baseReserve) {
            revert PairInvalidReserve(address(BASE));
        } else if (erc20 == QUOTE && QUOTE.balanceOf(address(this)) - amount < quoteReserve) {
            revert PairInvalidReserve(address(QUOTE));
        }

        erc20.safeTransfer(to, amount);
        emit Skim(_msgSender(), address(erc20), to, amount);
    }

    function emergencyCancel(uint256[] memory orderIds) external whenPaused onlyOwner {
        uint256 length = orderIds.length;
        for (uint256 i = 0; i < length;) {
            uint256 orderId = orderIds[i];
            Order memory order = _allOrders[orderId];
            if (order.owner == address(0)) continue;

            _cancelOrder(orderId, order, CloseType.EMERGENCY);
            unchecked {
                ++i;
            }
        }
    }

    function setPause(bool pause) external onlyOwner {
        if (pause) _pause();
        else _unpause();
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
