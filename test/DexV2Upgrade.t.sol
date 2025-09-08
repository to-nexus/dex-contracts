// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./DEXBase.t.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-5.2.0/proxy/utils/UUPSUpgradeable.sol";

import {CrossDexMakerVault} from "../src/CrossDexMakerVault.sol";
import {CrossDexMakerVaultFactory} from "../src/CrossDexMakerVaultFactory.sol";

import {V2CrossDexImpl} from "../src/V2CrossDexImpl.sol";
import {V2CrossDexRouter} from "../src/V2CrossDexRouter.sol";
import {V2MarketImpl} from "../src/V2MarketImpl.sol";
import {V2PairImpl} from "../src/V2PairImpl.sol";

contract DEXTradeTest is DEXBaseTest {
    CrossDexMakerVaultFactory public makerVaultFactory;

    address public v2CrossDexImpl;
    address public v2CrossDexRouter;
    address public v2MarketImpl;
    address public v2PairImpl;

    function setUp() external {
        _deploy(18, 18, 1e2, 1e6);
        // 1. deploy market vault factory
        makerVaultFactory = new CrossDexMakerVaultFactory();
        address _makerVaultFactory = address(makerVaultFactory);

        // 2. deploy v2Logics
        v2CrossDexImpl = address(new V2CrossDexImpl());
        v2CrossDexRouter = address(new V2CrossDexRouter());
        v2MarketImpl = address(new V2MarketImpl());
        v2PairImpl = address(new V2PairImpl());

        // 3. upgrade CrossDexImpl
        vm.prank(OWNER);
        UUPSUpgradeable(address(CROSS_DEX)).upgradeToAndCall(
            v2CrossDexImpl, abi.encodeCall(V2CrossDexImpl.reinitialize, (v2MarketImpl, v2PairImpl))
        );

        // 4. upgrade CrossDexRouter
        vm.prank(OWNER);
        UUPSUpgradeable(address(ROUTER)).upgradeToAndCall(
            v2CrossDexRouter, abi.encodeCall(V2CrossDexRouter.reinitialize, (_makerVaultFactory))
        );

        // 5. upgrade MarketImpl
        vm.prank(OWNER);
        UUPSUpgradeable(address(MARKET)).upgradeToAndCall(v2MarketImpl, abi.encodeCall(V2MarketImpl.reinitialize, ()));

        // 6. upgrade PairImpl
        vm.prank(OWNER);
        UUPSUpgradeable(address(PAIR)).upgradeToAndCall(v2PairImpl, "");
    }

    function test_initial_status() external view {
        // check impl address
        V2CrossDexImpl crossDex = V2CrossDexImpl(address(CROSS_DEX));
        assertEq(v2MarketImpl, crossDex.marketImpl());
        assertEq(v2PairImpl, crossDex.pairImpl());
        // check maker vault factory
        V2CrossDexRouter router = V2CrossDexRouter(address(ROUTER));
        assertEq(address(makerVaultFactory), router.MAKER_VAULT_FACTORY());
        // check maker vault
        address user1 = address(bytes20("user1"));
        address vault = makerVaultFactory.makerVaultAddress(user1);
        assertEq(vault, router.makerVaultAddress(user1));
        assertEq(vault, makerVaultFactory.makerVaultAddress(user1));
    }

    function test_create_vault() external {
        address user1 = makeAddr("user1");
        address vault = makerVaultFactory.makerVaultAddress(user1);

        vm.expectEmit();
        emit CrossDexMakerVaultFactory.Created(user1, vault);
        makerVaultFactory.ensureMakerVault(user1);
    }

    function test_trade_and_check_vault_case1() external {
        // 1. user 1 sell limit order -> user 2 buy market order
        // 2. user1 vault have quote token
        // 3. user1 claim from vault -> user1 have quote token
        address user1 = makeAddr("user1");
        address user1Vault = makerVaultFactory.makerVaultAddress(user1);
        address user2 = makeAddr("user2");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(1000));
        vm.prank(OWNER);
        QUOTE.transfer(user2, _toQuote(1000));

        uint256 price = _toQuote(1);

        vm.prank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(user1);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(1000), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        // check user1, user1Vault balance (0)
        assertEq(BASE.balanceOf(user1), 0);
        assertEq(QUOTE.balanceOf(user1), 0);
        assertEq(BASE.balanceOf(user1Vault), 0);
        assertEq(QUOTE.balanceOf(user1Vault), 0);

        vm.prank(user2);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(user2);
        ROUTER.submitBuyMarket(address(PAIR), _toQuote(1000), 1);
        // check user2 balance (base token)
        assertEq(BASE.balanceOf(user2), _toBase(1000));
        assertEq(QUOTE.balanceOf(user2), 0);

        // check user1Vault balance (quote token)
        assertEq(BASE.balanceOf(user1Vault), 0);
        assertEq(QUOTE.balanceOf(user1Vault), _toQuote(1000));
        // check user1 balance (0)
        assertEq(BASE.balanceOf(user1), 0);
        assertEq(QUOTE.balanceOf(user1), 0);

        vm.prank(user1);
        address[] memory tokens = new address[](2);
        tokens[0] = address(BASE);
        tokens[1] = address(QUOTE);
        CrossDexMakerVault(payable(user1Vault)).claim(tokens);
        // check user1 balance (quote token)
        assertEq(BASE.balanceOf(user1), 0);
        assertEq(QUOTE.balanceOf(user1), _toQuote(1000));
        // check user1Vault balance (0)
        assertEq(BASE.balanceOf(user1Vault), 0);
        assertEq(QUOTE.balanceOf(user1Vault), 0);
    }

    function test_trade_and_check_vault_case2() external {
        // 1. user1,user2,user3,user4,user5 sell limit order -> user6 buy market order
        // 2. user1 vault have quote token
        // 3. user1 claim from vault -> user1 have quote token
        address user1 = makeAddr("user1");
        address user1Vault = makerVaultFactory.makerVaultAddress(user1);
        address user2 = makeAddr("user2");
        address user2Vault = makerVaultFactory.makerVaultAddress(user2);
        address user3 = makeAddr("user3");
        address user3Vault = makerVaultFactory.makerVaultAddress(user3);
        address user4 = makeAddr("user4");
        address user4Vault = makerVaultFactory.makerVaultAddress(user4);
        address user5 = makeAddr("user5");
        address user5Vault = makerVaultFactory.makerVaultAddress(user5);
        address user6 = makeAddr("user6");

        vm.prank(OWNER);
        BASE.transfer(user1, _toBase(1000));
        vm.prank(OWNER);
        BASE.transfer(user2, _toBase(1000));
        vm.prank(OWNER);
        BASE.transfer(user3, _toBase(1000));
        vm.prank(OWNER);
        BASE.transfer(user4, _toBase(1000));
        vm.prank(OWNER);
        BASE.transfer(user5, _toBase(1000));

        vm.prank(OWNER);
        QUOTE.transfer(user6, _toQuote(1000) * 5);

        uint256 price = _toQuote(1);

        vm.prank(user1);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(user1);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(1000), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        // check user1, user1Vault balance (0)
        assertEq(BASE.balanceOf(user1), 0);
        assertEq(QUOTE.balanceOf(user1), 0);
        assertEq(BASE.balanceOf(user1Vault), 0);
        assertEq(QUOTE.balanceOf(user1Vault), 0);

        vm.prank(user2);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(user2);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(1000), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        // check user2, user2Vault balance (0)
        assertEq(BASE.balanceOf(user2), 0);
        assertEq(QUOTE.balanceOf(user2), 0);
        assertEq(BASE.balanceOf(user2Vault), 0);
        assertEq(QUOTE.balanceOf(user2Vault), 0);

        vm.prank(user3);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(user3);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(1000), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        // check user3, user3Vault balance (0)
        assertEq(BASE.balanceOf(user3), 0);
        assertEq(QUOTE.balanceOf(user3), 0);
        assertEq(BASE.balanceOf(user3Vault), 0);
        assertEq(QUOTE.balanceOf(user3Vault), 0);

        vm.prank(user4);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(user4);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(1000), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        // check user4, user4Vault balance (0)
        assertEq(BASE.balanceOf(user4), 0);
        assertEq(QUOTE.balanceOf(user4), 0);
        assertEq(BASE.balanceOf(user4Vault), 0);
        assertEq(QUOTE.balanceOf(user4Vault), 0);

        vm.prank(user5);
        BASE.approve(address(ROUTER), type(uint256).max);
        vm.prank(user5);
        ROUTER.submitSellLimit(
            address(PAIR), price, _toBase(1000), IPair.LimitConstraints.GOOD_TILL_CANCEL, _searchPrices, 0
        );
        // check user5, user5Vault balance (0)
        assertEq(BASE.balanceOf(user5), 0);
        assertEq(QUOTE.balanceOf(user5), 0);
        assertEq(BASE.balanceOf(user5Vault), 0);
        assertEq(QUOTE.balanceOf(user5Vault), 0);

        vm.prank(user6);
        QUOTE.approve(address(ROUTER), type(uint256).max);
        vm.prank(user6);
        ROUTER.submitBuyMarket(address(PAIR), _toQuote(1000) * 5, 10);
        // check user2 balance (base token)
        assertEq(BASE.balanceOf(user6), _toBase(1000) * 5);
        assertEq(QUOTE.balanceOf(user6), 0);

        // check user1Vault balance (quote token)
        assertEq(BASE.balanceOf(user1Vault), 0);
        assertEq(QUOTE.balanceOf(user1Vault), _toQuote(1000));
        // check user1 balance (0)
        assertEq(BASE.balanceOf(user1), 0);
        assertEq(QUOTE.balanceOf(user1), 0);

        address[] memory tokens = new address[](2);
        tokens[0] = address(BASE);
        tokens[1] = address(QUOTE);
        CrossDexMakerVault(payable(user1Vault)).claim(tokens);
        CrossDexMakerVault(payable(user2Vault)).claim(tokens);
        CrossDexMakerVault(payable(user3Vault)).claim(tokens);
        CrossDexMakerVault(payable(user4Vault)).claim(tokens);
        CrossDexMakerVault(payable(user5Vault)).claim(tokens);
        // check user1 balance (quote token)
        assertEq(BASE.balanceOf(user1), 0);
        assertEq(QUOTE.balanceOf(user1), _toQuote(1000));
        // check user1Vault balance (0)
        assertEq(BASE.balanceOf(user1Vault), 0);
        assertEq(QUOTE.balanceOf(user1Vault), 0);

        // check user2 balance (quote token)
        assertEq(BASE.balanceOf(user2), 0);
        assertEq(QUOTE.balanceOf(user2), _toQuote(1000));
        // check user2Vault balance (0)
        assertEq(BASE.balanceOf(user2Vault), 0);
        assertEq(QUOTE.balanceOf(user2Vault), 0);

        // check user3 balance (quote token)
        assertEq(BASE.balanceOf(user3), 0);
        assertEq(QUOTE.balanceOf(user3), _toQuote(1000));
        // check user3Vault balance (0)
        assertEq(BASE.balanceOf(user3Vault), 0);
        assertEq(QUOTE.balanceOf(user3Vault), 0);

        // check user4 balance (quote token)
        assertEq(BASE.balanceOf(user4), 0);
        assertEq(QUOTE.balanceOf(user4), _toQuote(1000));
        // check user4Vault balance (0)
        assertEq(BASE.balanceOf(user4Vault), 0);
        assertEq(QUOTE.balanceOf(user4Vault), 0);

        // check user5 balance (quote token)
        assertEq(BASE.balanceOf(user5), 0);
        assertEq(QUOTE.balanceOf(user5), _toQuote(1000));
        // check user5Vault balance (0)
        assertEq(BASE.balanceOf(user5Vault), 0);
        assertEq(QUOTE.balanceOf(user5Vault), 0);
    }
}
