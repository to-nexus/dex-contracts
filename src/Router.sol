// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.1.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.1.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.1.0/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.1.0/utils/ReentrancyGuardUpgradeable.sol";

import {IPair} from "./interfaces/IPair.sol";

contract Router is ERC1967Proxy {
    constructor(address implementation)
        ERC1967Proxy(implementation, abi.encodeWithSelector(RouterImpl.initialize.selector))
    {}
}

contract RouterImpl is UUPSUpgradeable, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using SafeERC20 for IERC20;
    using Math for uint256;

    error RouterAlreadyAddedPair(address);
    error RouterInvalidPairAddress(address);

    event PairAdded(address indexed pair, address indexed base, address indexed quote);
    event PairRemoved(address indexed pair);

    struct Pair {
        IERC20 BASE;
        IERC20 QUOTE;
        uint256 DENOMINATOR;
    }

    mapping(address pair => Pair) public allPairs;

    receive() external payable {
        if (msg.value != 0) revert();
    }

    function initialize() external initializer {
        __Ownable_init(_msgSender());
        __ReentrancyGuard_init();
    }

    function addPair(address pair) external onlyOwner {
        if (allPairs[pair].DENOMINATOR != 0) revert RouterAlreadyAddedPair(pair);

        IPair ipair = IPair(pair);
        uint256 DENOMINATOR = ipair.DENOMINATOR();
        if (DENOMINATOR == 0) revert RouterInvalidPairAddress(pair);
        IERC20 BASE = ipair.BASE();
        IERC20 QUOTE = ipair.QUOTE();

        if (address(BASE) != address(0)) BASE.forceApprove(pair, type(uint256).max);
        if (address(QUOTE) != address(0)) QUOTE.forceApprove(pair, type(uint256).max);

        allPairs[pair] = Pair({BASE: BASE, QUOTE: QUOTE, DENOMINATOR: DENOMINATOR});
        emit PairAdded(pair, address(BASE), address(QUOTE));
    }

    function removePair(address pair) external onlyOwner {
        Pair memory pairInfo = allPairs[pair];
        if (pairInfo.DENOMINATOR == 0) revert RouterInvalidPairAddress(pair);

        IPair ipair = IPair(pair);
        IERC20 BASE = ipair.BASE();
        IERC20 QUOTE = ipair.QUOTE();

        if (address(BASE) != address(0)) BASE.forceApprove(pair, 0);
        if (address(QUOTE) != address(0)) QUOTE.forceApprove(pair, 0);

        delete allPairs[pair];
        emit PairRemoved(pair);
    }

    function sellLimitOrder(address pair, uint256 price, uint256 amount)
        external
        payable
        nonReentrant
        returns (uint256)
    {
        address owner = _msgSender();
        IERC20 BASE = allPairs[pair].BASE;

        BASE.safeTransferFrom(owner, address(this), amount);
        IPair.Order memory order =
            IPair.Order({_type: IPair.OrderType.SELL, owner: owner, price: price, amount: amount});

        return IPair(pair).limit(order);
    }

    function buyLimitOrder(address pair, uint256 price, uint256 amount)
        external
        payable
        nonReentrant
        returns (uint256)
    {
        address owner = _msgSender();
        Pair memory pairInfo = allPairs[pair];
        IERC20 QUOTE = pairInfo.QUOTE;

        uint256 value = Math.mulDiv(price, amount, pairInfo.DENOMINATOR);

        QUOTE.safeTransferFrom(owner, address(this), value);
        IPair.Order memory order = IPair.Order({_type: IPair.OrderType.BUY, owner: owner, price: price, amount: amount});

        return IPair(pair).limit(order);
    }

    function sellMarketOrder(address pair, uint256 amount) external payable nonReentrant {
        address owner = _msgSender();
        IERC20 BASE = allPairs[pair].BASE;

        BASE.safeTransferFrom(owner, address(this), amount);
        IPair.Order memory order = IPair.Order({_type: IPair.OrderType.SELL, owner: owner, price: 0, amount: 0});

        IPair(pair).market(order, amount);
    }

    function buyMarketOrder(address pair, uint256 amount) external payable nonReentrant {
        address owner = _msgSender();
        IERC20 QUOTE = allPairs[pair].QUOTE;

        QUOTE.safeTransferFrom(owner, address(this), amount);
        IPair.Order memory order = IPair.Order({_type: IPair.OrderType.BUY, owner: owner, price: 0, amount: 0});

        IPair(pair).market(order, amount);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}

    uint256[49] __gap;
}
