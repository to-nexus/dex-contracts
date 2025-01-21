// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/ERC20.sol";
import {Address} from "@openzeppelin-contracts-5.2.0/utils/Address.sol";

import {IRouter} from "./interfaces/IRouter.sol";

contract WCROSS is ERC20 {
    using Address for address payable;

    address payable public immutable ROUTER;

    constructor() ERC20("CrossDEX Wrapped Cross", "cdCROSS") {
        ROUTER = payable(_msgSender());
    }

    receive() external payable {
        _mint(_msgSender(), msg.value);
    }

    function _update(address from, address to, uint256 value) internal override {
        super._update(from, to, value);
        // burn 이 아닌경우 Router 또는 Pair가 아닌 경우에만 burn 후 transfer
        if (to != address(0)) {
            if (!(to == ROUTER || IRouter(ROUTER).isPair(to))) {
                _burn(to, value);
                payable(to).sendValue(value);
            }
        }
    }
}
