package gevm

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func stop(evm *EVM) {
	evm.StopFlag = true
}

// Arithmetic
func add(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Add(&a, &b)) // a + b
	evm.PC++
	evm.gasDec(3)
}

func sub(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Sub(&a, &b)) // a - b
	evm.PC++
	evm.gasDec(3)
}

func mul(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Mul(&a, &b)) // a * b
	evm.PC++
	evm.gasDec(3)
}

func div(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Div(&a, &b)) // a / b, returns 0 if b == 0
	evm.PC++
	evm.gasDec(5)
}

func sdiv(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SDiv(&a, &b)) // signed a / b, returns 0 if b == 0
	evm.PC++
	evm.gasDec(5)
}

func mod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Mod(&a, &b)) // a % b, returns 0 if b == 0
	evm.PC++
	evm.gasDec(5)
}

func smod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SMod(&a, &b)) // signed a % b, returns 0 if b == 0
	evm.PC++
	evm.gasDec(5)
}

func addmod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	mod := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).AddMod(&a, &b, &mod)) // a + b % mod, returns 0 if mod == 0
	evm.PC++
	evm.gasDec(8)
}

func mulmod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	mod := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).MulMod(&a, &b, &mod)) // a * b % mod, returns 0 if mod == 0
	evm.PC++
	evm.gasDec(8)
}

func exp(evm *EVM) {
	a, exponent := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Exp(&a, &exponent)) // a ^ exponent
	evm.PC++
	// gas to decrement = 10 + (50 * size_in_bytes(exponent)))
	evm.gasDec(10 + (50 * uint64(exponent.ByteLen())))
}

func signextend(evm *EVM) {
	b, num := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).ExtendSign(&num, &b))
	evm.PC++
	evm.gasDec(5)
}

// Comparisons
func lt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Lt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.gasDec(3)
}

func slt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Slt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.gasDec(3)
}

func gt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Gt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.gasDec(3)
}

func sgt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Sgt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.gasDec(3)
}

func eq(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Eq(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.gasDec(3)
}

func iszero(evm *EVM) {
	a := evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.IsZero() {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.gasDec(3)
}

// Bitwise Operations
func and(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).And(&a, &b))
	evm.PC++
	evm.gasDec(3)
}

func or(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Or(&a, &b))
	evm.PC++
	evm.gasDec(3)
}

func xor(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Xor(&a, &b))
	evm.PC++
	evm.gasDec(3)
}

func not(evm *EVM) {
	a := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Not(&a))
	evm.PC++
	evm.gasDec(3)
}

func _byte(evm *EVM) {
	i, x := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(x.Byte(&i))
	evm.PC++
	evm.gasDec(3)
}

func shl(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Lsh(&value, uint(shift.Uint64())))
	evm.PC++
	evm.gasDec(3)
}

func shr(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Rsh(&value, uint(shift.Uint64())))
	evm.PC++
	evm.gasDec(3)
}

func sar(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SRsh(&value, uint(shift.Uint64())))
	evm.PC++
	evm.gasDec(3)
}

// Hash function
func keccak256(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	sizeU256 := evm.Stack.Pop()

	offset, size := offsetU256.Uint64(), sizeU256.Uint64()

	value := evm.Memory.Access(offset, size)
	hash := crypto.Keccak256(value)
	evm.Stack.Push(uint256.NewInt(0).SetBytes(hash))
	evm.PC++

	// Gas cost calculations
	currentMemSize := uint64(evm.Memory.Len())
	currentMemCost := calcMemoryGasCost(currentMemSize)
	newMemSize := offset + size

	var memExpansionSize uint64
	if currentMemSize < newMemSize {
		memExpansionSize = newMemSize - currentMemSize
	}

	newMemCost := calcMemoryGasCost(currentMemSize + memExpansionSize)
	totalMemExpansionCost := newMemCost - currentMemCost

	minWordSize := toWordSize(size)
	staticGas := uint64(30)
	dynamicGas := 6*minWordSize + totalMemExpansionCost
	evm.gasDec(staticGas + dynamicGas)
}

