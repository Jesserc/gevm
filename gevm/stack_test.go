package gevm

import (
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name       string
		operations func(st *Stack) interface{}
		expected   interface{}
	}{
		{
			name: "Push and Pop",
			operations: func(st *Stack) interface{} {
				value := uint256.NewInt(1)
				st.Push(value)
				item := st.Pop()
				return item.Uint64()
			},
			expected: uint64(1),
		},
		{
			name: "Peek",
			operations: func(st *Stack) interface{} {
				value := uint256.NewInt(2)
				st.Push(value)
				item := st.Peek()
				return item.Uint64()
			},
			expected: uint64(2),
		},
		{
			name: "ToString",
			operations: func(st *Stack) interface{} {
				values := []uint64{1, 2, 3}
				for _, v := range values {
					st.Push(uint256.NewInt(v))
				}
				return st.ToString()
			},
			expected: "[0x3, 0x2, 0x1]",
		},
		{
			name: "Stack Underflow on Pop",
			operations: func(st *Stack) interface{} {
				defer func() { // We can either use this 'recover' approach or assert.Panics(...)
					if r := recover(); r != nil {
						assert.Equal(t, ErrStackUnderflow.Error(), r)
					}
				}()
				return st.Pop()
			},
			expected: nil, // we expect panic
		},
		{
			name: "Stack Underflow on Peek",
			operations: func(st *Stack) interface{} {
				defer func() { // We can either use this 'recover' approach or assert.Panics(...)
					if r := recover(); r != nil {
						assert.Equal(t, ErrStackUnderflow.Error(), r)
					}
				}()
				return st.Peek()
			},
			expected: nil, // we expect panic
		},
		{
			name: "Stack Overflow",
			operations: func(st *Stack) interface{} {
				defer func() { // We can either use this 'recover' approach or assert.Panics(...)
					if r := recover(); r != nil {
						assert.Equal(t, ErrStackOverflow.Error(), r)
					}
				}()
				for i := 0; i <= MAX_STACK_SIZE; i++ { // MAX_STACK_SIZE+1 items to cause a panic
					st.Push(uint256.NewInt(uint64(i)))
				}
				return nil
			},
			expected: nil, // we expect panic
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack()
			result := tt.operations(stack)
			if tt.expected != nil {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
