package gevm

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

// Flow
func stop(evm *EVM) {
	evm.StopFlag = true
}

// Arithmetic
func add(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Add(&a, &b)) // a + b
	evm.PC++
	evm.deductGas(3)
}

func sub(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Sub(&a, &b)) // a - b
	evm.PC++
	evm.deductGas(3)
}

func mul(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Mul(&a, &b)) // a * b
	evm.PC++
	evm.deductGas(5)
}

func div(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	if b.IsZero() {
		evm.Stack.Push(uint256.NewInt(0)) // return 0 if b == 0
	} else {
		evm.Stack.Push(new(uint256.Int).Div(&a, &b)) // a / b, returns 0 if b == 0
	}
	evm.PC++
	evm.deductGas(5)
}

func sdiv(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	if b.IsZero() {
		evm.Stack.Push(uint256.NewInt(0)) // return 0 if b == 0
	} else {
		evm.Stack.Push(new(uint256.Int).SDiv(&a, &b)) // signed a / b, returns 0 if b == 0
	}
	evm.PC++
	evm.deductGas(5)
}

func mod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Mod(&a, &b)) // a % b, returns 0 if b == 0
	evm.PC++
	evm.deductGas(5)
}

func smod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SMod(&a, &b)) // signed a % b, returns 0 if b == 0
	evm.PC++
	evm.deductGas(5)
}

func addmod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	mod := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).AddMod(&a, &b, &mod)) // a + b % mod, returns 0 if mod == 0
	evm.PC++
	evm.deductGas(8)
}

func mulmod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	mod := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).MulMod(&a, &b, &mod)) // a * b % mod, returns 0 if mod == 0
	evm.PC++
	evm.deductGas(8)
}

func exp(evm *EVM) {
	a, exponent := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Exp(&a, &exponent)) // a ^ exponent
	evm.PC++
	// gas to decrement = 10 + (50 * size_in_bytes(exponent)))
	evm.deductGas(10 + (50 * uint64(exponent.ByteLen())))
}

func signextend(evm *EVM) {
	b, num := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).ExtendSign(&num, &b))
	evm.PC++
	evm.deductGas(5)
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
	evm.deductGas(3)
}

func slt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Slt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.deductGas(3)
}

func gt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Gt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.deductGas(3)
}

func sgt(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Sgt(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.deductGas(3)
}

func eq(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.Eq(&b) {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.deductGas(3)
}

func iszero(evm *EVM) {
	a := evm.Stack.Pop()
	ret := uint256.NewInt(0)
	if a.IsZero() {
		ret = uint256.NewInt(1)
	}
	evm.Stack.Push(ret)
	evm.PC++
	evm.deductGas(3)
}

// Bitwise Operations
func and(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).And(&a, &b))
	evm.PC++
	evm.deductGas(3)
}

func or(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Or(&a, &b))
	evm.PC++
	evm.deductGas(3)
}

func xor(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Xor(&a, &b))
	evm.PC++
	evm.deductGas(3)
}

func not(evm *EVM) {
	a := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Not(&a))
	evm.PC++
	evm.deductGas(3)
}

func _byte(evm *EVM) {
	i, x := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(x.Byte(&i))
	evm.PC++
	evm.deductGas(3)
}

func shl(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Lsh(&value, uint(shift.Uint64())))
	evm.PC++
	evm.deductGas(3)
}

func shr(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Rsh(&value, uint(shift.Uint64())))
	evm.PC++
	evm.deductGas(3)
}

func sar(evm *EVM) {
	shift, value := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SRsh(&value, uint(shift.Uint64())))
	evm.PC++
	evm.deductGas(3)
}

// Hash function
func keccak256(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	sizeU256 := evm.Stack.Pop()

	offset, size := offsetU256.Uint64(), sizeU256.Uint64()

	value := evm.Memory.Access(offset, size)
	hash := crypto.Keccak256(value)
	evm.Stack.Push(uint256.NewInt(0).SetBytes(hash))

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

	wordSize := toWordSize(size)
	staticGas := uint64(30)
	dynamicGas := 6*wordSize + totalMemExpansionCost

	evm.PC++
	evm.deductGas(staticGas + dynamicGas)
	dgMap[KECCAK256] = dynamicGas
}

