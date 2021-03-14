package src

import "fmt"

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("list node:%d, next:%v\n", head.Val, head.Next)
		head = head.Next
	}
	fmt.Printf("=========\n")
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	//把链表倒序，耗时O(n)
	var newHead *ListNode = nil
	var totalSz int = 0
	for head != nil {
		nextNode := head.Next
		if newHead == nil {
			head.Next = nil
		} else {
			head.Next = newHead
		}
		newHead = head
		head = nextNode
		totalSz++
	}
	head = nil
	leftNodeSz := totalSz % k
	if leftNodeSz > 0 {
		//有多余的节点,重新倒序处理
		for leftNodeSz > 0 {
			nextNode := newHead.Next
			if head == nil {
				newHead.Next = nil
			} else {
				newHead.Next = head
			}
			head = newHead
			newHead = nextNode
			leftNodeSz--
		}
	}
	//按每k个节点进行顺序,耗时O(n)
	for newHead != nil {
		tmpHead := newHead
		tmpTail := newHead
		ks := k - 1
		for ks > 0 {
			tmpTail = tmpTail.Next
			ks--
		}
		newHead = tmpTail.Next
		tmpTail.Next = head
		head = tmpHead
	}
	return head
}
