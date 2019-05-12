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

func (cpu *CPU6502) adcCommand(address uint16) {
	//TODO:
	cpu.registerACC = cpu.registerACC + cpu.memory.read(address) + cpu.flagC
}

func (cpu *CPU6502) andCommand(address uint16) {
	cpu.registerACC = cpu.registerACC & cpu.memory.read(address)
}

func (cpu *CPU6502) aslCommand(address uint16) {
	//TODO:
	cpu.memory.write(address, cpu.memory.read(address)<<2)
}

func (cpu *CPU6502) bccCommand(address uint16) {
	if cpu.flagC == 0 {
		cpu.registerPC = address
	}
}

func (cpu *CPU6502) bcsCommand(address uint16) {
	if cpu.flagZ == 1 {
		cpu.registerPC = address
	}
}

func (cpu *CPU6502) beqCommand(address uint16) {
	if cpu.flagC == 0 {
		cpu.registerPC = address
	}
}

//TODO: BIT

func (cpu *CPU6502) bmiCommand(address uint16) {
	if cpu.flagN == 1 {
		cpu.registerPC = address
	}
}

func (cpu *CPU6502) bneCommand(address uint16) {
	if cpu.flagZ == 0 {
		cpu.registerPC = address
	}
}

func (cpu *CPU6502) bplCommand(address uint16) {
	if cpu.flagN == 0 {
		cpu.registerPC = address
	}
}

//TODO: BRK

func (cpu *CPU6502) bvcCommand(address uint16) {
	if cpu.flagV == 0 {
		cpu.registerPC = address
	}
}

func (cpu *CPU6502) bvsCommand(address uint16) {
	if cpu.flagV == 1 {
		cpu.registerPC = address
	}
}

func (cpu *CPU6502) clcCommand(address uint16) {
	cpu.flagC = 0
}

func (cpu *CPU6502) cldCommand(address uint16) {
	cpu.flagD = 0
}

func (cpu *CPU6502) cliCommand(address uint16) {
	//TODO:
	cpu.flagI = 0
}

func (cpu *CPU6502) clvCommand(address uint16) {
	cpu.flagV = 0
}

//TODO: CMP

//TODO: CPX

//TODO: CPY

func (cpu *CPU6502) decCommand(address uint16) {
	cpu.memory.write(address, cpu.memory.read(address)-1)
}

func (cpu *CPU6502) dexCommand(address uint16) {
	cpu.registerX -= 1
}

func (cpu *CPU6502) deyCommand(address uint16) {
	cpu.registerY -= 1
}

//TODO: EOR

func (cpu *CPU6502) incCommand(address uint16) {
	cpu.memory.write(address, cpu.memory.read(address)+1)
}

func (cpu *CPU6502) inxCommand(address uint16) {
	cpu.registerX += 1
}

func (cpu *CPU6502) inyCommand(address uint16) {
	cpu.registerY += 1
}

func (cpu *CPU6502) jmpCommand(address uint16) {
	cpu.registerPC = address
}

//TODO: JSR

func (cpu *CPU6502) ldaCommand(address uint16) {
	cpu.registerACC = cpu.memory.read(address)
}

func (cpu *CPU6502) ldxCommand(address uint16) {
	cpu.registerX = cpu.memory.read(address)
}

func (cpu *CPU6502) ldyCommand(address uint16) {
	cpu.registerY = cpu.memory.read(address)
}

func (cpu *CPU6502) lsrCommand(address uint16) {
	cpu.memory.write(address, cpu.memory.read(address)>>2)
}

func (cpu *CPU6502) nopCommand(address uint16) {
}

//TODO: ORA

//TODO: PHA

//TODO: PHP

//TODO: PLA

//TODO: PLP

//TODO: ROL

//TODO: ROR

//TODO: RTI

//TODO: RTS

//TODO: SBC

func (cpu *CPU6502) secCommand(address uint16) {
	cpu.flagC = 1
}

func (cpu *CPU6502) sedCommand(address uint16) {
	cpu.flagD = 1
}

func (cpu *CPU6502) seiCommand(address uint16) {
	cpu.flagI = 1
}

func (cpu *CPU6502) staCommand(address uint16) {
	cpu.memory.write(address, cpu.registerACC)
}

func (cpu *CPU6502) stxCommand(address uint16) {
	cpu.memory.write(address, cpu.registerX)
}

func (cpu *CPU6502) styCommand(address uint16) {
	cpu.memory.write(address, cpu.registerY)
}

func (cpu *CPU6502) taxCommand(address uint16) {
	cpu.registerX = cpu.registerACC
}

func (cpu *CPU6502) tayCommand(address uint16) {
	cpu.registerY = cpu.registerACC
}

//TODO: TSX

func (cpu *CPU6502) txaCommand(address uint16) {
	cpu.registerACC = cpu.registerX
}

//TODO: TXS

func (cpu *CPU6502) tyaCommand(address uint16) {
	cpu.registerACC = cpu.registerY
}

func (mcu *CPU6502) buildInstruction(command command, addressing addressing, cycles uint16) instruction {
	return func() {
		argument, _ := addressing()
		//mcu.registerPC += extraCycles + cycles
		command(argument)
	}
}

func sum(a uint16, b int8) uint16 {
	return uint16(int32(a) + int32(b))
}
