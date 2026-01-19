// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.20;

import {List} from "../../src/lib/List.sol";
import {Test} from "forge-std/Test.sol";

contract ListWrapper {
    using List for List.U256;

    List.U256 internal _list;

    function empty() external view returns (bool) {
        return _list.empty();
    }

    function length() external view returns (uint256) {
        return List.length(_list);
    }

    function contains(uint256 data) external view returns (bool) {
        return _list.contains(data);
    }

    function at(uint256 index) external view returns (uint256) {
        return _list.at(index);
    }

    function values() external view returns (uint256[] memory) {
        return _list.values();
    }

    function findAscPrev(uint256 data, uint256[2] memory adjacent, uint256 findMaxCount)
        external
        view
        returns (uint256)
    {
        return _list.findAscPrev(data, adjacent, findMaxCount);
    }

    function findDescPrev(uint256 data, uint256[2] memory adjacent, uint256 findMaxCount)
        external
        view
        returns (uint256)
    {
        return _list.findDescPrev(data, adjacent, findMaxCount);
    }

    function insert(uint256 data, uint256 prev) external returns (bool) {
        return _list.insert(data, prev);
    }

    function push(uint256 data) external returns (bool) {
        return _list.push(data);
    }

    function remove(uint256 data) external returns (bool) {
        return _list.remove(data);
    }

    function head() external view returns (uint256) {
        return _list.head;
    }

    function tail() external view returns (uint256) {
        return _list.tail;
    }
}

