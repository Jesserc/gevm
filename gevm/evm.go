package gevm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

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

type ExecutionEnvironment struct {
	Stack     *Stack
	Memory    *Memory
	Storage   *Storage
	Transient *TransientStorage
}

type TransactionContext struct {
	Sender   common.Address // [20]byte, might change to common.Hash?
	Value    uint64
	Calldata []byte
}

type ChainConfig struct {
	ChainID  uint64
	GasLimit uint64
}

type EVM struct {
	ExecutionContext
	ExecutionEnvironment
	TransactionContext
	ChainConfig
}

func (evm *EVM) gasDec(gas uint64) {
	if evm.Gas < gas {
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
	jumpTable := NewJumpTable()
	index := 0
	for evm.continueExecution() {
		// fmt.Println("PC:", evm.PC)
		opcode := evm.Code[evm.PC]
		opStr := Opcode(opcode).String()
		index = int(evm.PC)
		fmt.Printf("%.2v -> %v\n", index, opStr)

		if opFunc, exists := jumpTable[Opcode(opcode)]; exists {
			opFunc(evm)
		} else {
			fmt.Printf("Unknown opcode: %#x\n", opcode)
			return
		}
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
	evm = New(common.Address{}, 0, 0, 0, 0, []byte{}, []byte{})
}

func New(sender common.Address, gas, value, chainID, gasLimit uint64, code, calldata []byte) *EVM {
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
