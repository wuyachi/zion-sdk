// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neo3_state_manager_abi

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
)

// Neo3StateManagerMetaData contains all meta data concerning the Neo3StateManager contract.
var Neo3StateManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ID\",\"type\":\"uint64\"}],\"name\":\"evtApproveRegisterStateValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ID\",\"type\":\"uint64\"}],\"name\":\"evtApproveRemoveStateValidator\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"approveRegisterStateValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"approveRemoveStateValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentStateValidator\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"Validator\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"StateValidators\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"registerStateValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"StateValidators\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"removeStateValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ca1c4d1b": "approveRegisterStateValidator(uint64,address)",
		"3473fd55": "approveRemoveStateValidator(uint64,address)",
		"770fa9ad": "getCurrentStateValidator()",
		"06fdde03": "name()",
		"f7531edd": "registerStateValidator(string[],address)",
		"d62c2f61": "removeStateValidator(string[],address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610302806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806306fdde03146100675780633473fd5514610085578063770fa9ad14610067578063ca1c4d1b14610085578063d62c2f61146100a5578063f7531edd146100a5575b600080fd5b61006f6100b3565b60405161007c9190610272565b60405180910390f35b6100986100933660046101dc565b6100b8565b60405161007c9190610267565b6100986100933660046100dc565b606090565b600092915050565b80356001600160a01b03811681146100d757600080fd5b919050565b60008060408084860312156100ef578283fd5b833567ffffffffffffffff80821115610106578485fd5b818601915086601f830112610119578485fd5b813560208282111561012d5761012d6102b6565b61013a818284020161028c565b82815281810190858301895b858110156101bc57813588018d603f820112610160578b8cfd5b8581013588811115610174576101746102b6565b610186601f8201601f1916880161028c565b8181528f8c838501011115610199578d8efd5b818c84018983013790810187018d90528552509284019290840190600101610146565b505080985050506101ce8189016100c0565b955050505050509250929050565b600080604083850312156101ee578182fd5b823567ffffffffffffffff81168114610205578283fd5b9150610213602084016100c0565b90509250929050565b60008151808452815b8181101561024157602081850181015186830182015201610225565b818111156102525782602083870101525b50601f01601f19169290920160200192915050565b901515815260200190565b600060208252610285602083018461021c565b9392505050565b60405181810167ffffffffffffffff811182821017156102ae576102ae6102b6565b604052919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220a8f2157c1581ab9dbaa50ba995d988d6b85ace3c199e4be64e3ef0d93855952864736f6c63430008000033",
}

// Neo3StateManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use Neo3StateManagerMetaData.ABI instead.
var Neo3StateManagerABI = Neo3StateManagerMetaData.ABI

// Deprecated: Use Neo3StateManagerMetaData.Sigs instead.
// Neo3StateManagerFuncSigs maps the 4-byte function signature to its string representation.
var Neo3StateManagerFuncSigs = Neo3StateManagerMetaData.Sigs

// Neo3StateManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Neo3StateManagerMetaData.Bin instead.
var Neo3StateManagerBin = Neo3StateManagerMetaData.Bin

// DeployNeo3StateManager deploys a new Ethereum contract, binding an instance of Neo3StateManager to it.
func DeployNeo3StateManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Neo3StateManager, error) {
	parsed, err := Neo3StateManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Neo3StateManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Neo3StateManager{Neo3StateManagerCaller: Neo3StateManagerCaller{contract: contract}, Neo3StateManagerTransactor: Neo3StateManagerTransactor{contract: contract}, Neo3StateManagerFilterer: Neo3StateManagerFilterer{contract: contract}}, nil
}

// Neo3StateManager is an auto generated Go binding around an Ethereum contract.
type Neo3StateManager struct {
	Neo3StateManagerCaller     // Read-only binding to the contract
	Neo3StateManagerTransactor // Write-only binding to the contract
	Neo3StateManagerFilterer   // Log filterer for contract events
}

