// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package taskavsregistrar

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

// ITaskAVSRegistrarBaseTypesAvsConfig is an auto generated low-level Go binding around an user-defined struct.
type ITaskAVSRegistrarBaseTypesAvsConfig struct {
	AggregatorOperatorSetId uint32
	ExecutorOperatorSetIds  []uint32
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// TaskAVSRegistrarMetaData contains all meta data concerning the TaskAVSRegistrar contract.
var TaskAVSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorToAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedOperators\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvsConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskAVSRegistrarBaseTypes.AvsConfig\",\"components\":[{\"name\":\"aggregatorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialConfig\",\"type\":\"tuple\",\"internalType\":\"structITaskAVSRegistrarBaseTypes.AvsConfig\",\"components\":[{\"name\":\"aggregatorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorAllowed\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorFromAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAvsConfig\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structITaskAVSRegistrarBaseTypes.AvsConfig\",\"components\":[{\"name\":\"aggregatorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AvsConfigSet\",\"inputs\":[{\"name\":\"aggregatorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DuplicateExecutorOperatorSetId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetIdsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAggregatorOperatorSetId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyInAllowlist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotInAllowlist\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051611d9b380380611d9b83398101604081905261002e91610145565b6001600160a01b03808416608052821660a05282828280808484610050610072565b50506001600160a01b031660c05250610067610072565b50505050505061018f565b5f54610100900460ff16156100dd5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161461012c575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610142575f5ffd5b50565b5f5f5f60608486031215610157575f5ffd5b83516101628161012e565b60208501519093506101738161012e565b60408501519092506101848161012e565b809150509250925092565b60805160a05160c051611bc66101d55f395f81816101b90152610a7c01525f81816101600152610d7901525f81816102a90152818161047b01526107ee0152611bc65ff3fe608060405234801561000f575f5ffd5b506004361061011c575f3560e01c80638481931d116100a9578063ca8aa7c71161006e578063ca8aa7c7146102a4578063d1f2e81d146102cb578063de1164bb146102de578063f2fde38b146102f6578063f91ff80c14610309575f5ffd5b80638481931d146102165780638da5cb5b14610236578063aaacc42514610247578063b52657871461025a578063c63fd50214610291575f5ffd5b806341f548f0116100ef57806341f548f01461019f5780634657e26a146101b45780636591666a146101db578063715018a6146101ee5780637fe94e16146101f6575f5ffd5b80630a4d3d29146101205780631017873a14610135578063303ca956146101485780633ec45c7e1461015b575b5f5ffd5b61013361012e3660046113c0565b61031c565b005b6101336101433660046113c0565b6103c6565b610133610156366004611439565b610470565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6101a7610502565b6040516101969190611495565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b6101336101e9366004611566565b6105b2565b6101336105e8565b6102096102043660046115b0565b6105fb565b60405161019691906115ca565b610229610224366004611615565b610624565b604051610196919061162e565b6096546001600160a01b0316610182565b61013361025536600461172d565b6106cd565b610281610268366004611615565b5f546201000090046001600160a01b0390811691161490565b6040519015158152602001610196565b61013361029f366004611786565b6107e3565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b6101336102d9366004611841565b61089c565b5f54610182906201000090046001600160a01b031681565b610133610304366004611615565b6108b0565b6102816103173660046113c0565b610926565b610324610959565b61034f8160c85f610334866109b3565b81526020019081526020015f20610a1690919063ffffffff16565b61036c57604051630444d2e160e21b815260040160405180910390fd5b6040805183516001600160a01b03908116825260208086015163ffffffff169083015283169101604051908190038120907f533bf6e1348e64eb9448930dece3436586c031d36722adbc7ccb479809128806905f90a35050565b6103ce610959565b6103f98160c85f6103de866109b3565b81526020019081526020015f20610a2a90919063ffffffff16565b610416576040516386e0613f60e01b815260040160405180910390fd5b6040805183516001600160a01b03908116825260208086015163ffffffff169083015283169101604051908190038120907ffe795219771c42bdbb61ef308cc2b33e1e35b35a3364499b99b2ec2287f20c8c905f90a35050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104b9576040516335e5ec5560e21b815260040160405180910390fd5b836001600160a01b03167ff8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d183836040516104f492919061187a565b60405180910390a250505050565b604080518082019091525f81526060602082015260408051808201825260fa805463ffffffff16825260fb805484516020828102820181019096528181529394929383860193909291908301828280156105a457602002820191905f5260205f20905f905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116105675790505b505050505081525050905090565b816105bc81610a3e565b6105d95760405163932d94f760e01b815260040160405180910390fd5b6105e38383610ae8565b505050565b6105f0610959565b6105f95f610b4f565b565b606061061e60c85f61060c856109b3565b81526020019081526020015f20610ba0565b92915050565b6001600160a01b0381165f90815260326020526040902080546060919061064a906118b7565b80601f0160208091040260200160405190810160405280929190818152602001828054610676906118b7565b80156106c15780601f10610698576101008083540402835291602001916106c1565b820191905f5260205f20905b8154815290600101906020018083116106a457829003601f168201915b50505050509050919050565b5f54610100900460ff16158080156106eb57505f54600160ff909116105b806107045750303b15801561070457505f5460ff166001145b61076c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff19166001179055801561078d575f805461ff0019166101001790555b610798848484610bac565b80156107dd575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461082c576040516335e5ec5560e21b815260040160405180910390fd5b6108398685858585610bed565b610844868585610cd5565b6108518685858585610e12565b856001600160a01b03167f9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d858560405161088c92919061187a565b60405180910390a2505050505050565b6108a4610959565b6108ad81610e2b565b50565b6108b8610959565b6001600160a01b03811661091d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610763565b6108ad81610b4f565b5f6109528260c85f610937876109b3565b81526020019081526020015f20610fa690919063ffffffff16565b9392505050565b6096546001600160a01b031633146105f95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610763565b5f815f0151826020015163ffffffff166040516020016109fe92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261061e906118ef565b5f610952836001600160a01b038416610fc7565b5f610952836001600160a01b0384166110aa565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af1158015610ac4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061061e9190611912565b6001600160a01b0382165f908152603260205260409020610b098282611975565b50816001600160a01b03167f0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b82604051610b43919061162e565b60405180910390a25050565b609680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60605f610952836110f6565b5f54610100900460ff16610bd25760405162461bcd60e51b815260040161076390611a2f565b610bdb8261114e565b610be48361117c565b6105e381610e2b565b5f5b63ffffffff8116841115610ccd5760fa5463ffffffff90811690869086908416818110610c1e57610c1e611a7a565b9050602002016020810190610c339190611a8e565b63ffffffff1603610cbb57604080518082019091525f546201000090046001600160a01b03168152610c9e9060208101878763ffffffff8616818110610c7b57610c7b611a7a565b9050602002016020810190610c909190611a8e565b63ffffffff16905287610926565b610cbb57604051630444d2e160e21b815260040160405180910390fd5b80610cc581611abb565b915050610bef565b505050505050565b5f5b63ffffffff81168211156107dd57604080518082019091525f80546201000090046001600160a01b031682529060208101858563ffffffff8616818110610d2057610d20611a7a565b9050602002016020810190610d359190611a8e565b63ffffffff90811690915260405163bd30a0b960e01b815282516001600160a01b0390811660048301526020840151909216602482015287821660448201529192507f0000000000000000000000000000000000000000000000000000000000000000169063bd30a0b990606401602060405180830381865afa158015610dbe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610de29190611912565b610dff5760405163815589fb60e01b815260040160405180910390fd5b5080610e0a81611abb565b915050610cd7565b5f610e1f82840184611adf565b9050610ccd8682610ae8565b5f81602001515111610e4f5760405162c10e4560e71b815260040160405180910390fd5b5f5b816020015151811015610f2b5781602001518181518110610e7457610e74611a7a565b602002602001015163ffffffff16825f015163ffffffff1603610eaa576040516311eef66d60e01b815260040160405180910390fd5b801580610f0657506020820151610ec2600183611b10565b81518110610ed257610ed2611a7a565b602002602001015163ffffffff1682602001518281518110610ef657610ef6611a7a565b602002602001015163ffffffff16115b610f2357604051631efe361b60e21b815260040160405180910390fd5b600101610e51565b50805160fa805463ffffffff191663ffffffff9092169190911781556020808301518051849392610f619260fb9291019061122a565b5050815160208301516040517f836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef519350610f9b929190611b23565b60405180910390a150565b6001600160a01b0381165f9081526001830160205260408120541515610952565b5f81815260018301602052604081205480156110a1575f610fe9600183611b10565b85549091505f90610ffc90600190611b10565b905081811461105b575f865f01828154811061101a5761101a611a7a565b905f5260205f200154905080875f01848154811061103a5761103a611a7a565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061106c5761106c611b7c565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061061e565b5f91505061061e565b5f8181526001830160205260408120546110ef57508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561061e565b505f61061e565b6060815f018054806020026020016040519081016040528092919081815260200182805480156106c157602002820191905f5260205f20905b81548152602001906001019080831161112f5750505050509050919050565b5f54610100900460ff166111745760405162461bcd60e51b815260040161076390611a2f565b61091d6111d2565b5f54610100900460ff166111a25760405162461bcd60e51b815260040161076390611a2f565b5f80546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b5050505050565b5f54610100900460ff166111f85760405162461bcd60e51b815260040161076390611a2f565b6105f95f54610100900460ff166112215760405162461bcd60e51b815260040161076390611a2f565b6105f933610b4f565b828054828255905f5260205f20906007016008900481019282156112c6579160200282015f5b8382111561129457835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302611250565b80156112c45782816101000a81549063ffffffff0219169055600401602081600301049283019260010302611294565b505b506112d29291506112d6565b5090565b5b808211156112d2575f81556001016112d7565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611320576113206112ea565b60405290565b604051601f8201601f191681016001600160401b038111828210171561134e5761134e6112ea565b604052919050565b80356001600160a01b038116811461136c575f5ffd5b919050565b803563ffffffff8116811461136c575f5ffd5b5f60408284031215611394575f5ffd5b61139c6112fe565b90506113a782611356565b81526113b560208301611371565b602082015292915050565b5f5f606083850312156113d1575f5ffd5b6113db8484611384565b91506113e960408401611356565b90509250929050565b5f5f83601f840112611402575f5ffd5b5081356001600160401b03811115611418575f5ffd5b6020830191508360208260051b8501011115611432575f5ffd5b9250929050565b5f5f5f5f6060858703121561144c575f5ffd5b61145585611356565b935061146360208601611356565b925060408501356001600160401b0381111561147d575f5ffd5b611489878288016113f2565b95989497509550505050565b6020808252825163ffffffff1682820152828101516040808401528051606084018190525f929190910190829060808501905b808310156114f15763ffffffff84511682526020820191506020840193506001830192506114c8565b5095945050505050565b5f82601f83011261150a575f5ffd5b81356001600160401b03811115611523576115236112ea565b611536601f8201601f1916602001611326565b81815284602083860101111561154a575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f60408385031215611577575f5ffd5b61158083611356565b915060208301356001600160401b0381111561159a575f5ffd5b6115a6858286016114fb565b9150509250929050565b5f604082840312156115c0575f5ffd5b6109528383611384565b602080825282518282018190525f918401906040840190835b8181101561160a5783516001600160a01b03168352602093840193909201916001016115e3565b509095945050505050565b5f60208284031215611625575f5ffd5b61095282611356565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60408284031215611673575f5ffd5b61167b6112fe565b905061168682611371565b815260208201356001600160401b038111156116a0575f5ffd5b8201601f810184136116b0575f5ffd5b80356001600160401b038111156116c9576116c96112ea565b8060051b6116d960208201611326565b918252602081840181019290810190878411156116f4575f5ffd5b6020850194505b8385101561171d5761170c85611371565b8252602094850194909101906116fb565b6020860152509295945050505050565b5f5f5f6060848603121561173f575f5ffd5b61174884611356565b925061175660208501611356565b915060408401356001600160401b03811115611770575f5ffd5b61177c86828701611663565b9150509250925092565b5f5f5f5f5f5f6080878903121561179b575f5ffd5b6117a487611356565b95506117b260208801611356565b945060408701356001600160401b038111156117cc575f5ffd5b6117d889828a016113f2565b90955093505060608701356001600160401b038111156117f6575f5ffd5b8701601f81018913611806575f5ffd5b80356001600160401b0381111561181b575f5ffd5b89602082840101111561182c575f5ffd5b60208201935080925050509295509295509295565b5f60208284031215611851575f5ffd5b81356001600160401b03811115611866575f5ffd5b61187284828501611663565b949350505050565b602080825281018290525f8360408301825b858110156114f15763ffffffff6118a284611371565b1682526020928301929091019060010161188c565b600181811c908216806118cb57607f821691505b6020821081036118e957634e487b7160e01b5f52602260045260245ffd5b50919050565b805160208083015191908110156118e9575f1960209190910360031b1b16919050565b5f60208284031215611922575f5ffd5b81518015158114610952575f5ffd5b601f8211156105e357805f5260205f20601f840160051c810160208510156119565750805b601f840160051c820191505b818110156111cb575f8155600101611962565b81516001600160401b0381111561198e5761198e6112ea565b6119a28161199c84546118b7565b84611931565b6020601f8211600181146119d4575f83156119bd5750848201515b5f19600385901b1c1916600184901b1784556111cb565b5f84815260208120601f198516915b82811015611a0357878501518255602094850194600190920191016119e3565b5084821015611a2057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215611a9e575f5ffd5b61095282611371565b634e487b7160e01b5f52601160045260245ffd5b5f63ffffffff821663ffffffff8103611ad657611ad6611aa7565b60010192915050565b5f60208284031215611aef575f5ffd5b81356001600160401b03811115611b04575f5ffd5b611872848285016114fb565b8181038181111561061e5761061e611aa7565b5f6040820163ffffffff85168352604060208401528084518083526060850191506020860192505f5b81811015611b7057835163ffffffff16835260209384019390920191600101611b4c565b50909695505050505050565b634e487b7160e01b5f52603160045260245ffdfea26469706673582212209ef5b80ba28a805b1b96bd1b0955794349ca4f97975f51660edff867fd695f9464736f6c634300081b0033",
}

// TaskAVSRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskAVSRegistrarMetaData.ABI instead.
var TaskAVSRegistrarABI = TaskAVSRegistrarMetaData.ABI

// TaskAVSRegistrarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskAVSRegistrarMetaData.Bin instead.
var TaskAVSRegistrarBin = TaskAVSRegistrarMetaData.Bin

// DeployTaskAVSRegistrar deploys a new Ethereum contract, binding an instance of TaskAVSRegistrar to it.
func DeployTaskAVSRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend, _allocationManager common.Address, _keyRegistrar common.Address, _permissionController common.Address) (common.Address, *types.Transaction, *TaskAVSRegistrar, error) {
	parsed, err := TaskAVSRegistrarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaskAVSRegistrarBin), backend, _allocationManager, _keyRegistrar, _permissionController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TaskAVSRegistrar{TaskAVSRegistrarCaller: TaskAVSRegistrarCaller{contract: contract}, TaskAVSRegistrarTransactor: TaskAVSRegistrarTransactor{contract: contract}, TaskAVSRegistrarFilterer: TaskAVSRegistrarFilterer{contract: contract}}, nil
}

// TaskAVSRegistrar is an auto generated Go binding around an Ethereum contract.
type TaskAVSRegistrar struct {
	TaskAVSRegistrarCaller     // Read-only binding to the contract
	TaskAVSRegistrarTransactor // Write-only binding to the contract
	TaskAVSRegistrarFilterer   // Log filterer for contract events
}

// TaskAVSRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskAVSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskAVSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskAVSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskAVSRegistrarSession struct {
	Contract     *TaskAVSRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskAVSRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskAVSRegistrarCallerSession struct {
	Contract *TaskAVSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TaskAVSRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskAVSRegistrarTransactorSession struct {
	Contract     *TaskAVSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TaskAVSRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskAVSRegistrarRaw struct {
	Contract *TaskAVSRegistrar // Generic contract binding to access the raw methods on
}

// TaskAVSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskAVSRegistrarCallerRaw struct {
	Contract *TaskAVSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// TaskAVSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskAVSRegistrarTransactorRaw struct {
	Contract *TaskAVSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskAVSRegistrar creates a new instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrar(address common.Address, backend bind.ContractBackend) (*TaskAVSRegistrar, error) {
	contract, err := bindTaskAVSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrar{TaskAVSRegistrarCaller: TaskAVSRegistrarCaller{contract: contract}, TaskAVSRegistrarTransactor: TaskAVSRegistrarTransactor{contract: contract}, TaskAVSRegistrarFilterer: TaskAVSRegistrarFilterer{contract: contract}}, nil
}

// NewTaskAVSRegistrarCaller creates a new read-only instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*TaskAVSRegistrarCaller, error) {
	contract, err := bindTaskAVSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarCaller{contract: contract}, nil
}

// NewTaskAVSRegistrarTransactor creates a new write-only instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskAVSRegistrarTransactor, error) {
	contract, err := bindTaskAVSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarTransactor{contract: contract}, nil
}

