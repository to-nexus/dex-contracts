// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.5.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.5.0/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin-contracts-5.5.0/utils/Address.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";

import {PausableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.5.0/utils/PausableUpgradeable.sol";

import {BPS_DENOMINATOR, IFeeController} from "./interfaces/IFeeController.sol";
import {IMarketV3} from "./interfaces/IMarketV3.sol";
import {IOwnable} from "./interfaces/IOwnable.sol";
import {IPairV3} from "./interfaces/IPairV3.sol";
import {List} from "./lib/List.sol";

contract PairImplV3 is IPairV3, IOwnable, UUPSUpgradeable, PausableUpgradeable {
    using SafeERC20 for IERC20;
    using Math for uint256;
    using List for List.U256;
    using Address for address;

    error PairInvalidReserve(address);
    error PairInvalidAccountReserve(address, address);
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
    event FeeCollect(uint256 indexed orderId, address indexed owner, uint256 amount, uint256 fee, uint256 value);
    event TickSizeUpdated(uint256 beforeLotSize, uint256 newLotSize, uint256 beforeTickSize, uint256 newTickSize);
    event Skim(address indexed caller, address indexed erc20, address indexed to, uint256 amount);
    event FeeControllerUpdated(address indexed before, address indexed current);

    // slots
    // keccak256(abi.encode(uint256(keccak256("crossdex.pair.matchedprice")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant _matchedPriceSlot = 0xfd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd00;

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
    mapping(uint256 price => List.U256) private _sellOrders; // price => sell order id list (For the same price, orders will be stored in chronological order.)
    mapping(uint256 price => List.U256) private _buyOrders; //  price => buy order id list (For the same price, orders will be stored in chronological order.)
    mapping(uint256 orderId => Order) private _allOrders;
    mapping(address account => uint256[2]) private _accountReserves; // 0: sell (BASE), 1: buy (QUOTE)

    // Pair-specific fee configuration
    IFeeController public feeController;

    uint256[24] private __gap;

    modifier onlyOwner() {
        // The owner of the Pair is the owner of the Market contract that deployed this Pair.
        _checkOwner();
        _;
    }

    modifier onlyRouter() {
        if (_msgSender() != ROUTER) revert PairInvalidRouter(_msgSender());
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
        uint256 _lotSize, // lot size for base token
        address _feeController,
        bytes memory feeControllerInitData
    ) external initializer {
        __Pausable_init();

        if (router == address(0)) revert PairInvalidInitializeData("router");
        if (quote == address(0)) revert PairInvalidInitializeData("quote");
        if (base == address(0)) revert PairInvalidInitializeData("base");
        if (_tickSize == 0) revert PairInvalidInitializeData("tickSize");
        if (_lotSize == 0) revert PairInvalidInitializeData("lotSize");
        if (_feeController == address(0)) revert PairInvalidInitializeData("feeController");

        MARKET = _msgSender();
        ROUTER = router;
        QUOTE = IERC20(quote);
        BASE = IERC20(base);
        DENOMINATOR = 10 ** IERC20Metadata(base).decimals();

        if (_tickSize * _lotSize % DENOMINATOR != 0) revert PairInvalidTickSize(_tickSize, _lotSize, DENOMINATOR);

        tickSize = _tickSize;
        lotSize = _lotSize;
        minTradeVolume = Math.mulDiv(_tickSize, _lotSize, DENOMINATOR);

        feeController = IFeeController(_feeController);
        _feeControllerInitialize(feeControllerInitData);
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

    function calcBuyVolumeWithFee(uint256 volume) external returns (uint256 buyVolume) {
        return _feeControllerCalcBuyVolumeWithFee(false, volume);
    }

    function orderById(uint256 id) external view returns (Order memory) {
        return _allOrders[id];
    }

    function accountReserves(address account) external view returns (uint256 base, uint256 quote) {
        return (_accountReserves[account][uint8(OrderSide.SELL)], _accountReserves[account][uint8(OrderSide.BUY)]);
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
    ) external override whenNotPaused onlyRouter returns (uint256 orderId) {
        // Check the conditions of the entered quantity.
        if (order.price == 0 || order.price % tickSize != 0) revert PairInvalidPrice(order.price);
        if (order.amount == 0 || order.amount % lotSize != 0) revert PairInvalidAmount(order.amount);

        orderId = ++_orderIdCounter;
        (bool isSellOrder) = order.side == OrderSide.SELL;

        (bool done, uint256 mustRemainQuoteAmount) = isSellOrder
            ? _executeSellOrder(orderId, order, maxMatchCount)
            : _executeBuyOrder(orderId, order, 0, maxMatchCount);

        if (done) {
            emit OrderClosed(orderId, CloseType.COMPLETED, block.timestamp);
            // Return the remaining balance.
            // The quantity returned as a quote can occur under all conditions except FILL_OR_KILL,
            // so it is refunded collectively at the end of the function before termination.
            if (isSellOrder && order.amount != 0) BASE.safeTransfer(order.owner, order.amount);
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
                // previous price validation
                if (
                    prevPrice != 0
                        && ((isSellOrder && order.price < prevPrice) || (!isSellOrder && order.price > prevPrice))
                ) revert PairInvalidPrevPrice(order.side, order.price, prevPrice);
                // update price
                _prices[uint8(order.side)].insert(order.price, prevPrice);

                // update order
                _allOrders[orderId] = order;

                if (isSellOrder) {
                    // Set maker fee bps for V2 compatibility (used in cancel refund)
                    _allOrders[orderId].feeBps = _feeControllerSellerMakerFeeBps();
                    _addBaseReserve(order.owner, order.amount);
                    _sellOrders[order.price].push(orderId);
                } else {
                    // Set maker fee bps for V2 compatibility (used in cancel refund)
                    _allOrders[orderId].feeBps = _feeControllerBuyerMakerFeeBps();
                    // For V2 RouterV2, the fee is already included in the transferred amount
                    // So we use the actual received amount instead of calculating fee again
                    uint256 reserveQuoteAmount = _feeControllerCalcBuyVolumeWithFeeOrder(true, order);
                    _addQuoteReserve(order.owner, reserveQuoteAmount);
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
            (, uint256 mustRemainQuoteAmount) = _executeBuyOrder(orderId, order, spendAmount, maxMatchCount);
            _returnRemainQuote(order.owner, mustRemainQuoteAmount);
        }

        emit OrderClosed(orderId, CloseType.COMPLETED, block.timestamp);
    }

    function cancelOrder(address caller, uint256[] calldata orderIds) external override onlyRouter {
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
        returns (bool, uint256)
    {
        if (order.side != OrderSide.SELL) revert PairInvalidOrderSide(OrderSide.BUY);

        // 1. Verify that the required tokens for the order have been deposited.
        if (BASE.balanceOf(address(this)) < baseReserve + order.amount) revert PairInvalidReserve(address(BASE));
        emit OrderCreated(order.owner, orderId, order.side, order.price, order.amount, block.timestamp);

        // 2. If there are immediately tradable orders, execute the trade.
        //    For a SELL order, search from the most expensive price in the BUY list
        //    and only match with buy orders that have a price equal to or higher than the input price.
        (bool done, uint256 earnQuoteAmount) = _matchSellOrder(orderId, order, maxMatchCount);
        if (earnQuoteAmount != 0) {
            uint256 takerFee = _feeControllerSettleFees();
            _exchangeSellOrder(orderId, order.owner, earnQuoteAmount, takerFee);
        }
        return (done, 0);
    }

    function _executeBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 spendQuoteAmount, // Set this value if it is a Market Order.
        uint256 maxMatchCount
    )
        private
        returns (bool, uint256)
    {
        if (order.side != OrderSide.BUY) revert PairInvalidOrderSide(OrderSide.SELL);

        // 1. Verify that the required tokens for the order have been deposited.
        uint256 skimQuoteAmount;
        {
            uint256 buyVolumeWithFee = spendQuoteAmount == 0
                ? _feeControllerCalcBuyVolumeWithFeeOrder(false, order)
                : _feeControllerCalcBuyVolumeWithFee(false, spendQuoteAmount);
            bool ok;
            (ok, skimQuoteAmount) = Math.trySub(QUOTE.balanceOf(address(this)), quoteReserve + buyVolumeWithFee);
            if (!ok) revert PairInvalidReserve(address(QUOTE));
            emit OrderCreated(order.owner, orderId, order.side, order.price, order.amount, block.timestamp);
        }
        // 2. If there are immediately tradable orders, execute the trade.
        //    For a BUY order, search from the cheapest price in the SELL list
        //    and only match with sell orders that have a price equal to or lower than the input price.
        (bool done, uint256 buyBaseAmount, uint256 useQuoteAmount) =
            _matchBuyOrder(orderId, order, spendQuoteAmount, maxMatchCount);

        // 3. Transfer the immediately settled BASE tokens.
        if (buyBaseAmount != 0) {
            // Process buyer fee
            _exchangeBuyOrder(orderId, order.owner, buyBaseAmount, useQuoteAmount, 0);
            uint256 takerFee = _feeControllerSettleFees();
            if (takerFee != 0) {
                emit FeeCollect(orderId, order.owner, useQuoteAmount, takerFee, useQuoteAmount - takerFee);
            }
        }

        return (done, skimQuoteAmount);
    }

    struct MatchSellCache {
        uint256 price;
        uint256 earnQuoteAmount;
        uint256 quoteReserve;
        bool done;
    }

    // For a SELL order, search from the most expensive price in the BUY list
    // and only execute trades where order.price is equal to or lower than the buy order price.
    /// @return done Whether the order has been completely matched or the maxMatchCount has reached 0.
    /// @return earnQuoteAmount The amount of QUOTE earned from the sale.
    function _matchSellOrder(uint256 orderId, Order memory order, uint256 maxMatchCount)
        private
        setLatest
        returns (bool, uint256)
    {
        MatchSellCache memory cache =
            MatchSellCache({price: 0, earnQuoteAmount: 0, quoteReserve: quoteReserve, done: false});

        // cache storage immutables to memory
        List.U256 storage _buyPrices = _prices[uint8(OrderSide.BUY)];
        while (!_buyPrices.empty()) {
            cache.price = _buyPrices.at(0);
            if (cache.price < order.price) break;
            List.U256 storage _orders = _buyOrders[cache.price];
            _cacheLatestPrice(cache.price);

            while (List.length(_orders) != 0) {
                uint256 makerId = _orders.at(0);
                Order storage maker = _allOrders[makerId];

                // Update the settled quantity.
                // For sell order matching buy order: sell order is taker, buy order is maker
                (address makerOwner, uint256 tradeAmount, uint256 makerFee) =
                    _matchOrderAmount(orderId, order, makerId, maker, cache.price, _orders);
                uint256 tradeQuoteAmount = Math.mulDiv(cache.price, tradeAmount, DENOMINATOR);

                // Trade executed.
                _exchangeBuyOrder(makerId, makerOwner, tradeAmount, tradeQuoteAmount, makerFee);

                // Update information.
                cache.earnQuoteAmount += tradeQuoteAmount;
                cache.quoteReserve = _subQuoteReserve(makerOwner, tradeQuoteAmount + makerFee, true, cache.quoteReserve);
                if (order.amount == 0 || --maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // Although the `while` loop has not yet ended,
                        // if `cOrder` and the last `target.amount` in `orders` are the same,
                        // `_orders` may be empty.
                        if (!_buyPrices.remove(cache.price)) revert PairUnknownPrices(OrderSide.BUY, cache.price);
                    }
                    cache.done = true;
                    break;
                }
            }
            if (cache.done) break;
            // Reaching this point means that all orders at the given `price` have been matched,
            // so remove `price` from `_buyPrices`.
            if (!_buyPrices.remove(cache.price)) revert PairUnknownPrices(OrderSide.BUY, cache.price);
        }
        if (cache.quoteReserve != quoteReserve) quoteReserve = cache.quoteReserve;
        return (cache.done, cache.earnQuoteAmount);
    }

    // avoid stack too deep `_matchBuyOrder function`
    struct MatchBuyCache {
        uint256 price;
        uint256 matchedBaseAmount;
        uint256 useQuoteAmount;
        uint256 baseReserve;
        uint256 totalTakerFee;
        bool done;
    }

    // For a BUY order, search from the cheapest price in the SELL list
    // and only execute trades where the sell order price is equal to or lower than the input price.
    /// @return done Whether the order has been completely matched or the maxMatchCount has reached 0.
    /// @return matchedBaseAmount The amount of BASE purchased.
    /// @return useQuoteAmount The amount of QUOTE used for the purchase.
    function _matchBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 quoteAmount, // Quote amount to be used for Market trades.
        uint256 maxMatchCount
    )
        private
        setLatest
        returns (bool, uint256, uint256)
    {
        MatchBuyCache memory cache = MatchBuyCache({
            price: 0, matchedBaseAmount: 0, useQuoteAmount: 0, baseReserve: baseReserve, totalTakerFee: 0, done: false
        });

        List.U256 storage _sellPrices = _prices[uint8(OrderSide.SELL)];
        while (!_sellPrices.empty()) {
            cache.price = _sellPrices.at(0);
            if (cache.price > order.price) break;
            List.U256 storage _orders = _sellOrders[cache.price];
            _cacheLatestPrice(cache.price);

            // If it is a Market trade, calculate the maximum quantity that can be purchased at the given price.
            if (quoteAmount != 0) {
                uint256 buyAmount = Math.mulDiv(quoteAmount - cache.useQuoteAmount, DENOMINATOR, cache.price);
                buyAmount -= buyAmount % lotSize;
                order.amount = buyAmount;
                if (buyAmount == 0) {
                    cache.done = true;
                    break;
                }
            }

            while (List.length(_orders) != 0) {
                uint256 makerId = _orders.at(0);
                Order storage maker = _allOrders[makerId];

                // Update the settled quantity.
                // For buy order matching sell order: buy order is taker, sell order is maker
                (address makerOwner, uint256 tradeAmount, uint256 makerFee) =
                    _matchOrderAmount(orderId, order, makerId, maker, cache.price, _orders);
                uint256 tradeQuoteAmount = Math.mulDiv(cache.price, tradeAmount, DENOMINATOR);
                // Trade executed. ( Calculate using the fee rate at the time the seller registered the sale.)
                _exchangeSellOrder(makerId, makerOwner, tradeQuoteAmount, makerFee);

                // Update information.
                cache.matchedBaseAmount += tradeAmount;
                cache.useQuoteAmount += tradeQuoteAmount;
                cache.baseReserve = _subBaseReserve(makerOwner, tradeAmount, true, cache.baseReserve);
                if (order.amount == 0 || --maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // Although the `while` loop has not yet ended,
                        // if `cOrder` and the last `target.amount` in `orders` are the same,
                        // `_orders` may be empty.
                        if (!_sellPrices.remove(cache.price)) revert PairUnknownPrices(OrderSide.SELL, cache.price);
                    }
                    cache.done = true;
                    break;
                }
            }
            if (cache.done) break;
            // Reaching this point means that all orders at the given `price` have been matched,
            // so remove `price` from `_sellPrices`.
            if (!_sellPrices.remove(cache.price)) revert PairUnknownPrices(OrderSide.SELL, cache.price);
        }
        if (cache.baseReserve != baseReserve) baseReserve = cache.baseReserve;
        return (cache.done, cache.matchedBaseAmount, cache.useQuoteAmount);
    }

    function _matchOrderAmount(
        uint256 takerId,
        Order memory taker,
        uint256 makerId,
        Order storage maker,
        uint256 price,
        List.U256 storage _orders
    ) private returns (address makerOwner, uint256 tradeAmount, uint256 makerFee) {
        tradeAmount = Math.min(taker.amount, maker.amount);
        uint256 tradeQuoteAmount = Math.mulDiv(price, tradeAmount, DENOMINATOR);
        (makerOwner, makerFee) =
        (maker.owner, _feeControllerRecodeMatch(takerId, taker, maker, tradeAmount, tradeQuoteAmount));

        (uint256 sellId, uint256 buyId) = (taker.side == OrderSide.SELL ? (takerId, makerId) : (makerId, takerId));
        emit OrderMatched(sellId, buyId, price, tradeAmount, block.timestamp);

        // If the entire quantity of target is traded, remove the data.
        if (tradeAmount == maker.amount) {
            _removeOrder(makerId, CloseType.COMPLETED, _orders);
        } else {
            unchecked {
                maker.amount -= tradeAmount;
            }
        }

        // If the entire quantity of order is traded, trigger termination.
        if (tradeAmount == taker.amount) {
            taker.amount = 0;
        } else {
            unchecked {
                taker.amount -= tradeAmount;
            }
        }
    }

    function _cancelOrder(uint256 orderId, Order memory order, CloseType _type) private {
        List.U256 storage _orders;
        bool isSellOrder = order.side == OrderSide.SELL;

        // Return the tokens held by the contract.
        if (isSellOrder) {
            _orders = _sellOrders[order.price];
            uint256 amount = order.amount;
            if (amount != 0) {
                BASE.safeTransfer(order.owner, amount);
                _subBaseReserve(order.owner, amount, false, baseReserve);
            }
        } else {
            _orders = _buyOrders[order.price];
            uint256 returnQuoteAmount = Math.mulDiv(order.price, order.amount, DENOMINATOR);
            if (returnQuoteAmount != 0) {
                if (order.feeBps != 0) {
                    returnQuoteAmount += Math.mulDiv(returnQuoteAmount, order.feeBps, BPS_DENOMINATOR);
                }
                QUOTE.safeTransfer(order.owner, returnQuoteAmount);
                _subQuoteReserve(order.owner, returnQuoteAmount, false, quoteReserve);
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

    // LimitBuy
    function _addQuoteReserve(address account, uint256 amount) private {
        (bool ok, uint256 newReserve) = Math.tryAdd(quoteReserve, amount);
        if (!ok) revert PairInvalidReserve(address(QUOTE));
        quoteReserve = newReserve;

        (ok, newReserve) = Math.tryAdd(_accountReserves[account][1], amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(QUOTE));
        _accountReserves[account][1] = newReserve;
    }

    // For LimitBuy order cancel or Matched From Sell
    function _subQuoteReserve(address account, uint256 amount, bool inMatching, uint256 reserve)
        private
        returns (uint256 newQuoteReserve)
    {
        (bool ok, uint256 newReserve) = Math.trySub(_accountReserves[account][1], amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(QUOTE));
        _accountReserves[account][1] = newReserve;

        (ok, newQuoteReserve) = Math.trySub(reserve, amount);
        if (!ok) revert PairInvalidReserve(address(QUOTE));
        if (!inMatching) quoteReserve = newQuoteReserve;
    }

    // LimitSell
    function _addBaseReserve(address account, uint256 amount) private {
        (bool ok, uint256 newReserve) = Math.tryAdd(baseReserve, amount);
        if (!ok) revert PairInvalidReserve(address(BASE));
        baseReserve = newReserve;

        (ok, newReserve) = Math.tryAdd(_accountReserves[account][0], amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(BASE));
        _accountReserves[account][0] = newReserve;
    }

    // For LimitSell order cancel or Matched From Buy
    function _subBaseReserve(address account, uint256 amount, bool inMatching, uint256 reserve)
        private
        returns (uint256 newBaseReserve)
    {
        (bool ok, uint256 newReserve) = Math.trySub(_accountReserves[account][0], amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(BASE));
        _accountReserves[account][0] = newReserve;

        (ok, newBaseReserve) = Math.trySub(reserve, amount);
        if (!ok) revert PairInvalidReserve(address(BASE));
        if (!inMatching) baseReserve = newBaseReserve;
    }

    function _returnRemainQuote(address to, uint256 mustRemainQuoteAmount) private {
        uint256 currentBalance = QUOTE.balanceOf(address(this));
        uint256 requiredAmount = mustRemainQuoteAmount + quoteReserve;

        if (currentBalance > requiredAmount) {
            uint256 remainQuoteAmount = currentBalance - requiredAmount;
            QUOTE.safeTransfer(to, remainQuoteAmount);
        }
    }

    function _exchangeSellOrder(uint256 orderId, address _owner, uint256 amount, uint256 fee) private {
        if (fee == 0) {
            QUOTE.safeTransfer(_owner, amount);
        } else {
            uint256 value = amount - fee;
            emit FeeCollect(orderId, _owner, amount, fee, value);
            QUOTE.safeTransfer(_owner, value);
        }
    }

    function _exchangeBuyOrder(
        uint256 orderId,
        address _owner,
        uint256 buyBaseAmount,
        uint256 useQuoteAmount,
        uint256 fee
    ) private {
        BASE.safeTransfer(_owner, buyBaseAmount);
        if (fee > 0) emit FeeCollect(orderId, _owner, useQuoteAmount, fee, useQuoteAmount - fee);
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
            if (_latestPrice != matchedPrice) matchedPrice = _latestPrice;
            if (matchedAt != block.timestamp) matchedAt = block.timestamp;
        }
    }

    function _checkOwner() private view {
        if (_msgSender() != IOwnable(MARKET).owner()) revert OwnableUnauthorizedAccount(_msgSender());
    }

    //    ##   #    # ##### #    #  ####  #####  # ######   ##   ##### #  ####  #    #
    //   #  #  #    #   #   #    # #    # #    # #     #   #  #    #   # #    # ##   #
    //  #    # #    #   #   ###### #    # #    # #    #   #    #   #   # #    # # #  #
    //  ###### #    #   #   #    # #    # #####  #   #    ######   #   # #    # #  # #
    //  #    # #    #   #   #    # #    # #   #  #  #     #    #   #   # #    # #   ##
    //  #    #  ####    #   #    #  ####  #    # # ###### #    #   #   #  ####  #    #

    function setTickSize(uint256 _lotSize, uint256 _tickSize) external {
        IMarketV3(MARKET).checkTickSizeRoles(_msgSender());

        if (_tickSize == 0) revert PairInvalidInitializeData("tickSize");
        if (_lotSize == 0) revert PairInvalidInitializeData("lotSize");
        if (_tickSize * _lotSize % DENOMINATOR != 0) revert PairInvalidTickSize(_tickSize, _lotSize, DENOMINATOR);
        emit TickSizeUpdated(lotSize, _lotSize, tickSize, _tickSize);

        lotSize = _lotSize;
        tickSize = _tickSize;
        minTradeVolume = Math.mulDiv(_tickSize, _lotSize, DENOMINATOR);
    }

    function setFeeController(address newFeeController, bytes memory feeControllerInitData) external {
        if (_msgSender() != MARKET) _checkOwner(); // only owner
        IMarketV3(MARKET).checkFeeControllerAllowed(newFeeController);

        if (address(feeController) != newFeeController) {
            emit FeeControllerUpdated(address(feeController), newFeeController);
            feeController = IFeeController(newFeeController);
        }
        _feeControllerInitialize(feeControllerInitData);
    }

    function skim(IERC20 erc20, address to, uint256 amount) external onlyOwner {
        if (amount == 0) return;

        if (erc20 == BASE && BASE.balanceOf(address(this)) < baseReserve + amount) {
            revert PairInvalidReserve(address(BASE));
        } else if (erc20 == QUOTE && QUOTE.balanceOf(address(this)) < quoteReserve + amount) {
            revert PairInvalidReserve(address(QUOTE));
        }

        erc20.safeTransfer(to, amount);
        emit Skim(_msgSender(), address(erc20), to, amount);
    }

    function emergencyCancelOrder(uint256[] calldata orderIds) external whenPaused onlyOwner {
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

    // ─────────────────────────────────────────────────────────────────────────────
    // FeeController delegatecall wrappers
    // ─────────────────────────────────────────────────────────────────────────────

    function _feeControllerInitialize(bytes memory initData) private {
        Address.functionDelegateCall(
            address(feeController), abi.encodeCall(IFeeController.initialize, (address(QUOTE), DENOMINATOR, initData))
        );
    }

    function _feeControllerCalcBuyVolumeWithFee(bool isMaker, uint256 volume) private returns (uint256) {
        // calcBuyVolumeWithFee(bool,uint256)
        bytes memory result = Address.functionDelegateCall(
            address(feeController), abi.encodeCall(IFeeController.calcBuyVolumeWithFee, (isMaker, volume))
        );
        return abi.decode(result, (uint256));
    }

    function _feeControllerCalcBuyVolumeWithFeeOrder(bool isMaker, Order memory order) private returns (uint256) {
        // calcBuyVolumeWithFee(bool,(uint8,address,uint32,uint256,uint256)) - Order struct
        bytes memory result = Address.functionDelegateCall(
            address(feeController), abi.encodeCall(IFeeController.calcBuyVolumeWithFeeByOrder, (isMaker, order))
        );
        return abi.decode(result, (uint256));
    }

    function _feeControllerRecodeMatch(
        uint256 takerId,
        Order memory taker,
        Order memory maker,
        uint256 tradeAmount,
        uint256 tradeQuoteAmount
    ) private returns (uint256) {
        bytes memory result = Address.functionDelegateCall(
            address(feeController),
            abi.encodeCall(IFeeController.recodeMatch, (takerId, taker, maker, tradeAmount, tradeQuoteAmount))
        );
        return abi.decode(result, (uint256));
    }

    function _feeControllerSettleFees() private returns (uint256) {
        bytes memory result =
            Address.functionDelegateCall(address(feeController), abi.encodeCall(IFeeController.settleFees, ()));
        return abi.decode(result, (uint256));
    }

    function _feeControllerSellerMakerFeeBps() private returns (uint32) {
        bytes memory result =
            Address.functionDelegateCall(address(feeController), abi.encodeCall(IFeeController.sellerMakerFeeBps, ()));
        return abi.decode(result, (uint32));
    }

    function _feeControllerBuyerMakerFeeBps() private returns (uint32) {
        bytes memory result =
            Address.functionDelegateCall(address(feeController), abi.encodeCall(IFeeController.buyerMakerFeeBps, ()));
        return abi.decode(result, (uint32));
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
