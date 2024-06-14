package gevm

import (
	"errors"
	"fmt"
)

var (
	ErrStackOverflow  = errors.New("Stack overflow")
	ErrStackUnderflow = errors.New("Stack underflow")
)

const MAX_STACK_SIZE = 1024

type Stack []any

func (s *Stack) Push(value any) {
	if len(*s) == MAX_STACK_SIZE {
		panic(ErrStackOverflow.Error())
	}
	*s = append(*s, value)
}

func (s *Stack) Pop() any {
	if len(*s) == 0 {
		panic(ErrStackUnderflow.Error())
	}

	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element
}

func (s Stack) String() string {
	var d string
	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			d += fmt.Sprintf("%v <first\n", s[i])
		} else if i == 0 {
			d += fmt.Sprintf("%v <last", s[i])
		} else {
			d += fmt.Sprintf("%v\n", s[i])
		}
	}
	return d
}

func NewStack(size int) *Stack {
	s := make(Stack, size)
	return &s
}
