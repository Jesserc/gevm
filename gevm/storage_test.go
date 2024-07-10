package gevm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	tests := []struct {
		name     string
		slot     int
		value    common.Hash
		testFunc func(storage *Storage, slot int, value common.Hash) (any, any) // I use (any, any) and not (common.Hash, bool) because the last test case returns (bool, bool)
		want     any
		want2    any
	}{
		{
			name:  "TestStorage_Load",
			value: common.HexToHash("0x20"),
			slot:  1,
			testFunc: func(storage *Storage, slot int, value common.Hash) (any, any) {
				storage.Store(slot, value)
				loadedValue, isWarm := storage.Load(slot)
				return loadedValue, isWarm
			},
			want:  common.HexToHash("0x20"),
			want2: true,
		},
		{
			name: "TestStorage_Load_NonExistentKey",
			slot: 1,
			testFunc: func(storage *Storage, slot int, value common.Hash) (any, any) {
				loadedValue, isWarm := storage.Load(slot)
				return loadedValue, isWarm
			},
			want:  common.Hash{},
			want2: false,
		},
		{
			name:  "TestStorage_Store",
			value: common.HexToHash("0x20"),
			slot:  1,
			testFunc: func(storage *Storage, slot int, value common.Hash) (any, any) {
				isWarm := storage.Store(slot, value)
				storedValue, _ := storage.data[slot]
				return storedValue, isWarm
			},
			want:  common.HexToHash("0x20"),
			want2: false,
		},
		{
			name:  "TestStorage_Get",
			value: common.HexToHash("0x20"),
			slot:  1,
			testFunc: func(storage *Storage, slot int, value common.Hash) (any, any) {
				storage.Store(slot, value)
				storedValue, isWarm := storage.Get(slot)
				return storedValue, isWarm
			},
			want:  common.HexToHash("0x20"),
			want2: true,
		},
		{
			name: "TestStorage_Get_NonExistentKey",
			slot: 1,
			testFunc: func(storage *Storage, slot int, value common.Hash) (any, any) {
				storedValue, isWarm := storage.Get(slot)
				return storedValue, isWarm
			},
			want:  common.Hash{},
			want2: false,
		},
		{
			name: "TestStorage_Load_WarmsUpKey",
			slot: 1,
			testFunc: func(storage *Storage, slot int, value common.Hash) (any, any) {
				_ = common.HexToHash("0x20")
				initialWarm := storage.cache[slot]
				storage.Load(slot)
				finalWarm := storage.cache[slot]
				return initialWarm, finalWarm
			},
			want:  false,
			want2: true,
		},
		{
			name:  "TestStorage_Store_WarmsUpKey",
			value: common.HexToHash("0x20"),
			slot:  1,
			testFunc: func(storage *Storage, slot int, value common.Hash) (any, any) {
				initialWarm := storage.cache[slot]
				storage.Store(slot, value)
				finalWarm := storage.cache[slot]
				return initialWarm, finalWarm
			},
			want:  false,
			want2: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := NewStorage()
			got, got2 := tt.testFunc(storage, tt.slot, tt.value)

			if tt.want != nil {
				assert.Equal(t, tt.want, got)
			}
			if tt.want2 != nil {
				assert.Equal(t, tt.want2, got2)
			}
		})
	}
}
func TestCalcSstoreGasCost(t *testing.T) {
	tests := []struct {
		name           string
		slot           int
		newValue       common.Hash
		setup          func(evm *EVM) // Setup function to initialize the EVM storage
		expectedGas    uint64
		expectedWarm   bool
		expectedRefund uint64
	}{
		{
			name:     "No-Op",
			slot:     1,
			newValue: common.HexToHash("0x1"),
			setup: func(evm *EVM) {
				evm.Storage.Store(1, common.HexToHash("0x1"))
			},
			expectedGas:    200,
			expectedWarm:   true,
			expectedRefund: 0,
		},
		{
			name:           "New Slot Creation",
			slot:           2,
			newValue:       common.HexToHash("0x1"),
			setup:          func(evm *EVM) {},
			expectedGas:    20_000,
			expectedWarm:   false,
			expectedRefund: 0,
		},
		{
			name:     "Slot Deletion",
			slot:     3,
			newValue: common.Hash{},
			setup: func(evm *EVM) {
				evm.Storage.Store(3, common.HexToHash("0x1"))
			},
			expectedGas:    5000,
			expectedWarm:   true,
			expectedRefund: 15_000,
		},
		{
			name:     "Slot Update",
			slot:     4,
			newValue: common.HexToHash("0x2"),
			setup: func(evm *EVM) {
				evm.Storage.Store(4, common.HexToHash("0x1"))
			},
			expectedGas:    5000,
			expectedWarm:   true,
			expectedRefund: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evm := setupEVM()

			// Setup the EVM storage
			tt.setup(evm)

			// Reset refund counter
			evm.Refund = 0

			// Calculate gas cost
			gasCost, isWarm := calcSstoreGasCost(evm, tt.slot, tt.newValue)

			assert.Equal(t, tt.expectedGas, gasCost, "Gas cost mismatch")
			assert.Equal(t, tt.expectedWarm, isWarm, "Warm storage mismatch")
			assert.Equal(t, tt.expectedRefund, evm.Refund, "Refund mismatch")
		})
	}
}
