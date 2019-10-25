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
	var data = "0xf878808504a817c800827530a0494e54334734716d704a48533362627a68754432646151316d6e59334d6f4634880de0b6b3a76400008026a0cbc209a7cc0699ff1dd31d9a5ac32559705f14d53fd84ba0bd828d02959ec71ca055f9b06147e97a90d1aff79b4c20f9f5fb77045c3d63b9b0e002328ceebff477"
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
