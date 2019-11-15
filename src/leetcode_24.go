package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var h *ListNode = nil
	rh := head
	m := head //1
	var t *ListNode = nil
	if head != nil {
		t = head.Next //2
	}

	for m != nil && t != nil {
		m.Next = t.Next
		t.Next = m
		if h != nil {
			h.Next = t
		} else {
			rh = t
		}

		h = m
		m = h.Next
		if m != nil {
			t = m.Next
		} else {
			t = nil
		}
	}
	return rh
}

func main() {

}
