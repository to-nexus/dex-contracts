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
    address public pairCreator1;
    address public pairCreator2;
    address public unauthorized;
    address public router;
    address public quote;
    address public feeCollector;
    address public base;

    bytes32 public constant PAIR_CREATOR_ROLE = keccak256("PAIR_CREATOR_ROLE");
    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

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
        pairCreator1 = makeAddr("pairCreator1");
        pairCreator2 = makeAddr("pairCreator2");
        unauthorized = makeAddr("unauthorized");
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

        // Deploy Verse8MarketOwner with multiple pair creators
        address[] memory creators = new address[](2);
        creators[0] = pairCreator1;
        creators[1] = pairCreator2;
        verse8Owner = new Verse8MarketOwner(owner, creators);

        // Transfer market ownership to Verse8MarketOwner
        market.transferOwnership(address(verse8Owner));

        // Mock tokens
        vm.etch(quote, hex"00");
        vm.etch(base, hex"00");
        vm.mockCall(base, abi.encodeWithSignature("decimals()"), abi.encode(uint8(18)));
    }

    // Test initial state and role assignments
    function test_InitialRoleAssignments() public view {
        // Check that owner has DEFAULT_ADMIN_ROLE
        assertTrue(verse8Owner.hasRole(DEFAULT_ADMIN_ROLE, owner));

        // Check that owner has PAIR_CREATOR_ROLE
        assertTrue(verse8Owner.hasRole(PAIR_CREATOR_ROLE, owner));

        // Check that pairCreators have PAIR_CREATOR_ROLE
        assertTrue(verse8Owner.hasRole(PAIR_CREATOR_ROLE, pairCreator1));
        assertTrue(verse8Owner.hasRole(PAIR_CREATOR_ROLE, pairCreator2));

        // Check that unauthorized user has no roles
        assertFalse(verse8Owner.hasRole(DEFAULT_ADMIN_ROLE, unauthorized));
        assertFalse(verse8Owner.hasRole(PAIR_CREATOR_ROLE, unauthorized));
    }

    // Test createPair function with PAIR_CREATOR_ROLE
    function test_OwnerCanCreatePair() public {
        vm.startPrank(owner);
        address pair = verse8Owner.createPair(address(market), base, 1e18, 1e18);
        assertNotEq(pair, address(0));
        vm.stopPrank();
    }

    function test_PairCreator1CanCreatePair() public {
        vm.startPrank(pairCreator1);
        address pair = verse8Owner.createPair(address(market), base, 1e18, 1e18);
        assertNotEq(pair, address(0));
        vm.stopPrank();
    }

    function test_PairCreator2CanCreatePair() public {
        vm.startPrank(pairCreator2);
        address pair = verse8Owner.createPair(address(market), base, 1e18, 1e18);
        assertNotEq(pair, address(0));
        vm.stopPrank();
    }

    function test_RevertWhen_UnauthorizedTriesToCreatePair() public {
        vm.startPrank(unauthorized);
        vm.expectRevert();
        verse8Owner.createPair(address(market), base, 1e18, 1e18);
        vm.stopPrank();
    }

    // Test createPairs batch function
    function test_PairCreatorCanCreateMultiplePairs() public {
        address base2 = makeAddr("base2");
        vm.etch(base2, hex"00");
        vm.mockCall(base2, abi.encodeWithSignature("decimals()"), abi.encode(uint8(18)));

        Verse8MarketOwner.CreatePairArgs[] memory args = new Verse8MarketOwner.CreatePairArgs[](2);
        args[0] = Verse8MarketOwner.CreatePairArgs({market: address(market), base: base, tickSize: 1e18, lotSize: 1e18});
        args[1] =
            Verse8MarketOwner.CreatePairArgs({market: address(market), base: base2, tickSize: 2e18, lotSize: 2e18});

        vm.startPrank(pairCreator1);
        address[] memory pairs = verse8Owner.createPairs(args);
        assertEq(pairs.length, 2);
        assertNotEq(pairs[0], address(0));
        assertNotEq(pairs[1], address(0));
        vm.stopPrank();
    }

    // Test setFeeCollector function through execute
    function test_OwnerCanSetFeeCollector() public {
        address newFeeCollector = makeAddr("newFeeCollector");
        bytes memory data = abi.encodeCall(IMarket.setFeeCollector, (newFeeCollector));

        vm.startPrank(owner);
        verse8Owner.execute(address(market), 0, data);
        vm.stopPrank();

        assertEq(market.feeCollector(), newFeeCollector);
    }

    function test_RevertWhen_PairCreatorTriesToSetFeeCollector() public {
        address newFeeCollector = makeAddr("newFeeCollector");
        bytes memory data = abi.encodeCall(IMarket.setFeeCollector, (newFeeCollector));

        vm.startPrank(pairCreator1);
        vm.expectRevert();
        verse8Owner.execute(address(market), 0, data);
        vm.stopPrank();
    }

    // Test setFeeBps function through execute
    function test_OwnerCanSetFeeBps() public {
        uint256 newFeeBps = 50;
        bytes memory data = abi.encodeCall(IMarket.setFeeBps, (newFeeBps));

        vm.startPrank(owner);
        verse8Owner.execute(address(market), 0, data);
        vm.stopPrank();

        assertEq(market.feeBps(), newFeeBps);
    }

    function test_RevertWhen_PairCreatorTriesToSetFeeBps() public {
        uint256 newFeeBps = 50;
        bytes memory data = abi.encodeCall(IMarket.setFeeBps, (newFeeBps));

        vm.startPrank(pairCreator1);
        vm.expectRevert();
        verse8Owner.execute(address(market), 0, data);
        vm.stopPrank();
    }

    // Test that only owner can call administrative functions
    function test_OwnerCanCallAdministrativeFunctions() public {
        // Test that owner can call any function through execute
        address newFeeCollector = makeAddr("newFeeCollector");
        bytes memory data = abi.encodeCall(IMarket.setFeeCollector, (newFeeCollector));

        vm.startPrank(owner);
        bytes memory result = verse8Owner.execute(address(market), 0, data);
        vm.stopPrank();

        assertEq(market.feeCollector(), newFeeCollector);
    }

    function test_RevertWhen_PairCreatorTriesToCallAdministrativeFunctions() public {
        address newFeeCollector = makeAddr("newFeeCollector");
        bytes memory data = abi.encodeCall(IMarket.setFeeCollector, (newFeeCollector));

        vm.startPrank(pairCreator1);
        vm.expectRevert();
        verse8Owner.execute(address(market), 0, data);
        vm.stopPrank();
    }

    // Test executeBatch function
    function test_OwnerCanExecuteBatch() public {
        address newFeeCollector = makeAddr("newFeeCollector");
        uint256 newFeeBps = 75;

        Verse8MarketOwner.ExecuteBatchArgs[] memory calls = new Verse8MarketOwner.ExecuteBatchArgs[](2);
        calls[0] = Verse8MarketOwner.ExecuteBatchArgs({
            to: address(market),
            value: 0,
            data: abi.encodeCall(IMarket.setFeeCollector, (newFeeCollector))
        });
        calls[1] = Verse8MarketOwner.ExecuteBatchArgs({
            to: address(market),
            value: 0,
            data: abi.encodeCall(IMarket.setFeeBps, (newFeeBps))
        });

        vm.startPrank(owner);
        bytes[] memory results = verse8Owner.executeBatch(calls);
        vm.stopPrank();

        assertEq(results.length, 2);
        assertEq(market.feeCollector(), newFeeCollector);
        assertEq(market.feeBps(), newFeeBps);
    }

    function test_RevertWhen_PairCreatorTriesToExecuteBatch() public {
        Verse8MarketOwner.ExecuteBatchArgs[] memory calls = new Verse8MarketOwner.ExecuteBatchArgs[](1);
        calls[0] = Verse8MarketOwner.ExecuteBatchArgs({
            to: address(market),
            value: 0,
            data: abi.encodeCall(IMarket.setFeeBps, (100))
        });

        vm.startPrank(pairCreator1);
        vm.expectRevert();
        verse8Owner.executeBatch(calls);
        vm.stopPrank();
    }

    function test_RevertWhen_UnauthorizedTriesToExecute() public {
        bytes memory data = abi.encodeCall(IMarket.setFeeBps, (100));

        vm.startPrank(unauthorized);
        vm.expectRevert();
        verse8Owner.execute(address(market), 0, data);
        vm.stopPrank();
    }

    // Test that execute reverts on failed calls
    function test_RevertWhen_ExecuteCallFails() public {
        // Try to call a non-existent function
        bytes memory invalidData = abi.encodeWithSignature("nonExistentFunction()");

        vm.startPrank(owner);
        vm.expectRevert(abi.encodeWithSelector(Verse8MarketOwner.Verse8MarketOwner__CallFailed.selector));
        verse8Owner.execute(address(market), 0, invalidData);
        vm.stopPrank();
    }

    // Test role management
    function test_OwnerCanGrantPairCreatorRole() public {
        address newPairCreator = makeAddr("newPairCreator");

        vm.startPrank(owner);
        verse8Owner.grantRole(PAIR_CREATOR_ROLE, newPairCreator);
        vm.stopPrank();

        assertTrue(verse8Owner.hasRole(PAIR_CREATOR_ROLE, newPairCreator));

        // New pair creator should be able to create pairs
        vm.startPrank(newPairCreator);
        address pair = verse8Owner.createPair(address(market), base, 1e18, 1e18);
        assertNotEq(pair, address(0));
        vm.stopPrank();
    }

    function test_OwnerCanRevokePairCreatorRole() public {
        vm.startPrank(owner);
        verse8Owner.revokeRole(PAIR_CREATOR_ROLE, pairCreator1);
        vm.stopPrank();

        assertFalse(verse8Owner.hasRole(PAIR_CREATOR_ROLE, pairCreator1));

        // Revoked pair creator should not be able to create pairs
        vm.startPrank(pairCreator1);
        vm.expectRevert();
        verse8Owner.createPair(address(market), base, 1e18, 1e18);
        vm.stopPrank();
    }

    function test_RevertWhen_PairCreatorTriesToGrantRole() public {
        address newPairCreator = makeAddr("newPairCreator");

        vm.startPrank(pairCreator1);
        vm.expectRevert();
        verse8Owner.grantRole(PAIR_CREATOR_ROLE, newPairCreator);
        vm.stopPrank();
    }

    // Test comprehensive access control scenarios
    function test_ComprehensiveAccessControl() public {
        // 1. PAIR_CREATOR_ROLE can createPair but cannot execute
        vm.startPrank(pairCreator1);

        // Should succeed: createPair
        address pair = verse8Owner.createPair(address(market), base, 1e18, 1e18);
        assertNotEq(pair, address(0));

        // Should fail: execute
        bytes memory data = abi.encodeCall(IMarket.setFeeBps, (100));
        vm.expectRevert();
        verse8Owner.execute(address(market), 0, data);

        vm.stopPrank();

        // 2. DEFAULT_ADMIN_ROLE (owner) can do everything
        vm.startPrank(owner);

        // Should succeed: createPair
        address base2 = makeAddr("base2");
        vm.etch(base2, hex"00");
        vm.mockCall(base2, abi.encodeWithSignature("decimals()"), abi.encode(uint8(18)));
        address pair2 = verse8Owner.createPair(address(market), base2, 1e18, 1e18);
        assertNotEq(pair2, address(0));

        // Should succeed: execute
        verse8Owner.execute(address(market), 0, data);
        assertEq(market.feeBps(), 100);

        vm.stopPrank();

        // 3. Unauthorized user cannot do anything
        vm.startPrank(unauthorized);

        // Should fail: createPair
        vm.expectRevert();
        verse8Owner.createPair(address(market), base, 1e18, 1e18);

        // Should fail: execute
        vm.expectRevert();
        verse8Owner.execute(address(market), 0, data);

        vm.stopPrank();
    }
}
