package main

import (
	"fmt"
	"go-test/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("abc") // push 字符串会报错，接口已经限定了类型
	fmt.Println(q.Pop())
}
