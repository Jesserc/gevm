package gevm

import (
	"github.com/ethereum/go-ethereum/common"
)

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

// calcSstoreGasCost calculates the gas cost for the SSTORE operation in the EVM.
//
// This is a bit simplified compared to the actual implementation in Ethereum clients.
func calcSstoreGasCost(evm *EVM, slot int, newValue common.Hash) (gasCost uint64) {
	// Load the current value stored at the specified slot.
	currentValue, isWarm := evm.Storage.Get(slot)

	// Determine the access cost (cold vs warm)
	accessCost := uint64(2100)
	if isWarm {
		accessCost = 100
	}

	// If the current value is the same as the new value, it's a no-op.
	if currentValue == newValue {
		return accessCost
	}

	// Calculate the dynamic gas cost
	var dynamicCost uint64
	if currentValue == (common.Hash{}) && newValue != (common.Hash{}) {
		// Zero to non-zero
		dynamicCost = 20000
	} else if currentValue != (common.Hash{}) && newValue == (common.Hash{}) {
		// Non-zero to zero
		dynamicCost = 2900
		evm.addRefund(4800)
	} else {
		// Non-zero to non-zero
		dynamicCost = 2900
	}

	return accessCost + dynamicCost
}

func calcLogGasCost(topicCount, size, memExpansionCost uint64) uint64 {
	staticGas := uint64(375)
	return staticGas*topicCount + 8*size + memExpansionCost
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
