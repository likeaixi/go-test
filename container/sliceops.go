package container

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	var s []int // zero value for slice is nil
	fmt.Println(s)

	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, 2 * i + 1)
	}

	fmt.Println(s)

	// 知道值
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	// 不知道值, 知道个数, 容量
	s2 :=  make([]int, 16)
	s3 :=  make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	// 复制切片
	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)


	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]

	fmt.Println(tail)
	printSlice(s2)
}
