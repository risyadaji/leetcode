package main

// References: https://leetcode.com/problems/merge-two-sorted-lists/description

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{} // dummy head, will return start from next of the head
	head := merged

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	for list1 != nil {
		head.Next = list1
		head = head.Next
		list1 = list1.Next
	}
	for list2 != nil {
		head.Next = list2
		head = head.Next
		list2 = list2.Next
	}

	return merged.Next
}
