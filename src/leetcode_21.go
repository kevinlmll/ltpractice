package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var h *ListNode
	h1 := l1
	h2 := l2

	if h1 == nil && h2 == nil {
		return head
	}

}
