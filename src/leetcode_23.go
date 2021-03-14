package src

import (
	"sort"
)

// 普通的merge sort方式，算法复杂度是O(mn)
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var head *ListNode = nil
	var tail *ListNode = nil
	listIdx := make([]*ListNode, len(lists))
	for i := range lists {
		listIdx[i] = lists[i]
	}

	for {
		nilNum := 0
		var roundIdx int = 0
		var roundNode *ListNode = nil
		for i, ln := range listIdx {
			if ln == nil {
				nilNum += 1
				continue
			}
			if roundNode == nil {
				roundNode = ln
				roundIdx = i
			} else {
				if ln.Val < roundNode.Val {
					roundNode = ln
					roundIdx = i
				}
			}
		}
		if nilNum == len(listIdx) {
			break
		}
		listIdx[roundIdx] = roundNode.Next

		if head == nil {
			head = roundNode
			tail = roundNode
		} else {
			tail.Next = roundNode
			tail = roundNode
		}
	}
	return head
}

// 已知链表是排序的情况下，可以做优化
// 最优情况是O(m)，最差情况下是O(nm)，平均是介于两者中间
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var head *ListNode = nil
	var tail *ListNode = nil
	minList := make([]*ListNode, 0)
	for i := range lists {
		if lists[i] != nil {
			minList = append(minList, lists[i])
		}
	}
	if minList == nil || len(minList) == 0 {
		return nil
	}

	sort.Slice(minList, func(i, j int) bool {
		return minList[i].Val < minList[j].Val
	})

	for i := 0; i < len(minList); {
		//fmt.Printf("%d\n", i)
		nextIdx := i + 1
		if nextIdx >= len(minList) {
			//没有下一级队列，把当前队列遍历完即可
			if tail != nil {
				tail.Next = minList[i]
			} else {
				tail = minList[i]
				head = tail
			}
			break
		}
		if minList[i] == nil {
			//当前队列已经为空了
			i++
			continue
		}

		//遍历当前队列，直到Val>minList[nextIdx].Val
		node := minList[i]
		for node != nil && node.Val <= minList[nextIdx].Val {
			//fmt.Printf("current idx:%d, node:%+v\n", i, *node)
			if head == nil {
				tail = node
				head = node
			} else {
				tail.Next = node
				tail = node
			}
			node = node.Next
		}
		minList[i] = node
		if node != nil {
			j := i
			for j < len(minList)-1 {
				next := j + 1
				if minList[j].Val > minList[next].Val {
					tmp := minList[j]
					minList[j] = minList[next]
					minList[next] = tmp
				}
				j++
			}
		}
	}
	return head
}
