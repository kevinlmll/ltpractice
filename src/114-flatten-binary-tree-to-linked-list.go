package src

func flatten(root *TreeNode) {
	flattenSub(root)
}

func flattenSub(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	nl := flattenSub(root.Left)
	nr := flattenSub(root.Right)

	if nl == nil && nr == nil {
		return root
	}
	if nl == nil {
		return nr
	}
	if nr == nil {
		root.Right = root.Left
		root.Left = nil
		return nl
	}
	nl.Left = nil
	nl.Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return nr
}
