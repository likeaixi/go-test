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
	// amount "0x54b40b1f852bda000000"  salt "like"
	var voteList = []*VoteParams{
		{From: "INT3LJK4UctyCwv5mdvnpBYvMbRTZBia", VoteHash: "0xe65ffe860e86a567086fa58c136f81e3d4fd3a12dd6492e8f39e86b6ebde3716"},
		{From: "INT3JF1CSRxna54ukUTgyew1VyUppGcD", VoteHash: "0x7337af0cbf11c804af66b570340b2ae75802464e9d2a9d3a64a967fded50e33f"},
	}

	for _, v := range voteList {
		voteNextEpoch(v.From, v.VoteHash)
	}
}

func voteNextEpoch(from, hash string) {
	var r config.VoteRPC
	params := VoteParams{
		From:     from,
		VoteHash: hash,
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
