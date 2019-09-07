package container

import "fmt"

// 内建容器
// 数组是值类型，作为参数会拷贝数组
func printArray(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}


func printSlice(arr []int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3, grid)

	for _, v := range arr3 {
		fmt.Println(v)
	}

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	printArray(&arr1)
	printArray(&arr3)

	printSlice(arr1[:])
	printSlice(arr3[:])

	fmt.Println(arr1, arr2, arr3)
}
