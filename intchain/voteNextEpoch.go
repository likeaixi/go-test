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
		From:     "INT3MUHiVzxaNdG1RAD7zQimzSZBtErX",
		VoteHash: "0xbc9263f6e9fde4a9dba170a2b748d1e3e35055a7e0a98164d0c9cf00e6e12dd3", // amount "0x54b40b1f852bda000000"  salt "like"
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
