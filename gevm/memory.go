package gevm

type Memory struct {
	data []byte
}

func (mem *Memory) Access(offset, size uint64) (cpy []byte) {
	if size == 0 {
		return nil
	}

	if mem.Len() < int(offset+size) {
		cpy = make([]byte, offset+size)
		copy(cpy[:], mem.data[:])
		cpy = cpy[offset : offset+size]
		return
	}
	cpy = mem.data[offset : offset+size]
	return
}

func (mem *Memory) Load(offset uint64) []byte {
	return mem.Access(offset, 32)
}

func (mem *Memory) Store(offset uint64, value []byte) (expansionCost uint64) {
	// Current memory size and cost
	currentMemSize := uint64(mem.Len())
	currentCost := calcMemoryGasCost(currentMemSize)

	// New memory size needed to store value
	newMemSize := offset + uint64(len(value))

	// Handle initial allocation separately
	if currentMemSize == 0 {
		mem.data = make([]byte, 32)
		copy(mem.data, value[:])
		return calcMemoryGasCost(32)
	}

	if currentMemSize < newMemSize {
		expansionSize := newMemSize - currentMemSize
		if expansionSize > 0 {
			mem.data = append(mem.data, make([]byte, expansionSize)...)
		}
		newCost := calcMemoryGasCost(uint64(mem.Len()))
		expansionCost = newCost - currentCost
	}

	copy(mem.data[offset:newMemSize], value)
	return expansionCost
}

func (mem *Memory) Store32(offset uint64, value []byte) (expansionCost uint64) {
	// Current memory size and cost
	currentMemSize := uint64(mem.Len())
	currentCost := calcMemoryGasCost(currentMemSize)
	newMemSize := offset + 32

	// Handle initial allocation separately
	if currentMemSize == 0 {
		mem.data = make([]byte, 32)
		copy(mem.data, value[:])
		return calcMemoryGasCost(32)
	}

	if currentMemSize < newMemSize {
		expansionSize := newMemSize - currentMemSize
		if expansionSize > 0 {
			mem.data = append(mem.data, make([]byte, expansionSize)...)
		}
		newCost := calcMemoryGasCost(uint64(mem.Len()))
		expansionCost = newCost - currentCost
	}

	copy(mem.data[offset:offset+32], value)
	return expansionCost
}
func (mem *Memory) Data() []byte {
	return mem.data
}

func (mem *Memory) Len() int {
	return len(mem.data)
}

func NewMemory() *Memory {
	return &Memory{}
}
