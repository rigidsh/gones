package asm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

func NewProgram() *Program {
	return &Program{
		instructions: make([]Instruction, 0),
	}
}

type Program struct {
	instructions []Instruction
}

func (program *Program) AddInstruction(instruction Instruction) {
	program.instructions = append(program.instructions, instruction)
}

func (program Program) Compile() []byte {
	var buf bytes.Buffer
	for _, ins := range program.instructions {
		ins.write(&buf)
	}

	return buf.Bytes()
}

type Instruction interface {
	write(writer io.Writer)
}

type OperandType int

var ABSOperandType OperandType = 0
var ABSXOperandType OperandType = 1
var ABSYOperandType OperandType = 2
var ACCOperandType OperandType = 3
var IMMOperandType OperandType = 4
var IMPLOperandType OperandType = 5
var INDOperandType OperandType = 6
var INDXOperandType OperandType = 7
var INDYOperandType OperandType = 8
var RELOperandType OperandType = 9
var ZPOperandType OperandType = 10
var ZPXOperandType OperandType = 11
var ZPYOperandType OperandType = 12

type Operand interface {
	write(writer io.Writer)
	GetType() OperandType
}

type simpleOperand struct {
	operandType OperandType
	value       int
}

func (op *simpleOperand) write(writer io.Writer) {
	var buf []byte
	switch op.operandType {
	case IMPLOperandType:
		buf = []byte{}
	case ABSOperandType, ABSXOperandType, ABSYOperandType, INDOperandType, INDXOperandType, INDYOperandType:
		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(op.value))
	case RELOperandType:
		buf = []byte{byte(int8(op.value))}
	default:
		buf = []byte{byte(op.value)}
	}

	writer.Write(buf)
}

func (op *simpleOperand) GetType() OperandType {
	return op.operandType
}

type simpleInstruction struct {
	values  map[OperandType]byte
	operand Operand
}

func (ins *simpleInstruction) write(writer io.Writer) {
	cmd, ok := ins.values[ins.operand.GetType()]
	if !ok {
		panic(errors.New("invalid command"))
	}
	writer.Write([]byte{cmd})
	ins.operand.write(writer)
}

func ADC(operand Operand) Instruction {
	return &simpleInstruction{
		values: map[OperandType]byte{
			IMMOperandType:  0x69,
			ZPOperandType:   0x65,
			ZPXOperandType:  0x75,
			ABSOperandType:  0x6D,
			ABSXOperandType: 0x7D,
			ABSYOperandType: 0x79,
			INDXOperandType: 0x61,
			INDYOperandType: 0x71,
		},
		operand: operand,
	}
}

func AND(operand Operand) Instruction {
	return &simpleInstruction{
		values: map[OperandType]byte{
			IMMOperandType:  0x29,
			ZPOperandType:   0x25,
			ZPXOperandType:  0x35,
			ABSOperandType:  0x2D,
			ABSXOperandType: 0x3D,
			ABSYOperandType: 0x39,
			INDXOperandType: 0x21,
			INDYOperandType: 0x31,
		},
		operand: operand,
	}
}

func JMP(operand Operand) Instruction {
	return &simpleInstruction{
		values: map[OperandType]byte{
			ABSOperandType: 0x4C,
			INDOperandType: 0x6C,
		},
		operand: operand,
	}
}

func BCC(operand Operand) Instruction {
	return &simpleInstruction{
		values: map[OperandType]byte{
			RELOperandType: 0x90,
		},
		operand: operand,
	}
}

func Value(value uint8) Operand {
	return &simpleOperand{
		operandType: IMMOperandType,
		value:       int(value),
	}
}

func ABS(value uint16) Operand {
	return &simpleOperand{
		operandType: ABSOperandType,
		value:       int(value),
	}
}

func IND(value uint16) Operand {
	return &simpleOperand{
		operandType: INDOperandType,
		value:       int(value),
	}
}

func REL(value int8) Operand {
	return &simpleOperand{
		operandType: RELOperandType,
		value:       int(value),
	}
}
