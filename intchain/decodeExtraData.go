package main

import (
	"encoding/json"
	"fmt"
	"go-test/intchain/config"
	"go-test/intchain/utils"
)

type extraDataParams struct {
	ExtraData string `json:"extra_data"`
}

//curl -X POST --data '{"jsonrpc":"2.0","method":"tdm_decodeExtraData","params":["0x0108696e74636861696e0000000000002c0615cd7d205898b7400000000000000000000201149d2b5214e24104f648819a2b851970e02e0f73c50114dae66c6bca991012f6f62d4db42c79c60e3c2e44010114980aaa734720bdc9d93e3d4030a135aec4910bac0000000000000001011415a43260d301c016d636ec117eac56578f1157990000000000002c06000140439243a34a9ef90323f7db46d7800b7049fa4a50599956054d716de1cf37f7dc686c8a5c6510a17609cda05dd53737cc8635f15da4f7f48ccc1c288db0960e730100000000000000010101000000000000000100"],"id":1}' -H 'content-type: application/json;' http://127.0.0.1:7000/intchain
func main() {
	extraData, err := decodeExtraData()
	if err != nil {
		fmt.Printf("decode extra data err %v\n", err)
	} else {
		fmt.Printf("extra data %v\n", extraData)
	}

}

// 一个 validator
// 0x0108696e74636861696e000000000000000215d0c6db99932d4000000000000000000000011439a5bda9461256223334630b3de0ff2f2845bffd011492cc85e460cf0d691ed6857d01b9d411b15f9017010114354f5285dc9897de5432887232322d245426379700000000000000010114199f547bb335dacdf297d38f711981616666e4730000000000000002000140570a7d3807fb99a1c0a2208c5376c55a8c56b605791efe54075036c14e4132ae35ac12aeca39572efff2efa1de0ee036c672d24dfaf765bcadf6d09098ff677d0100000000000000010101000000000000000100
// 两个validator
//0x0108696e74636861696e00000000000364cf15d1c283039ae1800000000000000000001f0114ded52af4ce2d5caabdcf1b70aaee8bf7414b939701147ce0af7b98475a4641b31a47b563271df9cf69b1010114b1caf6a620176acfe86de579186905a119d2830c00000000000000010114fd734faf45b54b8c60e68bb29a0882ab72f2a06100000000000364cf000140484c1d0412e5af6d12596bd326c3f67af649530d0bc4202419fb768f2896c18e6edc351f7835ec260127eecca42fda55280ace2d224408e0724a78d2adb73b280100000000000000020101000000000000000300
// 三个 validator
//0x0108696e74636861696e000000000003713d15d1c7c6c2488b8000000000000000000020011415d14546048c69022a12e0c8605338e68a841fb30114ec299e752fd4b41f0ea3d5a4e2b350e6f3a1471b010114ca95018425ce58d2e041eb5d9eda115f181214f800000000000000010114f4fd5d4e44794f67fa797f7824dd706934970e3c000000000003713d010201405678a97feace6079d60753d636ad7358349fcd2b222e1ffff393359f2c272e5d132f435b18967010e881db0dc52055c8c50ff0ac660195c494aa85a7015c7de30100000000000000030101000000000000000300
func decodeExtraData() (extraData interface{}, err error) {
	var r config.ExtraDataRPC
	params := extraDataParams{
		ExtraData: "0x0108696e74636861696e0000000000001c2a15d2b3d1d7283e40000000000000000000010114c78b2314da16439219a06902fce632d18c83574c01148c03abd5d3940de08b20f810847ef7c8eacfed81010114d7eec6fe037e4fbc26c0974f2659a6bf525a32dd00000000000000010114f33e99d40046d797c19b0bda6163f5ad718df56c0000000000001c2a010b01400c299aa0cac39ea4bff9f729996886c9d71cbfcf2f34eccc0df1ac3be3fd752f09f3dde4fde9e64252dd5684fe2ec5b2eecd4217106cd02f235cb0a9bd61a0f70100000000000000060101000000000000003c00",
	}

	postData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "tdm_decodeExtraData",
		"params":  []interface{}{params.ExtraData},
		"id":      "1",
	}

	resp, err := utils.RpcRequest(postData)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(resp.Body), &r)
	if err != nil {
		return "", err
	} else {
		return r.Result, nil
	}
}
