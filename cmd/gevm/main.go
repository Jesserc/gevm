package main

import (
	"fmt"

	"github.com/Jesserc/gevm/gevm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
)

func main() {
	evm := gevm.NewEVM(common.HexToAddress("0x"), 21000, uint256.NewInt(2e5), []byte{}, []byte{})

	evm.Stack.Push(uint256.NewInt(0).SetBytes(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000FFFFFFFF")))
	evm.Stack.Push(uint256.NewInt(0).SetBytes(hexutil.MustDecode("0xFFFFFFFF00000000000000000000000000000000000000000000000000000000")))
	evm.Stack.Push(uint256.NewInt(32))

	fmt.Printf("EVM: \n%v\n", evm.Stack.ToString())
}
