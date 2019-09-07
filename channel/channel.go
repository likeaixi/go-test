package main

import (
	"fmt"
	"time"
)

// 不要通过共享内存来通信，通过通信来共享内存

//func worker(id int, c chan int) {
//	go func() {
//		for {
//			fmt.Printf("Worker %d received %c\n", id, <- c)
//		}
//	}()
//}

func worker(id int, c chan int) {
	//for {
	//	n, ok := <- c
	//	if !ok {
	//		break
	//	}
	//
	//	fmt.Printf("Worker %d received %d\n", id, n)
	//}

	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func chanDemo() {
	//var c chan int // c == nil

	var channels [10]chan<- int  // 只能收数据
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) // 带缓冲区的 channel

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo() // channel 可以做为一等公民
	//bufferedChannel() // 带缓冲区的 channel
	//channelClose() // 带 close 的 channel
}
