package gevm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/holiman/uint256"
)

var (
	ErrStackOverflow  = errors.New("stack overflow")
	ErrStackUnderflow = errors.New("stack underflow")
)

const MAX_STACK_SIZE = 1024

type Stack struct {
	data []uint256.Int
}

func (st *Stack) Push(value *uint256.Int) {
	if len(st.data) == MAX_STACK_SIZE {
		panic(ErrStackOverflow.Error())
	}
	st.data = append(st.data, *value)
}

func (st *Stack) Pop() uint256.Int {
	if len(st.data) == 0 {
		panic(ErrStackUnderflow.Error())
	}
	ret := st.data[len(st.data)-1]
	st.data = (st.data)[:len(st.data)-1]
	return ret
}

func (st *Stack) Peek() uint256.Int {
	if len(st.data) == 0 {
		panic(ErrStackUnderflow.Error())
	}
	ret := st.data[len(st.data)-1]
	return ret
}

func (st Stack) ToString() string {
	var d string
	if len(st.data) == 0 {
		d = "[]"
		return d
	}
	for i := len(st.data) - 1; i >= 0; i-- {
		if i == len(st.data)-1 {
			d = "["
		}
		d += fmt.Sprintf("%v, ", st.data[i].Hex())
		if i == 0 {
			d = strings.TrimRight(d, ", ")
			d += "]"
		}
	}
	return d
}

func (evm *EVM) Data() []uint256.Int {
	return evm.Stack.data
}

func NewStack() *Stack {
	return &Stack{data: make([]uint256.Int, 0)}
}
