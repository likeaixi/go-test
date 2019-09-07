package main // 包声明

import (
	"errors"
	"fmt"
	"time"
	//"os"
	"runtime"
) // 引入包

var name string // 不可导出
var Name string // 可导出

const(
	i1 = 1000
	pi1 = 3.1415
	prefix1 = "Go"
)
var(
	i int
	pi float32
	prefix string
)

const(
	x = iota
	y = iota
	z = iota
	w = iota
	u = iota
	v = iota
)

var arr [10]int // 声明一个数组


var m map[string]int // 声明一个 key 是字符串，值为 int类型的 map,使用之前使用 make 初始化

// 函数作为值、类型
type testInt func(int) bool // 声明一个函数类型

// 声明一个 person 类型
type person struct {
	name string
	age int
}

type Skills []string

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human   // 匿名字段，struct
	Skills  // 匿名字段，自定义类型 string slice
	int		// 内置类型作为匿名字段
	speciality string
}


func main() { // 函数
	/*这是一个很简单的程序*/ // 注释
	fmt.Printf("Hello, World\n") // 语句&表达式

	name1, name2, name3 := "like", "like2", "like3"
	fmt.Printf("HelloWorld: %s,%s,%s\n", name1, name2, name3)

	s := []byte(name1)
	s[0] = 'L'
	s2 := string(s)
	fmt.Printf("%s\n", s2)

	s3 := "L" + name1[1:] // 切片
	fmt.Printf("%s\n", s3)

	s4 := `hello
			world`
	fmt.Printf("%s\n",s4)

	b := true
	fmt.Printf("%v\n",b)

	a := 5
	d := 6
	c := a + d
	fmt.Printf("%d\n", c)

	// 错误类型 error, 包 errors
	err := errors.New("emit macho dwarf : elf header corrupted")

	if err != nil {
		fmt.Printf("%v\n",err)
	}

	i := 1234
	fmt.Printf("%v\n", i)

	j := int32(1)
	fmt.Printf("%v\n", j)

	f := float32(3.14)
	bytes := [5]byte{'h','e','l','l','o'}
	fmt.Printf("%f, %d\n", f, bytes) // 3.140000, [104 101 108 108 111]

	// 声明数组
	arr[0] = 42
	arr[1] = 43
	fmt.Printf("%d\n", arr[0])
	fmt.Printf("%d\n", arr[9]) //未赋值， 默认为0


	e := [3]int{1, 2, 3}
	g := [...]int{1, 2, 3,4} // 省略长度， Go会自动根据元素个数来计算长度
	fmt.Printf("%d, %d\n", e, g)

	doubleArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("%d\n", doubleArray)

	// 声明切片 slice，"动态数组",不指明长度, 内置函数 len, cap, append, copy
	slice := []byte{'a', 'b', 'c'}
	fmt.Printf("%d\n", slice)

	slice1 := bytes[2:5:5] // 切片可以从已经存在的数组或者 slice 中再次声明 array[i:j], 含头不含尾, 整个数组可以用 array[:],切片最后一个参数为 cap 即容量
	fmt.Printf("%d\n", slice1)

	testSlice := []string{"a", "b", "c"}

	copySlice := testSlice
	fmt.Printf("testSlice 地址：%p\n, copySlice 地址：%p\n", &testSlice, &copySlice)
	fmt.Printf("copySlice[0]= %v\n", copySlice[0])
	testSlice[0] = "d"

	fmt.Printf("copySlice[0]= %v\n", copySlice[0])






	// 声明 map
	m = make(map[string]int) // 无序，长度不固定，内置 len 方法返回 key 数量，
	fmt.Printf("%d\n", m)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3


	// 流程和函数

	// 流程控制， 1、条件判断，2、循环控制，3、无条件跳转

	if h := len(m); h > 10 {
		fmt.Printf("h is greater than 10\n")
	} else {
		fmt.Printf("h is less than 10\n")
	}

	// goto, 用goto跳转到必须在当前函数内定义的标签

	//l := 0
	//Here:
	//println(l)
	//l++
	//goto Here
	t1, t2 := time.Now(), time.Now()
	for i := 0; i <= 50000000; i++ {
		if i == 5000000 {
			t2 = time.Now()
			t3 := t2.Second() - t1.Second()
			fmt.Printf("%d\n", t3)
		}
	}

	sum := 1

	for ; sum < 100; {
		sum += sum
		if sum > 50 {
			continue
		}
	}
	fmt.Printf("%d\n", sum)

	sum1 := 1
	for sum1 < 100 {
		if sum1 > 50 {
			break
		}
		sum1 += sum1
	}
	fmt.Printf("%d\n", sum1)

	// for 配合 range 可以读取 slice 和 map 的数据

	for k, v := range m {
		fmt.Printf("m的 key 为%s ", k)
		fmt.Printf("m的 val 为%d\n", v)
	}

	// _用来丢弃不需要的返回值
	for _, v := range m {
		fmt.Printf("m的 val 为%d\n", v)
	}

	// switch
	l := 10

	switch l {
	case 1:
		fmt.Printf("l is equal to 1\n")
		fallthrough
	case 2:
		fmt.Printf("l is equal to 2\n")
		fallthrough
	case 3:
		fmt.Printf("l is equal to 3\n")
		fallthrough
	case 10:
		fmt.Printf("l is equal to 10\n")
		fallthrough
	default:
		fmt.Printf("default case\n")
	}


	// 函数 用func 声明

	o, p := 3, 4
	maxXY := max(o, p)
	fmt.Printf("max(%d, %d) = %d\n", o, p, maxXY)
	fmt.Printf("max(%d, %d) = %d\n", o, p, max(o,p))

	add, multiplied := SumAndProduct(o,p)
	fmt.Printf("SumAndProduct(%d, %d) = %d, %d\n", o, p, add, multiplied)

	myfunc(3,4,5,6)

	// 传值与传指针

	q := 1
	fmt.Printf("q = %d\n", q)
	q1 := add1(q)
	fmt.Printf("q1 = %d\n", q1)
	fmt.Printf("q = %d\n", q)

	q2 := add2(&q)
	fmt.Printf("q2 = %d\n", q2)
	fmt.Printf("q = %d\n", q)


	// defer 在函数退出前执行
	//for i := 0; i < 10; i++ {
	//	defer fmt.Printf("%d\n", i)
	//}

	// 函数作为值来传递

	ss := []int{1,2,3,4,5,6,7,8}
	fmt.Println("slice", ss)

	odd := filter(ss, isOdd)
	fmt.Println("odd elements of ss are: ", odd)

	even := filter(ss, isEven)
	fmt.Println("even elements of ss are: ", even)


	// Panic 和 Recover
	//var user = os.Getenv("USER")
	//fmt.Printf("user is %s\n", user)
	//if user != "" {
	//	panic("no value for $USER")
	//}

	// struct

	var P person // P 就是 person 类型的变量
	P.age = 18
	P.name = "like"
	fmt.Printf("The person's name is %s\n", P.name)

	// 其它初始化方式
	Q := person{"like", 25}
	fmt.Println(Q)

	// 无序初始化
	R := person{age: 18, name: "like"}
	fmt.Println(R)

	// 通过 new 函数分配一个指针

	S := new(person)
	fmt.Println(*S)

	// 匿名字段
	// 初始化一个学生 Jane
	jane := Student{Human:Human{"Jane", 18, 120}, speciality: "Biology"}

	fmt.Println("her name is ", jane.name)
	fmt.Println("her age is ", jane.age)
	fmt.Println("her weight is ", jane.weight)
	fmt.Println("her speciality is ", jane.speciality)

	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are", jane.Skills)
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills are ", jane.Skills)

	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)

	// 面向对象
	//method



	// 并发
	go say("World") // 3-5次
	say("Hello") // 5次

	// channels
	ci := make(chan int)
	//cs := make(chan string)
	//cf := make(chan interface{})


	ci <- 1
	ci <- 2




	sa := []int{1,2,3,4,5,6}

	ca := make(chan int)

	go sum22(sa[:len(sa)/2], ca)
	go sum22(sa[len(sa)/2:], ca)

	//close(ca)
	x, y := <-ca, <-ca
	fmt.Println(x,y,x+y)

	//for i := range ca {
	//	fmt.Println(i)
	//}

	// select
	//c1 := make(chan int)
	//quit := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c1)
	//	}
	//	quit <-0
	//}()
	//fibonacci(c1, quit)

	// 超时
	gc := make(chan int)
	goo := make(chan bool)

	go func() {
		for {
			select {
				case v := <- gc:
					fmt.Println(v)
				case <- time.After(5 * 1000 * time.Millisecond):
					fmt.Println("timeout")
					goo <- true
					break
			}
		}
	}()
	<- goo
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 多个返回值
func SumAndProduct(a, b int) (add int, multiplied int) {
	add = a + b
	multiplied = a * b
	return
}

// 变参, 接受不定数量的参数
func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("the number is %d\n", n)
	}
}

func add1(a int) int {
	a += 1
	return a
}

func add2(a *int) int {
	*a += 1
	return *a
}

func isOdd(i int) bool {
	if i % 2 == 0 {
		return false
	}
	return true
}

func isEven(i int) bool {
	if i % 2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, val := range slice {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

// 下面这个函数检查作为其参数的函数在执行时是否产生 panic
func throwPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f, 如果 f 中出现了 panic，那么就可以恢复回来
	return
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func sum22(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y := y, x + y
			fmt.Println(x,y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}