// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// BeefyClientCommitment is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientCommitment struct {
	BlockNumber    uint32
	ValidatorSetID uint64
	Payload        []BeefyClientPayloadItem
}

// BeefyClientMMRLeaf is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientMMRLeaf struct {
	Version              uint8
	ParentNumber         uint32
	ParentHash           [32]byte
	NextAuthoritySetID   uint64
	NextAuthoritySetLen  uint32
	NextAuthoritySetRoot [32]byte
	ParachainHeadsRoot   [32]byte
	MessageCommitment    [32]byte
}

// BeefyClientPayloadItem is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientPayloadItem struct {
	PayloadID [2]byte
	Data      []byte
}

// BeefyClientValidatorProof is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientValidatorProof struct {
	V       uint8
	R       [32]byte
	S       [32]byte
	Index   *big.Int
	Account common.Address
	Proof   [][32]byte
}

// BeefyClientValidatorSet is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientValidatorSet struct {
	Id     *big.Int
	Length *big.Int
	Root   [32]byte
}

// Uint16Array is an auto generated low-level Go binding around an user-defined struct.
type Uint16Array struct {
	Data   []*big.Int
	Length *big.Int
}

// BeefyClientMetaData contains all meta data concerning the BeefyClient contract.
var BeefyClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_randaoCommitDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_randaoCommitExpiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minNumRequiredSignatures\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialBeefyBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_initialValidatorSet\",\"type\":\"tuple\",\"internalType\":\"structBeefyClient.ValidatorSet\",\"components\":[{\"name\":\"id\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"length\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_nextValidatorSet\",\"type\":\"tuple\",\"internalType\":\"structBeefyClient.ValidatorSet\",\"components\":[{\"name\":\"id\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"length\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MMR_ROOT_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"commitPrevRandao\",\"inputs\":[{\"name\":\"commitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createFinalBitfield\",\"inputs\":[{\"name\":\"commitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bitfield\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createInitialBitfield\",\"inputs\":[{\"name\":\"bitsToSet\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"currentValidatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"id\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"length\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"usageCounters\",\"type\":\"tuple\",\"internalType\":\"structUint16Array\",\"components\":[{\"name\":\"data\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestBeefyBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestMMRRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minNumRequiredSignatures\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextValidatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"id\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"length\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"usageCounters\",\"type\":\"tuple\",\"internalType\":\"structUint16Array\",\"components\":[{\"name\":\"data\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randaoCommitDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randaoCommitExpiration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitFinal\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"tuple\",\"internalType\":\"structBeefyClient.Commitment\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"validatorSetID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payload\",\"type\":\"tuple[]\",\"internalType\":\"structBeefyClient.PayloadItem[]\",\"components\":[{\"name\":\"payloadID\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"bitfield\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeefyClient.ValidatorProof[]\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structBeefyClient.MMRLeaf\",\"components\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"parentNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextAuthoritySetID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nextAuthoritySetLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextAuthoritySetRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parachainHeadsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageCommitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"leafProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"leafProofOrder\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitInitial\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"tuple\",\"internalType\":\"structBeefyClient.Commitment\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"validatorSetID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payload\",\"type\":\"tuple[]\",\"internalType\":\"structBeefyClient.PayloadItem[]\",\"components\":[{\"name\":\"payloadID\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"bitfield\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBeefyClient.ValidatorProof\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tickets\",\"inputs\":[{\"name\":\"ticketID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validatorSetLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numRequiredSignatures\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"prevRandao\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bitfieldHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMMRLeafProof\",\"inputs\":[{\"name\":\"leafHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofOrder\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"NewMMRRoot\",\"inputs\":[{\"name\":\"mmrRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTicket\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CommitmentNotRelevant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IndexOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBitfield\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBitfieldLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMMRLeaf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMMRLeafProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMMRRootLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTicket\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughClaims\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PrevRandaoAlreadyCaptured\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PrevRandaoNotCaptured\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StaleCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TicketExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedCompactEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WaitPeriodNotOver\",\"inputs\":[]}]",
}

