package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kostanios/nft-mint/webapi/src/utils/abi"
	blockchainutils "github.com/kostanios/nft-mint/webapi/src/utils/blockchain"
	"math/big"
	"net/http"
	"os"
)

type MintRequest struct {
	Quantity int      `json:"quantity"`
	Names    []string `json:"names"`
}

func CalculatePrice(instance *abi.StrictMintNFT, quantity int, names []string) (*big.Int, error) {
	first15Price, err := instance.FirstTierPrice(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch first 15 price: %v", err)
	}

	remainingPrice, err := instance.SecondTierPrice(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch remaining price: %v", err)
	}

	customNamePrice, err := instance.CustomNamePrice(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch custom name price: %v", err)
	}

	counterBigInt, err := instance.TokenIdCounter(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch custom name price: %v", err)
	}

	totalPrice := big.NewInt(0)
	counter := counterBigInt.Int64()

	for i := 0; i < quantity; i++ {
		if counter < 15 {
			totalPrice.Add(totalPrice, first15Price)
			counter++
		} else {
			totalPrice.Add(totalPrice, remainingPrice)
		}
	}

	if len(names) > 0 {
		customNameTotal := new(big.Int).Mul(customNamePrice, big.NewInt(int64(len(names))))
		totalPrice.Add(totalPrice, customNameTotal)
	}

	return totalPrice, nil
}

func StrictMintNFT(client *ethclient.Client, contractAddress string, quantity int, names []string) (string, error) {
	if quantity <= 0 || quantity > 5 {
		return "", errors.New("quantity must be between 1 and 5")
	}

	auth, err := blockchainutils.GetAuthTransactor()

	contractAddr := common.HexToAddress(contractAddress)
	instance, err := abi.NewStrictMintNFT(contractAddr, client)
	if err != nil {
		return "", fmt.Errorf("failed to load contract: %v", err)
	}

	totalPrice, _ := CalculatePrice(instance, quantity, names)

	auth.Value = totalPrice

	bigQuantity := big.NewInt(int64(quantity))
	tx, err := instance.Mint(auth, bigQuantity, names)
	if err != nil {
		return "", fmt.Errorf("mint transaction failed: %v", err)
	}

	return tx.Hash().Hex(), nil
}

func StrictMintNFTHandler(w http.ResponseWriter, r *http.Request) {
	var request MintRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	client := blockchainutils.GetEthClient()

	defer client.Close()

	StrictMintContractAddress, StrictMintContractAddressExists := os.LookupEnv("STRICT_MINT_NFT_CONTRACT_ADDRESS")
	if !StrictMintContractAddressExists {
		http.Error(w, "MintContractAddress is incorrect", http.StatusInternalServerError)
		return
	}

	contractAddress := StrictMintContractAddress
	tx, err := StrictMintNFT(client, contractAddress, request.Quantity, request.Names)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to mint NFTs: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"txHash": tx})
}