// Neo3StateManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Neo3StateManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Neo3StateManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Neo3StateManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Neo3StateManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Neo3StateManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Neo3StateManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Neo3StateManagerSession struct {
	Contract     *Neo3StateManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Neo3StateManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Neo3StateManagerCallerSession struct {
	Contract *Neo3StateManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Neo3StateManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Neo3StateManagerTransactorSession struct {
	Contract     *Neo3StateManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Neo3StateManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Neo3StateManagerRaw struct {
	Contract *Neo3StateManager // Generic contract binding to access the raw methods on
}

// Neo3StateManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Neo3StateManagerCallerRaw struct {
	Contract *Neo3StateManagerCaller // Generic read-only contract binding to access the raw methods on
}

// Neo3StateManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Neo3StateManagerTransactorRaw struct {
	Contract *Neo3StateManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNeo3StateManager creates a new instance of Neo3StateManager, bound to a specific deployed contract.
func NewNeo3StateManager(address common.Address, backend bind.ContractBackend) (*Neo3StateManager, error) {
	contract, err := bindNeo3StateManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Neo3StateManager{Neo3StateManagerCaller: Neo3StateManagerCaller{contract: contract}, Neo3StateManagerTransactor: Neo3StateManagerTransactor{contract: contract}, Neo3StateManagerFilterer: Neo3StateManagerFilterer{contract: contract}}, nil
}

// NewNeo3StateManagerCaller creates a new read-only instance of Neo3StateManager, bound to a specific deployed contract.
func NewNeo3StateManagerCaller(address common.Address, caller bind.ContractCaller) (*Neo3StateManagerCaller, error) {
	contract, err := bindNeo3StateManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Neo3StateManagerCaller{contract: contract}, nil
}

// NewNeo3StateManagerTransactor creates a new write-only instance of Neo3StateManager, bound to a specific deployed contract.
func NewNeo3StateManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*Neo3StateManagerTransactor, error) {
	contract, err := bindNeo3StateManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Neo3StateManagerTransactor{contract: contract}, nil
}

// NewNeo3StateManagerFilterer creates a new log filterer instance of Neo3StateManager, bound to a specific deployed contract.
func NewNeo3StateManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*Neo3StateManagerFilterer, error) {
	contract, err := bindNeo3StateManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Neo3StateManagerFilterer{contract: contract}, nil
}

// bindNeo3StateManager binds a generic wrapper to an already deployed contract.
func bindNeo3StateManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Neo3StateManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Neo3StateManager *Neo3StateManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Neo3StateManager.Contract.Neo3StateManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Neo3StateManager *Neo3StateManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.Neo3StateManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Neo3StateManager *Neo3StateManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.Neo3StateManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Neo3StateManager *Neo3StateManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Neo3StateManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Neo3StateManager *Neo3StateManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Neo3StateManager *Neo3StateManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.contract.Transact(opts, method, params...)
}

// ApproveRegisterStateValidator is a paid mutator transaction binding the contract method 0xca1c4d1b.
//
// Solidity: function approveRegisterStateValidator(uint64 ID, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerTransactor) ApproveRegisterStateValidator(opts *bind.TransactOpts, ID uint64, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.contract.Transact(opts, "approveRegisterStateValidator", ID, Address)
}

// ApproveRegisterStateValidator is a paid mutator transaction binding the contract method 0xca1c4d1b.
//
// Solidity: function approveRegisterStateValidator(uint64 ID, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerSession) ApproveRegisterStateValidator(ID uint64, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.ApproveRegisterStateValidator(&_Neo3StateManager.TransactOpts, ID, Address)
}

// ApproveRegisterStateValidator is a paid mutator transaction binding the contract method 0xca1c4d1b.
//
// Solidity: function approveRegisterStateValidator(uint64 ID, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerTransactorSession) ApproveRegisterStateValidator(ID uint64, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.ApproveRegisterStateValidator(&_Neo3StateManager.TransactOpts, ID, Address)
}

