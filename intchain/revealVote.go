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
	var signList = []*SignAddress{
		{From: "INT3LJK4UctyCwv5mdvnpBYvMbRTZBia", PrivKey: "0xE01D15AA2E7A90EDEA0D1B02ACA531822D7733076962CA55F3C7D1D8FFA8DE1F"},
		{From: "INT3JF1CSRxna54ukUTgyew1VyUppGcD", PrivKey: "0x5F4F3EA977E2F4C5B12F56E6197427F2C0902B3254267BFD7C606F965E6005DD"},
	}

	var revealVoteList = []*RevealVoteParams{
		{From: "INT3LJK4UctyCwv5mdvnpBYvMbRTZBia", PubKey: "20ADAF53F032DF36DBE8C8672EAC5DE273DADE49D0503ED2F0826FC47894447F4620F9CA656A3A9EAC3DB2B6657897A6A6ED6415135A0E84DD3A28A75564264F7CCF03F3035EE5260A752C5ADDFE3F01B709A382AAC94B47FA760112C3538A337BD7D0CD06CF615496430CC6CE29E9539EE99A49AC16E760CA9FCC5AB226BA0A", Amount: "0x54b40b1f852bda000000", Salt: "like", Signature: ""},
		{From: "INT3JF1CSRxna54ukUTgyew1VyUppGcD", PubKey: "8AF71CD7B4BC233191F3B453E5E02E8E62B895158A0585DF6178719EF3BA6E900CB294DE120C5C7A3A491C20BB878E78EE750DD3240EC89DA22F4DD39F1B2D1B89F2BEA1EFD1E2EC363202A8D5EFD95B5F56D3C458B2C1AA3EFE74D0EF5FF7C0603ED613F79B9566C1378B49EFF5B6137B583066BEE04D313ABB697B94FA8722", Amount: "0x54b40b1f852bda000000", Salt: "like", Signature: ""},
	}

	for i, v := range signList {
		sign, err := signAddress(v.From, v.PrivKey)
		fmt.Printf("sign %v\n", sign)
		if err != nil {
			fmt.Printf("sign address err %v\n", err)
		} else {
			hash, err := revealVote(revealVoteList[i].From, revealVoteList[i].PubKey, revealVoteList[i].Amount, revealVoteList[i].Salt, sign)
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

func revealVote(from, pubkey, amount, salt, sign string) (hash string, err error) {
	var r config.RevealVoteRPC
	params := RevealVoteParams{
		From:      from,
		PubKey:    pubkey,
		Amount:    amount,
		Salt:      salt,
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
