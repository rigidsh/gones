package asm

import (
	"bytes"
)

var DefaultCompiler = &Compiler{}

type Compiler struct {
}

func (compiler *Compiler) Compile(program Program) []byte {
	var buf bytes.Buffer
	for _, ins := range program.instructions {
		ins.write(&buf)
	}

	return buf.Bytes()
}
