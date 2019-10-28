package main

import (
	"fmt"
	"github.com/mikemintang/go-curl"
	"go-test/intchain/config"
)

func main() {
	//url := "https://mainnet.infura.io/bWQAsi2JbfmO9YAoxOgm"
	url := config.RemoteConfig.RpcUrl
	headers := config.ConHeaders
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionByHash",
		//"method": "eth_getTransactionReceipt",
		"params": []interface{}{"0xec0ac17e85bef3b08ae3cf1cc1ef4a6adc7268f613b8704913b519304147372a"},
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
