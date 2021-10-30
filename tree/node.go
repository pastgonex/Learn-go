package tree

import "fmt"

// 结构体定义
// 改成大写则表示 public
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 虽然没有构造方法， 但是Go中可以写 工厂函数
func createNode(value int) *TreeNode {
	// 在c++中， 返回一个局部变量的地址会出问题
	// 但是在go语言中，可以返回一个局部变量的地址给别人使用
	return &TreeNode{Value: value}
}

// 为结构体写方法
// 括号中的是 接收者， 表示 print()是给node来接收的
func (node TreeNode) print() {
	fmt.Printf("%v\n", node.Value)
}

// go语言中参数是传值的，因此需要加一个*， 传递指针
// go语言的之间引用和间接引用都是. 不需要 ->
func (node *TreeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	// nil指针虽然不报错， 但是不能将这个value拿出来... 因此上面要return
	node.Value = value
}

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.print()
	node.Left.traverse()
	node.Right.traverse()
}
