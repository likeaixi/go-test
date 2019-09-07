package main

import "fmt"

func tryDefer() {
	defer fmt.Println(1)
	fmt.Println(2)
}

func main() {
	tryDefer()
}
