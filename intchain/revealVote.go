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
	fmt.Printf("sign %v\n", sign)
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
		From:    "INT3MUHiVzxaNdG1RAD7zQimzSZBtErX",
		PrivKey: "0xFD20E8F1E3D6FBD9056B15DB8692A4AD4DF211825C7AEDE1D562B42A3DF23EB7",
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
		From:      "INT3MUHiVzxaNdG1RAD7zQimzSZBtErX",
		PubKey:    "372EC175F6D52B91FF43493AC9E113790C0CC6AD6562A1AD7E581717C40F017A6587F3FED9DFA590CED46AE37B99617F06DB3C36BE68F1AB9005276439AB44A455DDE3C1472467D84F851B9974D46A6A525E24638D5101BE207258D91983C03A2FB021C4C50DFE95E6169C698879ED9EA4A6089B24C596C17F46489823EA7CF7",
		Amount:    "0x54b40b1f852bda000000",
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