// Ethereum environment opcodes
func address(evm *EVM) {
	evm.Stack.Push(uint256.MustFromHex(evm.Sender.Hex()))
	evm.PC++
	evm.gasDec(3)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// balance pushes a mocked balance value onto the stack.
func balance(evm *EVM) {
	_ = evm.Stack.Pop() // Pop the address from the stack, though it's not used here
	evm.Stack.Push(uint256.MustFromDecimal("99999999999"))
	evm.PC++
	evm.gasDec(2600) // '2600' here represents "address access cost" and can be 100 for warm access, but we are assuming cold access since this is a mocked version.
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// origin pushes a mocked address (evm.Sender) of the transaction origin onto the stack.
func origin(evm *EVM) {
	// We're using evm.Sender because this is a mocked version, but evm.sender may not always be the same as tx.origin in real world cases.
	evm.Stack.Push(uint256.MustFromHex(evm.Sender.Hex()))
	evm.PC++
	evm.gasDec(2)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// caller pushes a mocked address of the current transaction caller in a transaction call chain.
func caller(evm *EVM) {
	evm.Stack.Push(uint256.MustFromHex("0xBc73e0231621D6274671839f9dF8EE7E2C8A6f93"))
	evm.PC++
	evm.gasDec(2)
}

func callvalue(evm *EVM) {
	evm.Stack.Push(evm.Value)
	evm.PC++
	evm.gasDec(2)
}

func calldataload(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	offset := offsetU256.Uint64()

	calldata := getData(evm.Calldata, offset, 32)

	evm.Stack.Push(uint256.NewInt(0).SetBytes(calldata))
	evm.PC++
	evm.gasDec(3)
}

func calldatasize(evm *EVM) {
	calldatasize := uint64(len(evm.Calldata))
	evm.Stack.Push(uint256.NewInt(calldatasize))
	evm.PC++
	evm.gasDec(2)
}

func calldatacopy(evm *EVM) {
	destMemOffsetU256 := evm.Stack.Pop()
	offsetU256 := evm.Stack.Pop()
	sizeU256 := evm.Stack.Pop()

	destMemOffset, offset, size := destMemOffsetU256.Uint64(), offsetU256.Uint64(), sizeU256.Uint64()

	calldata := getData(evm.Calldata, offset, size)
	memExpansionCost := evm.Memory.Store(destMemOffset, calldata)

	minWordSize := toWordSize(size)
	staticGas := uint64(3)
	dynamicGas := staticGas*minWordSize + memExpansionCost

	evm.PC++
	evm.gasDec(dynamicGas)
}

func codesize(evm *EVM) {
	codesize := uint64(len(evm.Code))
	evm.Stack.Push(uint256.NewInt(0).SetUint64(codesize))
	evm.PC++
	evm.gasDec(2)
}

func codecopy(evm *EVM) {
	destMemOffsetU256 := evm.Stack.Pop()
	offsetU256 := evm.Stack.Pop()
	sizeU256 := evm.Stack.Pop()

	destMemOffset, offset, size := destMemOffsetU256.Uint64(), offsetU256.Uint64(), sizeU256.Uint64()

	codeCopy := getData(evm.Code, offset, size)
	memExpansionCost := evm.Memory.Store(destMemOffset, codeCopy)

	minWordSize := toWordSize(size)
	staticGas := uint64(3)
	dynamicGas := staticGas*minWordSize + memExpansionCost

	evm.PC++
	evm.gasDec(dynamicGas)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// gasprice pushes a mocked gas price (0) onto the stack.
func gasprice(evm *EVM) {
	evm.Stack.Push(uint256.NewInt(0)) // push 0x0
	evm.PC++
	evm.gasDec(2)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// extcodesize pushes a mocked code size (0) of an external account onto the stack.
func extcodesize(evm *EVM) {
	_ = evm.Stack.Pop()               // Pop external address off the stack
	evm.Stack.Push(uint256.NewInt(0)) // push 0x0
	evm.PC++
	evm.gasDec(2)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// extcodecopy copies a zero-length mocked byte slice to memory.
func extcodecopy(evm *EVM) {
	_ = evm.Stack.Pop() // Pop external address off the stack
	destMemOffsetU256 := evm.Stack.Pop()
	_ = evm.Stack.Pop() // Pop offset of the external code to copy. This isn't used.
	sizeU256 := evm.Stack.Pop()

	destMemOffset, size := destMemOffsetU256.Uint64(), sizeU256.Uint64()

	extCodeCopy := []byte{}                                          // mocked (no external code)
	memExpansionCost := evm.Memory.Store(destMemOffset, extCodeCopy) // extCode has zero length

	minWordSize := toWordSize(size)
	dynamicGas := 3*minWordSize + memExpansionCost + 2600 // '2600' here represents "address access cost" and can be 100 for warm access, but we are assuming cold access since this is a mocked version.

	evm.PC++
	evm.gasDec(dynamicGas)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// returndatasize pushes a mocked (zero-value) length of the return data onto the stack.
func returndatasize(evm *EVM) {
	evm.Stack.Push(uint256.NewInt(0))
	evm.PC++
	evm.gasDec(2)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// returndatacopy copies a zero-length mocked return data byte slice to memory
func returndatacopy(evm *EVM) {
	destMemOffsetU256 := evm.Stack.Pop()
	_ = evm.Stack.Pop() // Pop offset of the return data to copy. This isn't used.
	sizeU256 := evm.Stack.Pop()

	destMemOffset, size := destMemOffsetU256.Uint64(), sizeU256.Uint64()

	retDataCopy := []byte{}
	memExpansionCost := evm.Memory.Store(destMemOffset, retDataCopy)

	minWordSize := toWordSize(size)
	staticGas := uint64(3)
	dynamicGas := staticGas*minWordSize + memExpansionCost

	evm.PC++
	evm.gasDec(dynamicGas)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// blockhash pushes the hash (mocked hash) of one of the 256 most recent complete blocks onto the stack.
func blockhash(evm *EVM) {
	blockNumU256 := evm.Stack.Pop()
	if blockNumU256.Uint64() > 256 {
		panic("Only the last 256 blocks can be accessed")
	}
	evm.Stack.Push(uint256.MustFromHex("0x29045A592007D0C246EF02C2223570DA9522D0CF0F73282C79A1BC8F0BB2C238")) // push a mocked block hash
	evm.PC++
	evm.gasDec(20)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// coinbase pushes a block’s beneficiary address onto the stack.
func coinbase(evm *EVM) {
	evm.Stack.Push(uint256.MustFromHex("0x29045A592007D0C246EF02C2223570DA9522D0CF0F73282C79A1BC8F0BB2C238")) // push a mocked coinbase address
	evm.PC++
	evm.gasDec(2)
}

// Pop & Push opcodes
func pop(evm *EVM) {
	_ = evm.Stack.Pop()
	evm.PC += 2 // Increment the program counter by two to account for the POP opcode and the corresponding item being removed
	evm.gasDec(2)
}

func pushN(evm *EVM, size uint64) {
	if size == 0 {
		evm.Stack.Push(uint256.NewInt(0))
		evm.PC += 1
		evm.gasDec(2)
		return
	}
	if size > 32 {
		panic("exceeded the maximum allowable size of 32 bytes (full word)")
	}
	if size > uint64(len(evm.Code)) {
		// size = uint64(len(evm.Code)) - 1
		panic("push size exceeds remaining code size")
	}

	start := evm.PC + 1
	end := start + size

	if len(evm.Code) == 1 {
		panic("code size is one")
	}
	if start >= uint64(len(evm.Code)) {
		panic("invalid push, no code left")
	}
	dataBytes := evm.Code[start:end] // hex bytes
	v := uint256.NewInt(0).SetBytes(dataBytes)
	evm.Stack.Push(v)
	evm.PC += size + 1 // Move PC to the next opcode
	evm.gasDec(3)
}

// Memory operations
func mload(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	offset := offsetU256.Uint64()

	data := evm.Memory.Load(offset) // we will trim then zeros if any and left pad it before pushing to stack
	evm.Stack.Push(uint256.NewInt(0).SetBytes32(data))

	// Gas cost calculations
	totalMemExpansionCost := evm.Memory.Store(offset, evm.Memory.Load(offset))
	staticGas := uint64(3)
	dynamicGas := staticGas + totalMemExpansionCost
	evm.gasDec(staticGas + dynamicGas)
}

/*


	padSize := size // 22
	fmt.Println("evm code:", evm.Code)

	item := getData(evm.Code, uint64(evm.PC+1), uint64(size))
	fmt.Println("item before:", item)

	item = common.RightPadBytes(item, int(padSize))

	fmt.Println("item after:", item)
	evm.Stack.Push(uint256.NewInt(0).SetBytes(item))


*/
