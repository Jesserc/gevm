package gevm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type EVM struct {
	PC      int
	Stack   Stack
	Memory  Memory
	Storage Storage

	Sender   common.Address // [20]byte, might change to common.Hash?
	Program  []byte
	Gas      int
	Value    int
	Calldata []byte

	StopFlag   bool
	RevertFlag bool

	ReturnData []byte
	Logs       []byte
}

func (evm *EVM) gasDec(gas int) {
	if evm.Gas < gas {
		panic(fmt.Errorf("out of gas: tried to consume %d gas, but only %d gas remaining", gas, evm.Gas))
	}
	evm.Gas -= gas // decrement gas
}

func NewEVM(sender common.Address, gas, value int, program, calldata []byte) *EVM {
	return &EVM{
		PC: 0,
		Stack: Stack{
			data: make([]uint256.Int, 0),
		},
		Memory: Memory{},
		Storage: Storage{
			data:  make(map[int]common.Hash),
			cache: make([]int, 0),
		},
		Sender:     sender,
		Program:    program,
		Gas:        gas,
		Value:      value,
		Calldata:   calldata,
		StopFlag:   false,
		RevertFlag: false,
		ReturnData: []byte{},
		Logs:       []byte{},
	}
}
