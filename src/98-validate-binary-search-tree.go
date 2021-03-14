package src

func isValidBST(root *TreeNode) bool {
	r, _ := isValidBSTSub(root)
	return r
}

func isValidBSTSub(root *TreeNode) (bool, []int) {
	if root == nil {
		return true, nil
	}

	lb, lo := isValidBSTSub(root.Left)
	if lb != true {
		return false, nil
	}

	rb, ro := isValidBSTSub(root.Right)
	if rb != true {
		return false, nil
	}

	orders := make([]int, 0, 1+len(lo)+len(ro))
	if lo != nil {
		end := len(lo) - 1
		if lo[end] >= root.Val {
			return false, nil
		}
		orders = append(orders, lo...)
	}
	orders = append(orders, root.Val)
	if ro != nil {
		start := 0
		if ro[start] <= root.Val {
			return false, nil
		}
		orders = append(orders, ro...)
	}
	return true, orders
}
