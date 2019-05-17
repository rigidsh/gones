package main

import (
	"fmt"
	"github.com/rigidsh/gones/asm"
	"testing"
)

func TestADC_IMM(t *testing.T) {

	//ADC #8
	ram := []byte{0x69, 0x08}
	cpu := CreateCPU6502(RAMMemory(ram))
	cpu.registerPC = 0
	cpu.registerACC = 42
	cpu.DoIteration()

	if cpu.registerACC != 50 {
		t.Error("42 + 8 != 50")
	}
}

func TestADC_overflow_IMM(t *testing.T) {
	ram := asm.DefaultCompiler.Compile([]asm.Instruction{
		asm.ADC(asm.Value(1)),
	})
	cpu := CreateCPU6502(RAMMemory(ram))
	cpu.registerPC = 0
	cpu.registerACC = 0xFF
	cpu.DoIteration()

	if cpu.registerACC != 0 {
		t.Error("42 + 8 != 50")
	}

	if cpu.flagC != 1 {
		t.Error("flag C incorrect")
	}

	if cpu.flagV != 0 {
		t.Error("flag V incorrect")
	}

	if cpu.flagN != 0 {
		t.Error("flag N incorrect")
	}

	if cpu.flagZ != 0 {
		t.Error("flag Z incorrect")
	}
}

func TestADC_ZP(t *testing.T) {

	//ADC *4
	ram := []byte{0x65, 0x02, 0x08}
	cpu := CreateCPU6502(RAMMemory(ram))
	cpu.registerPC = 0
	cpu.registerACC = 42
	cpu.DoIteration()

	if cpu.registerACC != 50 {
		t.Error("42 + 8 != 50")
	}

	fmt.Println(cpu.registerACC)
}

func TestADC_ZPX(t *testing.T) {

	//ADC *4,X
	ram := []byte{0x75, 0x02, 0x00, 0x08}
	cpu := CreateCPU6502(RAMMemory(ram))
	cpu.registerPC = 0
	cpu.registerACC = 42
	cpu.registerX = 1
	cpu.DoIteration()

	if cpu.registerACC != 50 {
		t.Error("42 + 8 != 50")
	}

	fmt.Println(cpu.registerACC)
}