// NewTaskAVSRegistrarFilterer creates a new log filterer instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskAVSRegistrarFilterer, error) {
	contract, err := bindTaskAVSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarFilterer{contract: contract}, nil
}

// bindTaskAVSRegistrar binds a generic wrapper to an already deployed contract.
func bindTaskAVSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskAVSRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskAVSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) AllocationManager() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.AllocationManager(&_TaskAVSRegistrar.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) AllocationManager() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.AllocationManager(&_TaskAVSRegistrar.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) Avs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "avs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) Avs() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.Avs(&_TaskAVSRegistrar.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) Avs() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.Avs(&_TaskAVSRegistrar.CallOpts)
}

// GetAllowedOperators is a free data retrieval call binding the contract method 0x7fe94e16.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) GetAllowedOperators(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "getAllowedOperators", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedOperators is a free data retrieval call binding the contract method 0x7fe94e16.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) GetAllowedOperators(operatorSet OperatorSet) ([]common.Address, error) {
	return _TaskAVSRegistrar.Contract.GetAllowedOperators(&_TaskAVSRegistrar.CallOpts, operatorSet)
}

// GetAllowedOperators is a free data retrieval call binding the contract method 0x7fe94e16.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) GetAllowedOperators(operatorSet OperatorSet) ([]common.Address, error) {
	return _TaskAVSRegistrar.Contract.GetAllowedOperators(&_TaskAVSRegistrar.CallOpts, operatorSet)
}

