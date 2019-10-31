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
	//var fromList = []string{"INT3DspEEVQ3ngFuY338jAoEcNw61V1j", "INT39L38bTaQkmzoC6mTrjyb2hGy85J6", "INT3FkMb49rivDxEjVZ9a9HuAksL6iDR"}
	var fromList = []string{"INT3ChZyJUaziuVRhP8wDm3utHDAvQgt", "INT3Paf5F2XoQqwWYaxoib6NfnDoBUgU", "INT3D78CcNArCwSQAzAfyXZPedQLbBS4"}

	for _, v := range fromList {
		delegate(v)
	}

}

func delegate(from string) {
	var r config.DelegateRPC
	params := DelegateParams{
		From:      from,
		Candidate: "INT3JF1CSRxna54ukUTgyew1VyUppGcD",
		Amount:    "0x152d02c7e14af6800000", // 100000e18
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
			fmt.Printf("解析出错  err=%v\n", err)
		} else {
			fmt.Printf("delegate 成功   result=%v\n", r.Result)
		}

	}
}
