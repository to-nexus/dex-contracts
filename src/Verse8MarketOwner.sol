// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {AccessControlDefaultAdminRules} from
    "@openzeppelin-contracts-5.2.0/access/extensions/AccessControlDefaultAdminRules.sol";

import {IMarketV2} from "./interfaces/IMarket.sol";

/**
 * @title Verse8MarketOwner
 * @notice A contract that manages MarketImpl ownership with separated permissions
 * @dev Delegates createPair authority to a separate pairCreator,
 *      while the owner retains all other permissions.
 *      This contract is non-upgradeable.
 */
contract Verse8MarketOwner is AccessControlDefaultAdminRules {
    error Verse8MarketOwner__CallFailed();

    bytes32 private constant PAIR_CREATOR_ROLE = keccak256("PAIR_CREATOR_ROLE");

    constructor(address _owner, address[] memory creators) AccessControlDefaultAdminRules(0, _owner) {
        _grantRole(PAIR_CREATOR_ROLE, _owner); // Owner can also create pairs
        for (uint256 i = 0; i < creators.length; i++) {
            _grantRole(PAIR_CREATOR_ROLE, creators[i]);
        }
    }

    function createPair(address market, address base, uint256 tickSize, uint256 lotSize, bytes memory feeData)
        external
        onlyRole(PAIR_CREATOR_ROLE)
        returns (address)
    {
        return IMarketV2(market).createPair(base, tickSize, lotSize, feeData);
    }

    struct CreatePairArgs {
        address market;
        address base;
        uint256 tickSize;
        uint256 lotSize;
        bytes feeData;
    }

    function createPairs(CreatePairArgs[] calldata args)
        external
        onlyRole(PAIR_CREATOR_ROLE)
        returns (address[] memory)
    {
        address[] memory pairs = new address[](args.length);
        for (uint256 i = 0; i < args.length; i++) {
            CreatePairArgs memory arg = args[i];
            pairs[i] = IMarketV2(arg.market).createPair(arg.base, arg.tickSize, arg.lotSize, arg.feeData);
        }
        return pairs;
    }

    function execute(address to, uint256 value, bytes calldata data)
        external
        onlyRole(DEFAULT_ADMIN_ROLE) // onlyOwner
        returns (bytes memory)
    {
        (bool success, bytes memory result) = to.call{value: value}(data);
        if (!success) revert Verse8MarketOwner__CallFailed();
        return result;
    }

    struct ExecuteBatchArgs {
        address to;
        uint256 value;
        bytes data;
    }

    function executeBatch(ExecuteBatchArgs[] calldata calls)
        external
        onlyRole(DEFAULT_ADMIN_ROLE) // onlyOwner
        returns (bytes[] memory results)
    {
        results = new bytes[](calls.length);
        for (uint256 i = 0; i < calls.length; i++) {
            (bool success, bytes memory result) = calls[i].to.call{value: calls[i].value}(calls[i].data);
            if (!success) revert Verse8MarketOwner__CallFailed();
            results[i] = result;
        }
    }
}
