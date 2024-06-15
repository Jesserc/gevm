package gevm

import (
	"github.com/holiman/uint256"
)

func stop(evm *EVM) {
	evm.StopFlag = true
}

func add(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Add(a, b)) // a + b
	evm.PC += 1
	evm.gasDec(3)
}

func sub(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Sub(a, b)) // a - b
	evm.PC += 1
	evm.gasDec(3)
}

func mul(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Mul(a, b)) // a * b
	evm.PC += 1
	evm.gasDec(3)
}

func div(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Div(a, b)) // a / b, returns 0 if b == 0
	evm.PC += 1
	evm.gasDec(5)
}

func sdiv(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SDiv(a, b)) // signed a / b, returns 0 if b == 0
	evm.PC += 1
	evm.gasDec(5)
}

func mod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Mod(a, b)) // a % b, returns 0 if b == 0
	evm.PC += 1
	evm.gasDec(5)
}

func smod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).SMod(a, b)) // signed a % b, returns 0 if b == 0
	evm.PC += 1
	evm.gasDec(5)
}

func addmod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	mod := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).AddMod(a, b, mod)) // a + b % mod, returns 0 if mod == 0
	evm.PC += 1
	evm.gasDec(8)
}

func mulmod(evm *EVM) {
	a, b := evm.Stack.Pop(), evm.Stack.Pop()
	mod := evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).MulMod(a, b, mod)) // a * b % mod, returns 0 if mod == 0
	evm.PC += 1
	evm.gasDec(8)
}

func exp(evm *EVM) {
	a, exponent := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).Exp(a, exponent)) // a ^ exponent
	evm.PC += 1
	// gas to decrement = 10 + (50 * size_in_bytes(exponent)))
	evm.gasDec(10 + (50 * exponent.ByteLen()))
}

func signextend(evm *EVM) {
	b, num := evm.Stack.Pop(), evm.Stack.Pop()
	evm.Stack.Push(new(uint256.Int).ExtendSign(num, b))
	evm.PC += 1
	evm.gasDec(5)
}
