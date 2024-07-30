// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// IIntentDBCreateIntentOrderParam is an auto generated low-level Go binding around an user-defined struct.
type IIntentDBCreateIntentOrderParam struct {
	SrcIntentToken         common.Address
	UnderlyingAssets       common.Address
	UnderlyingAssetsAmount *big.Int
	PayOrderId             *big.Int
	BridgeFee              *big.Int
	Node                   common.Address
}

// IIntentDBExeIntentOrderParam is an auto generated low-level Go binding around an user-defined struct.
type IIntentDBExeIntentOrderParam struct {
	DstIntentToken common.Address
	PayOrderId     *big.Int
}

// IIntentDBUnderlyingAssetsInfo is an auto generated low-level Go binding around an user-defined struct.
type IIntentDBUnderlyingAssetsInfo struct {
	UnderlyingAssets       common.Address
	UnderlyingAssetsAmount *big.Int
	TokenDecimals          *big.Int
	BridgeFee              *big.Int
	PriceInBase            *big.Int
	PriceDecimals          *big.Int
}

// IIntentDBCallParam is an auto generated low-level Go binding around an user-defined struct.
type IIntentDBCallParam struct {
	Node         common.Address
	NodeCallData []byte
}

// IIntentDBIntentTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type IIntentDBIntentTokenInfo struct {
	IntentToken         common.Address
	IntentTokenAmount   *big.Int
	IntentTokenDecimals *big.Int
	BaseAssetsRate      *big.Int
	BASEASSETRATEBASE   *big.Int
}

// IntentDBMetaData contains all meta data concerning the IntentDB contract.
var IntentDBMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cancelOperator\",\"type\":\"address\"}],\"name\":\"IntentOrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIIntentToken\",\"name\":\"srcIntentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAssetsAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDecimals\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIIntentDB.UnderlyingAssetsInfo\",\"name\":\"underlyingAssetsInfo\",\"type\":\"tuple\"}],\"name\":\"IntentOrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vwManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIIntentToken\",\"name\":\"intentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"intentTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intentTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAssetsRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"BASE_ASSET_RATE_BASE\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIIntentDB.IntentTokenInfo\",\"name\":\"intentTokenInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAssetsAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDecimals\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIIntentDB.UnderlyingAssetsInfo\",\"name\":\"underlyingAssetsInfo\",\"type\":\"tuple\"}],\"name\":\"IntentOrderExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIIntentToken\",\"name\":\"srcIntentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAssetsAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"internalType\":\"structIIntentDB.CreateIntentOrderParam[]\",\"name\":\"cparams\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrderETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orderOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIIntentToken\",\"name\":\"srcIntentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAssetsAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"internalType\":\"structIIntentDB.CreateIntentOrderParam[]\",\"name\":\"cparams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeCallData\",\"type\":\"bytes\"}],\"internalType\":\"structIIntentDB.CallParam\",\"name\":\"callParam\",\"type\":\"tuple\"}],\"name\":\"createSrcIntentOrderETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dstIntentOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orderOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vwManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAssets\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAssetsAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceDecimals\",\"type\":\"uint256\"}],\"internalType\":\"structIIntentDB.UnderlyingAssetsInfo[]\",\"name\":\"underlyingAssetsInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"intentTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractIIntentToken\",\"name\":\"dstIntentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"}],\"internalType\":\"structIIntentDB.ExeIntentOrderParam[]\",\"name\":\"eparams\",\"type\":\"tuple[]\"}],\"name\":\"executeDstIntentOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"srcIntentOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"status\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"createTime\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"orderDataHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IntentDBABI is the input ABI used to generate the binding from.
// Deprecated: Use IntentDBMetaData.ABI instead.
var IntentDBABI = IntentDBMetaData.ABI

// IntentDB is an auto generated Go binding around an Ethereum contract.
type IntentDB struct {
	IntentDBCaller     // Read-only binding to the contract
	IntentDBTransactor // Write-only binding to the contract
	IntentDBFilterer   // Log filterer for contract events
}

