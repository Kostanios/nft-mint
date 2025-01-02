// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StrictMintNFTMetaData contains all meta data concerning the StrictMintNFT contract.
var StrictMintNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mintNFTContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"customName\",\"type\":\"string\"}],\"name\":\"MintedNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"customNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"customNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstTierPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstTierRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintNFTContract\",\"outputs\":[{\"internalType\":\"contractMintNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondTierPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setCustomName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_firstTierPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondTierPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_customNamePrice\",\"type\":\"uint256\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StrictMintNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use StrictMintNFTMetaData.ABI instead.
var StrictMintNFTABI = StrictMintNFTMetaData.ABI

// StrictMintNFT is an auto generated Go binding around an Ethereum contract.
type StrictMintNFT struct {
	StrictMintNFTCaller     // Read-only binding to the contract
	StrictMintNFTTransactor // Write-only binding to the contract
	StrictMintNFTFilterer   // Log filterer for contract events
}

// StrictMintNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrictMintNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrictMintNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrictMintNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrictMintNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrictMintNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrictMintNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrictMintNFTSession struct {
	Contract     *StrictMintNFT    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrictMintNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrictMintNFTCallerSession struct {
	Contract *StrictMintNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StrictMintNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrictMintNFTTransactorSession struct {
	Contract     *StrictMintNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StrictMintNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrictMintNFTRaw struct {
	Contract *StrictMintNFT // Generic contract binding to access the raw methods on
}

// StrictMintNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrictMintNFTCallerRaw struct {
	Contract *StrictMintNFTCaller // Generic read-only contract binding to access the raw methods on
}

// StrictMintNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrictMintNFTTransactorRaw struct {
	Contract *StrictMintNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrictMintNFT creates a new instance of StrictMintNFT, bound to a specific deployed contract.
func NewStrictMintNFT(address common.Address, backend bind.ContractBackend) (*StrictMintNFT, error) {
	contract, err := bindStrictMintNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFT{StrictMintNFTCaller: StrictMintNFTCaller{contract: contract}, StrictMintNFTTransactor: StrictMintNFTTransactor{contract: contract}, StrictMintNFTFilterer: StrictMintNFTFilterer{contract: contract}}, nil
}

// NewStrictMintNFTCaller creates a new read-only instance of StrictMintNFT, bound to a specific deployed contract.
func NewStrictMintNFTCaller(address common.Address, caller bind.ContractCaller) (*StrictMintNFTCaller, error) {
	contract, err := bindStrictMintNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFTCaller{contract: contract}, nil
}

// NewStrictMintNFTTransactor creates a new write-only instance of StrictMintNFT, bound to a specific deployed contract.
func NewStrictMintNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*StrictMintNFTTransactor, error) {
	contract, err := bindStrictMintNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFTTransactor{contract: contract}, nil
}

// NewStrictMintNFTFilterer creates a new log filterer instance of StrictMintNFT, bound to a specific deployed contract.
func NewStrictMintNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*StrictMintNFTFilterer, error) {
	contract, err := bindStrictMintNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFTFilterer{contract: contract}, nil
}

