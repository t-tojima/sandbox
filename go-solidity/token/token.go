// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x6060604052341561000f57600080fd5b6102828061001e6000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd8114610066578063313ce5671461008b57806370a08231146100b4578063a9059cbb146100d3575b600080fd5b341561007157600080fd5b610079610109565b60405190815260200160405180910390f35b341561009657600080fd5b61009e61010f565b60405160ff909116815260200160405180910390f35b34156100bf57600080fd5b610079600160a060020a0360043516610118565b34156100de57600080fd5b6100f5600160a060020a0360043516602435610133565b604051901515815260200160405180910390f35b60005481565b60015460ff1681565b600160a060020a031660009081526002602052604090205490565b6000600160a060020a038316151561014a57600080fd5b600160a060020a03331660009081526002602052604090205482111561016f57600080fd5b600160a060020a033316600090815260026020526040902054610198908363ffffffff61022e16565b600160a060020a0333811660009081526002602052604080822093909355908516815220546101cd908363ffffffff61024016565b600160a060020a0380851660008181526002602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b60008282111561023a57fe5b50900390565b60008282018381101561024f57fe5b93925050505600a165627a7a72305820a7a8dfd734eb0560cbfbdbefef11d275dfa635e6bc630432db0b6a1ad99d85830029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BasicToken *BasicTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BasicToken *BasicTokenSession) Decimals() (uint8, error) {
	return _BasicToken.Contract.Decimals(&_BasicToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BasicToken *BasicTokenCallerSession) Decimals() (uint8, error) {
	return _BasicToken.Contract.Decimals(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ContractReceiverABI is the input ABI used to generate the binding from.
const ContractReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractReceiverBin is the compiled bytecode used for deploying new contracts.
const ContractReceiverBin = `0x`

// DeployContractReceiver deploys a new Ethereum contract, binding an instance of ContractReceiver to it.
func DeployContractReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractReceiver, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractReceiver{ContractReceiverCaller: ContractReceiverCaller{contract: contract}, ContractReceiverTransactor: ContractReceiverTransactor{contract: contract}, ContractReceiverFilterer: ContractReceiverFilterer{contract: contract}}, nil
}

// ContractReceiver is an auto generated Go binding around an Ethereum contract.
type ContractReceiver struct {
	ContractReceiverCaller     // Read-only binding to the contract
	ContractReceiverTransactor // Write-only binding to the contract
	ContractReceiverFilterer   // Log filterer for contract events
}

// ContractReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractReceiverSession struct {
	Contract     *ContractReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractReceiverCallerSession struct {
	Contract *ContractReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractReceiverTransactorSession struct {
	Contract     *ContractReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractReceiverRaw struct {
	Contract *ContractReceiver // Generic contract binding to access the raw methods on
}

// ContractReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractReceiverCallerRaw struct {
	Contract *ContractReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ContractReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractReceiverTransactorRaw struct {
	Contract *ContractReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractReceiver creates a new instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiver(address common.Address, backend bind.ContractBackend) (*ContractReceiver, error) {
	contract, err := bindContractReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractReceiver{ContractReceiverCaller: ContractReceiverCaller{contract: contract}, ContractReceiverTransactor: ContractReceiverTransactor{contract: contract}, ContractReceiverFilterer: ContractReceiverFilterer{contract: contract}}, nil
}

// NewContractReceiverCaller creates a new read-only instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverCaller(address common.Address, caller bind.ContractCaller) (*ContractReceiverCaller, error) {
	contract, err := bindContractReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverCaller{contract: contract}, nil
}

// NewContractReceiverTransactor creates a new write-only instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractReceiverTransactor, error) {
	contract, err := bindContractReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverTransactor{contract: contract}, nil
}

// NewContractReceiverFilterer creates a new log filterer instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractReceiverFilterer, error) {
	contract, err := bindContractReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverFilterer{contract: contract}, nil
}

// bindContractReceiver binds a generic wrapper to an already deployed contract.
func bindContractReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractReceiver *ContractReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractReceiver.Contract.ContractReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractReceiver *ContractReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReceiver.Contract.ContractReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractReceiver *ContractReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractReceiver.Contract.ContractReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractReceiver *ContractReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractReceiver *ContractReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractReceiver *ContractReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractReceiver.Contract.contract.Transact(opts, method, params...)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ContractReceiver *ContractReceiverTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractReceiver.contract.Transact(opts, "tokenFallback", _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ContractReceiver *ContractReceiverSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractReceiver.Contract.TokenFallback(&_ContractReceiver.TransactOpts, _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ContractReceiver *ContractReceiverTransactorSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractReceiver.Contract.TokenFallback(&_ContractReceiver.TransactOpts, _from, _value, _data)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Basic *ERC20BasicCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Basic *ERC20BasicSession) Decimals() (uint8, error) {
	return _ERC20Basic.Contract.Decimals(&_ERC20Basic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Basic *ERC20BasicCallerSession) Decimals() (uint8, error) {
	return _ERC20Basic.Contract.Decimals(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC223ABI is the input ABI used to generate the binding from.
const ERC223ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC223Transfer\",\"type\":\"event\"}]"

// ERC223Bin is the compiled bytecode used for deploying new contracts.
const ERC223Bin = `0x`

// DeployERC223 deploys a new Ethereum contract, binding an instance of ERC223 to it.
func DeployERC223(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC223, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC223Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC223{ERC223Caller: ERC223Caller{contract: contract}, ERC223Transactor: ERC223Transactor{contract: contract}, ERC223Filterer: ERC223Filterer{contract: contract}}, nil
}

// ERC223 is an auto generated Go binding around an Ethereum contract.
type ERC223 struct {
	ERC223Caller     // Read-only binding to the contract
	ERC223Transactor // Write-only binding to the contract
	ERC223Filterer   // Log filterer for contract events
}

// ERC223Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC223Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC223Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC223Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC223Session struct {
	Contract     *ERC223           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC223CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC223CallerSession struct {
	Contract *ERC223Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC223TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC223TransactorSession struct {
	Contract     *ERC223Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC223Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC223Raw struct {
	Contract *ERC223 // Generic contract binding to access the raw methods on
}

// ERC223CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC223CallerRaw struct {
	Contract *ERC223Caller // Generic read-only contract binding to access the raw methods on
}

// ERC223TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC223TransactorRaw struct {
	Contract *ERC223Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC223 creates a new instance of ERC223, bound to a specific deployed contract.
func NewERC223(address common.Address, backend bind.ContractBackend) (*ERC223, error) {
	contract, err := bindERC223(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC223{ERC223Caller: ERC223Caller{contract: contract}, ERC223Transactor: ERC223Transactor{contract: contract}, ERC223Filterer: ERC223Filterer{contract: contract}}, nil
}

// NewERC223Caller creates a new read-only instance of ERC223, bound to a specific deployed contract.
func NewERC223Caller(address common.Address, caller bind.ContractCaller) (*ERC223Caller, error) {
	contract, err := bindERC223(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223Caller{contract: contract}, nil
}

// NewERC223Transactor creates a new write-only instance of ERC223, bound to a specific deployed contract.
func NewERC223Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC223Transactor, error) {
	contract, err := bindERC223(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223Transactor{contract: contract}, nil
}

// NewERC223Filterer creates a new log filterer instance of ERC223, bound to a specific deployed contract.
func NewERC223Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC223Filterer, error) {
	contract, err := bindERC223(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC223Filterer{contract: contract}, nil
}

// bindERC223 binds a generic wrapper to an already deployed contract.
func bindERC223(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223 *ERC223Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223.Contract.ERC223Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223 *ERC223Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223.Contract.ERC223Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223 *ERC223Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223.Contract.ERC223Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223 *ERC223CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223 *ERC223TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223 *ERC223TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC223 *ERC223Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC223 *ERC223Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC223.Contract.BalanceOf(&_ERC223.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC223 *ERC223CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC223.Contract.BalanceOf(&_ERC223.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_ERC223 *ERC223Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_ERC223 *ERC223Session) Decimals() (uint8, error) {
	return _ERC223.Contract.Decimals(&_ERC223.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_ERC223 *ERC223CallerSession) Decimals() (uint8, error) {
	return _ERC223.Contract.Decimals(&_ERC223.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC223 *ERC223Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC223 *ERC223Session) Name() (string, error) {
	return _ERC223.Contract.Name(&_ERC223.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC223 *ERC223CallerSession) Name() (string, error) {
	return _ERC223.Contract.Name(&_ERC223.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC223 *ERC223Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC223 *ERC223Session) Symbol() (string, error) {
	return _ERC223.Contract.Symbol(&_ERC223.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC223 *ERC223CallerSession) Symbol() (string, error) {
	return _ERC223.Contract.Symbol(&_ERC223.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_supply uint256)
func (_ERC223 *ERC223Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_supply uint256)
func (_ERC223 *ERC223Session) TotalSupply() (*big.Int, error) {
	return _ERC223.Contract.TotalSupply(&_ERC223.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_supply uint256)
func (_ERC223 *ERC223CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC223.Contract.TotalSupply(&_ERC223.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(to address, value uint256, data bytes) returns(ok bool)
func (_ERC223 *ERC223Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "transfer", to, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(to address, value uint256, data bytes) returns(ok bool)
func (_ERC223 *ERC223Session) Transfer(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223.Contract.Transfer(&_ERC223.TransactOpts, to, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(to address, value uint256, data bytes) returns(ok bool)
func (_ERC223 *ERC223TransactorSession) Transfer(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223.Contract.Transfer(&_ERC223.TransactOpts, to, value, data)
}

// ERC223ERC223TransferIterator is returned from FilterERC223Transfer and is used to iterate over the raw logs and unpacked data for ERC223Transfer events raised by the ERC223 contract.
type ERC223ERC223TransferIterator struct {
	Event *ERC223ERC223Transfer // Event containing the contract specifics and raw log

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
func (it *ERC223ERC223TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC223ERC223Transfer)
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
		it.Event = new(ERC223ERC223Transfer)
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
func (it *ERC223ERC223TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC223ERC223TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC223ERC223Transfer represents a ERC223Transfer event raised by the ERC223 contract.
type ERC223ERC223Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC223Transfer is a free log retrieval operation binding the contract event 0x9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1.
//
// Solidity: event ERC223Transfer(_from indexed address, _to indexed address, _value uint256, _data bytes)
func (_ERC223 *ERC223Filterer) FilterERC223Transfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC223ERC223TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC223.contract.FilterLogs(opts, "ERC223Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC223ERC223TransferIterator{contract: _ERC223.contract, event: "ERC223Transfer", logs: logs, sub: sub}, nil
}

// WatchERC223Transfer is a free log subscription operation binding the contract event 0x9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1.
//
// Solidity: event ERC223Transfer(_from indexed address, _to indexed address, _value uint256, _data bytes)
func (_ERC223 *ERC223Filterer) WatchERC223Transfer(opts *bind.WatchOpts, sink chan<- *ERC223ERC223Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC223.contract.WatchLogs(opts, "ERC223Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC223ERC223Transfer)
				if err := _ERC223.contract.UnpackLog(event, "ERC223Transfer", log); err != nil {
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

// ERC223TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC223 contract.
type ERC223TransferIterator struct {
	Event *ERC223Transfer // Event containing the contract specifics and raw log

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
func (it *ERC223TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC223Transfer)
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
		it.Event = new(ERC223Transfer)
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
func (it *ERC223TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC223TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC223Transfer represents a Transfer event raised by the ERC223 contract.
type ERC223Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC223 *ERC223Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC223TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC223.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC223TransferIterator{contract: _ERC223.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC223 *ERC223Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC223Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC223.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC223Transfer)
				if err := _ERC223.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC223TokenABI is the input ABI used to generate the binding from.
const ERC223TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC223Transfer\",\"type\":\"event\"}]"

// ERC223TokenBin is the compiled bytecode used for deploying new contracts.
const ERC223TokenBin = `0x6060604052341561000f57600080fd5b61083c8061001e6000396000f3006060604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461008757806318160ddd14610111578063313ce5671461013657806370a082311461015f57806395d89b411461017e578063a9059cbb14610191578063be45fd62146101c7575b600080fd5b341561009257600080fd5b61009a61022c565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100d65780820151838201526020016100be565b50505050905090810190601f1680156101035780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561011c57600080fd5b6101246102d4565b60405190815260200160405180910390f35b341561014157600080fd5b6101496102da565b60405160ff909116815260200160405180910390f35b341561016a57600080fd5b610124600160a060020a03600435166102e3565b341561018957600080fd5b61009a6102fe565b341561019c57600080fd5b6101b3600160a060020a0360043516602435610371565b604051901515815260200160405180910390f35b34156101d257600080fd5b6101b360048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506103ad95505050505050565b6102346107fe565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102ca5780601f1061029f576101008083540402835291602001916102ca565b820191906000526020600020905b8154815290600101906020018083116102ad57829003601f168201915b5050505050905090565b60055490565b60045460ff1690565b600160a060020a031660009081526001602052604090205490565b6103066107fe565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102ca5780601f1061029f576101008083540402835291602001916102ca565b600061037b6107fe565b610384846103e1565b1561039b57610394848483610401565b91506103a6565b61039484848361067d565b5092915050565b60006103b8846103e1565b156103cf576103c8848484610401565b90506103da565b6103c884848461067d565b9392505050565b6000813b818111156103f657600191506103fb565b600091505b50919050565b6000808361040e336102e3565b101561041957600080fd5b61043284610426336102e3565b9063ffffffff6107dd16565b600160a060020a03331660009081526001602052604090205561046484610458876102e3565b9063ffffffff6107ef16565b600160a060020a03861660008181526001602052604090819020929092558692509063c0ee0b8a90339087908790518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156105165780820151838201526020016104fe565b50505050905090810190601f1680156105435780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b151561056357600080fd5b6102c65a03f1151561057457600080fd5b50505084600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8660405190815260200160405180910390a384600160a060020a031633600160a060020a03167f9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1868660405182815260406020820181815290820183818151815260200191508051906020019080838360005b8381101561063757808201518382015260200161061f565b50505050905090810190601f1680156106645780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3506001949350505050565b600082610689336102e3565b101561069457600080fd5b6106a183610426336102e3565b600160a060020a0333166000908152600160205260409020556106c783610458866102e3565b600160a060020a0380861660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a383600160a060020a031633600160a060020a03167f9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1858560405182815260406020820181815290820183818151815260200191508051906020019080838360005b83811015610798578082015183820152602001610780565b50505050905090810190601f1680156107c55780820380516001836020036101000a031916815260200191505b50935050505060405180910390a35060019392505050565b6000828211156107e957fe5b50900390565b6000828201838110156103da57fe5b602060405190810160405260008152905600a165627a7a723058209ff2c7bc3b539adca14cbe50472998dcc3aa9178c106328411e9107fe9a6fa190029`

// DeployERC223Token deploys a new Ethereum contract, binding an instance of ERC223Token to it.
func DeployERC223Token(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC223Token, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC223TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC223Token{ERC223TokenCaller: ERC223TokenCaller{contract: contract}, ERC223TokenTransactor: ERC223TokenTransactor{contract: contract}, ERC223TokenFilterer: ERC223TokenFilterer{contract: contract}}, nil
}

// ERC223Token is an auto generated Go binding around an Ethereum contract.
type ERC223Token struct {
	ERC223TokenCaller     // Read-only binding to the contract
	ERC223TokenTransactor // Write-only binding to the contract
	ERC223TokenFilterer   // Log filterer for contract events
}

// ERC223TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC223TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC223TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC223TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC223TokenSession struct {
	Contract     *ERC223Token      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC223TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC223TokenCallerSession struct {
	Contract *ERC223TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC223TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC223TokenTransactorSession struct {
	Contract     *ERC223TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC223TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC223TokenRaw struct {
	Contract *ERC223Token // Generic contract binding to access the raw methods on
}

// ERC223TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC223TokenCallerRaw struct {
	Contract *ERC223TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC223TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC223TokenTransactorRaw struct {
	Contract *ERC223TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC223Token creates a new instance of ERC223Token, bound to a specific deployed contract.
func NewERC223Token(address common.Address, backend bind.ContractBackend) (*ERC223Token, error) {
	contract, err := bindERC223Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC223Token{ERC223TokenCaller: ERC223TokenCaller{contract: contract}, ERC223TokenTransactor: ERC223TokenTransactor{contract: contract}, ERC223TokenFilterer: ERC223TokenFilterer{contract: contract}}, nil
}

// NewERC223TokenCaller creates a new read-only instance of ERC223Token, bound to a specific deployed contract.
func NewERC223TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC223TokenCaller, error) {
	contract, err := bindERC223Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenCaller{contract: contract}, nil
}

// NewERC223TokenTransactor creates a new write-only instance of ERC223Token, bound to a specific deployed contract.
func NewERC223TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC223TokenTransactor, error) {
	contract, err := bindERC223Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenTransactor{contract: contract}, nil
}

// NewERC223TokenFilterer creates a new log filterer instance of ERC223Token, bound to a specific deployed contract.
func NewERC223TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC223TokenFilterer, error) {
	contract, err := bindERC223Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenFilterer{contract: contract}, nil
}

// bindERC223Token binds a generic wrapper to an already deployed contract.
func bindERC223Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223Token *ERC223TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223Token.Contract.ERC223TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223Token *ERC223TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223Token.Contract.ERC223TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223Token *ERC223TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223Token.Contract.ERC223TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223Token *ERC223TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223Token *ERC223TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223Token *ERC223TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC223Token *ERC223TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC223Token *ERC223TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC223Token.Contract.BalanceOf(&_ERC223Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC223Token *ERC223TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC223Token.Contract.BalanceOf(&_ERC223Token.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_ERC223Token *ERC223TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC223Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_ERC223Token *ERC223TokenSession) Decimals() (uint8, error) {
	return _ERC223Token.Contract.Decimals(&_ERC223Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_ERC223Token *ERC223TokenCallerSession) Decimals() (uint8, error) {
	return _ERC223Token.Contract.Decimals(&_ERC223Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC223Token *ERC223TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC223Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC223Token *ERC223TokenSession) Name() (string, error) {
	return _ERC223Token.Contract.Name(&_ERC223Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC223Token *ERC223TokenCallerSession) Name() (string, error) {
	return _ERC223Token.Contract.Name(&_ERC223Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC223Token *ERC223TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC223Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC223Token *ERC223TokenSession) Symbol() (string, error) {
	return _ERC223Token.Contract.Symbol(&_ERC223Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC223Token *ERC223TokenCallerSession) Symbol() (string, error) {
	return _ERC223Token.Contract.Symbol(&_ERC223Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_ERC223Token *ERC223TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_ERC223Token *ERC223TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC223Token.Contract.TotalSupply(&_ERC223Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_ERC223Token *ERC223TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC223Token.Contract.TotalSupply(&_ERC223Token.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC223Token *ERC223TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223Token.contract.Transact(opts, "transfer", _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC223Token *ERC223TokenSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223Token.Contract.Transfer(&_ERC223Token.TransactOpts, _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC223Token *ERC223TokenTransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223Token.Contract.Transfer(&_ERC223Token.TransactOpts, _to, _value, _data)
}

// ERC223TokenERC223TransferIterator is returned from FilterERC223Transfer and is used to iterate over the raw logs and unpacked data for ERC223Transfer events raised by the ERC223Token contract.
type ERC223TokenERC223TransferIterator struct {
	Event *ERC223TokenERC223Transfer // Event containing the contract specifics and raw log

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
func (it *ERC223TokenERC223TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC223TokenERC223Transfer)
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
		it.Event = new(ERC223TokenERC223Transfer)
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
func (it *ERC223TokenERC223TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC223TokenERC223TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC223TokenERC223Transfer represents a ERC223Transfer event raised by the ERC223Token contract.
type ERC223TokenERC223Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC223Transfer is a free log retrieval operation binding the contract event 0x9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1.
//
// Solidity: event ERC223Transfer(_from indexed address, _to indexed address, _value uint256, _data bytes)
func (_ERC223Token *ERC223TokenFilterer) FilterERC223Transfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC223TokenERC223TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC223Token.contract.FilterLogs(opts, "ERC223Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenERC223TransferIterator{contract: _ERC223Token.contract, event: "ERC223Transfer", logs: logs, sub: sub}, nil
}

// WatchERC223Transfer is a free log subscription operation binding the contract event 0x9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1.
//
// Solidity: event ERC223Transfer(_from indexed address, _to indexed address, _value uint256, _data bytes)
func (_ERC223Token *ERC223TokenFilterer) WatchERC223Transfer(opts *bind.WatchOpts, sink chan<- *ERC223TokenERC223Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC223Token.contract.WatchLogs(opts, "ERC223Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC223TokenERC223Transfer)
				if err := _ERC223Token.contract.UnpackLog(event, "ERC223Transfer", log); err != nil {
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

// ERC223TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC223Token contract.
type ERC223TokenTransferIterator struct {
	Event *ERC223TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC223TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC223TokenTransfer)
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
		it.Event = new(ERC223TokenTransfer)
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
func (it *ERC223TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC223TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC223TokenTransfer represents a Transfer event raised by the ERC223Token contract.
type ERC223TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC223Token *ERC223TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC223TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC223Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenTransferIterator{contract: _ERC223Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC223Token *ERC223TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC223TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC223Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC223TokenTransfer)
				if err := _ERC223Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// MyTokenABI is the input ABI used to generate the binding from.
const MyTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC223Transfer\",\"type\":\"event\"}]"

// MyTokenBin is the compiled bytecode used for deploying new contracts.
const MyTokenBin = `0x606060405260408051908101604052600781527f4d79546f6b656e000000000000000000000000000000000000000000000000006020820152600690805161004b9291602001906100d3565b5060408051908101604052600381527f4d544b0000000000000000000000000000000000000000000000000000000000602082015260079080516100939291602001906100d3565b5060066008556402540be40060095534156100ad57600080fd5b6009546005819055600160a060020a03331660009081526001602052604090205561016e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011457805160ff1916838001178555610141565b82800160010185558215610141579182015b82811115610141578251825591602001919060010190610126565b5061014d929150610151565b5090565b61016b91905b8082111561014d5760008155600101610157565b90565b6108358061017d6000396000f30060606040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461009257806318160ddd1461011c5780632ff2e9dc14610141578063313ce5671461015457806370a082311461016757806395d89b4114610186578063a9059cbb14610199578063be45fd62146101cf575b600080fd5b341561009d57600080fd5b6100a5610234565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100e15780820151838201526020016100c9565b50505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561012757600080fd5b61012f6102d2565b60405190815260200160405180910390f35b341561014c57600080fd5b61012f6102d8565b341561015f57600080fd5b61012f6102de565b341561017257600080fd5b61012f600160a060020a03600435166102e4565b341561019157600080fd5b6100a56102ff565b34156101a457600080fd5b6101bb600160a060020a036004351660243561036a565b604051901515815260200160405180910390f35b34156101da57600080fd5b6101bb60048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506103a695505050505050565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102ca5780601f1061029f576101008083540402835291602001916102ca565b820191906000526020600020905b8154815290600101906020018083116102ad57829003601f168201915b505050505081565b60055490565b60095481565b60085481565b600160a060020a031660009081526001602052604090205490565b60078054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102ca5780601f1061029f576101008083540402835291602001916102ca565b60006103746107f7565b61037d846103da565b156103945761038d8484836103fa565b915061039f565b61038d848483610676565b5092915050565b60006103b1846103da565b156103c8576103c18484846103fa565b90506103d3565b6103c1848484610676565b9392505050565b6000813b818111156103ef57600191506103f4565b600091505b50919050565b60008083610407336102e4565b101561041257600080fd5b61042b8461041f336102e4565b9063ffffffff6107d616565b600160a060020a03331660009081526001602052604090205561045d84610451876102e4565b9063ffffffff6107e816565b600160a060020a03861660008181526001602052604090819020929092558692509063c0ee0b8a90339087908790518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561050f5780820151838201526020016104f7565b50505050905090810190601f16801561053c5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b151561055c57600080fd5b6102c65a03f1151561056d57600080fd5b50505084600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8660405190815260200160405180910390a384600160a060020a031633600160a060020a03167f9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1868660405182815260406020820181815290820183818151815260200191508051906020019080838360005b83811015610630578082015183820152602001610618565b50505050905090810190601f16801561065d5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3506001949350505050565b600082610682336102e4565b101561068d57600080fd5b61069a8361041f336102e4565b600160a060020a0333166000908152600160205260409020556106c083610451866102e4565b600160a060020a0380861660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a383600160a060020a031633600160a060020a03167f9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1858560405182815260406020820181815290820183818151815260200191508051906020019080838360005b83811015610791578082015183820152602001610779565b50505050905090810190601f1680156107be5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a35060019392505050565b6000828211156107e257fe5b50900390565b6000828201838110156103d357fe5b602060405190810160405260008152905600a165627a7a72305820aeee72c11d6b293d0a9834f935e837d205440ac62d1aa67952fa5def3146aaf20029`

// DeployMyToken deploys a new Ethereum contract, binding an instance of MyToken to it.
func DeployMyToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}, MyTokenFilterer: MyTokenFilterer{contract: contract}}, nil
}

// MyToken is an auto generated Go binding around an Ethereum contract.
type MyToken struct {
	MyTokenCaller     // Read-only binding to the contract
	MyTokenTransactor // Write-only binding to the contract
	MyTokenFilterer   // Log filterer for contract events
}

// MyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyTokenSession struct {
	Contract     *MyToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyTokenCallerSession struct {
	Contract *MyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyTokenTransactorSession struct {
	Contract     *MyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyTokenRaw struct {
	Contract *MyToken // Generic contract binding to access the raw methods on
}

// MyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyTokenCallerRaw struct {
	Contract *MyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyTokenTransactorRaw struct {
	Contract *MyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyToken creates a new instance of MyToken, bound to a specific deployed contract.
func NewMyToken(address common.Address, backend bind.ContractBackend) (*MyToken, error) {
	contract, err := bindMyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}, MyTokenFilterer: MyTokenFilterer{contract: contract}}, nil
}

// NewMyTokenCaller creates a new read-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenCaller(address common.Address, caller bind.ContractCaller) (*MyTokenCaller, error) {
	contract, err := bindMyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenCaller{contract: contract}, nil
}

// NewMyTokenTransactor creates a new write-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MyTokenTransactor, error) {
	contract, err := bindMyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransactor{contract: contract}, nil
}

// NewMyTokenFilterer creates a new log filterer instance of MyToken, bound to a specific deployed contract.
func NewMyTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MyTokenFilterer, error) {
	contract, err := bindMyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyTokenFilterer{contract: contract}, nil
}

// bindMyToken binds a generic wrapper to an already deployed contract.
func bindMyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.MyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyToken.Contract.contract.Transact(opts, method, params...)
}

// INITIAL_SUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_MyToken *MyTokenCaller) INITIAL_SUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIAL_SUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_MyToken *MyTokenSession) INITIAL_SUPPLY() (*big.Int, error) {
	return _MyToken.Contract.INITIAL_SUPPLY(&_MyToken.CallOpts)
}

// INITIAL_SUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_MyToken *MyTokenCallerSession) INITIAL_SUPPLY() (*big.Int, error) {
	return _MyToken.Contract.INITIAL_SUPPLY(&_MyToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MyToken *MyTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MyToken *MyTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MyToken *MyTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_MyToken *MyTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_MyToken *MyTokenSession) Decimals() (*big.Int, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_MyToken *MyTokenCallerSession) Decimals() (*big.Int, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenCallerSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenCallerSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MyToken *MyTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MyToken *MyTokenSession) TotalSupply() (*big.Int, error) {
	return _MyToken.Contract.TotalSupply(&_MyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MyToken *MyTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MyToken.Contract.TotalSupply(&_MyToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(success bool)
func (_MyToken *MyTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "transfer", _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(success bool)
func (_MyToken *MyTokenSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(success bool)
func (_MyToken *MyTokenTransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, _to, _value, _data)
}

// MyTokenERC223TransferIterator is returned from FilterERC223Transfer and is used to iterate over the raw logs and unpacked data for ERC223Transfer events raised by the MyToken contract.
type MyTokenERC223TransferIterator struct {
	Event *MyTokenERC223Transfer // Event containing the contract specifics and raw log

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
func (it *MyTokenERC223TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenERC223Transfer)
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
		it.Event = new(MyTokenERC223Transfer)
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
func (it *MyTokenERC223TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenERC223TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenERC223Transfer represents a ERC223Transfer event raised by the MyToken contract.
type MyTokenERC223Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC223Transfer is a free log retrieval operation binding the contract event 0x9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1.
//
// Solidity: event ERC223Transfer(_from indexed address, _to indexed address, _value uint256, _data bytes)
func (_MyToken *MyTokenFilterer) FilterERC223Transfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*MyTokenERC223TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _MyToken.contract.FilterLogs(opts, "ERC223Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenERC223TransferIterator{contract: _MyToken.contract, event: "ERC223Transfer", logs: logs, sub: sub}, nil
}

// WatchERC223Transfer is a free log subscription operation binding the contract event 0x9bfafdc2ae8835972d7b64ef3f8f307165ac22ceffde4a742c52da5487f45fd1.
//
// Solidity: event ERC223Transfer(_from indexed address, _to indexed address, _value uint256, _data bytes)
func (_MyToken *MyTokenFilterer) WatchERC223Transfer(opts *bind.WatchOpts, sink chan<- *MyTokenERC223Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "ERC223Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenERC223Transfer)
				if err := _MyToken.contract.UnpackLog(event, "ERC223Transfer", log); err != nil {
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

// MyTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyToken contract.
type MyTokenTransferIterator struct {
	Event *MyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenTransfer)
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
		it.Event = new(MyTokenTransfer)
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
func (it *MyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenTransfer represents a Transfer event raised by the MyToken contract.
type MyTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_MyToken *MyTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*MyTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _MyToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransferIterator{contract: _MyToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_MyToken *MyTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenTransfer)
				if err := _MyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// RadexABI is the input ABI used to generate the binding from.
const RadexABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sellToken\",\"type\":\"address\"},{\"name\":\"buyToken\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"priceMul\",\"type\":\"uint256\"},{\"name\":\"priceDiv\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"erc20token\",\"type\":\"address\"},{\"name\":\"erc223token\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"commitments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"commitmentsOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeMultiplier\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sellToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_buyToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_priceMul\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_priceDiv\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"NewOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"MarketMaker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_soldTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_boughtTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"}]"

// RadexBin is the compiled bytecode used for deploying new contracts.
const RadexBin = `0x606060405260028054600160a060020a03191690556000600655341561002457600080fd5b600160008190558054600160a060020a033316600160a060020a0319909116179055610f7b806100556000396000f3006060604052600436106100ab5763ffffffff60e060020a6000350416631e9a695081146100b0578063514fcac7146100d4578063518e62db146100ea5780636f3b60d61461012a578063aa6773541461013d578063adb6f63f14610162578063b60d428814610187578063c0ee0b8a1461018f578063c23f001f146101f4578063cc7e930c14610219578063e5a70ef71461023e578063ef46e0ca14610251578063f7888aec1461026a575b600080fd5b34156100bb57600080fd5b6100d2600160a060020a036004351660243561028f565b005b34156100df57600080fd5b6100d260043561043e565b34156100f557600080fd5b610118600160a060020a0360043581169060243516604435606435608435610579565b60405190815260200160405180910390f35b341561013557600080fd5b610118610852565b341561014857600080fd5b6100d2600160a060020a0360043581169060243516610858565b341561016d57600080fd5b610118600160a060020a036004358116906024351661095b565b6100d2610978565b341561019a57600080fd5b6100d260048035600160a060020a03169060248035919060649060443590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610a2395505050505050565b34156101ff57600080fd5b610118600160a060020a0360043581169060243516610aca565b341561022457600080fd5b610118600160a060020a0360043581169060243516610ae7565b341561024957600080fd5b610118610b13565b341561025c57600080fd5b6100d2600435602435610b19565b341561027557600080fd5b610118600160a060020a0360043581169060243516610ec0565b600081151561029d57600080fd5b5033600160a060020a038181166000908152600360209081526040808320938716835292905220548211156102d157600080fd5b600160a060020a03808216600090815260036020908152604080832093871683529290522054610307908363ffffffff610eec16565b600160a060020a03808316600090815260036020908152604080832088851680855292529091209290925560025416141561037257600160a060020a03811682156108fc0283604051600060405180830381858888f19350505050151561036d57600080fd5b6103eb565b82600160a060020a031663a9059cbb828460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156103cf57600080fd5b6102c65a03f115156103e057600080fd5b505050604051805150505b33600160a060020a031683600160a060020a03167fc2b4a290c20fb28939d29f102514fbffd2b73c059ffba8b78250c94161d5fcc6844260405191825260208201526040908101905180910390a3505050565b60008181526005602052604090206003810154151561045c57600080fd5b805433600160a060020a0390811691161461047657600080fd5b6003810154600160a060020a0333811660009081526004602090815260408083206001870154909416835292905220546104b59163ffffffff610eec16565b600160a060020a0333811660008181526004602090815260408083206001880180548716855290835281842096909655600380880154948452825280832095549094168252939093529120546105109163ffffffff610efe16565b600160a060020a033381166000908152600360209081526040808320600187015490941683529290528190209190915582907faa2a006d186dc89a3d6851834c85d16226329ac40bcad6531a5756a38ed1c90f9042905190815260200160405180910390a25050565b600083151561058757600080fd5b82151561059357600080fd5b81151561059f57600080fd5b84600160a060020a031686600160a060020a031614156105be57600080fd5b600160a060020a033381166000908152600360209081526040808320938a1683529290522054849010156105f157600080fd5b61061182610605868663ffffffff610f1416565b9063ffffffff610f3816565b151561061c57600080fd5b50600680546001810190915560c06040519081016040528033600160a060020a0316815260200187600160a060020a0316815260200186600160a060020a0316815260200185815260200184815260200183815250600560008381526020019081526020016000206000820151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055606082015181600301556080820151816004015560a082015160059091015550600160a060020a033381166000908152600360209081526040808320938a168352929052205461076f908563ffffffff610eec16565b600160a060020a033381166000818152600360209081526040808320948c16808452948252808320959095559181526004825283812092815291905220546107bd908563ffffffff610efe16565b600160a060020a0333811660008181526004602090815260408083208c86168085529252918290209490945591881692917f506cb18801692179462f2b7b4b6f10bea763984f94a1f1f3cd858e3a645b2329908590899089908990429051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a495945050505050565b60065481565b600154600090819033600160a060020a0390811691161461087857600080fd5b83915081600160a060020a03166318160ddd6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156108c157600080fd5b6102c65a03f115156108d257600080fd5b5050506040518051915050600160a060020a03821663095ea7b3848360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561093a57600080fd5b6102c65a03f1151561094b57600080fd5b5050506040518051505050505050565b600460209081526000928352604080842090915290825290205481565b600160a060020a033381166000908152600360209081526040808320600254909416835292905220546109b1903463ffffffff610efe16565b600160a060020a033381166000818152600360209081526040808320600280548716855292529182902094909455925490929116907fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d790349042905191825260208201526040908101905180910390a3565b600160a060020a0380841660009081526003602090815260408083203390941683529290522054610a5a908363ffffffff610efe16565b600160a060020a03808516600081815260036020908152604080832033909516808452949091529081902093909355917fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d790859042905191825260208201526040908101905180910390a3505050565b600360209081526000928352604080842090915290825290205481565b600160a060020a0380821660009081526004602090815260408083209386168352929052205492915050565b60005481565b6000806000600654851115610b2d57600080fd5b600560008681526020019081526020016000209250610b618360050154610605856004015487610f1490919063ffffffff16565b9150831515610b6f57600080fd5b8383600301541015610b8057600080fd5b825433600160a060020a0390811691161415610b9b57600080fd5b600160a060020a03338116600090815260036020908152604080832060028801549094168352929052205482901015610bd357600080fd5b600054610be790859063ffffffff610f3816565b8354600160a060020a039081166000908152600360209081526040808320600289015490941683529290522054909150610c27908363ffffffff610efe16565b8354600160a060020a03908116600090815260036020818152604080842060028a0180548716865290835281852096909655338516845291815281832094549093168252929091522054610c81908363ffffffff610eec16565b600160a060020a03338116600090815260036020908152604080832060028901548516845290915280822093909355600186015490911681522054610cd8908290610ccc9087610efe565b9063ffffffff610eec16565b600160a060020a03338116600090815260036020818152604080842060018a018054871686529083528185209690965588548516845291815281832094549093168252929091522054610d31908263ffffffff610efe16565b8354600160a060020a0390811660009081526003602090815260408083206001890180548616855290835281842095909555875484168352600482528083209454909316825292909252902054610d8e908563ffffffff610eec16565b8354600160a060020a0390811660009081526004602090815260408083206001890154909416835292905220556003830154610dd0908563ffffffff610eec16565b600384018190551515610e1257847feb3b41be12efa25c0c65725eeff2e2ad0250e2e5de2aa939e18f36aa656e98804260405190815260200160405180910390a25b82548590600160a060020a039081169033167fe6f814da7244d1ae6c61b54b5684858ba39cad7b9a91884be10060664987d75487864260405180848152602001838152602001828152602001935050505060405180910390a460018301548354600160a060020a0391821691167fac476a0560d6039841f7de6a83293b862be45e3182ab744e10e048952fe7e4e3834260405191825260208201526040908101905180910390a35050505050565b600160a060020a0380821660009081526003602090815260408083209386168352929052205492915050565b600082821115610ef857fe5b50900390565b600082820183811015610f0d57fe5b9392505050565b6000828202831580610f305750828482811515610f2d57fe5b04145b1515610f0d57fe5b6000808284811515610f4657fe5b049493505050505600a165627a7a7230582054531e0381624fe5c2e2b8099e6012aedf0424d23b52bfb4184336b667d9cf0e0029`

// DeployRadex deploys a new Ethereum contract, binding an instance of Radex to it.
func DeployRadex(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Radex, error) {
	parsed, err := abi.JSON(strings.NewReader(RadexABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RadexBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Radex{RadexCaller: RadexCaller{contract: contract}, RadexTransactor: RadexTransactor{contract: contract}, RadexFilterer: RadexFilterer{contract: contract}}, nil
}

// Radex is an auto generated Go binding around an Ethereum contract.
type Radex struct {
	RadexCaller     // Read-only binding to the contract
	RadexTransactor // Write-only binding to the contract
	RadexFilterer   // Log filterer for contract events
}

// RadexCaller is an auto generated read-only Go binding around an Ethereum contract.
type RadexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RadexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RadexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RadexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RadexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RadexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RadexSession struct {
	Contract     *Radex            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RadexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RadexCallerSession struct {
	Contract *RadexCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RadexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RadexTransactorSession struct {
	Contract     *RadexTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RadexRaw is an auto generated low-level Go binding around an Ethereum contract.
type RadexRaw struct {
	Contract *Radex // Generic contract binding to access the raw methods on
}

// RadexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RadexCallerRaw struct {
	Contract *RadexCaller // Generic read-only contract binding to access the raw methods on
}

// RadexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RadexTransactorRaw struct {
	Contract *RadexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRadex creates a new instance of Radex, bound to a specific deployed contract.
func NewRadex(address common.Address, backend bind.ContractBackend) (*Radex, error) {
	contract, err := bindRadex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Radex{RadexCaller: RadexCaller{contract: contract}, RadexTransactor: RadexTransactor{contract: contract}, RadexFilterer: RadexFilterer{contract: contract}}, nil
}

// NewRadexCaller creates a new read-only instance of Radex, bound to a specific deployed contract.
func NewRadexCaller(address common.Address, caller bind.ContractCaller) (*RadexCaller, error) {
	contract, err := bindRadex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RadexCaller{contract: contract}, nil
}

// NewRadexTransactor creates a new write-only instance of Radex, bound to a specific deployed contract.
func NewRadexTransactor(address common.Address, transactor bind.ContractTransactor) (*RadexTransactor, error) {
	contract, err := bindRadex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RadexTransactor{contract: contract}, nil
}

// NewRadexFilterer creates a new log filterer instance of Radex, bound to a specific deployed contract.
func NewRadexFilterer(address common.Address, filterer bind.ContractFilterer) (*RadexFilterer, error) {
	contract, err := bindRadex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RadexFilterer{contract: contract}, nil
}

// bindRadex binds a generic wrapper to an already deployed contract.
func bindRadex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RadexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Radex *RadexRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Radex.Contract.RadexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Radex *RadexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Radex.Contract.RadexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Radex *RadexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Radex.Contract.RadexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Radex *RadexCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Radex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Radex *RadexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Radex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Radex *RadexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Radex.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_Radex *RadexCaller) BalanceOf(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Radex.contract.Call(opts, out, "balanceOf", token, user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_Radex *RadexSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _Radex.Contract.BalanceOf(&_Radex.CallOpts, token, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_Radex *RadexCallerSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _Radex.Contract.BalanceOf(&_Radex.CallOpts, token, user)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances( address,  address) constant returns(uint256)
func (_Radex *RadexCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Radex.contract.Call(opts, out, "balances", arg0, arg1)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances( address,  address) constant returns(uint256)
func (_Radex *RadexSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Radex.Contract.Balances(&_Radex.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances( address,  address) constant returns(uint256)
func (_Radex *RadexCallerSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Radex.Contract.Balances(&_Radex.CallOpts, arg0, arg1)
}

// Commitments is a free data retrieval call binding the contract method 0xadb6f63f.
//
// Solidity: function commitments( address,  address) constant returns(uint256)
func (_Radex *RadexCaller) Commitments(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Radex.contract.Call(opts, out, "commitments", arg0, arg1)
	return *ret0, err
}

// Commitments is a free data retrieval call binding the contract method 0xadb6f63f.
//
// Solidity: function commitments( address,  address) constant returns(uint256)
func (_Radex *RadexSession) Commitments(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Radex.Contract.Commitments(&_Radex.CallOpts, arg0, arg1)
}

// Commitments is a free data retrieval call binding the contract method 0xadb6f63f.
//
// Solidity: function commitments( address,  address) constant returns(uint256)
func (_Radex *RadexCallerSession) Commitments(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Radex.Contract.Commitments(&_Radex.CallOpts, arg0, arg1)
}

// CommitmentsOf is a free data retrieval call binding the contract method 0xcc7e930c.
//
// Solidity: function commitmentsOf(token address, user address) constant returns(uint256)
func (_Radex *RadexCaller) CommitmentsOf(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Radex.contract.Call(opts, out, "commitmentsOf", token, user)
	return *ret0, err
}

// CommitmentsOf is a free data retrieval call binding the contract method 0xcc7e930c.
//
// Solidity: function commitmentsOf(token address, user address) constant returns(uint256)
func (_Radex *RadexSession) CommitmentsOf(token common.Address, user common.Address) (*big.Int, error) {
	return _Radex.Contract.CommitmentsOf(&_Radex.CallOpts, token, user)
}

// CommitmentsOf is a free data retrieval call binding the contract method 0xcc7e930c.
//
// Solidity: function commitmentsOf(token address, user address) constant returns(uint256)
func (_Radex *RadexCallerSession) CommitmentsOf(token common.Address, user common.Address) (*big.Int, error) {
	return _Radex.Contract.CommitmentsOf(&_Radex.CallOpts, token, user)
}

// FeeMultiplier is a free data retrieval call binding the contract method 0xe5a70ef7.
//
// Solidity: function feeMultiplier() constant returns(uint256)
func (_Radex *RadexCaller) FeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Radex.contract.Call(opts, out, "feeMultiplier")
	return *ret0, err
}

// FeeMultiplier is a free data retrieval call binding the contract method 0xe5a70ef7.
//
// Solidity: function feeMultiplier() constant returns(uint256)
func (_Radex *RadexSession) FeeMultiplier() (*big.Int, error) {
	return _Radex.Contract.FeeMultiplier(&_Radex.CallOpts)
}

// FeeMultiplier is a free data retrieval call binding the contract method 0xe5a70ef7.
//
// Solidity: function feeMultiplier() constant returns(uint256)
func (_Radex *RadexCallerSession) FeeMultiplier() (*big.Int, error) {
	return _Radex.Contract.FeeMultiplier(&_Radex.CallOpts)
}

// LatestOrderId is a free data retrieval call binding the contract method 0x6f3b60d6.
//
// Solidity: function latestOrderId() constant returns(uint256)
func (_Radex *RadexCaller) LatestOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Radex.contract.Call(opts, out, "latestOrderId")
	return *ret0, err
}

// LatestOrderId is a free data retrieval call binding the contract method 0x6f3b60d6.
//
// Solidity: function latestOrderId() constant returns(uint256)
func (_Radex *RadexSession) LatestOrderId() (*big.Int, error) {
	return _Radex.Contract.LatestOrderId(&_Radex.CallOpts)
}

// LatestOrderId is a free data retrieval call binding the contract method 0x6f3b60d6.
//
// Solidity: function latestOrderId() constant returns(uint256)
func (_Radex *RadexCallerSession) LatestOrderId() (*big.Int, error) {
	return _Radex.Contract.LatestOrderId(&_Radex.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(orderId uint256) returns()
func (_Radex *RadexTransactor) CancelOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _Radex.contract.Transact(opts, "cancelOrder", orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(orderId uint256) returns()
func (_Radex *RadexSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Radex.Contract.CancelOrder(&_Radex.TransactOpts, orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(orderId uint256) returns()
func (_Radex *RadexTransactorSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Radex.Contract.CancelOrder(&_Radex.TransactOpts, orderId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x518e62db.
//
// Solidity: function createOrder(sellToken address, buyToken address, amount uint256, priceMul uint256, priceDiv uint256) returns(orderId uint256)
func (_Radex *RadexTransactor) CreateOrder(opts *bind.TransactOpts, sellToken common.Address, buyToken common.Address, amount *big.Int, priceMul *big.Int, priceDiv *big.Int) (*types.Transaction, error) {
	return _Radex.contract.Transact(opts, "createOrder", sellToken, buyToken, amount, priceMul, priceDiv)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x518e62db.
//
// Solidity: function createOrder(sellToken address, buyToken address, amount uint256, priceMul uint256, priceDiv uint256) returns(orderId uint256)
func (_Radex *RadexSession) CreateOrder(sellToken common.Address, buyToken common.Address, amount *big.Int, priceMul *big.Int, priceDiv *big.Int) (*types.Transaction, error) {
	return _Radex.Contract.CreateOrder(&_Radex.TransactOpts, sellToken, buyToken, amount, priceMul, priceDiv)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x518e62db.
//
// Solidity: function createOrder(sellToken address, buyToken address, amount uint256, priceMul uint256, priceDiv uint256) returns(orderId uint256)
func (_Radex *RadexTransactorSession) CreateOrder(sellToken common.Address, buyToken common.Address, amount *big.Int, priceMul *big.Int, priceDiv *big.Int) (*types.Transaction, error) {
	return _Radex.Contract.CreateOrder(&_Radex.TransactOpts, sellToken, buyToken, amount, priceMul, priceDiv)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xef46e0ca.
//
// Solidity: function executeOrder(orderId uint256, amount uint256) returns()
func (_Radex *RadexTransactor) ExecuteOrder(opts *bind.TransactOpts, orderId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Radex.contract.Transact(opts, "executeOrder", orderId, amount)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xef46e0ca.
//
// Solidity: function executeOrder(orderId uint256, amount uint256) returns()
func (_Radex *RadexSession) ExecuteOrder(orderId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Radex.Contract.ExecuteOrder(&_Radex.TransactOpts, orderId, amount)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xef46e0ca.
//
// Solidity: function executeOrder(orderId uint256, amount uint256) returns()
func (_Radex *RadexTransactorSession) ExecuteOrder(orderId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Radex.Contract.ExecuteOrder(&_Radex.TransactOpts, orderId, amount)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Radex *RadexTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Radex.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Radex *RadexSession) Fund() (*types.Transaction, error) {
	return _Radex.Contract.Fund(&_Radex.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Radex *RadexTransactorSession) Fund() (*types.Transaction, error) {
	return _Radex.Contract.Fund(&_Radex.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(token address, value uint256) returns()
func (_Radex *RadexTransactor) Redeem(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Radex.contract.Transact(opts, "redeem", token, value)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(token address, value uint256) returns()
func (_Radex *RadexSession) Redeem(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Radex.Contract.Redeem(&_Radex.TransactOpts, token, value)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(token address, value uint256) returns()
func (_Radex *RadexTransactorSession) Redeem(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Radex.Contract.Redeem(&_Radex.TransactOpts, token, value)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(erc20token address, erc223token address) returns()
func (_Radex *RadexTransactor) Register(opts *bind.TransactOpts, erc20token common.Address, erc223token common.Address) (*types.Transaction, error) {
	return _Radex.contract.Transact(opts, "register", erc20token, erc223token)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(erc20token address, erc223token address) returns()
func (_Radex *RadexSession) Register(erc20token common.Address, erc223token common.Address) (*types.Transaction, error) {
	return _Radex.Contract.Register(&_Radex.TransactOpts, erc20token, erc223token)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(erc20token address, erc223token address) returns()
func (_Radex *RadexTransactorSession) Register(erc20token common.Address, erc223token common.Address) (*types.Transaction, error) {
	return _Radex.Contract.Register(&_Radex.TransactOpts, erc20token, erc223token)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256,  bytes) returns()
func (_Radex *RadexTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Radex.contract.Transact(opts, "tokenFallback", _from, _value, arg2)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256,  bytes) returns()
func (_Radex *RadexSession) TokenFallback(_from common.Address, _value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Radex.Contract.TokenFallback(&_Radex.TransactOpts, _from, _value, arg2)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256,  bytes) returns()
func (_Radex *RadexTransactorSession) TokenFallback(_from common.Address, _value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Radex.Contract.TokenFallback(&_Radex.TransactOpts, _from, _value, arg2)
}

// RadexDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Radex contract.
type RadexDepositIterator struct {
	Event *RadexDeposit // Event containing the contract specifics and raw log

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
func (it *RadexDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RadexDeposit)
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
		it.Event = new(RadexDeposit)
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
func (it *RadexDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RadexDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RadexDeposit represents a Deposit event raised by the Radex contract.
type RadexDeposit struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(_token indexed address, _owner indexed address, _amount uint256, _time uint256)
func (_Radex *RadexFilterer) FilterDeposit(opts *bind.FilterOpts, _token []common.Address, _owner []common.Address) (*RadexDepositIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Radex.contract.FilterLogs(opts, "Deposit", _tokenRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &RadexDepositIterator{contract: _Radex.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(_token indexed address, _owner indexed address, _amount uint256, _time uint256)
func (_Radex *RadexFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *RadexDeposit, _token []common.Address, _owner []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Radex.contract.WatchLogs(opts, "Deposit", _tokenRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RadexDeposit)
				if err := _Radex.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// RadexMarketMakerIterator is returned from FilterMarketMaker and is used to iterate over the raw logs and unpacked data for MarketMaker events raised by the Radex contract.
type RadexMarketMakerIterator struct {
	Event *RadexMarketMaker // Event containing the contract specifics and raw log

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
func (it *RadexMarketMakerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RadexMarketMaker)
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
		it.Event = new(RadexMarketMaker)
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
func (it *RadexMarketMakerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RadexMarketMakerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RadexMarketMaker represents a MarketMaker event raised by the Radex contract.
type RadexMarketMaker struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketMaker is a free log retrieval operation binding the contract event 0xac476a0560d6039841f7de6a83293b862be45e3182ab744e10e048952fe7e4e3.
//
// Solidity: event MarketMaker(_owner indexed address, _token indexed address, _amount uint256, _time uint256)
func (_Radex *RadexFilterer) FilterMarketMaker(opts *bind.FilterOpts, _owner []common.Address, _token []common.Address) (*RadexMarketMakerIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Radex.contract.FilterLogs(opts, "MarketMaker", _ownerRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &RadexMarketMakerIterator{contract: _Radex.contract, event: "MarketMaker", logs: logs, sub: sub}, nil
}

// WatchMarketMaker is a free log subscription operation binding the contract event 0xac476a0560d6039841f7de6a83293b862be45e3182ab744e10e048952fe7e4e3.
//
// Solidity: event MarketMaker(_owner indexed address, _token indexed address, _amount uint256, _time uint256)
func (_Radex *RadexFilterer) WatchMarketMaker(opts *bind.WatchOpts, sink chan<- *RadexMarketMaker, _owner []common.Address, _token []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Radex.contract.WatchLogs(opts, "MarketMaker", _ownerRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RadexMarketMaker)
				if err := _Radex.contract.UnpackLog(event, "MarketMaker", log); err != nil {
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

// RadexNewOrderIterator is returned from FilterNewOrder and is used to iterate over the raw logs and unpacked data for NewOrder events raised by the Radex contract.
type RadexNewOrderIterator struct {
	Event *RadexNewOrder // Event containing the contract specifics and raw log

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
func (it *RadexNewOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RadexNewOrder)
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
		it.Event = new(RadexNewOrder)
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
func (it *RadexNewOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RadexNewOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RadexNewOrder represents a NewOrder event raised by the Radex contract.
type RadexNewOrder struct {
	Id        *big.Int
	Owner     common.Address
	SellToken common.Address
	BuyToken  common.Address
	Amount    *big.Int
	PriceMul  *big.Int
	PriceDiv  *big.Int
	Time      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewOrder is a free log retrieval operation binding the contract event 0x506cb18801692179462f2b7b4b6f10bea763984f94a1f1f3cd858e3a645b2329.
//
// Solidity: event NewOrder(_id uint256, _owner indexed address, _sellToken indexed address, _buyToken indexed address, _amount uint256, _priceMul uint256, _priceDiv uint256, _time uint256)
func (_Radex *RadexFilterer) FilterNewOrder(opts *bind.FilterOpts, _owner []common.Address, _sellToken []common.Address, _buyToken []common.Address) (*RadexNewOrderIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _sellTokenRule []interface{}
	for _, _sellTokenItem := range _sellToken {
		_sellTokenRule = append(_sellTokenRule, _sellTokenItem)
	}
	var _buyTokenRule []interface{}
	for _, _buyTokenItem := range _buyToken {
		_buyTokenRule = append(_buyTokenRule, _buyTokenItem)
	}

	logs, sub, err := _Radex.contract.FilterLogs(opts, "NewOrder", _ownerRule, _sellTokenRule, _buyTokenRule)
	if err != nil {
		return nil, err
	}
	return &RadexNewOrderIterator{contract: _Radex.contract, event: "NewOrder", logs: logs, sub: sub}, nil
}

// WatchNewOrder is a free log subscription operation binding the contract event 0x506cb18801692179462f2b7b4b6f10bea763984f94a1f1f3cd858e3a645b2329.
//
// Solidity: event NewOrder(_id uint256, _owner indexed address, _sellToken indexed address, _buyToken indexed address, _amount uint256, _priceMul uint256, _priceDiv uint256, _time uint256)
func (_Radex *RadexFilterer) WatchNewOrder(opts *bind.WatchOpts, sink chan<- *RadexNewOrder, _owner []common.Address, _sellToken []common.Address, _buyToken []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _sellTokenRule []interface{}
	for _, _sellTokenItem := range _sellToken {
		_sellTokenRule = append(_sellTokenRule, _sellTokenItem)
	}
	var _buyTokenRule []interface{}
	for _, _buyTokenItem := range _buyToken {
		_buyTokenRule = append(_buyTokenRule, _buyTokenItem)
	}

	logs, sub, err := _Radex.contract.WatchLogs(opts, "NewOrder", _ownerRule, _sellTokenRule, _buyTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RadexNewOrder)
				if err := _Radex.contract.UnpackLog(event, "NewOrder", log); err != nil {
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

// RadexOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Radex contract.
type RadexOrderCancelledIterator struct {
	Event *RadexOrderCancelled // Event containing the contract specifics and raw log

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
func (it *RadexOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RadexOrderCancelled)
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
		it.Event = new(RadexOrderCancelled)
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
func (it *RadexOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RadexOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RadexOrderCancelled represents a OrderCancelled event raised by the Radex contract.
type RadexOrderCancelled struct {
	Id   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xaa2a006d186dc89a3d6851834c85d16226329ac40bcad6531a5756a38ed1c90f.
//
// Solidity: event OrderCancelled(_id indexed uint256, _time uint256)
func (_Radex *RadexFilterer) FilterOrderCancelled(opts *bind.FilterOpts, _id []*big.Int) (*RadexOrderCancelledIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Radex.contract.FilterLogs(opts, "OrderCancelled", _idRule)
	if err != nil {
		return nil, err
	}
	return &RadexOrderCancelledIterator{contract: _Radex.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xaa2a006d186dc89a3d6851834c85d16226329ac40bcad6531a5756a38ed1c90f.
//
// Solidity: event OrderCancelled(_id indexed uint256, _time uint256)
func (_Radex *RadexFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *RadexOrderCancelled, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Radex.contract.WatchLogs(opts, "OrderCancelled", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RadexOrderCancelled)
				if err := _Radex.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// RadexOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Radex contract.
type RadexOrderFulfilledIterator struct {
	Event *RadexOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *RadexOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RadexOrderFulfilled)
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
		it.Event = new(RadexOrderFulfilled)
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
func (it *RadexOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RadexOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RadexOrderFulfilled represents a OrderFulfilled event raised by the Radex contract.
type RadexOrderFulfilled struct {
	Id   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0xeb3b41be12efa25c0c65725eeff2e2ad0250e2e5de2aa939e18f36aa656e9880.
//
// Solidity: event OrderFulfilled(_id indexed uint256, _time uint256)
func (_Radex *RadexFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, _id []*big.Int) (*RadexOrderFulfilledIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Radex.contract.FilterLogs(opts, "OrderFulfilled", _idRule)
	if err != nil {
		return nil, err
	}
	return &RadexOrderFulfilledIterator{contract: _Radex.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0xeb3b41be12efa25c0c65725eeff2e2ad0250e2e5de2aa939e18f36aa656e9880.
//
// Solidity: event OrderFulfilled(_id indexed uint256, _time uint256)
func (_Radex *RadexFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *RadexOrderFulfilled, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Radex.contract.WatchLogs(opts, "OrderFulfilled", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RadexOrderFulfilled)
				if err := _Radex.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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

// RadexTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Radex contract.
type RadexTradeIterator struct {
	Event *RadexTrade // Event containing the contract specifics and raw log

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
func (it *RadexTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RadexTrade)
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
		it.Event = new(RadexTrade)
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
func (it *RadexTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RadexTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RadexTrade represents a Trade event raised by the Radex contract.
type RadexTrade struct {
	From         common.Address
	To           common.Address
	OrderId      *big.Int
	SoldTokens   *big.Int
	BoughtTokens *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xe6f814da7244d1ae6c61b54b5684858ba39cad7b9a91884be10060664987d754.
//
// Solidity: event Trade(_from indexed address, _to indexed address, _orderId indexed uint256, _soldTokens uint256, _boughtTokens uint256, _time uint256)
func (_Radex *RadexFilterer) FilterTrade(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _orderId []*big.Int) (*RadexTradeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Radex.contract.FilterLogs(opts, "Trade", _fromRule, _toRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return &RadexTradeIterator{contract: _Radex.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xe6f814da7244d1ae6c61b54b5684858ba39cad7b9a91884be10060664987d754.
//
// Solidity: event Trade(_from indexed address, _to indexed address, _orderId indexed uint256, _soldTokens uint256, _boughtTokens uint256, _time uint256)
func (_Radex *RadexFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *RadexTrade, _from []common.Address, _to []common.Address, _orderId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Radex.contract.WatchLogs(opts, "Trade", _fromRule, _toRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RadexTrade)
				if err := _Radex.contract.UnpackLog(event, "Trade", log); err != nil {
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

// RadexWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Radex contract.
type RadexWithdrawalIterator struct {
	Event *RadexWithdrawal // Event containing the contract specifics and raw log

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
func (it *RadexWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RadexWithdrawal)
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
		it.Event = new(RadexWithdrawal)
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
func (it *RadexWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RadexWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RadexWithdrawal represents a Withdrawal event raised by the Radex contract.
type RadexWithdrawal struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc2b4a290c20fb28939d29f102514fbffd2b73c059ffba8b78250c94161d5fcc6.
//
// Solidity: event Withdrawal(_token indexed address, _owner indexed address, _amount uint256, _time uint256)
func (_Radex *RadexFilterer) FilterWithdrawal(opts *bind.FilterOpts, _token []common.Address, _owner []common.Address) (*RadexWithdrawalIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Radex.contract.FilterLogs(opts, "Withdrawal", _tokenRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &RadexWithdrawalIterator{contract: _Radex.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc2b4a290c20fb28939d29f102514fbffd2b73c059ffba8b78250c94161d5fcc6.
//
// Solidity: event Withdrawal(_token indexed address, _owner indexed address, _amount uint256, _time uint256)
func (_Radex *RadexFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *RadexWithdrawal, _token []common.Address, _owner []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Radex.contract.WatchLogs(opts, "Withdrawal", _tokenRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RadexWithdrawal)
				if err := _Radex.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820378fd4ad51165f8c30ba521a1b99902c795e5b961c01fe7e4240ad56d51d1dbd0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b6107238061001e6000396000f3006060604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009d57806318160ddd146100d357806323b872dd146100f8578063313ce56714610120578063661884631461014957806370a082311461016b578063a9059cbb1461018a578063d73dd623146101ac578063dd62ed3e146101ce575b600080fd5b34156100a857600080fd5b6100bf600160a060020a03600435166024356101f3565b604051901515815260200160405180910390f35b34156100de57600080fd5b6100e661025f565b60405190815260200160405180910390f35b341561010357600080fd5b6100bf600160a060020a0360043581169060243516604435610265565b341561012b57600080fd5b6101336103e7565b60405160ff909116815260200160405180910390f35b341561015457600080fd5b6100bf600160a060020a03600435166024356103f0565b341561017657600080fd5b6100e6600160a060020a03600435166104ea565b341561019557600080fd5b6100bf600160a060020a0360043516602435610505565b34156101b757600080fd5b6100bf600160a060020a0360043516602435610600565b34156101d957600080fd5b6100e6600160a060020a03600435811690602435166106a4565b600160a060020a03338116600081815260036020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561027c57600080fd5b600160a060020a0384166000908152600260205260409020548211156102a157600080fd5b600160a060020a03808516600090815260036020908152604080832033909416835292905220548211156102d457600080fd5b600160a060020a0384166000908152600260205260409020546102fd908363ffffffff6106cf16565b600160a060020a038086166000908152600260205260408082209390935590851681522054610332908363ffffffff6106e116565b600160a060020a0380851660009081526002602090815260408083209490945587831682526003815283822033909316825291909152205461037a908363ffffffff6106cf16565b600160a060020a03808616600081815260036020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60015460ff1681565b600160a060020a0333811660009081526003602090815260408083209386168352929052908120548083111561044d57600160a060020a033381166000908152600360209081526040808320938816835292905290812055610484565b61045d818463ffffffff6106cf16565b600160a060020a033381166000908152600360209081526040808320938916835292905220555b600160a060020a0333811660008181526003602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526002602052604090205490565b6000600160a060020a038316151561051c57600080fd5b600160a060020a03331660009081526002602052604090205482111561054157600080fd5b600160a060020a03331660009081526002602052604090205461056a908363ffffffff6106cf16565b600160a060020a03338116600090815260026020526040808220939093559085168152205461059f908363ffffffff6106e116565b600160a060020a0380851660008181526002602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600360209081526040808320938616835292905290812054610638908363ffffffff6106e116565b600160a060020a0333811660008181526003602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b6000828211156106db57fe5b50900390565b6000828201838110156106f057fe5b93925050505600a165627a7a7230582046ccef263c3e5d4f77aea0e2ff6db7faef9e9c6aa1a85861afc826aa5ae6c2720029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StandardToken *StandardTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StandardToken *StandardTokenSession) Decimals() (uint8, error) {
	return _StandardToken.Contract.Decimals(&_StandardToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StandardToken *StandardTokenCallerSession) Decimals() (uint8, error) {
	return _StandardToken.Contract.Decimals(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
