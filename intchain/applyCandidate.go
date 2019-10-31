package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type CandidateParams struct {
	From            string `json:"from"`
	SecurityDeposit string `json:"security_deposit"`
	Commission      int    `json:"commission"`
}

func main() {
	applyCandidate()
}

func applyCandidate() {
	var r config.CandidateRPC
	params := CandidateParams{
		From:            "INT3JF1CSRxna54ukUTgyew1VyUppGcD",
		SecurityDeposit: "0x152d02c7e14af6800000", // 1e+23
		Commission:      10,                       // 0-100
	}
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "del_applyCandidate",
		"params":  []interface{}{params.From, params.SecurityDeposit, params.Commission},
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