// GetAvsConfig is a free data retrieval call binding the contract method 0x41f548f0.
//
// Solidity: function getAvsConfig() view returns((uint32,uint32[]))
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) GetAvsConfig(opts *bind.CallOpts) (ITaskAVSRegistrarBaseTypesAvsConfig, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "getAvsConfig")

	if err != nil {
		return *new(ITaskAVSRegistrarBaseTypesAvsConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ITaskAVSRegistrarBaseTypesAvsConfig)).(*ITaskAVSRegistrarBaseTypesAvsConfig)

	return out0, err

}

// GetAvsConfig is a free data retrieval call binding the contract method 0x41f548f0.
//
// Solidity: function getAvsConfig() view returns((uint32,uint32[]))
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) GetAvsConfig() (ITaskAVSRegistrarBaseTypesAvsConfig, error) {
	return _TaskAVSRegistrar.Contract.GetAvsConfig(&_TaskAVSRegistrar.CallOpts)
}

// GetAvsConfig is a free data retrieval call binding the contract method 0x41f548f0.
//
// Solidity: function getAvsConfig() view returns((uint32,uint32[]))
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) GetAvsConfig() (ITaskAVSRegistrarBaseTypesAvsConfig, error) {
	return _TaskAVSRegistrar.Contract.GetAvsConfig(&_TaskAVSRegistrar.CallOpts)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) GetOperatorSocket(opts *bind.CallOpts, operator common.Address) (string, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "getOperatorSocket", operator)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _TaskAVSRegistrar.Contract.GetOperatorSocket(&_TaskAVSRegistrar.CallOpts, operator)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _TaskAVSRegistrar.Contract.GetOperatorSocket(&_TaskAVSRegistrar.CallOpts, operator)
}

// IsOperatorAllowed is a free data retrieval call binding the contract method 0xf91ff80c.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) IsOperatorAllowed(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "isOperatorAllowed", operatorSet, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorAllowed is a free data retrieval call binding the contract method 0xf91ff80c.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) IsOperatorAllowed(operatorSet OperatorSet, operator common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.IsOperatorAllowed(&_TaskAVSRegistrar.CallOpts, operatorSet, operator)
}

