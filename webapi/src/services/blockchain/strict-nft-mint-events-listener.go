package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kostanios/nft-mint/webapi/src/utils/abi"
	"log"
	"net/http"
	"os"
)

func ListenToMintEvents() {
	ChainURL, chainURLExists := os.LookupEnv("CHAIN_WS_URL")
	if !chainURLExists {
		log.Fatalf("ChainURL is incorrect: %v", http.StatusInternalServerError)
		return
	}

	client, err := ethclient.Dial(ChainURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	StrictMintContractAddress, StrictMintContractAddressExists := os.LookupEnv("STRICT_MINT_NFT_CONTRACT_ADDRESS")
	if !StrictMintContractAddressExists {
		log.Fatalf("MintContractAddress is incorrect")
		return
	}

	contractAddress := common.HexToAddress(StrictMintContractAddress)
	instance, err := abi.NewStrictMintNFT(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}

	eventChannel := make(chan *abi.StrictMintNFTMintedNFT)
	subscription, err := instance.WatchMintedNFT(nil, eventChannel)
	if err != nil {
		log.Fatalf("Failed to subscribe to mint events: %v", err)
	}

	for {
		select {
		case event := <-eventChannel:
			log.Printf("Mint event detected: Owner %s, TokenId %d", event.Owner.Hex(), event.TokenId)
		case err := <-subscription.Err():
			log.Fatalf("Subscription error: %v", err)
		}
	}
}
