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

// // IVWManagerVWExecuteParam is an auto generated low-level Go binding around an user-defined struct.
type IVWManagerVWExecuteParam struct {
	Code             *big.Int
	GasTokenPrice    *big.Int
	PriorityFee      *big.Int
	GasLimit         *big.Int
	Manager          common.Address
	Service          common.Address
	GasToken         common.Address
	FeeReceiver      common.Address
	IsGateway        bool
	Data             []byte
	ServiceSignature []byte
	Proof            [][32]byte
}

// VWManagerServiceFeeParam is an auto generated low-level Go binding around an user-defined struct.
type VWManagerServiceFeeParam struct {
	GasToken      common.Address
	GasTokenPrice *big.Int
	PriorityFee   *big.Int
	GasLimit      *big.Int
	FeeReceiver   common.Address
}

// VWManagerMetaData contains all meta data concerning the VWManager contract.
var VWManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"}],\"name\":\"DomainSeparatorCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifyContract\",\"type\":\"address\"}],\"name\":\"DomainSeparatorConfiged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reseter\",\"type\":\"address\"}],\"name\":\"ResetterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"protocolFeeOpened\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"feeProportion\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"TxCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resCode\",\"type\":\"uint256\"}],\"name\":\"TxExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"VWCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"VWOwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resetter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"internalType\":\"structVWManagerService.FeeParam\",\"name\":\"fParam\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"approveResetter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedResetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"codeToCancel\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"internalType\":\"structVWManagerService.FeeParam\",\"name\":\"fParam\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"cancelTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"internalType\":\"structVWManagerService.FeeParam\",\"name\":\"fParam\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_protocolFeeOpened\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_feeVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeProportion\",\"type\":\"uint256\"}],\"name\":\"configFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_support\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_verifyingContract\",\"type\":\"address\"}],\"name\":\"configSrcChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isGateway\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"serviceSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIVWManager.VWExecuteParam\",\"name\":\"vweParam\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"resCode\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeProportion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"internalType\":\"structVWManagerService.FeeParam\",\"name\":\"fParam\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"result\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifyEIP1271Signature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resCode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isGateway\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"serviceSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIVWManager.VWExecuteParam\",\"name\":\"vweParam\",\"type\":\"tuple\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatileService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VWManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use VWManagerMetaData.ABI instead.
var VWManagerABI = VWManagerMetaData.ABI

// VWManager is an auto generated Go binding around an Ethereum contract.
type VWManager struct {
	VWManagerCaller     // Read-only binding to the contract
	VWManagerTransactor // Write-only binding to the contract
	VWManagerFilterer   // Log filterer for contract events
}

// VWManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VWManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VWManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VWManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VWManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VWManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VWManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VWManagerSession struct {
	Contract     *VWManager        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VWManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VWManagerCallerSession struct {
	Contract *VWManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VWManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VWManagerTransactorSession struct {
	Contract     *VWManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VWManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VWManagerRaw struct {
	Contract *VWManager // Generic contract binding to access the raw methods on
}

// VWManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VWManagerCallerRaw struct {
	Contract *VWManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VWManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VWManagerTransactorRaw struct {
	Contract *VWManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVWManager creates a new instance of VWManager, bound to a specific deployed contract.
func NewVWManager(address common.Address, backend bind.ContractBackend) (*VWManager, error) {
	contract, err := bindVWManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VWManager{VWManagerCaller: VWManagerCaller{contract: contract}, VWManagerTransactor: VWManagerTransactor{contract: contract}, VWManagerFilterer: VWManagerFilterer{contract: contract}}, nil
}

// NewVWManagerCaller creates a new read-only instance of VWManager, bound to a specific deployed contract.
func NewVWManagerCaller(address common.Address, caller bind.ContractCaller) (*VWManagerCaller, error) {
	contract, err := bindVWManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VWManagerCaller{contract: contract}, nil
}

// NewVWManagerTransactor creates a new write-only instance of VWManager, bound to a specific deployed contract.
func NewVWManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VWManagerTransactor, error) {
	contract, err := bindVWManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VWManagerTransactor{contract: contract}, nil
}

// NewVWManagerFilterer creates a new log filterer instance of VWManager, bound to a specific deployed contract.
func NewVWManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VWManagerFilterer, error) {
	contract, err := bindVWManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VWManagerFilterer{contract: contract}, nil
}

// bindVWManager binds a generic wrapper to an already deployed contract.
func bindVWManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VWManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VWManager *VWManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VWManager.Contract.VWManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VWManager *VWManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VWManager.Contract.VWManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VWManager *VWManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VWManager.Contract.VWManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VWManager *VWManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VWManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VWManager *VWManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VWManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VWManager *VWManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VWManager.Contract.contract.Transact(opts, method, params...)
}

