package gevm

import "fmt"

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
	KECCAK256 Opcode = 0x20
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
	GAS            Opcode = 0x5A
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
	MSIZE   Opcode = 0x59
	MCOPY   Opcode = 0x5e
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
	PUSH0  Opcode = 0x5F
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

func (op Opcode) String() string {
	switch op {
	case STOP:
		return "STOP"
	case ADD:
		return "ADD"
	case MUL:
		return "MUL"
	case SUB:
		return "SUB"
	case DIV:
		return "DIV"
	case SDIV:
		return "SDIV"
	case MOD:
		return "MOD"
	case SMOD:
		return "SMOD"
	case ADDMOD:
		return "ADDMOD"
	case MULMOD:
		return "MULMOD"
	case EXP:
		return "EXP"
	case SIGNEXTEND:
		return "SIGNEXTEND"
	case LT:
		return "LT"
	case GT:
		return "GT"
	case SLT:
		return "SLT"
	case SGT:
		return "SGT"
	case EQ:
		return "EQ"
	case ISZERO:
		return "ISZERO"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case XOR:
		return "XOR"
	case NOT:
		return "NOT"
	case BYTE:
		return "BYTE"
	case SHL:
		return "SHL"
	case SHR:
		return "SHR"
	case SAR:
		return "SAR"
	case KECCAK256:
		return "KECCAK256"
	case ADDRESS:
		return "ADDRESS"
	case BALANCE:
		return "BALANCE"
	case ORIGIN:
		return "ORIGIN"
	case CALLER:
		return "CALLER"
	case CALLVALUE:
		return "CALLVALUE"
	case CALLDATALOAD:
		return "CALLDATALOAD"
	case CALLDATASIZE:
		return "CALLDATASIZE"
	case CALLDATACOPY:
		return "CALLDATACOPY"
	case CODESIZE:
		return "CODESIZE"
	case CODECOPY:
		return "CODECOPY"
	case GASPRICE:
		return "GASPRICE"
	case EXTCODESIZE:
		return "EXTCODESIZE"
	case EXTCODECOPY:
		return "EXTCODECOPY"
	case RETURNDATASIZE:
		return "RETURNDATASIZE"
	case RETURNDATACOPY:
		return "RETURNDATACOPY"
	case EXTCODEHASH:
		return "EXTCODEHASH"
	case BLOCKHASH:
		return "BLOCKHASH"
	case COINBASE:
		return "COINBASE"
	case TIMESTAMP:
		return "TIMESTAMP"
	case NUMBER:
		return "NUMBER"
	case DIFFICULTY:
		return "DIFFICULTY"
	case GASLIMIT:
		return "GASLIMIT"
	case CHAINID:
		return "CHAINID"
	case SELFBALANCE:
		return "SELFBALANCE"
	// case BASEFEE:
	// 	return "BASEFEE"
	case POP:
		return "POP"
	case MLOAD:
		return "MLOAD"
	case MSTORE:
		return "MSTORE"
	case MSTORE8:
		return "MSTORE8"
	case SLOAD:
		return "SLOAD"
	case SSTORE:
		return "SSTORE"
	case JUMP:
		return "JUMP"
	case JUMPI:
		return "JUMPI"
	case PC:
		return "PC"
	case MSIZE:
		return "MSIZE"
	case MCOPY:
		return "MCOPY"
	case GAS:
		return "GAS"
	case JUMPDEST:
		return "JUMPDEST"
	case PUSH1:
		return "PUSH1"
	case PUSH2:
		return "PUSH2"
	case PUSH3:
		return "PUSH3"
	case PUSH4:
		return "PUSH4"
	case PUSH5:
		return "PUSH5"
	case PUSH6:
		return "PUSH6"
	case PUSH7:
		return "PUSH7"
	case PUSH8:
		return "PUSH8"
	case PUSH9:
		return "PUSH9"
	case PUSH10:
		return "PUSH10"
	case PUSH11:
		return "PUSH11"
	case PUSH12:
		return "PUSH12"
	case PUSH13:
		return "PUSH13"
	case PUSH14:
		return "PUSH14"
	case PUSH15:
		return "PUSH15"
	case PUSH16:
		return "PUSH16"
	case PUSH17:
		return "PUSH17"
	case PUSH18:
		return "PUSH18"
	case PUSH19:
		return "PUSH19"
	case PUSH20:
		return "PUSH20"
	case PUSH21:
		return "PUSH21"
	case PUSH22:
		return "PUSH22"
	case PUSH23:
		return "PUSH23"
	case PUSH24:
		return "PUSH24"
	case PUSH25:
		return "PUSH25"
	case PUSH26:
		return "PUSH26"
	case PUSH27:
		return "PUSH27"
	case PUSH28:
		return "PUSH28"
	case PUSH29:
		return "PUSH29"
	case PUSH30:
		return "PUSH30"
	case PUSH31:
		return "PUSH31"
	case PUSH32:
		return "PUSH32"
	case DUP1:
		return "DUP1"
	case DUP2:
		return "DUP2"
	case DUP3:
		return "DUP3"
	case DUP4:
		return "DUP4"
	case DUP5:
		return "DUP5"
	case DUP6:
		return "DUP6"
	case DUP7:
		return "DUP7"
	case DUP8:
		return "DUP8"
	case DUP9:
		return "DUP9"
	case DUP10:
		return "DUP10"
	case DUP11:
		return "DUP11"
	case DUP12:
		return "DUP12"
	case DUP13:
		return "DUP13"
	case DUP14:
		return "DUP14"
	case DUP15:
		return "DUP15"
	case DUP16:
		return "DUP16"
	case SWAP1:
		return "SWAP1"
	case SWAP2:
		return "SWAP2"
	case SWAP3:
		return "SWAP3"
	case SWAP4:
		return "SWAP4"
	case SWAP5:
		return "SWAP5"
	case SWAP6:
		return "SWAP6"
	case SWAP7:
		return "SWAP7"
	case SWAP8:
		return "SWAP8"
	case SWAP9:
		return "SWAP9"
	case SWAP10:
		return "SWAP10"
	case SWAP11:
		return "SWAP11"
	case SWAP12:
		return "SWAP12"
	case SWAP13:
		return "SWAP13"
	case SWAP14:
		return "SWAP14"
	case SWAP15:
		return "SWAP15"
	case SWAP16:
		return "SWAP16"
	case LOG0:
		return "LOG0"
	case LOG1:
		return "LOG1"
	case LOG2:
		return "LOG2"
	case LOG3:
		return "LOG3"
	case LOG4:
		return "LOG4"
	case CREATE:
		return "CREATE"
	case CALL:
		return "CALL"
	case CALLCODE:
		return "CALLCODE"
	case RETURN:
		return "RETURN"
	case DELEGATECALL:
		return "DELEGATECALL"
	case CREATE2:
		return "CREATE2"
	case STATICCALL:
		return "STATICCALL"
	case REVERT:
		return "REVERT"
	case INVALID:
		return "INVALID"
	case SELFDESTRUCT:
		return "SELFDESTRUCT"
	case PUSH0:
		return "PUSH0"
	case BASEFEE:
		return "BASEFEE"
	default:
		return fmt.Sprintf("UNKNOWN_OPCODE(0x%x)", byte(op))
	}
}

