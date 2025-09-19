// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avstaskhook

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

// ITaskMailboxTypesTaskParams is an auto generated low-level Go binding around an user-defined struct.
type ITaskMailboxTypesTaskParams struct {
	RefundCollector     common.Address
	ExecutorOperatorSet OperatorSet
	Payload             []byte
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// AVSTaskHookMetaData contains all meta data concerning the AVSTaskHook contract.
var AVSTaskHookMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateTaskFee\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handlePostTaskCreation\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handlePostTaskResultSubmission\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatePreTaskCreation\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatePreTaskResultSubmission\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103f38061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806309c5c4501461005957806351fe30981461006c578063ba33565d1461007e578063e06cd27e14610092578063e082467e146100c6575b5f5ffd5b61006a6100673660046100d4565b50565b005b61006a61007a366004610292565b5050565b61006a61008c3660046102dd565b50505050565b6100a56100a036600461035b565b505f90565b6040516bffffffffffffffffffffffff909116815260200160405180910390f35b61006a61007a366004610395565b5f602082840312156100e4575f5ffd5b5035919050565b80356001600160a01b0381168114610101575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff8111828210171561013d5761013d610106565b60405290565b6040805190810167ffffffffffffffff8111828210171561013d5761013d610106565b5f82601f830112610175575f5ffd5b813567ffffffffffffffff81111561018f5761018f610106565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156101be576101be610106565b6040528181528382016020018510156101d5575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f8183036080811215610202575f5ffd5b61020a61011a565b9150610215836100eb565b82526040601f1982011215610228575f5ffd5b50610231610143565b61023d602084016100eb565b8152604083013563ffffffff81168114610255575f5ffd5b602082810191909152820152606082013567ffffffffffffffff81111561027a575f5ffd5b61028684828501610166565b60408301525092915050565b5f5f604083850312156102a3575f5ffd5b6102ac836100eb565b9150602083013567ffffffffffffffff8111156102c7575f5ffd5b6102d3858286016101f1565b9150509250929050565b5f5f5f5f608085870312156102f0575f5ffd5b6102f9856100eb565b935060208501359250604085013567ffffffffffffffff81111561031b575f5ffd5b61032787828801610166565b925050606085013567ffffffffffffffff811115610343575f5ffd5b61034f87828801610166565b91505092959194509250565b5f6020828403121561036b575f5ffd5b813567ffffffffffffffff811115610381575f5ffd5b61038d848285016101f1565b949350505050565b5f5f604083850312156103a6575f5ffd5b6103af836100eb565b94602093909301359350505056fea2646970667358221220f0175f76a360416beb135abb16804e8eb41829011145692e3d1cdfc29270712e64736f6c634300081b0033",
}

// AVSTaskHookABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSTaskHookMetaData.ABI instead.
var AVSTaskHookABI = AVSTaskHookMetaData.ABI

// AVSTaskHookBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AVSTaskHookMetaData.Bin instead.
var AVSTaskHookBin = AVSTaskHookMetaData.Bin

// DeployAVSTaskHook deploys a new Ethereum contract, binding an instance of AVSTaskHook to it.
func DeployAVSTaskHook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AVSTaskHook, error) {
	parsed, err := AVSTaskHookMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AVSTaskHookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AVSTaskHook{AVSTaskHookCaller: AVSTaskHookCaller{contract: contract}, AVSTaskHookTransactor: AVSTaskHookTransactor{contract: contract}, AVSTaskHookFilterer: AVSTaskHookFilterer{contract: contract}}, nil
}

// AVSTaskHook is an auto generated Go binding around an Ethereum contract.
type AVSTaskHook struct {
	AVSTaskHookCaller     // Read-only binding to the contract
	AVSTaskHookTransactor // Write-only binding to the contract
	AVSTaskHookFilterer   // Log filterer for contract events
}

// AVSTaskHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSTaskHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSTaskHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSTaskHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSTaskHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSTaskHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSTaskHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSTaskHookSession struct {
	Contract     *AVSTaskHook      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AVSTaskHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSTaskHookCallerSession struct {
	Contract *AVSTaskHookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AVSTaskHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSTaskHookTransactorSession struct {
	Contract     *AVSTaskHookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AVSTaskHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSTaskHookRaw struct {
	Contract *AVSTaskHook // Generic contract binding to access the raw methods on
}

// AVSTaskHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSTaskHookCallerRaw struct {
	Contract *AVSTaskHookCaller // Generic read-only contract binding to access the raw methods on
}

// AVSTaskHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSTaskHookTransactorRaw struct {
	Contract *AVSTaskHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSTaskHook creates a new instance of AVSTaskHook, bound to a specific deployed contract.
func NewAVSTaskHook(address common.Address, backend bind.ContractBackend) (*AVSTaskHook, error) {
	contract, err := bindAVSTaskHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHook{AVSTaskHookCaller: AVSTaskHookCaller{contract: contract}, AVSTaskHookTransactor: AVSTaskHookTransactor{contract: contract}, AVSTaskHookFilterer: AVSTaskHookFilterer{contract: contract}}, nil
}

// NewAVSTaskHookCaller creates a new read-only instance of AVSTaskHook, bound to a specific deployed contract.
func NewAVSTaskHookCaller(address common.Address, caller bind.ContractCaller) (*AVSTaskHookCaller, error) {
	contract, err := bindAVSTaskHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHookCaller{contract: contract}, nil
}

// NewAVSTaskHookTransactor creates a new write-only instance of AVSTaskHook, bound to a specific deployed contract.
func NewAVSTaskHookTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSTaskHookTransactor, error) {
	contract, err := bindAVSTaskHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHookTransactor{contract: contract}, nil
}

