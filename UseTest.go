// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"}],\"name\":\"ExistingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_string1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_string2\",\"type\":\"string\"}],\"name\":\"getHex_Hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"args\",\"type\":\"string[]\"}],\"name\":\"printf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610746806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80639a329a2614610046578063dd005a8314610076578063de4dce4f146100a6575b600080fd5b610060600480360381019061005b9190610404565b6100c2565b60405161006d91906104fb565b60405180910390f35b610090600480360381019061008b919061051d565b6100d6565b60405161009d91906104fb565b60405180910390f35b6100c060048036038101906100bb91906105f3565b610181565b005b60606100ce83836100c2565b905092915050565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd005a8384846040518363ffffffff1660e01b8152600401610133929190610620565b600060405180830381865afa158015610150573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061017991906106c7565b905092915050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61022b826101e2565b810181811067ffffffffffffffff8211171561024a576102496101f3565b5b80604052505050565b600061025d6101c4565b90506102698282610222565b919050565b600067ffffffffffffffff821115610289576102886101f3565b5b610292826101e2565b9050602081019050919050565b82818337600083830152505050565b60006102c16102bc8461026e565b610253565b9050828152602081018484840111156102dd576102dc6101dd565b5b6102e884828561029f565b509392505050565b600082601f830112610305576103046101d8565b5b81356103158482602086016102ae565b91505092915050565b600067ffffffffffffffff821115610339576103386101f3565b5b602082029050602081019050919050565b600080fd5b600061036261035d8461031e565b610253565b905080838252602082019050602084028301858111156103855761038461034a565b5b835b818110156103cc57803567ffffffffffffffff8111156103aa576103a96101d8565b5b8086016103b789826102f0565b85526020850194505050602081019050610387565b5050509392505050565b600082601f8301126103eb576103ea6101d8565b5b81356103fb84826020860161034f565b91505092915050565b6000806040838503121561041b5761041a6101ce565b5b600083013567ffffffffffffffff811115610439576104386101d3565b5b610445858286016102f0565b925050602083013567ffffffffffffffff811115610466576104656101d3565b5b610472858286016103d6565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104b657808201518184015260208101905061049b565b60008484015250505050565b60006104cd8261047c565b6104d78185610487565b93506104e7818560208601610498565b6104f0816101e2565b840191505092915050565b6000602082019050818103600083015261051581846104c2565b905092915050565b60008060408385031215610534576105336101ce565b5b600083013567ffffffffffffffff811115610552576105516101d3565b5b61055e858286016102f0565b925050602083013567ffffffffffffffff81111561057f5761057e6101d3565b5b61058b858286016102f0565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105c082610595565b9050919050565b6105d0816105b5565b81146105db57600080fd5b50565b6000813590506105ed816105c7565b92915050565b600060208284031215610609576106086101ce565b5b6000610617848285016105de565b91505092915050565b6000604082019050818103600083015261063a81856104c2565b9050818103602083015261064e81846104c2565b90509392505050565b600061066a6106658461026e565b610253565b905082815260208101848484011115610686576106856101dd565b5b610691848285610498565b509392505050565b600082601f8301126106ae576106ad6101d8565b5b81516106be848260208601610657565b91505092915050565b6000602082840312156106dd576106dc6101ce565b5b600082015167ffffffffffffffff8111156106fb576106fa6101d3565b5b61070784828501610699565b9150509291505056fea2646970667358221220e5ba2a345c681ca1c199e491c7fa3bb51ddc1921ab551d3c479e2ab56377203364736f6c63430008120033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) view returns(string)
func (_Api *ApiCaller) GetHexHash(opts *bind.CallOpts, _string1 string, _string2 string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getHex_Hash", _string1, _string2)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) view returns(string)
func (_Api *ApiSession) GetHexHash(_string1 string, _string2 string) (string, error) {
	return _Api.Contract.GetHexHash(&_Api.CallOpts, _string1, _string2)
}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) view returns(string)
func (_Api *ApiCallerSession) GetHexHash(_string1 string, _string2 string) (string, error) {
	return _Api.Contract.GetHexHash(&_Api.CallOpts, _string1, _string2)
}

// Printf is a free data retrieval call binding the contract method 0x9a329a26.
//
// Solidity: function printf(string format, string[] args) pure returns(string)
func (_Api *ApiCaller) Printf(opts *bind.CallOpts, format string, args []string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "printf", format, args)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Printf is a free data retrieval call binding the contract method 0x9a329a26.
//
// Solidity: function printf(string format, string[] args) pure returns(string)
func (_Api *ApiSession) Printf(format string, args []string) (string, error) {
	return _Api.Contract.Printf(&_Api.CallOpts, format, args)
}

// Printf is a free data retrieval call binding the contract method 0x9a329a26.
//
// Solidity: function printf(string format, string[] args) pure returns(string)
func (_Api *ApiCallerSession) Printf(format string, args []string) (string, error) {
	return _Api.Contract.Printf(&_Api.CallOpts, format, args)
}

// ExistingContract is a paid mutator transaction binding the contract method 0xde4dce4f.
//
// Solidity: function ExistingContract(address _t) returns()
func (_Api *ApiTransactor) ExistingContract(opts *bind.TransactOpts, _t common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "ExistingContract", _t)
}

// ExistingContract is a paid mutator transaction binding the contract method 0xde4dce4f.
//
// Solidity: function ExistingContract(address _t) returns()
func (_Api *ApiSession) ExistingContract(_t common.Address) (*types.Transaction, error) {
	return _Api.Contract.ExistingContract(&_Api.TransactOpts, _t)
}

// ExistingContract is a paid mutator transaction binding the contract method 0xde4dce4f.
//
// Solidity: function ExistingContract(address _t) returns()
func (_Api *ApiTransactorSession) ExistingContract(_t common.Address) (*types.Transaction, error) {
	return _Api.Contract.ExistingContract(&_Api.TransactOpts, _t)
}
