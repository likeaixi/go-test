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
		//"method":  "eth_getTransactionByHash",
		"method": "eth_getTransactionReceipt",
		"params": []interface{}{"0x02ee76e2d6683fbc63af2cf754e5bbe9730efe65f6e4e452dc67029200da182e"},
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
