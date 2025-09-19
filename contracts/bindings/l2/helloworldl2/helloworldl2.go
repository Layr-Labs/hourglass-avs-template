// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helloworldl2

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

// HelloWorldL2MetaData contains all meta data concerning the HelloWorldL2 contract.
var HelloWorldL2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getMessage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMessage\",\"inputs\":[{\"name\":\"newMessage\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60c0604052601360809081527f48656c6c6f20576f726c642066726f6d204c320000000000000000000000000060a0525f9061003b90826100e5565b50348015610047575f5ffd5b5061019f565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061007557607f821691505b60208210810361009357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156100e057805f5260205f20601f840160051c810160208510156100be5750805b601f840160051c820191505b818110156100dd575f81556001016100ca565b50505b505050565b81516001600160401b038111156100fe576100fe61004d565b6101128161010c8454610061565b84610099565b6020601f821160018114610144575f831561012d5750848201515b5f19600385901b1c1916600184901b1784556100dd565b5f84815260208120601f198516915b828110156101735787850151825560209485019460019092019101610153565b508482101561019057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b61037a806101ac5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063368b877214610038578063ce6d41de1461004d575b5f5ffd5b61004b61004636600461011d565b61006b565b005b61005561007a565b60405161006291906101d0565b60405180910390f35b5f6100768282610289565b5050565b60605f805461008890610205565b80601f01602080910402602001604051908101604052809291908181526020018280546100b490610205565b80156100ff5780601f106100d6576101008083540402835291602001916100ff565b820191905f5260205f20905b8154815290600101906020018083116100e257829003601f168201915b5050505050905090565b634e487b7160e01b5f52604160045260245ffd5b5f6020828403121561012d575f5ffd5b813567ffffffffffffffff811115610143575f5ffd5b8201601f81018413610153575f5ffd5b803567ffffffffffffffff81111561016d5761016d610109565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561019c5761019c610109565b6040528181528282016020018610156101b3575f5ffd5b816020840160208301375f91810160200191909152949350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b600181811c9082168061021957607f821691505b60208210810361023757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561028457805f5260205f20601f840160051c810160208510156102625750805b601f840160051c820191505b81811015610281575f815560010161026e565b50505b505050565b815167ffffffffffffffff8111156102a3576102a3610109565b6102b7816102b18454610205565b8461023d565b6020601f8211600181146102e9575f83156102d25750848201515b5f19600385901b1c1916600184901b178455610281565b5f84815260208120601f198516915b8281101561031857878501518255602094850194600190920191016102f8565b508482101561033557868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea264697066735822122023f98770f5013e0e165154cad1cd16c91c5538b83e37a9333af1d628cb5055cd64736f6c634300081b0033",
}

// HelloWorldL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloWorldL2MetaData.ABI instead.
var HelloWorldL2ABI = HelloWorldL2MetaData.ABI

// HelloWorldL2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelloWorldL2MetaData.Bin instead.
var HelloWorldL2Bin = HelloWorldL2MetaData.Bin

// DeployHelloWorldL2 deploys a new Ethereum contract, binding an instance of HelloWorldL2 to it.
func DeployHelloWorldL2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HelloWorldL2, error) {
	parsed, err := HelloWorldL2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelloWorldL2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HelloWorldL2{HelloWorldL2Caller: HelloWorldL2Caller{contract: contract}, HelloWorldL2Transactor: HelloWorldL2Transactor{contract: contract}, HelloWorldL2Filterer: HelloWorldL2Filterer{contract: contract}}, nil
}

// HelloWorldL2 is an auto generated Go binding around an Ethereum contract.
type HelloWorldL2 struct {
	HelloWorldL2Caller     // Read-only binding to the contract
	HelloWorldL2Transactor // Write-only binding to the contract
	HelloWorldL2Filterer   // Log filterer for contract events
}

// HelloWorldL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type HelloWorldL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloWorldL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloWorldL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloWorldL2Session struct {
	Contract     *HelloWorldL2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloWorldL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloWorldL2CallerSession struct {
	Contract *HelloWorldL2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HelloWorldL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloWorldL2TransactorSession struct {
	Contract     *HelloWorldL2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HelloWorldL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type HelloWorldL2Raw struct {
	Contract *HelloWorldL2 // Generic contract binding to access the raw methods on
}

// HelloWorldL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloWorldL2CallerRaw struct {
	Contract *HelloWorldL2Caller // Generic read-only contract binding to access the raw methods on
}

// HelloWorldL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloWorldL2TransactorRaw struct {
	Contract *HelloWorldL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloWorldL2 creates a new instance of HelloWorldL2, bound to a specific deployed contract.
func NewHelloWorldL2(address common.Address, backend bind.ContractBackend) (*HelloWorldL2, error) {
	contract, err := bindHelloWorldL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2{HelloWorldL2Caller: HelloWorldL2Caller{contract: contract}, HelloWorldL2Transactor: HelloWorldL2Transactor{contract: contract}, HelloWorldL2Filterer: HelloWorldL2Filterer{contract: contract}}, nil
}

// NewHelloWorldL2Caller creates a new read-only instance of HelloWorldL2, bound to a specific deployed contract.
func NewHelloWorldL2Caller(address common.Address, caller bind.ContractCaller) (*HelloWorldL2Caller, error) {
	contract, err := bindHelloWorldL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2Caller{contract: contract}, nil
}

// NewHelloWorldL2Transactor creates a new write-only instance of HelloWorldL2, bound to a specific deployed contract.
func NewHelloWorldL2Transactor(address common.Address, transactor bind.ContractTransactor) (*HelloWorldL2Transactor, error) {
	contract, err := bindHelloWorldL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2Transactor{contract: contract}, nil
}

// NewHelloWorldL2Filterer creates a new log filterer instance of HelloWorldL2, bound to a specific deployed contract.
func NewHelloWorldL2Filterer(address common.Address, filterer bind.ContractFilterer) (*HelloWorldL2Filterer, error) {
	contract, err := bindHelloWorldL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2Filterer{contract: contract}, nil
}

// bindHelloWorldL2 binds a generic wrapper to an already deployed contract.
func bindHelloWorldL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelloWorldL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorldL2 *HelloWorldL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloWorldL2.Contract.HelloWorldL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorldL2 *HelloWorldL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.HelloWorldL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorldL2 *HelloWorldL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.HelloWorldL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorldL2 *HelloWorldL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloWorldL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorldL2 *HelloWorldL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorldL2 *HelloWorldL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.contract.Transact(opts, method, params...)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_HelloWorldL2 *HelloWorldL2Caller) GetMessage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "getMessage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_HelloWorldL2 *HelloWorldL2Session) GetMessage() (string, error) {
	return _HelloWorldL2.Contract.GetMessage(&_HelloWorldL2.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_HelloWorldL2 *HelloWorldL2CallerSession) GetMessage() (string, error) {
	return _HelloWorldL2.Contract.GetMessage(&_HelloWorldL2.CallOpts)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_HelloWorldL2 *HelloWorldL2Transactor) SetMessage(opts *bind.TransactOpts, newMessage string) (*types.Transaction, error) {
	return _HelloWorldL2.contract.Transact(opts, "setMessage", newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_HelloWorldL2 *HelloWorldL2Session) SetMessage(newMessage string) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.SetMessage(&_HelloWorldL2.TransactOpts, newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_HelloWorldL2 *HelloWorldL2TransactorSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.SetMessage(&_HelloWorldL2.TransactOpts, newMessage)
}
