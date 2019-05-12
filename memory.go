package main

type RAMMemory []uint8

func (ram RAMMemory) read(address uint16) uint8 {
	return ram[address]
}

func (ram RAMMemory) read16(address uint16) uint16 {
	//TODO:
	return uint16(ram.read(address)) + (uint16(ram.read(address+1)) << 8)
}

func (ram RAMMemory) write(address uint16, value uint8) {
	ram[address] = value
}
