package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

func main() {
	hash, err := sendRawTransaction()
	if err != nil {
		fmt.Printf("sendRawTransaction err %v", err)
	}

	fmt.Printf("hash %v\n", hash)
}

func sendRawTransaction() (hash string, err error) {
	var r config.RawTransactionRPC
	var data = "f87a808504a817c800825208a0494e5433506b72317a4d6d6b336d6e467a6968483546346b4e784661764a6f34880de0b6b3a764000082307826a09f500b82ca945b5cb07e58d8ce7a1da018a48c883e3ba57f173e6dcbaa37e5f5a00e81ca9d05cbeb16cbe059d0dfcc6df62d02417fde926779f85b8c8b6264d75a"
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_sendRawTransaction",
		"params":  []interface{}{data},
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
