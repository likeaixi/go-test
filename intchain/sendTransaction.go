package main

import (
	"fmt"
	"github.com/mikemintang/go-curl"
)

//type postData struct {
//	JsonRPC string    `json:"jsonrpc"`
//	Method  string    `json:"method"`
//	Params  []Params  `json:"params"`
//	ID      string    `json:"id"`
//}
//
type TxParams struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data"`
}

func main() {
	//s := `curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0x37eeb099c6d751e7229df825d40629612e134f82","to": "0x6ee600fc8e27c562a63ff5e56d1b788bca6f5c2e","gas": "0x76c0","gasPrice": "0x2540be400","value": "0x1","data": ""}],"id":1}' -H 'content-type: application/json;' http://127.0.0.1:7000/intchain`

	url := "http://127.0.0.1:8545/intchain"
	//url := "http://127.0.0.1:7000/intchain"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// 部署合约的交易
	params := TxParams{
		//From: "0x37eeb099c6d751e7229df825d40629612e134f82",
		From:     "0x52bc4024a38f88d2ee1afde012eedead3bc960bf",
		To:       "",
		Gas:      "0x4c4b40",
		GasPrice: "",
		Value:    "",
		Data:     "0x6080604052348015600f57600080fd5b5060ae8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c6888fa114602d575b600080fd5b605660048036036020811015604157600080fd5b8101908080359060200190929190505050606c565b6040518082815260200191505060405180910390f35b600060078202905091905056fea265627a7a723158201e079298dbe44bac0461daf3f6a30cba5938e93c6b3dcf37204dda1c04a9567164736f6c634300050b0032",
	}

	sendTransaction(url, headers, params)

	//params = TxParams{
	//	//From: "0x37eeb099c6d751e7229df825d40629612e134f82",
	//	From:     "0x52bc4024a38f88d2ee1afde012eedead3bc960bf",
	//	To:       "0x6ee600fc8e27c562a63ff5e56d1b788bca6f5c2e",
	//	Gas:      "0x76c0",
	//	GasPrice: "0x2540be400",
	//	Value:    "0x1",
	//	Data:     "",
	//}
	//
	//sendTransactions(url, headers, params)

}

func sendTransaction(url string, headers map[string]string, params TxParams) (hash string, err error) {
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_sendTransaction",
		"params":  []interface{}{params},
		"id":      "1",
	}

	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	if err != nil {
		fmt.Printf("交易发送失败， err=%v\n", err)
	} else {
		if resp.IsOk() {
			fmt.Printf("交易发送成功, body=%v\n", resp.Body)
		} else {
			fmt.Printf("交易发送失败, raw=%v\n", resp.Raw)
		}
	}

	return "", nil
}

func sendTransactions(url string, headers map[string]string, params TxParams) (hash string, err error) {
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_sendTransaction",
		"params":  []interface{}{params},
		"id":      "1",
	}

	for i := 0; i < 100000; i++ {
		req := curl.NewRequest()
		resp, err := req.
			SetUrl(url).
			SetHeaders(headers).
			SetPostData(postData).
			Post()

		if err != nil {
			fmt.Printf("交易发送失败， err=%v\n", err)
		} else {
			if resp.IsOk() {
				fmt.Printf("交易发送成功, body=%v\n", resp.Body)
			} else {
				fmt.Printf("交易发送失败, raw=%v\n", resp.Raw)
			}
		}
	}

	return "", err
}
