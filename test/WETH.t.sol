// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin-contracts-5.5.0/token/ERC20/IERC20.sol";
import {Test} from "forge-std/Test.sol";

import {WETH} from "../src/WETH.sol";
import {DEXV3BaseTest} from "./DEXV3Base.t.sol";
import {T20} from "./mock/T20.sol";

contract WETHTest is DEXV3BaseTest {
    function setUp() external {
        _deployV3(18, 18, 1e2, 1e6);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // View Functions Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_name() external view {
        assertEq(CROSS.name(), "CrossDEX Wrapped CROSS");
    }

    function test_symbol() external view {
        assertEq(CROSS.symbol(), "CROSS");
    }

    function test_decimals() external view {
        assertEq(CROSS.decimals(), 18);
    }

    function test_ROUTER() external view {
        assertEq(CROSS.ROUTER(), address(ROUTER));
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Minting Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_receive_mintsTokens() external {
        uint256 amount = 1 ether;

        // USER1 is not a Pair, so WETH will be auto-unwrapped and ETH returned
        uint256 user1EthBefore = USER1.balance;
        vm.deal(USER1, amount);
        vm.prank(USER1);
        (bool success,) = address(CROSS).call{value: amount}("");
        assertTrue(success);

        // Verify: ETH should be returned to sender (auto-unwrap)
        assertEq(USER1.balance, user1EthBefore + amount, "ETH should be returned to sender");
        assertEq(CROSS.balanceOf(USER1), 0, "WETH balance should remain 0");
    }

    function test_mintTo_mintsToRecipient() external {
        uint256 amount = 1 ether;

        vm.deal(USER1, amount);
        vm.prank(USER1);
        CROSS.mintTo{value: amount}(address(PAIR));

        // Pair should hold the WETH
        assertEq(CROSS.balanceOf(address(PAIR)), amount);
    }

    function test_mintTo_exactAmount() external {
        uint256 amount = 123456789012345678;

        vm.deal(USER1, amount);
        vm.prank(USER1);
        CROSS.mintTo{value: amount}(address(PAIR));

        assertEq(CROSS.balanceOf(address(PAIR)), amount);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Auto-Unwrap Behavior Tests
    // ═══════════════════════════════════════════════════════════════════════════

    function test_transfer_toPair_noUnwrap() external {
        uint256 amount = 1 ether;

        // Mint WETH to PAIR
        vm.deal(USER1, amount);
        vm.prank(USER1);
        CROSS.mintTo{value: amount}(address(PAIR));

        // Create a second pair with a DIFFERENT base token
        vm.startPrank(OWNER);
        T20 base2 = new T20("BASE2", "BASE2", 18);
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);
        address pair2 = MARKET.createPair(address(base2), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
        vm.stopPrank();

        // Transfer from PAIR to another pair - should NOT unwrap
        uint256 pair2BalanceBefore = CROSS.balanceOf(pair2);

        vm.prank(address(PAIR));
        CROSS.transfer(pair2, amount);

        // pair2 should hold WETH (no unwrap)
        assertEq(CROSS.balanceOf(pair2), pair2BalanceBefore + amount);
    }

    function test_transfer_toNonPair_unwrapsAutomatically() external {
        uint256 amount = 1 ether;

        // Mint WETH to PAIR
        vm.deal(USER1, amount);
        vm.prank(USER1);
        CROSS.mintTo{value: amount}(address(PAIR));

        // Transfer from PAIR to EOA - should unwrap
        uint256 user2EthBefore = USER2.balance;
        uint256 user2WethBefore = CROSS.balanceOf(USER2);

        vm.prank(address(PAIR));
        CROSS.transfer(USER2, amount);

        // USER2 should receive native ETH, not WETH
        assertEq(CROSS.balanceOf(USER2), user2WethBefore);
        assertEq(USER2.balance, user2EthBefore + amount);
    }

    function test_transfer_toEOA_unwrapsAndSendsETH() external {
        uint256 amount = 0.5 ether;

        // Mint WETH to PAIR
        vm.deal(USER1, amount);
        vm.prank(USER1);
        CROSS.mintTo{value: amount}(address(PAIR));

        uint256 user1EthBefore = USER1.balance;

        vm.prank(address(PAIR));
        CROSS.transfer(USER1, amount);

        assertEq(USER1.balance, user1EthBefore + amount);
        assertEq(CROSS.balanceOf(USER1), 0);
    }

    function test_transferFrom_toPair_noUnwrap() external {
        uint256 amount = 1 ether;

        // Mint WETH to PAIR
        vm.deal(USER1, amount);
        vm.prank(USER1);
        CROSS.mintTo{value: amount}(address(PAIR));

        // Create second pair with a DIFFERENT base token
        vm.startPrank(OWNER);
        T20 base2 = new T20("BASE2", "BASE2", 18);
        bytes memory feeData =
            abi.encode(FEE_COLLECTOR, SELLER_MAKER_FEE, SELLER_TAKER_FEE, BUYER_MAKER_FEE, BUYER_TAKER_FEE);
        address pair2 = MARKET.createPair(address(base2), QUOTE_DECIMALS / 1e2, BASE_DECIMALS / 1e6, feeData);
        vm.stopPrank();

        // Approve and transferFrom
        vm.prank(address(PAIR));
        CROSS.approve(USER1, amount);

        vm.prank(USER1);
        CROSS.transferFrom(address(PAIR), pair2, amount);

        // pair2 should hold WETH
        assertEq(CROSS.balanceOf(pair2), amount);
    }

    function test_transferFrom_toNonPair_unwraps() external {
        uint256 amount = 1 ether;

        // Mint WETH to PAIR
        vm.deal(USER1, amount);
        vm.prank(USER1);
        CROSS.mintTo{value: amount}(address(PAIR));

        // Approve and transferFrom to EOA
        vm.prank(address(PAIR));
        CROSS.approve(USER1, amount);

        uint256 user2EthBefore = USER2.balance;

        vm.prank(USER1);
        CROSS.transferFrom(address(PAIR), USER2, amount);

        // USER2 should receive native ETH
        assertEq(USER2.balance, user2EthBefore + amount);
        assertEq(CROSS.balanceOf(USER2), 0);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Edge Cases
    // ═══════════════════════════════════════════════════════════════════════════

    function test_transfer_zeroAmount_succeeds() external {
        // Should not revert
        vm.prank(address(PAIR));
        CROSS.transfer(USER1, 0);
    }

    function test_totalSupply_tracksCorrectly() external {
        uint256 amount1 = 1 ether;
        uint256 amount2 = 0.5 ether;

        // Mint to PAIR
        vm.deal(USER1, amount1);
        vm.prank(USER1);
        CROSS.mintTo{value: amount1}(address(PAIR));

        assertEq(CROSS.totalSupply(), amount1);

        // Mint more to PAIR
        vm.deal(USER1, amount2);
        vm.prank(USER1);
        CROSS.mintTo{value: amount2}(address(PAIR));

        assertEq(CROSS.totalSupply(), amount1 + amount2);

        // Transfer to EOA (auto-unwrap/burn)
        vm.prank(address(PAIR));
        CROSS.transfer(USER2, amount1);

        // Supply should decrease after burn
        assertEq(CROSS.totalSupply(), amount2);
    }

    function test_mintTo_toNonPair_immediatelyUnwraps() external {
        uint256 amount = 1 ether;

        uint256 user1EthBefore = USER1.balance;

        vm.deal(USER2, amount);
        vm.prank(USER2);
        CROSS.mintTo{value: amount}(USER1);

        // Should be immediately unwrapped
        assertEq(CROSS.balanceOf(USER1), 0);
        assertEq(USER1.balance, user1EthBefore + amount);
    }

    function test_receive_toNonPair_immediatelyUnwraps() external {
        uint256 amount = 1 ether;

        vm.deal(USER1, amount);
        vm.prank(USER1);
        (bool success,) = address(CROSS).call{value: amount}("");
        assertTrue(success);

        // Should be immediately unwrapped back to USER1
        // The WETH is minted to msg.sender (USER1), then since USER1 is not a Pair,
        // it's burned and ETH is sent back to USER1
        assertEq(CROSS.balanceOf(USER1), 0);
        assertEq(USER1.balance, amount);
    }
}
