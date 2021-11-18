// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package header_sync_abi

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

// HeaderSyncABI is the input ABI used to generate the binding from.
const HeaderSyncABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"BlockHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"Height\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"NextValidatorsHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"InfoChainID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"BlockHeight\",\"type\":\"uint64\"}],\"name\":\"OKEpochSwitchInfoEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"blockHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"}],\"name\":\"syncHeader\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"Headers\",\"type\":\"bytes[]\"}],\"name\":\"syncBlockHeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"CrossChainMsgs\",\"type\":\"bytes[]\"}],\"name\":\"syncCrossChainMsg\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"GenesisHeader\",\"type\":\"bytes\"}],\"name\":\"syncGenesisHeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var HeaderSyncParsedABI, _ = abi.JSON(strings.NewReader(HeaderSyncABI))

// HeaderSyncFuncSigs maps the 4-byte function signature to its string representation.
var HeaderSyncFuncSigs = map[string]string{
	"72ce6700": "syncBlockHeader(uint64,address,bytes[])",
	"21b5cff5": "syncCrossChainMsg(uint64,address,bytes[])",
	"b5ace618": "syncGenesisHeader(uint64,bytes)",
}

// HeaderSync is an auto generated Go binding around an Ethereum contract.
type HeaderSync struct {
	HeaderSyncCaller     // Read-only binding to the contract
	HeaderSyncTransactor // Write-only binding to the contract
	HeaderSyncFilterer   // Log filterer for contract events
}

// HeaderSyncCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeaderSyncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderSyncTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeaderSyncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderSyncFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeaderSyncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderSyncSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeaderSyncSession struct {
	Contract     *HeaderSync       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderSyncCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeaderSyncCallerSession struct {
	Contract *HeaderSyncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HeaderSyncTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeaderSyncTransactorSession struct {
	Contract     *HeaderSyncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HeaderSyncRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeaderSyncRaw struct {
	Contract *HeaderSync // Generic contract binding to access the raw methods on
}

// HeaderSyncCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeaderSyncCallerRaw struct {
	Contract *HeaderSyncCaller // Generic read-only contract binding to access the raw methods on
}

// HeaderSyncTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeaderSyncTransactorRaw struct {
	Contract *HeaderSyncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeaderSync creates a new instance of HeaderSync, bound to a specific deployed contract.
func NewHeaderSync(address common.Address, backend bind.ContractBackend) (*HeaderSync, error) {
	contract, err := bindHeaderSync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HeaderSync{HeaderSyncCaller: HeaderSyncCaller{contract: contract}, HeaderSyncTransactor: HeaderSyncTransactor{contract: contract}, HeaderSyncFilterer: HeaderSyncFilterer{contract: contract}}, nil
}

// NewHeaderSyncCaller creates a new read-only instance of HeaderSync, bound to a specific deployed contract.
func NewHeaderSyncCaller(address common.Address, caller bind.ContractCaller) (*HeaderSyncCaller, error) {
	contract, err := bindHeaderSync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderSyncCaller{contract: contract}, nil
}

// NewHeaderSyncTransactor creates a new write-only instance of HeaderSync, bound to a specific deployed contract.
func NewHeaderSyncTransactor(address common.Address, transactor bind.ContractTransactor) (*HeaderSyncTransactor, error) {
	contract, err := bindHeaderSync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderSyncTransactor{contract: contract}, nil
}

// NewHeaderSyncFilterer creates a new log filterer instance of HeaderSync, bound to a specific deployed contract.
func NewHeaderSyncFilterer(address common.Address, filterer bind.ContractFilterer) (*HeaderSyncFilterer, error) {
	contract, err := bindHeaderSync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeaderSyncFilterer{contract: contract}, nil
}

// bindHeaderSync binds a generic wrapper to an already deployed contract.
func bindHeaderSync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeaderSyncABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderSync *HeaderSyncRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HeaderSync.Contract.HeaderSyncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderSync *HeaderSyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderSync.Contract.HeaderSyncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderSync *HeaderSyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderSync.Contract.HeaderSyncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderSync *HeaderSyncCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HeaderSync.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderSync *HeaderSyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderSync.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderSync *HeaderSyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderSync.Contract.contract.Transact(opts, method, params...)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x72ce6700.
//
// Solidity: function syncBlockHeader(uint64 ChainID, address Address, bytes[] Headers) returns(bool success)
func (_HeaderSync *HeaderSyncTransactor) SyncBlockHeader(opts *bind.TransactOpts, ChainID uint64, Address common.Address, Headers [][]byte) (*types.Transaction, error) {
	return _HeaderSync.contract.Transact(opts, "syncBlockHeader", ChainID, Address, Headers)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x72ce6700.
//
// Solidity: function syncBlockHeader(uint64 ChainID, address Address, bytes[] Headers) returns(bool success)
func (_HeaderSync *HeaderSyncSession) SyncBlockHeader(ChainID uint64, Address common.Address, Headers [][]byte) (*types.Transaction, error) {
	return _HeaderSync.Contract.SyncBlockHeader(&_HeaderSync.TransactOpts, ChainID, Address, Headers)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x72ce6700.
//
// Solidity: function syncBlockHeader(uint64 ChainID, address Address, bytes[] Headers) returns(bool success)
func (_HeaderSync *HeaderSyncTransactorSession) SyncBlockHeader(ChainID uint64, Address common.Address, Headers [][]byte) (*types.Transaction, error) {
	return _HeaderSync.Contract.SyncBlockHeader(&_HeaderSync.TransactOpts, ChainID, Address, Headers)
}

// SyncCrossChainMsg is a paid mutator transaction binding the contract method 0x21b5cff5.
//
// Solidity: function syncCrossChainMsg(uint64 ChainID, address Address, bytes[] CrossChainMsgs) returns(bool success)
func (_HeaderSync *HeaderSyncTransactor) SyncCrossChainMsg(opts *bind.TransactOpts, ChainID uint64, Address common.Address, CrossChainMsgs [][]byte) (*types.Transaction, error) {
	return _HeaderSync.contract.Transact(opts, "syncCrossChainMsg", ChainID, Address, CrossChainMsgs)
}

// SyncCrossChainMsg is a paid mutator transaction binding the contract method 0x21b5cff5.
//
// Solidity: function syncCrossChainMsg(uint64 ChainID, address Address, bytes[] CrossChainMsgs) returns(bool success)
func (_HeaderSync *HeaderSyncSession) SyncCrossChainMsg(ChainID uint64, Address common.Address, CrossChainMsgs [][]byte) (*types.Transaction, error) {
	return _HeaderSync.Contract.SyncCrossChainMsg(&_HeaderSync.TransactOpts, ChainID, Address, CrossChainMsgs)
}

// SyncCrossChainMsg is a paid mutator transaction binding the contract method 0x21b5cff5.
//
// Solidity: function syncCrossChainMsg(uint64 ChainID, address Address, bytes[] CrossChainMsgs) returns(bool success)
func (_HeaderSync *HeaderSyncTransactorSession) SyncCrossChainMsg(ChainID uint64, Address common.Address, CrossChainMsgs [][]byte) (*types.Transaction, error) {
	return _HeaderSync.Contract.SyncCrossChainMsg(&_HeaderSync.TransactOpts, ChainID, Address, CrossChainMsgs)
}

// SyncGenesisHeader is a paid mutator transaction binding the contract method 0xb5ace618.
//
// Solidity: function syncGenesisHeader(uint64 ChainID, bytes GenesisHeader) returns(bool success)
func (_HeaderSync *HeaderSyncTransactor) SyncGenesisHeader(opts *bind.TransactOpts, ChainID uint64, GenesisHeader []byte) (*types.Transaction, error) {
	return _HeaderSync.contract.Transact(opts, "syncGenesisHeader", ChainID, GenesisHeader)
}

// SyncGenesisHeader is a paid mutator transaction binding the contract method 0xb5ace618.
//
// Solidity: function syncGenesisHeader(uint64 ChainID, bytes GenesisHeader) returns(bool success)
func (_HeaderSync *HeaderSyncSession) SyncGenesisHeader(ChainID uint64, GenesisHeader []byte) (*types.Transaction, error) {
	return _HeaderSync.Contract.SyncGenesisHeader(&_HeaderSync.TransactOpts, ChainID, GenesisHeader)
}

// SyncGenesisHeader is a paid mutator transaction binding the contract method 0xb5ace618.
//
// Solidity: function syncGenesisHeader(uint64 ChainID, bytes GenesisHeader) returns(bool success)
func (_HeaderSync *HeaderSyncTransactorSession) SyncGenesisHeader(ChainID uint64, GenesisHeader []byte) (*types.Transaction, error) {
	return _HeaderSync.Contract.SyncGenesisHeader(&_HeaderSync.TransactOpts, ChainID, GenesisHeader)
}

// HeaderSyncOKEpochSwitchInfoEventIterator is returned from FilterOKEpochSwitchInfoEvent and is used to iterate over the raw logs and unpacked data for OKEpochSwitchInfoEvent events raised by the HeaderSync contract.
type HeaderSyncOKEpochSwitchInfoEventIterator struct {
	Event *HeaderSyncOKEpochSwitchInfoEvent // Event containing the contract specifics and raw log

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
func (it *HeaderSyncOKEpochSwitchInfoEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeaderSyncOKEpochSwitchInfoEvent)
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
		it.Event = new(HeaderSyncOKEpochSwitchInfoEvent)
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
func (it *HeaderSyncOKEpochSwitchInfoEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeaderSyncOKEpochSwitchInfoEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeaderSyncOKEpochSwitchInfoEvent represents a OKEpochSwitchInfoEvent event raised by the HeaderSync contract.
type HeaderSyncOKEpochSwitchInfoEvent struct {
	ChainID            uint64
	BlockHash          string
	Height             uint64
	NextValidatorsHash string
	InfoChainID        string
	BlockHeight        uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOKEpochSwitchInfoEvent is a free log retrieval operation binding the contract event 0xbfd2d7144ec37c6f85850914f6d172957dd090c508f40d540062f1cd06c0852a.
//
// Solidity: event OKEpochSwitchInfoEvent(uint64 chainID, string BlockHash, uint64 Height, string NextValidatorsHash, string InfoChainID, uint64 BlockHeight)
func (_HeaderSync *HeaderSyncFilterer) FilterOKEpochSwitchInfoEvent(opts *bind.FilterOpts) (*HeaderSyncOKEpochSwitchInfoEventIterator, error) {

	logs, sub, err := _HeaderSync.contract.FilterLogs(opts, "OKEpochSwitchInfoEvent")
	if err != nil {
		return nil, err
	}
	return &HeaderSyncOKEpochSwitchInfoEventIterator{contract: _HeaderSync.contract, event: "OKEpochSwitchInfoEvent", logs: logs, sub: sub}, nil
}

var OKEpochSwitchInfoEventTopicHash = "0xbfd2d7144ec37c6f85850914f6d172957dd090c508f40d540062f1cd06c0852a"

// WatchOKEpochSwitchInfoEvent is a free log subscription operation binding the contract event 0xbfd2d7144ec37c6f85850914f6d172957dd090c508f40d540062f1cd06c0852a.
//
// Solidity: event OKEpochSwitchInfoEvent(uint64 chainID, string BlockHash, uint64 Height, string NextValidatorsHash, string InfoChainID, uint64 BlockHeight)
func (_HeaderSync *HeaderSyncFilterer) WatchOKEpochSwitchInfoEvent(opts *bind.WatchOpts, sink chan<- *HeaderSyncOKEpochSwitchInfoEvent) (event.Subscription, error) {

	logs, sub, err := _HeaderSync.contract.WatchLogs(opts, "OKEpochSwitchInfoEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeaderSyncOKEpochSwitchInfoEvent)
				if err := _HeaderSync.contract.UnpackLog(event, "OKEpochSwitchInfoEvent", log); err != nil {
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

// ParseOKEpochSwitchInfoEvent is a log parse operation binding the contract event 0xbfd2d7144ec37c6f85850914f6d172957dd090c508f40d540062f1cd06c0852a.
//
// Solidity: event OKEpochSwitchInfoEvent(uint64 chainID, string BlockHash, uint64 Height, string NextValidatorsHash, string InfoChainID, uint64 BlockHeight)
func (_HeaderSync *HeaderSyncFilterer) ParseOKEpochSwitchInfoEvent(log types.Log) (*HeaderSyncOKEpochSwitchInfoEvent, error) {
	event := new(HeaderSyncOKEpochSwitchInfoEvent)
	if err := _HeaderSync.contract.UnpackLog(event, "OKEpochSwitchInfoEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HeaderSyncSyncHeaderIterator is returned from FilterSyncHeader and is used to iterate over the raw logs and unpacked data for SyncHeader events raised by the HeaderSync contract.
type HeaderSyncSyncHeaderIterator struct {
	Event *HeaderSyncSyncHeader // Event containing the contract specifics and raw log

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
func (it *HeaderSyncSyncHeaderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeaderSyncSyncHeader)
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
		it.Event = new(HeaderSyncSyncHeader)
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
func (it *HeaderSyncSyncHeaderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeaderSyncSyncHeaderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeaderSyncSyncHeader represents a SyncHeader event raised by the HeaderSync contract.
type HeaderSyncSyncHeader struct {
	ChainID     uint64
	Height      uint64
	BlockHash   string
	BlockHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSyncHeader is a free log retrieval operation binding the contract event 0xe4d5dbebcfbd7358435d9d612bc7b584bc51faf456160631d2537a0de2e1a236.
//
// Solidity: event syncHeader(uint64 chainID, uint64 height, string blockHash, uint256 BlockHeight)
func (_HeaderSync *HeaderSyncFilterer) FilterSyncHeader(opts *bind.FilterOpts) (*HeaderSyncSyncHeaderIterator, error) {

	logs, sub, err := _HeaderSync.contract.FilterLogs(opts, "syncHeader")
	if err != nil {
		return nil, err
	}
	return &HeaderSyncSyncHeaderIterator{contract: _HeaderSync.contract, event: "syncHeader", logs: logs, sub: sub}, nil
}

var SyncHeaderTopicHash = "0xe4d5dbebcfbd7358435d9d612bc7b584bc51faf456160631d2537a0de2e1a236"

// WatchSyncHeader is a free log subscription operation binding the contract event 0xe4d5dbebcfbd7358435d9d612bc7b584bc51faf456160631d2537a0de2e1a236.
//
// Solidity: event syncHeader(uint64 chainID, uint64 height, string blockHash, uint256 BlockHeight)
func (_HeaderSync *HeaderSyncFilterer) WatchSyncHeader(opts *bind.WatchOpts, sink chan<- *HeaderSyncSyncHeader) (event.Subscription, error) {

	logs, sub, err := _HeaderSync.contract.WatchLogs(opts, "syncHeader")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeaderSyncSyncHeader)
				if err := _HeaderSync.contract.UnpackLog(event, "syncHeader", log); err != nil {
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

// ParseSyncHeader is a log parse operation binding the contract event 0xe4d5dbebcfbd7358435d9d612bc7b584bc51faf456160631d2537a0de2e1a236.
//
// Solidity: event syncHeader(uint64 chainID, uint64 height, string blockHash, uint256 BlockHeight)
func (_HeaderSync *HeaderSyncFilterer) ParseSyncHeader(log types.Log) (*HeaderSyncSyncHeader, error) {
	event := new(HeaderSyncSyncHeader)
	if err := _HeaderSync.contract.UnpackLog(event, "syncHeader", log); err != nil {
		return nil, err
	}
	return event, nil
}

