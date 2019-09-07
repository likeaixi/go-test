package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /=2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	//scanner := bufio.NewScanner(file)
	//
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}

	printFileContecnt(file)
}

func printFileContecnt(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forver() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)

	printFile("basic/branch/abc.txt") // 参数是文件名，不太好，修改为 Reader/Writer

	s := `abacdsf

		sdfasfs
		saf	s	s`

	printFileContecnt(strings.NewReader(s))
	printFileContecnt(bytes.NewReader([]byte(s)))

	//forver()
}
