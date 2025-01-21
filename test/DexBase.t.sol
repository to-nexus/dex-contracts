// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin-contracts-5.2.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin-contracts-5.2.0/utils/math/Math.sol";
import {Test, console} from "forge-std/Test.sol";

import {CrossDexImpl} from "../src/CrossDexImpl.sol";
import {MarketImpl} from "../src/MarketImpl.sol";
import {PairImpl} from "../src/PairImpl.sol";
import {RouterImpl} from "../src/RouterImpl.sol";
import {WCROSS} from "../src/WCROSS.sol";
import {IPair} from "../src/interfaces/IPair.sol";

import {T20} from "./mock/T20.sol";

contract DEXBaseTest is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant FEE_COLLECTOR = address(bytes20("FEE_COLLECTOR"));

    CrossDexImpl public CROSS_DEX;
    RouterImpl public ROUTER;
    WCROSS public WCross;

    IERC20 public QUOTE;
    IERC20 public BASE;

    MarketImpl public MARKET;
    PairImpl public PAIR;

    uint256 public MAX_MATCH_COUNT = type(uint256).max;
    uint256 public MAKER_FEE_PERMIL;
    uint256 public TAKER_FEE_PERMIL;

    uint256 public QUOTE_DECIMALS;
    uint256 public BASE_DECIMALS;

    function _deploy(uint8 quote_decimals, uint8 base_decimals, uint256 quote_tick_size, uint256 base_tick_size)
        internal
    {
        vm.label(OWNER, "owner");
        vm.label(FEE_COLLECTOR, "feeCollector");

        vm.startPrank(OWNER);

        QUOTE_DECIMALS = 10 ** quote_decimals;
        BASE_DECIMALS = 10 ** base_decimals;

        {
            // deploy cross dex
            address crossDexImpl = address(new CrossDexImpl());
            ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
            CROSS_DEX = CrossDexImpl(address(proxy));
            CROSS_DEX.initialize(OWNER, MAX_MATCH_COUNT);
        }
        {
            // get contracts from CROSS_DEX
            ROUTER = RouterImpl(CROSS_DEX.ROUTER());
            WCross = WCROSS(ROUTER.WCross());
        }
        {
            // deploy base and quote tokens
            QUOTE = new T20("QUOTE", "QUOTE", quote_decimals);
            BASE = new T20("BASE", "BASE", base_decimals);
        }
        {
            // create market & Pair
            address market = CROSS_DEX.createMarket(OWNER, FEE_COLLECTOR, address(QUOTE));
            MARKET = MarketImpl(market);
            address pair = MARKET.createPair(
                address(BASE),
                QUOTE_DECIMALS / quote_tick_size,
                BASE_DECIMALS / base_tick_size,
                MAKER_FEE_PERMIL,
                TAKER_FEE_PERMIL
            );
            PAIR = PairImpl(pair);
        }

        vm.stopPrank();
    }

    function _toBase(uint256 x) internal view returns (uint256) {
        return x * BASE_DECIMALS;
    }

    function _toQuote(uint256 x) internal view returns (uint256) {
        return x * QUOTE_DECIMALS;
    }
}
