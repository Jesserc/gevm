package gevm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpcodeString(t *testing.T) {
	tests := []struct {
		op       Opcode
		expected string
	}{
		{STOP, "STOP"},
		{ADD, "ADD"},
		{SUB, "SUB"},
		// ... other opcodes
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.op.String())
		})
	}
}

func TestOpcodeGas(t *testing.T) {
	tests := []struct {
		op       Opcode
		expected uint64
	}{
		{STOP, 0},
		{ADD, 3},
		{SUB, 3},
		// ... other opcodes
	}

	for _, tt := range tests {
		t.Run(tt.op.String()+" Gas", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.op.Gas())
		})
	}
}