contract ListTest is Test {
    ListWrapper public list;

    function setUp() external {
        list = new ListWrapper();
    }

    // ==================== Basic Operations ====================

    function test_empty_onNewList() external view {
        assertTrue(list.empty());
        assertEq(list.length(), 0);
    }

    function test_empty_afterInsert() external {
        list.insert(100, 0);
        assertFalse(list.empty());
    }

    function test_length_tracksCorrectly() external {
        assertEq(list.length(), 0);

        list.insert(100, 0);
        assertEq(list.length(), 1);

        list.insert(200, 100);
        assertEq(list.length(), 2);

        list.insert(300, 200);
        assertEq(list.length(), 3);

        list.remove(200);
        assertEq(list.length(), 2);

        list.remove(100);
        assertEq(list.length(), 1);

        list.remove(300);
        assertEq(list.length(), 0);
    }

    function test_contains_returnsTrueForExisting() external {
        list.insert(100, 0);
        list.insert(200, 100);
        list.insert(300, 200);

        assertTrue(list.contains(100));
        assertTrue(list.contains(200));
        assertTrue(list.contains(300));
    }

    function test_contains_returnsFalseForNonExisting() external {
        list.insert(100, 0);
        list.insert(200, 100);

        assertFalse(list.contains(50));
        assertFalse(list.contains(150));
        assertFalse(list.contains(300));
    }

    function test_contains_returnsFalseForZero() external {
        list.insert(100, 0);
        assertFalse(list.contains(0));
    }

    function test_contains_returnsTrueForHead() external {
        list.insert(100, 0);
        assertTrue(list.contains(100));
        assertEq(list.head(), 100);
    }

    function test_at_returnsCorrectElement() external {
        list.insert(100, 0);
        list.insert(200, 100);
        list.insert(300, 200);

        assertEq(list.at(0), 100);
        assertEq(list.at(1), 200);
        assertEq(list.at(2), 300);
    }

    function test_at_revert_outOfBounds() external {
        list.insert(100, 0);
        list.insert(200, 100);

        vm.expectRevert(List.ListInvalidIndex.selector);
        list.at(2);

        vm.expectRevert(List.ListInvalidIndex.selector);
        list.at(100);
    }

    function test_at_revert_emptyList() external {
        vm.expectRevert(List.ListInvalidIndex.selector);
        list.at(0);
    }

    function test_values_returnsAll() external {
        list.insert(100, 0);
        list.insert(200, 100);
        list.insert(300, 200);

        uint256[] memory vals = list.values();
        assertEq(vals.length, 3);
        assertEq(vals[0], 100);
        assertEq(vals[1], 200);
        assertEq(vals[2], 300);
    }

    function test_values_emptyList() external view {
        uint256[] memory vals = list.values();
        assertEq(vals.length, 0);
    }

    // ==================== Insert Operations ====================

    function test_insert_atHead() external {
        list.insert(200, 0);
        list.insert(100, 0);

        assertEq(list.head(), 100);
        assertEq(list.at(0), 100);
        assertEq(list.at(1), 200);
    }

    function test_insert_atTail() external {
        list.insert(100, 0);
        list.insert(200, 100);
        list.insert(300, 200);

        assertEq(list.tail(), 300);
        assertEq(list.at(2), 300);
    }

    function test_insert_inMiddle() external {
        list.insert(100, 0);
        list.insert(300, 100);
        list.insert(200, 100); // Insert 200 after 100

        uint256[] memory vals = list.values();
        assertEq(vals[0], 100);
        assertEq(vals[1], 200);
        assertEq(vals[2], 300);
    }

    function test_insert_duplicate_returnsFalse() external {
        list.insert(100, 0);
        bool inserted = list.insert(100, 0);
        assertFalse(inserted);
        assertEq(list.length(), 1);
    }

    function test_insert_revert_zeroData() external {
        vm.expectRevert(List.ListZeroData.selector);
        list.insert(0, 0);
    }

    function test_insert_revert_invalidPrevNode() external {
        list.insert(100, 0);

        vm.expectRevert(List.ListInvalidPrevNode.selector);
        list.insert(200, 999); // 999 is not in the list
    }

    function test_insert_firstElement() external {
        bool inserted = list.insert(100, 0);
        assertTrue(inserted);
        assertEq(list.head(), 100);
        assertEq(list.tail(), 100);
        assertEq(list.length(), 1);
    }

    function test_insert_updatesHeadAndTail() external {
        list.insert(100, 0);
        assertEq(list.head(), 100);
        assertEq(list.tail(), 100);

        list.insert(200, 100);
        assertEq(list.head(), 100);
        assertEq(list.tail(), 200);

        list.insert(50, 0);
        assertEq(list.head(), 50);
        assertEq(list.tail(), 200);
    }

    // ==================== Push Operations ====================

    function test_push_appendsToTail() external {
        list.push(100);
        list.push(200);
        list.push(300);

        assertEq(list.head(), 100);
        assertEq(list.tail(), 300);

        uint256[] memory vals = list.values();
        assertEq(vals[0], 100);
        assertEq(vals[1], 200);
        assertEq(vals[2], 300);
    }

    function test_push_duplicate_returnsFalse() external {
        list.push(100);
        bool pushed = list.push(100);
        assertFalse(pushed);
        assertEq(list.length(), 1);
    }

    function test_push_revert_zeroData() external {
        vm.expectRevert(List.ListZeroData.selector);
        list.push(0);
    }

    // ==================== Remove Operations ====================

    function test_remove_head() external {
        list.insert(100, 0);
        list.insert(200, 100);
        list.insert(300, 200);

        list.remove(100);

        assertEq(list.head(), 200);
        assertEq(list.length(), 2);
        assertFalse(list.contains(100));
    }

    function test_remove_tail() external {
        list.insert(100, 0);
        list.insert(200, 100);
        list.insert(300, 200);

        list.remove(300);

        assertEq(list.tail(), 200);
        assertEq(list.length(), 2);
        assertFalse(list.contains(300));
    }

    function test_remove_middle() external {
        list.insert(100, 0);
        list.insert(200, 100);
        list.insert(300, 200);

        list.remove(200);

        assertEq(list.length(), 2);
        assertFalse(list.contains(200));

        uint256[] memory vals = list.values();
        assertEq(vals[0], 100);
        assertEq(vals[1], 300);
    }

    function test_remove_nonExistent_returnsFalse() external {
        list.insert(100, 0);
        bool removed = list.remove(999);
        assertFalse(removed);
        assertEq(list.length(), 1);
    }

    function test_remove_revert_zeroData() external {
        vm.expectRevert(List.ListZeroData.selector);
        list.remove(0);
    }

    function test_remove_lastElement() external {
        list.insert(100, 0);
        list.remove(100);

        assertTrue(list.empty());
        assertEq(list.head(), 0);
        assertEq(list.tail(), 0);
    }

    function test_remove_allowsReinsert() external {
        list.insert(100, 0);
        list.remove(100);

        bool inserted = list.insert(100, 0);
        assertTrue(inserted);
        assertTrue(list.contains(100));
    }

    // ==================== findAscPrev Tests ====================

    function test_findAscPrev_emptyList() external view {
        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findAscPrev(100, adjacent, 10);
        assertEq(prev, 0);
    }

    function test_findAscPrev_lessThanHead() external {
        list.push(100);
        list.push(200);
        list.push(300);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findAscPrev(50, adjacent, 10);
        assertEq(prev, 0);
    }

    function test_findAscPrev_greaterThanTail() external {
        list.push(100);
        list.push(200);
        list.push(300);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findAscPrev(400, adjacent, 10);
        assertEq(prev, 300);
    }

    function test_findAscPrev_inMiddle() external {
        list.push(100);
        list.push(200);
        list.push(400);
        list.push(500);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findAscPrev(300, adjacent, 10);
        assertEq(prev, 200);
    }

    function test_findAscPrev_existingElement() external {
        list.push(100);
        list.push(200);
        list.push(300);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findAscPrev(200, adjacent, 10);
        assertEq(prev, 100);
    }

    function test_findAscPrev_withAdjacentHint() external {
        list.push(100);
        list.push(200);
        list.push(300);
        list.push(400);
        list.push(500);

        uint256[2] memory adjacent = [uint256(300), uint256(400)];
        uint256 prev = list.findAscPrev(350, adjacent, 10);
        assertEq(prev, 300);
    }

    function test_findAscPrev_revert_zeroData() external {
        list.push(100);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.expectRevert(List.ListZeroData.selector);
        list.findAscPrev(0, adjacent, 10);
    }

    function test_findAscPrev_revert_invalidFindMaxCount() external {
        list.push(100);
        list.push(200);
        list.push(300);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.expectRevert(List.ListInvalidFindMaxCount.selector);
        list.findAscPrev(150, adjacent, 0);
    }

    function test_findAscPrev_revert_failToFindPrev() external {
        list.push(100);
        list.push(200);
        list.push(300);
        list.push(400);
        list.push(500);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.expectRevert(List.ListFailToFindPrev.selector);
        list.findAscPrev(350, adjacent, 1); // maxCount too low
    }

    function test_findAscPrev_adjacentBackwardSearch() external {
        list.push(100);
        list.push(200);
        list.push(300);
        list.push(400);
        list.push(500);

        // Give a hint that's higher than the target, triggering backward search
        uint256[2] memory adjacent = [uint256(400), uint256(500)];
        uint256 prev = list.findAscPrev(250, adjacent, 10);
        assertEq(prev, 200);
    }

    // ==================== findDescPrev Tests ====================

    function test_findDescPrev_emptyList() external view {
        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findDescPrev(100, adjacent, 10);
        assertEq(prev, 0);
    }

    function test_findDescPrev_greaterThanHead() external {
        // Descending list: 500, 400, 300
        list.push(500);
        list.push(400);
        list.push(300);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findDescPrev(600, adjacent, 10);
        assertEq(prev, 0);
    }

    function test_findDescPrev_lessThanTail() external {
        // Descending list: 500, 400, 300
        list.push(500);
        list.push(400);
        list.push(300);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findDescPrev(200, adjacent, 10);
        assertEq(prev, 300);
    }

    function test_findDescPrev_inMiddle() external {
        // Descending list: 500, 400, 200, 100
        list.push(500);
        list.push(400);
        list.push(200);
        list.push(100);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findDescPrev(300, adjacent, 10);
        assertEq(prev, 400);
    }

    function test_findDescPrev_existingElement() external {
        list.push(500);
        list.push(400);
        list.push(300);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        uint256 prev = list.findDescPrev(400, adjacent, 10);
        assertEq(prev, 500);
    }

    function test_findDescPrev_withAdjacentHint() external {
        list.push(500);
        list.push(400);
        list.push(300);
        list.push(200);
        list.push(100);

        uint256[2] memory adjacent = [uint256(300), uint256(200)];
        uint256 prev = list.findDescPrev(250, adjacent, 10);
        assertEq(prev, 300);
    }

    function test_findDescPrev_revert_zeroData() external {
        list.push(500);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.expectRevert(List.ListZeroData.selector);
        list.findDescPrev(0, adjacent, 10);
    }

    function test_findDescPrev_revert_invalidFindMaxCount() external {
        list.push(500);
        list.push(400);
        list.push(300);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.expectRevert(List.ListInvalidFindMaxCount.selector);
        list.findDescPrev(350, adjacent, 0);
    }

    function test_findDescPrev_revert_failToFindPrev() external {
        list.push(500);
        list.push(400);
        list.push(300);
        list.push(200);
        list.push(100);

        uint256[2] memory adjacent = [uint256(0), uint256(0)];
        vm.expectRevert(List.ListFailToFindPrev.selector);
        list.findDescPrev(350, adjacent, 1); // maxCount too low
    }

    function test_findDescPrev_adjacentBackwardSearch() external {
        list.push(500);
        list.push(400);
        list.push(300);
        list.push(200);
        list.push(100);

        // Give a hint that's lower than the target, triggering backward search
        uint256[2] memory adjacent = [uint256(200), uint256(100)];
        uint256 prev = list.findDescPrev(350, adjacent, 10);
        assertEq(prev, 400);
    }

    // ==================== Edge Cases ====================

    function test_singleElement_operations() external {
        list.insert(100, 0);

        assertTrue(list.contains(100));
        assertEq(list.head(), 100);
        assertEq(list.tail(), 100);
        assertEq(list.at(0), 100);
        assertEq(list.length(), 1);

        uint256[] memory vals = list.values();
        assertEq(vals.length, 1);
        assertEq(vals[0], 100);
    }

    function test_largeList() external {
        uint256 count = 100;
        for (uint256 i = 1; i <= count; i++) {
            list.push(i * 10);
        }

        assertEq(list.length(), count);
        assertEq(list.head(), 10);
        assertEq(list.tail(), count * 10);

        // Remove from middle
        list.remove(500);
        assertEq(list.length(), count - 1);
        assertFalse(list.contains(500));

        // Verify list integrity
        uint256[] memory vals = list.values();
        assertEq(vals.length, count - 1);
    }

    function test_insertRemoveSequence() external {
        // Insert and remove in various orders
        list.insert(100, 0);
        list.insert(200, 100);
        list.insert(300, 200);

        list.remove(200);
        list.insert(250, 100);

        uint256[] memory vals = list.values();
        assertEq(vals.length, 3);
        assertEq(vals[0], 100);
        assertEq(vals[1], 250);
        assertEq(vals[2], 300);
    }

    function test_headOnlyLinked() external {
        // When head has no next (single element), it should still be "linked"
        list.insert(100, 0);
        assertTrue(list.contains(100));

        list.insert(200, 100);
        assertTrue(list.contains(100));
        assertTrue(list.contains(200));
    }
}
