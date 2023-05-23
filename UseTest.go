// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UseTest

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"}],\"name\":\"ExistingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_string1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_string2\",\"type\":\"string\"}],\"name\":\"getHex_Hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"args\",\"type\":\"string[]\"}],\"name\":\"printf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"}],\"name\":\"setA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610b90806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80639a329a261461005c578063d46300fd1461008c578063dd005a83146100aa578063de4dce4f146100da578063ee919d50146100f6575b600080fd5b6100766004803603810190610071919061063c565b610126565b6040516100839190610733565b60405180910390f35b6100946101d1565b6040516100a1919061076e565b60405180910390f35b6100c460048036038101906100bf9190610789565b610268565b6040516100d19190610733565b60405180910390f35b6100f460048036038101906100ef919061085f565b610313565b005b610110600480360381019061010b91906108b8565b610356565b60405161011d919061076e565b60405180910390f35b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a329a2684846040518363ffffffff1660e01b81526004016101839291906109f1565b600060405180830381865afa1580156101a0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906101c99190610a98565b905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630dbe671f6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561023f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102639190610af6565b905090565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd005a8384846040518363ffffffff1660e01b81526004016102c5929190610b23565b600060405180830381865afa1580156102e2573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061030b9190610a98565b905092915050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee919d50836040518263ffffffff1660e01b81526004016103b2919061076e565b6020604051808303816000875af11580156103d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f59190610af6565b9050919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104638261041a565b810181811067ffffffffffffffff821117156104825761048161042b565b5b80604052505050565b60006104956103fc565b90506104a1828261045a565b919050565b600067ffffffffffffffff8211156104c1576104c061042b565b5b6104ca8261041a565b9050602081019050919050565b82818337600083830152505050565b60006104f96104f4846104a6565b61048b565b90508281526020810184848401111561051557610514610415565b5b6105208482856104d7565b509392505050565b600082601f83011261053d5761053c610410565b5b813561054d8482602086016104e6565b91505092915050565b600067ffffffffffffffff8211156105715761057061042b565b5b602082029050602081019050919050565b600080fd5b600061059a61059584610556565b61048b565b905080838252602082019050602084028301858111156105bd576105bc610582565b5b835b8181101561060457803567ffffffffffffffff8111156105e2576105e1610410565b5b8086016105ef8982610528565b855260208501945050506020810190506105bf565b5050509392505050565b600082601f83011261062357610622610410565b5b8135610633848260208601610587565b91505092915050565b6000806040838503121561065357610652610406565b5b600083013567ffffffffffffffff8111156106715761067061040b565b5b61067d85828601610528565b925050602083013567ffffffffffffffff81111561069e5761069d61040b565b5b6106aa8582860161060e565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106ee5780820151818401526020810190506106d3565b60008484015250505050565b6000610705826106b4565b61070f81856106bf565b935061071f8185602086016106d0565b6107288161041a565b840191505092915050565b6000602082019050818103600083015261074d81846106fa565b905092915050565b6000819050919050565b61076881610755565b82525050565b6000602082019050610783600083018461075f565b92915050565b600080604083850312156107a05761079f610406565b5b600083013567ffffffffffffffff8111156107be576107bd61040b565b5b6107ca85828601610528565b925050602083013567ffffffffffffffff8111156107eb576107ea61040b565b5b6107f785828601610528565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061082c82610801565b9050919050565b61083c81610821565b811461084757600080fd5b50565b60008135905061085981610833565b92915050565b60006020828403121561087557610874610406565b5b60006108838482850161084a565b91505092915050565b61089581610755565b81146108a057600080fd5b50565b6000813590506108b28161088c565b92915050565b6000602082840312156108ce576108cd610406565b5b60006108dc848285016108a3565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061092d826106b4565b6109378185610911565b93506109478185602086016106d0565b6109508161041a565b840191505092915050565b60006109678383610922565b905092915050565b6000602082019050919050565b6000610987826108e5565b61099181856108f0565b9350836020820285016109a385610901565b8060005b858110156109df57848403895281516109c0858261095b565b94506109cb8361096f565b925060208a019950506001810190506109a7565b50829750879550505050505092915050565b60006040820190508181036000830152610a0b81856106fa565b90508181036020830152610a1f818461097c565b90509392505050565b6000610a3b610a36846104a6565b61048b565b905082815260208101848484011115610a5757610a56610415565b5b610a628482856106d0565b509392505050565b600082601f830112610a7f57610a7e610410565b5b8151610a8f848260208601610a28565b91505092915050565b600060208284031215610aae57610aad610406565b5b600082015167ffffffffffffffff811115610acc57610acb61040b565b5b610ad884828501610a6a565b91505092915050565b600081519050610af08161088c565b92915050565b600060208284031215610b0c57610b0b610406565b5b6000610b1a84828501610ae1565b91505092915050565b60006040820190508181036000830152610b3d81856106fa565b90508181036020830152610b5181846106fa565b9050939250505056fea26469706673582212200fd61768f7e28d0cb71c2ad11a7330e2da976028fb6e084194d043eacc59bb2a64736f6c63430008120033",
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

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256 result)
func (_Api *ApiCaller) GetA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256 result)
func (_Api *ApiSession) GetA() (*big.Int, error) {
	return _Api.Contract.GetA(&_Api.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256 result)
func (_Api *ApiCallerSession) GetA() (*big.Int, error) {
	return _Api.Contract.GetA(&_Api.CallOpts)
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
// Solidity: function printf(string format, string[] args) view returns(string)
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
// Solidity: function printf(string format, string[] args) view returns(string)
func (_Api *ApiSession) Printf(format string, args []string) (string, error) {
	return _Api.Contract.Printf(&_Api.CallOpts, format, args)
}

// Printf is a free data retrieval call binding the contract method 0x9a329a26.
//
// Solidity: function printf(string format, string[] args) view returns(string)
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

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _a) returns(uint256 result)
func (_Api *ApiTransactor) SetA(opts *bind.TransactOpts, _a *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setA", _a)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _a) returns(uint256 result)
func (_Api *ApiSession) SetA(_a *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetA(&_Api.TransactOpts, _a)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _a) returns(uint256 result)
func (_Api *ApiTransactorSession) SetA(_a *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetA(&_Api.TransactOpts, _a)
}
