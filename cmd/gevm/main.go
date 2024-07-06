package main

import (
	"fmt"

	"github.com/Jesserc/gevm/gevm"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	evm := gevm.New(common.HexToAddress("0x"), 500000, 2e5, 1, 8000000, []byte{}, []byte{})

	// evm.Stack.Push(uint256.NewInt(0).SetBytes(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000FFFFFFFF")))
	// evm.Stack.Push(uint256.NewInt(0).SetBytes(hexutil.MustDecode("0xFFFFFFFF00000000000000000000000000000000000000000000000000000000")))
	// evm.Stack.Push(uint256.NewInt(32))

	SIMPLE_ADD := []byte{0x60, 0x42, 0x60, 0xFF, 0x01}
	SIMPLE_PUSH := []byte{0x60, 0x42}
	evm.Code = SIMPLE_PUSH
	evm.Code = SIMPLE_ADD

	SIMPLE_REVERT := []byte{0x60, 0x1f, 0x60, 0x01, 0x01, 0x60, 0x00, 0x60, 0x00, 0xFD, 0x60, 0x20}
	evm.Code = SIMPLE_REVERT

	SIMPLE_STORE := []byte{0x60, 0x20, 0x5f, 0x55}
	SIMPLE_STORE2 := []byte{0x60, 0x20, 0x5f, 0x55, 0x60, 0xa, 0x5f, 0x55}
	evm.Code = SIMPLE_STORE
	evm.Code = SIMPLE_STORE2

	evm.Run()
	fmt.Printf("EVM stack: \n%v\n", evm.Stack.ToString())

	fmt.Println()
	// fmt.Println()
	// d := evm.Data()[0]
	// fmt.Println("EVM: ", &d)
	fmt.Println("EVM storage: ", evm.Storage)
}
