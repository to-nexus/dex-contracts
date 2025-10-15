// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";

import {PausableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/utils/PausableUpgradeable.sol";

import {BPS_DENOMINATOR, IMarketV2, NO_FEE_BPS} from "./interfaces/IMarket.sol";
import {IOwnable} from "./interfaces/IOwnable.sol";
import {IPairV2} from "./interfaces/IPair.sol";
import {List} from "./lib/List.sol";

contract PairImplV2 is IPairV2, IOwnable, UUPSUpgradeable, PausableUpgradeable {
    using SafeERC20 for IERC20;
    using Math for uint256;
    using List for List.U256;

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
    error PairInvalidFeeBps();
    error PairInvalidFeeStructure(uint32 makerFee, uint32 takerFee);

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
    event PairFeesUpdated(uint32 sellerMakerFee, uint32 sellerTakerFee, uint32 buyerMakerFee, uint32 buyerTakerFee);
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
    mapping(uint256 price => List.U256) private _sellOrders; // price => sell order id list (For the same price, orders will be stored in chronological order.)
    mapping(uint256 price => List.U256) private _buyOrders; //  price => buy order id list (For the same price, orders will be stored in chronological order.)
    mapping(uint256 orderId => Order) private _allOrders;
    mapping(address account => uint256[2]) private _accountReserves; // 0: sell (BASE), 1: buy (QUOTE)

    // Pair-specific fee configuration
    FeeConfig public feeConfig;

    uint256[31] private __gap;

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
        uint256 _lotSize, // lot size for base token
        bytes memory feeData // 4개 수수료를 인코딩한 데이터
    ) external initializer {
        __Pausable_init();

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

        // 4가지 수수료 디코딩
        (uint32 sellerMakerFeeBps_, uint32 sellerTakerFeeBps_, uint32 buyerMakerFeeBps_, uint32 buyerTakerFeeBps_) =
            abi.decode(feeData, (uint32, uint32, uint32, uint32));
        _setFeeBps(sellerMakerFeeBps_, sellerTakerFeeBps_, buyerMakerFeeBps_, buyerTakerFeeBps_);
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

    function getEffectiveFees()
        external
        view
        returns (uint32 sellerMakerFeeBps, uint32 sellerTakerFeeBps, uint32 buyerMakerFeeBps, uint32 buyerTakerFeeBps)
    {
        FeeConfig memory feeInfos = _resolveEffectiveFees();
        sellerMakerFeeBps = feeInfos.sellerMakerFeeBps;
        sellerTakerFeeBps = feeInfos.sellerTakerFeeBps;
        buyerMakerFeeBps = feeInfos.buyerMakerFeeBps;
        buyerTakerFeeBps = feeInfos.buyerTakerFeeBps;
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
                    _addBaseReserve(order.owner, order.amount);
                    // For sell orders, a fee is charged when acting as a maker.
                    _allOrders[orderId].feeBps = _sellerMakerFeeBps();
                    _sellOrders[order.price].push(orderId);
                } else {
                    uint32 feeBps = _buyerMakerFeeBps();
                    // For V2 RouterV2, the fee is already included in the transferred amount
                    // So we use the actual received amount instead of calculating fee again
                    uint256 receivedAmount = QUOTE.balanceOf(address(this)) - mustRemainQuoteAmount;
                    _addQuoteReserve(order.owner, receivedAmount);
                    _allOrders[orderId].feeBps = feeBps;
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

        // 3. Immediately transfer the proceeds from the trade to the seller.
        if (earnQuoteAmount != 0) {
            (address feeCollector, uint32 feeBps) = (_feeCollector(), _sellerTakerFeeBps());
            // quoteReserve -= earnQuoteAmount; // todo
            uint256 fee = _exchangeSellOrder(orderId, order.owner, earnQuoteAmount, feeCollector, feeBps);
            if (fee != 0) QUOTE.safeTransfer(feeCollector, fee);
        }

        return (done, 0);
    }

    function _executeBuyOrder(
        uint256 orderId,
        Order memory order,
        uint256 spendQuoteAmount, // Set this value if it is a Market Order.
        uint256 maxMatchCount
    ) private returns (bool, uint256) {
        if (order.side != OrderSide.BUY) revert PairInvalidOrderSide(OrderSide.SELL);

        // Verify the conditions of the entered quantity.
        uint256 quoteAmount;
        if (spendQuoteAmount != 0) quoteAmount = spendQuoteAmount;
        else quoteAmount = Math.mulDiv(order.price, order.amount, DENOMINATOR);

        // 1. Verify that the required tokens for the order have been deposited.
        uint256 skimQuoteAmount;
        {
            bool ok;

            // For limit orders, calculate fee separately since RouterV2 includes it
            uint32 feeBps = _buyerTakerFeeBps(); // Always use taker fee since even limit orders can be immediately matched
            uint256 fee = feeBps != 0 ? Math.mulDiv(quoteAmount, feeBps, BPS_DENOMINATOR) : 0;

            uint256 quoteBalance = QUOTE.balanceOf(address(this));
            (ok, skimQuoteAmount) = Math.trySub(quoteBalance, quoteReserve + quoteAmount + fee);
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
            uint256 fee = _exchangeBuyOrder(
                orderId, order.owner, buyBaseAmount, useQuoteAmount, _feeCollector(), _buyerTakerFeeBps()
            );
            if (fee != 0) QUOTE.safeTransfer(_feeCollector(), fee);

            // maker 수수료가 다른 경우 ( == 더 싼 경우) 해당 토큰을 skimQuoteAmount에서 차감한다. (트랜잭션이 종료될때 돌려준다.)
            // Market orders are always taker, so no fee adjustment needed
            if (spendQuoteAmount == 0 && _buyerMakerFeeBps() != _buyerTakerFeeBps()) {
                // maker 의 수수료가 더 싸야 하기 때문에 항상 성공해야 하며, 실패한 경우는 수수료율이 잘못 설정된 경우이다.
                // MarketImpl 의 수수료도 사용하고 PairImpl 의 수수료도 사용하는경우 발생할 수 있다.
                (bool ok, uint256 diff) =
                    Math.trySub(fee, Math.mulDiv(useQuoteAmount, _buyerMakerFeeBps(), BPS_DENOMINATOR));
                if (!ok) revert PairInvalidFeeBps();

                (ok, skimQuoteAmount) = Math.trySub(skimQuoteAmount, diff);
                if (!ok) revert PairInvalidReserve(address(QUOTE));
            }
        }

        return (done, skimQuoteAmount);
    }

    struct MatchSellCache {
        uint256 price;
        uint256 earnQuoteAmount;
        uint256 totalTargetFee;
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
        MatchSellCache memory cache = MatchSellCache({price: 0, earnQuoteAmount: 0, totalTargetFee: 0});
        // cache storage immutables to memory
        List.U256 storage _buyPrices = _prices[uint8(OrderSide.BUY)];
        while (!_buyPrices.empty()) {
            cache.price = _buyPrices.at(0);
            if (cache.price < order.price) break;
            List.U256 storage _orders = _buyOrders[cache.price];
            _cacheLatestPrice(cache.price);

            while (List.length(_orders) != 0) {
                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // Update the settled quantity.
                // For sell order matching buy order: sell order is taker, buy order is maker
                (address targetOwner, uint256 tradeAmount, uint32 targetFeeBps) =
                    _matchOrderAmount(orderId, order, targetId, target, cache.price, _orders);
                uint256 tradeQuoteAmount = Math.mulDiv(cache.price, tradeAmount, DENOMINATOR);

                // Trade executed.
                uint256 fee = _exchangeBuyOrder(
                    targetId, targetOwner, tradeAmount, tradeQuoteAmount, _feeCollector(), targetFeeBps
                );
                cache.totalTargetFee += fee;

                // Update information.
                cache.earnQuoteAmount += tradeQuoteAmount;
                _subQuoteReserve(targetOwner, tradeQuoteAmount + fee);
                if (order.amount == 0 || --maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // Although the `while` loop has not yet ended,
                        // if `cOrder` and the last `target.amount` in `orders` are the same,
                        // `_orders` may be empty.
                        if (!_buyPrices.remove(cache.price)) revert PairUnknownPrices(OrderSide.BUY, cache.price);
                    }
                    if (cache.totalTargetFee != 0) QUOTE.safeTransfer(_feeCollector(), cache.totalTargetFee);
                    return (true, cache.earnQuoteAmount);
                }
            }
            // Reaching this point means that all orders at the given `price` have been matched,
            // so remove `price` from `_buyPrices`.
            if (!_buyPrices.remove(cache.price)) revert PairUnknownPrices(OrderSide.BUY, cache.price);
        }
        if (cache.totalTargetFee != 0) QUOTE.safeTransfer(_feeCollector(), cache.totalTargetFee);
        return (false, cache.earnQuoteAmount);
    }

    // avoid stack too deep `_matchBuyOrder function`
    struct MatchBuyCache {
        uint256 price;
        uint256 matchedBaseAmount;
        uint256 useQuoteAmount;
        uint256 totalFee;
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
    ) private setLatest returns (bool, uint256, uint256) {
        MatchBuyCache memory cache = MatchBuyCache({price: 0, matchedBaseAmount: 0, useQuoteAmount: 0, totalFee: 0});

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
                    if (cache.totalFee != 0) QUOTE.safeTransfer(_feeCollector(), cache.totalFee);
                    return (true, cache.matchedBaseAmount, cache.useQuoteAmount);
                }
            }

            while (List.length(_orders) != 0) {
                uint256 targetId = _orders.at(0);
                Order storage target = _allOrders[targetId];

                // Update the settled quantity.
                // For buy order matching sell order: buy order is taker, sell order is maker
                (address targetOwner, uint256 tradeAmount, uint32 targetFeeBps) =
                    _matchOrderAmount(orderId, order, targetId, target, cache.price, _orders);
                uint256 tradeQuoteAmount = Math.mulDiv(cache.price, tradeAmount, DENOMINATOR);

                // Trade executed. ( Calculate using the fee rate at the time the seller registered the sale.)
                cache.totalFee +=
                    _exchangeSellOrder(targetId, targetOwner, tradeQuoteAmount, _feeCollector(), targetFeeBps);

                // Update information.
                cache.matchedBaseAmount += tradeAmount;
                cache.useQuoteAmount += tradeQuoteAmount;
                _subBaseReserve(targetOwner, tradeAmount, targetFeeBps, tradeQuoteAmount);
                if (order.amount == 0 || --maxMatchCount == 0) {
                    if (_orders.empty()) {
                        // Although the `while` loop has not yet ended,
                        // if `cOrder` and the last `target.amount` in `orders` are the same,
                        // `_orders` may be empty.
                        if (!_sellPrices.remove(cache.price)) revert PairUnknownPrices(OrderSide.SELL, cache.price);
                    }
                    if (cache.totalFee != 0) QUOTE.safeTransfer(_feeCollector(), cache.totalFee);
                    return (true, cache.matchedBaseAmount, cache.useQuoteAmount);
                }
            }
            // Reaching this point means that all orders at the given `price` have been matched,
            // so remove `price` from `_sellPrices`.
            if (!_sellPrices.remove(cache.price)) revert PairUnknownPrices(OrderSide.SELL, cache.price);
        }
        if (cache.totalFee != 0) QUOTE.safeTransfer(_feeCollector(), cache.totalFee);
        return (false, cache.matchedBaseAmount, cache.useQuoteAmount);
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
            uint256 amount = order.amount;
            if (amount != 0) {
                BASE.safeTransfer(order.owner, amount);
                _subBaseReserve(order.owner, amount, order.feeBps, 0); // Base 의 reserve 만 제거하면 되므로 거래 수량은 0
            }
        } else {
            _orders = _buyOrders[order.price];
            uint256 returnQuoteAmount = Math.mulDiv(order.price, order.amount, DENOMINATOR);
            if (returnQuoteAmount != 0) {
                if (order.feeBps != 0) {
                    returnQuoteAmount += Math.mulDiv(returnQuoteAmount, order.feeBps, BPS_DENOMINATOR);
                }
                QUOTE.safeTransfer(order.owner, returnQuoteAmount);
                _subQuoteReserve(order.owner, returnQuoteAmount);
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
        if (!ok) revert PairInvalidAccountReserve(account, address(QUOTE));
        quoteReserve = newReserve;

        (ok, newReserve) = Math.tryAdd(_accountReserves[account][1], amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(QUOTE));
        _accountReserves[account][1] = newReserve;
    }

    // For LimitBuy order cancel or Matched From Sell
    // todo: in matching, after update quoteReserve
    function _subQuoteReserve(address account, uint256 amount) private {
        (bool ok, uint256 newReserve) = Math.trySub(quoteReserve, amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(QUOTE));
        quoteReserve = newReserve;

        (ok, newReserve) = Math.trySub(_accountReserves[account][1], amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(QUOTE));
        _accountReserves[account][1] = newReserve;
    }

    // LimitSell
    function _addBaseReserve(address account, uint256 amount) private {
        (bool ok, uint256 newReserve) = Math.tryAdd(baseReserve, amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(BASE));
        baseReserve = newReserve;

        (ok, newReserve) = Math.tryAdd(_accountReserves[account][0], amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(BASE));
        _accountReserves[account][0] = newReserve;
    }

    // For LimitSell order cancel or Matched From Buy
    /// @return remaining quote fee
    // todo: in matching, after update quoteReserve
    function _subBaseReserve(address account, uint256 amount, uint32 feeBps, uint256 tradeVolume)
        private
        returns (uint256)
    {
        (bool ok, uint256 newReserve) = Math.trySub(baseReserve, amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(BASE));
        baseReserve = newReserve;

        (ok, newReserve) = Math.trySub(_accountReserves[account][0], amount);
        if (!ok) revert PairInvalidAccountReserve(account, address(BASE));
        _accountReserves[account][0] = newReserve;

        if (feeBps != 0 && tradeVolume != 0) {
            uint256 fee = Math.mulDiv(tradeVolume, feeBps, BPS_DENOMINATOR);
            _subQuoteReserve(account, fee);
            return fee;
        }
        return 0;
    }

    function _returnRemainQuote(address to, uint256 mustRemainQuoteAmount) private {
        uint256 currentBalance = QUOTE.balanceOf(address(this));
        uint256 requiredAmount = mustRemainQuoteAmount + quoteReserve;

        if (currentBalance > requiredAmount) {
            uint256 remainQuoteAmount = currentBalance - requiredAmount;
            QUOTE.safeTransfer(to, remainQuoteAmount);
        }
    }

    function _exchangeSellOrder(uint256 orderId, address _owner, uint256 amount, address feeCollector, uint32 feeBps)
        private
        returns (uint256)
    {
        if (feeBps == 0) {
            QUOTE.safeTransfer(_owner, amount);
            return 0;
        } else {
            uint256 fee = Math.mulDiv(amount, feeBps, BPS_DENOMINATOR);
            uint256 value = amount - fee;
            emit FeeCollect(orderId, _owner, amount, feeCollector, feeBps, fee, value);

            QUOTE.safeTransfer(_owner, value);
            return fee;
        }
    }

    function _exchangeBuyOrder(
        uint256 orderId,
        address _owner,
        uint256 buyBaseAmount,
        uint256 useQuoteAmount,
        address feeCollector,
        uint32 feeBps
    ) private returns (uint256) {
        BASE.safeTransfer(_owner, buyBaseAmount);

        if (feeBps == 0) {
            return 0;
        } else {
            uint256 fee = Math.mulDiv(useQuoteAmount, feeBps, BPS_DENOMINATOR);
            emit FeeCollect(orderId, _owner, useQuoteAmount, feeCollector, feeBps, fee, useQuoteAmount - fee);
            return fee;
        }
    }

    function _cacheFeeInfos() private {
        IMarketV2 market = IMarketV2(MARKET);
        address feeCollector = market.feeCollector();
        FeeConfig memory feeInfos = _resolveEffectiveFees();

        assembly {
            tstore(_feeCollectorSlot, feeCollector)
            tstore(_feeBpsSlot, feeInfos)
        }
    }

    function _resolveEffectiveFees() private view returns (FeeConfig memory feeInfos) {
        IMarketV2.FeeConfig memory defaultFeeBps = IMarketV2(MARKET).getFeeConfig();
        feeInfos.sellerMakerFeeBps =
            feeConfig.sellerMakerFeeBps == NO_FEE_BPS ? defaultFeeBps.sellerMakerFeeBps : feeConfig.sellerMakerFeeBps;
        feeInfos.sellerTakerFeeBps =
            feeConfig.sellerTakerFeeBps == NO_FEE_BPS ? defaultFeeBps.sellerTakerFeeBps : feeConfig.sellerTakerFeeBps;
        feeInfos.buyerMakerFeeBps =
            feeConfig.buyerMakerFeeBps == NO_FEE_BPS ? defaultFeeBps.buyerMakerFeeBps : feeConfig.buyerMakerFeeBps;
        feeInfos.buyerTakerFeeBps =
            feeConfig.buyerTakerFeeBps == NO_FEE_BPS ? defaultFeeBps.buyerTakerFeeBps : feeConfig.buyerTakerFeeBps;
    }

    function _feeCollector() private view returns (address feeCollector) {
        assembly {
            feeCollector := tload(_feeCollectorSlot)
        }
    }

    function _sellerMakerFeeBps() private view returns (uint32 feeBps) {
        FeeConfig memory feeInfos;
        assembly {
            feeInfos := tload(_feeBpsSlot)
        }
        feeBps = feeInfos.sellerMakerFeeBps;
    }

    function _sellerTakerFeeBps() private view returns (uint32 feeBps) {
        FeeConfig memory feeInfos;
        assembly {
            feeInfos := tload(_feeBpsSlot)
        }
        feeBps = feeInfos.sellerTakerFeeBps;
    }

    function _buyerMakerFeeBps() private view returns (uint32 feeBps) {
        FeeConfig memory feeInfos;
        assembly {
            feeInfos := tload(_feeBpsSlot)
        }
        feeBps = feeInfos.buyerMakerFeeBps;
    }

    function _buyerTakerFeeBps() private view returns (uint32 feeBps) {
        FeeConfig memory feeInfos;
        assembly {
            feeInfos := tload(_feeBpsSlot)
        }
        feeBps = feeInfos.buyerTakerFeeBps;
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
        if (_latestPrice != 0 && _latestPrice != matchedPrice) matchedPrice = _latestPrice;
        if (matchedAt != block.timestamp) matchedAt = block.timestamp;
    }

    function _setFeeBps(
        uint32 sellerMakerFeeBps_,
        uint32 sellerTakerFeeBps_,
        uint32 buyerMakerFeeBps_,
        uint32 buyerTakerFeeBps_
    ) private {
        // range check
        if (sellerMakerFeeBps_ >= BPS_DENOMINATOR && sellerMakerFeeBps_ != NO_FEE_BPS) revert PairInvalidFeeBps();
        if (sellerTakerFeeBps_ >= BPS_DENOMINATOR && sellerTakerFeeBps_ != NO_FEE_BPS) revert PairInvalidFeeBps();
        if (buyerMakerFeeBps_ >= BPS_DENOMINATOR && buyerMakerFeeBps_ != NO_FEE_BPS) revert PairInvalidFeeBps();
        if (buyerTakerFeeBps_ >= BPS_DENOMINATOR && buyerTakerFeeBps_ != NO_FEE_BPS) revert PairInvalidFeeBps();

        // logical check - taker fee must be >= maker fee
        if (sellerTakerFeeBps_ < sellerMakerFeeBps_ && sellerMakerFeeBps_ != NO_FEE_BPS) {
            revert PairInvalidFeeStructure(sellerMakerFeeBps_, sellerTakerFeeBps_);
        }
        if (buyerTakerFeeBps_ < buyerMakerFeeBps_ && buyerMakerFeeBps_ != NO_FEE_BPS) {
            revert PairInvalidFeeStructure(buyerMakerFeeBps_, buyerTakerFeeBps_);
        }

        // set
        feeConfig = FeeConfig({
            sellerMakerFeeBps: sellerMakerFeeBps_,
            sellerTakerFeeBps: sellerTakerFeeBps_,
            buyerMakerFeeBps: buyerMakerFeeBps_,
            buyerTakerFeeBps: buyerTakerFeeBps_
        });
        emit PairFeesUpdated(sellerMakerFeeBps_, sellerTakerFeeBps_, buyerMakerFeeBps_, buyerTakerFeeBps_);
    }

    //    ##   #    # ##### #    #  ####  #####  # ######   ##   ##### #  ####  #    #
    //   #  #  #    #   #   #    # #    # #    # #     #   #  #    #   # #    # ##   #
    //  #    # #    #   #   ###### #    # #    # #    #   #    #   #   # #    # # #  #
    //  ###### #    #   #   #    # #    # #####  #   #    ######   #   # #    # #  # #
    //  #    # #    #   #   #    # #    # #   #  #  #     #    #   #   # #    # #   ##
    //  #    #  ####    #   #    #  ####  #    # # ###### #    #   #   #  ####  #    #

    function setTickSize(uint256 _lotSize, uint256 _tickSize) external {
        IMarketV2(MARKET).checkTickSizeRoles(_msgSender());

        if (_tickSize == 0) revert PairInvalidInitializeData("tickSize");
        if (_lotSize == 0) revert PairInvalidInitializeData("lotSize");
        if (_tickSize * _lotSize % DENOMINATOR != 0) revert PairInvalidTickSize(_tickSize, _lotSize, DENOMINATOR);
        emit TickSizeUpdated(lotSize, _lotSize, tickSize, _tickSize);

        lotSize = _lotSize;
        tickSize = _tickSize;
        minTradeVolume = Math.mulDiv(_tickSize, _lotSize, DENOMINATOR);
    }

    function setPairFees(
        uint32 sellerMakerFeeBps_,
        uint32 sellerTakerFeeBps_,
        uint32 buyerMakerFeeBps_,
        uint32 buyerTakerFeeBps_
    ) external onlyOwner {
        _setFeeBps(sellerMakerFeeBps_, sellerTakerFeeBps_, buyerMakerFeeBps_, buyerTakerFeeBps_);
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

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
