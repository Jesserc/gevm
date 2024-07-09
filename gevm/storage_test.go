package gevm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	tests := []struct {
		name      string
		testFunc  func(*Storage) (any, any) // I use (any, any) and not (common.Hash, bool) because the last test case returns (bool, bool)
		expected  any
		expected2 any
	}{
		{
			name: "TestStorage_Load",
			testFunc: func(storage *Storage) (any, any) {
				key := 1
				value := common.HexToHash("0x20")
				storage.Store(key, value)
				loadedValue, isWarm := storage.Load(key)
				return loadedValue, isWarm
			},
			expected:  common.HexToHash("0x20"),
			expected2: true,
		},
		{
			name: "TestStorage_Load_NonExistentKey",
			testFunc: func(storage *Storage) (any, any) {
				key := 1
				loadedValue, isWarm := storage.Load(key)
				return loadedValue, isWarm
			},
			expected:  common.Hash{},
			expected2: false,
		},
		{
			name: "TestStorage_Store",
			testFunc: func(storage *Storage) (any, any) {
				key := 1
				value := common.HexToHash("0x20")
				isWarm := storage.Store(key, value)
				storedValue, _ := storage.data[key]
				return storedValue, isWarm
			},
			expected:  common.HexToHash("0x20"),
			expected2: false,
		},
		{
			name: "TestStorage_Get",
			testFunc: func(storage *Storage) (any, any) {
				key := 1
				value := common.HexToHash("0x20")
				storage.Store(key, value)
				storedValue, isWarm := storage.Get(key)
				return storedValue, isWarm
			},
			expected:  common.HexToHash("0x20"),
			expected2: true,
		},
		{
			name: "TestStorage_Get_NonExistentKey",
			testFunc: func(storage *Storage) (any, any) {
				key := 1
				storedValue, isWarm := storage.Get(key)
				return storedValue, isWarm
			},
			expected:  common.Hash{},
			expected2: false,
		},
		{
			name: "TestStorage_Load_WarmsUpKey",
			testFunc: func(storage *Storage) (any, any) {
				key := 1
				_ = common.HexToHash("0x20")
				initialWarm := storage.cache[key]
				storage.Load(key)
				finalWarm := storage.cache[key]
				return initialWarm, finalWarm
			},
			expected:  false,
			expected2: true,
		},
		{
			name: "TestStorage_Store_WarmsUpKey",
			testFunc: func(storage *Storage) (any, any) {
				key := 1
				value := common.HexToHash("0x20")
				initialWarm := storage.cache[key]
				storage.Store(key, value)
				finalWarm := storage.cache[key]
				return initialWarm, finalWarm
			},
			expected:  false,
			expected2: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := NewStorage()
			got, got2 := tt.testFunc(storage)

			if tt.expected != nil {
				assert.Equal(t, tt.expected, got)
			}
			if tt.expected2 != nil {
				assert.Equal(t, tt.expected2, got2)
			}
		})
	}
}
