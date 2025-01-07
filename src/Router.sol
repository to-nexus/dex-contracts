// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.1.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.1.0/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.1.0/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.1.0/utils/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.1.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.1.0/utils/math/Math.sol";
import {IPair} from "./Pair.sol";

contract Router is ERC1967Proxy {
    constructor(address implementation)
        ERC1967Proxy(implementation, abi.encodeWithSelector(RouterImpl.initialize.selector))
    {}
}

contract RouterImpl is UUPSUpgradeable, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using SafeERC20 for IERC20;
    using Math for uint256;

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

    function addPair(IPair pair) external onlyOwner {
        if (allPairs[address(pair)].DENOMINATOR != 0) revert();
        uint256 DENOMINATOR = pair.DENOMINATOR();
        if (DENOMINATOR == 0) revert();

        IERC20 BASE = pair.BASE();
        IERC20 QUOTE = pair.QUOTE();

        if (address(BASE) != address(0)) BASE.forceApprove(address(pair), type(uint256).max);
        if (address(QUOTE) != address(0)) QUOTE.forceApprove(address(pair), type(uint256).max);

        allPairs[address(pair)] = Pair({BASE: BASE, QUOTE: QUOTE, DENOMINATOR: DENOMINATOR});
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
