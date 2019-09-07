package main

import "fmt"

func lengthOfNoRepeatSubStr(s string) int {
	//lastOccurred := make(map[byte]int)
	lastOccurred := make(map[rune]int)
	start := 0
	maxLenght := 0
	//for i, ch := range []byte(s) {
	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i - start + 1 > maxLenght {
			maxLenght = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLenght
}

func main() {
	m := map[string]string {
		"name": "like",
		"name1": "like1",
		"name2": "like2",
		"name3": "like3",
		"name4": "like4",
	}

	m2 := make(map[string]int) // empty

	var m3 map[string]int // nil

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v) // 无序
	}

	fmt.Println(m["name"])
	fmt.Println(m["nam"]) // 取不存在的 key,拿到的是 zero value

	if name, ok := m["nam"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("Key does not exist")
	}

	delete(m, "name")

	name, ok := m["name"]

	fmt.Println(name, ok)

	// map的 key
	// map 使用 hash 表，必须可以比较相等
	// 除以 slice, map, function 的内建类型都可以作为 key
	// Struct 类型不包含上述字段，也可以作为 key


	// 找出最长不含有重复字符的子串
	// abcabcaa -> abc
	// aaaa -> a

	fmt.Println(lengthOfNoRepeatSubStr("abdsfsdfs"))
	fmt.Println(lengthOfNoRepeatSubStr("abc"))
	fmt.Println(lengthOfNoRepeatSubStr(""))
	fmt.Println(lengthOfNoRepeatSubStr("aaaaa"))
	fmt.Println(lengthOfNoRepeatSubStr("abccba"))
	fmt.Println(lengthOfNoRepeatSubStr("这是测试字符串"))
	fmt.Println(lengthOfNoRepeatSubStr("一二三二一"))
}
