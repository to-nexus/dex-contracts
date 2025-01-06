// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "./DoubleLinkedList.sol";

library DESCList {
    using DoubleLinkedList for DoubleLinkedList.U256;

    error DESCListPushFailed();
    error DESCListZeroData();

    struct U256 {
        DoubleLinkedList.U256 _inner;
    }

    function empty(U256 storage _list) internal view returns (bool) {
        return _list._inner.empty();
    }

    function length(U256 storage _list) internal view returns (uint256) {
        return DoubleLinkedList.length(_list._inner);
    }

    function contains(U256 storage _list, uint256 _data) internal view returns (bool) {
        if (_data == 0) revert DESCListZeroData();
        return _list._inner.contains(_data);
    }

    function at(U256 storage _list, uint256 _index) internal view returns (uint256) {
        return _list._inner.at(_index);
    }

    function next(U256 storage _list, uint256 value) internal view returns (uint256) {
        return _list._inner.nodes[value].next;
    }

    function values(U256 storage _list) internal view returns (uint256[] memory) {
        return _list._inner.values();
    }

    function push(U256 storage _list, uint256 _data, uint256 _search) internal returns (bool) {
        if (_data == 0) revert DESCListZeroData();

        DoubleLinkedList.U256 storage list = _list._inner;
        if (list.contains(_data)) {
            return false;
        }

        if (_data > list.head) {
            return list.insert(_data, 0);
        } else if (_data < list.tail) {
            return list.insert(_data, list.tail);
        } else {
            uint256 current = list.contains(_search) ? _search : list.head;
            if (current > _data) {
                while (current != 0) {
                    if (_data > current) {
                        return list.insert(_data, list.nodes[current].prev);
                    }
                    current = list.nodes[current].next;
                }
            } else {
                while (current != 0) {
                    if (_data < current) {
                        return list.insert(_data, current);
                    }
                    current = list.nodes[current].prev;
                }
            }
            revert DESCListPushFailed();
        }
    }

    function push(U256 storage _list, uint256 _data) internal returns (bool) {
        if (_data == 0) revert DESCListZeroData();
        return push(_list, _data, _list._inner.head);
    }

    function remove(U256 storage _list, uint256 _data) internal returns (bool) {
        if (_data == 0) revert DESCListZeroData();
        return _list._inner.remove(_data);
    }
}
