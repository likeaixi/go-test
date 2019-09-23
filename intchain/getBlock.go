package main

import (
	"encoding/json"
	"fmt"
	"github.com/mikemintang/go-curl"
	"math/big"
)

type Result struct {
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasused"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	MainChainNumber  string        `json:"mainchainNumber"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Sha3Uncles       string        `json:"sha2Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	Transactions     []interface{} `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
}

type Body struct {
	JsonRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  Result `json:"result"`
}

var blockNumber *big.Int
var blockTime *big.Int

func main() {
	//curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}' -H 'content-type: application/json;' http://127.0.0.1:7000/intchain

	getBlock("latest", -1)
	fmt.Printf("blockNumber=%v\n", blockNumber)

	searchNumber := big.NewInt(0)
	shift := big.NewInt(2000)
	searchNumber.Sub(blockNumber, shift)
	fmt.Printf("searchNumber=%v\n", searchNumber)

	for i := 0; big.NewInt(int64(i)).Cmp(shift) < 0; i++ {
		getBlock("0x"+(big.NewInt(0).Add(searchNumber, big.NewInt(int64(i)))).Text(16), i)
	}

}

func getBlock(n interface{}, i int) {
	url := "http://127.0.0.1:8545/intchain"
	//url := "http://127.0.0.1:7000/intchain"

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params":  []interface{}{n, true},
		"id":      "1",
	}

	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	var r Body
	if err != nil {
		fmt.Printf(" 获取区块失败，err=%v\n", err)
	} else {
		if resp.IsOk() {
			//fmt.Printf("获取区块成功，body=%v\n", resp.Body)
			if err := json.Unmarshal([]byte(resp.Body), &r); err == nil {
				//fmt.Printf("Result %v\n", r)
				result := r.Result
				if result.Number != "" {
					if i > 0 {
						bNumber := hexToBigInt(result.Number)
						txNumber := len(result.Transactions)
						t := hexToBigInt(result.Timestamp)
						cTime := big.NewInt(0)

						if txNumber > 500 {
							fmt.Printf("index=%v, blockNumber=%v, blockHash=%v, transactionCount=%v, costTime=%v, miner=%v\n\n", i, bNumber, result.Hash, txNumber, cTime.Sub(t, blockTime), result.Miner)
						}

						blockTime = t
					} else if i == 0 {
						blockTime = hexToBigInt(result.Timestamp)
					} else {
						blockNumber = hexToBigInt(r.Result.Number)
					}
				}

			} else {
				fmt.Printf("解析数据失败， err=%v\n", err)
			}
		} else {
			fmt.Printf("获取区块失败，raw=%v\n", resp.Raw)
		}
	}
}

func hexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)
	return n
}
