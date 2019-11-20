package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type TxParams struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data"`
}

func main() {
	//sendTransaction("INT39du2xJEKp6Kz55m9XXt1ZgAqgTjr", "INT3JDmnFuFP12AZZGERnBDZh6Ao44NZ")

	//sendTransactions()

	//for _, v := range config.PrivValidators {
	//	hash, err := sendTransaction("INT3CpFuk2cJ1te9WZV1w8Y3wkQCcA5Z", v.Address)
	//	if err != nil {
	//		fmt.Printf("交易发送失败， err %v", err)
	//	}
	//	fmt.Printf("hash %v\n", hash)
	//}
	//
	//for _, v := range config.DelegateAddress {
	//	hash, err := sendTransaction("INT3CpFuk2cJ1te9WZV1w8Y3wkQCcA5Z", v)
	//	if err != nil {
	//		fmt.Printf("交易发送失败， err %v", err)
	//	}
	//	fmt.Printf("hash %v\n", hash)
	//}
}

func sendTransaction(from, to string) (hash string, err error) {
	var r *config.RPCBody

	params := TxParams{
		From:     from,
		To:       to,
		Gas:      "0x76c0",
		GasPrice: "0x2540be400",
		//Value:    "0x1",
		Value: "0x33b2e3c9fd0803ce8000000", // 1000000000e+18
		//Value: "0x152d02c7e14af6800000", // 100000e+18
		Data: "",
	}

	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_sendTransaction",
		"params":  []interface{}{params},
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

	return r.Result, err
}

func sendTransactions() (hash string, err error) {
	var r *config.RPCBody

	params := TxParams{
		From:     "INT3CpFuk2cJ1te9WZV1w8Y3wkQCcA5Z",
		To:       "INT38cqij9EtpxCZk9wrRY8u5U9reZAp",
		Gas:      "0x76c0",
		GasPrice: "0x2540be400",
		Value:    "0x1",
		//Value: "0x33b2e3c9fd0803ce8000000", // 1000000000e+18
		//Value: "0x152d02c7e14af6800000", // 100000e+18
		Data: "",
	}

	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_sendTransaction",
		"params":  []interface{}{params},
		"id":      "1",
	}

	for i := 0; i < 100000; i++ {
		resp, err := utils.RpcRequest(postData)

		if err != nil {
			fmt.Printf("交易发送失败， err=%v\n", err)
		} else {
			err := json.Unmarshal([]byte(resp.Body), &r)
			if err != nil {
				fmt.Printf("交易发送失败， err %v", err)
			}
			fmt.Printf("交易发送成功， hash %v", r.Result)
		}
	}

	return "", err
}
