package main

import (
	"fmt"
	"github.com/mikemintang/go-curl"
)

func main() {
	//url := "http://127.0.0.1:7000/intchain"
	url := "https://mainnet.infura.io/bWQAsi2JbfmO9YAoxOgm"
	headers := map[string]string{
		"Content-type": "application/json",
	}
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		//"method":  "eth_getTransactionByHash",
		"method": "eth_getTransactionReceipt",
		"params": []interface{}{"0x7f15a26618873a799ae6427a215727fe93dbe52c8fe85d0c4702f49913f35912"},
		"id":     "1",
	}

	req := curl.NewRequest()

	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetPostData(postData).
		Post()
	if err != nil {
		fmt.Printf("获取交易失败，err=%v\n", err)
	} else {
		if resp.IsOk() {
			fmt.Printf("获取交易成功，body=%v\n", resp.Body)
		} else {
			fmt.Printf("获取交易失败， raw=%v\n", resp.Raw)
		}
	}

}
