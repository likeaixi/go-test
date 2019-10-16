package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type DelegateParams struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func main() {
	delegate()
}

func delegate() {
	var r config.DelegateRPC
	params := DelegateParams{
		From:   "3LPkVNtCACVZTRtQV4xNA1WefbiiNoxvd3",
		To:     "32H8py5Jg396p7QNDUwTwkeVod15ksxne5",
		Amount: "0x3635c9adc5dea00000", // 1000e18
	}
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "del_delegate",
		"params":  []interface{}{params.From, params.To, params.Amount},
		"id":      "1",
	}
	resp, err := utils.RpcRequest(postData)
	if err != nil {
		fmt.Printf("delegate失败, err=%v\n", err)
	} else {
		fmt.Printf("resp=%v\n", resp.Body)
		err := json.Unmarshal([]byte(resp.Body), &r)
		if err != nil {
			fmt.Printf("delegate 成功 result=%v\n", r.Result)
		}

	}
}