// bindStrictMintNFT binds a generic wrapper to an already deployed contract.
func bindStrictMintNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrictMintNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrictMintNFT *StrictMintNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrictMintNFT.Contract.StrictMintNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrictMintNFT *StrictMintNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.StrictMintNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrictMintNFT *StrictMintNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.StrictMintNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrictMintNFT *StrictMintNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrictMintNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrictMintNFT *StrictMintNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrictMintNFT *StrictMintNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTSession) MAXSUPPLY() (*big.Int, error) {
	return _StrictMintNFT.Contract.MAXSUPPLY(&_StrictMintNFT.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _StrictMintNFT.Contract.MAXSUPPLY(&_StrictMintNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StrictMintNFT *StrictMintNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StrictMintNFT.Contract.BalanceOf(&_StrictMintNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StrictMintNFT.Contract.BalanceOf(&_StrictMintNFT.CallOpts, owner)
}

// CustomNamePrice is a free data retrieval call binding the contract method 0xefa2a5a1.
//
// Solidity: function customNamePrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCaller) CustomNamePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "customNamePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CustomNamePrice is a free data retrieval call binding the contract method 0xefa2a5a1.
//
// Solidity: function customNamePrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTSession) CustomNamePrice() (*big.Int, error) {
	return _StrictMintNFT.Contract.CustomNamePrice(&_StrictMintNFT.CallOpts)
}

// CustomNamePrice is a free data retrieval call binding the contract method 0xefa2a5a1.
//
// Solidity: function customNamePrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCallerSession) CustomNamePrice() (*big.Int, error) {
	return _StrictMintNFT.Contract.CustomNamePrice(&_StrictMintNFT.CallOpts)
}

// CustomNames is a free data retrieval call binding the contract method 0xc440e011.
//
// Solidity: function customNames(uint256 ) view returns(string)
func (_StrictMintNFT *StrictMintNFTCaller) CustomNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "customNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CustomNames is a free data retrieval call binding the contract method 0xc440e011.
//
// Solidity: function customNames(uint256 ) view returns(string)
func (_StrictMintNFT *StrictMintNFTSession) CustomNames(arg0 *big.Int) (string, error) {
	return _StrictMintNFT.Contract.CustomNames(&_StrictMintNFT.CallOpts, arg0)
}

// CustomNames is a free data retrieval call binding the contract method 0xc440e011.
//
// Solidity: function customNames(uint256 ) view returns(string)
func (_StrictMintNFT *StrictMintNFTCallerSession) CustomNames(arg0 *big.Int) (string, error) {
	return _StrictMintNFT.Contract.CustomNames(&_StrictMintNFT.CallOpts, arg0)
}

// FirstTierPrice is a free data retrieval call binding the contract method 0x5fa17324.
//
// Solidity: function firstTierPrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCaller) FirstTierPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "firstTierPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstTierPrice is a free data retrieval call binding the contract method 0x5fa17324.
//
// Solidity: function firstTierPrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTSession) FirstTierPrice() (*big.Int, error) {
	return _StrictMintNFT.Contract.FirstTierPrice(&_StrictMintNFT.CallOpts)
}

// FirstTierPrice is a free data retrieval call binding the contract method 0x5fa17324.
//
// Solidity: function firstTierPrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCallerSession) FirstTierPrice() (*big.Int, error) {
	return _StrictMintNFT.Contract.FirstTierPrice(&_StrictMintNFT.CallOpts)
}

// FirstTierRemaining is a free data retrieval call binding the contract method 0xa11d9b0c.
//
// Solidity: function firstTierRemaining() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCaller) FirstTierRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "firstTierRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstTierRemaining is a free data retrieval call binding the contract method 0xa11d9b0c.
//
// Solidity: function firstTierRemaining() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTSession) FirstTierRemaining() (*big.Int, error) {
	return _StrictMintNFT.Contract.FirstTierRemaining(&_StrictMintNFT.CallOpts)
}