// BeefyClientABI is the input ABI used to generate the binding from.
// Deprecated: Use BeefyClientMetaData.ABI instead.
var BeefyClientABI = BeefyClientMetaData.ABI

// BeefyClient is an auto generated Go binding around an Ethereum contract.
type BeefyClient struct {
	BeefyClientCaller     // Read-only binding to the contract
	BeefyClientTransactor // Write-only binding to the contract
	BeefyClientFilterer   // Log filterer for contract events
}

// BeefyClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeefyClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeefyClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeefyClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeefyClientSession struct {
	Contract     *BeefyClient      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeefyClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeefyClientCallerSession struct {
	Contract *BeefyClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BeefyClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeefyClientTransactorSession struct {
	Contract     *BeefyClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BeefyClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeefyClientRaw struct {
	Contract *BeefyClient // Generic contract binding to access the raw methods on
}

// BeefyClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeefyClientCallerRaw struct {
	Contract *BeefyClientCaller // Generic read-only contract binding to access the raw methods on
}

// BeefyClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeefyClientTransactorRaw struct {
	Contract *BeefyClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeefyClient creates a new instance of BeefyClient, bound to a specific deployed contract.
func NewBeefyClient(address common.Address, backend bind.ContractBackend) (*BeefyClient, error) {
	contract, err := bindBeefyClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeefyClient{BeefyClientCaller: BeefyClientCaller{contract: contract}, BeefyClientTransactor: BeefyClientTransactor{contract: contract}, BeefyClientFilterer: BeefyClientFilterer{contract: contract}}, nil
}

// NewBeefyClientCaller creates a new read-only instance of BeefyClient, bound to a specific deployed contract.
func NewBeefyClientCaller(address common.Address, caller bind.ContractCaller) (*BeefyClientCaller, error) {
	contract, err := bindBeefyClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeefyClientCaller{contract: contract}, nil
}

// NewBeefyClientTransactor creates a new write-only instance of BeefyClient, bound to a specific deployed contract.
func NewBeefyClientTransactor(address common.Address, transactor bind.ContractTransactor) (*BeefyClientTransactor, error) {
	contract, err := bindBeefyClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeefyClientTransactor{contract: contract}, nil
}

// NewBeefyClientFilterer creates a new log filterer instance of BeefyClient, bound to a specific deployed contract.
func NewBeefyClientFilterer(address common.Address, filterer bind.ContractFilterer) (*BeefyClientFilterer, error) {
	contract, err := bindBeefyClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeefyClientFilterer{contract: contract}, nil
}

// bindBeefyClient binds a generic wrapper to an already deployed contract.
func bindBeefyClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeefyClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeefyClient *BeefyClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeefyClient.Contract.BeefyClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeefyClient *BeefyClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyClient.Contract.BeefyClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeefyClient *BeefyClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeefyClient.Contract.BeefyClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeefyClient *BeefyClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeefyClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeefyClient *BeefyClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeefyClient *BeefyClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeefyClient.Contract.contract.Transact(opts, method, params...)
}

// MMRROOTID is a free data retrieval call binding the contract method 0x0a7c8faa.
//
// Solidity: function MMR_ROOT_ID() view returns(bytes2)
func (_BeefyClient *BeefyClientCaller) MMRROOTID(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "MMR_ROOT_ID")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// MMRROOTID is a free data retrieval call binding the contract method 0x0a7c8faa.
//
// Solidity: function MMR_ROOT_ID() view returns(bytes2)
func (_BeefyClient *BeefyClientSession) MMRROOTID() ([2]byte, error) {
	return _BeefyClient.Contract.MMRROOTID(&_BeefyClient.CallOpts)
}

// MMRROOTID is a free data retrieval call binding the contract method 0x0a7c8faa.
//
// Solidity: function MMR_ROOT_ID() view returns(bytes2)
func (_BeefyClient *BeefyClientCallerSession) MMRROOTID() ([2]byte, error) {
	return _BeefyClient.Contract.MMRROOTID(&_BeefyClient.CallOpts)
}

