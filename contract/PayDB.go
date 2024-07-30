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

// IPayDBCallParam is an auto generated low-level Go binding around an user-defined struct.
type IPayDBCallParam struct {
	Node         common.Address
	NodeCallData []byte
}

// IPayDBCreatePayOrderParam is an auto generated low-level Go binding around an user-defined struct.
type IPayDBCreatePayOrderParam struct {
	AmountIn   *big.Int
	AmountOut  *big.Int
	PayOrderId *big.Int
	BridgeFee  *big.Int
	TokenIn    common.Address
	TokenOut   common.Address
	Node       common.Address
}

// IPayDBExePayOrderParam is an auto generated low-level Go binding around an user-defined struct.
type IPayDBExePayOrderParam struct {
	AmountOut  *big.Int
	TokenOut   common.Address
	PayOrderId *big.Int
}

// IPayDBVwOrderDetail is an auto generated low-level Go binding around an user-defined struct.
type IPayDBVwOrderDetail struct {
	Code     *big.Int
	Service  common.Address
	DataHash [32]byte
}

// IVWManagerVWExecuteParam is an auto generated low-level Go binding around an user-defined struct.
// type IVWManagerVWExecuteParam struct {
// 	Code             *big.Int
// 	GasTokenPrice    *big.Int
// 	PriorityFee      *big.Int
// 	GasLimit         *big.Int
// 	Manager          common.Address
// 	Service          common.Address
// 	GasToken         common.Address
// 	FeeReceiver      common.Address
// 	IsGateway        bool
// 	Data             []byte
// 	ServiceSignature []byte
// 	Proof            [][32]byte
// }

// PayDBMetaData contains all meta data concerning the PayDB contract.
var PayDBMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"IsolateOrderExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cancelOperator\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"wfHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"wfHash\",\"type\":\"bytes32\"}],\"name\":\"OrderExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VWFailedReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vwmanager\",\"type\":\"address\"}],\"name\":\"VWManagerSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"internalType\":\"structIPayDB.CreatePayOrderParam[]\",\"name\":\"cparams\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"workFlowHashs\",\"type\":\"bytes32[]\"}],\"name\":\"cancelOrderETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orderOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"internalType\":\"structIPayDB.CreatePayOrderParam[]\",\"name\":\"cparams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIPayDB.VwOrderDetail\",\"name\":\"vwDetail\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeCallData\",\"type\":\"bytes\"}],\"internalType\":\"structIPayDB.CallParam\",\"name\":\"callParam\",\"type\":\"tuple\"}],\"name\":\"createSrcOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orderOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"internalType\":\"structIPayDB.CreatePayOrderParam[]\",\"name\":\"cparams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIPayDB.VwOrderDetail\",\"name\":\"vwDetail\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"nodeCallData\",\"type\":\"bytes\"}],\"internalType\":\"structIPayDB.CallParam\",\"name\":\"callParam\",\"type\":\"tuple\"}],\"name\":\"createSrcOrderETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dstOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orderOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"}],\"internalType\":\"structIPayDB.ExePayOrderParam[]\",\"name\":\"eparams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isGateway\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"serviceSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIVWManager.VWExecuteParam\",\"name\":\"vwExeParam\",\"type\":\"tuple\"}],\"name\":\"executeDstOrderETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"}],\"internalType\":\"structIPayDB.ExePayOrderParam[]\",\"name\":\"eparams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isGateway\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"serviceSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIVWManager.VWExecuteParam\",\"name\":\"vwExeParam\",\"type\":\"tuple\"}],\"name\":\"executeIsolateOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"}],\"internalType\":\"structIPayDB.ExePayOrderParam[]\",\"name\":\"eparams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isGateway\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"serviceSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIVWManager.VWExecuteParam\",\"name\":\"vwExeParam\",\"type\":\"tuple\"}],\"name\":\"executeIsolateOrderETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"srcOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"status\",\"type\":\"uint96\"},{\"internalType\":\"bytes32\",\"name\":\"orderDataHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orderOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payOrderId\",\"type\":\"uint256\"}],\"internalType\":\"structIPayDB.ExePayOrderParam[]\",\"name\":\"eparams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isGateway\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"serviceSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIVWManager.VWExecuteParam\",\"name\":\"vwExeParam\",\"type\":\"tuple\"}],\"name\":\"tryExecuteDstOrderETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PayDBABI is the input ABI used to generate the binding from.
// Deprecated: Use PayDBMetaData.ABI instead.
var PayDBABI = PayDBMetaData.ABI

// PayDB is an auto generated Go binding around an Ethereum contract.
type PayDB struct {
	PayDBCaller     // Read-only binding to the contract
	PayDBTransactor // Write-only binding to the contract
	PayDBFilterer   // Log filterer for contract events
}

// PayDBCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayDBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayDBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayDBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayDBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayDBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayDBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PayDBSession struct {
	Contract     *PayDB            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayDBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayDBCallerSession struct {
	Contract *PayDBCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PayDBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayDBTransactorSession struct {
	Contract     *PayDBTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayDBRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayDBRaw struct {
	Contract *PayDB // Generic contract binding to access the raw methods on
}

// PayDBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayDBCallerRaw struct {
	Contract *PayDBCaller // Generic read-only contract binding to access the raw methods on
}

// PayDBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayDBTransactorRaw struct {
	Contract *PayDBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayDB creates a new instance of PayDB, bound to a specific deployed contract.
func NewPayDB(address common.Address, backend bind.ContractBackend) (*PayDB, error) {
	contract, err := bindPayDB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PayDB{PayDBCaller: PayDBCaller{contract: contract}, PayDBTransactor: PayDBTransactor{contract: contract}, PayDBFilterer: PayDBFilterer{contract: contract}}, nil
}

// NewPayDBCaller creates a new read-only instance of PayDB, bound to a specific deployed contract.
func NewPayDBCaller(address common.Address, caller bind.ContractCaller) (*PayDBCaller, error) {
	contract, err := bindPayDB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayDBCaller{contract: contract}, nil
}

// NewPayDBTransactor creates a new write-only instance of PayDB, bound to a specific deployed contract.
func NewPayDBTransactor(address common.Address, transactor bind.ContractTransactor) (*PayDBTransactor, error) {
	contract, err := bindPayDB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayDBTransactor{contract: contract}, nil
}

// NewPayDBFilterer creates a new log filterer instance of PayDB, bound to a specific deployed contract.
func NewPayDBFilterer(address common.Address, filterer bind.ContractFilterer) (*PayDBFilterer, error) {
	contract, err := bindPayDB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayDBFilterer{contract: contract}, nil
}

// bindPayDB binds a generic wrapper to an already deployed contract.
func bindPayDB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PayDBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayDB *PayDBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PayDB.Contract.PayDBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayDB *PayDBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayDB.Contract.PayDBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayDB *PayDBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayDB.Contract.PayDBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayDB *PayDBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PayDB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayDB *PayDBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayDB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayDB *PayDBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayDB.Contract.contract.Transact(opts, method, params...)
}

