package main

import (
	tree "Learn_go/01_Go_basic_and_PrograTming_thinking/C_Object_oriented/01_StructsAndMethods"
	"fmt"
)

func main() {
	root := tree.TreeNode{}
	root1 := tree.treeNode{value: 1, left: &treeNode{}} // 可以指定具体参数
	// root2 := treeNode{2, nil} //参数少不行
	root3 := treeNode{3, &treeNode{}, &treeNode{}}
	root.left = new(treeNode) // new出来的直接返回地址
	root.right = new(treeNode)
	fmt.Printf("%v\n%v\n%v\nroot.left = %v, root.right = %v\n",
		root, root1, root3, root.left, root.right)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.left.right = createNode(999) //返回一个实例的地址
	fmt.Println(*root.left.right)

	root.print()

	// 如果接收者是一个 指针， 那么就会把root的地址传给方法，root本身就是一个地址
	// 如果接收者是一个 值，那么就会将root解析出来，然后拷贝一份给 方法
	root.setValue(888)
	root.print()

	pRoot := &root
	fmt.Print("pRoot := &root -> pRoot.value: ")
	pRoot.print()
	pRoot.setValue(299)
	fmt.Print("pRoot.setvalue(299) -> pRoot.value: ")
	pRoot.print()
	fmt.Print("pRoot.setvalue(299) -> root.value: ")
	root.print()

	var nilRoot *treeNode
	nilRoot.setValue(100)

	root.traverse()
}
