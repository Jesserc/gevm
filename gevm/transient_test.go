package gevm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTransientStorage(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func(*TransientStorage) any // I use ( any) and not (common.Hash, bool) because the last test case returns (bool, bool)
		expected any
	}{
		{
			name: "TestTransientStorage_Load",
			testFunc: func(ts *TransientStorage) any {
				key := 1
				value := common.HexToHash("0x20")
				ts.Store(key, value)
				loadedValue := ts.Load(key)
				return loadedValue
			},
			expected: common.HexToHash("0x20"),
		},
		{
			name: "TestTransientStorage_Load_NonExistentKey",
			testFunc: func(ts *TransientStorage) any {
				key := 1
				loadedValue := ts.Load(key)
				return loadedValue
			},
			expected: common.Hash{},
		},
		{
			name: "TestTransientStorage_Store",
			testFunc: func(ts *TransientStorage) any {
				key := 1
				value := common.HexToHash("0x20")
				ts.Store(key, value)
				storedValue := ts.data[key]
				return storedValue
			},
			expected: common.HexToHash("0x20"),
		},
		{
			name: "TestTransientStorage_Clear",
			testFunc: func(ts *TransientStorage) any {
				key := 1
				value := common.HexToHash("0x1")
				ts.Store(key, value)

				// Clear the storage
				ts.Clear()
				loadedValue := ts.Load(key)
				return loadedValue
			},
			expected: common.Hash{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := NewTransientStorage()
			got := tt.testFunc(ts)

			if tt.expected != nil {
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}

func TestTransientStorage_Load(t *testing.T) {
	storage := NewTransientStorage()
	key := 1
	value := common.HexToHash("0x1")

	// Store the value first
	storage.Store(key, value)

	// Load the value
	loadedValue := storage.Load(key)

	assert.Equal(t, value, loadedValue, "Loaded value should be equal to stored value")
}

func TestTransientStorage_Load_NonExistentKey(t *testing.T) {
	storage := NewTransientStorage()
	key := 1

	// Load a value for a key that doesn't exist
	loadedValue := storage.Load(key)

	assert.Equal(t, common.Hash{}, loadedValue, "Loaded value for non-existent key should be an empty hash")
}

func TestTransientStorage_Store(t *testing.T) {
	storage := NewTransientStorage()
	key := 1
	value := common.HexToHash("0x1")

	// Store the value
	storage.Store(key, value)

	// Verify the stored value
	storedValue := storage.data[key]
	assert.Equal(t, value, storedValue, "Stored value should be equal to the input value")
}

func TestTransientStorage_Clear(t *testing.T) {
	storage := NewTransientStorage()
	key := 1
	value := common.HexToHash("0x1")

	// Store the value
	storage.Store(key, value)

	// Clear the storage
	storage.Clear()

	// Verify the storage is cleared
	loadedValue := storage.Load(key)
	assert.Equal(t, common.Hash{}, loadedValue, "Loaded value after clear should be an empty hash")
}
