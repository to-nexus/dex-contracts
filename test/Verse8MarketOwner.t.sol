// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {Test} from "forge-std/Test.sol";

import {MarketImpl} from "../src/MarketImpl.sol";
import {PairImpl} from "../src/PairImpl.sol";
import {Verse8MarketOwner} from "../src/Verse8MarketOwner.sol";
import {IMarket} from "../src/interfaces/IMarket.sol";

contract Verse8MarketOwnerTest is Test {
    Verse8MarketOwner public verse8Owner;
    MarketImpl public marketImpl;
    MarketImpl public market;
    PairImpl public pairImpl;

    address public owner;
    address public pairCreator;
    address public router;
    address public quote;
    address public feeCollector;
    address public base;

    // Mock CrossDex functions
    mapping(address => address) public pairToMarket;

    function pairCreated(address pair) external {
        pairToMarket[pair] = msg.sender;
    }

    function checkTickSizeRoles(address) external view {
        // Mock: allow all for testing
    }

    function setUp() public {
        owner = makeAddr("owner");
        pairCreator = makeAddr("pairCreator");
        router = makeAddr("router");
        quote = makeAddr("quote");
        feeCollector = makeAddr("feeCollector");
        base = makeAddr("base");

        // Deploy implementations
        marketImpl = new MarketImpl();
        pairImpl = new PairImpl();

        // Deploy Market proxy with deployer as temporary owner
        bytes memory initData =
            abi.encodeCall(MarketImpl.initialize, (address(this), router, quote, address(pairImpl), feeCollector, 30));
        ERC1967Proxy marketProxy = new ERC1967Proxy(address(marketImpl), initData);
        market = MarketImpl(address(marketProxy));

        // Deploy Verse8MarketOwner
        verse8Owner = new Verse8MarketOwner(owner, address(market), pairCreator);

        // Transfer market ownership to Verse8MarketOwner
        market.transferOwnership(address(verse8Owner));

        // Mock tokens
        vm.etch(quote, hex"00");
        vm.etch(base, hex"00");
        vm.mockCall(base, abi.encodeWithSignature("decimals()"), abi.encode(uint8(18)));
    }

    function test_InitialState() public view {
        assertEq(verse8Owner.owner(), owner);
        assertEq(verse8Owner.pairCreator(), pairCreator);
        assertEq(verse8Owner.market(), address(market));
        assertEq(market.owner(), address(verse8Owner));
    }

    function test_PairCreatorCanCreatePair() public {
        vm.startPrank(pairCreator);

        address pair = verse8Owner.createPair(base, 1e18, 1e18);

        assertNotEq(pair, address(0));
        vm.stopPrank();
    }

    function test_OwnerCanCreatePair() public {
        vm.startPrank(owner);

        address pair = verse8Owner.createPair(base, 1e18, 1e18);

        assertNotEq(pair, address(0));
        vm.stopPrank();
    }

    function test_RevertWhen_UnauthorizedTriesToCreatePair() public {
        address unauthorized = makeAddr("unauthorized");

        vm.startPrank(unauthorized);
        vm.expectRevert(
            abi.encodeWithSelector(Verse8MarketOwner.Verse8MarketOwnerUnauthorizedPairCreator.selector, unauthorized)
        );
        verse8Owner.createPair(base, 1e18, 1e18);
        vm.stopPrank();
    }

    function test_OwnerCanSetFeeCollector() public {
        address newFeeCollector = makeAddr("newFeeCollector");

        vm.startPrank(owner);
        verse8Owner.setFeeCollector(newFeeCollector);
        vm.stopPrank();

        assertEq(market.feeCollector(), newFeeCollector);
    }

    function test_RevertWhen_PairCreatorTriesToSetFeeCollector() public {
        address newFeeCollector = makeAddr("newFeeCollector");

        vm.startPrank(pairCreator);
        vm.expectRevert();
        verse8Owner.setFeeCollector(newFeeCollector);
        vm.stopPrank();
    }

    function test_OwnerCanSetFeeBps() public {
        uint256 newFeeBps = 50;

        vm.startPrank(owner);
        verse8Owner.setFeeBps(newFeeBps);
        vm.stopPrank();

        assertEq(market.feeBps(), newFeeBps);
    }

    function test_RevertWhen_PairCreatorTriesToSetFeeBps() public {
        uint256 newFeeBps = 50;

        vm.startPrank(pairCreator);
        vm.expectRevert();
        verse8Owner.setFeeBps(newFeeBps);
        vm.stopPrank();
    }

    function test_OwnerCanChangePairCreator() public {
        address newPairCreator = makeAddr("newPairCreator");

        vm.startPrank(owner);
        verse8Owner.setPairCreator(newPairCreator);
        vm.stopPrank();

        assertEq(verse8Owner.pairCreator(), newPairCreator);

        // New pair creator can create pairs
        vm.startPrank(newPairCreator);
        address newBase = makeAddr("newBase");
        vm.etch(newBase, hex"00");
        vm.mockCall(newBase, abi.encodeWithSignature("decimals()"), abi.encode(uint8(18)));
        address pair = verse8Owner.createPair(newBase, 1e18, 1e18);
        assertNotEq(pair, address(0));
        vm.stopPrank();
    }

    function test_CanCreatePairView() public {
        assertTrue(verse8Owner.canCreatePair(owner));
        assertTrue(verse8Owner.canCreatePair(pairCreator));
        address random = makeAddr("random");
        assertFalse(verse8Owner.canCreatePair(random));
    }

    function test_OwnerCanUpgradeMarket() public {
        MarketImpl newMarketImpl = new MarketImpl();

        vm.startPrank(owner);
        verse8Owner.upgradeMarket(address(newMarketImpl));
        vm.stopPrank();

        // Market should still be functional after upgrade
        assertEq(market.owner(), address(verse8Owner));
    }

    function test_OwnerCanTransferMarketOwnership() public {
        address newMarketOwner = makeAddr("newMarketOwner");

        vm.startPrank(owner);
        verse8Owner.transferMarketOwnership(newMarketOwner);
        vm.stopPrank();

        assertEq(market.owner(), newMarketOwner);
    }

    function test_RevertWhen_PairCreatorTriesToTransferMarketOwnership() public {
        address newMarketOwner = makeAddr("newMarketOwner");

        vm.startPrank(pairCreator);
        vm.expectRevert();
        verse8Owner.transferMarketOwnership(newMarketOwner);
        vm.stopPrank();
    }

    function test_OwnerCanTransferOwnership() public {
        address newOwner = makeAddr("newOwner");

        vm.startPrank(owner);
        verse8Owner.transferOwnership(newOwner);
        vm.stopPrank();

        assertEq(verse8Owner.owner(), newOwner);
    }

    function test_RevertWhen_PairCreatorTriesToTransferOwnership() public {
        address newOwner = makeAddr("newOwner");

        vm.startPrank(pairCreator);
        vm.expectRevert();
        verse8Owner.transferOwnership(newOwner);
        vm.stopPrank();
    }

    function test_OwnerCanSetMarket() public {
        // Deploy new market
        MarketImpl newMarketImpl = new MarketImpl();
        bytes memory initData =
            abi.encodeCall(MarketImpl.initialize, (address(this), router, quote, address(pairImpl), feeCollector, 30));
        ERC1967Proxy newMarketProxy = new ERC1967Proxy(address(newMarketImpl), initData);
        address newMarket = address(newMarketProxy);

        vm.startPrank(owner);
        vm.expectEmit(true, true, false, false);
        emit Verse8MarketOwner.MarketSet(address(market), newMarket);
        verse8Owner.setMarket(newMarket);
        vm.stopPrank();

        assertEq(verse8Owner.market(), newMarket);
    }

    function test_RevertWhen_PairCreatorTriesToSetMarket() public {
        address newMarket = makeAddr("newMarket");

        vm.startPrank(pairCreator);
        vm.expectRevert();
        verse8Owner.setMarket(newMarket);
        vm.stopPrank();
    }

    function test_RevertWhen_UnauthorizedTriesToSetMarket() public {
        address unauthorized = makeAddr("unauthorized");
        address newMarket = makeAddr("newMarket");

        vm.startPrank(unauthorized);
        vm.expectRevert();
        verse8Owner.setMarket(newMarket);
        vm.stopPrank();
    }

    function test_RevertWhen_SetMarketWithZeroAddress() public {
        vm.startPrank(owner);
        vm.expectRevert(
            abi.encodeWithSelector(Verse8MarketOwner.Verse8MarketOwnerInvalidMarketAddress.selector, address(0))
        );
        verse8Owner.setMarket(address(0));
        vm.stopPrank();
    }
}