// CreateFinalBitfield is a free data retrieval call binding the contract method 0x8ab81d13.
//
// Solidity: function createFinalBitfield(bytes32 commitmentHash, uint256[] bitfield) view returns(uint256[])
func (_BeefyClient *BeefyClientCaller) CreateFinalBitfield(opts *bind.CallOpts, commitmentHash [32]byte, bitfield []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "createFinalBitfield", commitmentHash, bitfield)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CreateFinalBitfield is a free data retrieval call binding the contract method 0x8ab81d13.
//
// Solidity: function createFinalBitfield(bytes32 commitmentHash, uint256[] bitfield) view returns(uint256[])
func (_BeefyClient *BeefyClientSession) CreateFinalBitfield(commitmentHash [32]byte, bitfield []*big.Int) ([]*big.Int, error) {
	return _BeefyClient.Contract.CreateFinalBitfield(&_BeefyClient.CallOpts, commitmentHash, bitfield)
}

// CreateFinalBitfield is a free data retrieval call binding the contract method 0x8ab81d13.
//
// Solidity: function createFinalBitfield(bytes32 commitmentHash, uint256[] bitfield) view returns(uint256[])
func (_BeefyClient *BeefyClientCallerSession) CreateFinalBitfield(commitmentHash [32]byte, bitfield []*big.Int) ([]*big.Int, error) {
	return _BeefyClient.Contract.CreateFinalBitfield(&_BeefyClient.CallOpts, commitmentHash, bitfield)
}

// CreateInitialBitfield is a free data retrieval call binding the contract method 0x5da57fe9.
//
// Solidity: function createInitialBitfield(uint256[] bitsToSet, uint256 length) pure returns(uint256[])
func (_BeefyClient *BeefyClientCaller) CreateInitialBitfield(opts *bind.CallOpts, bitsToSet []*big.Int, length *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "createInitialBitfield", bitsToSet, length)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CreateInitialBitfield is a free data retrieval call binding the contract method 0x5da57fe9.
//
// Solidity: function createInitialBitfield(uint256[] bitsToSet, uint256 length) pure returns(uint256[])
func (_BeefyClient *BeefyClientSession) CreateInitialBitfield(bitsToSet []*big.Int, length *big.Int) ([]*big.Int, error) {
	return _BeefyClient.Contract.CreateInitialBitfield(&_BeefyClient.CallOpts, bitsToSet, length)
}

