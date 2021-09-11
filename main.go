package main

import (
	"flag"

	nes "github.com/halfguru/nesgomulator/nes"
)

func main() {
	flag.Parse()
	program := []uint8{0xA9, 0xC0, 0xAA, 0xE8, 0x00}
	cpu := nes.NewCpu()
	cpu.Run(program)
}
