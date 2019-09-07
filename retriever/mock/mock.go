package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	//panic("implement me")
	return fmt.Sprintf("Retriever : {Contents=%s}", r.Contents)
}

// 如果是值接收者，不会修改 Contents, 改为指针接收者
func (r *Retriever) Post(url string, from map[string]string) string {
	//panic("implement me")
	r.Contents = from["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}


