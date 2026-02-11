// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {ERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/ERC20.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";

import {IRouter} from "./interfaces/IRouter.sol";

contract WETH is ERC20 {
    using Address for address payable;

    address payable public immutable ROUTER;

    constructor() ERC20("CrossDEX Wrapped CROSS", "CROSS") {
        ROUTER = payable(_msgSender());
    }

    receive() external payable {
        _mint(_msgSender(), msg.value);
    }

    function mintTo(address to) external payable {
        _mint(to, msg.value);
    }

    function _update(address from, address to, uint256 value) internal override {
        super._update(from, to, value);
        // [If it is not a burn case] If to is not the Pair, perform burn and then transfer.
        if (to != address(0)) {
            if (!IRouter(ROUTER).isPair(to)) {
                _burn(to, value);
                payable(to).sendValue(value);
            }
        }
    }
}
