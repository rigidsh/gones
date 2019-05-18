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
