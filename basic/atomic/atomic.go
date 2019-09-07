package main

import (
	"fmt"
	"sync"
	"time"
)

// 传统的同步机制 （通过共享内存进行通信）

// WaitGroup
// Mutex
// Cond

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increatement() {
	fmt.Println("safe increatement")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++

	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increatement()
	go func() {
		a.increatement()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
