// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin-contracts-5.2.0/access/Ownable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";

import "./interfaces/ITickSizeSetter.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";
import {SafeCast} from "@openzeppelin-contracts-5.2.0/utils/math/SafeCast.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableSet.sol";

contract TickSizeSetter is Ownable {
    using SafeCast for int256;
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;

    error TickSizeSetterZeroInput(bytes32 field);
    error TickSizeSetterInvalidPrice(uint256 index);
    error TickSizeSetterNotResolvedDeciamls(address pair, uint8 decimals);
    error TickSizeSetterInvalidAccess();
    error TickSizeSetterInvalidPair(address pair);

    event UpdateIntervalChanged(uint256 interval);
    event AddSkipPair(address indexed pair);
    event RemoveSkipPair(address indexed pair);
    event SizeFormatsChanged(SizeFormat[] formats);
    event ResolvedSizesByPairUpdated(address pair, ResolvedSize[] sizes);
    event ResolvedSizesByPairRemoved(address pair);

    struct PairConfig {
        address quote;
        address base;
        uint8 quoteDecimals;
        uint8 baseDecimals;
    }

    struct SizeFormat {
        /// @dev
        /// - 0.1 : unit = 1, scale = -1
        /// - 5   : unit = 5, scale = 0
        /// - 10  : unit = 1, scale = 1
        uint8 minPriceUnit;
        int8 minPriceScale;
        uint8 tickSizeUnit;
        int8 tickSizeScale;
        uint8 lotSizeUnit;
        int8 lotSizeScale;
    }

    struct ResolvedSize {
        uint256 minPrice;
        uint256 tickSize;
        uint256 lotSize;
    }

    // keccak256(abi.encode(uint256(keccak256("crossdex.ticksizesetter.updatetimestamp")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant UPDATE_TIMESTAMP = 0x72e6793f966897c43b0e33775d05baa76ff9110155a84246f9453c16d5d68200;
    // keccak256(abi.encode(uint256(keccak256("crossdex.ticksizesetter.pricetimestamp")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant PRICE_TIMESTAMP = 0x41a83abda80d11e65ee82cd8e180c62e956a7775c30101f5415ca33e35d07000;
    uint256 private constant UPDATE_MIN_GAS_LEFT = 515_000;

    ICrossDexForTickSizeSetter public immutable CROSS_DEX;
    IRouterForTickSizeSetter public immutable ROUTER;

    uint256 public updateInterval = 86400; // 1 day

    mapping(address pair => PairConfig) public pairConfigs;
    mapping(address pair => uint256) public lastUpdateTimestamp;

    EnumerableSet.AddressSet private _skipPairs;

    SizeFormat[] public sizeFormats;

    EnumerableSet.UintSet private _allDecimals;
    mapping(uint8 decimals => ResolvedSize[]) public resolvedSizes;
    mapping(address pair => ResolvedSize[]) public resolvedSizesByPair;

    modifier updateTimestamp() {
        _setTimestamp();
        _;
        _clearTimestamp();
    }

    constructor(address owner_, address _crossDex) Ownable(owner_) {
        if (_crossDex == address(0)) revert TickSizeSetterZeroInput("crossDex");
        CROSS_DEX = ICrossDexForTickSizeSetter(_crossDex);
        ROUTER = IRouterForTickSizeSetter(CROSS_DEX.ROUTER());
        // gte ~ lt price | tickSize | lotSize
        // -----------------------------------
        //      ~ 0.1     | 0.0001  | 1
        // 0.1  ~ 1       | 0.001   | 1
        // 1    ~ 10      | 0.01    | 1
        // 10   ~ 50      | 0.1     | 0.5
        // 50   ~ 100     | 0.5     | 0.1
        // 100  ~ 500     | 1       | 0.01
        // 500  ~ 1000    | 5       | 0.001
        // 1000 ~         | 10      | 0.0001
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: -1, // 0.1
                tickSizeUnit: 1,
                tickSizeScale: -4, // 0.0001
                lotSizeUnit: 1,
                lotSizeScale: 0 // 1
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 0, // 1
                tickSizeUnit: 1,
                tickSizeScale: -3, // 0.001
                lotSizeUnit: 1,
                lotSizeScale: 0 // 1
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 1, // 10
                tickSizeUnit: 1,
                tickSizeScale: -2, // 0.01
                lotSizeUnit: 1,
                lotSizeScale: 0 // 1
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 5,
                minPriceScale: 1, // 50
                tickSizeUnit: 1,
                tickSizeScale: -1, // 0.1
                lotSizeUnit: 5,
                lotSizeScale: -1 // 0.5
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 2, // 100
                tickSizeUnit: 5,
                tickSizeScale: -1, // 0.5
                lotSizeUnit: 1,
                lotSizeScale: -1 // 0.1
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 5,
                minPriceScale: 2, // 500
                tickSizeUnit: 1,
                tickSizeScale: 0, // 1
                lotSizeUnit: 1,
                lotSizeScale: -2 // 0.01
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 3, // 1000
                tickSizeUnit: 5,
                tickSizeScale: 0, // 5
                lotSizeUnit: 1,
                lotSizeScale: -3 // 0.001
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 0,
                minPriceScale: 0, // ~~~
                tickSizeUnit: 1,
                tickSizeScale: 2, // 100
                lotSizeUnit: 1,
                lotSizeScale: -4 // 0.0001
            })
        );
        _resolveSize(18);
    }

    function getSizeFormats() external view returns (SizeFormat[] memory) {
        return sizeFormats;
    }

    function allSkipPairs() external view returns (address[] memory) {
        return _skipPairs.values();
    }

    function allDecimals() external view returns (uint256[] memory) {
        return _allDecimals.values();
    }

    function tickSizeByPrice(address pair, uint256 price) public view returns (uint256 tickSize, uint256 lotSize) {
        PairConfig memory config = pairConfigs[pair];

        if (resolvedSizesByPair[pair].length != 0) {
            ResolvedSize[] memory resolved = resolvedSizesByPair[pair];
            (bool found, uint256 index) = _findPriceIndex(resolved, price);
            if (!found) revert TickSizeSetterNotResolvedDeciamls(pair, config.quoteDecimals);

            tickSize = resolved[index].tickSize;
            lotSize = resolved[index].lotSize;
        } else {
            ResolvedSize[] memory resolved = resolvedSizes[config.quoteDecimals];
            (bool found, uint256 index) = _findPriceIndex(resolved, price);
            if (!found) revert TickSizeSetterNotResolvedDeciamls(pair, config.quoteDecimals);

            if (config.quoteDecimals == config.baseDecimals) {
                tickSize = resolved[index].tickSize;
                lotSize = resolved[index].lotSize;
            } else {
                tickSize = resolved[index].tickSize;
                lotSize = resolvedSizes[config.baseDecimals][index].lotSize;
            }
        }
    }

    function calcTimestamp(uint256 timestamp) public view returns (uint256) {
        uint256 mod = timestamp % updateInterval;
        if (mod == 0) return timestamp;
        return timestamp - (mod);
    }

    // Returns the earliest market and start index that require an update.
    // If the market is returned as address(0), it means all pairs
    // have already been updated for the current interval.
    function updatable() external view returns (address, uint256) {
        uint256 updateTime = calcTimestamp(block.timestamp);
        uint256 priceTime = updateTime - updateInterval;

        (, address[] memory markets) = CROSS_DEX.allMarkets();
        uint256 mLen = markets.length;
        unchecked {
            for (uint256 i = 0; i < mLen; ++i) {
                address market = markets[i];
                (, address[] memory pairs) = IMarketForTickSizeSetter(market).allPairs();
                uint256 pLen = pairs.length;
                for (uint256 index = 0; index < pLen; ++index) {
                    address pair = pairs[index];
                    if (_skipPairs.contains(pair)) continue;
                    if (lastUpdateTimestamp[pair] != updateTime && _getPrice(pair, priceTime) != 0) {
                        return (market, index);
                    }
                }
            }
        }
        return (address(0), 0);
    }

    // Iterate over all pairs in the CrossDex contract and update their tick sizes.
    function allUpdates() external updateTimestamp {
        (address[] memory quotes, address[] memory markets) = CROSS_DEX.allMarkets();
        uint256 length = markets.length;
        for (uint256 i = 0; i < length;) {
            address market = markets[i];
            address quote = quotes[i];
            uint8 quoteDecimals = IERC20Metadata(quote).decimals();
            if (_allDecimals.add(quoteDecimals)) _resolveSize(quoteDecimals);
            if (_update(market, 0, type(uint256).max, quote, quoteDecimals)) break;
            unchecked {
                ++i;
            }
        }
    }

    // As the number of pairs increases, allUpdates may run out of gas
    // and stop midway without updating all pairs.
    // This function updates tick sizes by iterating over pairs in the market starting
    // from a specific pair.
    function update(address market, uint256 startIndex, uint256 count) external updateTimestamp {
        if (count == 0) count = type(uint256).max;

        address quote = IMarketForTickSizeSetter(market).QUOTE();
        uint8 quoteDecimals = IERC20Metadata(quote).decimals();
        if (_allDecimals.add(quoteDecimals)) _resolveSize(quoteDecimals);
        _update(market, startIndex, startIndex + count, quote, quoteDecimals);
    }

    function _findPriceIndex(ResolvedSize[] memory resolved, uint256 price) private pure returns (bool, uint256) {
        uint256 length = resolved.length;
        if (length == 0) return (false, 0);
        unchecked {
            // If the target value is greater than the price at index array - 1, the last tick must be used,
            // so there's no need to compare it — subtract 1 in advance.
            --length;

            // // Since the resolved array is expected to be short, a linear search is used.
            for (uint256 i = 0; i < length; ++i) {
                if (price < resolved[i].minPrice) return (true, i);
            }
        }
        return (true, length);
    }

    // Iterate over all pairs in the market and update those that have not been updated
    // in the current interval and have a non-zero matchedPrice.
    // Returns true if the loop exited early because of low remaining gas.
    function _update(address market, uint256 startIndex, uint256 endIndex, address quote, uint8 quoteDecimals)
        private
        returns (bool)
    {
        uint256 priceTime = _getPriceTimestamp();
        (address[] memory bases, address[] memory pairs) = IMarketForTickSizeSetter(market).allPairs();
        uint256 length = Math.min(pairs.length, endIndex);
        for (uint256 i = startIndex; i < length;) {
            // Break the loop if the remaining gas is too low to avoid running out of gas
            // and reverting the transaction.
            if (gasleft() < UPDATE_MIN_GAS_LEFT) return true;

            address pair = pairs[i];
            // check latest update timestamp each pair
            if (!_skipPairs.contains(pair) && _tryUpdateTimestamp(pair)) {
                address base = bases[i];
                uint8 baseDecimals = pairConfigs[pair].baseDecimals;
                // set pair config if not set
                if (baseDecimals == 0) {
                    baseDecimals = IERC20Metadata(base).decimals();
                    if (_allDecimals.add(baseDecimals)) _resolveSize(baseDecimals);
                    pairConfigs[pair] =
                        PairConfig({quote: quote, base: base, quoteDecimals: quoteDecimals, baseDecimals: baseDecimals});
                }

                uint256 price = _getPrice(pair, priceTime);
                IPairForTickSizeSetter ipair = IPairForTickSizeSetter(pair);
                if (price != 0 && !ipair.paused()) {
                    (uint256 tickSize, uint256 lotSize) = tickSizeByPrice(pair, price);
                    (uint256 tick, uint256 lot) = ipair.tickSizes();
                    if (tick != tickSize || lot != lotSize) ipair.setTickSize(lotSize, tickSize);
                }
            }

            unchecked {
                ++i;
            }
        }
        return false;
    }

    function _tryUpdateTimestamp(address pair) private returns (bool) {
        uint256 updateTime = _getUpdateTimestamp();
        bool ok = lastUpdateTimestamp[pair] < updateTime;

        if (ok) lastUpdateTimestamp[pair] = updateTime;
        return ok;
    }

    function _getUpdateTimestamp() private view returns (uint256) {
        uint256 updateTime;
        assembly {
            updateTime := tload(UPDATE_TIMESTAMP)
        }
        if (updateTime == 0) revert TickSizeSetterInvalidAccess();
        return updateTime;
    }

    function _getPriceTimestamp() private view returns (uint256) {
        uint256 priceTime;
        assembly {
            priceTime := tload(PRICE_TIMESTAMP)
        }
        if (priceTime == 0) revert TickSizeSetterInvalidAccess();
        return priceTime;
    }

    function _getPrice(address pair, uint256 priceTime) private view returns (uint256) {
        uint256 price = ROUTER.matchedPriceAt(pair, priceTime);
        if (price == 0) price = IPairForTickSizeSetter(pair).matchedPrice();
        return price;
    }

    function _setTimestamp() private {
        uint256 updateTime = calcTimestamp(block.timestamp);
        uint256 priceTime = updateTime - updateInterval;
        assembly {
            tstore(UPDATE_TIMESTAMP, updateTime)
            tstore(PRICE_TIMESTAMP, priceTime)
        }
    }

    function _clearTimestamp() private {
        assembly {
            tstore(UPDATE_TIMESTAMP, 0)
            tstore(PRICE_TIMESTAMP, 0)
        }
    }

    function _resolveSize(uint8 decimals) private {
        delete(resolvedSizes[decimals]);
        ResolvedSize[] storage resolved = resolvedSizes[decimals];

        uint256 length = sizeFormats.length;
        for (uint256 i = 0; i < length;) {
            SizeFormat memory format = sizeFormats[i];
            resolved.push(
                ResolvedSize({
                    minPrice: _calcValue(decimals, format.minPriceUnit, format.minPriceScale),
                    tickSize: _calcValue(decimals, format.tickSizeUnit, format.tickSizeScale),
                    lotSize: _calcValue(decimals, format.lotSizeUnit, format.lotSizeScale)
                })
            );
            unchecked {
                ++i;
            }
        }
    }

    function _calcValue(uint8 decimals, uint8 unit, int8 scale) private pure returns (uint256) {
        int256 _scale = SafeCast.toInt256(uint256(decimals)) + int256(scale);
        return unit * (10 ** _scale.toUint256());
    }

    function _checkPair(address pair) private view {
        if (pair == address(0)) revert TickSizeSetterZeroInput("pair");
        if (!CROSS_DEX.isMarket(CROSS_DEX.pairToMarket(pair))) revert TickSizeSetterInvalidPair(pair);
    }

    function forceUpdate(address pair, uint256 lotSize, uint256 tickSize) external onlyOwner {
        _checkPair(pair);
        IPairForTickSizeSetter(pair).setTickSize(lotSize, tickSize);
        if (_skipPairs.add(pair)) emit AddSkipPair(pair);
    }

    function removeSkipPair(address pair) external onlyOwner {
        if (!_skipPairs.remove(pair)) revert TickSizeSetterInvalidPair(pair);
        emit RemoveSkipPair(pair);
    }

    function setUpdateInterval(uint256 interval) external onlyOwner {
        if (interval == 0) revert TickSizeSetterZeroInput("interval");
        updateInterval = interval;
        emit UpdateIntervalChanged(interval);
    }

    function setSizeFormats(SizeFormat[] calldata formats) external onlyOwner {
        uint256 length = formats.length;
        if (length == 0) revert TickSizeSetterZeroInput("formats");

        delete(sizeFormats);
        uint256 beforePrice;

        unchecked {
            // To avoid validating the last price during setup—since it will always be used
            // subtract 1 beforehand.
            --length;
        }
        for (uint256 i = 0; i < length;) {
            uint256 price = _calcValue(18, formats[i].minPriceUnit, formats[i].minPriceScale);
            if (price <= beforePrice) revert TickSizeSetterInvalidPrice(i);
            beforePrice = price;

            sizeFormats.push(formats[i]);
            unchecked {
                ++i;
            }
        }
        // Add the final value that was intentionally excluded earlier.
        sizeFormats.push(formats[length]);

        uint256 lenDecimals = _allDecimals.length();
        for (uint256 i = 0; i < lenDecimals;) {
            _resolveSize(SafeCast.toUint8(_allDecimals.at(i)));
            unchecked {
                ++i;
            }
        }

        emit SizeFormatsChanged(formats);
    }

    function setResolvedSizesByPair(address pair, ResolvedSize[] calldata sizes) external onlyOwner {
        _checkPair(pair);

        uint256 length = sizes.length;
        if (length == 0) revert TickSizeSetterZeroInput("sizes");

        delete (resolvedSizesByPair[pair]);
        ResolvedSize[] storage resolved = resolvedSizesByPair[pair];

        uint256 beforePrice;
        unchecked {
            // To avoid validating the last price during setup—since it will always be used
            // subtract 1 beforehand.
            --length;

            for (uint256 i = 0; i < length; ++i) {
                if (sizes[i].minPrice <= beforePrice) revert TickSizeSetterInvalidPrice(i);
                beforePrice = sizes[i].minPrice;

                resolved.push(sizes[i]);
            }
            resolved.push(sizes[length]);
        }
        emit ResolvedSizesByPairUpdated(pair, sizes);
    }

    function removeResolvedSizesByPair(address pair) external onlyOwner {
        _checkPair(pair);

        if (resolvedSizesByPair[pair].length == 0) revert TickSizeSetterInvalidPair(pair);
        delete( resolvedSizesByPair[pair]);
        emit ResolvedSizesByPairRemoved(pair);
    }
}
