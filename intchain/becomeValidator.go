package main

import (
	"fmt"
	"github.com/mikemintang/go-curl"
)

func main() {
	getCurrentEpochNumber()
}

func getCurrentEpochNumber() (number string) {
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params":  []interface{}{},
		"id":      "1",
	}

	//resp, err := rpcRequest(postData)
	//if err == nil {
	//	if resp.IsOk() {
	//
	//	}else {
	//		fmt.Printf()
	//	}
	//}else {
	//	fmt.Printf("获取Epoch失败：err=%v\n", err)
	//}

	return "0x1"
}

//func getEpoch(url string, headers map[string]string)(epoch Epoch) {
//
//}

func rpcRequest(postData map[string]interface{}) (resp *curl.Response, err error) {
	url := "http://127.0.0.1:7000/intchain"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	req := curl.NewRequest()

	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	if err == nil {

	} else {

	}
}
