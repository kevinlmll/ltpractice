package src

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	retHead := head
	moveFirst := head
	moveSecond := head
	moveThird := head

	//n是保证有效的
	for n > 0 {
		moveFirst = moveFirst.Next
		n--
	}

	for moveFirst != nil {
		moveFirst = moveFirst.Next
		moveSecond = moveSecond.Next
		if n > 0 {
			moveThird = moveThird.Next
		}
		n++
	}

	//这时moveSecond是需要删除的结点
	if moveSecond == head {
		retHead = head.Next
	} else {
		moveThird.Next = moveSecond.Next
	}
	return retHead
}