// FirstTierRemaining is a free data retrieval call binding the contract method 0xa11d9b0c.
//
// Solidity: function firstTierRemaining() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCallerSession) FirstTierRemaining() (*big.Int, error) {
	return _StrictMintNFT.Contract.FirstTierRemaining(&_StrictMintNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StrictMintNFT *StrictMintNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StrictMintNFT *StrictMintNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _StrictMintNFT.Contract.GetApproved(&_StrictMintNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StrictMintNFT *StrictMintNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _StrictMintNFT.Contract.GetApproved(&_StrictMintNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StrictMintNFT *StrictMintNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StrictMintNFT *StrictMintNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _StrictMintNFT.Contract.IsApprovedForAll(&_StrictMintNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StrictMintNFT *StrictMintNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _StrictMintNFT.Contract.IsApprovedForAll(&_StrictMintNFT.CallOpts, owner, operator)
}

// MintCount is a free data retrieval call binding the contract method 0xed9ec888.
//
// Solidity: function mintCount(address ) view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCaller) MintCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "mintCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCount is a free data retrieval call binding the contract method 0xed9ec888.
//
// Solidity: function mintCount(address ) view returns(uint256)
func (_StrictMintNFT *StrictMintNFTSession) MintCount(arg0 common.Address) (*big.Int, error) {
	return _StrictMintNFT.Contract.MintCount(&_StrictMintNFT.CallOpts, arg0)
}

// MintCount is a free data retrieval call binding the contract method 0xed9ec888.
//
// Solidity: function mintCount(address ) view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCallerSession) MintCount(arg0 common.Address) (*big.Int, error) {
	return _StrictMintNFT.Contract.MintCount(&_StrictMintNFT.CallOpts, arg0)
}

// MintNFTContract is a free data retrieval call binding the contract method 0xe90f04c7.
//
// Solidity: function mintNFTContract() view returns(address)
func (_StrictMintNFT *StrictMintNFTCaller) MintNFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "mintNFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintNFTContract is a free data retrieval call binding the contract method 0xe90f04c7.
//
// Solidity: function mintNFTContract() view returns(address)
func (_StrictMintNFT *StrictMintNFTSession) MintNFTContract() (common.Address, error) {
	return _StrictMintNFT.Contract.MintNFTContract(&_StrictMintNFT.CallOpts)
}

// MintNFTContract is a free data retrieval call binding the contract method 0xe90f04c7.
//
// Solidity: function mintNFTContract() view returns(address)
func (_StrictMintNFT *StrictMintNFTCallerSession) MintNFTContract() (common.Address, error) {
	return _StrictMintNFT.Contract.MintNFTContract(&_StrictMintNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrictMintNFT *StrictMintNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrictMintNFT *StrictMintNFTSession) Name() (string, error) {
	return _StrictMintNFT.Contract.Name(&_StrictMintNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrictMintNFT *StrictMintNFTCallerSession) Name() (string, error) {
	return _StrictMintNFT.Contract.Name(&_StrictMintNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrictMintNFT *StrictMintNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrictMintNFT *StrictMintNFTSession) Owner() (common.Address, error) {
	return _StrictMintNFT.Contract.Owner(&_StrictMintNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrictMintNFT *StrictMintNFTCallerSession) Owner() (common.Address, error) {
	return _StrictMintNFT.Contract.Owner(&_StrictMintNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StrictMintNFT *StrictMintNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StrictMintNFT *StrictMintNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StrictMintNFT.Contract.OwnerOf(&_StrictMintNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StrictMintNFT *StrictMintNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StrictMintNFT.Contract.OwnerOf(&_StrictMintNFT.CallOpts, tokenId)
}

// SecondTierPrice is a free data retrieval call binding the contract method 0x91516cce.
//
// Solidity: function secondTierPrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCaller) SecondTierPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "secondTierPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondTierPrice is a free data retrieval call binding the contract method 0x91516cce.
//
// Solidity: function secondTierPrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTSession) SecondTierPrice() (*big.Int, error) {
	return _StrictMintNFT.Contract.SecondTierPrice(&_StrictMintNFT.CallOpts)
}

// SecondTierPrice is a free data retrieval call binding the contract method 0x91516cce.
//
// Solidity: function secondTierPrice() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCallerSession) SecondTierPrice() (*big.Int, error) {
	return _StrictMintNFT.Contract.SecondTierPrice(&_StrictMintNFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StrictMintNFT *StrictMintNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StrictMintNFT *StrictMintNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StrictMintNFT.Contract.SupportsInterface(&_StrictMintNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StrictMintNFT *StrictMintNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StrictMintNFT.Contract.SupportsInterface(&_StrictMintNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrictMintNFT *StrictMintNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrictMintNFT *StrictMintNFTSession) Symbol() (string, error) {
	return _StrictMintNFT.Contract.Symbol(&_StrictMintNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrictMintNFT *StrictMintNFTCallerSession) Symbol() (string, error) {
	return _StrictMintNFT.Contract.Symbol(&_StrictMintNFT.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCaller) TokenIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "tokenIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTSession) TokenIdCounter() (*big.Int, error) {
	return _StrictMintNFT.Contract.TokenIdCounter(&_StrictMintNFT.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_StrictMintNFT *StrictMintNFTCallerSession) TokenIdCounter() (*big.Int, error) {
	return _StrictMintNFT.Contract.TokenIdCounter(&_StrictMintNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StrictMintNFT *StrictMintNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _StrictMintNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StrictMintNFT *StrictMintNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _StrictMintNFT.Contract.TokenURI(&_StrictMintNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StrictMintNFT *StrictMintNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _StrictMintNFT.Contract.TokenURI(&_StrictMintNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.Approve(&_StrictMintNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.Approve(&_StrictMintNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x8a1bbf86.
//
// Solidity: function mint(uint256 quantity, string[] names) payable returns()
func (_StrictMintNFT *StrictMintNFTTransactor) Mint(opts *bind.TransactOpts, quantity *big.Int, names []string) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "mint", quantity, names)
}

// Mint is a paid mutator transaction binding the contract method 0x8a1bbf86.
//
// Solidity: function mint(uint256 quantity, string[] names) payable returns()
func (_StrictMintNFT *StrictMintNFTSession) Mint(quantity *big.Int, names []string) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.Mint(&_StrictMintNFT.TransactOpts, quantity, names)
}

// Mint is a paid mutator transaction binding the contract method 0x8a1bbf86.
//
// Solidity: function mint(uint256 quantity, string[] names) payable returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) Mint(quantity *big.Int, names []string) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.Mint(&_StrictMintNFT.TransactOpts, quantity, names)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrictMintNFT *StrictMintNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrictMintNFT *StrictMintNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrictMintNFT.Contract.RenounceOwnership(&_StrictMintNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrictMintNFT.Contract.RenounceOwnership(&_StrictMintNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SafeTransferFrom(&_StrictMintNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SafeTransferFrom(&_StrictMintNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StrictMintNFT *StrictMintNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StrictMintNFT *StrictMintNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SafeTransferFrom0(&_StrictMintNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SafeTransferFrom0(&_StrictMintNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StrictMintNFT *StrictMintNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StrictMintNFT *StrictMintNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SetApprovalForAll(&_StrictMintNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SetApprovalForAll(&_StrictMintNFT.TransactOpts, operator, approved)
}

// SetCustomName is a paid mutator transaction binding the contract method 0x43d2cbcb.
//
// Solidity: function setCustomName(uint256 tokenId, string name) payable returns()
func (_StrictMintNFT *StrictMintNFTTransactor) SetCustomName(opts *bind.TransactOpts, tokenId *big.Int, name string) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "setCustomName", tokenId, name)
}

// SetCustomName is a paid mutator transaction binding the contract method 0x43d2cbcb.
//
// Solidity: function setCustomName(uint256 tokenId, string name) payable returns()
func (_StrictMintNFT *StrictMintNFTSession) SetCustomName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SetCustomName(&_StrictMintNFT.TransactOpts, tokenId, name)
}

// SetCustomName is a paid mutator transaction binding the contract method 0x43d2cbcb.
//
// Solidity: function setCustomName(uint256 tokenId, string name) payable returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) SetCustomName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SetCustomName(&_StrictMintNFT.TransactOpts, tokenId, name)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _firstTierPrice, uint256 _secondTierPrice, uint256 _customNamePrice) returns()
func (_StrictMintNFT *StrictMintNFTTransactor) SetPrices(opts *bind.TransactOpts, _firstTierPrice *big.Int, _secondTierPrice *big.Int, _customNamePrice *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "setPrices", _firstTierPrice, _secondTierPrice, _customNamePrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _firstTierPrice, uint256 _secondTierPrice, uint256 _customNamePrice) returns()
func (_StrictMintNFT *StrictMintNFTSession) SetPrices(_firstTierPrice *big.Int, _secondTierPrice *big.Int, _customNamePrice *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SetPrices(&_StrictMintNFT.TransactOpts, _firstTierPrice, _secondTierPrice, _customNamePrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _firstTierPrice, uint256 _secondTierPrice, uint256 _customNamePrice) returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) SetPrices(_firstTierPrice *big.Int, _secondTierPrice *big.Int, _customNamePrice *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.SetPrices(&_StrictMintNFT.TransactOpts, _firstTierPrice, _secondTierPrice, _customNamePrice)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.TransferFrom(&_StrictMintNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.TransferFrom(&_StrictMintNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrictMintNFT *StrictMintNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StrictMintNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrictMintNFT *StrictMintNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.TransferOwnership(&_StrictMintNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrictMintNFT *StrictMintNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrictMintNFT.Contract.TransferOwnership(&_StrictMintNFT.TransactOpts, newOwner)
}

// StrictMintNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StrictMintNFT contract.
type StrictMintNFTApprovalIterator struct {
	Event *StrictMintNFTApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrictMintNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrictMintNFTApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrictMintNFTApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrictMintNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrictMintNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrictMintNFTApproval represents a Approval event raised by the StrictMintNFT contract.
type StrictMintNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StrictMintNFT *StrictMintNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StrictMintNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StrictMintNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFTApprovalIterator{contract: _StrictMintNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StrictMintNFT *StrictMintNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StrictMintNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StrictMintNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrictMintNFTApproval)
				if err := _StrictMintNFT.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StrictMintNFT *StrictMintNFTFilterer) ParseApproval(log types.Log) (*StrictMintNFTApproval, error) {
	event := new(StrictMintNFTApproval)
	if err := _StrictMintNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrictMintNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StrictMintNFT contract.
type StrictMintNFTApprovalForAllIterator struct {
	Event *StrictMintNFTApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrictMintNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrictMintNFTApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrictMintNFTApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrictMintNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrictMintNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrictMintNFTApprovalForAll represents a ApprovalForAll event raised by the StrictMintNFT contract.
type StrictMintNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StrictMintNFT *StrictMintNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StrictMintNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StrictMintNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFTApprovalForAllIterator{contract: _StrictMintNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StrictMintNFT *StrictMintNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StrictMintNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StrictMintNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrictMintNFTApprovalForAll)
				if err := _StrictMintNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StrictMintNFT *StrictMintNFTFilterer) ParseApprovalForAll(log types.Log) (*StrictMintNFTApprovalForAll, error) {
	event := new(StrictMintNFTApprovalForAll)
	if err := _StrictMintNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrictMintNFTMintedNFTIterator is returned from FilterMintedNFT and is used to iterate over the raw logs and unpacked data for MintedNFT events raised by the StrictMintNFT contract.
type StrictMintNFTMintedNFTIterator struct {
	Event *StrictMintNFTMintedNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrictMintNFTMintedNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrictMintNFTMintedNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrictMintNFTMintedNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrictMintNFTMintedNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrictMintNFTMintedNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrictMintNFTMintedNFT represents a MintedNFT event raised by the StrictMintNFT contract.
type StrictMintNFTMintedNFT struct {
	TokenId    *big.Int
	Owner      common.Address
	CustomName string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintedNFT is a free log retrieval operation binding the contract event 0x1c5600493310ae85dfff03e7b25db65b89bb0886c463a9aa05825c60facd78c9.
//
// Solidity: event MintedNFT(uint256 indexed tokenId, address indexed owner, string customName)
func (_StrictMintNFT *StrictMintNFTFilterer) FilterMintedNFT(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*StrictMintNFTMintedNFTIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StrictMintNFT.contract.FilterLogs(opts, "MintedNFT", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFTMintedNFTIterator{contract: _StrictMintNFT.contract, event: "MintedNFT", logs: logs, sub: sub}, nil
}

// WatchMintedNFT is a free log subscription operation binding the contract event 0x1c5600493310ae85dfff03e7b25db65b89bb0886c463a9aa05825c60facd78c9.
//
// Solidity: event MintedNFT(uint256 indexed tokenId, address indexed owner, string customName)
func (_StrictMintNFT *StrictMintNFTFilterer) WatchMintedNFT(opts *bind.WatchOpts, sink chan<- *StrictMintNFTMintedNFT) (event.Subscription, error) {

	//var tokenIdRule []interface{}
	//for _, tokenIdItem := range tokenId {
	//	tokenIdRule = append(tokenIdRule, tokenIdItem)
	//}
	//var ownerRule []interface{}
	//for _, ownerItem := range owner {
	//	ownerRule = append(ownerRule, ownerItem)
	//}

	logs, sub, err := _StrictMintNFT.contract.WatchLogs(opts, "MintedNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrictMintNFTMintedNFT)
				if err := _StrictMintNFT.contract.UnpackLog(event, "MintedNFT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintedNFT is a log parse operation binding the contract event 0x1c5600493310ae85dfff03e7b25db65b89bb0886c463a9aa05825c60facd78c9.
//
// Solidity: event MintedNFT(uint256 indexed tokenId, address indexed owner, string customName)
func (_StrictMintNFT *StrictMintNFTFilterer) ParseMintedNFT(log types.Log) (*StrictMintNFTMintedNFT, error) {
	event := new(StrictMintNFTMintedNFT)
	if err := _StrictMintNFT.contract.UnpackLog(event, "MintedNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrictMintNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StrictMintNFT contract.
type StrictMintNFTOwnershipTransferredIterator struct {
	Event *StrictMintNFTOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrictMintNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrictMintNFTOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrictMintNFTOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrictMintNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrictMintNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrictMintNFTOwnershipTransferred represents a OwnershipTransferred event raised by the StrictMintNFT contract.
type StrictMintNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrictMintNFT *StrictMintNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StrictMintNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrictMintNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFTOwnershipTransferredIterator{contract: _StrictMintNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrictMintNFT *StrictMintNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StrictMintNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrictMintNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrictMintNFTOwnershipTransferred)
				if err := _StrictMintNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrictMintNFT *StrictMintNFTFilterer) ParseOwnershipTransferred(log types.Log) (*StrictMintNFTOwnershipTransferred, error) {
	event := new(StrictMintNFTOwnershipTransferred)
	if err := _StrictMintNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrictMintNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StrictMintNFT contract.
type StrictMintNFTTransferIterator struct {
	Event *StrictMintNFTTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrictMintNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrictMintNFTTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrictMintNFTTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrictMintNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrictMintNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrictMintNFTTransfer represents a Transfer event raised by the StrictMintNFT contract.
type StrictMintNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StrictMintNFT *StrictMintNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StrictMintNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StrictMintNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StrictMintNFTTransferIterator{contract: _StrictMintNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StrictMintNFT *StrictMintNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StrictMintNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StrictMintNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrictMintNFTTransfer)
				if err := _StrictMintNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StrictMintNFT *StrictMintNFTFilterer) ParseTransfer(log types.Log) (*StrictMintNFTTransfer, error) {
	event := new(StrictMintNFTTransfer)
	if err := _StrictMintNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
