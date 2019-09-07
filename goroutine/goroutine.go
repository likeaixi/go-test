package main

import (
	"fmt"
	"time"
)

func main() { // 也是一个 goroutine
	//var a [10]int

	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello from goroutine %d\n", i) // printf是一个 io 操作，内部有协程之间的切换，所以打的顺序不确定

				//a[i]++ // 不会自动交出控制权
				//runtime.Gosched() //手动交出控制权
			}
		}(i)
	}

	time.Sleep(time.Minute)
	//fmt.Print(a)
}
