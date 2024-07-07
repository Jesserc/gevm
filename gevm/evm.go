package gevm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type dynamicGasMap map[Opcode]uint64

func (d dynamicGasMap) Gas(op Opcode) (hasDyGas bool, gas uint64) {
	gas, hasDyGas = d[op]
	return hasDyGas, gas
}

var (
	dgMap = make(dynamicGasMap)
)

// ExecutionContext represents the execution context during EVM execution.
type ExecutionContext struct {
	PC         uint64
	Code       []byte
	Gas        uint64
	Refund     uint64
	StopFlag   bool
	RevertFlag bool
	ReturnData []byte
	LogRecord  *LogRecord
}

// ExecutionEnvironment encapsulates the EVM execution environment including the stack, memory, storage, and transient storage.
type ExecutionEnvironment struct {
	Stack     *Stack
	Memory    *Memory
	Storage   *Storage
	Transient *TransientStorage
}

// TransactionContext holds transaction-specific information during EVM execution.
type TransactionContext struct {
	Sender   common.Address // [20]byte, might change to common.Hash?
	Value    uint64
	Calldata []byte
}

// ChainConfig stores network configuration parameters.
type ChainConfig struct {
	ChainID  uint64
	GasLimit uint64
}

// EVM represents an Ethereum Virtual Machine instance.
type EVM struct {
	ExecutionContext
	ExecutionEnvironment
	TransactionContext
	ChainConfig
}

func (evm *EVM) deductGas(gas uint64) {
	if evm.Gas < gas || evm.Gas <= 0 {
		panic(fmt.Errorf("out of gas: tried to consume %d gas, but only %d gas remaining", gas, evm.Gas))
	}
	evm.Gas -= gas // decrement gas
}
func (evm *EVM) continueExecution() bool {
	if int(evm.PC) > len(evm.Code)-1 {
		return false
	}
	if evm.StopFlag {
		return false
	}
	if evm.RevertFlag {
		return false
	}
	if evm.Gas == 0 {
		return false
	}
	return true
}

func (evm *EVM) Run() {
	fmt.Println("#### Trace ####")

	jumpTable := NewJumpTable()
	for evm.continueExecution() {
		currentPC := evm.PC
		opcode := evm.Code[currentPC]
		op := Opcode(opcode)
		if opFunc, exists := jumpTable[Opcode(opcode)]; exists {
			opFunc(evm)
		} else {
			fmt.Printf("Unknown opcode: %#x\n", opcode)
			return
		}

		fmt.Println("Opcode:", op)
		// fmt.Println("Value:",)
		fmt.Println("Stack:", evm.Stack.ToString())
		var gCost uint64
		if has, cost := dgMap.Gas(op); has {
			gCost = cost
		} else {
			gCost = op.Gas()
		}

		fmt.Println("Gas Cost:", gCost)
		fmt.Println("Memory:", hexutil.Encode(evm.Memory.data))
		fmt.Println("Storage:", evm.Storage.data)
		fmt.Println("Return Data:", hexutil.Encode(evm.ReturnData))
		fmt.Println("PC:", currentPC)
		fmt.Println()
	}
}

func (evm *EVM) AddRefund(refund uint64) {
	evm.Refund += refund
}

func (evm *EVM) SubRefund(gas uint64) {
	if gas > evm.Refund {
		panic(fmt.Sprintf("Refund counter below zero (gas: %d > refund: %d)", gas, evm.Refund))
	}
	evm.Refund -= gas
}

func (evm *EVM) Reset() {
	//lint:ignore SA4006 ignore unused code warning for this variable
	evm = NewEVM(common.Address{}, 0, 0, 0, 0, []byte{}, []byte{})
}

func NewEVM(sender common.Address, gas, value, chainID, gasLimit uint64, code, calldata []byte) *EVM {
	return &EVM{
		ExecutionContext: ExecutionContext{
			PC:         0,
			Code:       []byte{},
			Gas:        gas,
			Refund:     0,
			StopFlag:   false,
			RevertFlag: false,
			ReturnData: []byte{},
			LogRecord:  nil,
		},
		ExecutionEnvironment: ExecutionEnvironment{
			Stack:     NewStack(),
			Memory:    NewMemory(),
			Storage:   NewStorage(),
			Transient: NewTransientStorage(),
		},
		TransactionContext: TransactionContext{
			Sender:   common.Address{},
			Value:    0,
			Calldata: []byte{},
		},
		ChainConfig: ChainConfig{
			ChainID:  1,
			GasLimit: gas,
		},
	}
}
