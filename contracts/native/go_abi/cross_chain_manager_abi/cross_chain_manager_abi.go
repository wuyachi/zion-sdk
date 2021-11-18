// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cross_chain_manager_abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CrossChainManagerABI is the input ABI used to generate the binding from.
const CrossChainManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"merkleValueHex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"BlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"makeProof\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"SourceChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Height\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"Proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"RelayerAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Extra\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"HeaderOrCrossChainMsg\",\"type\":\"bytes\"}],\"name\":\"importOuterTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var CrossChainManagerParsedABI, _ = abi.JSON(strings.NewReader(CrossChainManagerABI))

// CrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var CrossChainManagerFuncSigs = map[string]string{
	"5b60b01e": "importOuterTransfer(uint64,uint32,bytes,bytes,bytes,bytes)",
}

// CrossChainManager is an auto generated Go binding around an Ethereum contract.
type CrossChainManager struct {
	CrossChainManagerCaller     // Read-only binding to the contract
	CrossChainManagerTransactor // Write-only binding to the contract
	CrossChainManagerFilterer   // Log filterer for contract events
}

// CrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainManagerSession struct {
	Contract     *CrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainManagerCallerSession struct {
	Contract *CrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainManagerTransactorSession struct {
	Contract     *CrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainManagerRaw struct {
	Contract *CrossChainManager // Generic contract binding to access the raw methods on
}

// CrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainManagerCallerRaw struct {
	Contract *CrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainManagerTransactorRaw struct {
	Contract *CrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChainManager creates a new instance of CrossChainManager, bound to a specific deployed contract.
func NewCrossChainManager(address common.Address, backend bind.ContractBackend) (*CrossChainManager, error) {
	contract, err := bindCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainManager{CrossChainManagerCaller: CrossChainManagerCaller{contract: contract}, CrossChainManagerTransactor: CrossChainManagerTransactor{contract: contract}, CrossChainManagerFilterer: CrossChainManagerFilterer{contract: contract}}, nil
}

// NewCrossChainManagerCaller creates a new read-only instance of CrossChainManager, bound to a specific deployed contract.
func NewCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*CrossChainManagerCaller, error) {
	contract, err := bindCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerCaller{contract: contract}, nil
}

// NewCrossChainManagerTransactor creates a new write-only instance of CrossChainManager, bound to a specific deployed contract.
func NewCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainManagerTransactor, error) {
	contract, err := bindCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerTransactor{contract: contract}, nil
}

// NewCrossChainManagerFilterer creates a new log filterer instance of CrossChainManager, bound to a specific deployed contract.
func NewCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainManagerFilterer, error) {
	contract, err := bindCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerFilterer{contract: contract}, nil
}

// bindCrossChainManager binds a generic wrapper to an already deployed contract.
func bindCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainManager *CrossChainManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CrossChainManager.Contract.CrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainManager *CrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainManager.Contract.CrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainManager *CrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainManager.Contract.CrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainManager *CrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainManager *CrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainManager *CrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0x5b60b01e.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes RelayerAddress, bytes Extra, bytes HeaderOrCrossChainMsg) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactor) ImportOuterTransfer(opts *bind.TransactOpts, SourceChainID uint64, Height uint32, Proof []byte, RelayerAddress []byte, Extra []byte, HeaderOrCrossChainMsg []byte) (*types.Transaction, error) {
	return _CrossChainManager.contract.Transact(opts, "importOuterTransfer", SourceChainID, Height, Proof, RelayerAddress, Extra, HeaderOrCrossChainMsg)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0x5b60b01e.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes RelayerAddress, bytes Extra, bytes HeaderOrCrossChainMsg) returns(bool success)
func (_CrossChainManager *CrossChainManagerSession) ImportOuterTransfer(SourceChainID uint64, Height uint32, Proof []byte, RelayerAddress []byte, Extra []byte, HeaderOrCrossChainMsg []byte) (*types.Transaction, error) {
	return _CrossChainManager.Contract.ImportOuterTransfer(&_CrossChainManager.TransactOpts, SourceChainID, Height, Proof, RelayerAddress, Extra, HeaderOrCrossChainMsg)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0x5b60b01e.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes RelayerAddress, bytes Extra, bytes HeaderOrCrossChainMsg) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactorSession) ImportOuterTransfer(SourceChainID uint64, Height uint32, Proof []byte, RelayerAddress []byte, Extra []byte, HeaderOrCrossChainMsg []byte) (*types.Transaction, error) {
	return _CrossChainManager.Contract.ImportOuterTransfer(&_CrossChainManager.TransactOpts, SourceChainID, Height, Proof, RelayerAddress, Extra, HeaderOrCrossChainMsg)
}

// CrossChainManagerMakeProofIterator is returned from FilterMakeProof and is used to iterate over the raw logs and unpacked data for MakeProof events raised by the CrossChainManager contract.
type CrossChainManagerMakeProofIterator struct {
	Event *CrossChainManagerMakeProof // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerMakeProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerMakeProof)
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
		it.Event = new(CrossChainManagerMakeProof)
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
func (it *CrossChainManagerMakeProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerMakeProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerMakeProof represents a MakeProof event raised by the CrossChainManager contract.
type CrossChainManagerMakeProof struct {
	MerkleValueHex string
	BlockHeight    uint64
	Key            string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMakeProof is a free log retrieval operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManager *CrossChainManagerFilterer) FilterMakeProof(opts *bind.FilterOpts) (*CrossChainManagerMakeProofIterator, error) {

	logs, sub, err := _CrossChainManager.contract.FilterLogs(opts, "makeProof")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerMakeProofIterator{contract: _CrossChainManager.contract, event: "makeProof", logs: logs, sub: sub}, nil
}

var MakeProofTopicHash = "0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d"

// WatchMakeProof is a free log subscription operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManager *CrossChainManagerFilterer) WatchMakeProof(opts *bind.WatchOpts, sink chan<- *CrossChainManagerMakeProof) (event.Subscription, error) {

	logs, sub, err := _CrossChainManager.contract.WatchLogs(opts, "makeProof")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerMakeProof)
				if err := _CrossChainManager.contract.UnpackLog(event, "makeProof", log); err != nil {
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

// ParseMakeProof is a log parse operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManager *CrossChainManagerFilterer) ParseMakeProof(log types.Log) (*CrossChainManagerMakeProof, error) {
	event := new(CrossChainManagerMakeProof)
	if err := _CrossChainManager.contract.UnpackLog(event, "makeProof", log); err != nil {
		return nil, err
	}
	return event, nil
}

