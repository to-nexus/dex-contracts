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

import {WCROSS} from "./WCROSS.sol";

import {ICrossDex} from "./interfaces/ICrossDex.sol";
import {IOwnable} from "./interfaces/IOwnable.sol";
import {IPair} from "./interfaces/IPair.sol";
import {IRouter, IRouterInitializer} from "./interfaces/IRouter.sol";
import {IWCROSS} from "./interfaces/IWCROSS.sol";

contract RouterImpl is IRouter, IRouterInitializer, UUPSUpgradeable, ContextUpgradeable, ReentrancyGuardUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;
    using Math for uint256;

    error RouterInitializeData(bytes32);
    error RouterInvalidPairAddress(address);
    error RouterInvalidValue();

    event PairAdded(address indexed pair, address indexed base, address indexed quote);
    event PairRemoved(address indexed pair);

    address public CROSS_DEX; // immutable
    IWCROSS public WCross; // immutable

    uint256 public maxMatchCount;

    uint256[47] private __gap;

    modifier checkValue() {
        _;
        if (address(this).balance != 0) revert RouterInvalidValue();
    }

    modifier validPair(address pair) {
        if (!isPair(pair)) revert RouterInvalidPairAddress(pair);
        _;
    }

    modifier onlyOwner() {
        if (_msgSender() != IOwnable(CROSS_DEX).owner()) revert IOwnable.OwnableUnauthorizedAccount(_msgSender());
        _;
    }

    constructor() {
        _disableInitializers();
    }

    receive() external payable checkValue {}

    function initialize(uint256 _maxMatchCount) external override initializer {
        if (_maxMatchCount == 0) revert RouterInitializeData("maxMatchCount");

        CROSS_DEX = _msgSender();
        WCross = IWCROSS(payable(address(new WCROSS())));
        maxMatchCount = _maxMatchCount;

        __Context_init();
        __ReentrancyGuard_init();
    }

    function isPair(address pair) public view override returns (bool) {
        return ICrossDex(CROSS_DEX).pairToMarket(pair) != address(0);
    }

    function limitSell(
        address pair,
        uint256 price,
        uint256 amount,
        IPair.LimitConstraints constraints,
        uint256 searchPrice,
        uint256 _maxMatchCount
    ) external payable nonReentrant validPair(pair) checkValue returns (uint256) {
        address owner = _msgSender();
        IPair.TokenConfig memory info = IPair(pair).getTokenConfig();

        IERC20 BASE = info.BASE;
        if (address(BASE) == address(WCross)) WCross.mintTo{value: amount}(address(pair));
        else BASE.safeTransferFrom(owner, pair, amount);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.SELL, owner: owner, feePermil: 0, price: price, amount: amount});
        return IPair(pair).limit(order, constraints, searchPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function limitBuy(
        address pair,
        uint256 price,
        uint256 amount,
        IPair.LimitConstraints constraints,
        uint256 searchPrice,
        uint256 _maxMatchCount
    ) external payable nonReentrant validPair(pair) checkValue returns (uint256) {
        address owner = _msgSender();
        IPair.TokenConfig memory info = IPair(pair).getTokenConfig();

        IERC20 QUOTE = info.QUOTE;
        uint256 volume = Math.mulDiv(price, amount, info.DENOMINATOR);
        if (address(QUOTE) == address(WCross)) WCross.mintTo{value: volume}(address(pair));
        else QUOTE.safeTransferFrom(owner, address(pair), volume);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.BUY, owner: owner, feePermil: 0, price: price, amount: amount});
        return IPair(pair).limit(order, constraints, searchPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function marketSell(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        checkValue
    {
        address owner = _msgSender();
        IPair.TokenConfig memory info = IPair(pair).getTokenConfig();

        IERC20 BASE = info.BASE;
        if (address(BASE) == address(WCross)) WCross.mintTo{value: amount}(address(pair));
        else BASE.safeTransferFrom(owner, address(pair), amount);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.SELL, owner: owner, feePermil: 0, price: 0, amount: 0});
        IPair(pair).market(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function marketBuy(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        checkValue
    {
        address owner = _msgSender();
        IPair.TokenConfig memory info = IPair(pair).getTokenConfig();

        IERC20 QUOTE = info.QUOTE;
        if (address(QUOTE) == address(WCross)) WCross.mintTo{value: amount}(address(pair));
        else QUOTE.safeTransferFrom(owner, address(pair), amount);

        IPair.Order memory order =
            IPair.Order({side: IPair.OrderSide.BUY, owner: owner, feePermil: 0, price: 0, amount: 0});
        IPair(pair).market(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function cancel(address pair, uint256[] calldata orderIds) external validPair(pair) {
        IPair(pair).cancel(_msgSender(), orderIds);
    }

    function _toMaxMatchCount(uint256 _maxMatchCount) private view returns (uint256) {
        return _maxMatchCount == 0 || _maxMatchCount > maxMatchCount ? maxMatchCount : _maxMatchCount;
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
