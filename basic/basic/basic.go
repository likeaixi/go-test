package main

import (
	"fmt"
	"math"
)

var (
	aa = 3
	bb = 4
	ss = "s"
)

func variableZeroValue() {
	var (
		a int
		s string

	)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a,b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func consts() {
	const(
		filename string = "abc.txt"
	 	a, b = 3, 4
	 )
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 特殊常量，枚举类型
func enums() {
	 const(
	 	cpp = iota
	 	java
	 	python
	 	golang
	 	javascript
	 )

	 const(
	 	b = 1 << (10 * iota)
	 	kb
	 	mb
	 	gb
	 	tb
	 	pb
	 )

	 fmt.Println(cpp, javascript, java, python, golang)
	 fmt.Println(b, kb, mb, gb, tb, pb)
}


func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}


func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableShorter()
	consts()
	enums()
}
