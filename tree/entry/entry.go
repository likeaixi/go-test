package main

import (
	"fmt"
	"go-test/tree"
	"golang.org/x/tools/container/intsets"
)

// 包如何扩充系统类型或者别人的类型

// 定义别名, 见 queue/entry/main.go


// 使用组合
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}


// go get 获取第三方库
// gopm 来获取无法下载的包

// go build 来编译下载的包
// go install 产生 pkg 文件和可执行文件
// go run 直接编译运行

func testSpase()  {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(2)

	fmt.Println(s.Has(1))
	fmt.Println(s.Has(2))
}

// go 语言公支持封装，不支持继承和多态

// go 语言没有 class, 只有 struct

// 无论地址还是结构本身，一律使用 . 访问

func main() {
	var root tree.Node
	fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}

	root.Right.Left = new(tree.Node)

	fmt.Println(root.Right.Left)

	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Right.Left.Print()

	//root.SetValue(200)
	root.Print()
	fmt.Println(root)

	//nodes := []tree.Node{
	//	{Value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	root.Traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()


	testSpase()
}
