package main

import (
	"fmt"
	"sync"
)

// 不要通过共享内存来通信，通过通信来共享内存，添加一个 channel 来通信


func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)

	return w
}

func chanDemo() {
	//var c chan int // c == nil
	var wg sync.WaitGroup // 系统提供的 WaitGroup 来实现多人等待

	var workers [10]worker  // 只能收数据
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}



	wg.Add(20)
	for i , worker := range workers {
		worker.in <- 'a' + i
	}
	//for _, worker := range workers {
	//	<-worker.done
	//}

	for i , worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()

	//for _, worker := range workers {
	//	<-worker.done
	//}
	//time.Sleep(time.Millisecond)
}


func main() {
	chanDemo() // channel 可以做为一等公民
}
