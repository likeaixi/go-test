package container

import "fmt"


// 内建容器
// slice 切片, slice本身没有数据，是对底层 array的一个view

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])

	s1 := arr[2:]
	fmt.Println("s1", s1)
	s2 := arr[:]
	fmt.Println("s2", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	//fmt.Println(s2)
	//fmt.Println(s2[:5])
	//fmt.Println(s2[:5][2:])

	fmt.Println("Extending slice")
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2) // slice 可以向后扩展，不可以向前扩展，s[i]不可以超越 len(s), 向后扩展不可以超越底层数组 cap(s)

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",s2, len(s2), cap(s2))


	// 添加元素时，如果超越 cap(s)，则系统会重新分配更大的底层数组
	// 由于是值传递，必须接收 append 返回值
	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)

	// s4, s5 no longer view arr
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	fmt.Println("arr = ", arr)
}
