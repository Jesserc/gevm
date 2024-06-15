package gevm

import (
	"math"
)

type Memory struct {
	data []byte
}

func (mem *Memory) Access(offset, size int) (nMem []byte) {
	if mem.Len() < offset+size {
		nMem = make([]byte, offset+size)
		copy(nMem[:], mem.data[:])
		nMem = nMem[offset : offset+size]
		return
	}
	nMem = mem.data[offset : offset+size]
	return
}

func (mem *Memory) Load(offset int) []byte {
	return mem.Access(offset, 32)
}

func (mem *Memory) Store(offset int, value []byte) int {
	expansionCost := 0 // memory expansion cost

	if mem.Len() <= offset+len(value) {
		expansionSize := 0

		if mem.Len() == 0 {
			expansionSize = 32
			mem.data = make([]byte, 32)
		}

		if mem.Len() < offset+len(value) {
			expansionSize += (offset + len(value)) - mem.Cap()
			if expansionSize > 0 {
				mem.data = append(mem.data, make([]byte, expansionSize)...)
				/*
					// original memory expansion logic
					def calc_memory_expansion_gas(memory_byte_size):
					memory_size_word = (memory_byte_size + 31) / 32
					memory_cost = (memory_size_word ** 2) / 512 + (3 * memory_size_word)
					return round(memory_cost)

				*/
				expansionCost = int(math.Pow(float64(expansionSize), 2)) // simplified expansion cost
			}
		}
	}
	copy(mem.data[offset:offset+len(value)], value)
	return expansionCost
}

func (mem *Memory) Data() []byte {
	return mem.data
}

func (mem *Memory) Len() int {
	return len(mem.data)
}

func (mem *Memory) Cap() int {
	return cap(mem.data)
}

func NewMemory() *Memory {
	return &Memory{}
}
