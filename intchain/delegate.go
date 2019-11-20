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
	var fromList = []string{"INT39NQ6EoRUqK6ypvmqPx7j7ZsskGN4", "INT3HQ652WaqNVhuDAYiNd1iKRbgBeo3", "INT3MxpSzfLKEQDQVG9ZruyjivKR4Dpz"}

	for _, v := range fromList {
		for i, p := range config.PrivValidators {
			if i == 0 {

				hash, err := delegate(v, p.Address)
				if err != nil {
					fmt.Printf(" delegate err %v\n", err)
				}

				fmt.Printf("delegate hash %v\n", hash)
			}
		}
	}

}

func delegate(from, candidate string) (hash string, err error) {
	var r config.DelegateRPC
	params := DelegateParams{
		From:      from,
		Candidate: candidate,
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
		return "", err
	}

	err = json.Unmarshal([]byte(resp.Body), &r)
	if err != nil {
		return "", err
	}

	return r.Result, nil
}