// ApproveRemoveStateValidator is a paid mutator transaction binding the contract method 0x3473fd55.
//
// Solidity: function approveRemoveStateValidator(uint64 ID, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerTransactor) ApproveRemoveStateValidator(opts *bind.TransactOpts, ID uint64, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.contract.Transact(opts, "approveRemoveStateValidator", ID, Address)
}

// ApproveRemoveStateValidator is a paid mutator transaction binding the contract method 0x3473fd55.
//
// Solidity: function approveRemoveStateValidator(uint64 ID, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerSession) ApproveRemoveStateValidator(ID uint64, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.ApproveRemoveStateValidator(&_Neo3StateManager.TransactOpts, ID, Address)
}

// ApproveRemoveStateValidator is a paid mutator transaction binding the contract method 0x3473fd55.
//
// Solidity: function approveRemoveStateValidator(uint64 ID, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerTransactorSession) ApproveRemoveStateValidator(ID uint64, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.ApproveRemoveStateValidator(&_Neo3StateManager.TransactOpts, ID, Address)
}

// GetCurrentStateValidator is a paid mutator transaction binding the contract method 0x770fa9ad.
//
// Solidity: function getCurrentStateValidator() returns(bytes Validator)
func (_Neo3StateManager *Neo3StateManagerTransactor) GetCurrentStateValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neo3StateManager.contract.Transact(opts, "getCurrentStateValidator")
}

// GetCurrentStateValidator is a paid mutator transaction binding the contract method 0x770fa9ad.
//
// Solidity: function getCurrentStateValidator() returns(bytes Validator)
func (_Neo3StateManager *Neo3StateManagerSession) GetCurrentStateValidator() (*types.Transaction, error) {
	return _Neo3StateManager.Contract.GetCurrentStateValidator(&_Neo3StateManager.TransactOpts)
}

// GetCurrentStateValidator is a paid mutator transaction binding the contract method 0x770fa9ad.
//
// Solidity: function getCurrentStateValidator() returns(bytes Validator)
func (_Neo3StateManager *Neo3StateManagerTransactorSession) GetCurrentStateValidator() (*types.Transaction, error) {
	return _Neo3StateManager.Contract.GetCurrentStateValidator(&_Neo3StateManager.TransactOpts)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string Name)
func (_Neo3StateManager *Neo3StateManagerTransactor) Name(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neo3StateManager.contract.Transact(opts, "name")
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string Name)
func (_Neo3StateManager *Neo3StateManagerSession) Name() (*types.Transaction, error) {
	return _Neo3StateManager.Contract.Name(&_Neo3StateManager.TransactOpts)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string Name)
func (_Neo3StateManager *Neo3StateManagerTransactorSession) Name() (*types.Transaction, error) {
	return _Neo3StateManager.Contract.Name(&_Neo3StateManager.TransactOpts)
}

// RegisterStateValidator is a paid mutator transaction binding the contract method 0xf7531edd.
//
// Solidity: function registerStateValidator(string[] StateValidators, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerTransactor) RegisterStateValidator(opts *bind.TransactOpts, StateValidators []string, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.contract.Transact(opts, "registerStateValidator", StateValidators, Address)
}

// RegisterStateValidator is a paid mutator transaction binding the contract method 0xf7531edd.
//
// Solidity: function registerStateValidator(string[] StateValidators, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerSession) RegisterStateValidator(StateValidators []string, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.RegisterStateValidator(&_Neo3StateManager.TransactOpts, StateValidators, Address)
}

// RegisterStateValidator is a paid mutator transaction binding the contract method 0xf7531edd.
//
// Solidity: function registerStateValidator(string[] StateValidators, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerTransactorSession) RegisterStateValidator(StateValidators []string, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.RegisterStateValidator(&_Neo3StateManager.TransactOpts, StateValidators, Address)
}

