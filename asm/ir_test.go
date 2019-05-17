package asm

import (
	"testing"
)

func TestADC_ABS(t *testing.T) {
	program := NewProgram()
	program.AddInstruction(ADC(ABS(0xFEDB)))
	result := program.Compile()
	if result[0] != 0x6D || result[1] != 0xDB || result[2] != 0xFE {
		t.Error()
	}
}

func TestBCC(t *testing.T) {
	program := NewProgram()
	program.AddInstruction(BCC(REL(-4)))
	result := program.Compile()
	if result[0] != 0x90 || result[1] != 0xFC {
		t.Error()
	}
}
