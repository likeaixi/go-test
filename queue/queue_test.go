package queue

import (
	"fmt"
	"go-google/queue"
)

// 测试例子
func ExampleQueue_Pop() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())


	// output:
	// 1
	// 2
	// false
	// 3
	// true
}
