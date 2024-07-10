# gevm

An Ethereum Virtual Machine (EVM) implementation from scratch, written in Go.

> [!WARNING]
> This implementation is for educational purposes and not for production use.

This project is an isolated EVM implementation, meaning it has no state for accounts. However, memory, storage, and event logs are tracked, although everything resets after each EVM execution.

## Running the EVM

To run the EVM:

1. Navigate to the command directory:
   ```sh
   cd cmd/gevm
   ```
2. Execute the main program:

   ```sh
   go run main.go
   ```

   Some simple example bytecodes are provided.
   There is also an example from an actual compiled contract, which is commented out. Uncomment it to run it.

## Dynamic Gas Calculation

Dynamic gas calculation is supported (memory expansion cost and storage operations). Functions for this are located in `gevm/common.go`, and the `dgMap[Opcode]uint64` holds records of each opcode that has dynamic gas. The dynamic gas is calculated at runtime for any opcode that has dynamic gas during execution, and the `dgMap` is updated to store this gas cost.

## EVM Structure

The main EVM files are located in `/gevm`. Here are some key structures:

- `ExecutionRuntime` represents the execution runtime during EVM execution.

  ```go
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
  ```

- `ExecutionEnvironment` encapsulates the EVM execution data environment, including the stack, memory, storage, and transient storage.

  ```go
  type ExecutionEnvironment struct {
      Stack     *Stack
      Memory    *Memory
      Storage   *Storage
      Transient *TransientStorage
  }
  ```

- `TransactionContext` holds transaction-specific information during EVM execution.

  ```go
  type TransactionContext struct {
      Sender   common.Address
      Value    uint64
      Calldata []byte
  }
  ```

- `ChainConfig` stores network configuration parameters.

  ```go
  type ChainConfig struct {
      ChainID  uint64
      GasLimit uint64
  }
  ```

- `Block` represents a block.
  ```go
  type Block struct {
      Coinbase  common.Address
      GasPrice  uint64
      Number    uint64
      Timestamp time.Time
      BaseFee   uint64
  }
  ```

All these are initialized with the `NewEVM` function found in `gevm/evm.go`.

## Tests

To run unit tests:

```sh
go test -v
```

For a better understanding of the project, explore the files in `/gevm`.

## Supported Opcodes

The implementation supports 131 out of the 143 EVM opcodes.

## Unsupported Opcodes

The following opcodes are not supported:

- CREATE
- CALL
- CALLCODE
- DELEGATECALL
- CREATE2
- EXTCODEHASH
- DIFFICULTY
- STATICCALL
- SELFBALANCE
- SELFDESTRUCT
- EIP-4844 opcodes: BLOBHASH, BLOBBASEFEE

These opcodes typically require state management. All other opcodes are supported, including EIP-1153 transient storage opcodes `TLOAD` and `TSTORE`.

## Contributing

Contributions are welcome! Please feel free to create an issue or submit a pull request.
