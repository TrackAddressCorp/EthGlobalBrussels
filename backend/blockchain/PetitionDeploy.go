// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// PetitionContractMetaData contains all meta data concerning the PetitionContract contract.
var PetitionContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_petition_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_petition_text\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get_petition_text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_petition_title\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_signers_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_world_id\",\"type\":\"string\"}],\"name\":\"sign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161106b38038061106b833981810160405281019061003191906101d6565b815f908161003f9190610459565b50806001908161004f9190610459565b505f60025f6101000a8154816cffffffffffffffffffffffffff02191690836cffffffffffffffffffffffffff1602179055505050610528565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100e8826100a2565b810181811067ffffffffffffffff82111715610107576101066100b2565b5b80604052505050565b5f610119610089565b905061012582826100df565b919050565b5f67ffffffffffffffff821115610144576101436100b2565b5b61014d826100a2565b9050602081019050919050565b8281835e5f83830152505050565b5f61017a6101758461012a565b610110565b9050828152602081018484840111156101965761019561009e565b5b6101a184828561015a565b509392505050565b5f82601f8301126101bd576101bc61009a565b5b81516101cd848260208601610168565b91505092915050565b5f80604083850312156101ec576101eb610092565b5b5f83015167ffffffffffffffff81111561020957610208610096565b5b610215858286016101a9565b925050602083015167ffffffffffffffff81111561023657610235610096565b5b610242858286016101a9565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061029a57607f821691505b6020821081036102ad576102ac610256565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261030f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826102d4565b61031986836102d4565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61035d61035861035384610331565b61033a565b610331565b9050919050565b5f819050919050565b61037683610343565b61038a61038282610364565b8484546102e0565b825550505050565b5f90565b61039e610392565b6103a981848461036d565b505050565b5b818110156103cc576103c15f82610396565b6001810190506103af565b5050565b601f821115610411576103e2816102b3565b6103eb846102c5565b810160208510156103fa578190505b61040e610406856102c5565b8301826103ae565b50505b505050565b5f82821c905092915050565b5f6104315f1984600802610416565b1980831691505092915050565b5f6104498383610422565b9150826002028217905092915050565b6104628261024c565b67ffffffffffffffff81111561047b5761047a6100b2565b5b6104858254610283565b6104908282856103d0565b5f60209050601f8311600181146104c1575f84156104af578287015190505b6104b9858261043e565b865550610520565b601f1984166104cf866102b3565b5f5b828110156104f6578489015182556001820191506020850194506020810190506104d1565b86831015610513578489015161050f601f891682610422565b8355505b6001600288020188555050505b505050505050565b610b36806105355f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806331c075181461004e57806379d6348d1461006c578063991936f4146100885780639d737b63146100a6575b5f80fd5b6100566100c4565b60405161006391906104de565b60405180910390f35b6100866004803603810190610081919061063b565b610154565b005b6100906103af565b60405161009d91906104de565b60405180910390f35b6100ae61043e565b6040516100bb919061069a565b60405180910390f35b6060600180546100d3906106e0565b80601f01602080910402602001604051908101604052809291908181526020018280546100ff906106e0565b801561014a5780601f106101215761010080835404028352916020019161014a565b820191905f5260205f20905b81548152906001019060200180831161012d57829003601f168201915b5050505050905090565b629896806cffffffffffffffffffffffffff1660025f9054906101000a90046cffffffffffffffffffffffffff166cffffffffffffffffffffffffff16106101d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101c890610780565b60405180910390fd5b6040518060a001604052808281526020014281526020016040518060400160405280600381526020017f4e2f41000000000000000000000000000000000000000000000000000000000081525081526020015f61ffff1681526020016040518060400160405280600381526020017f4e2f410000000000000000000000000000000000000000000000000000000000815250815250600360025f9054906101000a90046cffffffffffffffffffffffffff166cffffffffffffffffffffffffff166298968081106102a5576102a461079e565b5b600502015f820151815f0190816102bc9190610968565b506020820151816001015560408201518160020190816102dc9190610968565b506060820151816003015f6101000a81548161ffff021916908361ffff16021790555060808201518160040190816103149190610968565b5090505060025f81819054906101000a90046cffffffffffffffffffffffffff168092919061034290610a7c565b91906101000a8154816cffffffffffffffffffffffffff02191690836cffffffffffffffffffffffffff1602179055505060016302faf083826040516103889190610aea565b90815260200160405180910390205f6101000a81548160ff02191690831515021790555050565b60605f80546103bd906106e0565b80601f01602080910402602001604051908101604052809291908181526020018280546103e9906106e0565b80156104345780601f1061040b57610100808354040283529160200191610434565b820191905f5260205f20905b81548152906001019060200180831161041757829003601f168201915b5050505050905090565b5f60025f9054906101000a90046cffffffffffffffffffffffffff166cffffffffffffffffffffffffff16905090565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6104b08261046e565b6104ba8185610478565b93506104ca818560208601610488565b6104d381610496565b840191505092915050565b5f6020820190508181035f8301526104f681846104a6565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61054d82610496565b810181811067ffffffffffffffff8211171561056c5761056b610517565b5b80604052505050565b5f61057e6104fe565b905061058a8282610544565b919050565b5f67ffffffffffffffff8211156105a9576105a8610517565b5b6105b282610496565b9050602081019050919050565b828183375f83830152505050565b5f6105df6105da8461058f565b610575565b9050828152602081018484840111156105fb576105fa610513565b5b6106068482856105bf565b509392505050565b5f82601f8301126106225761062161050f565b5b81356106328482602086016105cd565b91505092915050565b5f602082840312156106505761064f610507565b5b5f82013567ffffffffffffffff81111561066d5761066c61050b565b5b6106798482850161060e565b91505092915050565b5f819050919050565b61069481610682565b82525050565b5f6020820190506106ad5f83018461068b565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806106f757607f821691505b60208210810361070a576107096106b3565b5b50919050565b7f546865207065746974696f6e20686173207265616368656420746865206d61785f8201527f696d756d206e756d626572206f66207369676e6572732e000000000000000000602082015250565b5f61076a603783610478565b915061077582610710565b604082019050919050565b5f6020820190508181035f8301526107978161075e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826107ec565b61083186836107ec565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61086c61086761086284610682565b610849565b610682565b9050919050565b5f819050919050565b61088583610852565b61089961089182610873565b8484546107f8565b825550505050565b5f90565b6108ad6108a1565b6108b881848461087c565b505050565b5b818110156108db576108d05f826108a5565b6001810190506108be565b5050565b601f821115610920576108f1816107cb565b6108fa846107dd565b81016020851015610909578190505b61091d610915856107dd565b8301826108bd565b50505b505050565b5f82821c905092915050565b5f6109405f1984600802610925565b1980831691505092915050565b5f6109588383610931565b9150826002028217905092915050565b6109718261046e565b67ffffffffffffffff81111561098a57610989610517565b5b61099482546106e0565b61099f8282856108df565b5f60209050601f8311600181146109d0575f84156109be578287015190505b6109c8858261094d565b865550610a2f565b601f1984166109de866107cb565b5f5b82811015610a05578489015182556001820191506020850194506020810190506109e0565b86831015610a225784890151610a1e601f891682610931565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6cffffffffffffffffffffffffff82169050919050565b5f610a8682610a64565b91506cffffffffffffffffffffffffff8203610aa557610aa4610a37565b5b600182019050919050565b5f81905092915050565b5f610ac48261046e565b610ace8185610ab0565b9350610ade818560208601610488565b80840191505092915050565b5f610af58284610aba565b91508190509291505056fea2646970667358221220617384bcd6899c41905c0c64dc51e364e9b99132961adeb724428c58db306bbd64736f6c634300081a0033",
}

// PetitionContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PetitionContractMetaData.ABI instead.
var PetitionContractABI = PetitionContractMetaData.ABI

// PetitionContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PetitionContractMetaData.Bin instead.
var PetitionContractBin = PetitionContractMetaData.Bin

// DeployPetitionContract deploys a new Ethereum contract, binding an instance of PetitionContract to it.
func DeployPetitionContract(auth *bind.TransactOpts, backend bind.ContractBackend, _petition_title string, _petition_text string) (common.Address, *types.Transaction, *PetitionContract, error) {
	parsed, err := PetitionContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PetitionContractBin), backend, _petition_title, _petition_text)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PetitionContract{PetitionContractCaller: PetitionContractCaller{contract: contract}, PetitionContractTransactor: PetitionContractTransactor{contract: contract}, PetitionContractFilterer: PetitionContractFilterer{contract: contract}}, nil
}

// PetitionContract is an auto generated Go binding around an Ethereum contract.
type PetitionContract struct {
	PetitionContractCaller     // Read-only binding to the contract
	PetitionContractTransactor // Write-only binding to the contract
	PetitionContractFilterer   // Log filterer for contract events
}

// PetitionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PetitionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PetitionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PetitionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PetitionContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PetitionContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PetitionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PetitionContractSession struct {
	Contract     *PetitionContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PetitionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PetitionContractCallerSession struct {
	Contract *PetitionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PetitionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PetitionContractTransactorSession struct {
	Contract     *PetitionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PetitionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PetitionContractRaw struct {
	Contract *PetitionContract // Generic contract binding to access the raw methods on
}

// PetitionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PetitionContractCallerRaw struct {
	Contract *PetitionContractCaller // Generic read-only contract binding to access the raw methods on
}

// PetitionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PetitionContractTransactorRaw struct {
	Contract *PetitionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPetitionContract creates a new instance of PetitionContract, bound to a specific deployed contract.
func NewPetitionContract(address common.Address, backend bind.ContractBackend) (*PetitionContract, error) {
	contract, err := bindPetitionContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PetitionContract{PetitionContractCaller: PetitionContractCaller{contract: contract}, PetitionContractTransactor: PetitionContractTransactor{contract: contract}, PetitionContractFilterer: PetitionContractFilterer{contract: contract}}, nil
}

// NewPetitionContractCaller creates a new read-only instance of PetitionContract, bound to a specific deployed contract.
func NewPetitionContractCaller(address common.Address, caller bind.ContractCaller) (*PetitionContractCaller, error) {
	contract, err := bindPetitionContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PetitionContractCaller{contract: contract}, nil
}

// NewPetitionContractTransactor creates a new write-only instance of PetitionContract, bound to a specific deployed contract.
func NewPetitionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PetitionContractTransactor, error) {
	contract, err := bindPetitionContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PetitionContractTransactor{contract: contract}, nil
}

