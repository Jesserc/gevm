package main

import (
	"fmt"

	"github.com/Jesserc/gevm/gevm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
)

func main() {

	fmt.Println("=== EVM State Before Operations ===")
	evm := gevm.NewEVM(common.HexToAddress("0x"), 21000, uint256.NewInt(2e5), []byte{}, []byte{})
	fmt.Printf("EVM: %v\n", evm.Stack)
	fmt.Println()
	fmt.Println()
	/*
		fmt.Println("=== Stack ===")
		// stack := gevm.NewStack()
		evm.Stack.Push(uint256.MustFromDecimal("1"))
		evm.Stack.Push(uint256.MustFromDecimal("2"))
		evm.Stack.Push(uint256.MustFromDecimal("3"))

		fmt.Println(evm.Stack.ToString())
		fmt.Println()

		fmt.Println("=== Memory ===") */
	// memory := gevm.NewMemory()
	// data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// evm.Memory.Store(0, data)
	// fmt.Println(evm.Memory.Access(0, 10))
	// fmt.Println(evm.Memory.Load(9)) // [10..31]

	// evm.Memory.Store(5, []byte{11, 12, 13, 14, 15})
	// fmt.Println(evm.Memory.Access(0, 10))
	// fmt.Println(evm.Memory.Load(0)) //[0..32]

	/* 	epc := evm.Memory.Store(0, []byte{0x01, 0x02, 0x03, 0x04})
	   	fmt.Println("First evm.Memory expansion cost:", epc) // 0, 32 memSize is allocated for free first
	   	data := evm.Memory.Load(0)
	   	fmt.Println(data)

	   	fmt.Println(evm.Memory.Cap(), "old cap")
	   	// epc = evm.Memory.Store(0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33})
	   	epc = evm.Memory.Store(0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100})
	   	fmt.Println("Second evm.Memory expansion cost:", epc)
	   	fmt.Println(evm.Memory.Cap(), "new cap")
	   	fmt.Println()

	   	fmt.Println("=== Storage ===")
	   	// storage := gevm.NewStorage()

	   	warm, v := evm.Storage.Load(0)                // 0x0
	   	fmt.Println("Warm access for slot 0?:", warm) // false
	   	fmt.Println("Storage value:", v)

	   	evm.Storage.Store(0, common.HexToHash("0xa")) // store 10
	   	warm, v = evm.Storage.Load(0)                 // 0xa
	   	fmt.Println("Warm access for slot 0?:", warm) // true
	   	fmt.Println("Storage value:", v)

	   	evm.Storage.Store(292929, common.HexToHash("0x20")) // store 10
	   	warm, v = evm.Storage.Load(292929)
	   	fmt.Println("Warm access for slot 292929?:", warm) // false
	   	fmt.Println("Storage value:", v)                   // [0]
	   	fmt.Println()
	*/
	cost := evm.Memory.Store(0x0, hexutil.MustDecode("0xFFFFFFFF"))
	fmt.Println("Memory expansion cost:", cost)
	// evm.Memory.Store(0x0, []byte("0xFFFFFFFF"))
	// evm.Stack.Push(uint256.MustFromHex("0x4"))
	evm.Calldata = []byte("0xffff")
	evm.Stack.Push(uint256.MustFromHex("0x0"))

	fmt.Println("=== EVM State After Operations ===")
	fmt.Printf("EVM Stack: %v\n", evm.Stack)
	fmt.Printf("EVM Memory: %v\n", evm.Memory)
	fmt.Printf("EVM: %v\n", evm)

	// fmt.Println()
}
