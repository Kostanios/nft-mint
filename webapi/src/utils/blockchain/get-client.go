package blockchainutils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

func GetEthClient() *ethclient.Client {
	ChainURL, chainURLExists := os.LookupEnv("CHAIN_API_URL")
	if !chainURLExists {
		log.Fatalf("ChainURL is incorrect")
		return nil
	}

	client, err := ethclient.Dial(ChainURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client")
		return nil
	}

	return client
}