// NewPetitionContractFilterer creates a new log filterer instance of PetitionContract, bound to a specific deployed contract.
func NewPetitionContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PetitionContractFilterer, error) {
	contract, err := bindPetitionContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PetitionContractFilterer{contract: contract}, nil
}

// bindPetitionContract binds a generic wrapper to an already deployed contract.
func bindPetitionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PetitionContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PetitionContract *PetitionContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PetitionContract.Contract.PetitionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PetitionContract *PetitionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PetitionContract.Contract.PetitionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PetitionContract *PetitionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PetitionContract.Contract.PetitionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PetitionContract *PetitionContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PetitionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PetitionContract *PetitionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PetitionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PetitionContract *PetitionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PetitionContract.Contract.contract.Transact(opts, method, params...)
}

// GetPetitionText is a free data retrieval call binding the contract method 0x31c07518.
//
// Solidity: function get_petition_text() view returns(string)
func (_PetitionContract *PetitionContractCaller) GetPetitionText(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PetitionContract.contract.Call(opts, &out, "get_petition_text")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPetitionText is a free data retrieval call binding the contract method 0x31c07518.
//
// Solidity: function get_petition_text() view returns(string)
func (_PetitionContract *PetitionContractSession) GetPetitionText() (string, error) {
	return _PetitionContract.Contract.GetPetitionText(&_PetitionContract.CallOpts)
}

// GetPetitionText is a free data retrieval call binding the contract method 0x31c07518.
//
// Solidity: function get_petition_text() view returns(string)
func (_PetitionContract *PetitionContractCallerSession) GetPetitionText() (string, error) {
	return _PetitionContract.Contract.GetPetitionText(&_PetitionContract.CallOpts)
}

// GetPetitionTitle is a free data retrieval call binding the contract method 0x991936f4.
//
// Solidity: function get_petition_title() view returns(string)
func (_PetitionContract *PetitionContractCaller) GetPetitionTitle(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PetitionContract.contract.Call(opts, &out, "get_petition_title")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPetitionTitle is a free data retrieval call binding the contract method 0x991936f4.
//
// Solidity: function get_petition_title() view returns(string)
func (_PetitionContract *PetitionContractSession) GetPetitionTitle() (string, error) {
	return _PetitionContract.Contract.GetPetitionTitle(&_PetitionContract.CallOpts)
}

// GetPetitionTitle is a free data retrieval call binding the contract method 0x991936f4.
//
// Solidity: function get_petition_title() view returns(string)
func (_PetitionContract *PetitionContractCallerSession) GetPetitionTitle() (string, error) {
	return _PetitionContract.Contract.GetPetitionTitle(&_PetitionContract.CallOpts)
}

// GetSignersCount is a free data retrieval call binding the contract method 0x9d737b63.
//
// Solidity: function get_signers_count() view returns(uint256)
func (_PetitionContract *PetitionContractCaller) GetSignersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PetitionContract.contract.Call(opts, &out, "get_signers_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignersCount is a free data retrieval call binding the contract method 0x9d737b63.
//
// Solidity: function get_signers_count() view returns(uint256)
func (_PetitionContract *PetitionContractSession) GetSignersCount() (*big.Int, error) {
	return _PetitionContract.Contract.GetSignersCount(&_PetitionContract.CallOpts)
}

// GetSignersCount is a free data retrieval call binding the contract method 0x9d737b63.
//
// Solidity: function get_signers_count() view returns(uint256)
func (_PetitionContract *PetitionContractCallerSession) GetSignersCount() (*big.Int, error) {
	return _PetitionContract.Contract.GetSignersCount(&_PetitionContract.CallOpts)
}

// Sign is a paid mutator transaction binding the contract method 0x79d6348d.
//
// Solidity: function sign(string _world_id) returns()
func (_PetitionContract *PetitionContractTransactor) Sign(opts *bind.TransactOpts, _world_id string) (*types.Transaction, error) {
	return _PetitionContract.contract.Transact(opts, "sign", _world_id)
}

// Sign is a paid mutator transaction binding the contract method 0x79d6348d.
//
// Solidity: function sign(string _world_id) returns()
func (_PetitionContract *PetitionContractSession) Sign(_world_id string) (*types.Transaction, error) {
	return _PetitionContract.Contract.Sign(&_PetitionContract.TransactOpts, _world_id)
}

// Sign is a paid mutator transaction binding the contract method 0x79d6348d.
//
// Solidity: function sign(string _world_id) returns()
func (_PetitionContract *PetitionContractTransactorSession) Sign(_world_id string) (*types.Transaction, error) {
	return _PetitionContract.Contract.Sign(&_PetitionContract.TransactOpts, _world_id)
}
