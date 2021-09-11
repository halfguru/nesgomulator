package nes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_5_ops_working_together(t *testing.T) {
	program := []uint8{0xA9, 0xC0, 0xAA, 0xE8, 0x00}
	cpu := NewCpu()
	cpu.Run(program)
	assert.Equal(t, cpu.X, uint8(0xC1), "X flag should be set at 0xC1")
}

func Test_lda_zeropage(t *testing.T) {
	program := []uint8{0xa5, 0x10, 0x00}
	cpu := NewCpu()
	cpu.memory.Write(0x10, 0x55)
	cpu.Run(program)
	assert.Equal(t, cpu.A, uint8(0x55), "A flag should be set at 0x55")
}
