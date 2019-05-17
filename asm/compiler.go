package asm

import (
	"bytes"
)

var DefaultCompiler = &Compiler{}

type Compiler struct {
}

func (compiler *Compiler) Compile(program []Instruction) []byte {
	var buf bytes.Buffer
	for _, ins := range program {
		ins.write(&buf)
	}

	return buf.Bytes()
}
