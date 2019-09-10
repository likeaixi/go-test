package main

type Config struct {
	rpcUrl      string
	contentType string
}

var localConfig = Config{
	rpcUrl:      "http://127.0.0.1:8545/intchain",
	contentType: "application/json",
}

var remoteConfig = Config{
	rpcUrl:      "http://45.125.32.219:8545/intchain",
	contentType: "application/json",
}
