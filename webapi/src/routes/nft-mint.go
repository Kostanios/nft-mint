package nftmint

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kostanios/nft-mint/webapi/src/utils/abi"
	"math/big"
	"net/http"
	"os"
	"strconv"
)

func MintNFT(client *ethclient.Client, contractAddress string) (string, error) {
	PrivateKey, PrivateKeyExists := os.LookupEnv("PRIVATE_KEY")
	if !PrivateKeyExists {
		return "", errors.New("private key is not exist")
	}

	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		return "", errors.New("invalid private key")
	}

	ChainId, ChainIdExists := os.LookupEnv("CHAIN_ID")
	if !ChainIdExists {
		return "", errors.New("chain id is incorrect")
	}

	chainId, err := strconv.ParseInt(ChainId, 10, 64)
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %v", err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	instance, err := abi.NewMintNFT(contractAddr, client)
	if err != nil {
		return "", fmt.Errorf("failed to load contract: %v", err)
	}

	tx, err := instance.Mint(auth)
	if err != nil {
		return "", fmt.Errorf("mint transaction failed: %v", err)
	}

	return tx.Hash().Hex(), nil
}

func MintNFTHandler(w http.ResponseWriter, r *http.Request) {
	ChainURL, chainURLExists := os.LookupEnv("CHAIN_API_URL")
	if !chainURLExists {
		http.Error(w, "ChainURL is incorrect", http.StatusInternalServerError)
		return
	}

	client, err := ethclient.Dial(ChainURL)
	if err != nil {
		http.Error(w, "Failed to connect to Ethereum client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	MintContractAddress, MintContractAddressExists := os.LookupEnv("MINT_NFT_CONTRACT_ADDRESS")
	if !MintContractAddressExists {
		http.Error(w, "MintContractAddress is incorrect", http.StatusInternalServerError)
		return
	}

	contractAddress := MintContractAddress

	tx, err := MintNFT(client, contractAddress)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to mint NFT: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"txHash": tx})
}
