package nes

type Opcode struct {
	code     uint8
	mnemonic string
	len      uint8
	cycles   uint8
	mode     AddressingMode
}

var CPU_OP_CODES = []Opcode{

	{0x29, "AND", 2, 2, IMMEDIATE},
	{0x25, "AND", 2, 3, ZEROPAGE},
	{0x35, "AND", 2, 4, ZEROPAGE_X},
	{0x2D, "AND", 3, 4, ABSOLUTE},
	{0x3D, "AND", 3, 4 /*+1 if page crossed*/, ABSOLUTE_X},
	{0x39, "AND", 3, 4 /*+1 if page crossed*/, ABSOLUTE_Y},
	{0x21, "AND", 2, 6, INDIRECT_X},
	{0x31, "AND", 2, 5 /*+1 if page crossed*/, INDIRECT_Y},

	{0x00, "BRK", 1, 7, NONEADDRESSING},
	{0xAA, "TAX", 1, 2, NONEADDRESSING},
	{0xE8, "INX", 1, 2, NONEADDRESSING},

	{0xA9, "LDA", 2, 2, IMMEDIATE},
	{0xA5, "LDA", 2, 3, ZEROPAGE},
	{0xB5, "LDA", 2, 4, ZEROPAGE_X},
	{0xAD, "LDA", 3, 4, ABSOLUTE},
	{0xBD, "LDA", 3, 4 /*+1 if page crossed*/, ABSOLUTE_X},
	{0xB9, "LDA", 3, 4 /*+1 if page crossed*/, ABSOLUTE_Y},
	{0xA1, "LDA", 2, 6, INDIRECT_X},
	{0xB1, "LDA", 2, 5 /*+1 if page crossed*/, INDIRECT_Y},

	{0x84, "STA", 2, 3, ZEROPAGE},
	{0x95, "STA", 2, 4, ZEROPAGE_X},
	{0x8d, "STA", 3, 4, ABSOLUTE},
	{0x9d, "STA", 3, 5, ABSOLUTE_X},
	{0x99, "STA", 3, 5, ABSOLUTE_Y},
	{0x81, "STA", 2, 6, INDIRECT_X},
	{0x91, "STA", 2, 6, INDIRECT_Y},
}

func NewOpcodeMap() map[uint8]Opcode {
	opcodes_map := make(map[uint8]Opcode)
	for _, cpuop := range CPU_OP_CODES {
		opcodes_map[cpuop.code] = cpuop
	}
	return opcodes_map
}

var OPCODES_MAP = map[uint8]Opcode{
	1: CPU_OP_CODES[0],
}
