package gevm

type Opcode byte

const (
	STOP Opcode = 0x0
)

// Math
const (
	ADD        Opcode = 0x1
	MUL        Opcode = 0x2
	SUB        Opcode = 0x3
	DIV        Opcode = 0x4
	SDIV       Opcode = 0x5
	MOD        Opcode = 0x6
	SMOD       Opcode = 0x7
	ADDMOD     Opcode = 0x8
	MULMOD     Opcode = 0x9
	EXP        Opcode = 0xA
	SIGNEXTEND Opcode = 0xB
)

// Comparisons
const (
	LT     Opcode = 0x10
	GT     Opcode = 0x11
	SLT    Opcode = 0x12
	SGT    Opcode = 0x13
	EQ     Opcode = 0x14
	ISZERO Opcode = 0x15
)

// Logic
const (
	AND Opcode = 0x16
	OR  Opcode = 0x17
	XOR Opcode = 0x18
	NOT Opcode = 0x19
)

// Bit Operations
const (
	BYTE Opcode = 0x1A
	SHL  Opcode = 0x1B
	SHR  Opcode = 0x1C
	SAR  Opcode = 0x1D
)

// Hash
const (
	SHA3 Opcode = 0x20
)

// Ethereum State
const (
	ADDRESS        Opcode = 0x30
	BALANCE        Opcode = 0x31
	ORIGIN         Opcode = 0x32
	CALLER         Opcode = 0x33
	CALLVALUE      Opcode = 0x34
	CALLDATALOAD   Opcode = 0x35
	CALLDATASIZE   Opcode = 0x36
	CALLDATACOPY   Opcode = 0x37
	CODESIZE       Opcode = 0x38
	CODECOPY       Opcode = 0x39
	GASPRICE       Opcode = 0x3A
	EXTCODESIZE    Opcode = 0x3B
	EXTCODECOPY    Opcode = 0x3C
	RETURNDATASIZE Opcode = 0x3D
	RETURNDATACOPY Opcode = 0x3E
	EXTCODEHASH    Opcode = 0x3F
	BLOCKHASH      Opcode = 0x40
	COINBASE       Opcode = 0x41
	TIMESTAMP      Opcode = 0x42
	NUMBER         Opcode = 0x43
	DIFFICULTY     Opcode = 0x44
	GASLIMIT       Opcode = 0x45
	CHAINID        Opcode = 0x46
	SELFBALANCE    Opcode = 0x47
	BASEFEE        Opcode = 0x48
)

// Stack Pop
const (
	POP Opcode = 0x50
)

// Memory
const (
	MLOAD   Opcode = 0x51
	MSTORE  Opcode = 0x52
	MSTORE8 Opcode = 0x53
)

// Storage
const (
	SLOAD  Opcode = 0x54
	SSTORE Opcode = 0x55
)

// Jump
const (
	JUMP     Opcode = 0x56
	JUMPI    Opcode = 0x57
	PC       Opcode = 0x58
	JUMPDEST Opcode = 0x5B
)

// Transient Storage
const (
	TLOAD  Opcode = 0x5c
	TSTORE Opcode = 0x5d
)

// Push
const (
	PUSH1  Opcode = 0x60
	PUSH2  Opcode = 0x61
	PUSH3  Opcode = 0x62
	PUSH4  Opcode = 0x63
	PUSH5  Opcode = 0x64
	PUSH6  Opcode = 0x65
	PUSH7  Opcode = 0x66
	PUSH8  Opcode = 0x67
	PUSH9  Opcode = 0x68
	PUSH10 Opcode = 0x69
	PUSH11 Opcode = 0x6A
	PUSH12 Opcode = 0x6B
	PUSH13 Opcode = 0x6C
	PUSH14 Opcode = 0x6D
	PUSH15 Opcode = 0x6E
	PUSH16 Opcode = 0x6F
	PUSH17 Opcode = 0x70
	PUSH18 Opcode = 0x71
	PUSH19 Opcode = 0x72
	PUSH20 Opcode = 0x73
	PUSH21 Opcode = 0x74
	PUSH22 Opcode = 0x75
	PUSH23 Opcode = 0x76
	PUSH24 Opcode = 0x77
	PUSH25 Opcode = 0x78
	PUSH26 Opcode = 0x79
	PUSH27 Opcode = 0x7A
	PUSH28 Opcode = 0x7B
	PUSH29 Opcode = 0x7C
	PUSH30 Opcode = 0x7D
	PUSH31 Opcode = 0x7E
	PUSH32 Opcode = 0x7F
)

// Dup
const (
	DUP1  Opcode = 0x80
	DUP2  Opcode = 0x81
	DUP3  Opcode = 0x82
	DUP4  Opcode = 0x83
	DUP5  Opcode = 0x84
	DUP6  Opcode = 0x85
	DUP7  Opcode = 0x86
	DUP8  Opcode = 0x87
	DUP9  Opcode = 0x88
	DUP10 Opcode = 0x89
	DUP11 Opcode = 0x8A
	DUP12 Opcode = 0x8B
	DUP13 Opcode = 0x8C
	DUP14 Opcode = 0x8D
	DUP15 Opcode = 0x8E
	DUP16 Opcode = 0x8F
)

// Swap
const (
	SWAP1  Opcode = 0x90
	SWAP2  Opcode = 0x91
	SWAP3  Opcode = 0x92
	SWAP4  Opcode = 0x93
	SWAP5  Opcode = 0x94
	SWAP6  Opcode = 0x95
	SWAP7  Opcode = 0x96
	SWAP8  Opcode = 0x97
	SWAP9  Opcode = 0x98
	SWAP10 Opcode = 0x99
	SWAP11 Opcode = 0x9A
	SWAP12 Opcode = 0x9B
	SWAP13 Opcode = 0x9C
	SWAP14 Opcode = 0x9D
	SWAP15 Opcode = 0x9E
	SWAP16 Opcode = 0x9F
)

// Log
const (
	LOG0 Opcode = 0xA0
	LOG1 Opcode = 0xA1
	LOG2 Opcode = 0xA2
	LOG3 Opcode = 0xA3
	LOG4 Opcode = 0xA4
)

// Contract
const (
	CREATE       Opcode = 0xF0
	CALL         Opcode = 0xF1
	CALLCODE     Opcode = 0xF2 // legacy, NOT supported by us, fixed by DELEGATECALL
	RETURN       Opcode = 0xF3
	DELEGATECALL Opcode = 0xF4
	CREATE2      Opcode = 0xF5
	STATICCALL   Opcode = 0xFA
	REVERT       Opcode = 0xFD
	INVALID      Opcode = 0xFE
	SELFDESTRUCT Opcode = 0xFF
)
