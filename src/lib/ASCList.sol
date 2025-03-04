// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "./List.sol";

library ASCList {
    using List for List.U256;

    error ASCListPushFailed();

    struct U256 {
        List.U256 _inner;
    }

    function empty(U256 storage _list) internal view returns (bool) {
        return _list._inner.empty();
    }

    function length(U256 storage _list) internal view returns (uint256) {
        return List.length(_list._inner);
    }

    function contains(U256 storage _list, uint256 _data) internal view returns (bool) {
        return _list._inner.contains(_data);
    }

    function at(U256 storage _list, uint256 _index) internal view returns (uint256) {
        return _list._inner.at(_index);
    }

    function values(U256 storage _list) internal view returns (uint256[] memory) {
        return _list._inner.values();
    }

    function push(U256 storage _list, uint256 _data, uint256[2] memory _searchs) internal returns (bool) {
        List.U256 storage list = _list._inner;
        if (list.contains(_data)) return false;

        if (_data < list.head || list.empty()) {
            return list.insert(_data, 0);
        } else if (_data > list.tail) {
            return list.insert(_data, list.tail);
        } else {
            uint256 current = _checkContainsForPush(_list, _searchs);
            if (current < _data) {
                while (current != 0) {
                    current = list.nodes[current].next;
                    if (_data < current) return list.insert(_data, list.nodes[current].prev);
                }
            } else {
                while (current != 0) {
                    current = list.nodes[current].prev;
                    if (_data > current) return list.insert(_data, current);
                }
            }
            revert ASCListPushFailed();
        }
    }

    function push(U256 storage _list, uint256 _data) internal returns (bool) {
        uint256[2] memory _searchs;
        _searchs[0] = _list._inner.tail;
        return push(_list, _data, _searchs);
    }

    function remove(U256 storage _list, uint256 _data) internal returns (bool) {
        return _list._inner.remove(_data);
    }

    function _checkContainsForPush(U256 storage _list, uint256[2] memory _searchs) private view returns (uint256) {
        if (contains(_list, _searchs[0])) return (_searchs[0]);
        else if (contains(_list, _searchs[1])) return (_searchs[1]);
        else return _list._inner.head;
    }
}
