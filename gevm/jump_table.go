package gevm

type JumpTable map[Opcode]func(*EVM)

func NewJumpTable() JumpTable {
	jumpTable := map[Opcode]func(*EVM){
		STOP:           stop,
		ADD:            add,
		MUL:            mul,
		SUB:            sub,
		DIV:            div,
		SDIV:           sdiv,
		MOD:            mod,
		SMOD:           smod,
		ADDMOD:         addmod,
		MULMOD:         mulmod,
		EXP:            exp,
		SIGNEXTEND:     signextend,
		LT:             lt,
		GT:             gt,
		SLT:            slt,
		SGT:            sgt,
		EQ:             eq,
		ISZERO:         iszero,
		AND:            and,
		OR:             or,
		XOR:            xor,
		NOT:            not,
		BYTE:           _byte,
		SHL:            shl,
		SHR:            shr,
		SAR:            sar,
		KECCAK256:      keccak256,
		ADDRESS:        address,
		BALANCE:        balance,
		ORIGIN:         origin,
		CALLER:         caller,
		CALLVALUE:      callvalue,
		CALLDATALOAD:   calldataload,
		CALLDATASIZE:   calldatasize,
		CALLDATACOPY:   calldatacopy,
		CODESIZE:       codesize,
		CODECOPY:       codecopy,
		GASPRICE:       gasprice,
		EXTCODESIZE:    extcodesize,
		EXTCODECOPY:    extcodecopy,
		RETURNDATASIZE: returndatasize,
		RETURNDATACOPY: returndatacopy,
		BLOCKHASH:      blockhash,
		COINBASE:       coinbase,
		TIMESTAMP:      timestamp,
		NUMBER:         number,
		GASLIMIT:       gaslimit,
		CHAINID:        chainid,
		BASEFEE:        basefee,
		POP:            pop,
		PUSH0:          push0,
		MLOAD:          mload,
		MSTORE:         mstore,
		MSTORE8:        mstore8,
		MSIZE:          msize,
		MCOPY:          mcopy,
		SLOAD:          sload,
		SSTORE:         sstore,
		JUMP:           jump,
		JUMPI:          jumpi,
		PC:             pc,
		GAS:            gas,
		JUMPDEST:       jumpdest,
		RETURN:         _return,
		REVERT:         revert,
		INVALID:        invalid,
		LOG0:           log0,
		LOG1:           log1,
		LOG2:           log2,
		LOG3:           log3,
		LOG4:           log4,
		// SELFBALANCE:    selfbalance,
		// SELFDESTRUCT: selfdestruct,
	}

	// Add PUSH1 to PUSH32 opcodes
	for i := 0; i <= 31; i++ {
		i := uint8(i)
		opcode := PUSH1 + Opcode(i)
		jumpTable[opcode] = generatePushNFunc(i + 1)
	}

	// Add DUP1 to DUP16 opcodes
	for i := 0; i <= 15; i++ {
		i := uint8(i)
		opcode := DUP1 + Opcode(i)
		jumpTable[opcode] = generateDupNFunc(i + 1)
	}

	// Add SWAP1 to SWAP16 opcodes
	for i := 0; i <= 15; i++ {
		i := uint8(i)
		opcode := SWAP1 + Opcode(i)
		jumpTable[opcode] = generateSwapNFunc(i + 1)
	}

	return jumpTable
}
