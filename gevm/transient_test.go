package gevm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTransientStorage(t *testing.T) {
	tests := []struct {
		name     string
		slot     int
		value    common.Hash
		testFunc func(ts *TransientStorage, slot int, value common.Hash) any // I use ( any) and not (common.Hash, bool) because the last test case returns (bool, bool)
		want     any
	}{
		{
			name:  "TestTransientStorage_Load",
			slot:  0,
			value: common.HexToHash("0x20"),
			testFunc: func(ts *TransientStorage, slot int, value common.Hash) any {
				ts.Store(slot, value)
				loadedValue := ts.Load(slot)
				return loadedValue
			},
			want: common.HexToHash("0x20"),
		},
		{
			name: "TestTransientStorage_Load_NonExistentKey",
			slot: 5,
			testFunc: func(ts *TransientStorage, slot int, value common.Hash) any {
				loadedValue := ts.Load(slot)
				return loadedValue
			},
			want: common.Hash{},
		},
		{
			name:  "TestTransientStorage_Store",
			slot:  256,
			value: common.HexToHash("0xa"),
			testFunc: func(ts *TransientStorage, slot int, value common.Hash) any {
				ts.Store(slot, value)
				storedValue := ts.data[slot]
				return storedValue
			},
			want: common.HexToHash("0xa"),
		},
		{
			name:  "TestTransientStorage_Clear",
			slot:  0,
			value: common.HexToHash("0x20"),
			testFunc: func(ts *TransientStorage, slot int, value common.Hash) any {
				ts.Store(slot, value)
				// Clear the storage
				ts.Clear()
				loadedValue := ts.Load(slot)
				return loadedValue
			},
			want: common.Hash{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := NewTransientStorage()
			got := tt.testFunc(ts, tt.slot, tt.value)

			if tt.want != nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
