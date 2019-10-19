package config

type Config struct {
	RpcUrl  string
	Headers map[string]string
}

type Headers map[string]string

var LocalConfig = Config{
	//RpcUrl:     "http://127.0.0.1:8545/intchain",
	//RpcUrl:     "http://127.0.0.1:7000/intchain",
	RpcUrl: "http://127.0.0.1:6968/intchain",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

var RemoteConfig = Config{
	RpcUrl: "http://129.226.134.100:7000/intchain",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

var ConHeaders = Headers{
	"Content-Type": "application/json",
}

// RPC 返回数据的结构体
type RPCBody struct {
	JsonRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
}

// Epoch RPC 数据
type EpochNumberRPC struct {
	JsonRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
}

type EpochRPC struct {
	JsonRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  Epoch  `json:"result"`
}

type Epoch struct {
	Number           string      `json:"number"`
	RewardPerBlock   string      `json:"reward_per_block"`
	StartBlock       string      `json:"start_block"`
	EndBlock         string      `json:"end_block"`
	StartTime        string      `json:"start_time"`
	EndTime          string      `json:"end_time"`
	VoteStartBlock   string      `json:"vote_start_block"`
	VoteEndBlock     string      `json:"reward_end_block"`
	RevealStartBlock string      `json:"reveal_start_block"`
	RevealEndBlock   string      `json:"reveal_end_block"`
	Validators       []Validator `json:"validators"`
}

type Validator struct {
	Address     string `json:"address"`
	PublicKey   string `json:"public_key"`
	VotingPower string `json:"voting_power"`
	RemainEpoch string `json:"remain_epoch"`
}

// Delegation RPC  数据
type DelegateRPC struct {
	JsonRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
}

type CandidateRPC struct {
	JsonRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
}

// Chain RPC 数据
