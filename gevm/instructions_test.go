package gevm

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func setupEVM() *EVM {
	block := NewBlock(common.HexToAddress("0x"), 2, 1, 0, 1, time.Now())
	return NewEVM(common.HexToAddress("0x"), 1000, 2e5, 1, 8000000, []byte{}, []byte{}, block)
}

// Table-driven tests for EVM operations
func TestEVMOperations(t *testing.T) {
	tests := []struct {
		name         string
		initialStack []*uint256.Int
		op           func(evm *EVM)
		expected     *uint256.Int
		expectedPC   uint64
		expectedGas  uint64
	}{
		{
			name:         "Add positive numbers",
			initialStack: []*uint256.Int{uint256.NewInt(1), uint256.NewInt(2)},
			op:           add,
			expected:     uint256.NewInt(3),
			expectedPC:   1,
			expectedGas:  997,
		},
		{
			name:         "Sub positive numbers",
			initialStack: []*uint256.Int{uint256.NewInt(3), uint256.NewInt(5)},
			op:           sub,
			expected:     uint256.NewInt(2),
			expectedPC:   1,
			expectedGas:  997,
		},
		{
			name:         "Mul positive numbers",
			initialStack: []*uint256.Int{uint256.NewInt(3), uint256.NewInt(4)},
			op:           mul,
			expected:     uint256.NewInt(12),
			expectedPC:   1,
			expectedGas:  995,
		},
		{
			name:         "Div positive numbers",
			initialStack: []*uint256.Int{uint256.NewInt(2), uint256.NewInt(10)},
			op:           div,
			expected:     uint256.NewInt(5),
			expectedPC:   1,
			expectedGas:  995,
		},
		{
			name:         "And operation",
			initialStack: []*uint256.Int{uint256.NewInt(3), uint256.NewInt(6)},
			op:           and,
			expected:     uint256.NewInt(2),
			expectedPC:   1,
			expectedGas:  997,
		},
		{
			name:         "Or operation",
			initialStack: []*uint256.Int{uint256.NewInt(3), uint256.NewInt(6)},
			op:           or,
			expected:     uint256.NewInt(7),
			expectedPC:   1,
			expectedGas:  997,
		},
		{
			name:         "Xor operation",
			initialStack: []*uint256.Int{uint256.NewInt(3), uint256.NewInt(6)},
			op:           xor,
			expected:     uint256.NewInt(5),
			expectedPC:   1,
			expectedGas:  997,
		},
		{
			name:         "Not operation",
			initialStack: []*uint256.Int{uint256.NewInt(1)},
			op:           not,
			expected:     uint256.MustFromHex("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe"),
			expectedPC:   1,
			expectedGas:  997,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evm := setupEVM()
			for _, val := range tt.initialStack {
				evm.Stack.Push(val)
			}
			tt.op(evm)

			result := evm.Stack.Pop()
			assert.True(t, result.Eq(tt.expected), tt.name+" failed")
			assert.Equal(t, tt.expectedPC, evm.PC, "Program counter not incremented correctly in "+tt.name)
			assert.Equal(t, tt.expectedGas, evm.Gas, "Gas not deducted correctly in "+tt.name)
		})
	}
}