// ApprovedResetter is a free data retrieval call binding the contract method 0x9b4f5249.
//
// Solidity: function approvedResetter(address ) view returns(address)
func (_VWManager *VWManagerCaller) ApprovedResetter(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "approvedResetter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApprovedResetter is a free data retrieval call binding the contract method 0x9b4f5249.
//
// Solidity: function approvedResetter(address ) view returns(address)
func (_VWManager *VWManagerSession) ApprovedResetter(arg0 common.Address) (common.Address, error) {
	return _VWManager.Contract.ApprovedResetter(&_VWManager.CallOpts, arg0)
}

// ApprovedResetter is a free data retrieval call binding the contract method 0x9b4f5249.
//
// Solidity: function approvedResetter(address ) view returns(address)
func (_VWManager *VWManagerCallerSession) ApprovedResetter(arg0 common.Address) (common.Address, error) {
	return _VWManager.Contract.ApprovedResetter(&_VWManager.CallOpts, arg0)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_VWManager *VWManagerCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_VWManager *VWManagerSession) Deployer() (common.Address, error) {
	return _VWManager.Contract.Deployer(&_VWManager.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_VWManager *VWManagerCallerSession) Deployer() (common.Address, error) {
	return _VWManager.Contract.Deployer(&_VWManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0x9ff91054.
//
// Solidity: function domainSeparator(uint256 ) view returns(bytes32)
func (_VWManager *VWManagerCaller) DomainSeparator(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "domainSeparator", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0x9ff91054.
//
// Solidity: function domainSeparator(uint256 ) view returns(bytes32)
func (_VWManager *VWManagerSession) DomainSeparator(arg0 *big.Int) ([32]byte, error) {
	return _VWManager.Contract.DomainSeparator(&_VWManager.CallOpts, arg0)
}

// DomainSeparator is a free data retrieval call binding the contract method 0x9ff91054.
//
// Solidity: function domainSeparator(uint256 ) view returns(bytes32)
func (_VWManager *VWManagerCallerSession) DomainSeparator(arg0 *big.Int) ([32]byte, error) {
	return _VWManager.Contract.DomainSeparator(&_VWManager.CallOpts, arg0)
}

// FeeProportion is a free data retrieval call binding the contract method 0x2a3ecc41.
//
// Solidity: function feeProportion() view returns(uint256)
func (_VWManager *VWManagerCaller) FeeProportion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "feeProportion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeProportion is a free data retrieval call binding the contract method 0x2a3ecc41.
//
// Solidity: function feeProportion() view returns(uint256)
func (_VWManager *VWManagerSession) FeeProportion() (*big.Int, error) {
	return _VWManager.Contract.FeeProportion(&_VWManager.CallOpts)
}

// FeeProportion is a free data retrieval call binding the contract method 0x2a3ecc41.
//
// Solidity: function feeProportion() view returns(uint256)
func (_VWManager *VWManagerCallerSession) FeeProportion() (*big.Int, error) {
	return _VWManager.Contract.FeeProportion(&_VWManager.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_VWManager *VWManagerCaller) FeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "feeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_VWManager *VWManagerSession) FeeVault() (common.Address, error) {
	return _VWManager.Contract.FeeVault(&_VWManager.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_VWManager *VWManagerCallerSession) FeeVault() (common.Address, error) {
	return _VWManager.Contract.FeeVault(&_VWManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VWManager *VWManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VWManager *VWManagerSession) Owner() (common.Address, error) {
	return _VWManager.Contract.Owner(&_VWManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VWManager *VWManagerCallerSession) Owner() (common.Address, error) {
	return _VWManager.Contract.Owner(&_VWManager.CallOpts)
}

// OwnerWallet is a free data retrieval call binding the contract method 0x1463e592.
//
// Solidity: function ownerWallet(address ) view returns(address)
func (_VWManager *VWManagerCaller) OwnerWallet(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "ownerWallet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerWallet is a free data retrieval call binding the contract method 0x1463e592.
//
// Solidity: function ownerWallet(address ) view returns(address)
func (_VWManager *VWManagerSession) OwnerWallet(arg0 common.Address) (common.Address, error) {
	return _VWManager.Contract.OwnerWallet(&_VWManager.CallOpts, arg0)
}

// OwnerWallet is a free data retrieval call binding the contract method 0x1463e592.
//
// Solidity: function ownerWallet(address ) view returns(address)
func (_VWManager *VWManagerCallerSession) OwnerWallet(arg0 common.Address) (common.Address, error) {
	return _VWManager.Contract.OwnerWallet(&_VWManager.CallOpts, arg0)
}

// ProtocolFeeOpened is a free data retrieval call binding the contract method 0xd2bfc87c.
//
// Solidity: function protocolFeeOpened() view returns(bool)
func (_VWManager *VWManagerCaller) ProtocolFeeOpened(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "protocolFeeOpened")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProtocolFeeOpened is a free data retrieval call binding the contract method 0xd2bfc87c.
//
// Solidity: function protocolFeeOpened() view returns(bool)
func (_VWManager *VWManagerSession) ProtocolFeeOpened() (bool, error) {
	return _VWManager.Contract.ProtocolFeeOpened(&_VWManager.CallOpts)
}

// ProtocolFeeOpened is a free data retrieval call binding the contract method 0xd2bfc87c.
//
// Solidity: function protocolFeeOpened() view returns(bool)
func (_VWManager *VWManagerCallerSession) ProtocolFeeOpened() (bool, error) {
	return _VWManager.Contract.ProtocolFeeOpened(&_VWManager.CallOpts)
}

// Result is a free data retrieval call binding the contract method 0x3e6da5d5.
//
// Solidity: function result(address , uint256 ) view returns(uint256)
func (_VWManager *VWManagerCaller) Result(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "result", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Result is a free data retrieval call binding the contract method 0x3e6da5d5.
//
// Solidity: function result(address , uint256 ) view returns(uint256)
func (_VWManager *VWManagerSession) Result(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VWManager.Contract.Result(&_VWManager.CallOpts, arg0, arg1)
}

// Result is a free data retrieval call binding the contract method 0x3e6da5d5.
//
// Solidity: function result(address , uint256 ) view returns(uint256)
func (_VWManager *VWManagerCallerSession) Result(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VWManager.Contract.Result(&_VWManager.CallOpts, arg0, arg1)
}

// VerifyEIP1271Signature is a free data retrieval call binding the contract method 0x06fc0d10.
//
// Solidity: function verifyEIP1271Signature(address wallet, bytes32 digest, bytes signature) view returns(bool isValid)
func (_VWManager *VWManagerCaller) VerifyEIP1271Signature(opts *bind.CallOpts, wallet common.Address, digest [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "verifyEIP1271Signature", wallet, digest, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyEIP1271Signature is a free data retrieval call binding the contract method 0x06fc0d10.
//
// Solidity: function verifyEIP1271Signature(address wallet, bytes32 digest, bytes signature) view returns(bool isValid)
func (_VWManager *VWManagerSession) VerifyEIP1271Signature(wallet common.Address, digest [32]byte, signature []byte) (bool, error) {
	return _VWManager.Contract.VerifyEIP1271Signature(&_VWManager.CallOpts, wallet, digest, signature)
}

// VerifyEIP1271Signature is a free data retrieval call binding the contract method 0x06fc0d10.
//
// Solidity: function verifyEIP1271Signature(address wallet, bytes32 digest, bytes signature) view returns(bool isValid)
func (_VWManager *VWManagerCallerSession) VerifyEIP1271Signature(wallet common.Address, digest [32]byte, signature []byte) (bool, error) {
	return _VWManager.Contract.VerifyEIP1271Signature(&_VWManager.CallOpts, wallet, digest, signature)
}

// VolatileService is a free data retrieval call binding the contract method 0xdccc201e.
//
// Solidity: function volatileService() view returns(address)
func (_VWManager *VWManagerCaller) VolatileService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "volatileService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VolatileService is a free data retrieval call binding the contract method 0xdccc201e.
//
// Solidity: function volatileService() view returns(address)
func (_VWManager *VWManagerSession) VolatileService() (common.Address, error) {
	return _VWManager.Contract.VolatileService(&_VWManager.CallOpts)
}

// VolatileService is a free data retrieval call binding the contract method 0xdccc201e.
//
// Solidity: function volatileService() view returns(address)
func (_VWManager *VWManagerCallerSession) VolatileService() (common.Address, error) {
	return _VWManager.Contract.VolatileService(&_VWManager.CallOpts)
}

// WalletOwner is a free data retrieval call binding the contract method 0xd664a65d.
//
// Solidity: function walletOwner(address ) view returns(address)
func (_VWManager *VWManagerCaller) WalletOwner(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VWManager.contract.Call(opts, &out, "walletOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletOwner is a free data retrieval call binding the contract method 0xd664a65d.
//
// Solidity: function walletOwner(address ) view returns(address)
func (_VWManager *VWManagerSession) WalletOwner(arg0 common.Address) (common.Address, error) {
	return _VWManager.Contract.WalletOwner(&_VWManager.CallOpts, arg0)
}

// WalletOwner is a free data retrieval call binding the contract method 0xd664a65d.
//
// Solidity: function walletOwner(address ) view returns(address)
func (_VWManager *VWManagerCallerSession) WalletOwner(arg0 common.Address) (common.Address, error) {
	return _VWManager.Contract.WalletOwner(&_VWManager.CallOpts, arg0)
}

// ApproveResetter is a paid mutator transaction binding the contract method 0xb073d4ee.
//
// Solidity: function approveResetter(uint256 code, address wallet, address resetter, bool approved, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns(bytes32 dataHash)
func (_VWManager *VWManagerTransactor) ApproveResetter(opts *bind.TransactOpts, code *big.Int, wallet common.Address, resetter common.Address, approved bool, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "approveResetter", code, wallet, resetter, approved, fParam, signature)
}

// ApproveResetter is a paid mutator transaction binding the contract method 0xb073d4ee.
//
// Solidity: function approveResetter(uint256 code, address wallet, address resetter, bool approved, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns(bytes32 dataHash)
func (_VWManager *VWManagerSession) ApproveResetter(code *big.Int, wallet common.Address, resetter common.Address, approved bool, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.Contract.ApproveResetter(&_VWManager.TransactOpts, code, wallet, resetter, approved, fParam, signature)
}

// ApproveResetter is a paid mutator transaction binding the contract method 0xb073d4ee.
//
// Solidity: function approveResetter(uint256 code, address wallet, address resetter, bool approved, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns(bytes32 dataHash)
func (_VWManager *VWManagerTransactorSession) ApproveResetter(code *big.Int, wallet common.Address, resetter common.Address, approved bool, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.Contract.ApproveResetter(&_VWManager.TransactOpts, code, wallet, resetter, approved, fParam, signature)
}

// CancelTx is a paid mutator transaction binding the contract method 0x8329ee87.
//
// Solidity: function cancelTx(uint256 code, address wallet, uint256 codeToCancel, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns()
func (_VWManager *VWManagerTransactor) CancelTx(opts *bind.TransactOpts, code *big.Int, wallet common.Address, codeToCancel *big.Int, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "cancelTx", code, wallet, codeToCancel, fParam, signature)
}

// CancelTx is a paid mutator transaction binding the contract method 0x8329ee87.
//
// Solidity: function cancelTx(uint256 code, address wallet, uint256 codeToCancel, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns()
func (_VWManager *VWManagerSession) CancelTx(code *big.Int, wallet common.Address, codeToCancel *big.Int, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.Contract.CancelTx(&_VWManager.TransactOpts, code, wallet, codeToCancel, fParam, signature)
}

// CancelTx is a paid mutator transaction binding the contract method 0x8329ee87.
//
// Solidity: function cancelTx(uint256 code, address wallet, uint256 codeToCancel, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns()
func (_VWManager *VWManagerTransactorSession) CancelTx(code *big.Int, wallet common.Address, codeToCancel *big.Int, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.Contract.CancelTx(&_VWManager.TransactOpts, code, wallet, codeToCancel, fParam, signature)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x203788af.
//
// Solidity: function changeOwner(uint256 code, address wallet, address newOwner, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns()
func (_VWManager *VWManagerTransactor) ChangeOwner(opts *bind.TransactOpts, code *big.Int, wallet common.Address, newOwner common.Address, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "changeOwner", code, wallet, newOwner, fParam, signature)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x203788af.
//
// Solidity: function changeOwner(uint256 code, address wallet, address newOwner, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns()
func (_VWManager *VWManagerSession) ChangeOwner(code *big.Int, wallet common.Address, newOwner common.Address, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.Contract.ChangeOwner(&_VWManager.TransactOpts, code, wallet, newOwner, fParam, signature)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x203788af.
//
// Solidity: function changeOwner(uint256 code, address wallet, address newOwner, (address,uint256,uint256,uint256,address) fParam, bytes signature) returns()
func (_VWManager *VWManagerTransactorSession) ChangeOwner(code *big.Int, wallet common.Address, newOwner common.Address, fParam VWManagerServiceFeeParam, signature []byte) (*types.Transaction, error) {
	return _VWManager.Contract.ChangeOwner(&_VWManager.TransactOpts, code, wallet, newOwner, fParam, signature)
}

// ConfigFee is a paid mutator transaction binding the contract method 0x5eb47646.
//
// Solidity: function configFee(bool _protocolFeeOpened, address _feeVault, uint256 _feeProportion) returns()
func (_VWManager *VWManagerTransactor) ConfigFee(opts *bind.TransactOpts, _protocolFeeOpened bool, _feeVault common.Address, _feeProportion *big.Int) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "configFee", _protocolFeeOpened, _feeVault, _feeProportion)
}

// ConfigFee is a paid mutator transaction binding the contract method 0x5eb47646.
//
// Solidity: function configFee(bool _protocolFeeOpened, address _feeVault, uint256 _feeProportion) returns()
func (_VWManager *VWManagerSession) ConfigFee(_protocolFeeOpened bool, _feeVault common.Address, _feeProportion *big.Int) (*types.Transaction, error) {
	return _VWManager.Contract.ConfigFee(&_VWManager.TransactOpts, _protocolFeeOpened, _feeVault, _feeProportion)
}

// ConfigFee is a paid mutator transaction binding the contract method 0x5eb47646.
//
// Solidity: function configFee(bool _protocolFeeOpened, address _feeVault, uint256 _feeProportion) returns()
func (_VWManager *VWManagerTransactorSession) ConfigFee(_protocolFeeOpened bool, _feeVault common.Address, _feeProportion *big.Int) (*types.Transaction, error) {
	return _VWManager.Contract.ConfigFee(&_VWManager.TransactOpts, _protocolFeeOpened, _feeVault, _feeProportion)
}

// ConfigSrcChain is a paid mutator transaction binding the contract method 0x8ae2d740.
//
// Solidity: function configSrcChain(uint256 _chainId, bool _support, address _verifyingContract) returns()
func (_VWManager *VWManagerTransactor) ConfigSrcChain(opts *bind.TransactOpts, _chainId *big.Int, _support bool, _verifyingContract common.Address) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "configSrcChain", _chainId, _support, _verifyingContract)
}

// ConfigSrcChain is a paid mutator transaction binding the contract method 0x8ae2d740.
//
// Solidity: function configSrcChain(uint256 _chainId, bool _support, address _verifyingContract) returns()
func (_VWManager *VWManagerSession) ConfigSrcChain(_chainId *big.Int, _support bool, _verifyingContract common.Address) (*types.Transaction, error) {
	return _VWManager.Contract.ConfigSrcChain(&_VWManager.TransactOpts, _chainId, _support, _verifyingContract)
}

// ConfigSrcChain is a paid mutator transaction binding the contract method 0x8ae2d740.
//
// Solidity: function configSrcChain(uint256 _chainId, bool _support, address _verifyingContract) returns()
func (_VWManager *VWManagerTransactorSession) ConfigSrcChain(_chainId *big.Int, _support bool, _verifyingContract common.Address) (*types.Transaction, error) {
	return _VWManager.Contract.ConfigSrcChain(&_VWManager.TransactOpts, _chainId, _support, _verifyingContract)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xb054a9e8.
//
// Solidity: function createWallet(address owner) returns(address wallet)
func (_VWManager *VWManagerTransactor) CreateWallet(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "createWallet", owner)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xb054a9e8.
//
// Solidity: function createWallet(address owner) returns(address wallet)
func (_VWManager *VWManagerSession) CreateWallet(owner common.Address) (*types.Transaction, error) {
	return _VWManager.Contract.CreateWallet(&_VWManager.TransactOpts, owner)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xb054a9e8.
//
// Solidity: function createWallet(address owner) returns(address wallet)
func (_VWManager *VWManagerTransactorSession) CreateWallet(owner common.Address) (*types.Transaction, error) {
	return _VWManager.Contract.CreateWallet(&_VWManager.TransactOpts, owner)
}

// Execute is a paid mutator transaction binding the contract method 0x7612c9eb.
//
// Solidity: function execute(address wallet, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vweParam) returns(uint256 resCode)
func (_VWManager *VWManagerTransactor) Execute(opts *bind.TransactOpts, wallet common.Address, vweParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "execute", wallet, vweParam)
}

// Execute is a paid mutator transaction binding the contract method 0x7612c9eb.
//
// Solidity: function execute(address wallet, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vweParam) returns(uint256 resCode)
func (_VWManager *VWManagerSession) Execute(wallet common.Address, vweParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _VWManager.Contract.Execute(&_VWManager.TransactOpts, wallet, vweParam)
}

// Execute is a paid mutator transaction binding the contract method 0x7612c9eb.
//
// Solidity: function execute(address wallet, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vweParam) returns(uint256 resCode)
func (_VWManager *VWManagerTransactorSession) Execute(wallet common.Address, vweParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _VWManager.Contract.Execute(&_VWManager.TransactOpts, wallet, vweParam)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _deployer) returns()
func (_VWManager *VWManagerTransactor) Initialize(opts *bind.TransactOpts, _deployer common.Address) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "initialize", _deployer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _deployer) returns()
func (_VWManager *VWManagerSession) Initialize(_deployer common.Address) (*types.Transaction, error) {
	return _VWManager.Contract.Initialize(&_VWManager.TransactOpts, _deployer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _deployer) returns()
func (_VWManager *VWManagerTransactorSession) Initialize(_deployer common.Address) (*types.Transaction, error) {
	return _VWManager.Contract.Initialize(&_VWManager.TransactOpts, _deployer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VWManager *VWManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VWManager *VWManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _VWManager.Contract.RenounceOwnership(&_VWManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VWManager *VWManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VWManager.Contract.RenounceOwnership(&_VWManager.TransactOpts)
}

// ResetOwner is a paid mutator transaction binding the contract method 0xa622b9b4.
//
// Solidity: function resetOwner(uint256 code, address wallet, address newOwner, (address,uint256,uint256,uint256,address) fParam, bytes data) returns()
func (_VWManager *VWManagerTransactor) ResetOwner(opts *bind.TransactOpts, code *big.Int, wallet common.Address, newOwner common.Address, fParam VWManagerServiceFeeParam, data []byte) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "resetOwner", code, wallet, newOwner, fParam, data)
}

// ResetOwner is a paid mutator transaction binding the contract method 0xa622b9b4.
//
// Solidity: function resetOwner(uint256 code, address wallet, address newOwner, (address,uint256,uint256,uint256,address) fParam, bytes data) returns()
func (_VWManager *VWManagerSession) ResetOwner(code *big.Int, wallet common.Address, newOwner common.Address, fParam VWManagerServiceFeeParam, data []byte) (*types.Transaction, error) {
	return _VWManager.Contract.ResetOwner(&_VWManager.TransactOpts, code, wallet, newOwner, fParam, data)
}

// ResetOwner is a paid mutator transaction binding the contract method 0xa622b9b4.
//
// Solidity: function resetOwner(uint256 code, address wallet, address newOwner, (address,uint256,uint256,uint256,address) fParam, bytes data) returns()
func (_VWManager *VWManagerTransactorSession) ResetOwner(code *big.Int, wallet common.Address, newOwner common.Address, fParam VWManagerServiceFeeParam, data []byte) (*types.Transaction, error) {
	return _VWManager.Contract.ResetOwner(&_VWManager.TransactOpts, code, wallet, newOwner, fParam, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VWManager *VWManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VWManager *VWManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VWManager.Contract.TransferOwnership(&_VWManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VWManager *VWManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VWManager.Contract.TransferOwnership(&_VWManager.TransactOpts, newOwner)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xc5a5a6d9.
//
// Solidity: function verifyProof(uint256 resCode, address wallet, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vweParam) returns()
func (_VWManager *VWManagerTransactor) VerifyProof(opts *bind.TransactOpts, resCode *big.Int, wallet common.Address, vweParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _VWManager.contract.Transact(opts, "verifyProof", resCode, wallet, vweParam)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xc5a5a6d9.
//
// Solidity: function verifyProof(uint256 resCode, address wallet, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vweParam) returns()
func (_VWManager *VWManagerSession) VerifyProof(resCode *big.Int, wallet common.Address, vweParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _VWManager.Contract.VerifyProof(&_VWManager.TransactOpts, resCode, wallet, vweParam)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xc5a5a6d9.
//
// Solidity: function verifyProof(uint256 resCode, address wallet, (uint256,uint256,uint256,uint256,address,address,address,address,bool,bytes,bytes,bytes32[]) vweParam) returns()
func (_VWManager *VWManagerTransactorSession) VerifyProof(resCode *big.Int, wallet common.Address, vweParam IVWManagerVWExecuteParam) (*types.Transaction, error) {
	return _VWManager.Contract.VerifyProof(&_VWManager.TransactOpts, resCode, wallet, vweParam)
}

// VWManagerDomainSeparatorCanceledIterator is returned from FilterDomainSeparatorCanceled and is used to iterate over the raw logs and unpacked data for DomainSeparatorCanceled events raised by the VWManager contract.
type VWManagerDomainSeparatorCanceledIterator struct {
	Event *VWManagerDomainSeparatorCanceled // Event containing the contract specifics and raw log

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
func (it *VWManagerDomainSeparatorCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerDomainSeparatorCanceled)
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
		it.Event = new(VWManagerDomainSeparatorCanceled)
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
func (it *VWManagerDomainSeparatorCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerDomainSeparatorCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerDomainSeparatorCanceled represents a DomainSeparatorCanceled event raised by the VWManager contract.
type VWManagerDomainSeparatorCanceled struct {
	SrcChainID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDomainSeparatorCanceled is a free log retrieval operation binding the contract event 0x51100e26a826101a10edded2887953d6cd4ccda5aba13b5c18270c65bafb8781.
//
// Solidity: event DomainSeparatorCanceled(uint256 indexed srcChainID)
func (_VWManager *VWManagerFilterer) FilterDomainSeparatorCanceled(opts *bind.FilterOpts, srcChainID []*big.Int) (*VWManagerDomainSeparatorCanceledIterator, error) {

	var srcChainIDRule []interface{}
	for _, srcChainIDItem := range srcChainID {
		srcChainIDRule = append(srcChainIDRule, srcChainIDItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "DomainSeparatorCanceled", srcChainIDRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerDomainSeparatorCanceledIterator{contract: _VWManager.contract, event: "DomainSeparatorCanceled", logs: logs, sub: sub}, nil
}

// WatchDomainSeparatorCanceled is a free log subscription operation binding the contract event 0x51100e26a826101a10edded2887953d6cd4ccda5aba13b5c18270c65bafb8781.
//
// Solidity: event DomainSeparatorCanceled(uint256 indexed srcChainID)
func (_VWManager *VWManagerFilterer) WatchDomainSeparatorCanceled(opts *bind.WatchOpts, sink chan<- *VWManagerDomainSeparatorCanceled, srcChainID []*big.Int) (event.Subscription, error) {

	var srcChainIDRule []interface{}
	for _, srcChainIDItem := range srcChainID {
		srcChainIDRule = append(srcChainIDRule, srcChainIDItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "DomainSeparatorCanceled", srcChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerDomainSeparatorCanceled)
				if err := _VWManager.contract.UnpackLog(event, "DomainSeparatorCanceled", log); err != nil {
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

// ParseDomainSeparatorCanceled is a log parse operation binding the contract event 0x51100e26a826101a10edded2887953d6cd4ccda5aba13b5c18270c65bafb8781.
//
// Solidity: event DomainSeparatorCanceled(uint256 indexed srcChainID)
func (_VWManager *VWManagerFilterer) ParseDomainSeparatorCanceled(log types.Log) (*VWManagerDomainSeparatorCanceled, error) {
	event := new(VWManagerDomainSeparatorCanceled)
	if err := _VWManager.contract.UnpackLog(event, "DomainSeparatorCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerDomainSeparatorConfigedIterator is returned from FilterDomainSeparatorConfiged and is used to iterate over the raw logs and unpacked data for DomainSeparatorConfiged events raised by the VWManager contract.
type VWManagerDomainSeparatorConfigedIterator struct {
	Event *VWManagerDomainSeparatorConfiged // Event containing the contract specifics and raw log

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
func (it *VWManagerDomainSeparatorConfigedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerDomainSeparatorConfiged)
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
		it.Event = new(VWManagerDomainSeparatorConfiged)
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
func (it *VWManagerDomainSeparatorConfigedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerDomainSeparatorConfigedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerDomainSeparatorConfiged represents a DomainSeparatorConfiged event raised by the VWManager contract.
type VWManagerDomainSeparatorConfiged struct {
	SrcChainID     *big.Int
	VerifyContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDomainSeparatorConfiged is a free log retrieval operation binding the contract event 0xf0730fd2e459f960de2366aed998b8a66608257fb5e3150d3d8dc8dfce92fceb.
//
// Solidity: event DomainSeparatorConfiged(uint256 indexed srcChainID, address indexed verifyContract)
func (_VWManager *VWManagerFilterer) FilterDomainSeparatorConfiged(opts *bind.FilterOpts, srcChainID []*big.Int, verifyContract []common.Address) (*VWManagerDomainSeparatorConfigedIterator, error) {

	var srcChainIDRule []interface{}
	for _, srcChainIDItem := range srcChainID {
		srcChainIDRule = append(srcChainIDRule, srcChainIDItem)
	}
	var verifyContractRule []interface{}
	for _, verifyContractItem := range verifyContract {
		verifyContractRule = append(verifyContractRule, verifyContractItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "DomainSeparatorConfiged", srcChainIDRule, verifyContractRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerDomainSeparatorConfigedIterator{contract: _VWManager.contract, event: "DomainSeparatorConfiged", logs: logs, sub: sub}, nil
}

// WatchDomainSeparatorConfiged is a free log subscription operation binding the contract event 0xf0730fd2e459f960de2366aed998b8a66608257fb5e3150d3d8dc8dfce92fceb.
//
// Solidity: event DomainSeparatorConfiged(uint256 indexed srcChainID, address indexed verifyContract)
func (_VWManager *VWManagerFilterer) WatchDomainSeparatorConfiged(opts *bind.WatchOpts, sink chan<- *VWManagerDomainSeparatorConfiged, srcChainID []*big.Int, verifyContract []common.Address) (event.Subscription, error) {

	var srcChainIDRule []interface{}
	for _, srcChainIDItem := range srcChainID {
		srcChainIDRule = append(srcChainIDRule, srcChainIDItem)
	}
	var verifyContractRule []interface{}
	for _, verifyContractItem := range verifyContract {
		verifyContractRule = append(verifyContractRule, verifyContractItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "DomainSeparatorConfiged", srcChainIDRule, verifyContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerDomainSeparatorConfiged)
				if err := _VWManager.contract.UnpackLog(event, "DomainSeparatorConfiged", log); err != nil {
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

// ParseDomainSeparatorConfiged is a log parse operation binding the contract event 0xf0730fd2e459f960de2366aed998b8a66608257fb5e3150d3d8dc8dfce92fceb.
//
// Solidity: event DomainSeparatorConfiged(uint256 indexed srcChainID, address indexed verifyContract)
func (_VWManager *VWManagerFilterer) ParseDomainSeparatorConfiged(log types.Log) (*VWManagerDomainSeparatorConfiged, error) {
	event := new(VWManagerDomainSeparatorConfiged)
	if err := _VWManager.contract.UnpackLog(event, "DomainSeparatorConfiged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VWManager contract.
type VWManagerInitializedIterator struct {
	Event *VWManagerInitialized // Event containing the contract specifics and raw log

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
func (it *VWManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerInitialized)
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
		it.Event = new(VWManagerInitialized)
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
func (it *VWManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerInitialized represents a Initialized event raised by the VWManager contract.
type VWManagerInitialized struct {
	Deployer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e6.
//
// Solidity: event Initialized(address deployer)
func (_VWManager *VWManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*VWManagerInitializedIterator, error) {

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VWManagerInitializedIterator{contract: _VWManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e6.
//
// Solidity: event Initialized(address deployer)
func (_VWManager *VWManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VWManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerInitialized)
				if err := _VWManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e6.
//
// Solidity: event Initialized(address deployer)
func (_VWManager *VWManagerFilterer) ParseInitialized(log types.Log) (*VWManagerInitialized, error) {
	event := new(VWManagerInitialized)
	if err := _VWManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VWManager contract.
type VWManagerOwnershipTransferredIterator struct {
	Event *VWManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VWManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerOwnershipTransferred)
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
		it.Event = new(VWManagerOwnershipTransferred)
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
func (it *VWManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerOwnershipTransferred represents a OwnershipTransferred event raised by the VWManager contract.
type VWManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VWManager *VWManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VWManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerOwnershipTransferredIterator{contract: _VWManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VWManager *VWManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VWManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerOwnershipTransferred)
				if err := _VWManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VWManager *VWManagerFilterer) ParseOwnershipTransferred(log types.Log) (*VWManagerOwnershipTransferred, error) {
	event := new(VWManagerOwnershipTransferred)
	if err := _VWManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerResetterChangedIterator is returned from FilterResetterChanged and is used to iterate over the raw logs and unpacked data for ResetterChanged events raised by the VWManager contract.
type VWManagerResetterChangedIterator struct {
	Event *VWManagerResetterChanged // Event containing the contract specifics and raw log

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
func (it *VWManagerResetterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerResetterChanged)
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
		it.Event = new(VWManagerResetterChanged)
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
func (it *VWManagerResetterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerResetterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerResetterChanged represents a ResetterChanged event raised by the VWManager contract.
type VWManagerResetterChanged struct {
	Reseter common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterResetterChanged is a free log retrieval operation binding the contract event 0x67a5c84f35159d6aa301a7c05573fc2557ed7135eed8afd733039966c96b873e.
//
// Solidity: event ResetterChanged(address indexed reseter)
func (_VWManager *VWManagerFilterer) FilterResetterChanged(opts *bind.FilterOpts, reseter []common.Address) (*VWManagerResetterChangedIterator, error) {

	var reseterRule []interface{}
	for _, reseterItem := range reseter {
		reseterRule = append(reseterRule, reseterItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "ResetterChanged", reseterRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerResetterChangedIterator{contract: _VWManager.contract, event: "ResetterChanged", logs: logs, sub: sub}, nil
}

// WatchResetterChanged is a free log subscription operation binding the contract event 0x67a5c84f35159d6aa301a7c05573fc2557ed7135eed8afd733039966c96b873e.
//
// Solidity: event ResetterChanged(address indexed reseter)
func (_VWManager *VWManagerFilterer) WatchResetterChanged(opts *bind.WatchOpts, sink chan<- *VWManagerResetterChanged, reseter []common.Address) (event.Subscription, error) {

	var reseterRule []interface{}
	for _, reseterItem := range reseter {
		reseterRule = append(reseterRule, reseterItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "ResetterChanged", reseterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerResetterChanged)
				if err := _VWManager.contract.UnpackLog(event, "ResetterChanged", log); err != nil {
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

// ParseResetterChanged is a log parse operation binding the contract event 0x67a5c84f35159d6aa301a7c05573fc2557ed7135eed8afd733039966c96b873e.
//
// Solidity: event ResetterChanged(address indexed reseter)
func (_VWManager *VWManagerFilterer) ParseResetterChanged(log types.Log) (*VWManagerResetterChanged, error) {
	event := new(VWManagerResetterChanged)
	if err := _VWManager.contract.UnpackLog(event, "ResetterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the VWManager contract.
type VWManagerSetFeeIterator struct {
	Event *VWManagerSetFee // Event containing the contract specifics and raw log

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
func (it *VWManagerSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerSetFee)
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
		it.Event = new(VWManagerSetFee)
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
func (it *VWManagerSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerSetFee represents a SetFee event raised by the VWManager contract.
type VWManagerSetFee struct {
	ProtocolFeeOpened bool
	FeeVault          common.Address
	FeeProportion     *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x659336362acd38c16a3c9eddd487b6e122ddda32015683d56a8690848868db18.
//
// Solidity: event SetFee(bool indexed protocolFeeOpened, address indexed feeVault, uint256 indexed feeProportion)
func (_VWManager *VWManagerFilterer) FilterSetFee(opts *bind.FilterOpts, protocolFeeOpened []bool, feeVault []common.Address, feeProportion []*big.Int) (*VWManagerSetFeeIterator, error) {

	var protocolFeeOpenedRule []interface{}
	for _, protocolFeeOpenedItem := range protocolFeeOpened {
		protocolFeeOpenedRule = append(protocolFeeOpenedRule, protocolFeeOpenedItem)
	}
	var feeVaultRule []interface{}
	for _, feeVaultItem := range feeVault {
		feeVaultRule = append(feeVaultRule, feeVaultItem)
	}
	var feeProportionRule []interface{}
	for _, feeProportionItem := range feeProportion {
		feeProportionRule = append(feeProportionRule, feeProportionItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "SetFee", protocolFeeOpenedRule, feeVaultRule, feeProportionRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerSetFeeIterator{contract: _VWManager.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x659336362acd38c16a3c9eddd487b6e122ddda32015683d56a8690848868db18.
//
// Solidity: event SetFee(bool indexed protocolFeeOpened, address indexed feeVault, uint256 indexed feeProportion)
func (_VWManager *VWManagerFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *VWManagerSetFee, protocolFeeOpened []bool, feeVault []common.Address, feeProportion []*big.Int) (event.Subscription, error) {

	var protocolFeeOpenedRule []interface{}
	for _, protocolFeeOpenedItem := range protocolFeeOpened {
		protocolFeeOpenedRule = append(protocolFeeOpenedRule, protocolFeeOpenedItem)
	}
	var feeVaultRule []interface{}
	for _, feeVaultItem := range feeVault {
		feeVaultRule = append(feeVaultRule, feeVaultItem)
	}
	var feeProportionRule []interface{}
	for _, feeProportionItem := range feeProportion {
		feeProportionRule = append(feeProportionRule, feeProportionItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "SetFee", protocolFeeOpenedRule, feeVaultRule, feeProportionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerSetFee)
				if err := _VWManager.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x659336362acd38c16a3c9eddd487b6e122ddda32015683d56a8690848868db18.
//
// Solidity: event SetFee(bool indexed protocolFeeOpened, address indexed feeVault, uint256 indexed feeProportion)
func (_VWManager *VWManagerFilterer) ParseSetFee(log types.Log) (*VWManagerSetFee, error) {
	event := new(VWManagerSetFee)
	if err := _VWManager.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerTxCanceledIterator is returned from FilterTxCanceled and is used to iterate over the raw logs and unpacked data for TxCanceled events raised by the VWManager contract.
type VWManagerTxCanceledIterator struct {
	Event *VWManagerTxCanceled // Event containing the contract specifics and raw log

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
func (it *VWManagerTxCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerTxCanceled)
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
		it.Event = new(VWManagerTxCanceled)
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
func (it *VWManagerTxCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerTxCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerTxCanceled represents a TxCanceled event raised by the VWManager contract.
type VWManagerTxCanceled struct {
	Code *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTxCanceled is a free log retrieval operation binding the contract event 0xb59b0a9f996c69b3fe9afd164fde28d000d14e27569f68a82be4382e1ad5906d.
//
// Solidity: event TxCanceled(uint256 indexed code)
func (_VWManager *VWManagerFilterer) FilterTxCanceled(opts *bind.FilterOpts, code []*big.Int) (*VWManagerTxCanceledIterator, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "TxCanceled", codeRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerTxCanceledIterator{contract: _VWManager.contract, event: "TxCanceled", logs: logs, sub: sub}, nil
}

// WatchTxCanceled is a free log subscription operation binding the contract event 0xb59b0a9f996c69b3fe9afd164fde28d000d14e27569f68a82be4382e1ad5906d.
//
// Solidity: event TxCanceled(uint256 indexed code)
func (_VWManager *VWManagerFilterer) WatchTxCanceled(opts *bind.WatchOpts, sink chan<- *VWManagerTxCanceled, code []*big.Int) (event.Subscription, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "TxCanceled", codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerTxCanceled)
				if err := _VWManager.contract.UnpackLog(event, "TxCanceled", log); err != nil {
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

// ParseTxCanceled is a log parse operation binding the contract event 0xb59b0a9f996c69b3fe9afd164fde28d000d14e27569f68a82be4382e1ad5906d.
//
// Solidity: event TxCanceled(uint256 indexed code)
func (_VWManager *VWManagerFilterer) ParseTxCanceled(log types.Log) (*VWManagerTxCanceled, error) {
	event := new(VWManagerTxCanceled)
	if err := _VWManager.contract.UnpackLog(event, "TxCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerTxExecutedIterator is returned from FilterTxExecuted and is used to iterate over the raw logs and unpacked data for TxExecuted events raised by the VWManager contract.
type VWManagerTxExecutedIterator struct {
	Event *VWManagerTxExecuted // Event containing the contract specifics and raw log

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
func (it *VWManagerTxExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerTxExecuted)
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
		it.Event = new(VWManagerTxExecuted)
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
func (it *VWManagerTxExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerTxExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerTxExecuted represents a TxExecuted event raised by the VWManager contract.
type VWManagerTxExecuted struct {
	Wallet   common.Address
	Owner    common.Address
	Code     *big.Int
	RootHash [32]byte
	ResCode  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTxExecuted is a free log retrieval operation binding the contract event 0xc3b9e3e0453c4c483d190e681671e2fdf023fa55964f2b232816c7766502c981.
//
// Solidity: event TxExecuted(address indexed wallet, address indexed owner, uint256 indexed code, bytes32 rootHash, uint256 resCode)
func (_VWManager *VWManagerFilterer) FilterTxExecuted(opts *bind.FilterOpts, wallet []common.Address, owner []common.Address, code []*big.Int) (*VWManagerTxExecutedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "TxExecuted", walletRule, ownerRule, codeRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerTxExecutedIterator{contract: _VWManager.contract, event: "TxExecuted", logs: logs, sub: sub}, nil
}

// WatchTxExecuted is a free log subscription operation binding the contract event 0xc3b9e3e0453c4c483d190e681671e2fdf023fa55964f2b232816c7766502c981.
//
// Solidity: event TxExecuted(address indexed wallet, address indexed owner, uint256 indexed code, bytes32 rootHash, uint256 resCode)
func (_VWManager *VWManagerFilterer) WatchTxExecuted(opts *bind.WatchOpts, sink chan<- *VWManagerTxExecuted, wallet []common.Address, owner []common.Address, code []*big.Int) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "TxExecuted", walletRule, ownerRule, codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerTxExecuted)
				if err := _VWManager.contract.UnpackLog(event, "TxExecuted", log); err != nil {
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

// ParseTxExecuted is a log parse operation binding the contract event 0xc3b9e3e0453c4c483d190e681671e2fdf023fa55964f2b232816c7766502c981.
//
// Solidity: event TxExecuted(address indexed wallet, address indexed owner, uint256 indexed code, bytes32 rootHash, uint256 resCode)
func (_VWManager *VWManagerFilterer) ParseTxExecuted(log types.Log) (*VWManagerTxExecuted, error) {
	event := new(VWManagerTxExecuted)
	if err := _VWManager.contract.UnpackLog(event, "TxExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerVWCreatedIterator is returned from FilterVWCreated and is used to iterate over the raw logs and unpacked data for VWCreated events raised by the VWManager contract.
type VWManagerVWCreatedIterator struct {
	Event *VWManagerVWCreated // Event containing the contract specifics and raw log

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
func (it *VWManagerVWCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerVWCreated)
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
		it.Event = new(VWManagerVWCreated)
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
func (it *VWManagerVWCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerVWCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerVWCreated represents a VWCreated event raised by the VWManager contract.
type VWManagerVWCreated struct {
	Wallet common.Address
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVWCreated is a free log retrieval operation binding the contract event 0xbd706445a79ccc4525b3a5e1cac322a9dfe11a324422d6c08b072ed773dec3b4.
//
// Solidity: event VWCreated(address indexed wallet, address indexed owner)
func (_VWManager *VWManagerFilterer) FilterVWCreated(opts *bind.FilterOpts, wallet []common.Address, owner []common.Address) (*VWManagerVWCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "VWCreated", walletRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerVWCreatedIterator{contract: _VWManager.contract, event: "VWCreated", logs: logs, sub: sub}, nil
}

// WatchVWCreated is a free log subscription operation binding the contract event 0xbd706445a79ccc4525b3a5e1cac322a9dfe11a324422d6c08b072ed773dec3b4.
//
// Solidity: event VWCreated(address indexed wallet, address indexed owner)
func (_VWManager *VWManagerFilterer) WatchVWCreated(opts *bind.WatchOpts, sink chan<- *VWManagerVWCreated, wallet []common.Address, owner []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "VWCreated", walletRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerVWCreated)
				if err := _VWManager.contract.UnpackLog(event, "VWCreated", log); err != nil {
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

// ParseVWCreated is a log parse operation binding the contract event 0xbd706445a79ccc4525b3a5e1cac322a9dfe11a324422d6c08b072ed773dec3b4.
//
// Solidity: event VWCreated(address indexed wallet, address indexed owner)
func (_VWManager *VWManagerFilterer) ParseVWCreated(log types.Log) (*VWManagerVWCreated, error) {
	event := new(VWManagerVWCreated)
	if err := _VWManager.contract.UnpackLog(event, "VWCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VWManagerVWOwnerChangedIterator is returned from FilterVWOwnerChanged and is used to iterate over the raw logs and unpacked data for VWOwnerChanged events raised by the VWManager contract.
type VWManagerVWOwnerChangedIterator struct {
	Event *VWManagerVWOwnerChanged // Event containing the contract specifics and raw log

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
func (it *VWManagerVWOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VWManagerVWOwnerChanged)
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
		it.Event = new(VWManagerVWOwnerChanged)
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
func (it *VWManagerVWOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VWManagerVWOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VWManagerVWOwnerChanged represents a VWOwnerChanged event raised by the VWManager contract.
type VWManagerVWOwnerChanged struct {
	Wallet        common.Address
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVWOwnerChanged is a free log retrieval operation binding the contract event 0x2ac1104edeca8814839744eadd744b41cd16de4efa78c7cfe7c7ab19ce3fee8d.
//
// Solidity: event VWOwnerChanged(address indexed wallet, address indexed previousOwner, address indexed newOwner)
func (_VWManager *VWManagerFilterer) FilterVWOwnerChanged(opts *bind.FilterOpts, wallet []common.Address, previousOwner []common.Address, newOwner []common.Address) (*VWManagerVWOwnerChangedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VWManager.contract.FilterLogs(opts, "VWOwnerChanged", walletRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VWManagerVWOwnerChangedIterator{contract: _VWManager.contract, event: "VWOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchVWOwnerChanged is a free log subscription operation binding the contract event 0x2ac1104edeca8814839744eadd744b41cd16de4efa78c7cfe7c7ab19ce3fee8d.
//
// Solidity: event VWOwnerChanged(address indexed wallet, address indexed previousOwner, address indexed newOwner)
func (_VWManager *VWManagerFilterer) WatchVWOwnerChanged(opts *bind.WatchOpts, sink chan<- *VWManagerVWOwnerChanged, wallet []common.Address, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VWManager.contract.WatchLogs(opts, "VWOwnerChanged", walletRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VWManagerVWOwnerChanged)
				if err := _VWManager.contract.UnpackLog(event, "VWOwnerChanged", log); err != nil {
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

// ParseVWOwnerChanged is a log parse operation binding the contract event 0x2ac1104edeca8814839744eadd744b41cd16de4efa78c7cfe7c7ab19ce3fee8d.
//
// Solidity: event VWOwnerChanged(address indexed wallet, address indexed previousOwner, address indexed newOwner)
func (_VWManager *VWManagerFilterer) ParseVWOwnerChanged(log types.Log) (*VWManagerVWOwnerChanged, error) {
	event := new(VWManagerVWOwnerChanged)
	if err := _VWManager.contract.UnpackLog(event, "VWOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
