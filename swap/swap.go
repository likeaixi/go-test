package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/likeaixi/go-test/intchain/utils"
	"io"
)

// UpGetSwapTokenList 获得代币列表
type UpGetSwapTokenList struct {
	//(可选参数)
	Start int32
	Num   int32
}

// DownGetSwapTokenList 下发获得代币列表
type DownGetSwapTokenList struct {
	List []SwapTokenItem
}

// SwapTokenItem 单项列表
type SwapTokenItem struct {
	ID              int64
	TokenName       string
	TokenUnitName   string
	TokenFullName   string
	Icon            string
	ContractAddress string
	Remark          string
}

var key = "X9dsf_)#&334$R("

func main() {
	err := getSwapToken()
	if err != nil {
		fmt.Printf("get swap token err %v\n", err)
	}
}

func getSwapToken() error {
	var data = UpGetSwapTokenList{
		Start: 0,
		Num:   -1,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	fmt.Printf("json data %v\n", jsonData)
	fmt.Printf("string json data %v\n", string(jsonData))

	synData := string(jsonData) + "" + key
	fmt.Printf("syn data %v\n", md5.Sum([]byte(synData)))

	h := md5.New()
	io.WriteString(h, string(jsonData))
	io.WriteString(h, key)
	synData2 := h.Sum(nil)
	fmt.Printf("syn data2 %v\n", synData2)
	fmt.Printf("syn data2 string %v\n", hex.EncodeToString(synData2))

	postData := map[string]interface{}{
		"data": string(jsonData),
		"syn":  hex.EncodeToString(synData2),
	}

	result, err := utils.RpcRequest(postData)
	if err != nil {
		return err
	}

	fmt.Printf("result %v\n", result.Body)

	return nil
}
