// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "./List.sol";

library ASCList {
    using List for List.U256;

    error ASCListPushFailed();
    error ASCListZeroData();

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
        if (_data == 0) revert ASCListZeroData();
        return _list._inner.contains(_data);
    }

    function at(U256 storage _list, uint256 _index) internal view returns (uint256) {
        return _list._inner.at(_index);
    }

    function first(U256 storage _list) internal view returns (uint256) {
        return _list._inner.head;
    }

    function last(U256 storage _list) internal view returns (uint256) {
        return _list._inner.tail;
    }

    function values(U256 storage _list) internal view returns (uint256[] memory) {
        return _list._inner.values();
    }

    function push(U256 storage _list, uint256 _data, uint256 _search) internal returns (bool) {
        if (_data == 0) revert ASCListZeroData();

        List.U256 storage list = _list._inner;
        if (list.contains(_data)) return false;

        if (_data < list.head || list.empty()) {
            return list.insert(_data, 0);
        } else if (_data > list.tail) {
            return list.insert(_data, list.tail);
        } else {
            uint256 current = list.contains(_search) ? _search : list.head;
            if (current < _data) {
                while (current != 0) {
                    if (_data < current) return list.insert(_data, list.nodes[current].prev);
                    current = list.nodes[current].next;
                }
            } else {
                while (current != 0) {
                    if (_data > current) return list.insert(_data, current);
                    current = list.nodes[current].prev;
                }
            }
            revert ASCListPushFailed();
        }
    }

    function push(U256 storage _list, uint256 _data) internal returns (bool) {
        if (_data == 0) revert ASCListZeroData();
        return push(_list, _data, _list._inner.head);
    }

    function remove(U256 storage _list, uint256 _data) internal returns (bool) {
        if (_data == 0) revert ASCListZeroData();
        return _list._inner.remove(_data);
    }
}
