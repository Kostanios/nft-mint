package blockchainutils

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"strconv"
)

func GetAuthTransactor() (*bind.TransactOpts, error) {
	PrivateKey, PrivateKeyExists := os.LookupEnv("PRIVATE_KEY")
	if !PrivateKeyExists {
		return nil, errors.New("private key is not exist")
	}

	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		return nil, errors.New("invalid private key")
	}

	ChainId, ChainIdExists := os.LookupEnv("CHAIN_ID")
	if !ChainIdExists {
		return nil, errors.New("chain id is incorrect")
	}

	chainId, err := strconv.ParseInt(ChainId, 10, 64)
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	return auth, nil
}
