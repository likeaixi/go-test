package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type VoteParams struct {
	From     string `json:"from"`
	VoteHash string `json:"vote_hash"`
}

type VoteHashParams struct {
	From   string `jsn:"from"`
	PubKey string `json:"pub_key"`
	Amount string `json:"amount"`
	Salt   string `json:"salt"`
}

func main() {
	for i, v := range config.PrivValidators {
		hash, err := getVoteHash(v.Address, v.ConsPubKey)
		if err != nil {
			fmt.Printf("get vote hash err %v\n", err)
		} else if i == 1 || i == 5 {
			h, err := voteNextEpoch(v.Address, hash)
			if err != nil {
				fmt.Printf("vote next epoch err %v\n", err)
			}
			fmt.Printf("vote next epoch hash %v\n", h)
		}

	}
}

func getVoteHash(from, pubkey string) (hash string, err error) {
	var r config.VoteRPC
	params := VoteHashParams{
		From:   from,
		PubKey: pubkey,
		Amount: "0x54b40b1f852bda000000",
		Salt:   "intchain",
	}

	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "tdm_getVoteHash",
		"params":  []interface{}{params.From, params.PubKey, params.Amount, params.Salt},
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

func voteNextEpoch(from, hash string) (h string, err error) {
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
		return "", err
	}

	err = json.Unmarshal([]byte(resp.Body), &r)

	if err != nil {
		return "", err
	}

	return r.Result, nil

}
