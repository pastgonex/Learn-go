package main

import (
	"Learn_go/tree"
)

func main() {
	root := tree.Node{Value: 3}
	// root2 := treeNode{2, nil} //参数少不行
	//root3 := {3, &treeNode{}, &treeNode{}}
	root.Right = new(tree.Node)
	root.Traverse()
}
