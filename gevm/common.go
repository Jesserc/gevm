package gevm

import "github.com/ethereum/go-ethereum/common"

// toWordSize returns the number of 32-byte words required to hold a given size in bytes.
// In the EVM, 1 word is equal to 32 bytes.
// The function ensures that any fraction of a word will count as a full word.
func toWordSize(size uint64) uint64 {
	return (size + 31) / 32
}

// calcMemoryCost calculates the gas cost for memory expansion in the EVM.
// The cost is based on the number of 32-byte words required to store the given size in bytes.
// The formula used is from the Ethereum Yellow Paper:
// cost = (3 * words) + (words * words) / 512
// where 'words' is the number of 32-byte words required.
func calcMemoryGasCost(size uint64) uint64 {
	words := toWordSize(size)
	return (3 * words) + (words*words)/512
}

// getData returns a slice from the data based on the start and size and pads
// up to size with zero's. This function is overflow safe.
func getData(data []byte, start uint64, size uint64) []byte {
	length := uint64(len(data))
	if start > length {
		start = length
	}
	end := start + size
	if end > length {
		end = length
	}
	return common.RightPadBytes(data[start:end], int(size))
}

func calcLogGasCost(topicCount, size, memExpansionCost uint64) uint64 {
	staticGas := uint64(375)
	return staticGas*topicCount + 8*size + memExpansionCost
}
