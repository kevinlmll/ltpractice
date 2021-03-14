package src

func traverse(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	ls := traverse(root.Left, max)
	rs := traverse(root.Right, max)

	smax := ls
	if rs > smax {
		smax = rs
	}
	sval := root.Val + smax
	if root.Val > sval {
		sval = root.Val
	}

	aval := root.Val
	if aval+ls > aval {
		aval = aval + ls
	}
	if aval+rs > aval {
		aval = aval + rs
	}

	if sval > *max {
		*max = sval
	}
	if aval > *max {
		*max = aval
	}
	return sval
}

func maxPathSum(root *TreeNode) int {
	max := -1001
	traverse(root, &max)
	return max
}