// Ethereum environment (calldata, code, others) operations
func address(evm *EVM) {
	evm.Stack.Push(uint256.MustFromHex(evm.Sender.Hex()))
	evm.PC++
	evm.deductGas(2)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// balance pushes a mocked balance value onto the stack.
func balance(evm *EVM) {
	_ = evm.Stack.Pop() // Pop the address from the stack, though it's not used here
	evm.Stack.Push(uint256.MustFromDecimal("99999999999"))
	evm.PC++
	evm.deductGas(2600) // '2600' here represents "address access cost" and can be 100 for warm access, but we are assuming cold access since this is a mocked version.
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// origin pushes a mocked address (evm.Sender) of the transaction origin onto the stack.
func origin(evm *EVM) {
	// We're using evm.Sender because this is a mocked version, but evm.sender may not always be the same as tx.origin in real world cases.
	evm.Stack.Push(uint256.MustFromHex(evm.Sender.Hex()))
	evm.PC++
	evm.deductGas(2)
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// caller pushes a mocked address of the current transaction caller in a transaction call chain.
func caller(evm *EVM) {
	evm.Stack.Push(uint256.MustFromHex("0xBc73e0231621D6274671839f9dF8EE7E2C8A6f93"))
	evm.PC++
	evm.deductGas(2)
}

func callvalue(evm *EVM) {
	evm.Stack.Push(uint256.NewInt(evm.Value))
	evm.PC++
	evm.deductGas(2)
}

func calldatasize(evm *EVM) {
	calldatasize := uint64(len(evm.Calldata))
	evm.Stack.Push(uint256.NewInt(calldatasize))
	evm.PC++
	evm.deductGas(2)
}

func calldataload(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	offset := offsetU256.Uint64()

	calldata := getData(evm.Calldata, offset, 32)

	evm.Stack.Push(uint256.NewInt(0).SetBytes(calldata))
	evm.PC++
	evm.deductGas(3)
}

func calldatacopy(evm *EVM) {
	destMemOffsetU256 := evm.Stack.Pop()
	offsetU256 := evm.Stack.Pop()
	sizeU256 := evm.Stack.Pop()

	destMemOffset, offset, size := destMemOffsetU256.Uint64(), offsetU256.Uint64(), sizeU256.Uint64()

	calldata := getData(evm.Calldata, offset, size)
	memExpansionCost := evm.Memory.Store(destMemOffset, calldata)

	wordSize := toWordSize(size)
	staticGas := uint64(3)
	dynamicGas := staticGas*wordSize + memExpansionCost

	evm.PC++
	evm.deductGas(dynamicGas)
	dgMap[CALLDATACOPY] = dynamicGas
}

func codesize(evm *EVM) {
	codesize := uint64(len(evm.Code))
	evm.Stack.Push(uint256.NewInt(0).SetUint64(codesize))
	evm.PC++
	evm.deductGas(2)
}

func codecopy(evm *EVM) {
	destMemOffsetU256 := evm.Stack.Pop()
	offsetU256 := evm.Stack.Pop()
	sizeU256 := evm.Stack.Pop()

	destMemOffset, offset, size := destMemOffsetU256.Uint64(), offsetU256.Uint64(), sizeU256.Uint64()

	codeCopy := getData(evm.Code, offset, size)
	memExpansionCost := evm.Memory.Store(destMemOffset, codeCopy)

	wordSize := toWordSize(size)
	staticGas := uint64(3)
	dynamicGas := staticGas*wordSize + memExpansionCost

	evm.PC++
	evm.deductGas(dynamicGas)
	dgMap[CODECOPY] = dynamicGas
}

// gasprice pushes a mocked gas price (0) onto the stack.
func gasprice(evm *EVM) {
	gasPrice := uint256.NewInt(evm.Block.GasPrice)
	evm.Stack.Push(gasPrice)
	evm.PC++
	evm.deductGas(2)
}

// remaining gas (after this instruction).
func gas(evm *EVM) {
	evm.deductGas(2) // subtract gas first
	gasLeft := uint256.NewInt(evm.Gas)
	evm.Stack.Push(gasLeft)
	evm.PC++
}

// This is a mocked version and doesn't behave exactly as it would in a real EVM.
//
// extcodesize pushes a mocked code size (0) of an external account onto the stack.
func extcodesize(evm *EVM) {
	_ = evm.Stack.Pop()               // Pop external address off the stack
	evm.Stack.Push(uint256.NewInt(0)) // push 0x0
	evm.PC++
	evm.deductGas(2)
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

	wordSize := toWordSize(size)
	dynamicGas := 3*wordSize + memExpansionCost + 2600 // '2600' here represents "address access cost" and can be 100 for warm access, but we are assuming cold access since this is a mocked version.

	evm.PC++
	evm.deductGas(dynamicGas)
	dgMap[EXTCODECOPY] = dynamicGas
}

// returndatasize pushes a mocked (zero-value) length of the return data onto the stack.
func returndatasize(evm *EVM) {
	returnDataSize := uint256.NewInt(uint64(len(evm.ReturnData)))
	evm.Stack.Push(returnDataSize)
	evm.PC++
	evm.deductGas(2)
}

// returndatacopy copies a zero-length mocked return data byte slice to memory
func returndatacopy(evm *EVM) {
	destMemOffsetU256 := evm.Stack.Pop()
	_ = evm.Stack.Pop() // Pop offset of the return data to copy. This isn't used.
	sizeU256 := evm.Stack.Pop()

	destMemOffset, size := destMemOffsetU256.Uint64(), sizeU256.Uint64()

	retDataCopy := evm.ReturnData
	memExpansionCost := evm.Memory.Store(destMemOffset, retDataCopy)

	wordSize := toWordSize(size)
	staticGas := uint64(3)
	dynamicGas := staticGas*wordSize + memExpansionCost

	evm.PC++
	evm.deductGas(dynamicGas)
	dgMap[RETURNDATACOPY] = dynamicGas
}

// blockhash pushes a mocked hash of the current block onto the stack.
//
// Normally, this would be the hash of one of the 256 most recent blocks given by a block number, but for simplicity, we always use the current block.
func blockhash(evm *EVM) {
	blockNumU256 := evm.Stack.Pop()
	if blockNumU256.Uint64() > 256 {
		panic("Only the last 256 blocks can be accessed")
	}

	blockBytes, err := json.Marshal(evm.Block)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal block to JSON bytes: %v", err))
	}
	hash := crypto.Keccak256(blockBytes)
	evm.Stack.Push(uint256.NewInt(0).SetBytes(hash))
	evm.PC++
	evm.deductGas(20)
}

// coinbase pushes a block’s beneficiary address onto the stack.
func coinbase(evm *EVM) {
	coinBase := uint256.NewInt(0).SetBytes32(evm.Block.Coinbase[:])
	evm.Stack.Push(coinBase)
	evm.PC++
	evm.deductGas(2)
}

func timestamp(evm *EVM) {
	timeStamp := uint256.NewInt(uint64(evm.Block.Timestamp.Unix()))
	evm.Stack.Push(timeStamp)
	evm.PC++
	evm.deductGas(2)
}

func number(evm *EVM) {
	number := uint256.NewInt(uint64(evm.Block.Number))
	evm.Stack.Push(number)
	evm.PC++
	evm.deductGas(2)
}

func basefee(evm *EVM) {
	baseFee := uint256.NewInt(uint64(evm.Block.BaseFee))
	evm.Stack.Push(baseFee)
	evm.PC++
	evm.deductGas(2)
}

func gaslimit(evm *EVM) {
	evm.Stack.Push(uint256.NewInt(evm.ChainConfig.GasLimit))
	evm.PC++
	evm.deductGas(2)
}

func chainid(evm *EVM) {
	evm.Stack.Push(uint256.NewInt(evm.ChainConfig.ChainID))
	evm.PC++
	evm.deductGas(2)
}

// Pop, Push, Dup & swap operations
func pop(evm *EVM) {
	_ = evm.Stack.Pop()
	evm.PC++
	evm.deductGas(2)
}

func push0(evm *EVM) {
	evm.Stack.Push(uint256.NewInt(0))
	evm.PC++
	evm.deductGas(2)
}

func pushN(evm *EVM, n uint64) {
	if n < 1 || n > 32 {
		panic("Invalid push size, must be between 1 and 32")
	}
	if evm.PC+n >= uint64(len(evm.Code)) {
		panic("Push size exceeds remaining code size")
	}

	start := evm.PC + 1
	end := start + n
	dataBytes := evm.Code[start:end] // hex bytes
	v := uint256.NewInt(0).SetBytes(dataBytes)
	evm.Stack.Push(v)
	evm.PC += n + 1 // Move PC to the next opcode
	evm.deductGas(3)
}

func dupN(evm *EVM, n uint8) {
	if n < 1 || n > 16 {
		panic("Invalid dup size, must be between 1 and 16")
	}

	stackLen := len(evm.Stack.data)
	if stackLen < int(n) {
		panic("Insufficient stack size for dup operation")
	}
	// Access the n-th element from the top of the stack.
	// The stack is a slice, so elements are appended from the right
	// Index zero is the last element and index len(stack)-1 is the top and most recent element
	valueU256 := evm.Stack.data[stackLen-int(n)]
	// This achieves the same thing as above
	// valueU256 := evm.Stack.data[uint8(len(evm.Stack.data)-1)-(n-1)]

	evm.Stack.Push(&valueU256)

	evm.PC++
	evm.deductGas(3)
}

func swapN(evm *EVM, n uint8) {
	if n < 1 || n > 16 {
		panic("Invalid swap size, must be between 1 and 16")
	}

	stackLen := len(evm.Stack.data)
	if stackLen < int(n+1) {
		panic("Insufficient stack size for swap operation")
	}

	// We do this backward slice
	evm.Stack.data[stackLen-1], evm.Stack.data[stackLen-int(n+1)] = evm.Stack.data[stackLen-int(n+1)], evm.Stack.data[stackLen-1]

	evm.PC++
	evm.deductGas(3)
}

// Memory operations
func mload(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	offset := offsetU256.Uint64()

	data := evm.Memory.Load(offset)
	evm.Stack.Push(uint256.NewInt(0).SetBytes(data))

	// Gas cost calculations
	memExpansionCost := evm.Memory.Store32(offset, evm.Memory.Load(offset))
	staticGas := uint64(3)
	dynamicGas := staticGas + memExpansionCost

	evm.PC++
	evm.deductGas(dynamicGas)
	dgMap[MLOAD] = dynamicGas
}

func mstore(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	valueU256 := evm.Stack.Pop()

	data := make([]byte, 32)
	v := valueU256.Bytes32()
	copy(data, v[:])

	memExpansionCost := evm.Memory.Store32(offsetU256.Uint64(), data)
	staticGas := uint64(3)
	dynamicGas := staticGas + memExpansionCost

	evm.deductGas(dynamicGas)
	evm.PC++
	dgMap[MSTORE] = dynamicGas
}

func mstore8(evm *EVM) {
	offsetU256 := evm.Stack.Pop()
	valueU256 := evm.Stack.Pop()

	memExpansionCost := evm.Memory.Store(offsetU256.Uint64(), valueU256.Bytes()) // valueU256.Bytes() would be just a slice of one byte
	staticGas := uint64(3)
	dynamicGas := staticGas + memExpansionCost

	evm.deductGas(dynamicGas)
	evm.PC++
	dgMap[MSTORE8] = dynamicGas
}

func msize(evm *EVM) {
	evm.Stack.Push(uint256.NewInt(uint64(evm.Memory.Len())))
	evm.PC++
	evm.deductGas(2)
}

func mcopy(evm *EVM) {
	destMemOffsetU256 := evm.Stack.Pop()
	offsetU256 := evm.Stack.Pop()
	sizeU256 := evm.Stack.Pop()

	destMemOffset, offset, size := destMemOffsetU256.Uint64(), offsetU256.Uint64(), sizeU256.Uint64()

	memCopy := getData(evm.Memory.data, offset, size)
	memExpansionCost := evm.Memory.Store(destMemOffset, memCopy)

	wordSize := toWordSize(size)
	staticGas := uint64(3)
	dynamicGas := staticGas*wordSize + memExpansionCost

	evm.deductGas(dynamicGas)
	evm.PC++
	dgMap[MCOPY] = dynamicGas
}

// Storage operations
func sload(evm *EVM) {
	slotU256 := evm.Stack.Pop()
	v, isWarm := evm.Storage.Load(int(slotU256.Uint64()))

	valueU256 := uint256.NewInt(0).SetBytes32(v[:])
	evm.Stack.Push(valueU256)

	var dynamicGas uint64
	evm.PC++
	if isWarm {
		dynamicGas = 100
	} else {
		dynamicGas = 2100
	}
	evm.deductGas(dynamicGas)
	dgMap[SLOAD] = dynamicGas
}

func sstore(evm *EVM) {
	slotU256 := evm.Stack.Pop()
	valueU256 := evm.Stack.Pop()

	slot := int(slotU256.Uint64())
	newValue := common.BytesToHash(valueU256.Bytes())

	gasCost := calcSstoreGasCost(evm, slot, newValue)

	evm.deductGas(gasCost)
	evm.Storage.Store(slot, newValue)

	evm.PC++
	dgMap[SSTORE] = gasCost
}

// Transient storage operations
func tload(evm *EVM) {
	slotU256 := evm.Stack.Pop()
	v := evm.Transient.Load(int(slotU256.Uint64()))
	valueU256 := uint256.NewInt(0).SetBytes32(v[:])

	evm.deductGas(100)
	evm.Stack.Push(valueU256)
	evm.PC++
}

func tstore(evm *EVM) {
	slotU256 := evm.Stack.Pop()
	valueU256 := evm.Stack.Pop()
	v := common.BytesToHash(valueU256.Bytes())

	evm.deductGas(100)
	evm.Transient.Store(int(slotU256.Uint64()), v)
	evm.PC++
}

// Jump operations
func jump(evm *EVM) {
	newPCIndexU256 := evm.Stack.Pop()
	newPCIndex := newPCIndexU256.Uint64()
	if uint64(evm.Code[newPCIndex]) != uint64(JUMPDEST) {
		panic("Invalid jump destination")
	}
	evm.PC = newPCIndex
	evm.deductGas(8)
}

func jumpi(evm *EVM) {
	newPCIndexU256 := evm.Stack.Pop()
	valueU256 := evm.Stack.Pop()

	newPCIndex := newPCIndexU256.Uint64()

	if uint64(evm.Code[newPCIndex]) != uint64(JUMPDEST) {
		panic("Invalid jump destination")
	}

	if !valueU256.IsZero() {
		evm.PC = newPCIndex
	} else {
		evm.PC++
	}
	evm.deductGas(10)
}

func pc(evm *EVM) {
	evm.Stack.Push(uint256.NewInt(evm.PC))
	evm.PC++
	evm.deductGas(2)
}

func jumpdest(evm *EVM) {
	evm.PC++
	evm.deductGas(1)
}

// Execution control
func invalid(evm *EVM) {
	evm.StopFlag = true
	evm.Gas = 0 // consume available gas
}

func revert(evm *EVM) {
	destMemOffsetU256, sizeU256 := evm.Stack.Pop(), evm.Stack.Pop()

	destMemOffset, size := destMemOffsetU256.Uint64(), sizeU256.Uint64()
	evm.ReturnData = evm.Memory.Access(destMemOffset, size)

	evm.RevertFlag = true
	// evm.PC++
}

func _return(evm *EVM) {
	destMemOffsetU256, sizeU256 := evm.Stack.Pop(), evm.Stack.Pop()

	destMemOffset, size := destMemOffsetU256.Uint64(), sizeU256.Uint64()
	evm.ReturnData = evm.Memory.Access(destMemOffset, size)
	evm.StopFlag = true
	// evm.PC++
}

// Logging
func log0(evm *EVM) {
	offsetU256, sizeU256 := evm.Stack.Pop(), evm.Stack.Pop()

	offset, size := offsetU256.Uint64(), sizeU256.Uint64()
	data := evm.Memory.Access(offset, size)
	evm.LogRecord.AddLog([]common.Hash{}, data) // change []common.Hash{} to nil?

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

	evm.PC++
	dynamicGas := calcLogGasCost(0, size, totalMemExpansionCost)
	evm.deductGas(dynamicGas)
	dgMap[LOG0] = dynamicGas
}

func log1(evm *EVM) {
	offsetU256, sizeU256, topicU256 := evm.Stack.Pop(), evm.Stack.Pop(), evm.Stack.Pop()

	offset, size := offsetU256.Uint64(), sizeU256.Uint64()
	data := evm.Memory.Access(offset, size)
	topic := common.BytesToHash(topicU256.Bytes())
	evm.LogRecord.AddLog([]common.Hash{topic}, data)

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

	evm.PC++
	dynamicGas := calcLogGasCost(1, size, totalMemExpansionCost)
	evm.deductGas(dynamicGas)
	dgMap[LOG1] = dynamicGas
}

func log2(evm *EVM) {
	offsetU256, sizeU256 := evm.Stack.Pop(), evm.Stack.Pop()
	topic1U256, topic2U256 := evm.Stack.Pop(), evm.Stack.Pop()

	offset, size := offsetU256.Uint64(), sizeU256.Uint64()
	data := evm.Memory.Access(offset, size)
	topic1, topic2 := common.BytesToHash(topic1U256.Bytes()), common.BytesToHash(topic2U256.Bytes())
	evm.LogRecord.AddLog([]common.Hash{topic1, topic2}, data)

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

	evm.PC++
	dynamicGas := calcLogGasCost(2, size, totalMemExpansionCost)
	evm.deductGas(dynamicGas)
	dgMap[LOG2] = dynamicGas
}

func log3(evm *EVM) {
	offsetU256, sizeU256 := evm.Stack.Pop(), evm.Stack.Pop()
	topic1U256, topic2U256, topic3U256 := evm.Stack.Pop(), evm.Stack.Pop(), evm.Stack.Pop()

	offset, size := offsetU256.Uint64(), sizeU256.Uint64()
	data := evm.Memory.Access(offset, size)
	topic1, topic2, topic3 := common.BytesToHash(topic1U256.Bytes()), common.BytesToHash(topic2U256.Bytes()), common.BytesToHash(topic3U256.Bytes())
	evm.LogRecord.AddLog([]common.Hash{topic1, topic2, topic3}, data)

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

	evm.PC++
	dynamicGas := calcLogGasCost(3, size, totalMemExpansionCost)
	evm.deductGas(dynamicGas)
	dgMap[LOG3] = dynamicGas
}

func log4(evm *EVM) {
	offsetU256, sizeU256 := evm.Stack.Pop(), evm.Stack.Pop()
	topic1U256, topic2U256, topic3U256, topic4U256 := evm.Stack.Pop(), evm.Stack.Pop(), evm.Stack.Pop(), evm.Stack.Pop()

	offset, size := offsetU256.Uint64(), sizeU256.Uint64()
	data := evm.Memory.Access(offset, size)
	topic1, topic2, topic3, topic4 := common.BytesToHash(topic1U256.Bytes()), common.BytesToHash(topic2U256.Bytes()), common.BytesToHash(topic3U256.Bytes()), common.BytesToHash(topic4U256.Bytes())
	evm.LogRecord.AddLog([]common.Hash{topic1, topic2, topic3, topic4}, data)

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

	evm.PC++
	dynamicGas := calcLogGasCost(3, size, totalMemExpansionCost)
	evm.deductGas(dynamicGas)
	dgMap[LOG4] = dynamicGas
}

// This is used in jump_table.go
func generatePushNFunc(size uint8) func(*EVM) {
	return func(evm *EVM) {
		pushN(evm, uint64(size))
	}
}

func generateSwapNFunc(index uint8) func(*EVM) {
	return func(evm *EVM) {
		swapN(evm, index)
	}
}

func generateDupNFunc(index uint8) func(*EVM) {
	return func(evm *EVM) {
		dupN(evm, index)
	}
}