// DstOrder is a free data retrieval call binding the contract method 0xf45afa36.
//
// Solidity: function dstOrder(address , uint256 ) view returns(bytes32)
func (_PayDB *PayDBCaller) DstOrder(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PayDB.contract.Call(opts, &out, "dstOrder", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DstOrder is a free data retrieval call binding the contract method 0xf45afa36.
//
// Solidity: function dstOrder(address , uint256 ) view returns(bytes32)
func (_PayDB *PayDBSession) DstOrder(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _PayDB.Contract.DstOrder(&_PayDB.CallOpts, arg0, arg1)
}

// DstOrder is a free data retrieval call binding the contract method 0xf45afa36.
//
// Solidity: function dstOrder(address , uint256 ) view returns(bytes32)
func (_PayDB *PayDBCallerSession) DstOrder(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _PayDB.Contract.DstOrder(&_PayDB.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PayDB *PayDBCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PayDB.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PayDB *PayDBSession) Owner() (common.Address, error) {
	return _PayDB.Contract.Owner(&_PayDB.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PayDB *PayDBCallerSession) Owner() (common.Address, error) {
	return _PayDB.Contract.Owner(&_PayDB.CallOpts)
}

// SrcOrder is a free data retrieval call binding the contract method 0x672f0ea2.
//
// Solidity: function srcOrder(uint256 ) view returns(address node, uint96 status, bytes32 orderDataHash)
func (_PayDB *PayDBCaller) SrcOrder(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Node          common.Address
	Status        *big.Int
	OrderDataHash [32]byte
}, error) {
	var out []interface{}
	err := _PayDB.contract.Call(opts, &out, "srcOrder", arg0)

	outstruct := new(struct {
		Node          common.Address
		Status        *big.Int
		OrderDataHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Node = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OrderDataHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SrcOrder is a free data retrieval call binding the contract method 0x672f0ea2.
//
// Solidity: function srcOrder(uint256 ) view returns(address node, uint96 status, bytes32 orderDataHash)
func (_PayDB *PayDBSession) SrcOrder(arg0 *big.Int) (struct {
	Node          common.Address
	Status        *big.Int
	OrderDataHash [32]byte
}, error) {
	return _PayDB.Contract.SrcOrder(&_PayDB.CallOpts, arg0)
}

// SrcOrder is a free data retrieval call binding the contract method 0x672f0ea2.
//
// Solidity: function srcOrder(uint256 ) view returns(address node, uint96 status, bytes32 orderDataHash)
func (_PayDB *PayDBCallerSession) SrcOrder(arg0 *big.Int) (struct {
	Node          common.Address
	Status        *big.Int
	OrderDataHash [32]byte
}, error) {
	return _PayDB.Contract.SrcOrder(&_PayDB.CallOpts, arg0)
}

// CancelOrderETH is a paid mutator transaction binding the contract method 0x712293d3.
//
// Solidity: function cancelOrderETH(address sender, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, bytes32[] workFlowHashs) payable returns()
func (_PayDB *PayDBTransactor) CancelOrderETH(opts *bind.TransactOpts, sender common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, workFlowHashs [][32]byte) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "cancelOrderETH", sender, receiver, cparams, workFlowHashs)
}

// CancelOrderETH is a paid mutator transaction binding the contract method 0x712293d3.
//
// Solidity: function cancelOrderETH(address sender, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, bytes32[] workFlowHashs) payable returns()
func (_PayDB *PayDBSession) CancelOrderETH(sender common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, workFlowHashs [][32]byte) (*types.Transaction, error) {
	return _PayDB.Contract.CancelOrderETH(&_PayDB.TransactOpts, sender, receiver, cparams, workFlowHashs)
}

// CancelOrderETH is a paid mutator transaction binding the contract method 0x712293d3.
//
// Solidity: function cancelOrderETH(address sender, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, bytes32[] workFlowHashs) payable returns()
func (_PayDB *PayDBTransactorSession) CancelOrderETH(sender common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, workFlowHashs [][32]byte) (*types.Transaction, error) {
	return _PayDB.Contract.CancelOrderETH(&_PayDB.TransactOpts, sender, receiver, cparams, workFlowHashs)
}

// CreateSrcOrder is a paid mutator transaction binding the contract method 0x12b28b66.
//
// Solidity: function createSrcOrder(address orderOwner, address wallet, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, (uint256,address,bytes32) vwDetail, (address,bytes) callParam) returns()
func (_PayDB *PayDBTransactor) CreateSrcOrder(opts *bind.TransactOpts, orderOwner common.Address, wallet common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, vwDetail IPayDBVwOrderDetail, callParam IPayDBCallParam) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "createSrcOrder", orderOwner, wallet, receiver, cparams, vwDetail, callParam)
}

// CreateSrcOrder is a paid mutator transaction binding the contract method 0x12b28b66.
//
// Solidity: function createSrcOrder(address orderOwner, address wallet, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, (uint256,address,bytes32) vwDetail, (address,bytes) callParam) returns()
func (_PayDB *PayDBSession) CreateSrcOrder(orderOwner common.Address, wallet common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, vwDetail IPayDBVwOrderDetail, callParam IPayDBCallParam) (*types.Transaction, error) {
	return _PayDB.Contract.CreateSrcOrder(&_PayDB.TransactOpts, orderOwner, wallet, receiver, cparams, vwDetail, callParam)
}

// CreateSrcOrder is a paid mutator transaction binding the contract method 0x12b28b66.
//
// Solidity: function createSrcOrder(address orderOwner, address wallet, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, (uint256,address,bytes32) vwDetail, (address,bytes) callParam) returns()
func (_PayDB *PayDBTransactorSession) CreateSrcOrder(orderOwner common.Address, wallet common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, vwDetail IPayDBVwOrderDetail, callParam IPayDBCallParam) (*types.Transaction, error) {
	return _PayDB.Contract.CreateSrcOrder(&_PayDB.TransactOpts, orderOwner, wallet, receiver, cparams, vwDetail, callParam)
}

// CreateSrcOrderETH is a paid mutator transaction binding the contract method 0x3a1e9e4c.
//
// Solidity: function createSrcOrderETH(address orderOwner, address wallet, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, (uint256,address,bytes32) vwDetail, (address,bytes) callParam) payable returns()
func (_PayDB *PayDBTransactor) CreateSrcOrderETH(opts *bind.TransactOpts, orderOwner common.Address, wallet common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, vwDetail IPayDBVwOrderDetail, callParam IPayDBCallParam) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "createSrcOrderETH", orderOwner, wallet, receiver, cparams, vwDetail, callParam)
}

// CreateSrcOrderETH is a paid mutator transaction binding the contract method 0x3a1e9e4c.
//
// Solidity: function createSrcOrderETH(address orderOwner, address wallet, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, (uint256,address,bytes32) vwDetail, (address,bytes) callParam) payable returns()
func (_PayDB *PayDBSession) CreateSrcOrderETH(orderOwner common.Address, wallet common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, vwDetail IPayDBVwOrderDetail, callParam IPayDBCallParam) (*types.Transaction, error) {
	return _PayDB.Contract.CreateSrcOrderETH(&_PayDB.TransactOpts, orderOwner, wallet, receiver, cparams, vwDetail, callParam)
}

