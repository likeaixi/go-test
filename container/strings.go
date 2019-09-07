package main

import (
	"fmt"
	"unicode/utf8"
	//"strings"
)

// Fields(可以认空格), Split, Join
// Contains, Index
// ToLower, ToUpper
// Trim, TrimRight, TrimLeft

func main() {
	s := "Yes我是中国人!" // UTF-8
	fmt.Println(len(s))
	fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) { // []byte 获得字节
		fmt.Printf("%x ", b)
	}

	fmt.Println()

	for i, ch := range s { // ch is a rune （int32), 4字节整数， 进行 utf-8解码，然后转为 unicode
		fmt.Printf("(%d %x)\n", i, ch)
	}

	fmt.Println("Rune count：", utf8.RuneCountInString(s)) // 获得字符数量
	
	bytes := []byte(s)
	for len(bytes) > 0 { // len 获取字节长度
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c\n", ch)
	}


	for i, ch := range []rune(s) { // 使用 range 遍历 pos, rune 对
		fmt.Printf("(%d, %c)", i, ch)
	}

}
