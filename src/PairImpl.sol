// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";

import {PausableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/utils/PausableUpgradeable.sol";

import {IMarket} from "./interfaces/IMarket.sol";
import {IOwnable} from "./interfaces/IOwnable.sol";
import {IPair} from "./interfaces/IPair.sol";
// import {ASCList} from "./lib/ASCList.sol";
// import {DESCList} from "./lib/DESCList.sol";
import {List} from "./lib/List.sol";

contract PairImpl is IPair, IOwnable, UUPSUpgradeable, PausableUpgradeable {
    using SafeERC20 for IERC20;
    using Math for uint256;
    using List for List.U256;
    // using DESCList for DESCList.U256;
    // using ASCList for ASCList.U256;

    error PairInvalidReserve(address);
    error PairInvalidInitializeData(bytes32);
    error PairInvalidRouter(address);
    error PairInvalidOrderSide(OrderSide);
    error PairInvalidPrice(uint256);
    error PairInvalidPrevPrice(OrderSide, uint256, uint256);
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
    event FeeCollect(
        uint256 indexed orderId,
        address indexed owner,
        uint256 amount,
        address indexed recipient,
        uint256 feeBps,
        uint256 fee,
        uint256 value
    );
    event TickSizeUpdated(uint256 beforeLotSize, uint256 newLotSize, uint256 beforeTickSize, uint256 newTickSize);
    event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount);

    // slots
    // keccak256(abi.encode(uint256(keccak256("crossdex.pair.matchedprice")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant _matchedPriceSlot = 0xfd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd00;
    // keccak256(abi.encode(uint256(keccak256("crossdex.pair.feecollector")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant _feeCollectorSlot = 0xd6aa07baf8485abf9d26fecf4c935d75b50a73e678db02b944bd3ac875982300;
    // keccak256(abi.encode(uint256(keccak256("crossdex.pair.feebps")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant _feeBpsSlot = 0x1d2ff3fa6980aeeebca4e94965520da48983006e9b1115c1c853cbb10d943d00;

    address public MARKET; // immutable
    address public ROUTER; // immutable
    IERC20 public BASE; // immutable
    IERC20 public QUOTE; // immutable
    uint256 public DENOMINATOR; // immutable ( == 10 ** BASE.decimals())

    // reserves
    uint256 public baseReserve;
    uint256 public quoteReserve;

    // latest
    uint256 public matchedPrice;
    uint256 public matchedAt;

    // tick size
    uint256 public tickSize;
    uint256 public lotSize;
    uint256 public minTradeVolume; // [QUOTE] Math.mulDiv(_tickSize, _lotSize, DENOMINATOR);

    // orders
    uint256 private _orderIdCounter;
    List.U256[2] private _prices; // 0: sell, 1: buy (IPair.OrderSide)
    // ASCList.U256 private _sellPrices; // This is the price list for sell orders. It will be stored in ascending order. // 0: sell, 1: buy
    // DESCList.U256 private _buyPrices; // This is the price list for buy orders. It will be stored in descending order.
    mapping(uint256 price => List.U256) private _sellOrders; // price => sell order id list (For the same price, orders will be stored in chronological order.)
    mapping(uint256 price => List.U256) private _buyOrders; //  price => buy order id list (For the same price, orders will be stored in chronological order.)
    mapping(uint256 orderId => Order) private _allOrders;

    uint256[32] private __gap;

    modifier onlyOwner() {
        // The Pair is the same as the Owner of the Market.
        if (_msgSender() != owner()) revert OwnableUnauthorizedAccount(_msgSender());
        _;
    }

    modifier onlyRouter() {
        if (_msgSender() != ROUTER) revert PairInvalidRouter(_msgSender());
        _;
    }

    modifier cacheFeeInfos() {
        _cacheFeeInfos();
        _;
    }

    modifier setLatest() {
        _;
        _setLatest();
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address router,
        address quote,
        address base,
        uint256 _tickSize, // tick size for quote token
        uint256 _lotSize // lot size for base token
    ) external initializer {
        if (router == address(0)) revert PairInvalidInitializeData("router");
        if (quote == address(0)) revert PairInvalidInitializeData("quote");
        if (base == address(0)) revert PairInvalidInitializeData("base");
        if (_tickSize == 0) revert PairInvalidInitializeData("tickSize");
        if (_lotSize == 0) revert PairInvalidInitializeData("lotSize");

        MARKET = _msgSender();
        ROUTER = router;
        QUOTE = IERC20(quote);
        BASE = IERC20(base);
        DENOMINATOR = 10 ** IERC20Metadata(base).decimals();

        if (_tickSize * _lotSize % DENOMINATOR != 0) revert PairInvalidTickSize(_tickSize, _lotSize, DENOMINATOR);

        tickSize = _tickSize;
        lotSize = _lotSize;
        minTradeVolume = Math.mulDiv(_tickSize, _lotSize, DENOMINATOR);

        __Pausable_init();
    }

    //  #    # # ###### #    #  ####
    //  #    # # #      #    # #
    //  #    # # #####  #    #  ####
    //  #    # # #      # ## #      #
    //   #  #  # #      ##  ## #    #
    //    ##   # ###### #    #  ####

    function getConfig() external view returns (Config memory) {
        return Config({QUOTE: QUOTE, BASE: BASE, DENOMINATOR: DENOMINATOR});
    }

    function orderById(uint256 id) external view returns (Order memory) {
        return _allOrders[id];
    }

    function ticks() external view returns (uint256[] memory sellPrices, uint256[] memory buyPrices) {
        sellPrices = _prices[uint8(OrderSide.SELL)].values();
        buyPrices = _prices[uint8(OrderSide.BUY)].values();
    }

    function tickSizes() external view returns (uint256 tick, uint256 lot) {
        tick = tickSize;
        lot = lotSize;
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

    function owner() public view returns (address) {
        return IOwnable(MARKET).owner();
    }

    function findPrevPrice(OrderSide side, uint256 price, uint256[2] calldata adjacent, uint256 findMaxCount)
        external
        view
        returns (uint256)
    {
        if (side == OrderSide.SELL) {
            // For a SELL order, search from the most expensive price in the BUY list
            // and only match with buy orders that have a price equal to or higher than the input price.
            return _prices[uint8(side)].findASCPrev(price, adjacent, findMaxCount);
        } else {
            // For a BUY order, search from the cheapest price in the SELL list
            // and only match with sell orders that have a price equal to or lower than the input price.
            return _prices[uint8(side)].findDESCPrev(price, adjacent, findMaxCount);
        }
    }

    //  ###### #    # ######  ####  #    # ##### ######  ####
    //  #       #  #  #      #    # #    #   #   #      #
    //  #####    ##   #####  #      #    #   #   #####   ####
    //  #        ##   #      #      #    #   #   #           #
    //  #       #  #  #      #    # #    #   #   #      #    #
    //  ###### #    # ######  ####   ####    #   ######  ####

    function submitLimitOrder(
        Order memory order,
        LimitConstraints constraints,
        uint256 prevPrice,
        uint256 maxMatchCount
    ) external override whenNotPaused onlyRouter cacheFeeInfos returns (uint256 orderId) {
        // Check the conditions of the entered quantity.
        if (order.price == 0 || order.price % tickSize != 0) revert PairInvalidPrice(order.price);
        if (order.amount == 0 || order.amount % lotSize != 0) revert PairInvalidAmount(order.amount);

        orderId = ++_orderIdCounter;
        (bool isSellOrder) = order.side == OrderSide.SELL;

        uint256 mustRemainQuoteAmount = isSellOrder
            ? _executeSellOrder(orderId, order, maxMatchCount)
            : _executeBuyOrder(orderId, order, 0, maxMatchCount);

        if (order.amount == 0) {
            emit OrderClosed(orderId, CloseType.COMPLETED, block.timestamp);
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
                    if (prevPrice != 0 && order.price < prevPrice) {
                        // If the price is lowerthan the previous price, it is not possible to register.
                        revert PairInvalidPrevPrice(OrderSide.SELL, order.price, prevPrice);
                    }
                    baseReserve += order.amount;

                    order.feeBps = _feeBps(); // For sell orders, a fee is charged when acting as a maker.
                    _allOrders[orderId] = order;
                    _prices[uint8(OrderSide.SELL)].insert(order.price, prevPrice);
                    // ASCList.push(_sellPrices, order.price, adjacent);
                    _sellOrders[order.price].push(orderId);
                } else {
                    if (prevPrice != 0 && order.price > prevPrice) {
                        // If the price is higher than the previous price, it is not possible to register.
                        revert PairInvalidPrevPrice(OrderSide.BUY, order.price, prevPrice);
                    }
                    quoteReserve += Math.mulDiv(order.price, order.amount, DENOMINATOR);

                    order.feeBps = 0; // Buy orders have no fees.
                    _allOrders[orderId] = order;
                    _prices[uint8(OrderSide.BUY)].insert(order.price, prevPrice);
                    _buyOrders[order.price].push(orderId);
                }
            }
        }
        if (!isSellOrder) _returnRemainQuote(order.owner, mustRemainQuoteAmount);
    }

    function submitMarketOrder(Order memory order, uint256 spendAmount, uint256 maxMatchCount)
        external
        override
        whenNotPaused
        onlyRouter
        cacheFeeInfos
    {
        uint256 orderId = ++_orderIdCounter;
        if (order.side == OrderSide.SELL) {
            if (spendAmount == 0 || spendAmount % lotSize != 0) revert PairInvalidAmount(spendAmount);
            order.price = 0; // For selling, the price is lowered until the entire spendAmount is sold.
            order.amount = spendAmount;

            // Return the remaining balance.
            _executeSellOrder(orderId, order, maxMatchCount);
            if (order.amount != 0) BASE.safeTransfer(order.owner, order.amount);
        } else {
            if (spendAmount < minTradeVolume) revert PairInsufficientTradeVolume(spendAmount, minTradeVolume);
            order.price = type(uint256).max; // For buying, the price is increased until the entire spendAmount is exhausted.

            // Return the remaining balance.
            uint256 mustRemainQuoteAmount = _executeBuyOrder(orderId, order, spendAmount, maxMatchCount);
            _returnRemainQuote(order.owner, mustRemainQuoteAmount);
        }

        emit OrderClosed(orderId, CloseType.COMPLETED, block.timestamp);
    }

    function cancelOrder(address caller, uint256[] memory orderIds) external override onlyRouter {
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

    function _executeSellOrder(uint256 orderId, Order memory order, uint256 maxMatchCount) private returns (uint256) {
        if (order.side != OrderSide.SELL) revert PairInvalidOrderSide(order.side);

        // 1. Verify that the required tokens for the order have been deposited.
        if (BASE.balanceOf(address(this)) < baseReserve + order.amount) revert PairInvalidReserve(address(BASE));
        emit OrderCreated(order.owner, orderId, order.side, order.price, order.amount, block.timestamp);

        // 2. If there are immediately tradable orders, execute the trade.
        //    For a SELL order, search from the most expensive price in the BUY list
        //    and only match with buy orders that have a price equal to or higher than the input price.
        uint256 earnQuoteAmount = _matchSellOrder(orderId, order, maxMatchCount);

        // 3. Immediately transfer the proceeds from the trade to the seller.
        if (earnQuoteAmount != 0) {
            (address feeCollector, uint32 feeBps) = (_feeCollector(), _feeBps());
            quoteReserve -= earnQuoteAmount;
            _exchangeQuote(orderId, order.owner, earnQuoteAmount, feeCollector, feeBps);
        }

        return 0;
    }

    function _executeBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 spendQuoteAmount, // Set this value if it is a Market Order.
        uint256 maxMatchCount
    ) private returns (uint256) {
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
        uint256 buyBaseAmount = _matchBuyOrder(orderId, order, spendQuoteAmount, maxMatchCount);

        // 3. Transfer the immediately settled BASE tokens.
        if (buyBaseAmount != 0) {
            baseReserve -= buyBaseAmount;
            // Purchased tokens are not subject to fees.
            BASE.safeTransfer(order.owner, buyBaseAmount);
        }

        return skimQuoteAmount;
    }

    // For a SELL order, search from the most expensive price in the BUY list
    // and only execute trades where order.price is equal to or lower than the buy order price.
    function _matchSellOrder(uint256 orderId, Order memory order, uint256 maxMatchCount)
        private
        setLatest
        returns (uint256 earnQuoteAmount)
    {
        // (earnQuoteAmount) The amount of Quote earned from the sale.

        // cache storage immutables to memory
        (IERC20 base, uint256 denominator) = (BASE, DENOMINATOR);
        List.U256 storage _buyPrices = _prices[uint8(OrderSide.BUY)];
        while (!_buyPrices.empty()) {
            uint256 price = _buyPrices.at(0);
            if (price < order.price) break;
            List.U256 storage _orders = _buyOrders[price];
            _cacheLatestPrice(price);

            uint256 length = List.length(_orders);
            while (length != 0) {
                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // Update the settled quantity.
                (address targetOwner, uint256 tradeAmount,) =
                    _matchOrderAmount(orderId, order, targetId, target, price, _orders);

                // Trade executed.
                base.safeTransfer(targetOwner, tradeAmount);

                // Update information.
                earnQuoteAmount += Math.mulDiv(price, tradeAmount, denominator);
                if (order.amount == 0 || --maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // Although the `while` loop has not yet ended,
                        // if `cOrder` and the last `target.amount` in `orders` are the same,
                        // `_orders` may be empty.
                        if (!_buyPrices.remove(price)) revert PairUnknownPrices(OrderSide.BUY, price);
                    }
                    return earnQuoteAmount;
                }
                unchecked {
                    --length;
                }
            }
            // Reaching this point means that all orders at the given `price` have been matched,
            // so remove `price` from `_buyPrices`.
            if (!_buyPrices.remove(price)) revert PairUnknownPrices(OrderSide.BUY, price);
        }
        return earnQuoteAmount;
    }

    // For a BUY order, search from the cheapest price in the SELL list
    // and only execute trades where the sell order price is equal to or lower than the input price.
    function _matchBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 quoteAmount, // Quote amount to be used for Market trades.
        uint256 maxMatchCount
    ) private setLatest returns (uint256 matchedBaseAmount) {
        uint256 useQuoteAmount;

        List.U256 storage _sellPrices = _prices[uint8(OrderSide.SELL)];
        while (!_sellPrices.empty()) {
            uint256 price = _sellPrices.at(0);
            if (price > order.price) break;
            List.U256 storage _orders = _sellOrders[price];
            _cacheLatestPrice(price);

            // If it is a Market trade, calculate the maximum quantity that can be purchased at the given price.
            if (quoteAmount != 0) {
                uint256 buyAmount = Math.mulDiv(quoteAmount - useQuoteAmount, DENOMINATOR, price);
                buyAmount -= buyAmount % lotSize;
                order.amount = buyAmount;
                if (buyAmount == 0) return matchedBaseAmount;
            }

            uint256 length = List.length(_orders);
            while (length != 0) {
                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // Update the settled quantity.
                (address targetOwner, uint256 tradeAmount, uint32 targetFeeBps) =
                    _matchOrderAmount(orderId, order, targetId, target, price, _orders);
                uint256 tradeQuoteAmount = Math.mulDiv(price, tradeAmount, DENOMINATOR);

                // Trade executed. ( Calculate using the fee rate at the time the seller registered the sale.)
                _exchangeQuote(targetId, targetOwner, tradeQuoteAmount, _feeCollector(), targetFeeBps);

                // Update information.
                matchedBaseAmount += tradeAmount;
                useQuoteAmount += tradeQuoteAmount;
                if (order.amount == 0 || --maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // Although the `while` loop has not yet ended,
                        // if `cOrder` and the last `target.amount` in `orders` are the same,
                        // `_orders` may be empty.
                        if (!_sellPrices.remove(price)) revert PairUnknownPrices(OrderSide.SELL, price);
                    }
                    return matchedBaseAmount;
                }
                unchecked {
                    --length;
                }
            }
            // Reaching this point means that all orders at the given `price` have been matched,
            // so remove `price` from `_sellPrices`.
            if (!_sellPrices.remove(price)) revert PairUnknownPrices(OrderSide.SELL, price);
        }
        return matchedBaseAmount;
    }

    function _matchOrderAmount(
        uint256 orderId,
        Order memory order,
        uint256 targetId,
        Order storage target,
        uint256 price,
        List.U256 storage _orders
    ) private returns (address targetOwner, uint256 tradeAmount, uint32 targetFeeBps) {
        (targetOwner, tradeAmount, targetFeeBps) = (target.owner, Math.min(order.amount, target.amount), target.feeBps);

        (uint256 sellId, uint256 buyId) = (order.side == OrderSide.SELL ? (orderId, targetId) : (targetId, orderId));
        emit OrderMatched(sellId, buyId, price, tradeAmount, block.timestamp);

        // If the entire quantity of target is traded, remove the data.
        if (tradeAmount == target.amount) {
            _removeOrder(targetId, CloseType.COMPLETED, _orders);
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
        if (_orders.empty()) _prices[uint8(order.side)].remove(order.price);
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

    function _exchangeQuote(uint256 orderId, address _owner, uint256 amount, address feeCollector, uint32 feeBps)
        private
    {
        if (feeBps == 0) {
            QUOTE.safeTransfer(_owner, amount);
        } else {
            uint256 fee = Math.mulDiv(amount, feeBps, 10000);
            uint256 value = amount - fee;
            emit FeeCollect(orderId, _owner, amount, feeCollector, feeBps, fee, value);

            QUOTE.safeTransfer(feeCollector, fee);
            QUOTE.safeTransfer(_owner, value);
        }
    }

    function _cacheFeeInfos() private {
        IMarket market = IMarket(MARKET);
        (address feeCollector, uint32 feeBps) = (market.feeCollector(), market.feeBps());
        assembly {
            tstore(_feeCollectorSlot, feeCollector)
            tstore(_feeBpsSlot, feeBps)
        }
    }

    function _feeBps() private view returns (uint32 feeBps) {
        assembly {
            feeBps := tload(_feeBpsSlot)
        }
    }

    function _feeCollector() private view returns (address feeCollector) {
        assembly {
            feeCollector := tload(_feeCollectorSlot)
        }
    }

    function _cacheLatestPrice(uint256 price) private {
        assembly {
            tstore(_matchedPriceSlot, price)
        }
    }

    function _setLatest() private {
        uint256 _latestPrice;
        assembly {
            _latestPrice := tload(_matchedPriceSlot)
        }
        if (_latestPrice != 0) {
            matchedPrice = _latestPrice;
            matchedAt = block.timestamp;
        }
    }

    //    ##   #    # ##### #    #  ####  #####  # ######   ##   ##### #  ####  #    #
    //   #  #  #    #   #   #    # #    # #    # #     #   #  #    #   # #    # ##   #
    //  #    # #    #   #   ###### #    # #    # #    #   #    #   #   # #    # # #  #
    //  ###### #    #   #   #    # #    # #####  #   #    ######   #   # #    # #  # #
    //  #    # #    #   #   #    # #    # #   #  #  #     #    #   #   # #    # #   ##
    //  #    #  ####    #   #    #  ####  #    # # ###### #    #   #   #  ####  #    #

    function setTickSize(uint256 _lotSize, uint256 _tickSize) external {
        IMarket(MARKET).checkTickSizeRoles(_msgSender());

        if (_tickSize == 0) revert PairInvalidInitializeData("tickSize");
        if (_lotSize == 0) revert PairInvalidInitializeData("lotSize");
        if (_tickSize * _lotSize % DENOMINATOR != 0) revert PairInvalidTickSize(_tickSize, _lotSize, DENOMINATOR);
        emit TickSizeUpdated(lotSize, _lotSize, tickSize, _tickSize);

        lotSize = _lotSize;
        tickSize = _tickSize;
        minTradeVolume = Math.mulDiv(_tickSize, _lotSize, DENOMINATOR);
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

    function emergencyCancelOrder(uint256[] memory orderIds) external whenPaused onlyOwner {
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
