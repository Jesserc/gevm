package gevm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	tests := []struct {
		name     string
		offset   uint64
		size     uint64
		value    []byte
		testFunc func(mem *Memory, offset uint64, size uint64, value []byte) (got1 any, got2 any)
		want     any
		want2    any
	}{
		{
			name:   "TestMemory_Access",
			offset: 0,
			size:   4,
			value:  []byte{0x01, 0x02, 0x03, 0x04},
			testFunc: func(mem *Memory, offset uint64, size uint64, value []byte) (any, any) {
				mem.Store(offset, value)
				return mem.Access(offset, size), nil
			},
			want:  []byte{0x01, 0x02, 0x03, 0x04},
			want2: nil,
		},
		{
			name:   "TestMemory_Load",
			offset: 0,
			size:   32,
			value:  common.RightPadBytes([]byte{0x01, 0x02, 0x03, 0x04}, 32),
			testFunc: func(mem *Memory, offset uint64, size uint64, value []byte) (any, any) {
				mem.Store(offset, value)
				return mem.Load(offset), nil
			},
			want:  common.RightPadBytes([]byte{0x01, 0x02, 0x03, 0x04}, 32),
			want2: nil,
		},
		{
			name:   "TestMemory_Store",
			offset: 0,
			size:   4,
			value:  []byte{0x01, 0x02, 0x03, 0x04},
			testFunc: func(mem *Memory, offset uint64, size uint64, value []byte) (any, any) {
				expansionCost := mem.Store(offset, value)
				return expansionCost, mem.Access(offset, size)
			},
			want:  uint64(3),
			want2: []byte{0x01, 0x02, 0x03, 0x04},
		},
		{
			name:   "TestMemory_Store32",
			offset: 0,
			size:   32,
			value:  common.RightPadBytes([]byte{0x01, 0x02, 0x03, 0x04}, 32),
			testFunc: func(mem *Memory, offset uint64, size uint64, value []byte) (any, any) {
				expansionCost := mem.Store32(offset, value)
				return expansionCost, mem.Access(offset, size)
			},
			want:  uint64(3),
			want2: common.RightPadBytes([]byte{0x01, 0x02, 0x03, 0x04}, 32),
		},
		{
			name:   "TestMemory_Data",
			offset: 0,
			size:   4,
			value:  []byte{0x01, 0x02, 0x03, 0x04},
			testFunc: func(mem *Memory, offset uint64, size uint64, value []byte) (any, any) {
				mem.Store(offset, value)
				return mem.Data()[:4], nil
			},
			want:  []byte{0x01, 0x02, 0x03, 0x04},
			want2: nil,
		},
		{
			name:   "TestMemory_Len",
			offset: 0,
			size:   0,
			value:  nil,
			testFunc: func(mem *Memory, offset uint64, size uint64, value []byte) (any, any) {
				initialLen := mem.Len()
				mem.Store(offset, []byte{0x01, 0x02, 0x03, 0x04})
				finalLen := mem.Len()
				return initialLen, finalLen
			},
			want:  0,
			want2: 32, // 32 bytes of memory is always created at first initialization
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewMemory()
			got, got2 := tt.testFunc(mem, tt.offset, tt.size, tt.value)
			assert.Equal(t, tt.want, got)
			if tt.want2 != nil {
				assert.Equal(t, tt.want2, got2)
			}
		})
	}
}
