package gevm

import (
	"github.com/ethereum/go-ethereum/common"
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
	evm.PC += 1
	evm.gasDec(3)
}

func sub(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Sub(&a, &b)) // a - b
	evm.PC += 1
	evm.gasDec(3)
}

func mul(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Mul(&a, &b)) // a * b
	evm.PC += 1
	evm.gasDec(3)
}

func div(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Div(&a, &b)) // a / b, returns 0 if b == 0
	evm.PC += 1
	evm.gasDec(5)
}

func sdiv(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SDiv(&a, &b)) // signed a / b, returns 0 if b == 0
	evm.PC += 1
	evm.gasDec(5)
}

func mod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Mod(&a, &b)) // a % b, returns 0 if b == 0
	evm.PC += 1
	evm.gasDec(5)
}

func smod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SMod(&a, &b)) // signed a % b, returns 0 if b == 0
	evm.PC += 1
	evm.gasDec(5)
}

func addmod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	mod := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).AddMod(&a, &b, &mod)) // a + b % mod, returns 0 if mod == 0
	evm.PC += 1
	evm.gasDec(8)
}

func mulmod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	mod := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).MulMod(&a, &b, &mod)) // a * b % mod, returns 0 if mod == 0
	evm.PC += 1
	evm.gasDec(8)
}

func exp(evm *EVM) {
	a, exponent := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Exp(&a, &exponent)) // a ^ exponent
	evm.PC += 1
	// gas to decrement = 10 + (50 * size_in_bytes(exponent)))
	evm.gasDec(10 + (50 * uint64(exponent.ByteLen())))
}

func signextend(evm *EVM) {
	b, num := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).ExtendSign(&num, &b))
	evm.PC += 1
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
	evm.PC += 1
	evm.gasDec(3)
}

func slt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Slt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC += 1
	evm.gasDec(3)
}

func gt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Gt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC += 1
	evm.gasDec(3)
}

func sgt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Sgt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC += 1
	evm.gasDec(3)
}

func eq(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Eq(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC += 1
	evm.gasDec(3)
}

func iszero(evm *EVM) {
	a := evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.IsZero() {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC += 1
	evm.gasDec(3)
}

// Bitwise Operations
func and(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).And(&a, &b))
	evm.PC += 1
	evm.gasDec(3)
}

func or(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Or(&a, &b))
	evm.PC += 1
	evm.gasDec(3)
}

func xor(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Xor(&a, &b))
	evm.PC += 1
	evm.gasDec(3)
}

func not(evm *EVM) {
	a := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Not(&a))
	evm.PC += 1
	evm.gasDec(3)
}

func _byte(evm *EVM) {
	i, x := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(x.Byte(&i))
	evm.PC += 1
	evm.gasDec(3)
}

func shl(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Lsh(&value, uint(shift.Uint64())))
	evm.PC += 1
	evm.gasDec(3)
}

func shr(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Rsh(&value, uint(shift.Uint64())))
	evm.PC += 1
	evm.gasDec(3)
}

func sar(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SRsh(&value, uint(shift.Uint64())))
	evm.PC += 1
	evm.gasDec(3)
}

// Hash function
func keccak256(evm *EVM) {
	offset, size := evm.Stack.Pop(), evm.Stack.Pop()
	value := evm.Memory.Access(offset.Uint64(), size.Uint64())
	hash := crypto.Keccak256(value)
	evm.Stack.Push(uint256.NewInt(0).SetBytes(hash))
	evm.PC += 1

	// Gas cost calculations
	currentMemSize := uint64(evm.Memory.Len())
	currentMemCost := calcMemoryGasCost(currentMemSize)
	newMemSize := offset.Uint64() + size.Uint64()

	var memExpansionSize uint64
	if currentMemSize < newMemSize {
		memExpansionSize = newMemSize - currentMemSize
	}

	newMemCost := calcMemoryGasCost(currentMemSize + memExpansionSize)
	totalMemExpansionCost := newMemCost - currentMemCost

	minWordSize := toWordSize(size.Uint64())
	staticGas := uint64(30)
	dynamicGas := 6*minWordSize + totalMemExpansionCost
	evm.gasDec(staticGas + dynamicGas)
}

