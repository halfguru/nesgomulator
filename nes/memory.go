package nes

type Memory []uint8

func NewMemory() Memory {
	return make([]uint8, 0x10000)
}

func (mem Memory) Read(addr uint16) uint8 {
	return mem[addr]
}

func (mem Memory) Write(addr uint16, data uint8) {
	mem[addr] = data
}

func (mem Memory) Read_u16(pos uint16) uint16 {
	lo := uint16(mem.Read(pos))
	hi := uint16(mem.Read(pos + 1))
	return hi<<8 | lo
}

func (mem Memory) Write_u16(pos uint16, data uint16) {
	hi := uint8(data >> 8)
	lo := uint8(data & 0xff)
	mem.Write(pos, lo)
	mem.Write(pos+1, hi)
}

func (mem Memory) Load(program []uint8) {
	copy(mem[0x8000:], program)
	mem.Write_u16(0xFFFC, 0x8000)
}
