package main

// ref: https://leetcode.com/problems/middle-of-the-linked-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// assume fast is twice faster than slow
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next // twice faster
	}

	// when twice faster and fast finish faster, mean slow will be
	// in the middle of the process

	return slow
}

// First solution
// func middleNode(head *ListNode) *ListNode {
// 	var n, middle int
// 	curr := head

// 	// calculate total linkedList
// 	for curr.Next != nil {
// 		n++
// 		curr = curr.Next
// 	}

// 	middle = n / 2
// 	curr = head
// 	for middle > 0 {
// 		curr = curr.Next
// 		middle--
// 	}

// 	return curr
// }