// IsOperatorAllowed is a free data retrieval call binding the contract method 0xf91ff80c.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) IsOperatorAllowed(operatorSet OperatorSet, operator common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.IsOperatorAllowed(&_TaskAVSRegistrar.CallOpts, operatorSet, operator)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) KeyRegistrar() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.KeyRegistrar(&_TaskAVSRegistrar.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) KeyRegistrar() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.KeyRegistrar(&_TaskAVSRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) Owner() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.Owner(&_TaskAVSRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) Owner() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.Owner(&_TaskAVSRegistrar.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) PermissionController() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.PermissionController(&_TaskAVSRegistrar.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) PermissionController() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.PermissionController(&_TaskAVSRegistrar.CallOpts)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) SupportsAVS(opts *bind.CallOpts, _avs common.Address) (bool, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "supportsAVS", _avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.SupportsAVS(&_TaskAVSRegistrar.CallOpts, _avs)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.SupportsAVS(&_TaskAVSRegistrar.CallOpts, _avs)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x1017873a.
//
// Solidity: function addOperatorToAllowlist((address,uint32) operatorSet, address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) AddOperatorToAllowlist(opts *bind.TransactOpts, operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "addOperatorToAllowlist", operatorSet, operator)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x1017873a.
//
// Solidity: function addOperatorToAllowlist((address,uint32) operatorSet, address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) AddOperatorToAllowlist(operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.AddOperatorToAllowlist(&_TaskAVSRegistrar.TransactOpts, operatorSet, operator)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x1017873a.
//
// Solidity: function addOperatorToAllowlist((address,uint32) operatorSet, address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) AddOperatorToAllowlist(operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.AddOperatorToAllowlist(&_TaskAVSRegistrar.TransactOpts, operatorSet, operator)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address , uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "deregisterOperator", operator, arg1, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address , uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) DeregisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.DeregisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, arg1, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address , uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) DeregisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.DeregisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, arg1, operatorSetIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xaaacc425.
//
// Solidity: function initialize(address _avs, address _owner, (uint32,uint32[]) _initialConfig) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) Initialize(opts *bind.TransactOpts, _avs common.Address, _owner common.Address, _initialConfig ITaskAVSRegistrarBaseTypesAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "initialize", _avs, _owner, _initialConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xaaacc425.
//
// Solidity: function initialize(address _avs, address _owner, (uint32,uint32[]) _initialConfig) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) Initialize(_avs common.Address, _owner common.Address, _initialConfig ITaskAVSRegistrarBaseTypesAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.Initialize(&_TaskAVSRegistrar.TransactOpts, _avs, _owner, _initialConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xaaacc425.
//
// Solidity: function initialize(address _avs, address _owner, (uint32,uint32[]) _initialConfig) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) Initialize(_avs common.Address, _owner common.Address, _initialConfig ITaskAVSRegistrarBaseTypesAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.Initialize(&_TaskAVSRegistrar.TransactOpts, _avs, _owner, _initialConfig)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address , uint32[] operatorSetIds, bytes data) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "registerOperator", operator, arg1, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address , uint32[] operatorSetIds, bytes data) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) RegisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RegisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, arg1, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address , uint32[] operatorSetIds, bytes data) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) RegisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RegisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, arg1, operatorSetIds, data)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x0a4d3d29.
//
// Solidity: function removeOperatorFromAllowlist((address,uint32) operatorSet, address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) RemoveOperatorFromAllowlist(opts *bind.TransactOpts, operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "removeOperatorFromAllowlist", operatorSet, operator)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x0a4d3d29.
//
// Solidity: function removeOperatorFromAllowlist((address,uint32) operatorSet, address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) RemoveOperatorFromAllowlist(operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RemoveOperatorFromAllowlist(&_TaskAVSRegistrar.TransactOpts, operatorSet, operator)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x0a4d3d29.
//
// Solidity: function removeOperatorFromAllowlist((address,uint32) operatorSet, address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) RemoveOperatorFromAllowlist(operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RemoveOperatorFromAllowlist(&_TaskAVSRegistrar.TransactOpts, operatorSet, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RenounceOwnership(&_TaskAVSRegistrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RenounceOwnership(&_TaskAVSRegistrar.TransactOpts)
}

// SetAvsConfig is a paid mutator transaction binding the contract method 0xd1f2e81d.
//
// Solidity: function setAvsConfig((uint32,uint32[]) config) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) SetAvsConfig(opts *bind.TransactOpts, config ITaskAVSRegistrarBaseTypesAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "setAvsConfig", config)
}

// SetAvsConfig is a paid mutator transaction binding the contract method 0xd1f2e81d.
//
// Solidity: function setAvsConfig((uint32,uint32[]) config) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) SetAvsConfig(config ITaskAVSRegistrarBaseTypesAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.SetAvsConfig(&_TaskAVSRegistrar.TransactOpts, config)
}

// SetAvsConfig is a paid mutator transaction binding the contract method 0xd1f2e81d.
//
// Solidity: function setAvsConfig((uint32,uint32[]) config) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) SetAvsConfig(config ITaskAVSRegistrarBaseTypesAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.SetAvsConfig(&_TaskAVSRegistrar.TransactOpts, config)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TransferOwnership(&_TaskAVSRegistrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TransferOwnership(&_TaskAVSRegistrar.TransactOpts, newOwner)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) UpdateSocket(opts *bind.TransactOpts, operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "updateSocket", operator, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) UpdateSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.UpdateSocket(&_TaskAVSRegistrar.TransactOpts, operator, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) UpdateSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.UpdateSocket(&_TaskAVSRegistrar.TransactOpts, operator, socket)
}

