package main

import "fmt"

type CPU6502 struct {
	memory      Memory
	registerACC uint8
	registerPC  uint16
	registerS   uint8
	registerX   uint8
	registerY   uint8
	flagD       uint8
	flagC       uint8
	flagI       uint8
	flagV       uint8
	flagZ       uint8
	flagN       uint8

	commands [256]instruction
}

func CreateCPU6502(memory Memory) *CPU6502 {

	cpu := &CPU6502{
		memory: memory,
	}

	cpu.commands[0x65] = cpu.buildInstruction(cpu.adcCommand, cpu.addressingZP, 3)
	cpu.commands[0x69] = cpu.buildInstruction(cpu.adcCommand, cpu.addressingIMM, 3)
	cpu.commands[0x75] = cpu.buildInstruction(cpu.adcCommand, cpu.addressingZPX, 3)

	cpu.commands[0xA5] = cpu.buildInstruction(cpu.ldaCommand, cpu.addressingZP, 2)
	cpu.commands[0xA9] = cpu.buildInstruction(cpu.ldaCommand, cpu.addressingIMM, 2)
	cpu.commands[0x85] = cpu.buildInstruction(cpu.staCommand, cpu.addressingZP, 2)

	return cpu
}

func (cpu *CPU6502) extractNextInstruction() uint8 {
	data := cpu.memory.read(cpu.registerPC)
	cpu.registerPC++
	return data
}

func (cpu *CPU6502) extractNextInstruction16() uint16 {
	data := cpu.memory.read16(cpu.registerPC)
	cpu.registerPC += 2
	return data
}

func (cpu *CPU6502) Reset() {
	cpu.registerPC = 0x0800
	cpu.registerS = 0xff
	cpu.flagD = 0
	cpu.flagC = 0
	cpu.flagI = 0
	cpu.flagV = 0
	cpu.flagZ = 0
	cpu.flagN = 0

}

func (cpu *CPU6502) DoIteration() {
	instruction := cpu.extractNextInstruction()
	fmt.Printf("%#x\n", instruction)
	cpu.commands[instruction]()
}

type instruction func()
type command func(address uint16)
type addressing func() (address uint16, extraCycles uint16)

func (cpu *CPU6502) addressingZP() (address uint16, extraCycles uint16) {
	return uint16(cpu.extractNextInstruction()), 0
}

func (cpu *CPU6502) addressingREL() (address uint16, extraCycles uint16) {
	var rel = int8(cpu.extractNextInstruction())

	return sum(cpu.registerPC, rel), 0
}

func (cpu *CPU6502) addressingIMP() (address uint16, extraCycles uint16) {
	return 0, 0
}

func (cpu *CPU6502) addressingABS() (address uint16, extraCycles uint16) {
	return cpu.extractNextInstruction16(), 0
}

func (cpu *CPU6502) addressingACC() (address uint16, extraCycles uint16) {
	return uint16(cpu.registerACC), 0
}

func (cpu *CPU6502) addressingIMM() (address uint16, extraCycles uint16) {
	address = cpu.registerPC
	cpu.registerPC += 1
	return address, 0
}

func (cpu *CPU6502) addressingZPX() (address uint16, extraCycles uint16) {
	//TODO: probably & 0xFF
	return uint16(cpu.extractNextInstruction() + cpu.registerX), 0
}

func (cpu *CPU6502) addressingZPY() (address uint16, extraCycles uint16) {
	//TODO: probably & 0xFF
	return uint16(cpu.extractNextInstruction() + cpu.registerY), 0
}

func (cpu *CPU6502) addressingABSX() (address uint16, extraCycles uint16) {
	//TODO: extra cycle
	return cpu.extractNextInstruction16() + uint16(cpu.registerX), 0
}

func (cpu *CPU6502) addressingABSY() (address uint16, extraCycles uint16) {
	//TODO: extra cycle
	return cpu.extractNextInstruction16() + uint16(cpu.registerY), 0
}

func (cpu *CPU6502) addressingPREIDXIND() (address uint16, extraCycles uint16) {
	//TODO: extra cycle
	return cpu.memory.read16(uint16(cpu.extractNextInstruction() + cpu.registerX)), 0
}

func (cpu *CPU6502) addressingPOSTIDXIND() (address uint16, extraCycles uint16) {
	//TODO: extra cycle
	return cpu.memory.read16(cpu.extractNextInstruction16() + uint16(cpu.registerY)), 0
}

func (cpu *CPU6502) addressingINDABS() (address uint16, extraCycles uint16) {
	return cpu.memory.read16(cpu.extractNextInstruction16() + uint16(cpu.registerY)), 0
}

func (cpu *CPU6502) setAcc(data uint8) {
	cpu.registerACC = data

	cpu.flagN = (data >> 7) & 1
	cpu.flagZ = data & 0xff
}

func (cpu *CPU6502) pullStack() uint8 {
	cpu.registerS++

	return cpu.memory.read(uint16(cpu.registerS) + 0x0100)
}

func (cpu *CPU6502) pushStack(data uint8) {
	cpu.memory.write(uint16(cpu.registerS)+0x0100, data)

	cpu.registerS--
}

func (cpu *CPU6502) pushStack16(data uint16) {
	cpu.pushStack(uint8(cpu.registerPC >> 8))
	cpu.pushStack(uint8(cpu.registerPC))
}

func (cpu *CPU6502) buildInstruction(command command, addressing addressing, cycles uint16) instruction {
	return func() {
		argument, _ := addressing()
		command(argument)
	}
}

func sum(a uint16, b int8) uint16 {
	return uint16(int32(a) + int32(b))
}