// IntentDBCaller is an auto generated read-only Go binding around an Ethereum contract.
type IntentDBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentDBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IntentDBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentDBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IntentDBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentDBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IntentDBSession struct {
	Contract     *IntentDB         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IntentDBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IntentDBCallerSession struct {
	Contract *IntentDBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IntentDBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IntentDBTransactorSession struct {
	Contract     *IntentDBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IntentDBRaw is an auto generated low-level Go binding around an Ethereum contract.
type IntentDBRaw struct {
	Contract *IntentDB // Generic contract binding to access the raw methods on
}

// IntentDBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IntentDBCallerRaw struct {
	Contract *IntentDBCaller // Generic read-only contract binding to access the raw methods on
}

// IntentDBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IntentDBTransactorRaw struct {
	Contract *IntentDBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIntentDB creates a new instance of IntentDB, bound to a specific deployed contract.
func NewIntentDB(address common.Address, backend bind.ContractBackend) (*IntentDB, error) {
	contract, err := bindIntentDB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IntentDB{IntentDBCaller: IntentDBCaller{contract: contract}, IntentDBTransactor: IntentDBTransactor{contract: contract}, IntentDBFilterer: IntentDBFilterer{contract: contract}}, nil
}

// NewIntentDBCaller creates a new read-only instance of IntentDB, bound to a specific deployed contract.
func NewIntentDBCaller(address common.Address, caller bind.ContractCaller) (*IntentDBCaller, error) {
	contract, err := bindIntentDB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IntentDBCaller{contract: contract}, nil
}

// NewIntentDBTransactor creates a new write-only instance of IntentDB, bound to a specific deployed contract.
func NewIntentDBTransactor(address common.Address, transactor bind.ContractTransactor) (*IntentDBTransactor, error) {
	contract, err := bindIntentDB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IntentDBTransactor{contract: contract}, nil
}

// NewIntentDBFilterer creates a new log filterer instance of IntentDB, bound to a specific deployed contract.
func NewIntentDBFilterer(address common.Address, filterer bind.ContractFilterer) (*IntentDBFilterer, error) {
	contract, err := bindIntentDB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IntentDBFilterer{contract: contract}, nil
}

// bindIntentDB binds a generic wrapper to an already deployed contract.
func bindIntentDB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IntentDBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntentDB *IntentDBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IntentDB.Contract.IntentDBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntentDB *IntentDBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntentDB.Contract.IntentDBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntentDB *IntentDBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IntentDB.Contract.IntentDBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntentDB *IntentDBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IntentDB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntentDB *IntentDBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntentDB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntentDB *IntentDBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IntentDB.Contract.contract.Transact(opts, method, params...)
}

// DstIntentOrder is a free data retrieval call binding the contract method 0xb27bc423.
//
// Solidity: function dstIntentOrder(address , uint256 ) view returns(bytes32)
func (_IntentDB *IntentDBCaller) DstIntentOrder(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IntentDB.contract.Call(opts, &out, "dstIntentOrder", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DstIntentOrder is a free data retrieval call binding the contract method 0xb27bc423.
//
// Solidity: function dstIntentOrder(address , uint256 ) view returns(bytes32)
func (_IntentDB *IntentDBSession) DstIntentOrder(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _IntentDB.Contract.DstIntentOrder(&_IntentDB.CallOpts, arg0, arg1)
}

// DstIntentOrder is a free data retrieval call binding the contract method 0xb27bc423.
//
// Solidity: function dstIntentOrder(address , uint256 ) view returns(bytes32)
func (_IntentDB *IntentDBCallerSession) DstIntentOrder(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _IntentDB.Contract.DstIntentOrder(&_IntentDB.CallOpts, arg0, arg1)
}

// SrcIntentOrder is a free data retrieval call binding the contract method 0x24c6c33b.
//
// Solidity: function srcIntentOrder(uint256 ) view returns(address node, uint64 status, uint32 createTime, bytes32 orderDataHash)
func (_IntentDB *IntentDBCaller) SrcIntentOrder(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Node          common.Address
	Status        uint64
	CreateTime    uint32
	OrderDataHash [32]byte
}, error) {
	var out []interface{}
	err := _IntentDB.contract.Call(opts, &out, "srcIntentOrder", arg0)

	outstruct := new(struct {
		Node          common.Address
		Status        uint64
		CreateTime    uint32
		OrderDataHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Node = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.CreateTime = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.OrderDataHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SrcIntentOrder is a free data retrieval call binding the contract method 0x24c6c33b.
//
// Solidity: function srcIntentOrder(uint256 ) view returns(address node, uint64 status, uint32 createTime, bytes32 orderDataHash)
func (_IntentDB *IntentDBSession) SrcIntentOrder(arg0 *big.Int) (struct {
	Node          common.Address
	Status        uint64
	CreateTime    uint32
	OrderDataHash [32]byte
}, error) {
	return _IntentDB.Contract.SrcIntentOrder(&_IntentDB.CallOpts, arg0)
}

// SrcIntentOrder is a free data retrieval call binding the contract method 0x24c6c33b.
//
// Solidity: function srcIntentOrder(uint256 ) view returns(address node, uint64 status, uint32 createTime, bytes32 orderDataHash)
func (_IntentDB *IntentDBCallerSession) SrcIntentOrder(arg0 *big.Int) (struct {
	Node          common.Address
	Status        uint64
	CreateTime    uint32
	OrderDataHash [32]byte
}, error) {
	return _IntentDB.Contract.SrcIntentOrder(&_IntentDB.CallOpts, arg0)
}

// CancelOrderETH is a paid mutator transaction binding the contract method 0x253f8024.
//
// Solidity: function cancelOrderETH(address sender, address receiver, (address,address,uint256,uint256,uint256,address)[] cparams) payable returns()
func (_IntentDB *IntentDBTransactor) CancelOrderETH(opts *bind.TransactOpts, sender common.Address, receiver common.Address, cparams []IIntentDBCreateIntentOrderParam) (*types.Transaction, error) {
	return _IntentDB.contract.Transact(opts, "cancelOrderETH", sender, receiver, cparams)
}

// CancelOrderETH is a paid mutator transaction binding the contract method 0x253f8024.
//
// Solidity: function cancelOrderETH(address sender, address receiver, (address,address,uint256,uint256,uint256,address)[] cparams) payable returns()
func (_IntentDB *IntentDBSession) CancelOrderETH(sender common.Address, receiver common.Address, cparams []IIntentDBCreateIntentOrderParam) (*types.Transaction, error) {
	return _IntentDB.Contract.CancelOrderETH(&_IntentDB.TransactOpts, sender, receiver, cparams)
}

// CancelOrderETH is a paid mutator transaction binding the contract method 0x253f8024.
//
// Solidity: function cancelOrderETH(address sender, address receiver, (address,address,uint256,uint256,uint256,address)[] cparams) payable returns()
func (_IntentDB *IntentDBTransactorSession) CancelOrderETH(sender common.Address, receiver common.Address, cparams []IIntentDBCreateIntentOrderParam) (*types.Transaction, error) {
	return _IntentDB.Contract.CancelOrderETH(&_IntentDB.TransactOpts, sender, receiver, cparams)
}

// CreateSrcIntentOrderETH is a paid mutator transaction binding the contract method 0xca145f3c.
//
// Solidity: function createSrcIntentOrderETH(address orderOwner, address receiver, (address,address,uint256,uint256,uint256,address)[] cparams, (address,bytes) callParam) payable returns()
func (_IntentDB *IntentDBTransactor) CreateSrcIntentOrderETH(opts *bind.TransactOpts, orderOwner common.Address, receiver common.Address, cparams []IIntentDBCreateIntentOrderParam, callParam IIntentDBCallParam) (*types.Transaction, error) {
	return _IntentDB.contract.Transact(opts, "createSrcIntentOrderETH", orderOwner, receiver, cparams, callParam)
}

// CreateSrcIntentOrderETH is a paid mutator transaction binding the contract method 0xca145f3c.
//
// Solidity: function createSrcIntentOrderETH(address orderOwner, address receiver, (address,address,uint256,uint256,uint256,address)[] cparams, (address,bytes) callParam) payable returns()
func (_IntentDB *IntentDBSession) CreateSrcIntentOrderETH(orderOwner common.Address, receiver common.Address, cparams []IIntentDBCreateIntentOrderParam, callParam IIntentDBCallParam) (*types.Transaction, error) {
	return _IntentDB.Contract.CreateSrcIntentOrderETH(&_IntentDB.TransactOpts, orderOwner, receiver, cparams, callParam)
}

// CreateSrcIntentOrderETH is a paid mutator transaction binding the contract method 0xca145f3c.
//
// Solidity: function createSrcIntentOrderETH(address orderOwner, address receiver, (address,address,uint256,uint256,uint256,address)[] cparams, (address,bytes) callParam) payable returns()
func (_IntentDB *IntentDBTransactorSession) CreateSrcIntentOrderETH(orderOwner common.Address, receiver common.Address, cparams []IIntentDBCreateIntentOrderParam, callParam IIntentDBCallParam) (*types.Transaction, error) {
	return _IntentDB.Contract.CreateSrcIntentOrderETH(&_IntentDB.TransactOpts, orderOwner, receiver, cparams, callParam)
}

// ExecuteDstIntentOrder is a paid mutator transaction binding the contract method 0x30c7ea9a.
//
// Solidity: function executeDstIntentOrder(address orderOwner, address receiver, address vwManager, (address,uint256,uint256,uint256,uint256,uint256)[] underlyingAssetsInfos, (uint256,address,uint256)[] eparams) returns()
func (_IntentDB *IntentDBTransactor) ExecuteDstIntentOrder(opts *bind.TransactOpts, orderOwner common.Address, receiver common.Address, vwManager common.Address, underlyingAssetsInfos []IIntentDBUnderlyingAssetsInfo, eparams []IIntentDBExeIntentOrderParam) (*types.Transaction, error) {
	return _IntentDB.contract.Transact(opts, "executeDstIntentOrder", orderOwner, receiver, vwManager, underlyingAssetsInfos, eparams)
}

// ExecuteDstIntentOrder is a paid mutator transaction binding the contract method 0x30c7ea9a.
//
// Solidity: function executeDstIntentOrder(address orderOwner, address receiver, address vwManager, (address,uint256,uint256,uint256,uint256,uint256)[] underlyingAssetsInfos, (uint256,address,uint256)[] eparams) returns()
func (_IntentDB *IntentDBSession) ExecuteDstIntentOrder(orderOwner common.Address, receiver common.Address, vwManager common.Address, underlyingAssetsInfos []IIntentDBUnderlyingAssetsInfo, eparams []IIntentDBExeIntentOrderParam) (*types.Transaction, error) {
	return _IntentDB.Contract.ExecuteDstIntentOrder(&_IntentDB.TransactOpts, orderOwner, receiver, vwManager, underlyingAssetsInfos, eparams)
}

// ExecuteDstIntentOrder is a paid mutator transaction binding the contract method 0x30c7ea9a.
//
// Solidity: function executeDstIntentOrder(address orderOwner, address receiver, address vwManager, (address,uint256,uint256,uint256,uint256,uint256)[] underlyingAssetsInfos, (uint256,address,uint256)[] eparams) returns()
func (_IntentDB *IntentDBTransactorSession) ExecuteDstIntentOrder(orderOwner common.Address, receiver common.Address, vwManager common.Address, underlyingAssetsInfos []IIntentDBUnderlyingAssetsInfo, eparams []IIntentDBExeIntentOrderParam) (*types.Transaction, error) {
	return _IntentDB.Contract.ExecuteDstIntentOrder(&_IntentDB.TransactOpts, orderOwner, receiver, vwManager, underlyingAssetsInfos, eparams)
}

// IntentDBIntentOrderCancelledIterator is returned from FilterIntentOrderCancelled and is used to iterate over the raw logs and unpacked data for IntentOrderCancelled events raised by the IntentDB contract.
type IntentDBIntentOrderCancelledIterator struct {
	Event *IntentDBIntentOrderCancelled // Event containing the contract specifics and raw log

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
func (it *IntentDBIntentOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentDBIntentOrderCancelled)
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
		it.Event = new(IntentDBIntentOrderCancelled)
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
func (it *IntentDBIntentOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentDBIntentOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentDBIntentOrderCancelled represents a IntentOrderCancelled event raised by the IntentDB contract.
type IntentDBIntentOrderCancelled struct {
	Node           common.Address
	PayOrderId     *big.Int
	CancelOperator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIntentOrderCancelled is a free log retrieval operation binding the contract event 0x6c626ebbe679e89e616c8677ff2b83de2a927396df81bfa70e3fc6b896980e93.
//
// Solidity: event IntentOrderCancelled(address indexed node, uint256 indexed payOrderId, address indexed cancelOperator)
func (_IntentDB *IntentDBFilterer) FilterIntentOrderCancelled(opts *bind.FilterOpts, node []common.Address, payOrderId []*big.Int, cancelOperator []common.Address) (*IntentDBIntentOrderCancelledIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var payOrderIdRule []interface{}
	for _, payOrderIdItem := range payOrderId {
		payOrderIdRule = append(payOrderIdRule, payOrderIdItem)
	}
	var cancelOperatorRule []interface{}
	for _, cancelOperatorItem := range cancelOperator {
		cancelOperatorRule = append(cancelOperatorRule, cancelOperatorItem)
	}

	logs, sub, err := _IntentDB.contract.FilterLogs(opts, "IntentOrderCancelled", nodeRule, payOrderIdRule, cancelOperatorRule)
	if err != nil {
		return nil, err
	}
	return &IntentDBIntentOrderCancelledIterator{contract: _IntentDB.contract, event: "IntentOrderCancelled", logs: logs, sub: sub}, nil
}

// WatchIntentOrderCancelled is a free log subscription operation binding the contract event 0x6c626ebbe679e89e616c8677ff2b83de2a927396df81bfa70e3fc6b896980e93.
//
// Solidity: event IntentOrderCancelled(address indexed node, uint256 indexed payOrderId, address indexed cancelOperator)
func (_IntentDB *IntentDBFilterer) WatchIntentOrderCancelled(opts *bind.WatchOpts, sink chan<- *IntentDBIntentOrderCancelled, node []common.Address, payOrderId []*big.Int, cancelOperator []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var payOrderIdRule []interface{}
	for _, payOrderIdItem := range payOrderId {
		payOrderIdRule = append(payOrderIdRule, payOrderIdItem)
	}
	var cancelOperatorRule []interface{}
	for _, cancelOperatorItem := range cancelOperator {
		cancelOperatorRule = append(cancelOperatorRule, cancelOperatorItem)
	}

	logs, sub, err := _IntentDB.contract.WatchLogs(opts, "IntentOrderCancelled", nodeRule, payOrderIdRule, cancelOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentDBIntentOrderCancelled)
				if err := _IntentDB.contract.UnpackLog(event, "IntentOrderCancelled", log); err != nil {
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

// ParseIntentOrderCancelled is a log parse operation binding the contract event 0x6c626ebbe679e89e616c8677ff2b83de2a927396df81bfa70e3fc6b896980e93.
//
// Solidity: event IntentOrderCancelled(address indexed node, uint256 indexed payOrderId, address indexed cancelOperator)
func (_IntentDB *IntentDBFilterer) ParseIntentOrderCancelled(log types.Log) (*IntentDBIntentOrderCancelled, error) {
	event := new(IntentDBIntentOrderCancelled)
	if err := _IntentDB.contract.UnpackLog(event, "IntentOrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentDBIntentOrderCreatedIterator is returned from FilterIntentOrderCreated and is used to iterate over the raw logs and unpacked data for IntentOrderCreated events raised by the IntentDB contract.
type IntentDBIntentOrderCreatedIterator struct {
	Event *IntentDBIntentOrderCreated // Event containing the contract specifics and raw log

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
func (it *IntentDBIntentOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentDBIntentOrderCreated)
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
		it.Event = new(IntentDBIntentOrderCreated)
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
func (it *IntentDBIntentOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentDBIntentOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentDBIntentOrderCreated represents a IntentOrderCreated event raised by the IntentDB contract.
type IntentDBIntentOrderCreated struct {
	Owner                common.Address
	Node                 common.Address
	PayOrderId           *big.Int
	SrcIntentToken       common.Address
	Sender               common.Address
	Receiver             common.Address
	UnderlyingAssetsInfo IIntentDBUnderlyingAssetsInfo
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIntentOrderCreated is a free log retrieval operation binding the contract event 0x1d7d2e02229ef7737c54f80d91ad28f88e1ca5f4d0d45bd695c8a19d21384c9c.
//
// Solidity: event IntentOrderCreated(address indexed owner, address indexed node, uint256 indexed payOrderId, address srcIntentToken, address sender, address receiver, (address,uint256,uint256,uint256,uint256,uint256) underlyingAssetsInfo)
func (_IntentDB *IntentDBFilterer) FilterIntentOrderCreated(opts *bind.FilterOpts, owner []common.Address, node []common.Address, payOrderId []*big.Int) (*IntentDBIntentOrderCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var payOrderIdRule []interface{}
	for _, payOrderIdItem := range payOrderId {
		payOrderIdRule = append(payOrderIdRule, payOrderIdItem)
	}

	logs, sub, err := _IntentDB.contract.FilterLogs(opts, "IntentOrderCreated", ownerRule, nodeRule, payOrderIdRule)
	if err != nil {
		return nil, err
	}
	return &IntentDBIntentOrderCreatedIterator{contract: _IntentDB.contract, event: "IntentOrderCreated", logs: logs, sub: sub}, nil
}

// WatchIntentOrderCreated is a free log subscription operation binding the contract event 0x1d7d2e02229ef7737c54f80d91ad28f88e1ca5f4d0d45bd695c8a19d21384c9c.
//
// Solidity: event IntentOrderCreated(address indexed owner, address indexed node, uint256 indexed payOrderId, address srcIntentToken, address sender, address receiver, (address,uint256,uint256,uint256,uint256,uint256) underlyingAssetsInfo)
func (_IntentDB *IntentDBFilterer) WatchIntentOrderCreated(opts *bind.WatchOpts, sink chan<- *IntentDBIntentOrderCreated, owner []common.Address, node []common.Address, payOrderId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var payOrderIdRule []interface{}
	for _, payOrderIdItem := range payOrderId {
		payOrderIdRule = append(payOrderIdRule, payOrderIdItem)
	}

	logs, sub, err := _IntentDB.contract.WatchLogs(opts, "IntentOrderCreated", ownerRule, nodeRule, payOrderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentDBIntentOrderCreated)
				if err := _IntentDB.contract.UnpackLog(event, "IntentOrderCreated", log); err != nil {
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

// ParseIntentOrderCreated is a log parse operation binding the contract event 0x1d7d2e02229ef7737c54f80d91ad28f88e1ca5f4d0d45bd695c8a19d21384c9c.
//
// Solidity: event IntentOrderCreated(address indexed owner, address indexed node, uint256 indexed payOrderId, address srcIntentToken, address sender, address receiver, (address,uint256,uint256,uint256,uint256,uint256) underlyingAssetsInfo)
func (_IntentDB *IntentDBFilterer) ParseIntentOrderCreated(log types.Log) (*IntentDBIntentOrderCreated, error) {
	event := new(IntentDBIntentOrderCreated)
	if err := _IntentDB.contract.UnpackLog(event, "IntentOrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentDBIntentOrderExecutedIterator is returned from FilterIntentOrderExecuted and is used to iterate over the raw logs and unpacked data for IntentOrderExecuted events raised by the IntentDB contract.
type IntentDBIntentOrderExecutedIterator struct {
	Event *IntentDBIntentOrderExecuted // Event containing the contract specifics and raw log

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
func (it *IntentDBIntentOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentDBIntentOrderExecuted)
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
		it.Event = new(IntentDBIntentOrderExecuted)
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
func (it *IntentDBIntentOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentDBIntentOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentDBIntentOrderExecuted represents a IntentOrderExecuted event raised by the IntentDB contract.
type IntentDBIntentOrderExecuted struct {
	Executor             common.Address
	PayOrderId           *big.Int
	Owner                common.Address
	Receiver             common.Address
	VwManager            common.Address
	IntentTokenInfo      IIntentDBIntentTokenInfo
	UnderlyingAssetsInfo IIntentDBUnderlyingAssetsInfo
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIntentOrderExecuted is a free log retrieval operation binding the contract event 0x38d3b385322561af25b71bd8e4db2400f5c142d8846b27b1cbdae24b646b29b5.
//
// Solidity: event IntentOrderExecuted(address indexed executor, uint256 indexed payOrderId, address owner, address receiver, address vwManager, (address,uint256,uint256,uint256,uint256) intentTokenInfo, (address,uint256,uint256,uint256,uint256,uint256) underlyingAssetsInfo)
func (_IntentDB *IntentDBFilterer) FilterIntentOrderExecuted(opts *bind.FilterOpts, executor []common.Address, payOrderId []*big.Int) (*IntentDBIntentOrderExecutedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var payOrderIdRule []interface{}
	for _, payOrderIdItem := range payOrderId {
		payOrderIdRule = append(payOrderIdRule, payOrderIdItem)
	}

	logs, sub, err := _IntentDB.contract.FilterLogs(opts, "IntentOrderExecuted", executorRule, payOrderIdRule)
	if err != nil {
		return nil, err
	}
	return &IntentDBIntentOrderExecutedIterator{contract: _IntentDB.contract, event: "IntentOrderExecuted", logs: logs, sub: sub}, nil
}

// WatchIntentOrderExecuted is a free log subscription operation binding the contract event 0x38d3b385322561af25b71bd8e4db2400f5c142d8846b27b1cbdae24b646b29b5.
//
// Solidity: event IntentOrderExecuted(address indexed executor, uint256 indexed payOrderId, address owner, address receiver, address vwManager, (address,uint256,uint256,uint256,uint256) intentTokenInfo, (address,uint256,uint256,uint256,uint256,uint256) underlyingAssetsInfo)
func (_IntentDB *IntentDBFilterer) WatchIntentOrderExecuted(opts *bind.WatchOpts, sink chan<- *IntentDBIntentOrderExecuted, executor []common.Address, payOrderId []*big.Int) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var payOrderIdRule []interface{}
	for _, payOrderIdItem := range payOrderId {
		payOrderIdRule = append(payOrderIdRule, payOrderIdItem)
	}

	logs, sub, err := _IntentDB.contract.WatchLogs(opts, "IntentOrderExecuted", executorRule, payOrderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentDBIntentOrderExecuted)
				if err := _IntentDB.contract.UnpackLog(event, "IntentOrderExecuted", log); err != nil {
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

// ParseIntentOrderExecuted is a log parse operation binding the contract event 0x38d3b385322561af25b71bd8e4db2400f5c142d8846b27b1cbdae24b646b29b5.
//
// Solidity: event IntentOrderExecuted(address indexed executor, uint256 indexed payOrderId, address owner, address receiver, address vwManager, (address,uint256,uint256,uint256,uint256) intentTokenInfo, (address,uint256,uint256,uint256,uint256,uint256) underlyingAssetsInfo)
func (_IntentDB *IntentDBFilterer) ParseIntentOrderExecuted(log types.Log) (*IntentDBIntentOrderExecuted, error) {
	event := new(IntentDBIntentOrderExecuted)
	if err := _IntentDB.contract.UnpackLog(event, "IntentOrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
