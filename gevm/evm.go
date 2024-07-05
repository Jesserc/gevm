package gevm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type EVM struct {
	PC        uint64
	Stack     *Stack
	Memory    *Memory
	Storage   *Storage
	Transient *TransientStorage

	Sender   common.Address // [20]byte, might change to common.Hash?
	Code     []byte
	Gas      uint64
	Value    uint64
	Calldata []byte

	StopFlag   bool
	RevertFlag bool

	ReturnData []byte
	LogRecord  *LogRecord
}

func (evm *EVM) gasDec(gas uint64) {
	if evm.Gas < gas {
		panic(fmt.Errorf("out of gas: tried to consume %d gas, but only %d gas remaining", gas, evm.Gas))
	}
	evm.Gas -= gas // decrement gas
}

func NewEVM(sender common.Address, gas, value uint64, code, calldata []byte) *EVM {
	return &EVM{
		PC:         0,
		Stack:      NewStack(),
		Memory:     NewMemory(),
		Storage:    NewStorage(),
		Transient:  NewTransientStorage(),
		Sender:     sender,
		Code:       code,
		Gas:        gas,
		Value:      value,
		Calldata:   calldata,
		StopFlag:   false,
		RevertFlag: false,
		ReturnData: []byte{},
		LogRecord:  NewLogRecord(),
	}
}
