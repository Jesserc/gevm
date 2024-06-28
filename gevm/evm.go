package gevm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type EVM struct {
	PC      uint64
	Stack   *Stack
	Memory  *Memory
	Storage *Storage

	Sender   common.Address // [20]byte, might change to common.Hash?
	Code     []byte
	Gas      uint64
	Value    *uint256.Int
	Calldata []byte

	StopFlag   bool
	RevertFlag bool

	ReturnData []byte
	Logs       []byte
}

func (evm *EVM) gasDec(gas uint64) {
	if evm.Gas < gas {
		panic(fmt.Errorf("out of gas: tried to consume %d gas, but only %d gas remaining", gas, evm.Gas))
	}
	evm.Gas -= gas // decrement gas
}

func NewEVM(sender common.Address, gas uint64, value *uint256.Int, code, calldata []byte) *EVM {
	return &EVM{
		PC: 0,
		Stack: &Stack{
			data: make([]uint256.Int, 0),
		},
		Memory: &Memory{
			data: make([]byte, 0),
		},
		Storage: &Storage{
			data:  make(map[int]common.Hash),
			cache: make([]int, 0),
		},
		Sender:     sender,
		Code:       code,
		Gas:        gas,
		Value:      value,
		Calldata:   calldata,
		StopFlag:   false,
		RevertFlag: false,
		ReturnData: []byte{},
		Logs:       []byte{},
	}
}
