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
import {IPair} from "./interfaces/IPair.sol";
import {IRouter, IRouterInitializer} from "./interfaces/IRouter.sol";
import {IWETH} from "./interfaces/IWETH.sol";

contract CrossDexRouter is
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

    error RouterInvalidInputData(bytes32);
    error RouterInvalidPairAddress(address);
    error RouterInvalidValue();
    error RouterCancelLimitExceeded(uint256 length, uint256 limit);

    event FindPrevPriceCountChanged(uint256 indexed before, uint256 indexed current);
    event MaxMatchCountChanged(uint256 indexed before, uint256 indexed current);
    event CancelLimitChanged(uint256 indexed before, uint256 indexed current);

    address public CROSS_DEX; // immutable
    IWETH public CROSS; // immutable

    uint256 public findPrevPriceCount;
    uint256 public maxMatchCount;
    uint256 public cancelLimit;

    uint256[45] private __gap;

    modifier checkValue() {
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
        if (_findPrevPriceCount == 0) revert RouterInvalidInputData("findPrevPriceCount");
        if (_maxMatchCount == 0) revert RouterInvalidInputData("maxMatchCount");
        if (_cancelLimit == 0) revert RouterInvalidInputData("cancelLimit");

        CROSS_DEX = _msgSender();
        CROSS = IWETH(payable(address(new WETH())));
        findPrevPriceCount = _findPrevPriceCount;
        maxMatchCount = _maxMatchCount;
        cancelLimit = _cancelLimit;

        __Context_init();
        __ReentrancyGuard_init();
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
    ) external payable nonReentrant validPair(pair) checkValue returns (uint256) {
        address _owner = _msgSender();
        IPair _pair = IPair(pair);
        IPair.Config memory info = _pair.getConfig();

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
    ) external payable nonReentrant validPair(pair) checkValue returns (uint256) {
        address _owner = _msgSender();
        IPair _pair = IPair(pair);
        IPair.Config memory info = _pair.getConfig();
        uint256 prevPrice = _pair.findPrevPrice(IPair.OrderSide.BUY, price, adjacent, findPrevPriceCount);

        uint256 volume = Math.mulDiv(price, amount, info.DENOMINATOR);
        if (address(info.QUOTE) == address(CROSS)) CROSS.mintTo{value: volume}(pair);
        else info.QUOTE.safeTransferFrom(_owner, address(pair), volume);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.BUY, owner: _owner, feeBps: 0, price: price, amount: amount});
        return _pair.submitLimitOrder(order, constraints, prevPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function submitSellMarket(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        checkValue
    {
        address _owner = _msgSender();
        IPair.Config memory info = IPair(pair).getConfig();

        IERC20 BASE = info.BASE;
        if (address(BASE) == address(CROSS)) CROSS.mintTo{value: amount}(address(pair));
        else BASE.safeTransferFrom(_owner, address(pair), amount);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.SELL, owner: _owner, feeBps: 0, price: 0, amount: 0});
        IPair(pair).submitMarketOrder(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function submitBuyMarket(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        checkValue
    {
        address _owner = _msgSender();
        IPair.Config memory info = IPair(pair).getConfig();

        IERC20 QUOTE = info.QUOTE;
        if (address(QUOTE) == address(CROSS)) CROSS.mintTo{value: amount}(address(pair));
        else QUOTE.safeTransferFrom(_owner, address(pair), amount);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.BUY, owner: _owner, feeBps: 0, price: 0, amount: 0});
        IPair(pair).submitMarketOrder(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function cancelOrder(address pair, uint256[] calldata orderIds) external validPair(pair) {
        uint256 length = orderIds.length;
        if (length != 0) {
            if (length > cancelLimit) revert RouterCancelLimitExceeded(length, cancelLimit);
            IPair(pair).cancelOrder(_msgSender(), orderIds);
        }
    }

    function _toMaxMatchCount(uint256 _maxMatchCount) private view returns (uint256) {
        return _maxMatchCount == 0 || _maxMatchCount > maxMatchCount ? maxMatchCount : _maxMatchCount;
    }

    function setfindPrevPriceCount(uint256 _findPrevPriceCount) external onlyOwner {
        if (_findPrevPriceCount == 0) revert RouterInvalidInputData("findPrevPriceCount");
        emit FindPrevPriceCountChanged(findPrevPriceCount, _findPrevPriceCount);
        findPrevPriceCount = _findPrevPriceCount;
    }

    function setMaxMatchCount(uint256 _maxMatchCount) external onlyOwner {
        if (_maxMatchCount == 0) revert RouterInvalidInputData("findPrevPriceCount");
        emit MaxMatchCountChanged(maxMatchCount, _maxMatchCount);
        maxMatchCount = _maxMatchCount;
    }

    function setCancelLimit(uint256 _cancelLimit) external onlyOwner {
        if (_cancelLimit == 0) revert RouterInvalidInputData("cancelLimit");
        emit CancelLimitChanged(cancelLimit, _cancelLimit);
        cancelLimit = _cancelLimit;
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