// Ethereum environment opcodes
func address(evm *EVM) {
	evm.Stack.Push(uint256.MustFromHex(evm.Sender.Hex()))
	evm.PC += 1
	evm.gasDec(3)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// balance pushes a mocked balance value onto the stack.
func balance(evm *EVM) {
	_ = evm.Stack.Pop() // Pop the address from the stack, though it's not used here
	evm.Stack.Push(uint256.MustFromDecimal("99999999999"))
	evm.PC += 1

	// Decrease the gas by the amount needed for a cold balance check
	// Note: Normally, the gas cost could be 100 for warm access, but we are assuming cold access since this is a mocked version
	evm.gasDec(2600)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// origin pushes a mocked address (evm.Sender) of the transaction origin onto the stack.
func origin(evm *EVM) {
	// We're using evm.Sender because this is a mocked version, and evm.sender not always be the same as tx.origin in real world cases.
	evm.Stack.Push(uint256.MustFromHex(evm.Sender.Hex()))
	evm.PC += 1
	evm.gasDec(2)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// caller pushes a mocked address of the current transaction caller in a transaction call chain.
func caller(evm *EVM) {
	evm.Stack.Push(uint256.MustFromHex("0xBc73e0231621D6274671839f9dF8EE7E2C8A6f93"))
	evm.PC += 1
	evm.gasDec(2)
}

func callvalue(evm *EVM) {
	evm.Stack.Push(evm.Value)
	evm.PC += 1
	evm.gasDec(2)
}

func calldataload(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	offset := offsetU256.Uint64()

	delta := uint64(0)
	if offset+32 > uint64(len(evm.Calldata)) {
		delta = offset + 32 - uint64(len(evm.Calldata))
	}

	// calldata := make([]byte, 32)
	var calldata []byte
	if offset < uint64(len(evm.Calldata)) {
		// copy(calldata, evm.Calldata[i:min(i+32-delta, uint64(len(evm.Calldata)))]) // we can use this commented out code to achieve the same thing
		calldata = common.RightPadBytes(evm.Calldata[offset:min(offset+32-delta, uint64(len(evm.Calldata)))], 32)
	}

	evm.Stack.Push(uint256.NewInt(0).SetBytes(calldata))
	evm.PC++
	evm.gasDec(3)
}

func calldatasize(evm *EVM) {
	calldatasize := uint64(len(evm.Calldata))
	evm.Stack.Push(uint256.NewInt(calldatasize))
	evm.PC += 1
	evm.gasDec(2)
}

func calldatacopy(evm *EVM) {
	destMemOffsetU256 := evm.Stack.Pop()
	offsetU256 := evm.Stack.Pop() // 31
	sizeU256 := evm.Stack.Pop()   // 8

	destMemOffset, offset, size := destMemOffsetU256.Uint64(), offsetU256.Uint64(), sizeU256.Uint64()

	length := uint64(len(evm.Calldata))
	if offset > length {
		offset = length
	}

	end := offset + size
	if end > length {
		end = length
	}
	calldata := evm.Calldata[offset:end]
	memExpansionCost := evm.Memory.Store(destMemOffset, calldata)

	minWordSize := toWordSize(size)
	staticGas := uint64(3)
	dynamicGas := 3*minWordSize + memExpansionCost

	evm.PC += 1
	evm.gasDec(staticGas + dynamicGas)
}

func codesize(evm *EVM) {
	codesize := uint64(len(evm.Program))
	evm.Stack.Push(uint256.NewInt(0).SetUint64(codesize))
	evm.PC += 1
	evm.gasDec(2)
}
