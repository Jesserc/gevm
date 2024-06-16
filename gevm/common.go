package gevm

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