// NewAVSTaskHookFilterer creates a new log filterer instance of AVSTaskHook, bound to a specific deployed contract.
func NewAVSTaskHookFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSTaskHookFilterer, error) {
	contract, err := bindAVSTaskHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHookFilterer{contract: contract}, nil
}

// bindAVSTaskHook binds a generic wrapper to an already deployed contract.
func bindAVSTaskHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSTaskHookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSTaskHook *AVSTaskHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSTaskHook.Contract.AVSTaskHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSTaskHook *AVSTaskHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.AVSTaskHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSTaskHook *AVSTaskHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.AVSTaskHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSTaskHook *AVSTaskHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSTaskHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSTaskHook *AVSTaskHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSTaskHook *AVSTaskHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.contract.Transact(opts, method, params...)
}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) ) view returns(uint96)
func (_AVSTaskHook *AVSTaskHookCaller) CalculateTaskFee(opts *bind.CallOpts, arg0 ITaskMailboxTypesTaskParams) (*big.Int, error) {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "calculateTaskFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) ) view returns(uint96)
func (_AVSTaskHook *AVSTaskHookSession) CalculateTaskFee(arg0 ITaskMailboxTypesTaskParams) (*big.Int, error) {
	return _AVSTaskHook.Contract.CalculateTaskFee(&_AVSTaskHook.CallOpts, arg0)
}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) ) view returns(uint96)
func (_AVSTaskHook *AVSTaskHookCallerSession) CalculateTaskFee(arg0 ITaskMailboxTypesTaskParams) (*big.Int, error) {
	return _AVSTaskHook.Contract.CalculateTaskFee(&_AVSTaskHook.CallOpts, arg0)
}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address , (address,(address,uint32),bytes) ) view returns()
func (_AVSTaskHook *AVSTaskHookCaller) ValidatePreTaskCreation(opts *bind.CallOpts, arg0 common.Address, arg1 ITaskMailboxTypesTaskParams) error {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "validatePreTaskCreation", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address , (address,(address,uint32),bytes) ) view returns()
func (_AVSTaskHook *AVSTaskHookSession) ValidatePreTaskCreation(arg0 common.Address, arg1 ITaskMailboxTypesTaskParams) error {
	return _AVSTaskHook.Contract.ValidatePreTaskCreation(&_AVSTaskHook.CallOpts, arg0, arg1)
}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address , (address,(address,uint32),bytes) ) view returns()
func (_AVSTaskHook *AVSTaskHookCallerSession) ValidatePreTaskCreation(arg0 common.Address, arg1 ITaskMailboxTypesTaskParams) error {
	return _AVSTaskHook.Contract.ValidatePreTaskCreation(&_AVSTaskHook.CallOpts, arg0, arg1)
}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address , bytes32 , bytes , bytes ) view returns()
func (_AVSTaskHook *AVSTaskHookCaller) ValidatePreTaskResultSubmission(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 []byte, arg3 []byte) error {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "validatePreTaskResultSubmission", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address , bytes32 , bytes , bytes ) view returns()
func (_AVSTaskHook *AVSTaskHookSession) ValidatePreTaskResultSubmission(arg0 common.Address, arg1 [32]byte, arg2 []byte, arg3 []byte) error {
	return _AVSTaskHook.Contract.ValidatePreTaskResultSubmission(&_AVSTaskHook.CallOpts, arg0, arg1, arg2, arg3)
}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address , bytes32 , bytes , bytes ) view returns()
func (_AVSTaskHook *AVSTaskHookCallerSession) ValidatePreTaskResultSubmission(arg0 common.Address, arg1 [32]byte, arg2 []byte, arg3 []byte) error {
	return _AVSTaskHook.Contract.ValidatePreTaskResultSubmission(&_AVSTaskHook.CallOpts, arg0, arg1, arg2, arg3)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 ) returns()
func (_AVSTaskHook *AVSTaskHookTransactor) HandlePostTaskCreation(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.contract.Transact(opts, "handlePostTaskCreation", arg0)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 ) returns()
func (_AVSTaskHook *AVSTaskHookSession) HandlePostTaskCreation(arg0 [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandlePostTaskCreation(&_AVSTaskHook.TransactOpts, arg0)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 ) returns()
func (_AVSTaskHook *AVSTaskHookTransactorSession) HandlePostTaskCreation(arg0 [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandlePostTaskCreation(&_AVSTaskHook.TransactOpts, arg0)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xe082467e.
//
// Solidity: function handlePostTaskResultSubmission(address , bytes32 ) returns()
func (_AVSTaskHook *AVSTaskHookTransactor) HandlePostTaskResultSubmission(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.contract.Transact(opts, "handlePostTaskResultSubmission", arg0, arg1)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xe082467e.
//
// Solidity: function handlePostTaskResultSubmission(address , bytes32 ) returns()
func (_AVSTaskHook *AVSTaskHookSession) HandlePostTaskResultSubmission(arg0 common.Address, arg1 [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandlePostTaskResultSubmission(&_AVSTaskHook.TransactOpts, arg0, arg1)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xe082467e.
//
// Solidity: function handlePostTaskResultSubmission(address , bytes32 ) returns()
func (_AVSTaskHook *AVSTaskHookTransactorSession) HandlePostTaskResultSubmission(arg0 common.Address, arg1 [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandlePostTaskResultSubmission(&_AVSTaskHook.TransactOpts, arg0, arg1)
}
