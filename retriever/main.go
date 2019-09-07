package main

import (
	"fmt"
	"go-test/retriever/mock"
	"go-test/retriever/real"
	"time"
)


// 接口实现是隐式的，使用者定义接口
// 只要实现接口里面的方法


// 接口变量包含 实现者的类型和实现者的值/指针
// 接口变量自带指针
// 接口变量同样采用值传递，几乎不需要使用接口的指针
// 指针接收者实现只能以指针方式使用；值接收者都可

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)


	fmt.Printf("> %T, %v\n", r, r)
	fmt.Print("> Try switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UseerAgent:", v.UserAgent)
	}
	fmt.Println()
}

const url = "http://www.imooc.com"

type Poster interface {
	Post(url string, from map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name": "like",
			"source": "golang",
	})
}

// 组合接口
type RetrieverPoster interface {
	Retriever
	Poster
}


func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another faked imooc.com"})
	return s.Get(url)
}

// 查看接口变量
// 表示任何类型: interface{}
// Type Assertion
// Type Switch
func main() {
	var r Retriever
	retriever := &mock.Retriever{"this is a fake badu"}

	r = retriever
	fmt.Printf("%T, %v\n", r, r)
	r = &real.Retriever{"Mozilla/5.0", time.Minute}
	fmt.Printf("%T, %v\n", r, r)


	fmt.Println("Try session")
	fmt.Println(session(retriever))

	inspect(r)

	// type assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	}


	// 标准库里面用了很多组合接口
	//var wc io.WriteCloser




	// 常用系统接口
	/*
		1、Stringer
		2、Reader
		3、Writer
	*/

}
