package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		//panic("unsuppoted op:" + op)
		return 0, fmt.Errorf("unsuppoted op: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
	//q = a / b
	//r = a % b
	//return
}

func apply(op func(int, int) int, a, b int) int {
	p :=reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" + "(%d, %d)", opName, a ,b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap1(a, b int) (int, int) {
	return b, a
}


// 指针
func swap2(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	fmt.Println(eval(3,4, "x"))

	result, e := eval(3, 4, "x")
	if e != nil {
		fmt.Println("error:", e)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Println(q, r)

	//fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a, b int) int{
			return int(math.Pow(float64(a), float64(b)))
		},3, 4,
		))
	fmt.Println(sum(1,2,3,4,5))

	a, b := 3, 4
	a, b = swap1(a, b)
	fmt.Println(a, b)
}