// CreateInitialBitfield is a free data retrieval call binding the contract method 0x5da57fe9.
//
// Solidity: function createInitialBitfield(uint256[] bitsToSet, uint256 length) pure returns(uint256[])
func (_BeefyClient *BeefyClientCallerSession) CreateInitialBitfield(bitsToSet []*big.Int, length *big.Int) ([]*big.Int, error) {
	return _BeefyClient.Contract.CreateInitialBitfield(&_BeefyClient.CallOpts, bitsToSet, length)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x2cdea717.
//
// Solidity: function currentValidatorSet() view returns(uint128 id, uint128 length, bytes32 root, (uint256[],uint256) usageCounters)
func (_BeefyClient *BeefyClientCaller) CurrentValidatorSet(opts *bind.CallOpts) (struct {
	Id            *big.Int
	Length        *big.Int
	Root          [32]byte
	UsageCounters Uint16Array
}, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "currentValidatorSet")

	outstruct := new(struct {
		Id            *big.Int
		Length        *big.Int
		Root          [32]byte
		UsageCounters Uint16Array
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Root = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.UsageCounters = *abi.ConvertType(out[3], new(Uint16Array)).(*Uint16Array)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x2cdea717.
//
// Solidity: function currentValidatorSet() view returns(uint128 id, uint128 length, bytes32 root, (uint256[],uint256) usageCounters)
func (_BeefyClient *BeefyClientSession) CurrentValidatorSet() (struct {
	Id            *big.Int
	Length        *big.Int
	Root          [32]byte
	UsageCounters Uint16Array
}, error) {
	return _BeefyClient.Contract.CurrentValidatorSet(&_BeefyClient.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x2cdea717.
//
// Solidity: function currentValidatorSet() view returns(uint128 id, uint128 length, bytes32 root, (uint256[],uint256) usageCounters)
func (_BeefyClient *BeefyClientCallerSession) CurrentValidatorSet() (struct {
	Id            *big.Int
	Length        *big.Int
	Root          [32]byte
	UsageCounters Uint16Array
}, error) {
	return _BeefyClient.Contract.CurrentValidatorSet(&_BeefyClient.CallOpts)
}

// LatestBeefyBlock is a free data retrieval call binding the contract method 0x66ae69a0.
//
// Solidity: function latestBeefyBlock() view returns(uint64)
func (_BeefyClient *BeefyClientCaller) LatestBeefyBlock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "latestBeefyBlock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestBeefyBlock is a free data retrieval call binding the contract method 0x66ae69a0.
//
// Solidity: function latestBeefyBlock() view returns(uint64)
func (_BeefyClient *BeefyClientSession) LatestBeefyBlock() (uint64, error) {
	return _BeefyClient.Contract.LatestBeefyBlock(&_BeefyClient.CallOpts)
}

// LatestBeefyBlock is a free data retrieval call binding the contract method 0x66ae69a0.
//
// Solidity: function latestBeefyBlock() view returns(uint64)
func (_BeefyClient *BeefyClientCallerSession) LatestBeefyBlock() (uint64, error) {
	return _BeefyClient.Contract.LatestBeefyBlock(&_BeefyClient.CallOpts)
}

// LatestMMRRoot is a free data retrieval call binding the contract method 0x41c9634e.
//
// Solidity: function latestMMRRoot() view returns(bytes32)
func (_BeefyClient *BeefyClientCaller) LatestMMRRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "latestMMRRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestMMRRoot is a free data retrieval call binding the contract method 0x41c9634e.
//
// Solidity: function latestMMRRoot() view returns(bytes32)
func (_BeefyClient *BeefyClientSession) LatestMMRRoot() ([32]byte, error) {
	return _BeefyClient.Contract.LatestMMRRoot(&_BeefyClient.CallOpts)
}

// LatestMMRRoot is a free data retrieval call binding the contract method 0x41c9634e.
//
// Solidity: function latestMMRRoot() view returns(bytes32)
func (_BeefyClient *BeefyClientCallerSession) LatestMMRRoot() ([32]byte, error) {
	return _BeefyClient.Contract.LatestMMRRoot(&_BeefyClient.CallOpts)
}

// MinNumRequiredSignatures is a free data retrieval call binding the contract method 0x6f55bd32.
//
// Solidity: function minNumRequiredSignatures() view returns(uint256)
func (_BeefyClient *BeefyClientCaller) MinNumRequiredSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "minNumRequiredSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinNumRequiredSignatures is a free data retrieval call binding the contract method 0x6f55bd32.
//
// Solidity: function minNumRequiredSignatures() view returns(uint256)
func (_BeefyClient *BeefyClientSession) MinNumRequiredSignatures() (*big.Int, error) {
	return _BeefyClient.Contract.MinNumRequiredSignatures(&_BeefyClient.CallOpts)
}

// MinNumRequiredSignatures is a free data retrieval call binding the contract method 0x6f55bd32.
//
// Solidity: function minNumRequiredSignatures() view returns(uint256)
func (_BeefyClient *BeefyClientCallerSession) MinNumRequiredSignatures() (*big.Int, error) {
	return _BeefyClient.Contract.MinNumRequiredSignatures(&_BeefyClient.CallOpts)
}

// NextValidatorSet is a free data retrieval call binding the contract method 0x36667513.
//
// Solidity: function nextValidatorSet() view returns(uint128 id, uint128 length, bytes32 root, (uint256[],uint256) usageCounters)
func (_BeefyClient *BeefyClientCaller) NextValidatorSet(opts *bind.CallOpts) (struct {
	Id            *big.Int
	Length        *big.Int
	Root          [32]byte
	UsageCounters Uint16Array
}, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "nextValidatorSet")

	outstruct := new(struct {
		Id            *big.Int
		Length        *big.Int
		Root          [32]byte
		UsageCounters Uint16Array
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Root = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.UsageCounters = *abi.ConvertType(out[3], new(Uint16Array)).(*Uint16Array)

	return *outstruct, err

}

// NextValidatorSet is a free data retrieval call binding the contract method 0x36667513.
//
// Solidity: function nextValidatorSet() view returns(uint128 id, uint128 length, bytes32 root, (uint256[],uint256) usageCounters)
func (_BeefyClient *BeefyClientSession) NextValidatorSet() (struct {
	Id            *big.Int
	Length        *big.Int
	Root          [32]byte
	UsageCounters Uint16Array
}, error) {
	return _BeefyClient.Contract.NextValidatorSet(&_BeefyClient.CallOpts)
}

// NextValidatorSet is a free data retrieval call binding the contract method 0x36667513.
//
// Solidity: function nextValidatorSet() view returns(uint128 id, uint128 length, bytes32 root, (uint256[],uint256) usageCounters)
func (_BeefyClient *BeefyClientCallerSession) NextValidatorSet() (struct {
	Id            *big.Int
	Length        *big.Int
	Root          [32]byte
	UsageCounters Uint16Array
}, error) {
	return _BeefyClient.Contract.NextValidatorSet(&_BeefyClient.CallOpts)
}

// RandaoCommitDelay is a free data retrieval call binding the contract method 0x591d99ee.
//
// Solidity: function randaoCommitDelay() view returns(uint256)
func (_BeefyClient *BeefyClientCaller) RandaoCommitDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "randaoCommitDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandaoCommitDelay is a free data retrieval call binding the contract method 0x591d99ee.
//
// Solidity: function randaoCommitDelay() view returns(uint256)
func (_BeefyClient *BeefyClientSession) RandaoCommitDelay() (*big.Int, error) {
	return _BeefyClient.Contract.RandaoCommitDelay(&_BeefyClient.CallOpts)
}

// RandaoCommitDelay is a free data retrieval call binding the contract method 0x591d99ee.
//
// Solidity: function randaoCommitDelay() view returns(uint256)
func (_BeefyClient *BeefyClientCallerSession) RandaoCommitDelay() (*big.Int, error) {
	return _BeefyClient.Contract.RandaoCommitDelay(&_BeefyClient.CallOpts)
}

// RandaoCommitExpiration is a free data retrieval call binding the contract method 0xad209a9b.
//
// Solidity: function randaoCommitExpiration() view returns(uint256)
func (_BeefyClient *BeefyClientCaller) RandaoCommitExpiration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "randaoCommitExpiration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandaoCommitExpiration is a free data retrieval call binding the contract method 0xad209a9b.
//
// Solidity: function randaoCommitExpiration() view returns(uint256)
func (_BeefyClient *BeefyClientSession) RandaoCommitExpiration() (*big.Int, error) {
	return _BeefyClient.Contract.RandaoCommitExpiration(&_BeefyClient.CallOpts)
}

// RandaoCommitExpiration is a free data retrieval call binding the contract method 0xad209a9b.
//
// Solidity: function randaoCommitExpiration() view returns(uint256)
func (_BeefyClient *BeefyClientCallerSession) RandaoCommitExpiration() (*big.Int, error) {
	return _BeefyClient.Contract.RandaoCommitExpiration(&_BeefyClient.CallOpts)
}

// Tickets is a free data retrieval call binding the contract method 0xdf0dd0d5.
//
// Solidity: function tickets(bytes32 ticketID) view returns(uint64 blockNumber, uint32 validatorSetLen, uint32 numRequiredSignatures, uint256 prevRandao, bytes32 bitfieldHash)
func (_BeefyClient *BeefyClientCaller) Tickets(opts *bind.CallOpts, ticketID [32]byte) (struct {
	BlockNumber           uint64
	ValidatorSetLen       uint32
	NumRequiredSignatures uint32
	PrevRandao            *big.Int
	BitfieldHash          [32]byte
}, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "tickets", ticketID)

	outstruct := new(struct {
		BlockNumber           uint64
		ValidatorSetLen       uint32
		NumRequiredSignatures uint32
		PrevRandao            *big.Int
		BitfieldHash          [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ValidatorSetLen = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.NumRequiredSignatures = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.PrevRandao = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BitfieldHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Tickets is a free data retrieval call binding the contract method 0xdf0dd0d5.
//
// Solidity: function tickets(bytes32 ticketID) view returns(uint64 blockNumber, uint32 validatorSetLen, uint32 numRequiredSignatures, uint256 prevRandao, bytes32 bitfieldHash)
func (_BeefyClient *BeefyClientSession) Tickets(ticketID [32]byte) (struct {
	BlockNumber           uint64
	ValidatorSetLen       uint32
	NumRequiredSignatures uint32
	PrevRandao            *big.Int
	BitfieldHash          [32]byte
}, error) {
	return _BeefyClient.Contract.Tickets(&_BeefyClient.CallOpts, ticketID)
}

// Tickets is a free data retrieval call binding the contract method 0xdf0dd0d5.
//
// Solidity: function tickets(bytes32 ticketID) view returns(uint64 blockNumber, uint32 validatorSetLen, uint32 numRequiredSignatures, uint256 prevRandao, bytes32 bitfieldHash)
func (_BeefyClient *BeefyClientCallerSession) Tickets(ticketID [32]byte) (struct {
	BlockNumber           uint64
	ValidatorSetLen       uint32
	NumRequiredSignatures uint32
	PrevRandao            *big.Int
	BitfieldHash          [32]byte
}, error) {
	return _BeefyClient.Contract.Tickets(&_BeefyClient.CallOpts, ticketID)
}

// VerifyMMRLeafProof is a free data retrieval call binding the contract method 0xa401662b.
//
// Solidity: function verifyMMRLeafProof(bytes32 leafHash, bytes32[] proof, uint256 proofOrder) view returns(bool)
func (_BeefyClient *BeefyClientCaller) VerifyMMRLeafProof(opts *bind.CallOpts, leafHash [32]byte, proof [][32]byte, proofOrder *big.Int) (bool, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "verifyMMRLeafProof", leafHash, proof, proofOrder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMMRLeafProof is a free data retrieval call binding the contract method 0xa401662b.
//
// Solidity: function verifyMMRLeafProof(bytes32 leafHash, bytes32[] proof, uint256 proofOrder) view returns(bool)
func (_BeefyClient *BeefyClientSession) VerifyMMRLeafProof(leafHash [32]byte, proof [][32]byte, proofOrder *big.Int) (bool, error) {
	return _BeefyClient.Contract.VerifyMMRLeafProof(&_BeefyClient.CallOpts, leafHash, proof, proofOrder)
}

// VerifyMMRLeafProof is a free data retrieval call binding the contract method 0xa401662b.
//
// Solidity: function verifyMMRLeafProof(bytes32 leafHash, bytes32[] proof, uint256 proofOrder) view returns(bool)
func (_BeefyClient *BeefyClientCallerSession) VerifyMMRLeafProof(leafHash [32]byte, proof [][32]byte, proofOrder *big.Int) (bool, error) {
	return _BeefyClient.Contract.VerifyMMRLeafProof(&_BeefyClient.CallOpts, leafHash, proof, proofOrder)
}

// CommitPrevRandao is a paid mutator transaction binding the contract method 0xa77cf3d2.
//
// Solidity: function commitPrevRandao(bytes32 commitmentHash) returns()
func (_BeefyClient *BeefyClientTransactor) CommitPrevRandao(opts *bind.TransactOpts, commitmentHash [32]byte) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "commitPrevRandao", commitmentHash)
}

// CommitPrevRandao is a paid mutator transaction binding the contract method 0xa77cf3d2.
//
// Solidity: function commitPrevRandao(bytes32 commitmentHash) returns()
func (_BeefyClient *BeefyClientSession) CommitPrevRandao(commitmentHash [32]byte) (*types.Transaction, error) {
	return _BeefyClient.Contract.CommitPrevRandao(&_BeefyClient.TransactOpts, commitmentHash)
}

// CommitPrevRandao is a paid mutator transaction binding the contract method 0xa77cf3d2.
//
// Solidity: function commitPrevRandao(bytes32 commitmentHash) returns()
func (_BeefyClient *BeefyClientTransactorSession) CommitPrevRandao(commitmentHash [32]byte) (*types.Transaction, error) {
	return _BeefyClient.Contract.CommitPrevRandao(&_BeefyClient.TransactOpts, commitmentHash)
}

// SubmitFinal is a paid mutator transaction binding the contract method 0x40f29e47.
//
// Solidity: function submitFinal((uint32,uint64,(bytes2,bytes)[]) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs, (uint8,uint32,bytes32,uint64,uint32,bytes32,bytes32,bytes32) leaf, bytes32[] leafProof, uint256 leafProofOrder) returns()
func (_BeefyClient *BeefyClientTransactor) SubmitFinal(opts *bind.TransactOpts, commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof, leaf BeefyClientMMRLeaf, leafProof [][32]byte, leafProofOrder *big.Int) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "submitFinal", commitment, bitfield, proofs, leaf, leafProof, leafProofOrder)
}

// SubmitFinal is a paid mutator transaction binding the contract method 0x40f29e47.
//
// Solidity: function submitFinal((uint32,uint64,(bytes2,bytes)[]) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs, (uint8,uint32,bytes32,uint64,uint32,bytes32,bytes32,bytes32) leaf, bytes32[] leafProof, uint256 leafProofOrder) returns()
func (_BeefyClient *BeefyClientSession) SubmitFinal(commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof, leaf BeefyClientMMRLeaf, leafProof [][32]byte, leafProofOrder *big.Int) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitFinal(&_BeefyClient.TransactOpts, commitment, bitfield, proofs, leaf, leafProof, leafProofOrder)
}

// SubmitFinal is a paid mutator transaction binding the contract method 0x40f29e47.
//
// Solidity: function submitFinal((uint32,uint64,(bytes2,bytes)[]) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs, (uint8,uint32,bytes32,uint64,uint32,bytes32,bytes32,bytes32) leaf, bytes32[] leafProof, uint256 leafProofOrder) returns()
func (_BeefyClient *BeefyClientTransactorSession) SubmitFinal(commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof, leaf BeefyClientMMRLeaf, leafProof [][32]byte, leafProofOrder *big.Int) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitFinal(&_BeefyClient.TransactOpts, commitment, bitfield, proofs, leaf, leafProof, leafProofOrder)
}

// SubmitInitial is a paid mutator transaction binding the contract method 0xbb51f1eb.
//
// Solidity: function submitInitial((uint32,uint64,(bytes2,bytes)[]) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) returns()
func (_BeefyClient *BeefyClientTransactor) SubmitInitial(opts *bind.TransactOpts, commitment BeefyClientCommitment, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "submitInitial", commitment, bitfield, proof)
}

// SubmitInitial is a paid mutator transaction binding the contract method 0xbb51f1eb.
//
// Solidity: function submitInitial((uint32,uint64,(bytes2,bytes)[]) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) returns()
func (_BeefyClient *BeefyClientSession) SubmitInitial(commitment BeefyClientCommitment, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitInitial(&_BeefyClient.TransactOpts, commitment, bitfield, proof)
}

// SubmitInitial is a paid mutator transaction binding the contract method 0xbb51f1eb.
//
// Solidity: function submitInitial((uint32,uint64,(bytes2,bytes)[]) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) returns()
func (_BeefyClient *BeefyClientTransactorSession) SubmitInitial(commitment BeefyClientCommitment, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitInitial(&_BeefyClient.TransactOpts, commitment, bitfield, proof)
}

// BeefyClientNewMMRRootIterator is returned from FilterNewMMRRoot and is used to iterate over the raw logs and unpacked data for NewMMRRoot events raised by the BeefyClient contract.
type BeefyClientNewMMRRootIterator struct {
	Event *BeefyClientNewMMRRoot // Event containing the contract specifics and raw log

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
func (it *BeefyClientNewMMRRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeefyClientNewMMRRoot)
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
		it.Event = new(BeefyClientNewMMRRoot)
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
func (it *BeefyClientNewMMRRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeefyClientNewMMRRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeefyClientNewMMRRoot represents a NewMMRRoot event raised by the BeefyClient contract.
type BeefyClientNewMMRRoot struct {
	MmrRoot     [32]byte
	BlockNumber uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewMMRRoot is a free log retrieval operation binding the contract event 0xd95fe1258d152dc91c81b09380498adc76ed36a6079bcb2ed31eff622ae2d0f1.
//
// Solidity: event NewMMRRoot(bytes32 mmrRoot, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) FilterNewMMRRoot(opts *bind.FilterOpts) (*BeefyClientNewMMRRootIterator, error) {

	logs, sub, err := _BeefyClient.contract.FilterLogs(opts, "NewMMRRoot")
	if err != nil {
		return nil, err
	}
	return &BeefyClientNewMMRRootIterator{contract: _BeefyClient.contract, event: "NewMMRRoot", logs: logs, sub: sub}, nil
}

// WatchNewMMRRoot is a free log subscription operation binding the contract event 0xd95fe1258d152dc91c81b09380498adc76ed36a6079bcb2ed31eff622ae2d0f1.
//
// Solidity: event NewMMRRoot(bytes32 mmrRoot, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) WatchNewMMRRoot(opts *bind.WatchOpts, sink chan<- *BeefyClientNewMMRRoot) (event.Subscription, error) {

	logs, sub, err := _BeefyClient.contract.WatchLogs(opts, "NewMMRRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeefyClientNewMMRRoot)
				if err := _BeefyClient.contract.UnpackLog(event, "NewMMRRoot", log); err != nil {
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

// ParseNewMMRRoot is a log parse operation binding the contract event 0xd95fe1258d152dc91c81b09380498adc76ed36a6079bcb2ed31eff622ae2d0f1.
//
// Solidity: event NewMMRRoot(bytes32 mmrRoot, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) ParseNewMMRRoot(log types.Log) (*BeefyClientNewMMRRoot, error) {
	event := new(BeefyClientNewMMRRoot)
	if err := _BeefyClient.contract.UnpackLog(event, "NewMMRRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeefyClientNewTicketIterator is returned from FilterNewTicket and is used to iterate over the raw logs and unpacked data for NewTicket events raised by the BeefyClient contract.
type BeefyClientNewTicketIterator struct {
	Event *BeefyClientNewTicket // Event containing the contract specifics and raw log

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
func (it *BeefyClientNewTicketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeefyClientNewTicket)
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
		it.Event = new(BeefyClientNewTicket)
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
func (it *BeefyClientNewTicketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeefyClientNewTicketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeefyClientNewTicket represents a NewTicket event raised by the BeefyClient contract.
type BeefyClientNewTicket struct {
	Relayer     common.Address
	BlockNumber uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewTicket is a free log retrieval operation binding the contract event 0xbee983fc706c692efb9b0240bddc5666c010a53af55ed5fb42d226e7e4293869.
//
// Solidity: event NewTicket(address relayer, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) FilterNewTicket(opts *bind.FilterOpts) (*BeefyClientNewTicketIterator, error) {

	logs, sub, err := _BeefyClient.contract.FilterLogs(opts, "NewTicket")
	if err != nil {
		return nil, err
	}
	return &BeefyClientNewTicketIterator{contract: _BeefyClient.contract, event: "NewTicket", logs: logs, sub: sub}, nil
}

// WatchNewTicket is a free log subscription operation binding the contract event 0xbee983fc706c692efb9b0240bddc5666c010a53af55ed5fb42d226e7e4293869.
//
// Solidity: event NewTicket(address relayer, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) WatchNewTicket(opts *bind.WatchOpts, sink chan<- *BeefyClientNewTicket) (event.Subscription, error) {

	logs, sub, err := _BeefyClient.contract.WatchLogs(opts, "NewTicket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeefyClientNewTicket)
				if err := _BeefyClient.contract.UnpackLog(event, "NewTicket", log); err != nil {
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

// ParseNewTicket is a log parse operation binding the contract event 0xbee983fc706c692efb9b0240bddc5666c010a53af55ed5fb42d226e7e4293869.
//
// Solidity: event NewTicket(address relayer, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) ParseNewTicket(log types.Log) (*BeefyClientNewTicket, error) {
	event := new(BeefyClientNewTicket)
	if err := _BeefyClient.contract.UnpackLog(event, "NewTicket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
