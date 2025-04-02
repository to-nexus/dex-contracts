// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

library List {
    struct Node {
        uint256 prev;
        uint256 next;
    }

    error ListZeroData();
    error ListInvalidIndex();
    error ListInvalidPrevNode();
    error ListInvalidFindMaxCount();
    error ListFailToFindPrev();

    function _linked(Node memory node) private pure returns (bool) {
        return node.prev != 0 || node.next != 0;
    }

    struct U256 {
        uint256 length;
        uint256 head;
        uint256 tail;
        mapping(uint256 => Node) nodes;
    }

    function empty(U256 storage _list) internal view returns (bool) {
        return _list.length == 0;
    }

    function length(U256 storage _list) internal view returns (uint256) {
        return _list.length;
    }

    function contains(U256 storage _list, uint256 _data) internal view returns (bool) {
        if (_data == 0) return false;
        return _data == _list.head || _linked(_list.nodes[_data]);
    }

    function at(U256 storage _list, uint256 _index) internal view returns (uint256) {
        if (_index >= _list.length) revert ListInvalidIndex();

        uint256 data = _list.head;
        for (uint256 i = _index; i != 0;) {
            unchecked {
                --i;
            }
            data = _list.nodes[data].next;
        }
        return data;
    }

    function values(U256 storage _list) internal view returns (uint256[] memory) {
        uint256[] memory result = new uint256[](_list.length);
        uint256 _length = _list.length;
        uint256 data = _list.head;
        for (uint256 i = 0; i < _length;) {
            result[i] = data;
            data = _list.nodes[data].next;
            unchecked {
                ++i;
            }
        }
        return result;
    }

    function findASCPrev(U256 storage _list, uint256 _data, uint256[2] memory _adjacent, uint256 _findMaxCount)
        internal
        view
        returns (uint256)
    {
        if (_data == 0) revert ListZeroData();
        if (contains(_list, _data)) return _list.nodes[_data].prev;

        if (_data < _list.head || empty(_list)) {
            return 0;
        } else if (_data > _list.tail) {
            return _list.tail;
        } else {
            if (_findMaxCount == 0) revert ListInvalidFindMaxCount();
            uint256 current = _ensureAdjacentForPush(_list, _adjacent);
            unchecked {
                if (current < _data) {
                    while (current != 0 && --_findMaxCount != 0) {
                        current = _list.nodes[current].next;
                        if (current > _data) return _list.nodes[current].prev;
                    }
                } else {
                    while (current != 0 && --_findMaxCount != 0) {
                        current = _list.nodes[current].prev;
                        if (current < _data) return current;
                    }
                }
            }
            revert ListFailToFindPrev();
        }
    }

    function findDESCPrev(U256 storage _list, uint256 _data, uint256[2] memory _adjacent, uint256 _findMaxCount)
        internal
        view
        returns (uint256)
    {
        if (_data == 0) revert ListZeroData();
        if (contains(_list, _data)) return _list.nodes[_data].prev;

        if (_data > _list.head || empty(_list)) {
            return 0;
        } else if (_data < _list.tail) {
            return _list.tail;
        } else {
            if (_findMaxCount == 0) revert ListInvalidFindMaxCount();
            uint256 current = _ensureAdjacentForPush(_list, _adjacent);
            unchecked {
                if (current > _data) {
                    while (current != 0 && --_findMaxCount != 0) {
                        current = _list.nodes[current].next;
                        if (_data > current) return _list.nodes[current].prev;
                    }
                } else {
                    // current < _data
                    while (current != 0 && --_findMaxCount != 0) {
                        current = _list.nodes[current].prev;
                        if (_data < current) return current;
                    }
                }
            }
            revert ListFailToFindPrev();
        }
    }

    function insert(U256 storage _list, uint256 _data, uint256 _prev) internal returns (bool) {
        if (_data == 0) revert ListZeroData();
        if (contains(_list, _data)) return false;

        Node storage prevNode = _list.nodes[_prev];
        Node memory newNode;

        if (_prev == 0) {
            // change head
            uint256 _head = _list.head;
            _list.head = _data;

            if (_head != 0) _list.nodes[_head].prev = _data;
            newNode = Node({prev: 0, next: _head});
        } else {
            if (!(_linked(prevNode) || _prev == _list.head)) revert ListInvalidPrevNode();
            uint256 prevNext = prevNode.next;
            newNode = Node({prev: _prev, next: prevNext});
            if (prevNext != 0) _list.nodes[prevNext].prev = _data;
            prevNode.next = _data;
        }

        if (newNode.next == 0) _list.tail = _data;

        _list.nodes[_data] = newNode;
        ++_list.length;

        return true;
    }

    function push(U256 storage _list, uint256 _data) internal returns (bool) {
        return insert(_list, _data, _list.tail);
    }

    function remove(U256 storage _list, uint256 _data) internal returns (bool) {
        if (_data == 0) revert ListZeroData();
        if (!contains(_list, _data)) return false;

        Node storage node = _list.nodes[_data];
        Node storage prevNode = _list.nodes[node.prev];
        Node storage nextNode = _list.nodes[node.next];

        prevNode.next = node.next;
        nextNode.prev = node.prev;

        if (_list.head == _data) _list.head = node.next;
        if (_list.tail == _data) _list.tail = node.prev;

        delete _list.nodes[_data];
        --_list.length;
        return true;
    }

    function _ensureAdjacentForPush(U256 storage _list, uint256[2] memory _adjacent) private view returns (uint256) {
        if (contains(_list, _adjacent[0])) return (_adjacent[0]);
        else if (contains(_list, _adjacent[1])) return (_adjacent[1]);
        else return _list.head;
    }
}
