package utils

import (
	"errors"
	"fmt"
	"github.com/mikemintang/go-curl"
	"go-test/intchain/config"
)

func RpcRequest(postData map[string]interface{}) (result *curl.Response, err error) {
	url := config.LocalConfig.RpcUrl

	headers := config.ConHeaders

	req := curl.NewRequest()

	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetPostData(postData).
		Post()
	var r = &curl.Response{}
	if err != nil {
		return r, err
	} else {
		if resp.IsOk() {
			fmt.Printf("Rpc request sucess resp.body=%v\n", resp.Body)
			return resp, nil
		} else {
			return r, errors.New("RPC请求失败")
		}
	}
}
