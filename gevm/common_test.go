package gevm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToWordSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		testFunc func() uint64
		want     uint64
	}{
		{
			name: "Test_ToWordSize_32_Bytes",
			size: 32,
			want: 1,
		},
		{
			name: "Test_ToWordSize_96_Bytes",
			size: 96,
			want: 3,
		},
		{
			name: "Test_ToWordSize_0_Bytes",
			size: 0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]byte, tt.size)
			wordSize := toWordSize(uint64(len(data)))
			assert.Equal(t, tt.want, wordSize, "")
		})
	}
}

func TestCalcMemoryGasCost(t *testing.T) {
	tests := []struct {
		size uint64
		want uint64
	}{
		{size: 0, want: 0},
		{size: 1, want: 3},
		{size: 32, want: 3},
		{size: 33, want: 6},
		{size: 64, want: 6},
		{size: 65, want: 9},
		{size: 1024, want: 96 + (32*32)/512},
		{size: 1048576, want: 98304 + (32768*32768)/512},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Size: %v", tt.size), func(t *testing.T) {
			got := calcMemoryGasCost(tt.size)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalcLogGasCost(t *testing.T) {
	tests := []struct {
		memExpansionCost uint64
		topicCount       uint64
		sizeInMem        uint64
		want             uint64
	}{
		{sizeInMem: 0, topicCount: 1, memExpansionCost: 0, want: 375},
		{sizeInMem: 1, topicCount: 1, memExpansionCost: 3, want: 386},
		{sizeInMem: 64, topicCount: 1, memExpansionCost: 6, want: 893},
		{sizeInMem: 1024, topicCount: 1, memExpansionCost: 98, want: 8665},
		{sizeInMem: 2048, topicCount: 1, memExpansionCost: 2195456, want: 2212215},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Log size in memory: %v", tt.sizeInMem), func(t *testing.T) {
			got := calcLogGasCost(tt.topicCount, tt.sizeInMem, tt.memExpansionCost)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetData(t *testing.T) {
	tests := []struct {
		name  string
		data  []byte
		start uint64
		size  uint64
		want  []byte
	}{
		{
			name:  "Normal case",
			data:  []byte{1, 2, 3, 4, 5},
			start: 1,
			size:  3,
			want:  []byte{2, 3, 4},
		},
		{
			name:  "Start out of bounds",
			data:  []byte{1, 2, 3, 4, 5},
			start: 10,
			size:  3,
			want:  []byte{0, 0, 0},
		},
		{
			name:  "End out of bounds",
			data:  []byte{1, 2, 3, 4, 5},
			start: 3,
			size:  10,
			want:  []byte{4, 5, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:  "Start and end out of bounds",
			data:  []byte{1, 2, 3, 4, 5},
			start: 10,
			size:  10,
			want:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:  "Zero size",
			data:  []byte{1, 2, 3, 4, 5},
			start: 2,
			size:  0,
			want:  []byte{},
		},
		{
			name:  "Empty data",
			data:  []byte{},
			start: 0,
			size:  3,
			want:  []byte{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getData(tt.data, tt.start, tt.size)
			assert.Equal(t, tt.want, result, "Test case %s failed", tt.name)
		})
	}
}
