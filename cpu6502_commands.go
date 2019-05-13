package main

func (cpu *CPU6502) adcCommand(address uint16) {
	tmp := uint16(cpu.registerACC) + uint16(cpu.memory.read(address)) + uint16(cpu.flagC)

	if tmp > 0xFF {
		cpu.flagC = 1
	} else {
		cpu.flagC = 0
	}

	//TODO: check
	if cpu.registerACC^cpu.memory.read(address)&0x80 == 0 && cpu.registerACC^uint8(tmp)&0x80 != 0 {
		cpu.flagV = 1
	} else {
		cpu.flagV = 0
	}

	cpu.setAcc(uint8(tmp))
}

func (cpu *CPU6502) andCommand(address uint16) {
	cpu.setAcc(cpu.registerACC & cpu.memory.read(address))
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

func (cpu *CPU6502) jsrCommand(address uint16) {
	cpu.pushStack16(cpu.registerPC)
	cpu.registerPC = address
}

func (cpu *CPU6502) ldaCommand(address uint16) {
	cpu.setAcc(cpu.memory.read(address))
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

func (cpu *CPU6502) phaCommand(address uint16) {
	cpu.pushStack(cpu.registerACC)
}

//TODO: PHP

func (cpu *CPU6502) plaCommand(address uint16) {
	cpu.setAcc(cpu.pullStack())
}

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
	cpu.setAcc(cpu.registerX)
}

//TODO: TXS

func (cpu *CPU6502) tyaCommand(address uint16) {
	cpu.setAcc(cpu.registerY)
}
