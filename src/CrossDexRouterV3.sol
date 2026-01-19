// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {UUPSUpgradeable} from "@openzeppelin-contracts-5.5.0/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.5.0/utils/structs/EnumerableSet.sol";

import {ReentrancyGuardTransient} from "@openzeppelin-contracts-5.5.0/utils/ReentrancyGuardTransient.sol";
import {ContextUpgradeable} from "@openzeppelin-contracts-upgradeable-5.5.0/utils/ContextUpgradeable.sol";

import {WETH} from "./WETH.sol";
import {ICrossDexV3} from "./interfaces/ICrossDexV3.sol";
import {IOwnable} from "./interfaces/IOwnable.sol";
import {IPairV3} from "./interfaces/IPairV3.sol";
import {IRouterV3} from "./interfaces/IRouterV3.sol";
import {IWETH} from "./interfaces/IWETH.sol";

contract CrossDexRouterV3 is UUPSUpgradeable, ContextUpgradeable, ReentrancyGuardTransient, IRouterV3, IOwnable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;
    using Math for uint256;

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

    uint256[43] private __gap;

    modifier checkSubmit() {
        _checkAccountCode(_msgSender());
        _;
        _checkNoRemainingValue();
    }

    function _checkNoRemainingValue() private view {
        if (address(this).balance != 0) revert RouterInvalidValue();
    }

    modifier validPair(address pair) {
        _checkValidPair(pair);
        _;
    }

    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    function _checkValidPair(address pair) private view {
        if (!isPair(pair)) revert RouterInvalidPairAddress(pair);
    }

    function _checkOwner() private view {
        if (_msgSender() != owner()) revert OwnableUnauthorizedAccount(_msgSender());
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

        if (_findPrevPriceCount == 0) revert RouterInvalidInputData("findPrevPriceCount");
        if (_maxMatchCount == 0) revert RouterInvalidInputData("maxMatchCount");
        if (_cancelLimit == 0) revert RouterInvalidInputData("cancelLimit");

        CROSS_DEX = _msgSender();
        CROSS = IWETH(payable(address(new WETH())));
        findPrevPriceCount = _findPrevPriceCount;
        maxMatchCount = _maxMatchCount;
        cancelLimit = _cancelLimit;
    }

    /**
     * @dev Calculate the total QUOTE volume required (including buyer taker fee)
     * @param pair The pair contract address
     * @param volume The base quote volume user wants to spend
     * @return The total volume needed including fee
     * @notice Not a view function due to delegatecall. Use eth_call for gas-free queries.
     */
    function getRequiredBuyVolume(address pair, uint256 volume) external validPair(pair) returns (uint256) {
        return IPairV3(pair).calcBuyVolumeWithFee(volume);
    }

    function submitSellLimit(
        address pair,
        uint256 price,
        uint256 amount,
        IPairV3.LimitConstraints constraints,
        uint256[2] memory adjacent,
        uint256 _maxMatchCount
    ) external payable nonReentrant checkSubmit validPair(pair) returns (uint256) {
        address _owner = _msgSender();
        IPairV3 _pair = IPairV3(pair);
        IPairV3.Config memory info = _pair.getConfig();

        uint256 prevPrice = _pair.findPrevPrice(IPairV3.OrderSide.SELL, price, adjacent, findPrevPriceCount);

        if (address(info.BASE) == address(CROSS)) CROSS.mintTo{value: amount}(pair);
        else info.BASE.safeTransferFrom(_owner, pair, amount);

        IPairV3.Order memory order =
            IPairV3.Order({side: IPairV3.OrderSide.SELL, owner: _owner, feeBps: 0, price: price, amount: amount});
        return _pair.submitLimitOrder(order, constraints, prevPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function submitBuyLimit(
        address pair,
        uint256 price,
        uint256 amount,
        IPairV3.LimitConstraints constraints,
        uint256[2] calldata adjacent,
        uint256 _maxMatchCount
    ) external payable nonReentrant checkSubmit validPair(pair) returns (uint256) {
        address _owner = _msgSender();
        IPairV3 _pair = IPairV3(pair);
        IPairV3.Config memory info = _pair.getConfig();
        uint256 prevPrice = _pair.findPrevPrice(IPairV3.OrderSide.BUY, price, adjacent, findPrevPriceCount);

        {
            uint256 volume = Math.mulDiv(price, amount, info.DENOMINATOR);
            // Use taker fee since limit order can be immediately matched and become taker
            volume = IPairV3(pair).calcBuyVolumeWithFee(volume);

            if (address(info.QUOTE) == address(CROSS)) CROSS.mintTo{value: volume}(pair);
            else info.QUOTE.safeTransferFrom(_owner, pair, volume);
        }

        IPairV3.Order memory order =
            IPairV3.Order({side: IPairV3.OrderSide.BUY, owner: _owner, feeBps: 0, price: price, amount: amount});
        return _pair.submitLimitOrder(order, constraints, prevPrice, _toMaxMatchCount(_maxMatchCount));
    }

    function submitSellMarket(address pair, uint256 amount, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        checkSubmit
        validPair(pair)
    {
        address _owner = _msgSender();
        IPairV3.Config memory info = IPairV3(pair).getConfig();

        IERC20 baseToken = info.BASE;
        if (address(baseToken) == address(CROSS)) CROSS.mintTo{value: amount}(pair);
        else baseToken.safeTransferFrom(_owner, pair, amount);

        IPairV3.Order memory order =
            IPairV3.Order({side: IPairV3.OrderSide.SELL, owner: _owner, feeBps: 0, price: 0, amount: 0});
        IPairV3(pair).submitMarketOrder(order, amount, _toMaxMatchCount(_maxMatchCount));
    }

    function submitBuyMarket(address pair, uint256 quoteVolume, uint256 _maxMatchCount)
        external
        payable
        nonReentrant
        checkSubmit
        validPair(pair)
    {
        address _owner = _msgSender();
        IPairV3.Config memory info = IPairV3(pair).getConfig();

        IERC20 quoteToken = info.QUOTE;
        {
            uint256 volume = IPairV3(pair).calcBuyVolumeWithFee(quoteVolume);
            if (address(quoteToken) == address(CROSS)) CROSS.mintTo{value: volume}(pair);
            else quoteToken.safeTransferFrom(_owner, pair, volume);
        }

        IPairV3.Order memory order =
            IPairV3.Order({side: IPairV3.OrderSide.BUY, owner: _owner, feeBps: 0, price: 0, amount: 0});
        IPairV3(pair).submitMarketOrder(order, quoteVolume, _toMaxMatchCount(_maxMatchCount));
    }

    function cancelOrder(address pair, uint256[] calldata orderIds) external nonReentrant validPair(pair) {
        uint256 length = orderIds.length;
        if (length != 0) {
            if (length > cancelLimit) revert RouterCancelLimitExceeded(length, cancelLimit);
            IPairV3(pair).cancelOrder(_msgSender(), orderIds);
        }
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Public view functions
    // ─────────────────────────────────────────────────────────────────────────────

    function isPair(address pair) public view override returns (bool) {
        return ICrossDexV3(CROSS_DEX).pairToMarket(pair) != address(0);
    }

    function owner() public view override returns (address) {
        return IOwnable(CROSS_DEX).owner();
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Private functions
    // ─────────────────────────────────────────────────────────────────────────────

    function _toMaxMatchCount(uint256 _maxMatchCount) private view returns (uint256) {
        return _maxMatchCount == 0 || _maxMatchCount > maxMatchCount ? maxMatchCount : _maxMatchCount;
    }

    /**
     * @dev Checks if the account has contract code and blocks contract accounts unless whitelisted
     *
     * WARNING: This check can be bypassed by contracts calling the router from within their
     * constructor, as account.code.length is zero during construction.
     */
    function _checkAccountCode(address account) private view {
        if (whitelistedCodeAccounts.contains(account)) return;
        if (account.code.length != 0) revert RouterContractAccountBlocked(account);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Admin functions
    // ─────────────────────────────────────────────────────────────────────────────

    function setFindPrevPriceCount(uint256 _findPrevPriceCount) external onlyOwner {
        if (_findPrevPriceCount == 0) revert RouterInvalidInputData("findPrevPriceCount");
        emit FindPrevPriceCountChanged(findPrevPriceCount, _findPrevPriceCount);
        findPrevPriceCount = _findPrevPriceCount;
    }

    function setMaxMatchCount(uint256 _maxMatchCount) external onlyOwner {
        if (_maxMatchCount == 0) revert RouterInvalidInputData("maxMatchCount");
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
