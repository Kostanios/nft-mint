package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kostanios/nft-mint/webapi/src/utils/abi"
	blockchainutils "github.com/kostanios/nft-mint/webapi/src/utils/blockchain"
	"os"

	"math/big"
	"net/http"
)

func ListNFTsHandler(w http.ResponseWriter, r *http.Request) {
	client := blockchainutils.GetEthClient()
	defer client.Close()

	MintContractAddress, MintContractAddressExists := os.LookupEnv("MINT_NFT_CONTRACT_ADDRESS")
	if !MintContractAddressExists {
		http.Error(w, "MintContractAddress is incorrect", http.StatusInternalServerError)
		return
	}

	StrictMintContractAddress, StrictMintContractAddressExists := os.LookupEnv("STRICT_MINT_NFT_CONTRACT_ADDRESS")
	if !StrictMintContractAddressExists {
		http.Error(w, "MintContractAddress is incorrect", http.StatusInternalServerError)
		return
	}

	data, err := ListNFTs(client, MintContractAddress, StrictMintContractAddress)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list NFTs: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ListNFTs(client *ethclient.Client, contract1Address, contract2Address string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	contract1, err := abi.NewMintNFT(common.HexToAddress(contract1Address), client)
	if err != nil {
		return nil, fmt.Errorf("failed to load first contract: %v", err)
	}
	counter1, err := contract1.TokenIdCounter(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch counter from first contract: %v", err)
	}
	result["mintContractCounter"] = counter1.Int64()

	contract2, err := abi.NewStrictMintNFT(common.HexToAddress(contract2Address), client)
	if err != nil {
		return nil, fmt.Errorf("failed to load second contract: %v", err)
	}
	counter2, err := contract2.TokenIdCounter(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch counter from second contract: %v", err)
	}

	var tokenNames []string
	for i := int64(0); i < counter2.Int64(); i++ {
		tokenID := big.NewInt(i)
		name, err := contract2.CustomNames(nil, tokenID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch token name for token ID %d: %v", i, err)
		}
		tokenNames = append(tokenNames, name)
	}

	result["strictMintContractCounter"] = counter2.Int64()
	result["strictMintContractTokenNames"] = tokenNames

	return result, nil
}
