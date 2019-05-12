package main

type CPU interface {
	DoIteration()
}

type Memory interface {
	read(address uint16) uint8
	read16(address uint16) uint16
	write(address uint16, value uint8)
}

type PPU interface {
}
