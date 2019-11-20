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
	for i, v := range config.PrivValidators {
		sign, err := signAddress(v.Address, v.ConsPrivKey)

		if err != nil {
			fmt.Printf("sign address err %v\n", err)
		} else if i == 0 || i == 1 || i == 5 {
			hash, err := revealVote(v.Address, v.ConsPubKey, sign)
			if err != nil {
				fmt.Printf("reveal vote err %v\n", err)
			}

			fmt.Printf("hash %v\n", hash)
		}
	}

}

func signAddress(from, priv string) (sign string, err error) {
	var r config.SignAddressRPC

	params := SignAddress{
		From:    from,
		PrivKey: priv,
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

func revealVote(from, pubkey, sign string) (hash string, err error) {
	var r config.RevealVoteRPC
	params := RevealVoteParams{
		From:      from,
		PubKey:    pubkey,
		Amount:    "0x54b40b1f852bda000000",
		Salt:      "intchain",
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
		err := json.Unmarshal([]byte(resp.Body), &r)
		if err != nil {
			return "", err
		} else {
			return r.Result, nil
		}

	}

}
