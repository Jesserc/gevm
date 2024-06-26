package gevm

import (
	"errors"
	"fmt"

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

func (st Stack) ToString() string {
	var d string
	for i := len(st.data) - 1; i >= 0; i-- {
		if i == len(st.data)-1 {
			d += fmt.Sprintf("%v <first\n", st.data[i].String())
		} else if i == 0 {
			d += fmt.Sprintf("%v <last", st.data[i].String())
		} else {
			d += fmt.Sprintf("%v\n", st.data[i].String())
		}
	}
	return d
}

func NewStack() *Stack {
	return &Stack{}
}
