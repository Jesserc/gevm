package main

import (
	"fmt"

	"github.com/Jesserc/gevm/gevm"
)

func main() {
	fmt.Println("=== Stack ===")
	stack := gevm.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack)
	fmt.Println()

	fmt.Println("=== Memory ===")
	memory := gevm.NewMemory()
	// data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// memory.Store(0, data)
	// fmt.Println(memory.Access(0, 10))
	// fmt.Println(memory.Load(9)) // [10..31]

	// memory.Store(5, []byte{11, 12, 13, 14, 15})
	// fmt.Println(memory.Access(0, 10))
	// fmt.Println(memory.Load(0)) //[0..32]

	memory.Store(0, []byte{0x01, 0x02, 0x03, 0x04})
	data := memory.Load(0)
	fmt.Println(data)

	epc := memory.Store(0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100})
	fmt.Println("Memory expansion cost:", epc)
	fmt.Println()

	fmt.Println("=== Storage ===")
	storage := gevm.NewStorage()

	warm, v := storage.Load(0)
	fmt.Println("Warm access?:", warm)
	fmt.Println("Storage value:", v)

	storage.Store(0, []byte("Hello Ethereum"))
	warm, v = storage.Load(0)
	fmt.Println("Warm access?:", warm)
	fmt.Println("Storage value:", v)

	warm, v = storage.Load(292929)
	fmt.Println("Warm access?:", warm) // false
	fmt.Println("Storage value:", v)   // [0]
}