// RemoveStateValidator is a paid mutator transaction binding the contract method 0xd62c2f61.
//
// Solidity: function removeStateValidator(string[] StateValidators, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerTransactor) RemoveStateValidator(opts *bind.TransactOpts, StateValidators []string, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.contract.Transact(opts, "removeStateValidator", StateValidators, Address)
}

// RemoveStateValidator is a paid mutator transaction binding the contract method 0xd62c2f61.
//
// Solidity: function removeStateValidator(string[] StateValidators, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerSession) RemoveStateValidator(StateValidators []string, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.RemoveStateValidator(&_Neo3StateManager.TransactOpts, StateValidators, Address)
}

// RemoveStateValidator is a paid mutator transaction binding the contract method 0xd62c2f61.
//
// Solidity: function removeStateValidator(string[] StateValidators, address Address) returns(bool success)
func (_Neo3StateManager *Neo3StateManagerTransactorSession) RemoveStateValidator(StateValidators []string, Address common.Address) (*types.Transaction, error) {
	return _Neo3StateManager.Contract.RemoveStateValidator(&_Neo3StateManager.TransactOpts, StateValidators, Address)
}

// Neo3StateManagerEvtApproveRegisterStateValidatorIterator is returned from FilterEvtApproveRegisterStateValidator and is used to iterate over the raw logs and unpacked data for EvtApproveRegisterStateValidator events raised by the Neo3StateManager contract.
type Neo3StateManagerEvtApproveRegisterStateValidatorIterator struct {
	Event *Neo3StateManagerEvtApproveRegisterStateValidator // Event containing the contract specifics and raw log

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
func (it *Neo3StateManagerEvtApproveRegisterStateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Neo3StateManagerEvtApproveRegisterStateValidator)
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
		it.Event = new(Neo3StateManagerEvtApproveRegisterStateValidator)
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
func (it *Neo3StateManagerEvtApproveRegisterStateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Neo3StateManagerEvtApproveRegisterStateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Neo3StateManagerEvtApproveRegisterStateValidator represents a EvtApproveRegisterStateValidator event raised by the Neo3StateManager contract.
type Neo3StateManagerEvtApproveRegisterStateValidator struct {
	ID  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEvtApproveRegisterStateValidator is a free log retrieval operation binding the contract event 0xa2d1b53f6d5ea7964c14d81acde072ef679b771780c18a93206cd3319c6621d6.
//
// Solidity: event evtApproveRegisterStateValidator(uint64 ID)
func (_Neo3StateManager *Neo3StateManagerFilterer) FilterEvtApproveRegisterStateValidator(opts *bind.FilterOpts) (*Neo3StateManagerEvtApproveRegisterStateValidatorIterator, error) {

	logs, sub, err := _Neo3StateManager.contract.FilterLogs(opts, "evtApproveRegisterStateValidator")
	if err != nil {
		return nil, err
	}
	return &Neo3StateManagerEvtApproveRegisterStateValidatorIterator{contract: _Neo3StateManager.contract, event: "evtApproveRegisterStateValidator", logs: logs, sub: sub}, nil
}

// WatchEvtApproveRegisterStateValidator is a free log subscription operation binding the contract event 0xa2d1b53f6d5ea7964c14d81acde072ef679b771780c18a93206cd3319c6621d6.
//
// Solidity: event evtApproveRegisterStateValidator(uint64 ID)
func (_Neo3StateManager *Neo3StateManagerFilterer) WatchEvtApproveRegisterStateValidator(opts *bind.WatchOpts, sink chan<- *Neo3StateManagerEvtApproveRegisterStateValidator) (event.Subscription, error) {

	logs, sub, err := _Neo3StateManager.contract.WatchLogs(opts, "evtApproveRegisterStateValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Neo3StateManagerEvtApproveRegisterStateValidator)
				if err := _Neo3StateManager.contract.UnpackLog(event, "evtApproveRegisterStateValidator", log); err != nil {
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

// ParseEvtApproveRegisterStateValidator is a log parse operation binding the contract event 0xa2d1b53f6d5ea7964c14d81acde072ef679b771780c18a93206cd3319c6621d6.
//
// Solidity: event evtApproveRegisterStateValidator(uint64 ID)
func (_Neo3StateManager *Neo3StateManagerFilterer) ParseEvtApproveRegisterStateValidator(log types.Log) (*Neo3StateManagerEvtApproveRegisterStateValidator, error) {
	event := new(Neo3StateManagerEvtApproveRegisterStateValidator)
	if err := _Neo3StateManager.contract.UnpackLog(event, "evtApproveRegisterStateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Neo3StateManagerEvtApproveRemoveStateValidatorIterator is returned from FilterEvtApproveRemoveStateValidator and is used to iterate over the raw logs and unpacked data for EvtApproveRemoveStateValidator events raised by the Neo3StateManager contract.
type Neo3StateManagerEvtApproveRemoveStateValidatorIterator struct {
	Event *Neo3StateManagerEvtApproveRemoveStateValidator // Event containing the contract specifics and raw log

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
func (it *Neo3StateManagerEvtApproveRemoveStateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Neo3StateManagerEvtApproveRemoveStateValidator)
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
		it.Event = new(Neo3StateManagerEvtApproveRemoveStateValidator)
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
func (it *Neo3StateManagerEvtApproveRemoveStateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Neo3StateManagerEvtApproveRemoveStateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Neo3StateManagerEvtApproveRemoveStateValidator represents a EvtApproveRemoveStateValidator event raised by the Neo3StateManager contract.
type Neo3StateManagerEvtApproveRemoveStateValidator struct {
	ID  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEvtApproveRemoveStateValidator is a free log retrieval operation binding the contract event 0x32050489af2e3d6571e4c66ba0ca36ccf2b8feed7a7059d23aab90998d263604.
//
// Solidity: event evtApproveRemoveStateValidator(uint64 ID)
func (_Neo3StateManager *Neo3StateManagerFilterer) FilterEvtApproveRemoveStateValidator(opts *bind.FilterOpts) (*Neo3StateManagerEvtApproveRemoveStateValidatorIterator, error) {

	logs, sub, err := _Neo3StateManager.contract.FilterLogs(opts, "evtApproveRemoveStateValidator")
	if err != nil {
		return nil, err
	}
	return &Neo3StateManagerEvtApproveRemoveStateValidatorIterator{contract: _Neo3StateManager.contract, event: "evtApproveRemoveStateValidator", logs: logs, sub: sub}, nil
}

// WatchEvtApproveRemoveStateValidator is a free log subscription operation binding the contract event 0x32050489af2e3d6571e4c66ba0ca36ccf2b8feed7a7059d23aab90998d263604.
//
// Solidity: event evtApproveRemoveStateValidator(uint64 ID)
func (_Neo3StateManager *Neo3StateManagerFilterer) WatchEvtApproveRemoveStateValidator(opts *bind.WatchOpts, sink chan<- *Neo3StateManagerEvtApproveRemoveStateValidator) (event.Subscription, error) {

	logs, sub, err := _Neo3StateManager.contract.WatchLogs(opts, "evtApproveRemoveStateValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Neo3StateManagerEvtApproveRemoveStateValidator)
				if err := _Neo3StateManager.contract.UnpackLog(event, "evtApproveRemoveStateValidator", log); err != nil {
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

// ParseEvtApproveRemoveStateValidator is a log parse operation binding the contract event 0x32050489af2e3d6571e4c66ba0ca36ccf2b8feed7a7059d23aab90998d263604.
//
// Solidity: event evtApproveRemoveStateValidator(uint64 ID)
func (_Neo3StateManager *Neo3StateManagerFilterer) ParseEvtApproveRemoveStateValidator(log types.Log) (*Neo3StateManagerEvtApproveRemoveStateValidator, error) {
	event := new(Neo3StateManagerEvtApproveRemoveStateValidator)
	if err := _Neo3StateManager.contract.UnpackLog(event, "evtApproveRemoveStateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
