package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type DelegateParams struct {
	From      string `json:"from"`
	Candidate string `json:"candidate"`
	Amount    string `json:"amount"`
}

func main() {
	delegate()
}

func delegate() {
	var r config.DelegateRPC
	params := DelegateParams{
		From:      "INT361cPNXzfLP1Rn52vitzWBU5x5Nz7",
		Candidate: "INT3DWAx9WMF7gscfaHGFByTEpjBt5Xs",
		Amount:    "0x3635c9adc5dea00000", // 1000e18
	}
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "del_delegate",
		"params":  []interface{}{params.From, params.Candidate, params.Amount},
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