func (op Opcode) Gas() uint64 {
	switch op {
	case STOP:
		return 0
	case ADD, SUB, LT, GT, SLT, SGT, EQ, ISZERO, AND, OR, XOR, NOT, BYTE, SHL, SHR, SAR:
	case MUL, DIV, SDIV, MOD, SMOD, SIGNEXTEND:
		return 5
	case ADDMOD, MULMOD:
		return 8
	case EXP:
		return 10
	case KECCAK256:
		return 30
	case ADDRESS, ORIGIN, CALLER, CALLVALUE, CALLDATALOAD, CALLDATASIZE, CODESIZE, GASPRICE, RETURNDATASIZE, EXTCODESIZE, BLOCKHASH, COINBASE, TIMESTAMP, NUMBER, DIFFICULTY, GASLIMIT, CHAINID, SELFBALANCE, BASEFEE:
		return 2
	case BALANCE:
		return 2600
	case CALLDATACOPY, CODECOPY, RETURNDATACOPY, EXTCODECOPY:
		return 3
	case EXTCODEHASH:
		return 400
	case POP:
		return 2
	case MLOAD, MSTORE:
		return 3
	case MSTORE8:
		return 3
	case SLOAD:
		return 800
	case SSTORE:
		return 21000 // simplified, actual cost depends on context
	case JUMP:
		return 8
	case JUMPI:
		return 10
	case PC, MSIZE, GAS:
		return 2
	case JUMPDEST:
		return 1
	case PUSH1, PUSH2, PUSH3, PUSH4, PUSH5, PUSH6, PUSH7, PUSH8, PUSH9, PUSH10, PUSH11, PUSH12, PUSH13, PUSH14, PUSH15, PUSH16, PUSH17, PUSH18, PUSH19, PUSH20, PUSH21, PUSH22, PUSH23, PUSH24, PUSH25, PUSH26, PUSH27, PUSH28, PUSH29, PUSH30, PUSH31, PUSH32:
		return 3
	case DUP1, DUP2, DUP3, DUP4, DUP5, DUP6, DUP7, DUP8, DUP9, DUP10, DUP11, DUP12, DUP13, DUP14, DUP15, DUP16:
		return 3
	case SWAP1, SWAP2, SWAP3, SWAP4, SWAP5, SWAP6, SWAP7, SWAP8, SWAP9, SWAP10, SWAP11, SWAP12, SWAP13, SWAP14, SWAP15, SWAP16:
		return 3
	case LOG0:
		return 375
	case LOG1:
		return 750
	case LOG2:
		return 1125
	case LOG3:
		return 1500
	case LOG4:
		return 1875
	case CREATE:
		return 32000
	case CALL:
		return 700 // simplified, actual cost depends on context
	case CALLCODE:
		return 700 // simplified, actual cost depends on context
	case RETURN:
		return 0
	case DELEGATECALL:
		return 700 // simplified, actual cost depends on context
	case CREATE2:
		return 32000
	case STATICCALL:
		return 700 // simplified, actual cost depends on context
	case REVERT:
		return 0
	case INVALID:
		return 0
	case SELFDESTRUCT:
		return 5000 // simplified, actual cost depends on context
	case 0x5f:
		return 2
	default:
		return 0
	}
	return 0
}
