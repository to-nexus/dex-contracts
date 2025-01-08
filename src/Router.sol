// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.1.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.1.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";

import {EnumerableSet} from "@openzeppelin-contracts-5.1.0/utils/structs/EnumerableSet.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.1.0/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.1.0/utils/ReentrancyGuardUpgradeable.sol";

import {IPair} from "./interfaces/IPair.sol";

contract Router is ERC1967Proxy {
    constructor(address implementation, uint256 _maxMatchCount)
        ERC1967Proxy(implementation, abi.encodeWithSelector(RouterImpl.initialize.selector, _maxMatchCount))
    {}
}

contract RouterImpl is UUPSUpgradeable, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;
    using Math for uint256;

    error RouterInitializeData(bytes32);
    error RouterAlreadyAddedPair(address);
    error RouterInvalidPairAddress(address);

    event PairAdded(address indexed pair, address indexed base, address indexed quote);
    event PairRemoved(address indexed pair);

    struct PairInfo {
        IERC20 QUOTE;
        IERC20 BASE;
        uint256 QUOTE_DENOMINATOR;
        uint256 BASE_DENOMINATOR;
    }

    uint256 public maxMatchCount;
    EnumerableSet.AddressSet private _allPairs;
    mapping(address pair => PairInfo) public pairInfo;

    modifier validPair(address pair) {
        if (!_allPairs.contains(pair)) revert RouterInvalidPairAddress(pair);
        _;
    }

    receive() external payable {
        if (msg.value != 0) revert();
    }

    function initialize(uint256 _maxMatchCount) external initializer {
        if (_maxMatchCount == 0) revert RouterInitializeData("maxMatchCount");

        maxMatchCount = _maxMatchCount;

        __Ownable_init(_msgSender());
        __ReentrancyGuard_init();
    }

    function sellLimitOrder(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        returns (uint256)
    {
        address owner = _msgSender();

        _pairInfo(pair).BASE.safeTransferFrom(owner, address(this), amount);

        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.SELL, owner: owner, feePermile: 0, price: price, amount: amount});
        return IPair(pair).limit(order, searchPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function buyLimitOrder(address pair, uint256 price, uint256 amount, uint256 searchPrice, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
        returns (uint256)
    {
        address owner = _msgSender();

        PairInfo memory info = _pairInfo(pair);
        uint256 volumn = Math.mulDiv(price, amount, info.BASE_DENOMINATOR);
        info.QUOTE.safeTransferFrom(owner, address(this), volumn);

        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.BUY, owner: owner, feePermile: 0, price: price, amount: amount});
        return IPair(pair).limit(order, searchPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function sellMarketOrder(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
    {
        address owner = _msgSender();

        _pairInfo(pair).BASE.safeTransferFrom(owner, address(this), amount);

        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.SELL, owner: owner, feePermile: 0, price: 0, amount: 0});
        IPair(pair).market(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function buyMarketOrder(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        validPair(pair)
    {
        address owner = _msgSender();

        _pairInfo(pair).QUOTE.safeTransferFrom(owner, address(this), amount);

        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.BUY, owner: owner, feePermile: 0, price: 0, amount: 0});
        IPair(pair).market(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function cancelOrder(address pair, uint256[] calldata orderIds) external validPair(pair) {
        IPair(pair).cancel(_msgSender(), orderIds);
    }

    function _pairInfo(address pair) private view returns (PairInfo memory) {
        return pairInfo[pair];
    }

    function _toMaxMatchCount(uint256 _maxMatchCount) private view returns (uint256) {
        return _maxMatchCount == 0 || _maxMatchCount > maxMatchCount ? maxMatchCount : _maxMatchCount;
    }

    function addPair(address pair) external onlyOwner {
        if (!_allPairs.add(pair)) revert RouterAlreadyAddedPair(pair);

        IPair ipair = IPair(pair);
        uint256 QUOTE_DENOMINATOR = ipair.QUOTE_DENOMINATOR();
        uint256 BASE_DENOMINATOR = ipair.BASE_DENOMINATOR();
        IERC20 BASE = ipair.BASE();
        IERC20 QUOTE = ipair.QUOTE();
        if (
            QUOTE_DENOMINATOR == 0 || BASE_DENOMINATOR == 0 || address(QUOTE) == address(0)
                || address(BASE) == address(0)
        ) revert RouterInvalidPairAddress(pair);

        BASE.forceApprove(pair, type(uint256).max);
        QUOTE.forceApprove(pair, type(uint256).max);

        pairInfo[pair] = PairInfo({
            QUOTE: QUOTE,
            BASE: BASE,
            QUOTE_DENOMINATOR: QUOTE_DENOMINATOR,
            BASE_DENOMINATOR: BASE_DENOMINATOR
        });
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

    uint256[48] __gap;
}
