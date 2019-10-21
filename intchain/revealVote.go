package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type SignAddress struct {
	From    string `json:"from"`
	PrivKey string `json:"priv_key"`
}

type RevealVoteParams struct {
	From      string `josn:"from"`
	PubKey    string `json:"pub_key"`
	Amount    string `json:"amount"`
	Salt      string `json:"salt"`
	Signature string `json:"signature"`
}

func main() {
	sign, err := signAddress()
	if err != nil {
		fmt.Printf("sign address err %v\n", err)
	} else {
		hash, err := revealVote(sign)
		if err != nil {
			fmt.Printf("reveal vote err %v\n", err)
		}
		fmt.Printf("hash %v\n", hash)
	}

}

func signAddress() (sign string, err error) {
	var r config.SignAddressRPC

	params := SignAddress{
		From:    "INT3MccCA7EtMzijJa2zjxoiSYzbNLE4",
		PrivKey: "719D72D7CE74236FF967954960B086C9A15B5512D506268E855B947269B25587",
	}

	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "chain_signAddress",
		"params":  []interface{}{params.From, params.PrivKey},
		"id":      "1",
	}

	resp, err := utils.RpcRequest(postData)
	if err != nil {
		return "", err
	} else {
		fmt.Printf("resp %v\n", resp.Body)
		err := json.Unmarshal([]byte(resp.Body), &r)
		if err != nil {
			return "", err
		} else {
			return r.Result, nil
		}

	}
}

func revealVote(sign string) (hash string, err error) {
	var r config.RevealVoteRPC
	params := RevealVoteParams{
		From:      "INT3MccCA7EtMzijJa2zjxoiSYzbNLE4",
		PubKey:    "618CEAF6AD449B826E2521222A94426B82800202332251F0929EC47B36A647C65E00D2EA34C07A8EF7953C2E1555D8321449423CCFB0B64BB13090E7A433114D68F1C1891BAA20101E5CC8E2B10E207F5D21D1A1116547E1EED5E92FDFE4F5E58119C5267B82AE06BBA5016827396B74E1ECDCC3801746242CA24C7749EB2F88",
		Amount:    "0x152d02c7e14af6800000",
		Salt:      "like",
		Signature: sign,
	}

	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "tdm_revealVote",
		"params":  []interface{}{params.From, params.PubKey, params.Amount, params.Salt, params.Signature},
		"id":      "1",
	}

	resp, err := utils.RpcRequest(postData)
	if err != nil {
		return "", err
	} else {
		fmt.Printf("resp %v\n", resp.Body)
		err := json.Unmarshal([]byte(resp.Body), &r)
		if err != nil {
			return "", err
		} else {
			return r.Result, nil
		}

	}

}
