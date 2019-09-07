package tree

// 只有使用指针才可以改变结构体内容, nil 指针也可以调用方法
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	// 所以这里不用判断这里的left 和 right 是否为空指针
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}