// CreateSrcOrderETH is a paid mutator transaction binding the contract method 0x3a1e9e4c.
//
// Solidity: function createSrcOrderETH(address orderOwner, address wallet, address receiver, (uint256,uint256,uint256,uint256,address,address,address)[] cparams, (uint256,address,bytes32) vwDetail, (address,bytes) callParam) payable returns()
func (_PayDB *PayDBTransactorSession) CreateSrcOrderETH(orderOwner common.Address, wallet common.Address, receiver common.Address, cparams []IPayDBCreatePayOrderParam, vwDetail IPayDBVwOrderDetail, callParam IPayDBCallParam) (*types.Transaction, error) {
	return _PayDB.Contract.CreateSrcOrderETH(&_PayDB.TransactOpts, orderOwner, wallet, receiver, cparams, vwDetail, callParam)
}

// ExecuteDstOrderETH is a paid mutator transaction binding the contract method 0x91d22b18.
//
// Solidity: function executeDstOrderETH(address orderOwner, address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBTransactor) ExecuteDstOrderETH(opts *bind.TransactOpts, orderOwner common.Address, receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "executeDstOrderETH", orderOwner, receiver, wallet, eparams, vwExeParam)
}

// ExecuteDstOrderETH is a paid mutator transaction binding the contract method 0x91d22b18.
//
// Solidity: function executeDstOrderETH(address orderOwner, address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBSession) ExecuteDstOrderETH(orderOwner common.Address, receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.Contract.ExecuteDstOrderETH(&_PayDB.TransactOpts, orderOwner, receiver, wallet, eparams, vwExeParam)
}

// ExecuteDstOrderETH is a paid mutator transaction binding the contract method 0x91d22b18.
//
// Solidity: function executeDstOrderETH(address orderOwner, address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBTransactorSession) ExecuteDstOrderETH(orderOwner common.Address, receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.Contract.ExecuteDstOrderETH(&_PayDB.TransactOpts, orderOwner, receiver, wallet, eparams, vwExeParam)
}

// ExecuteIsolateOrder is a paid mutator transaction binding the contract method 0x6264d7e0.
//
// Solidity: function executeIsolateOrder(address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) returns()
func (_PayDB *PayDBTransactor) ExecuteIsolateOrder(opts *bind.TransactOpts, receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "executeIsolateOrder", receiver, wallet, eparams, vwExeParam)
}

// ExecuteIsolateOrder is a paid mutator transaction binding the contract method 0x6264d7e0.
//
// Solidity: function executeIsolateOrder(address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) returns()
func (_PayDB *PayDBSession) ExecuteIsolateOrder(receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.Contract.ExecuteIsolateOrder(&_PayDB.TransactOpts, receiver, wallet, eparams, vwExeParam)
}

// ExecuteIsolateOrder is a paid mutator transaction binding the contract method 0x6264d7e0.
//
// Solidity: function executeIsolateOrder(address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) returns()
func (_PayDB *PayDBTransactorSession) ExecuteIsolateOrder(receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.Contract.ExecuteIsolateOrder(&_PayDB.TransactOpts, receiver, wallet, eparams, vwExeParam)
}

// ExecuteIsolateOrderETH is a paid mutator transaction binding the contract method 0xdc7fd8fe.
//
// Solidity: function executeIsolateOrderETH(address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBTransactor) ExecuteIsolateOrderETH(opts *bind.TransactOpts, receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "executeIsolateOrderETH", receiver, wallet, eparams, vwExeParam)
}

// ExecuteIsolateOrderETH is a paid mutator transaction binding the contract method 0xdc7fd8fe.
//
// Solidity: function executeIsolateOrderETH(address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBSession) ExecuteIsolateOrderETH(receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.Contract.ExecuteIsolateOrderETH(&_PayDB.TransactOpts, receiver, wallet, eparams, vwExeParam)
}

// ExecuteIsolateOrderETH is a paid mutator transaction binding the contract method 0xdc7fd8fe.
//
// Solidity: function executeIsolateOrderETH(address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBTransactorSession) ExecuteIsolateOrderETH(receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.Contract.ExecuteIsolateOrderETH(&_PayDB.TransactOpts, receiver, wallet, eparams, vwExeParam)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PayDB *PayDBTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PayDB *PayDBSession) RenounceOwnership() (*types.Transaction, error) {
	return _PayDB.Contract.RenounceOwnership(&_PayDB.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PayDB *PayDBTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PayDB.Contract.RenounceOwnership(&_PayDB.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PayDB *PayDBTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PayDB *PayDBSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PayDB.Contract.TransferOwnership(&_PayDB.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PayDB *PayDBTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PayDB.Contract.TransferOwnership(&_PayDB.TransactOpts, newOwner)
}

// TryExecuteDstOrderETH is a paid mutator transaction binding the contract method 0x5e8952b6.
//
// Solidity: function tryExecuteDstOrderETH(address orderOwner, address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBTransactor) TryExecuteDstOrderETH(opts *bind.TransactOpts, orderOwner common.Address, receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.contract.Transact(opts, "tryExecuteDstOrderETH", orderOwner, receiver, wallet, eparams, vwExeParam)
}

// TryExecuteDstOrderETH is a paid mutator transaction binding the contract method 0x5e8952b6.
//
// Solidity: function tryExecuteDstOrderETH(address orderOwner, address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBSession) TryExecuteDstOrderETH(orderOwner common.Address, receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.Contract.TryExecuteDstOrderETH(&_PayDB.TransactOpts, orderOwner, receiver, wallet, eparams, vwExeParam)
}

