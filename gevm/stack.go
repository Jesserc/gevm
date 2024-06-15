package gevm

import (
	"errors"
	"fmt"
)

var (
	ErrStackOverflow  = errors.New("stack overflow")
	ErrStackUnderflow = errors.New("stack underflow")
)

const MAX_STACK_SIZE = 1024

type Stack struct {
	data []byte
}

func (st *Stack) Push(value []byte) {
	if len(st.data) == MAX_STACK_SIZE {
		panic(ErrStackOverflow.Error())
	}
	st.data = append(st.data, value...)
}

func (st *Stack) Pop() []byte {
	if len(st.data) == 0 {
		panic(ErrStackUnderflow.Error())
	}
	element := st.data[len(st.data)-1]
	st.data = (st.data)[:len(st.data)-1]
	return []byte{element}
}

func (st Stack) String() string {
	var d string
	for i := len(st.data) - 1; i >= 0; i-- {
		if i == len(st.data)-1 {
			d += fmt.Sprintf("%v <first\n", st.data[i])
		} else if i == 0 {
			d += fmt.Sprintf("%v <last", st.data[i])
		} else {
			d += fmt.Sprintf("%v\n", st.data[i])
		}
	}
	return d
}

func NewStack() *Stack {
	return &Stack{}
}