// TaskAVSRegistrarAvsConfigSetIterator is returned from FilterAvsConfigSet and is used to iterate over the raw logs and unpacked data for AvsConfigSet events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarAvsConfigSetIterator struct {
	Event *TaskAVSRegistrarAvsConfigSet // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarAvsConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarAvsConfigSet)
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
		it.Event = new(TaskAVSRegistrarAvsConfigSet)
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
func (it *TaskAVSRegistrarAvsConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarAvsConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarAvsConfigSet represents a AvsConfigSet event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarAvsConfigSet struct {
	AggregatorOperatorSetId uint32
	ExecutorOperatorSetIds  []uint32
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAvsConfigSet is a free log retrieval operation binding the contract event 0x836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef51.
//
// Solidity: event AvsConfigSet(uint32 aggregatorOperatorSetId, uint32[] executorOperatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterAvsConfigSet(opts *bind.FilterOpts) (*TaskAVSRegistrarAvsConfigSetIterator, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "AvsConfigSet")
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarAvsConfigSetIterator{contract: _TaskAVSRegistrar.contract, event: "AvsConfigSet", logs: logs, sub: sub}, nil
}

// WatchAvsConfigSet is a free log subscription operation binding the contract event 0x836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef51.
//
// Solidity: event AvsConfigSet(uint32 aggregatorOperatorSetId, uint32[] executorOperatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchAvsConfigSet(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarAvsConfigSet) (event.Subscription, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "AvsConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarAvsConfigSet)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "AvsConfigSet", log); err != nil {
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

// ParseAvsConfigSet is a log parse operation binding the contract event 0x836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef51.
//
// Solidity: event AvsConfigSet(uint32 aggregatorOperatorSetId, uint32[] executorOperatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseAvsConfigSet(log types.Log) (*TaskAVSRegistrarAvsConfigSet, error) {
	event := new(TaskAVSRegistrarAvsConfigSet)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "AvsConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarInitializedIterator struct {
	Event *TaskAVSRegistrarInitialized // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarInitialized)
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
		it.Event = new(TaskAVSRegistrarInitialized)
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
func (it *TaskAVSRegistrarInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarInitialized represents a Initialized event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaskAVSRegistrarInitializedIterator, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarInitializedIterator{contract: _TaskAVSRegistrar.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarInitialized) (event.Subscription, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarInitialized)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseInitialized(log types.Log) (*TaskAVSRegistrarInitialized, error) {
	event := new(TaskAVSRegistrarInitialized)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorAddedToAllowlistIterator is returned from FilterOperatorAddedToAllowlist and is used to iterate over the raw logs and unpacked data for OperatorAddedToAllowlist events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorAddedToAllowlistIterator struct {
	Event *TaskAVSRegistrarOperatorAddedToAllowlist // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorAddedToAllowlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorAddedToAllowlist)
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
		it.Event = new(TaskAVSRegistrarOperatorAddedToAllowlist)
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
func (it *TaskAVSRegistrarOperatorAddedToAllowlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorAddedToAllowlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorAddedToAllowlist represents a OperatorAddedToAllowlist event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorAddedToAllowlist struct {
	OperatorSet OperatorSet
	Operator    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToAllowlist is a free log retrieval operation binding the contract event 0xfe795219771c42bdbb61ef308cc2b33e1e35b35a3364499b99b2ec2287f20c8c.
//
// Solidity: event OperatorAddedToAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorAddedToAllowlist(opts *bind.FilterOpts, operatorSet []OperatorSet, operator []common.Address) (*TaskAVSRegistrarOperatorAddedToAllowlistIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorAddedToAllowlist", operatorSetRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorAddedToAllowlistIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorAddedToAllowlist", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToAllowlist is a free log subscription operation binding the contract event 0xfe795219771c42bdbb61ef308cc2b33e1e35b35a3364499b99b2ec2287f20c8c.
//
// Solidity: event OperatorAddedToAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorAddedToAllowlist(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorAddedToAllowlist, operatorSet []OperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorAddedToAllowlist", operatorSetRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorAddedToAllowlist)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorAddedToAllowlist", log); err != nil {
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

// ParseOperatorAddedToAllowlist is a log parse operation binding the contract event 0xfe795219771c42bdbb61ef308cc2b33e1e35b35a3364499b99b2ec2287f20c8c.
//
// Solidity: event OperatorAddedToAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorAddedToAllowlist(log types.Log) (*TaskAVSRegistrarOperatorAddedToAllowlist, error) {
	event := new(TaskAVSRegistrarOperatorAddedToAllowlist)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorAddedToAllowlist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorDeregisteredIterator struct {
	Event *TaskAVSRegistrarOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorDeregistered)
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
		it.Event = new(TaskAVSRegistrarOperatorDeregistered)
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
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorDeregistered represents a OperatorDeregistered event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorDeregistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorDeregisteredIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorDeregistered)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorDeregistered(log types.Log) (*TaskAVSRegistrarOperatorDeregistered, error) {
	event := new(TaskAVSRegistrarOperatorDeregistered)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorRegisteredIterator struct {
	Event *TaskAVSRegistrarOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorRegistered)
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
		it.Event = new(TaskAVSRegistrarOperatorRegistered)
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
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorRegistered represents a OperatorRegistered event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorRegistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorRegisteredIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorRegistered)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorRegistered(log types.Log) (*TaskAVSRegistrarOperatorRegistered, error) {
	event := new(TaskAVSRegistrarOperatorRegistered)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorRemovedFromAllowlistIterator is returned from FilterOperatorRemovedFromAllowlist and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromAllowlist events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorRemovedFromAllowlistIterator struct {
	Event *TaskAVSRegistrarOperatorRemovedFromAllowlist // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorRemovedFromAllowlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorRemovedFromAllowlist)
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
		it.Event = new(TaskAVSRegistrarOperatorRemovedFromAllowlist)
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
func (it *TaskAVSRegistrarOperatorRemovedFromAllowlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorRemovedFromAllowlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorRemovedFromAllowlist represents a OperatorRemovedFromAllowlist event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorRemovedFromAllowlist struct {
	OperatorSet OperatorSet
	Operator    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromAllowlist is a free log retrieval operation binding the contract event 0x533bf6e1348e64eb9448930dece3436586c031d36722adbc7ccb479809128806.
//
// Solidity: event OperatorRemovedFromAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorRemovedFromAllowlist(opts *bind.FilterOpts, operatorSet []OperatorSet, operator []common.Address) (*TaskAVSRegistrarOperatorRemovedFromAllowlistIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorRemovedFromAllowlist", operatorSetRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorRemovedFromAllowlistIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorRemovedFromAllowlist", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromAllowlist is a free log subscription operation binding the contract event 0x533bf6e1348e64eb9448930dece3436586c031d36722adbc7ccb479809128806.
//
// Solidity: event OperatorRemovedFromAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorRemovedFromAllowlist(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorRemovedFromAllowlist, operatorSet []OperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorRemovedFromAllowlist", operatorSetRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorRemovedFromAllowlist)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorRemovedFromAllowlist", log); err != nil {
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

// ParseOperatorRemovedFromAllowlist is a log parse operation binding the contract event 0x533bf6e1348e64eb9448930dece3436586c031d36722adbc7ccb479809128806.
//
// Solidity: event OperatorRemovedFromAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorRemovedFromAllowlist(log types.Log) (*TaskAVSRegistrarOperatorRemovedFromAllowlist, error) {
	event := new(TaskAVSRegistrarOperatorRemovedFromAllowlist)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorRemovedFromAllowlist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorSocketSetIterator is returned from FilterOperatorSocketSet and is used to iterate over the raw logs and unpacked data for OperatorSocketSet events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorSocketSetIterator struct {
	Event *TaskAVSRegistrarOperatorSocketSet // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorSocketSet)
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
		it.Event = new(TaskAVSRegistrarOperatorSocketSet)
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
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorSocketSet represents a OperatorSocketSet event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorSocketSet struct {
	Operator common.Address
	Socket   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketSet is a free log retrieval operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorSocketSet(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorSocketSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorSocketSetIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorSocketSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketSet is a free log subscription operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorSocketSet(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorSocketSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorSocketSet)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
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

// ParseOperatorSocketSet is a log parse operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorSocketSet(log types.Log) (*TaskAVSRegistrarOperatorSocketSet, error) {
	event := new(TaskAVSRegistrarOperatorSocketSet)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOwnershipTransferredIterator struct {
	Event *TaskAVSRegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOwnershipTransferred)
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
		it.Event = new(TaskAVSRegistrarOwnershipTransferred)
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
func (it *TaskAVSRegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaskAVSRegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOwnershipTransferredIterator{contract: _TaskAVSRegistrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOwnershipTransferred)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOwnershipTransferred(log types.Log) (*TaskAVSRegistrarOwnershipTransferred, error) {
	event := new(TaskAVSRegistrarOwnershipTransferred)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
