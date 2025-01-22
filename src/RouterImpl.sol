// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.2.0/utils/structs/EnumerableSet.sol";

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.2.0/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.2.0/utils/ReentrancyGuardUpgradeable.sol";

import {WCROSS} from "./WCROSS.sol";

import {IPair} from "./interfaces/IPair.sol";
import {IRouter, IRouterInitializer} from "./interfaces/IRouter.sol";

contract RouterImpl is IRouter, IRouterInitializer, UUPSUpgradeable, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;
    using Address for address payable;
    using Math for uint256;

    error RouterInvalidCrossDex(address);
    error RouterInitializeData(bytes32);
    error RouterAlreadyAddedPair(address);
    error RouterInvalidPairAddress(address);
    error RouterInvalidValue();

    event PairAdded(address indexed pair, address indexed base, address indexed quote);
    event PairRemoved(address indexed pair);

    struct PairInfo {
        IERC20 BASE;
        IERC20 QUOTE;
        uint256 DENOMINATOR;
    }

    address public CROSS_DEX; // immutable
    address payable public WCross; // immutable

    uint256 public maxMatchCount;
    EnumerableSet.AddressSet private _allPairs;
    mapping(address pair => PairInfo) public pairInfo;

    uint256[46] private __gap;

    modifier onlyCrossDex() {
        if (_msgSender() != CROSS_DEX) revert RouterInvalidCrossDex(_msgSender());
        _;
    }

    modifier validPair(address pair) {
        if (!_allPairs.contains(pair)) revert RouterInvalidPairAddress(pair);
        _;
    }

    modifier checkValue() {
        _;
        if (address(this).balance != 0) revert RouterInvalidValue();
    }

    receive() external payable checkValue {}

    function initialize(address _owner, uint256 _maxMatchCount) external override initializer {
        if (_maxMatchCount == 0) revert RouterInitializeData("maxMatchCount");

        CROSS_DEX = _msgSender();
        WCross = payable(address(new WCROSS()));
        maxMatchCount = _maxMatchCount;

        __Ownable_init(_owner);
        __ReentrancyGuard_init();
    }

    function isPair(address pair) external view override returns (bool) {
        return _allPairs.contains(pair);
    }

    function limitSell(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        checkValue
        returns (uint256)
    {
        address owner = _msgSender();

        PairInfo memory info = _pairInfo(pair);

        IERC20 BASE = info.BASE;
        if (address(BASE) == WCross) WCross.sendValue(amount);
        else BASE.safeTransferFrom(owner, address(this), amount);

        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.SELL, owner: owner, feePermil: 0, price: price, amount: amount});
        return IPair(pair).limit(order, searchPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function limitBuy(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        checkValue
        returns (uint256)
    {
        address owner = _msgSender();

        PairInfo memory info = _pairInfo(pair);
        uint256 volume = Math.mulDiv(price, amount, info.DENOMINATOR);

        IERC20 QUOTE = info.QUOTE;
        if (address(QUOTE) == WCross) WCross.sendValue(volume);
        else QUOTE.safeTransferFrom(owner, address(this), volume);

        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.BUY, owner: owner, feePermil: 0, price: price, amount: amount});
        return IPair(pair).limit(order, searchPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function marketSell(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        checkValue
    {
        address owner = _msgSender();

        PairInfo memory info = _pairInfo(pair);

        IERC20 BASE = info.BASE;
        if (address(BASE) == WCross) WCross.sendValue(amount);
        else BASE.safeTransferFrom(owner, address(this), amount);

        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.SELL, owner: owner, feePermil: 0, price: 0, amount: 0});
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

        PairInfo memory info = _pairInfo(pair);

        IERC20 QUOTE = info.QUOTE;
        if (address(QUOTE) == WCross) WCross.sendValue(amount);
        else QUOTE.safeTransferFrom(owner, address(this), amount);

        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.BUY, owner: owner, feePermil: 0, price: 0, amount: 0});
        IPair(pair).market(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function cancel(address pair, uint256[] calldata orderIds) external validPair(pair) {
        IPair(pair).cancel(_msgSender(), orderIds);
    }

    function _pairInfo(address pair) private view returns (PairInfo memory) {
        return pairInfo[pair];
    }

    function _toMaxMatchCount(uint256 _maxMatchCount) private view returns (uint256) {
        return _maxMatchCount == 0 || _maxMatchCount > maxMatchCount ? maxMatchCount : _maxMatchCount;
    }

    function addPair(address pair) external override onlyCrossDex {
        if (!_allPairs.add(pair)) revert RouterAlreadyAddedPair(pair);

        IPair iPair = IPair(pair);
        uint256 DENOMINATOR = iPair.DENOMINATOR();
        IERC20 BASE = iPair.BASE();
        IERC20 QUOTE = iPair.QUOTE();
        if (DENOMINATOR == 0 || address(QUOTE) == address(0) || address(BASE) == address(0)) {
            revert RouterInvalidPairAddress(pair);
        }

        BASE.forceApprove(pair, type(uint256).max);
        QUOTE.forceApprove(pair, type(uint256).max);

        pairInfo[pair] = PairInfo({BASE: BASE, QUOTE: QUOTE, DENOMINATOR: DENOMINATOR});
        emit PairAdded(pair, address(BASE), address(QUOTE));
    }

    function removePair(address pair) external onlyOwner {
        if (!_allPairs.remove(pair)) revert RouterInvalidPairAddress(pair);
        PairInfo memory info = _pairInfo(pair);
        info.BASE.forceApprove(pair, 0);
        info.QUOTE.forceApprove(pair, 0);

        delete pairInfo[pair];
        emit PairRemoved(pair);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
