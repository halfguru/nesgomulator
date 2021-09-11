package nes

import (
	"errors"

	"github.com/golang/glog"
)

type Cpu struct {
	A               uint8
	X               uint8
	Y               uint8
	status          uint8
	program_counter uint16
	memory          Memory
	opcodes         [0xFF]func()
}

type AddressingMode string

const (
	IMMEDIATE      AddressingMode = "Immediate"
	ZEROPAGE       AddressingMode = "ZeroPage"
	ZEROPAGE_X     AddressingMode = "ZeroPage_X"
	ZEROPAGE_Y     AddressingMode = "ZeroPage_Y"
	ABSOLUTE       AddressingMode = "Absolute"
	ABSOLUTE_X     AddressingMode = "Absolute_X"
	ABSOLUTE_Y     AddressingMode = "Absolute_Y"
	INDIRECT_X     AddressingMode = "Indirect_X"
	INDIRECT_Y     AddressingMode = "Indirect_Y"
	NONEADDRESSING AddressingMode = "NoneAddressing"
)

func NewCpu() Cpu {
	return Cpu{
		A:               0,
		X:               0,
		status:          0,
		program_counter: 0,
		memory:          NewMemory(),
	}
}

func (cpu *Cpu) Reset() {
	cpu.A = 0
	cpu.X = 0
	cpu.status = 0
	cpu.program_counter = cpu.memory.Read_u16(0xFFFC)
}

func (cpu *Cpu) GetOperandAddress(mode AddressingMode) (uint16, error) {
	switch mode {
	case IMMEDIATE:
		return cpu.program_counter, nil

	case ZEROPAGE:
		return uint16(cpu.memory.Read(cpu.program_counter)), nil

	case ABSOLUTE:
		return cpu.memory.Read_u16(cpu.program_counter), nil

	case ZEROPAGE_X:
		pos := cpu.memory.Read(cpu.program_counter)
		addr := pos + cpu.X
		return uint16(addr), nil

	case ZEROPAGE_Y:
		pos := cpu.memory.Read(cpu.program_counter)
		addr := pos + cpu.Y
		return uint16(addr), nil

	case ABSOLUTE_X:
		base := cpu.memory.Read_u16(cpu.program_counter)
		addr := base + uint16(cpu.X)
		return addr, nil

	case ABSOLUTE_Y:
		base := cpu.memory.Read_u16(cpu.program_counter)
		addr := base + uint16(cpu.Y)
		return addr, nil

	case INDIRECT_X:
		base := cpu.memory.Read(cpu.program_counter)
		ptr := base + cpu.X
		lo := cpu.memory.Read(uint16(ptr))
		hi := cpu.memory.Read(uint16(ptr + 1))
		return uint16(hi)<<8 | uint16(lo), nil

	case INDIRECT_Y:
		base := cpu.memory.Read(cpu.program_counter)
		ptr := base + cpu.Y
		lo := cpu.memory.Read(uint16(ptr))
		hi := cpu.memory.Read(uint16(ptr + 1))
		return uint16(hi)<<8 | uint16(lo), nil

	default:
		return 0, errors.New("Operand Address not found")
	}
}

// If result bit 7 is set, set the status Negative flag
func (cpu *Cpu) setNegative(result uint8) {
	if result&byte(0b10000000) != 0 {
		cpu.status = cpu.status | byte(0b10000000)
	} else {
		cpu.status = cpu.status & byte(0b01111111)
	}
}

// If result is not set, set the status Zero flag
func (cpu *Cpu) setZero(result uint8) {
	if result == 0 {
		cpu.status = cpu.status | byte(0b00000010)
	} else {
		cpu.status = cpu.status & byte(0b11111101)
	}
}

// A logical AND is performed, bit by bit, on the accumulator contents using the contents of a byte of memory.
func (cpu *Cpu) AND(mode AddressingMode) {
	glog.Info("AND")
	addr, _ := cpu.GetOperandAddress(mode)
	value := cpu.memory.Read(addr)
	cpu.setZero(value & cpu.A)
	cpu.setNegative(value & cpu.A)
}

// This operation shifts all the bits of the accumulator or memory contents one bit left.
// Bit 0 is set to 0 and bit 7 is placed in the carry flag. The effect of this operation
// is to multiply the memory contents by 2 (ignoring 2's complement considerations),
// setting the carry if the result will not fit in 8 bits.
func (cpu *Cpu) ASL(mode AddressingMode) {
	data := cpu.A
	if data>>7 == 1 {

	}
}

// Load byte of mmroy into accumulator seting the zero and negative flags
func (cpu *Cpu) LDA(mode AddressingMode) {
	glog.Info("LDA")
	addr, _ := cpu.GetOperandAddress(mode)
	value := cpu.memory.Read(addr)
	cpu.A = value
	cpu.setZero(cpu.A)
	cpu.setNegative(cpu.A)
}

func (cpu *Cpu) STA(mode AddressingMode) {
	glog.Info("STA")
	addr, _ := cpu.GetOperandAddress(mode)
	cpu.memory.Write(addr, cpu.A)
}

// Copies the current ctents of the accumulator into he X register
// and sets the zo and egative flags as appropriate.
func (cpu *Cpu) TAX() {
	glog.Info("TA")
	cpu.X = cpu.A
	cpu.setZero(cpu.X)
	cpu.setNegative(cpu.X)
}

// Adds one to thX regiter setting the zero and negative flags as appropriate.
func (cpu *Cpu) INX() {
	glog.Info("INX")
	cpu.X += 1
	cpu.setZero(cpu.X)
	cpu.setNegative(cpu.X)
}

func (cpu *Cpu) Run(program []uint8) {
	cpu.memory.Load(program)
	cpu.Reset()
	opcode_map := NewOpcodeMap()

	for {
		code := cpu.memory.Read(uint16(cpu.program_counter))
		cpu.program_counter += 1
		opcode := opcode_map[code]

		switch code {

		case 0x29, 0x25, 0x35, 0x2D, 0x3D, 0x39, 0x21, 0x31:
			cpu.AND(opcode.mode)

		case 0x0A, 0x06, 0x16, 0x0E, 0x1E:
			cpu.ASL(opcode.mode)

		case 0xA9, 0xA5, 0xB5, 0xAD, 0xBD, 0xB9, 0xA1, 0xB1:
			cpu.LDA(opcode.mode)

		case 0x85, 0x95, 0x8D, 0x9D, 0x99, 0x81, 0x91:
			cpu.STA(opcode.mode)

		case 0xAA:
			cpu.TAX()

		case 0xE8:
			cpu.INX()

		case 0x00:
			return

		default:
			return
		}

		cpu.program_counter += uint16((opcode.len - 1))
	}

}
