package tree

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.Traverse()
	node.Right.Traverse()
}
