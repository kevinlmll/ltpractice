package src

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connect(root.Left)
	connect(root.Right)

	nl := root.Left
	nr := root.Right
	for nl != nil {
		nl.Next = nr
		nl = nl.Right
		nr = nr.Left
	}
	return root
}
