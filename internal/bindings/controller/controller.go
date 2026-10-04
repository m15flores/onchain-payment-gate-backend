// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controller

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"authorizedBack_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizedBack\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRelease\",\"inputs\":[{\"name\":\"to_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVault\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ReleaseRequested\",\"inputs\":[{\"name\":\"to_\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount_\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedBack is a free data retrieval call binding the contract method 0xa5025798.
//
// Solidity: function authorizedBack() view returns(address)
func (_Controller *ControllerCaller) AuthorizedBack(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "authorizedBack")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthorizedBack is a free data retrieval call binding the contract method 0xa5025798.
//
// Solidity: function authorizedBack() view returns(address)
func (_Controller *ControllerSession) AuthorizedBack() (common.Address, error) {
	return _Controller.Contract.AuthorizedBack(&_Controller.CallOpts)
}

// AuthorizedBack is a free data retrieval call binding the contract method 0xa5025798.
//
// Solidity: function authorizedBack() view returns(address)
func (_Controller *ControllerCallerSession) AuthorizedBack() (common.Address, error) {
	return _Controller.Contract.AuthorizedBack(&_Controller.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Controller *ControllerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Controller *ControllerSession) Vault() (common.Address, error) {
	return _Controller.Contract.Vault(&_Controller.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Controller *ControllerCallerSession) Vault() (common.Address, error) {
	return _Controller.Contract.Vault(&_Controller.CallOpts)
}

// RequestRelease is a paid mutator transaction binding the contract method 0x9b4e7007.
//
// Solidity: function requestRelease(address to_, uint256 amount_) returns()
func (_Controller *ControllerTransactor) RequestRelease(opts *bind.TransactOpts, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "requestRelease", to_, amount_)
}

// RequestRelease is a paid mutator transaction binding the contract method 0x9b4e7007.
//
// Solidity: function requestRelease(address to_, uint256 amount_) returns()
func (_Controller *ControllerSession) RequestRelease(to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.RequestRelease(&_Controller.TransactOpts, to_, amount_)
}

// RequestRelease is a paid mutator transaction binding the contract method 0x9b4e7007.
//
// Solidity: function requestRelease(address to_, uint256 amount_) returns()
func (_Controller *ControllerTransactorSession) RequestRelease(to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.RequestRelease(&_Controller.TransactOpts, to_, amount_)
}

// ControllerReleaseRequestedIterator is returned from FilterReleaseRequested and is used to iterate over the raw logs and unpacked data for ReleaseRequested events raised by the Controller contract.
type ControllerReleaseRequestedIterator struct {
	Event *ControllerReleaseRequested // Event containing the contract specifics and raw log

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
func (it *ControllerReleaseRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerReleaseRequested)
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
		it.Event = new(ControllerReleaseRequested)
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
func (it *ControllerReleaseRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerReleaseRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerReleaseRequested represents a ReleaseRequested event raised by the Controller contract.
type ControllerReleaseRequested struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReleaseRequested is a free log retrieval operation binding the contract event 0xb119911fe330a2404b309b712fc308317955383d897319b83fe8f4f805e63d23.
//
// Solidity: event ReleaseRequested(address indexed to_, uint256 amount_)
func (_Controller *ControllerFilterer) FilterReleaseRequested(opts *bind.FilterOpts, to_ []common.Address) (*ControllerReleaseRequestedIterator, error) {

	var to_Rule []interface{}
	for _, to_Item := range to_ {
		to_Rule = append(to_Rule, to_Item)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "ReleaseRequested", to_Rule)
	if err != nil {
		return nil, err
	}
	return &ControllerReleaseRequestedIterator{contract: _Controller.contract, event: "ReleaseRequested", logs: logs, sub: sub}, nil
}

// WatchReleaseRequested is a free log subscription operation binding the contract event 0xb119911fe330a2404b309b712fc308317955383d897319b83fe8f4f805e63d23.
//
// Solidity: event ReleaseRequested(address indexed to_, uint256 amount_)
func (_Controller *ControllerFilterer) WatchReleaseRequested(opts *bind.WatchOpts, sink chan<- *ControllerReleaseRequested, to_ []common.Address) (event.Subscription, error) {

	var to_Rule []interface{}
	for _, to_Item := range to_ {
		to_Rule = append(to_Rule, to_Item)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "ReleaseRequested", to_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerReleaseRequested)
				if err := _Controller.contract.UnpackLog(event, "ReleaseRequested", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseReleaseRequested is a log parse operation binding the contract event 0xb119911fe330a2404b309b712fc308317955383d897319b83fe8f4f805e63d23.
//
// Solidity: event ReleaseRequested(address indexed to_, uint256 amount_)
func (_Controller *ControllerFilterer) ParseReleaseRequested(log types.Log) (*ControllerReleaseRequested, error) {
	event := new(ControllerReleaseRequested)
	if err := _Controller.contract.UnpackLog(event, "ReleaseRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
