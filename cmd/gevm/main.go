package main

import (
	"fmt"

	"github.com/Jesserc/gevm/gevm"
)

func main() {
	stack := gevm.NewStack(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack)
}