// TryExecuteDstOrderETH is a paid mutator transaction binding the contract method 0x5e8952b6.
//
// Solidity: function tryExecuteDstOrderETH(address orderOwner, address receiver, address wallet, (uint256,address,uint256)[] eparams, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vwExeParam) payable returns()
func (_PayDB *PayDBTransactorSession) TryExecuteDstOrderETH(orderOwner common.Address, receiver common.Address, wallet common.Address, eparams []IPayDBExePayOrderParam, vwExeParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _PayDB.Contract.TryExecuteDstOrderETH(&_PayDB.TransactOpts, orderOwner, receiver, wallet, eparams, vwExeParam)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PayDB *PayDBTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PayDB.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PayDB *PayDBSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PayDB.Contract.Fallback(&_PayDB.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PayDB *PayDBTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PayDB.Contract.Fallback(&_PayDB.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PayDB *PayDBTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayDB.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PayDB *PayDBSession) Receive() (*types.Transaction, error) {
	return _PayDB.Contract.Receive(&_PayDB.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PayDB *PayDBTransactorSession) Receive() (*types.Transaction, error) {
	return _PayDB.Contract.Receive(&_PayDB.TransactOpts)
}

// PayDBIsolateOrderExecutedIterator is returned from FilterIsolateOrderExecuted and is used to iterate over the raw logs and unpacked data for IsolateOrderExecuted events raised by the PayDB contract.
type PayDBIsolateOrderExecutedIterator struct {
	Event *PayDBIsolateOrderExecuted // Event containing the contract specifics and raw log

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
func (it *PayDBIsolateOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayDBIsolateOrderExecuted)
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
		it.Event = new(PayDBIsolateOrderExecuted)
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
func (it *PayDBIsolateOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayDBIsolateOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayDBIsolateOrderExecuted represents a IsolateOrderExecuted event raised by the PayDB contract.
type PayDBIsolateOrderExecuted struct {
	Executor  common.Address
	Wallet    common.Address
	Receiver  common.Address
	AmountOut *big.Int
	TokenOut  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIsolateOrderExecuted is a free log retrieval operation binding the contract event 0xd9f6383b7dcb1b7b8744b1e10d5a9164ff3a188f6c99ef24848e9987a93237a7.
//
// Solidity: event IsolateOrderExecuted(address indexed executor, address wallet, address receiver, uint256 amountOut, address tokenOut)
func (_PayDB *PayDBFilterer) FilterIsolateOrderExecuted(opts *bind.FilterOpts, executor []common.Address) (*PayDBIsolateOrderExecutedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _PayDB.contract.FilterLogs(opts, "IsolateOrderExecuted", executorRule)
	if err != nil {
		return nil, err
	}
	return &PayDBIsolateOrderExecutedIterator{contract: _PayDB.contract, event: "IsolateOrderExecuted", logs: logs, sub: sub}, nil
}

// WatchIsolateOrderExecuted is a free log subscription operation binding the contract event 0xd9f6383b7dcb1b7b8744b1e10d5a9164ff3a188f6c99ef24848e9987a93237a7.
//
// Solidity: event IsolateOrderExecuted(address indexed executor, address wallet, address receiver, uint256 amountOut, address tokenOut)
func (_PayDB *PayDBFilterer) WatchIsolateOrderExecuted(opts *bind.WatchOpts, sink chan<- *PayDBIsolateOrderExecuted, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _PayDB.contract.WatchLogs(opts, "IsolateOrderExecuted", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayDBIsolateOrderExecuted)
				if err := _PayDB.contract.UnpackLog(event, "IsolateOrderExecuted", log); err != nil {
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

// ParseIsolateOrderExecuted is a log parse operation binding the contract event 0xd9f6383b7dcb1b7b8744b1e10d5a9164ff3a188f6c99ef24848e9987a93237a7.
//
// Solidity: event IsolateOrderExecuted(address indexed executor, address wallet, address receiver, uint256 amountOut, address tokenOut)
func (_PayDB *PayDBFilterer) ParseIsolateOrderExecuted(log types.Log) (*PayDBIsolateOrderExecuted, error) {
	event := new(PayDBIsolateOrderExecuted)
	if err := _PayDB.contract.UnpackLog(event, "IsolateOrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayDBOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the PayDB contract.
type PayDBOrderCancelledIterator struct {
	Event *PayDBOrderCancelled // Event containing the contract specifics and raw log

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
func (it *PayDBOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayDBOrderCancelled)
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
		it.Event = new(PayDBOrderCancelled)
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
func (it *PayDBOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayDBOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayDBOrderCancelled represents a OrderCancelled event raised by the PayDB contract.
type PayDBOrderCancelled struct {
	Node           common.Address
	PayOrderId     *big.Int
	CancelOperator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xd26d0eaf100efe01fc00473fcd7c0d820f51e58b0ff954d6cdb7d984e000fb1c.
//
// Solidity: event OrderCancelled(address indexed node, uint256 indexed payOrderId, address indexed cancelOperator)
func (_PayDB *PayDBFilterer) FilterOrderCancelled(opts *bind.FilterOpts, node []common.Address, payOrderId []*big.Int, cancelOperator []common.Address) (*PayDBOrderCancelledIterator, error) {

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

	logs, sub, err := _PayDB.contract.FilterLogs(opts, "OrderCancelled", nodeRule, payOrderIdRule, cancelOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PayDBOrderCancelledIterator{contract: _PayDB.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xd26d0eaf100efe01fc00473fcd7c0d820f51e58b0ff954d6cdb7d984e000fb1c.
//
// Solidity: event OrderCancelled(address indexed node, uint256 indexed payOrderId, address indexed cancelOperator)
func (_PayDB *PayDBFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *PayDBOrderCancelled, node []common.Address, payOrderId []*big.Int, cancelOperator []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PayDB.contract.WatchLogs(opts, "OrderCancelled", nodeRule, payOrderIdRule, cancelOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayDBOrderCancelled)
				if err := _PayDB.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0xd26d0eaf100efe01fc00473fcd7c0d820f51e58b0ff954d6cdb7d984e000fb1c.
//
// Solidity: event OrderCancelled(address indexed node, uint256 indexed payOrderId, address indexed cancelOperator)
func (_PayDB *PayDBFilterer) ParseOrderCancelled(log types.Log) (*PayDBOrderCancelled, error) {
	event := new(PayDBOrderCancelled)
	if err := _PayDB.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayDBOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the PayDB contract.
type PayDBOrderCreatedIterator struct {
	Event *PayDBOrderCreated // Event containing the contract specifics and raw log

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
func (it *PayDBOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayDBOrderCreated)
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
		it.Event = new(PayDBOrderCreated)
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
func (it *PayDBOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayDBOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayDBOrderCreated represents a OrderCreated event raised by the PayDB contract.
type PayDBOrderCreated struct {
	Owner      common.Address
	Node       common.Address
	PayOrderId *big.Int
	Sender     common.Address
	Receiver   common.Address
	AmountIn   *big.Int
	TokenIn    common.Address
	AmountOut  *big.Int
	TokenOut   common.Address
	Code       *big.Int
	WfHash     [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0xd5e17b22fcc50d4f6449df5335edf7467e5a7c9bc8dbc79288e93a25ca6a90c5.
//
// Solidity: event OrderCreated(address indexed owner, address indexed node, uint256 indexed payOrderId, address sender, address receiver, uint256 amountIn, address tokenIn, uint256 amountOut, address tokenOut, uint256 code, bytes32 wfHash)
func (_PayDB *PayDBFilterer) FilterOrderCreated(opts *bind.FilterOpts, owner []common.Address, node []common.Address, payOrderId []*big.Int) (*PayDBOrderCreatedIterator, error) {

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

	logs, sub, err := _PayDB.contract.FilterLogs(opts, "OrderCreated", ownerRule, nodeRule, payOrderIdRule)
	if err != nil {
		return nil, err
	}
	return &PayDBOrderCreatedIterator{contract: _PayDB.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0xd5e17b22fcc50d4f6449df5335edf7467e5a7c9bc8dbc79288e93a25ca6a90c5.
//
// Solidity: event OrderCreated(address indexed owner, address indexed node, uint256 indexed payOrderId, address sender, address receiver, uint256 amountIn, address tokenIn, uint256 amountOut, address tokenOut, uint256 code, bytes32 wfHash)
func (_PayDB *PayDBFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *PayDBOrderCreated, owner []common.Address, node []common.Address, payOrderId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PayDB.contract.WatchLogs(opts, "OrderCreated", ownerRule, nodeRule, payOrderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayDBOrderCreated)
				if err := _PayDB.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0xd5e17b22fcc50d4f6449df5335edf7467e5a7c9bc8dbc79288e93a25ca6a90c5.
//
// Solidity: event OrderCreated(address indexed owner, address indexed node, uint256 indexed payOrderId, address sender, address receiver, uint256 amountIn, address tokenIn, uint256 amountOut, address tokenOut, uint256 code, bytes32 wfHash)
func (_PayDB *PayDBFilterer) ParseOrderCreated(log types.Log) (*PayDBOrderCreated, error) {
	event := new(PayDBOrderCreated)
	if err := _PayDB.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayDBOrderExecutedIterator is returned from FilterOrderExecuted and is used to iterate over the raw logs and unpacked data for OrderExecuted events raised by the PayDB contract.
type PayDBOrderExecutedIterator struct {
	Event *PayDBOrderExecuted // Event containing the contract specifics and raw log

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
func (it *PayDBOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayDBOrderExecuted)
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
		it.Event = new(PayDBOrderExecuted)
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
func (it *PayDBOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayDBOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayDBOrderExecuted represents a OrderExecuted event raised by the PayDB contract.
type PayDBOrderExecuted struct {
	Executor   common.Address
	PayOrderId *big.Int
	Owner      common.Address
	Receiver   common.Address
	AmountOut  *big.Int
	TokenOut   common.Address
	Code       *big.Int
	WfHash     [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderExecuted is a free log retrieval operation binding the contract event 0x035725451e5bd2b60362063dbcdea18faf7d366572f9ab6b817a73a5da53b6a2.
//
// Solidity: event OrderExecuted(address indexed executor, uint256 indexed payOrderId, address owner, address receiver, uint256 amountOut, address tokenOut, uint256 code, bytes32 wfHash)
func (_PayDB *PayDBFilterer) FilterOrderExecuted(opts *bind.FilterOpts, executor []common.Address, payOrderId []*big.Int) (*PayDBOrderExecutedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var payOrderIdRule []interface{}
	for _, payOrderIdItem := range payOrderId {
		payOrderIdRule = append(payOrderIdRule, payOrderIdItem)
	}

	logs, sub, err := _PayDB.contract.FilterLogs(opts, "OrderExecuted", executorRule, payOrderIdRule)
	if err != nil {
		return nil, err
	}
	return &PayDBOrderExecutedIterator{contract: _PayDB.contract, event: "OrderExecuted", logs: logs, sub: sub}, nil
}

// WatchOrderExecuted is a free log subscription operation binding the contract event 0x035725451e5bd2b60362063dbcdea18faf7d366572f9ab6b817a73a5da53b6a2.
//
// Solidity: event OrderExecuted(address indexed executor, uint256 indexed payOrderId, address owner, address receiver, uint256 amountOut, address tokenOut, uint256 code, bytes32 wfHash)
func (_PayDB *PayDBFilterer) WatchOrderExecuted(opts *bind.WatchOpts, sink chan<- *PayDBOrderExecuted, executor []common.Address, payOrderId []*big.Int) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var payOrderIdRule []interface{}
	for _, payOrderIdItem := range payOrderId {
		payOrderIdRule = append(payOrderIdRule, payOrderIdItem)
	}

	logs, sub, err := _PayDB.contract.WatchLogs(opts, "OrderExecuted", executorRule, payOrderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayDBOrderExecuted)
				if err := _PayDB.contract.UnpackLog(event, "OrderExecuted", log); err != nil {
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

// ParseOrderExecuted is a log parse operation binding the contract event 0x035725451e5bd2b60362063dbcdea18faf7d366572f9ab6b817a73a5da53b6a2.
//
// Solidity: event OrderExecuted(address indexed executor, uint256 indexed payOrderId, address owner, address receiver, uint256 amountOut, address tokenOut, uint256 code, bytes32 wfHash)
func (_PayDB *PayDBFilterer) ParseOrderExecuted(log types.Log) (*PayDBOrderExecuted, error) {
	event := new(PayDBOrderExecuted)
	if err := _PayDB.contract.UnpackLog(event, "OrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayDBOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PayDB contract.
type PayDBOwnershipTransferredIterator struct {
	Event *PayDBOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PayDBOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayDBOwnershipTransferred)
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
		it.Event = new(PayDBOwnershipTransferred)
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
func (it *PayDBOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayDBOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayDBOwnershipTransferred represents a OwnershipTransferred event raised by the PayDB contract.
type PayDBOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PayDB *PayDBFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PayDBOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PayDB.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PayDBOwnershipTransferredIterator{contract: _PayDB.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PayDB *PayDBFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PayDBOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PayDB.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayDBOwnershipTransferred)
				if err := _PayDB.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PayDB *PayDBFilterer) ParseOwnershipTransferred(log types.Log) (*PayDBOwnershipTransferred, error) {
	event := new(PayDBOwnershipTransferred)
	if err := _PayDB.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayDBVWFailedReasonIterator is returned from FilterVWFailedReason and is used to iterate over the raw logs and unpacked data for VWFailedReason events raised by the PayDB contract.
type PayDBVWFailedReasonIterator struct {
	Event *PayDBVWFailedReason // Event containing the contract specifics and raw log

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
func (it *PayDBVWFailedReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayDBVWFailedReason)
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
		it.Event = new(PayDBVWFailedReason)
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
func (it *PayDBVWFailedReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayDBVWFailedReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayDBVWFailedReason represents a VWFailedReason event raised by the PayDB contract.
type PayDBVWFailedReason struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVWFailedReason is a free log retrieval operation binding the contract event 0x3d3068871f601e731d9e5810e5d2947adee0ab0e84672c6888e1aa6696c0fd82.
//
// Solidity: event VWFailedReason(string reason)
func (_PayDB *PayDBFilterer) FilterVWFailedReason(opts *bind.FilterOpts) (*PayDBVWFailedReasonIterator, error) {

	logs, sub, err := _PayDB.contract.FilterLogs(opts, "VWFailedReason")
	if err != nil {
		return nil, err
	}
	return &PayDBVWFailedReasonIterator{contract: _PayDB.contract, event: "VWFailedReason", logs: logs, sub: sub}, nil
}

// WatchVWFailedReason is a free log subscription operation binding the contract event 0x3d3068871f601e731d9e5810e5d2947adee0ab0e84672c6888e1aa6696c0fd82.
//
// Solidity: event VWFailedReason(string reason)
func (_PayDB *PayDBFilterer) WatchVWFailedReason(opts *bind.WatchOpts, sink chan<- *PayDBVWFailedReason) (event.Subscription, error) {

	logs, sub, err := _PayDB.contract.WatchLogs(opts, "VWFailedReason")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayDBVWFailedReason)
				if err := _PayDB.contract.UnpackLog(event, "VWFailedReason", log); err != nil {
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

// ParseVWFailedReason is a log parse operation binding the contract event 0x3d3068871f601e731d9e5810e5d2947adee0ab0e84672c6888e1aa6696c0fd82.
//
// Solidity: event VWFailedReason(string reason)
func (_PayDB *PayDBFilterer) ParseVWFailedReason(log types.Log) (*PayDBVWFailedReason, error) {
	event := new(PayDBVWFailedReason)
	if err := _PayDB.contract.UnpackLog(event, "VWFailedReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayDBVWManagerSetIterator is returned from FilterVWManagerSet and is used to iterate over the raw logs and unpacked data for VWManagerSet events raised by the PayDB contract.
type PayDBVWManagerSetIterator struct {
	Event *PayDBVWManagerSet // Event containing the contract specifics and raw log

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
func (it *PayDBVWManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayDBVWManagerSet)
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
		it.Event = new(PayDBVWManagerSet)
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
func (it *PayDBVWManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayDBVWManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayDBVWManagerSet represents a VWManagerSet event raised by the PayDB contract.
type PayDBVWManagerSet struct {
	Vwmanager common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVWManagerSet is a free log retrieval operation binding the contract event 0xde42365cd1f410c30b2fed1f7a9a5bf9c53cd60f056877f9872a754a0b6021f2.
//
// Solidity: event VWManagerSet(address indexed vwmanager)
func (_PayDB *PayDBFilterer) FilterVWManagerSet(opts *bind.FilterOpts, vwmanager []common.Address) (*PayDBVWManagerSetIterator, error) {

	var vwmanagerRule []interface{}
	for _, vwmanagerItem := range vwmanager {
		vwmanagerRule = append(vwmanagerRule, vwmanagerItem)
	}

	logs, sub, err := _PayDB.contract.FilterLogs(opts, "VWManagerSet", vwmanagerRule)
	if err != nil {
		return nil, err
	}
	return &PayDBVWManagerSetIterator{contract: _PayDB.contract, event: "VWManagerSet", logs: logs, sub: sub}, nil
}

// WatchVWManagerSet is a free log subscription operation binding the contract event 0xde42365cd1f410c30b2fed1f7a9a5bf9c53cd60f056877f9872a754a0b6021f2.
//
// Solidity: event VWManagerSet(address indexed vwmanager)
func (_PayDB *PayDBFilterer) WatchVWManagerSet(opts *bind.WatchOpts, sink chan<- *PayDBVWManagerSet, vwmanager []common.Address) (event.Subscription, error) {

	var vwmanagerRule []interface{}
	for _, vwmanagerItem := range vwmanager {
		vwmanagerRule = append(vwmanagerRule, vwmanagerItem)
	}

	logs, sub, err := _PayDB.contract.WatchLogs(opts, "VWManagerSet", vwmanagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayDBVWManagerSet)
				if err := _PayDB.contract.UnpackLog(event, "VWManagerSet", log); err != nil {
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

// ParseVWManagerSet is a log parse operation binding the contract event 0xde42365cd1f410c30b2fed1f7a9a5bf9c53cd60f056877f9872a754a0b6021f2.
//
// Solidity: event VWManagerSet(address indexed vwmanager)
func (_PayDB *PayDBFilterer) ParseVWManagerSet(log types.Log) (*PayDBVWManagerSet, error) {
	event := new(PayDBVWManagerSet)
	if err := _PayDB.contract.UnpackLog(event, "VWManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
