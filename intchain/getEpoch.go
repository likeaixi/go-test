package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

func main() {
	epochNumber, err := getCurrentEpochNumber()
	if err != nil {
		fmt.Printf("获取epoch number失败 err=%v\n", err)
	}
	fmt.Printf("epochNumber=%v\n", epochNumber)

	epoch, err := getEpoch("0x1F")
	if err != nil {
		fmt.Printf("获取epoch 失败 err=%v\n", err)
	}
	fmt.Printf("epoch=%v\n", epoch)
}

func getCurrentEpochNumber() (number interface{}, err error) {
	var r config.EpochNumberRPC
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "tdm_getCurrentEpochNumber",
		"params":  []interface{}{},
		"id":      "1",
	}
	resp, err := utils.RpcRequest(postData)
	if err != nil {
		fmt.Printf("获取 Epoch Number失败, err=%v\n", err)
		return "-1", err
	} else {
		fmt.Printf("resp=%v\n", resp.Body)
		err := json.Unmarshal([]byte(resp.Body), &r)
		if err != nil {
			return "-1", err
		}
		return r.Result, nil
	}
}

func getEpoch(epochNumber interface{}) (epoch config.Epoch, err error) {
	var epochR config.EpochRPC

	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "tdm_getEpoch",
		"params":  []interface{}{epochNumber},
		"id":      "1",
	}
	resp, err := utils.RpcRequest(postData)
	if err != nil {
		fmt.Printf("获取 Epoch 失败, err=%v\n", err)
		return epochR.Result, err
	} else {
		fmt.Printf("resp=%v\n", resp.Body)
		err := json.Unmarshal([]byte(resp.Body), &epochR)
		if err != nil {
			return epochR.Result, err
		}

		return epochR.Result, nil
	}
}
