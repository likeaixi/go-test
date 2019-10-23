package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type GetTransactionCountParams struct {
	From string `json:"from"`
	Tag  string `json:"tag"`
}

func main() {
	getTransactionCount()
}

func getTransactionCount() {
	var r config.DelegateRPC
	params := GetTransactionCountParams{
		From: "INT3EYrA3Z3wA1P3AhWyvbRFPyPznqbR",
		Tag:  "latest",
	}
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionCount",
		"params":  []interface{}{params.From, params.Tag},
		"id":      "1",
	}
	resp, err := utils.RpcRequest(postData)
	if err != nil {
		fmt.Printf("失败, err=%v\n", err)
	} else {
		fmt.Printf("resp=%v\n", resp.Body)
		err := json.Unmarshal([]byte(resp.Body), &r)
		if err != nil {
			fmt.Printf("解析出错  err=%v\n", err)
		} else {
			fmt.Printf("成功   result=%v\n", r.Result)
		}

	}
}
