// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin-contracts-5.2.0/access/Ownable.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeCast} from "@openzeppelin-contracts-5.2.0/utils/math/SafeCast.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableSet.sol";

interface ICrossDex {
    function allMarkets() external view returns (address[] memory, address[] memory);
}

interface IMarket {
    function QUOTE() external view returns (address);
    function allPairs() external view returns (address[] memory, address[] memory);
}

interface IPair {
    function matchedPrice() external view returns (uint256);
    function setTickSize(uint256 _lotSize, uint256 _tickSize) external;
}

contract TickSizeSetter is Ownable {
    using SafeCast for int256;
    using EnumerableSet for EnumerableSet.UintSet;

    error TickSizeSetterInvalidInitializeData(bytes32 field);

    struct PairConfig {
        address quote;
        address base;
        uint8 quoteDecimals;
        uint8 baseDecimals;
    }

    struct SizeFormat {
        /// @dev
        /// 0.1 일 경우 unit = 1, scale = -1
        /// 5 일 경우 unit = 5, scale = 0
        /// 10 일 경우 unit = 1, scale = 1
        uint8 minPriceUnit;
        int16 minPriceScale;
        uint8 tickSizeUnit;
        int16 tickSizeScale;
        uint8 lotSizeUnit;
        int16 lotSizeScale;
    }

    struct ResolvedSize {
        uint256 minPrice;
        uint256 tickSize;
        uint256 lotSize;
    }

    // keccak256(abi.encode(uint256(keccak256("crossdex.ticsizesetter.currentupdatetimestamp")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CURRENT_UPDATE_TIMESTAMP =
        0xfd0e5d4f9b88892d3b04349a0e2bc0d1359414c21932fcd7d5a523a6c0a5cd00;

    ICrossDex public immutable CROSS_DEX;
    uint256 public updateInterval = 86400; // 1 day

    mapping(address pair => PairConfig) public pairConfigs;
    mapping(address pair => uint256) public lastUpdateTimestamp;

    SizeFormat[] public sizeFormats;

    EnumerableSet.UintSet private _allDecimals;
    mapping(uint8 decimals => ResolvedSize[]) public resolvedSizes;

    constructor(address owner_, address _crossDex) Ownable(owner_) {
        if (_crossDex == address(0)) revert TickSizeSetterInvalidInitializeData("crossDex");
        CROSS_DEX = ICrossDex(_crossDex);

        // gte ~ lt price | ticSize | lotSize
        // -----------------------------------
        //      ~ 0.1     | 0.0001  | 1
        // 0.1  ~ 1       | 0.001   | 1
        // 1    ~ 10      | 0.01    | 1
        // 10   ~ 50      | 0.1     | 1
        // 50   ~ 100     | 0.5     | 0.1
        // 100  ~ 500     | 1       | 0.01
        // 500  ~ 1000    | 5       | 0.001
        // 1000 ~         | 10      | 0.0001
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: -1,
                tickSizeUnit: 1,
                tickSizeScale: -4,
                lotSizeUnit: 1,
                lotSizeScale: 0
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 0,
                tickSizeUnit: 1,
                tickSizeScale: -3,
                lotSizeUnit: 1,
                lotSizeScale: 0
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 1,
                tickSizeUnit: 1,
                tickSizeScale: -2,
                lotSizeUnit: 1,
                lotSizeScale: 0
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 5,
                minPriceScale: 1,
                tickSizeUnit: 1,
                tickSizeScale: -1,
                lotSizeUnit: 5,
                lotSizeScale: -1
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 2,
                tickSizeUnit: 5,
                tickSizeScale: -1,
                lotSizeUnit: 1,
                lotSizeScale: -1
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 5,
                minPriceScale: 2,
                tickSizeUnit: 1,
                tickSizeScale: 0,
                lotSizeUnit: 1,
                lotSizeScale: -2
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 1,
                minPriceScale: 3,
                tickSizeUnit: 5,
                tickSizeScale: 0,
                lotSizeUnit: 1,
                lotSizeScale: -3
            })
        );
        sizeFormats.push(
            SizeFormat({
                minPriceUnit: 0,
                minPriceScale: 0,
                tickSizeUnit: 1,
                tickSizeScale: 2,
                lotSizeUnit: 1,
                lotSizeScale: -4
            })
        );
        _resolveSize(18);
    }

    function getSizeFormats() external view returns (SizeFormat[] memory) {
        return sizeFormats;
    }

    function findPriceIndex(uint8 quoteDecimals, uint256 price) public view returns (uint256) {
        ResolvedSize[] memory resolved = resolvedSizes[quoteDecimals];
        uint256 length = resolved.length;
        if (length == 0) revert();
        unchecked {
            --length;

            for (uint256 i = 0; i < length; ++i) {
                if (price < resolved[i].minPrice) return i;
            }
        }
        return length;
    }

    function allUpdates() public {
        (address[] memory quotes, address[] memory markets) = CROSS_DEX.allMarkets();
        uint256 length = markets.length;
        for (uint256 i = 0; i < length;) {
            address market = markets[i];
            address quote = quotes[i];
            uint8 quoteDecimals = IERC20Metadata(quote).decimals();
            if (_allDecimals.add(quoteDecimals)) _resolveSize(quoteDecimals);
            _update(market, 0, quote, quoteDecimals);
            unchecked {
                ++i;
            }
        }
        _clearUpdateTimestamp();
    }

    function update(address market, uint256 startIndex) external {
        address quote = IMarket(market).QUOTE();
        uint8 quoteDecimals = IERC20Metadata(quote).decimals();
        if (_allDecimals.add(quoteDecimals)) _resolveSize(quoteDecimals);
        _update(market, startIndex, quote, quoteDecimals);
        _clearUpdateTimestamp();
    }

    function manualUpdate(address pair, uint256 lotSize, uint256 tickSize) external onlyOwner {
        // Logic to force update the tick size
        IPair(pair).setTickSize(lotSize, tickSize);
    }

    function _update(address market, uint256 startIndex, address quote, uint8 quoteDecimals) private {
        (address[] memory bases, address[] memory pairs) = IMarket(market).allPairs();
        uint256 length = pairs.length;
        for (uint256 i = startIndex; i < length;) {
            address pair = pairs[i];

            // check latest update timestamp each pair
            if (_tryUpdateTimestamp(pair)) {
                address base = bases[i];
                uint8 baseDecimals = pairConfigs[pair].baseDecimals;
                // set pair config if not set
                if (baseDecimals == 0) {
                    baseDecimals = IERC20Metadata(base).decimals();
                    if (_allDecimals.add(baseDecimals)) _resolveSize(baseDecimals);
                    pairConfigs[pair] =
                        PairConfig({quote: quote, base: base, quoteDecimals: quoteDecimals, baseDecimals: baseDecimals});
                }

                uint256 tickSize;
                uint256 lotSize;

                // get pair price and index from resolvedSizes
                uint256 price = IPair(pair).matchedPrice();
                if (price != 0) {
                    uint256 index = findPriceIndex(quoteDecimals, price);

                    if (quoteDecimals == baseDecimals) {
                        ResolvedSize memory resolved = resolvedSizes[quoteDecimals][index];
                        tickSize = resolved.tickSize;
                        lotSize = resolved.lotSize;
                    } else {
                        tickSize = resolvedSizes[quoteDecimals][index].tickSize;
                        lotSize = resolvedSizes[baseDecimals][index].lotSize;
                    }

                    IPair(pair).setTickSize(lotSize, tickSize);
                }
            }
            unchecked {
                ++i;
            }
        }
    }

    function _tryUpdateTimestamp(address pair) private returns (bool) {
        uint256 updateTime = _calcUpdateTimestamp();
        bool ok = lastUpdateTimestamp[pair] < updateTime;
        if (ok) lastUpdateTimestamp[pair] = updateTime;
        return ok;
    }

    function _calcUpdateTimestamp() private returns (uint256) {
        uint256 updateTime;
        assembly {
            updateTime := tload(CURRENT_UPDATE_TIMESTAMP)
        }
        if (updateTime == 0) {
            updateTime = block.timestamp - (block.timestamp % updateInterval);
            assembly {
                tstore(CURRENT_UPDATE_TIMESTAMP, updateTime)
            }
        }
        return updateTime;
    }

    function _clearUpdateTimestamp() private {
        assembly {
            tstore(CURRENT_UPDATE_TIMESTAMP, 0)
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

    function _calcValue(uint8 decimals, uint8 unit, int16 scale) private pure returns (uint256) {
        int256 _scale = SafeCast.toInt256(uint256(decimals)) + int256(scale);
        return unit * (10 ** _scale.toUint256());
    }
}
