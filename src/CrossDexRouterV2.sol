// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableSet.sol";

import {ContextUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/utils/ContextUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.2.0/utils/ReentrancyGuardUpgradeable.sol";

import {WETH} from "./WETH.sol";

import {ICrossDex} from "./interfaces/ICrossDex.sol";
import {IOwnable} from "./interfaces/IOwnable.sol";
import {IPair, IPairV2} from "./interfaces/IPair.sol";
import {IRouter, IRouterInitializer} from "./interfaces/IRouter.sol";
import {IWETH} from "./interfaces/IWETH.sol";

contract CrossDexRouterV2 is
    IRouter,
    IRouterInitializer,
    IOwnable,
    UUPSUpgradeable,
    ContextUpgradeable,
    ReentrancyGuardUpgradeable
{
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;
    using Math for uint256;

    uint32 private constant BPS_DENOMINATOR = 10000;

    error RouterInvalidInputData(bytes32);
    error RouterInvalidPairAddress(address);
    error RouterInvalidValue();
    error RouterCancelLimitExceeded(uint256 length, uint256 limit);
    error RouterContractAccountBlocked(address);

    event FindPrevPriceCountChanged(uint256 indexed before, uint256 indexed current);
    event MaxMatchCountChanged(uint256 indexed before, uint256 indexed current);
    event CancelLimitChanged(uint256 indexed before, uint256 indexed current);
    event WhitelistedCodeAccountSet(address indexed account, bool whitelisted);

    address public CROSS_DEX; // immutable
    IWETH public CROSS; // immutable

    uint256 public findPrevPriceCount;
    uint256 public maxMatchCount;
    uint256 public cancelLimit;

    EnumerableSet.AddressSet private whitelistedCodeAccounts;

    uint256[44] private __gap;

    modifier submitCheck() {
        _checkAccountCode(_msgSender());
        _;
        if (address(this).balance != 0) revert RouterInvalidValue();
    }

    modifier validPair(address pair) {
        if (!isPair(pair)) revert RouterInvalidPairAddress(pair);
        _;
    }

    modifier onlyOwner() {
        if (_msgSender() != owner()) revert OwnableUnauthorizedAccount(_msgSender());
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(uint256 _findPrevPriceCount, uint256 _maxMatchCount, uint256 _cancelLimit)
        external
        override
        initializer
    {
        __Context_init();
        __ReentrancyGuard_init();

        if (_findPrevPriceCount == 0) revert RouterInvalidInputData("findPrevPriceCount");
        if (_maxMatchCount == 0) revert RouterInvalidInputData("maxMatchCount");
        if (_cancelLimit == 0) revert RouterInvalidInputData("cancelLimit");

        CROSS_DEX = _msgSender();
        CROSS = IWETH(payable(address(new WETH())));
        findPrevPriceCount = _findPrevPriceCount;
        maxMatchCount = _maxMatchCount;
        cancelLimit = _cancelLimit;
    }

    function isPair(address pair) public view override returns (bool) {
        return ICrossDex(CROSS_DEX).pairToMarket(pair) != address(0);
    }

    function owner() public view override returns (address) {
        return IOwnable(CROSS_DEX).owner();
    }

    function submitSellLimit(
        address pair,
        uint256 price,
        uint256 amount,
        IPair.LimitConstraints constraints,
        uint256[2] memory adjacent,
        uint256 _maxMatchCount
    ) external payable nonReentrant submitCheck validPair(pair) returns (uint256) {
        address _owner = _msgSender();
        IPairV2 _pair = IPairV2(pair);
        IPairV2.Config memory info = _pair.getConfig();

        uint256 prevPrice = _pair.findPrevPrice(IPair.OrderSide.SELL, price, adjacent, findPrevPriceCount);

        if (address(info.BASE) == address(CROSS)) CROSS.mintTo{value: amount}(pair);
        else info.BASE.safeTransferFrom(_owner, pair, amount);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.SELL, owner: _owner, feeBps: 0, price: price, amount: amount});
        return _pair.submitLimitOrder(order, constraints, prevPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function submitBuyLimit(
        address pair,
        uint256 price,
        uint256 amount,
        IPair.LimitConstraints constraints,
        uint256[2] calldata adjacent,
        uint256 _maxMatchCount
    ) external payable nonReentrant submitCheck validPair(pair) returns (uint256) {
        address _owner = _msgSender();
        IPairV2 _pair = IPairV2(pair);
        IPairV2.Config memory info = _pair.getConfig();
        uint256 prevPrice = _pair.findPrevPrice(IPair.OrderSide.BUY, price, adjacent, findPrevPriceCount);

        {
            uint256 volume = Math.mulDiv(price, amount, info.DENOMINATOR);
            // Use taker fee since limit order can be immediately matched and become taker
            volume = _calculateRequireBuyVolume(pair, volume);

            if (address(info.QUOTE) == address(CROSS)) CROSS.mintTo{value: volume}(pair);
            else info.QUOTE.safeTransferFrom(_owner, address(pair), volume);
        }

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.BUY, owner: _owner, feeBps: 0, price: price, amount: amount});
        return _pair.submitLimitOrder(order, constraints, prevPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function submitSellMarket(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        submitCheck
        validPair(pair)
    {
        address _owner = _msgSender();
        IPairV2.Config memory info = IPairV2(pair).getConfig();

        IERC20 BASE = info.BASE;
        if (address(BASE) == address(CROSS)) CROSS.mintTo{value: amount}(address(pair));
        else BASE.safeTransferFrom(_owner, address(pair), amount);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.SELL, owner: _owner, feeBps: 0, price: 0, amount: 0});
        IPairV2(pair).submitMarketOrder(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function submitBuyMarket(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        submitCheck
        validPair(pair)
    {
        address _owner = _msgSender();
        IPairV2.Config memory info = IPairV2(pair).getConfig();

        IERC20 QUOTE = info.QUOTE;
        {
            uint256 volume = _calculateRequireBuyVolume(pair, amount);
            if (address(QUOTE) == address(CROSS)) CROSS.mintTo{value: volume}(address(pair));
            else QUOTE.safeTransferFrom(_owner, address(pair), volume);
        }

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.BUY, owner: _owner, feeBps: 0, price: 0, amount: 0});
        IPairV2(pair).submitMarketOrder(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function cancelOrder(address pair, uint256[] calldata orderIds) external nonReentrant validPair(pair) {
        uint256 length = orderIds.length;
        if (length != 0) {
            if (length > cancelLimit) revert RouterCancelLimitExceeded(length, cancelLimit);
            IPairV2(pair).cancelOrder(_msgSender(), orderIds);
        }
    }

    /**
     * @dev Calculate the volume including buyer taker fee for buy orders
     * Use taker fee since both limit and market orders can be matched immediately
     * @param pair The pair contract address
     * @param baseVolume The base volume without fee
     * @return The volume including fee
     */
    function _calculateRequireBuyVolume(address pair, uint256 baseVolume) private view returns (uint256) {
        (,,, uint32 buyerTakerFeeBps) = IPairV2(pair).getEffectiveFees();
        return Math.mulDiv(baseVolume, BPS_DENOMINATOR + buyerTakerFeeBps, BPS_DENOMINATOR);
    }

    function _toMaxMatchCount(uint256 _maxMatchCount) private view returns (uint256) {
        return _maxMatchCount == 0 || _maxMatchCount > maxMatchCount ? maxMatchCount : _maxMatchCount;
    }

    function _checkAccountCode(address account) private view {
        if (whitelistedCodeAccounts.contains(account)) return;
        if (account.code.length != 0) revert RouterContractAccountBlocked(account);
    }

    function setFindPrevPriceCount(uint256 _findPrevPriceCount) external onlyOwner {
        if (_findPrevPriceCount == 0) revert RouterInvalidInputData("findPrevPriceCount");
        emit FindPrevPriceCountChanged(findPrevPriceCount, _findPrevPriceCount);
        findPrevPriceCount = _findPrevPriceCount;
    }

    function setMaxMatchCount(uint256 _maxMatchCount) external onlyOwner {
        if (_maxMatchCount == 0) revert RouterInvalidInputData("_maxMatchCount");
        emit MaxMatchCountChanged(maxMatchCount, _maxMatchCount);
        maxMatchCount = _maxMatchCount;
    }

    function setCancelLimit(uint256 _cancelLimit) external onlyOwner {
        if (_cancelLimit == 0) revert RouterInvalidInputData("cancelLimit");
        emit CancelLimitChanged(cancelLimit, _cancelLimit);
        cancelLimit = _cancelLimit;
    }

    function setWhitelistedCodeAccount(address[] memory accounts, bool whitelisted) external onlyOwner {
        if (whitelisted) {
            for (uint256 i = 0; i < accounts.length;) {
                address account = accounts[i];
                if (whitelistedCodeAccounts.add(account)) emit WhitelistedCodeAccountSet(account, true);
                unchecked {
                    ++i;
                }
            }
        } else {
            for (uint256 i = 0; i < accounts.length;) {
                address account = accounts[i];
                if (whitelistedCodeAccounts.remove(account)) emit WhitelistedCodeAccountSet(account, false);
                unchecked {
                    ++i;
                }
            }
        }
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
