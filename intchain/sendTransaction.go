package main

import (
	"fmt"
	"github.com/mikemintang/go-curl"
	"go-test/intchain/config"
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

	//url := config.RemoteConfig.RpcUrl
	url := config.LocalConfig.RpcUrl

	headers := config.ConHeaders

	// 部署合约的交易
	params := TxParams{
		//From: "0x37eeb099c6d751e7229df825d40629612e134f82",
		From:     "0x6388a9d239ad1afcc7ca8d5c4ef2e1caeb588aca",
		To:       "",
		Gas:      "0x4c4b40",
		GasPrice: "",
		Value:    "",
		Data:     "0x60606040523415600e57600080fd5b60d38061001c6000396000f300606060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c6888fa181146043575b600080fd5b3415604d57600080fd5b60566004356068565b60405190815260200160405180910390f35b60007f24abdb5865df5079dcc5ac590ff6f01d5c16edbc5fab4e195d9febd1114503da8260070260405190815260200160405180910390a150600702905600a165627a7a7230582093e2ed97247c1e6c1d67df14ef4f815ae6eaeaedd13aa85f4cce214602a91a310029",
	}

	//sendTransaction(url, headers, params)

	params = TxParams{
		//From: "0x37eeb099c6d751e7229df825d40629612e134f82",
		From:     "INT32UrbTe22LvT3ocrnSryBMqwoedg8",
		To:       "INT3Pkr1zMmk3mnFzihH5F4kNxFavJo4",
		Gas:      "0x76c0",
		GasPrice: "0x2540be400",
		//Value:    "0x1",
		//Value: "0x33b2e3c9fd0803ce8000000",  // 1000000000e+18
		Value: "0x152d02c7e14af6800000", // 100000e+18
		Data:  "",
	}

	sendTransactions(url, headers, params)

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

	for i := 0; i < 1; i++ {
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
