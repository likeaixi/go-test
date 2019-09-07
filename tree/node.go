package  tree

import "fmt"



// c 语言变量创建在栈上，函数退出，变量立即销毁，分配在堆上，需要手动释放，  Java 使用 new  大部分分配在堆上，有垃圾回收
// 结构体创建在堆上还是栈上？ 不需要知道！ 栈上不参与垃圾回收，堆上参与垃圾回收
type Node struct {
	Value int
	Left, Right *Node
}

// 一个目录一个包, main包 包含可执行入口
// 大写 public
// 小写 private

// 结构定义的方法必须全部放在同一个包内，可以是不同的文件

// 结构体的方法，前面有一个接收者
func (node Node) Print() {
	fmt.Print(node.Value)
}


// 值接收者 vs 指针接收者

// 要改变内容必须使用指针接收者
// 结构过大也要考虑使用指针接收者
// 一致性建议：如果有指针接收者，最好都是指针接收者

// 值接收者是 go 语言特有
// 值/指针接收者均可接收值/指针




// 只有使用指针才可以改变结构体内容, nil 指针也可以调用方法
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}

	node.Value = value
}


func CreateNode(value int) *Node { // 自定义工厂函数， 返回局部变量的地址
	return &Node{Value: value}
}



