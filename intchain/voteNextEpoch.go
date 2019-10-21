package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type VoteParams struct {
	From     string `josn:"from"`
	VoteHash string `json:"vote_hash"`
}

func main() {
	voteNextEpoch()
}

func voteNextEpoch() {
	var r config.VoteRPC
	params := VoteParams{
		From:     "INT3MccCA7EtMzijJa2zjxoiSYzbNLE4",
		VoteHash: "0x0911233283683a3147056ab038726d2af60bde0d41f6910084c69e55f0b984f4", // amount "0x152d02c7e14af6800000"  salt "like"
	}

	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "tdm_voteNextEpoch",
		"params":  []interface{}{params.From, params.VoteHash},
		"id":      "1",
	}

	resp, err := utils.RpcRequest(postData)
	if err != nil {
		fmt.Printf("vote 失败 err %v\n", err)
	} else {
		fmt.Printf("resp %v\n", resp.Body)
		err := json.Unmarshal([]byte(resp.Body), &r)
		if err != nil {
			fmt.Printf("解析出错 err %v\n", err)
		} else {
			fmt.Printf("vote 成功 resutl %v\n", r.Result)
		}

	}

}
