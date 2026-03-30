// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.5.0/proxy/ERC1967/ERC1967Proxy.sol";
import {Script, console} from "forge-std/Script.sol";

import {CrossDexImplV3} from "../src/CrossDexImplV3.sol";
import {CrossDexRouterV3} from "../src/CrossDexRouterV3.sol";
import {FeeControllerV2Compat} from "../src/FeeControllerV2Compat.sol";
import {FeeControllerV3Dist} from "../src/FeeControllerV3Dist.sol";
import {FeeControllerV3Split} from "../src/FeeControllerV3Split.sol";
import {MarketImplV3} from "../src/MarketImplV3.sol";
import {PairImplV3} from "../src/PairImplV3.sol";

/// @title DeployV3
/// @notice Deployment script for V3 contracts
/// @dev Usage:
///   1. Deploy implementations: forge script script/DeployV3.s.sol:DeployV3 --sig "deployImplementations()" --broadcast
///   2. Deploy proxy: forge script script/DeployV3.s.sol:DeployV3 --sig "deployProxy(address,address,address,address,address)" <args> --broadcast
contract DeployV3 is Script {
    // ─────────────────────────────────────────────────────────────────────────────
    // Structs
    // ─────────────────────────────────────────────────────────────────────────────

    struct Implementations {
        address crossDexImpl;
        address routerImpl;
        address marketImpl;
        address pairImpl;
        address feeControllerV2Compat;
        address feeControllerV3Dist;
        address feeControllerV3Split;
    }

    struct ProxyDeployment {
        address crossDex;
        address router;
    }

    struct InitParams {
        uint256 findPrevPriceCount;
        uint256 maxMatchCount;
        uint256 cancelLimit;
        address tickSizeSetter;
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Deploy Implementations
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Deploy all implementation contracts
    /// @return impls Struct containing all implementation addresses
    function deployImplementations() external returns (Implementations memory impls) {
        vm.startBroadcast();

        impls = _deployImplementations();

        vm.stopBroadcast();

        _logImplementations(impls);
    }

    function upgradeCrossDex(
        address crossDex,
        address router,
        address crossDexImpl,
        address routerImpl,
        address marketImpl,
        address pairImpl,
        address[] memory feeControllers
    ) external {
        vm.startBroadcast();
        CrossDexImplV3(crossDex)
            .upgradeToAndCall(
                crossDexImpl, abi.encodeCall(CrossDexImplV3.reInitialize, (marketImpl, pairImpl, feeControllers))
            );
        CrossDexRouterV3(router).upgradeToAndCall(routerImpl, abi.encodeCall(CrossDexRouterV3.reInitialize, ()));
        vm.stopBroadcast();
    }

    function upgradeMarket(
        address market,
        address marketImpl,
        address pairImpl,
        address feeController,
        bytes memory feeControllerInitData
    ) external {
        vm.startBroadcast();
        MarketImplV3(market)
            .upgradeToAndCall(marketImpl, abi.encodeCall(MarketImplV3.reInitialize, (pairImpl, feeController)));
        (, address[] memory pairs) = MarketImplV3(market).allPairs();
        for (uint256 i = 0; i < pairs.length; ++i) {
            address pair = pairs[i];
            PairImplV3(pair)
                .upgradeToAndCall(
                    pairImpl, abi.encodeCall(PairImplV3.reInitialize, (feeController, feeControllerInitData))
                );
        }
        vm.stopBroadcast();
    }

    function _deployImplementations() internal returns (Implementations memory impls) {
        impls.routerImpl = address(new CrossDexRouterV3());
        impls.marketImpl = address(new MarketImplV3());
        impls.pairImpl = address(new PairImplV3());
        impls.crossDexImpl = address(new CrossDexImplV3());
        impls.feeControllerV2Compat = address(new FeeControllerV2Compat());
        impls.feeControllerV3Dist = address(new FeeControllerV3Dist());
        impls.feeControllerV3Split = address(new FeeControllerV3Split());
    }

    function _logImplementations(Implementations memory impls) internal pure {
        console.log("=== Implementation Deployment ===");
        console.log("CrossDexImplV3 impl:", impls.crossDexImpl);
        console.log("CrossDexRouterV3 impl:", impls.routerImpl);
        console.log("MarketImplV3 impl:", impls.marketImpl);
        console.log("PairImplV3 impl:", impls.pairImpl);
        console.log("FeeControllerV2Compat:", impls.feeControllerV2Compat);
        console.log("FeeControllerV3Dist:", impls.feeControllerV3Dist);
        console.log("FeeControllerV3Split:", impls.feeControllerV3Split);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Deploy Proxy
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Deploy CrossDex proxy and initialize the system
    /// @param crossDexImpl CrossDexImplV3 implementation address
    /// @param routerImpl CrossDexRouterV3 implementation address
    /// @param marketImpl MarketImplV3 implementation address
    /// @param pairImpl PairImplV3 implementation address
    /// @param feeController FeeController implementation address to allow
    /// @return deployment Struct containing proxy addresses
    function deployProxy(
        address crossDexImpl,
        address routerImpl,
        address marketImpl,
        address pairImpl,
        address feeController
    ) external returns (ProxyDeployment memory deployment) {
        InitParams memory params = _loadInitParams();

        vm.startBroadcast();

        deployment = _deployProxy(crossDexImpl, routerImpl, marketImpl, pairImpl, feeController, params);

        vm.stopBroadcast();

        _logProxyDeployment(deployment, feeController);
    }

    function _loadInitParams() internal view returns (InitParams memory params) {
        params.findPrevPriceCount = vm.envOr("FIND_PREV_PRICE_COUNT", type(uint256).max);
        params.maxMatchCount = vm.envOr("MAX_MATCH_COUNT", type(uint256).max);
        params.cancelLimit = vm.envOr("CANCEL_LIMIT", type(uint256).max);
        params.tickSizeSetter = vm.envOr("TICK_SIZE_SETTER", address(0));
    }

    function _deployProxy(
        address crossDexImpl,
        address routerImpl,
        address marketImpl,
        address pairImpl,
        address feeController,
        InitParams memory params
    ) internal returns (ProxyDeployment memory deployment) {
        // Deploy CrossDex proxy
        ERC1967Proxy proxy = new ERC1967Proxy(crossDexImpl, hex"");
        CrossDexImplV3 crossDex = CrossDexImplV3(address(proxy));

        // Initialize
        crossDex.initialize(
            msg.sender,
            routerImpl,
            params.findPrevPriceCount,
            params.maxMatchCount,
            params.cancelLimit,
            marketImpl,
            pairImpl,
            params.tickSizeSetter
        );

        // Allow fee controller
        if (feeController != address(0)) crossDex.setFeeControllerAllow(feeController, true);

        deployment.crossDex = address(crossDex);
        deployment.router = crossDex.ROUTER();
    }

    function _logProxyDeployment(ProxyDeployment memory deployment, address feeController) internal pure {
        console.log("=== Proxy Deployment ===");
        console.log("CrossDex proxy:", deployment.crossDex);
        console.log("Router proxy:", deployment.router);
        if (feeController != address(0)) console.log("FeeController allowed:", feeController);
        console.log("");
        console.log("Next steps:");
        console.log("1. Create Market: CrossDex.createMarket(owner, quote, feeController, message)");
        console.log("2. Create Pair: Market.createPair(base, tickSize, lotSize, feeInitData)");
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Helper: Deploy All
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Deploy everything in one transaction (implementations + proxy)
    /// @param useFeeControllerV3Split If true, use V3Split fee controller; otherwise use V2Compat
    /// @return impls Implementation addresses
    /// @return deployment Proxy addresses
    function deployAll(bool useFeeControllerV3Split)
        external
        returns (Implementations memory impls, ProxyDeployment memory deployment)
    {
        vm.startBroadcast();

        impls = _deployImplementations();

        address feeController = useFeeControllerV3Split ? impls.feeControllerV3Split : impls.feeControllerV2Compat;

        InitParams memory params = InitParams({
            findPrevPriceCount: type(uint256).max,
            maxMatchCount: type(uint256).max,
            cancelLimit: type(uint256).max,
            tickSizeSetter: address(0)
        });

        deployment =
            _deployProxy(impls.crossDexImpl, impls.routerImpl, impls.marketImpl, impls.pairImpl, feeController, params);

        vm.stopBroadcast();

        _logImplementations(impls);
        console.log("");
        _logProxyDeployment(deployment, feeController);
    }

    // ─────────────────────────────────────────────────────────────────────────────
    // Helper: Deploy & Upgrade All
    // ─────────────────────────────────────────────────────────────────────────────

    /// @notice Deploy fresh implementations and upgrade all existing proxies in one transaction.
    /// @param crossDex Existing CrossDex proxy address
    /// @param oldFeeControllers Previous fee controller addresses to remove from allowed list
    /// @return impls Newly deployed implementation addresses
    function deployAndUpgradeAll(address crossDex, address[] memory oldFeeControllers)
        external
        returns (Implementations memory impls)
    {
        vm.startBroadcast();

        impls = _deployImplementations();
        _upgradeAll(crossDex, impls, oldFeeControllers);

        vm.stopBroadcast();

        _logImplementations(impls);
    }

    /// @notice Upgrade all existing proxies with pre-deployed implementations.
    /// @param crossDex Existing CrossDex proxy address
    /// @param crossDexImpl New CrossDexImplV3 implementation
    /// @param routerImpl New CrossDexRouterV3 implementation
    /// @param marketImpl New MarketImplV3 implementation
    /// @param pairImpl New PairImplV3 implementation
    /// @param oldFeeControllers Previous fee controller addresses to remove from allowed list
    /// @param newFeeControllers New fee controller addresses to add to allowed list
    function upgradeAll(
        address crossDex,
        address crossDexImpl,
        address routerImpl,
        address marketImpl,
        address pairImpl,
        address[] memory oldFeeControllers,
        address[] memory newFeeControllers
    ) external {
        Implementations memory impls;
        impls.crossDexImpl = crossDexImpl;
        impls.routerImpl = routerImpl;
        impls.marketImpl = marketImpl;
        impls.pairImpl = pairImpl;

        vm.startBroadcast();
        _upgradeAll(crossDex, impls, oldFeeControllers);

        CrossDexImplV3 dex = CrossDexImplV3(crossDex);
        for (uint256 i = 0; i < newFeeControllers.length; ++i) {
            dex.setFeeControllerAllow(newFeeControllers[i], true);
        }

        vm.stopBroadcast();
    }

    /// @notice Upgrade a single Market and all its Pairs to the implementations currently set in CrossDex.
    ///         Must be called by the Market owner (may differ from CrossDex owner).
    /// @param crossDex CrossDex proxy address (reads marketImpl / pairImpl from it)
    /// @param marketProxy Market proxy address to upgrade
    function upgradeMarketFromDex(address crossDex, address marketProxy) external {
        CrossDexImplV3 dex = CrossDexImplV3(crossDex);
        address newMarketImpl = dex.marketImpl();
        address newPairImpl = dex.pairImpl();

        vm.startBroadcast();

        MarketImplV3 market = MarketImplV3(marketProxy);
        market.upgradeToAndCall(newMarketImpl, hex"");
        market.setPairImpl(newPairImpl);

        (, address[] memory pairs) = market.allPairs();
        for (uint256 j = 0; j < pairs.length; ++j) {
            PairImplV3(pairs[j]).upgradeToAndCall(newPairImpl, hex"");
        }

        vm.stopBroadcast();

        console.log("=== Market Upgrade Complete ===");
        console.log("Market:", marketProxy);
        console.log("Pairs upgraded:", pairs.length);
        console.log("Using marketImpl:", newMarketImpl);
        console.log("Using pairImpl:", newPairImpl);
    }

    /// @notice Migrate a Market's fee controller to a new address, preserving each Pair's existing config.
    ///         Reads each Pair's current fee config via getFeeControllerConfig(), then re-initializes
    ///         with the new fee controller address and the same config data.
    /// @param marketProxy Market proxy address
    /// @param newFeeController New fee controller implementation address (must be allowed in CrossDex)
    function migrateFeeController(address marketProxy, address newFeeController) external {
        MarketImplV3 market = MarketImplV3(marketProxy);
        (, address[] memory pairs) = market.allPairs();

        bytes[] memory configs = new bytes[](pairs.length);
        for (uint256 i = 0; i < pairs.length; ++i) {
            (, configs[i]) = PairImplV3(pairs[i]).getFeeControllerConfig();
        }

        vm.startBroadcast();

        market.setFeeController(new address[](0), newFeeController, hex"");

        for (uint256 i = 0; i < pairs.length; ++i) {
            PairImplV3(pairs[i]).setFeeController(newFeeController, configs[i]);
        }

        vm.stopBroadcast();

        console.log("=== FeeController Migration Complete ===");
        console.log("Market:", marketProxy);
        console.log("New FeeController:", newFeeController);
        console.log("Pairs migrated:", pairs.length);
    }

    function _upgradeAll(
        address crossDex,
        Implementations memory impls,
        address[] memory oldFeeControllers
    ) internal {
        CrossDexImplV3 dex = CrossDexImplV3(crossDex);
        address router = dex.ROUTER();

        // Upgrade CrossDex & Router proxies
        dex.upgradeToAndCall(impls.crossDexImpl, hex"");
        CrossDexRouterV3(router).upgradeToAndCall(impls.routerImpl, hex"");

        // Update impl references stored in CrossDex
        dex.setMarketImpl(impls.marketImpl);
        dex.setPairImpl(impls.pairImpl);

        // Remove old fee controllers
        for (uint256 i = 0; i < oldFeeControllers.length; ++i) {
            dex.setFeeControllerAllow(oldFeeControllers[i], false);
        }

        // Allow newly deployed fee controllers
        if (impls.feeControllerV2Compat != address(0)) {
            dex.setFeeControllerAllow(impls.feeControllerV2Compat, true);
        }
        if (impls.feeControllerV3Dist != address(0)) {
            dex.setFeeControllerAllow(impls.feeControllerV3Dist, true);
        }
        if (impls.feeControllerV3Split != address(0)) {
            dex.setFeeControllerAllow(impls.feeControllerV3Split, true);
        }

        console.log("=== CrossDex Upgrade Complete ===");
        console.log("CrossDex:", crossDex);
        console.log("Router:", router);
    }
}
