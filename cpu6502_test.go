package main

import (
	"fmt"
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

	fmt.Println(cpu.registerACC)
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
