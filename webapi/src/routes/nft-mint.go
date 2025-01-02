package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kostanios/nft-mint/webapi/src/utils/abi"
	blockchainutils "github.com/kostanios/nft-mint/webapi/src/utils/blockchain"
	"net/http"
	"os"
)

func MintNFT(client *ethclient.Client, contractAddress string) (string, error) {
	auth, err := blockchainutils.GetAuthTransactor()

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
	client := blockchainutils.GetEthClient()
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
