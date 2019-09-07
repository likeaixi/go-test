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
	From  	 string   `json:"from"`
	To    	 string   `json:"to"`
	Gas   	 string   `json:"gas"`
	GasPrice string   `json:"gasPrice"`
	Value    string   `json:"value"`
	Data     string   `json:"data"`

}

func main() {
	//s := `curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0x37eeb099c6d751e7229df825d40629612e134f82","to": "0x6ee600fc8e27c562a63ff5e56d1b788bca6f5c2e","gas": "0x76c0","gasPrice": "0x2540be400","value": "0x1","data": ""}],"id":1}' -H 'content-type: application/json;' http://127.0.0.1:7000/intchain`

	url := "http://127.0.0.1:8545/intchain"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	params := TxParams{
		//From: "0x37eeb099c6d751e7229df825d40629612e134f82",
		From: "0x52bc4024a38f88d2ee1afde012eedead3bc960bf",
		To: "0x6ee600fc8e27c562a63ff5e56d1b788bca6f5c2e",
		Gas: "0x76c0",
		GasPrice: "0x2540be400",
		Value: "0x1",
		Data: "",
	}

	postData := map[string]interface{} {
		"jsonrpc": "2.0",
		"method": "eth_sendTransaction",
		"params": []interface{}{params},
		"id": "1",
	}

	for i := 0; i < 1000000; i++ {
		req := curl.NewRequest()
		resp, err := req.
			SetUrl(url).
			SetHeaders(headers).
			SetPostData(postData).
			Post()

		if err != nil {
			fmt.Printf("交易发送失败， err=%v\n",err)
		}else {
			if resp.IsOk() {
				fmt.Printf("交易发送成功, body=%v\n", resp.Body)
			} else {
				fmt.Printf("交易发送失败, raw=%v\n", resp.Raw)
			}
		}
	}

}
