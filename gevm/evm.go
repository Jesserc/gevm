package gevm

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// dynamicGasMap maps opcodes to their dynamic gas costs.
type dynamicGasMap map[Opcode]uint64

// DynamicGas returns the gas cost for a given opcode.
func (d dynamicGasMap) DynamicGas(op Opcode) (bool, uint64) {
	gas, hasDyGas := d[op]
	return hasDyGas, gas
}

var (
	dgMap = make(dynamicGasMap)
)

// ExecutionRuntime represents the execution runtime during EVM execution.
type ExecutionRuntime struct {
	PC         uint64
	Code       []byte
	Gas        uint64
	Refund     uint64
	StopFlag   bool
	RevertFlag bool
	ReturnData []byte
	LogRecord  *LogRecord
	Block      *Block
}

// ExecutionEnvironment encapsulates the EVM execution data environment.
type ExecutionEnvironment struct {
	Stack     *Stack
	Memory    *Memory
	Storage   *Storage
	Transient *TransientStorage
}

// TransactionContext holds transaction-specific information during EVM execution.
type TransactionContext struct {
	Sender   common.Address
	Value    uint64
	Calldata []byte
}

// ChainConfig stores network configuration parameters.
type ChainConfig struct {
	ChainID  uint64
	GasLimit uint64
}

// Block represents a block.
type Block struct {
	Coinbase  common.Address
	GasPrice  uint64
	Number    uint64
	Timestamp time.Time
	BaseFee   uint64
}

// NewBlock creates a new block instance.
func NewBlock(coinbase common.Address, gasPrice, number, difficulty, baseFee uint64, timeStamp time.Time) *Block {
	return &Block{
		Coinbase:  coinbase,
		GasPrice:  gasPrice,
		Number:    number,
		Timestamp: timeStamp,
		BaseFee:   baseFee,
	}
}

// EVM represents an Ethereum Virtual Machine instance.
type EVM struct {
	ExecutionRuntime
	ExecutionEnvironment
	TransactionContext
	ChainConfig
}

func (evm *EVM) deductGas(gas uint64) {
	if evm.Gas < gas || evm.Gas <= 0 {
		panic(fmt.Errorf("out of gas: tried to consume %d gas, but only %d gas remaining", gas, evm.Gas))
	}
	evm.Gas -= gas // deduct gas
}

// continueExecution checks if the EVM should continue execution.
func (evm *EVM) continueExecution() bool {
	return int(evm.PC) <= len(evm.Code)-1 && // Check if PC is within code bounds
		!evm.StopFlag && // Check if STOP instruction was encountered
		!evm.RevertFlag // Check if REVERT instruction was encountered
	//	&& evm.Gas > 0 // Check if there's enough gas to continue
}

func (evm *EVM) Run() {
	fmt.Println("#### Trace ####")

	// Initialize the jump table containing all opcode implementations
	jumpTable := NewJumpTable()
	var totalGasUsed uint64

	// Main execution loop
	for evm.continueExecution() {
		// Get the current program counter and opcode
		currentPC := evm.PC
		opcode := evm.Code[currentPC]
		op := Opcode(opcode)

		// Execute the opcode if it exists in the jump table
		if opFunc, exists := jumpTable[op]; exists {
			opFunc(evm)
		} else {
			// Handle unknown opcodes
			fmt.Printf("Unknown opcode: %#x\n", opcode)
			return
		}

		// Calculate gas cost for the operation
		var gCost uint64
		if has, cost := dgMap.DynamicGas(op); has {
			// Use dynamic gas cost if available
			gCost = cost
		} else {
			// Use static gas cost defined for the opcode
			gCost = op.Gas()
		}
		totalGasUsed += gCost // accumulate total gas used

		// Log the current EVM state after executing the opcode
		logEVMState(evm, op, gCost, currentPC)
	}

	// Total gas consumed
	totalGasUsed -= evm.Refund // minus refund, if any

	// Log emitted logs and some stats
	LogEVMLogs(totalGasUsed, evm)
}

func (evm *EVM) addRefund(refund uint64) {
	evm.Refund += refund
}

func (evm *EVM) subRefund(gas uint64) {
	if gas > evm.Refund {
		panic(fmt.Sprintf("Refund counter below zero (gas: %d > refund: %d)", gas, evm.Refund))
	}
	evm.Refund -= gas
}

func NewEVM(sender common.Address, gas, value, chainID, gasLimit uint64, code, calldata []byte, blockInfo *Block) *EVM {
	return &EVM{
		ExecutionRuntime: ExecutionRuntime{
			PC:         0,
			Code:       []byte{},
			Gas:        gas,
			Refund:     0,
			StopFlag:   false,
			RevertFlag: false,
			ReturnData: []byte{},
			LogRecord:  NewLogRecord(),
			Block:      blockInfo,
		},
		ExecutionEnvironment: ExecutionEnvironment{
			Stack:     NewStack(),
			Memory:    NewMemory(),
			Storage:   NewStorage(),
			Transient: NewTransientStorage(),
		},
		TransactionContext: TransactionContext{
			Sender:   common.Address{},
			Value:    0,
			Calldata: []byte{},
		},
		ChainConfig: ChainConfig{
			ChainID:  1,
			GasLimit: gas,
		},
	}
}

func logEVMState(evm *EVM, op Opcode, gasCost uint64, currentPC uint64) {
	fmt.Println("Opcode:", op)
	fmt.Println("Stack:", evm.Stack.ToString())
	fmt.Println("Gas Cost:", gasCost)
	fmt.Println("Memory:", hexutil.Encode(evm.Memory.data))
	fmt.Println("Storage:", evm.Storage.data)
	fmt.Println("Return Data:", hexutil.Encode(evm.ReturnData))
	fmt.Println("PC:", currentPC)
	fmt.Println()
	// fmt.Println("Value:")
}

func LogEVMLogs(totalGasUsed uint64, evm *EVM) {
	fmt.Println("#### LOGS ####")
	fmt.Println("Total gas used:", totalGasUsed)
	fmt.Println("Total memory allocations:", toWordSize(uint64(len(evm.Memory.data))))
	fmt.Println("Allocated bytes in memory:", len(evm.Memory.data))
	fmt.Println("Total storage allocations:", len(evm.Storage.data))
	fmt.Println("Total storage gas refund:", evm.Refund)
	fmt.Println("Logs:\n", evm.LogRecord)
	fmt.Println("Chain ID:", evm.ChainID)
	fmt.Println("Gas Limit:", evm.GasLimit)
	fmt.Println("Coinbase:", evm.Block.Coinbase)
}
