// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin-contracts-5.5.0/utils/math/Math.sol";
import {Test} from "forge-std/Test.sol";

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {CrossDexRouterV3} from "../src/CrossDexRouterV3.sol";
import {FeeControllerV2Compat} from "../src/FeeControllerV2Compat.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";
import {WETH} from "../src/WETH.sol";

import {BPS_DENOMINATOR} from "../src/interfaces/IFeeController.sol";

import {T20} from "./mock/T20.sol";

contract DEXV3BaseTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));
    address public constant USER1 = address(bytes20("USER1"));
    address public constant USER2 = address(bytes20("USER2"));

    CrossDexImplV3 public CROSS_DEX;
    CrossDexRouterV3 public ROUTER;
    WETH public CROSS;

    IERC20 public QUOTE;
    IERC20 public BASE;

    MarketImplV3 public MARKET;
    PairImplV3 public PAIR;
    FeeControllerV2Compat public FEE_CONTROLLER;

    uint256 public FIND_PREV_PRICE_COUNT = type(uint256).max;
    uint256 public MAX_MATCH_COUNT = type(uint256).max;
    uint256 public CANCEL_LIMIT = type(uint256).max;

    // Default fees (matching V2 pattern: seller fees only by default)
    uint32 public SELLER_MAKER_FEE = 20; // 0.2%
    uint32 public SELLER_TAKER_FEE = 30; // 0.3%
    uint32 public BUYER_MAKER_FEE = 0;
    uint32 public BUYER_TAKER_FEE = 0;

    uint256 public QUOTE_DECIMALS;
    uint256 public BASE_DECIMALS;

    uint256[2] internal _searchPrices;

    function _deployV3(uint8 quote_decimals, uint8 base_decimals, uint256 quote_tick_size, uint256 base_tick_size)
        internal
    {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");
        vm.label(USER1, "user1");
        vm.label(USER2, "user2");

        vm.startPrank(OWNER);

        QUOTE_DECIMALS = 10 ** quote_decimals;
        BASE_DECIMALS = 10 ** base_decimals;

        {
            // Deploy FeeController implementation
            FEE_CONTROLLER = new FeeControllerV2Compat();
        }
        {
            // deploy impl contracts (using V3 versions)
            address routerImpl = address(new CrossDexRouterV3());
            address marketImpl = address(new MarketImplV3());
            address pairImpl = address(new PairImplV3());

            // deploy cross dex
            address crossDexImpl = address(new CrossDexImplV3());
            ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
            CROSS_DEX = CrossDexImplV3(address(proxy));
            CROSS_DEX.initialize(
                OWNER,
                routerImpl,
                FIND_PREV_PRICE_COUNT,
                MAX_MATCH_COUNT,
                CANCEL_LIMIT,
                marketImpl,
                pairImpl,
                address(0) // tickSizeSetter
            );

            // Allow fee controller
            CROSS_DEX.setFeeControllerAllow(address(FEE_CONTROLLER), true);
        }
        {
            // get contracts from CROSS_DEX
            ROUTER = CrossDexRouterV3(CROSS_DEX.ROUTER());
            CROSS = WETH(payable(address(ROUTER.CROSS())));
        }
        {
            // deploy base and quote tokens
            QUOTE = new T20("QUOTE", "QUOTE", quote_decimals);
            BASE = new T20("BASE", "BASE", base_decimals);
        }
        {
            // create market with FeeController
            address market = CROSS_DEX.createMarket(OWNER, address(QUOTE), address(FEE_CONTROLLER), "");
            MARKET = MarketImplV3(market);

            // Encode fee controller init data for pair creation
            bytes memory feeControllerInitData =
                abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);

            // create pair
            address pair = MARKET.createPair(
                address(BASE), QUOTE_DECIMALS / quote_tick_size, BASE_DECIMALS / base_tick_size, feeControllerInitData
            );
            PAIR = PairImplV3(pair);
        }

        // Setup initial balances
        QUOTE.transfer(USER1, _toQuote(50000));
        QUOTE.transfer(USER2, _toQuote(50000));
        BASE.transfer(USER1, _toBase(50000));
        BASE.transfer(USER2, _toBase(50000));

        vm.stopPrank();

        // Setup approvals
        vm.prank(USER1);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER1);
        BASE.approve(address(ROUTER), type(uint256).max);

        vm.prank(USER2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(USER2);
        BASE.approve(address(ROUTER), type(uint256).max);
    }

    function _toBase(uint256 x) internal view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) internal view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }

    function _toTradeVolume(uint256 price, uint256 amount) internal view returns (uint256) {
        return Math.mulDiv(price, amount, BASE_DECIMALS);
    }

    function _calcFee(uint256 amount, uint32 bps) internal pure returns (uint256) {
        return Math.mulDiv(amount, bps, BPS_DENOMINATOR);
    }
}
