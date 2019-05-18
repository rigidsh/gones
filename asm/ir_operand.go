package asm

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
