package asm

var adcCommandMap = map[OperandType]byte{
	IMMOperandType:  0x69,
	ZPOperandType:   0x65,
	ZPXOperandType:  0x75,
	ABSOperandType:  0x6D,
	ABSXOperandType: 0x7D,
	ABSYOperandType: 0x79,
	INDXOperandType: 0x61,
	INDYOperandType: 0x71,
}

var andCommandMap = map[OperandType]byte{
	IMMOperandType:  0x29,
	ZPOperandType:   0x25,
	ZPXOperandType:  0x35,
	ABSOperandType:  0x2D,
	ABSXOperandType: 0x3D,
	ABSYOperandType: 0x39,
	INDXOperandType: 0x21,
	INDYOperandType: 0x31,
}

var aslCommandMap = map[OperandType]byte{
	ACCOperandType:  0x0A,
	ZPOperandType:   0x06,
	ZPXOperandType:  0x16,
	ABSOperandType:  0x0E,
	ABSXOperandType: 0x1E,
}

var bccCommandMap = map[OperandType]byte{
	RELOperandType: 0x90,
}

var bcsCommandMap = map[OperandType]byte{
	RELOperandType: 0xB0,
}

var beqCommandMap = map[OperandType]byte{
	RELOperandType: 0xF0,
}

var bitCommandMap = map[OperandType]byte{
	ZPOperandType:  0x24,
	ABSOperandType: 0x2C,
}

var bmiCommandMap = map[OperandType]byte{
	RELOperandType: 0x30,
}

var bneCommandMap = map[OperandType]byte{
	RELOperandType: 0xD0,
}

var bplCommandMap = map[OperandType]byte{
	RELOperandType: 0x10,
}

var brkCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x00,
}

var bvcCommandMap = map[OperandType]byte{
	RELOperandType: 0x50,
}

var bvsCommandMap = map[OperandType]byte{
	RELOperandType: 0x70,
}

var clcCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x18,
}

var cldCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xD8,
}

var cliCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x58,
}

var clvCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xB8,
}

var cmpCommandMap = map[OperandType]byte{
	IMMOperandType:  0xC9,
	ZPOperandType:   0xC5,
	ZPXOperandType:  0xD5,
	ABSOperandType:  0xCD,
	ABSXOperandType: 0xDD,
	ABSYOperandType: 0xD9,
	INDXOperandType: 0xC1,
	INDYOperandType: 0xD1,
}

var cpxCommandMap = map[OperandType]byte{
	IMMOperandType: 0xE0,
	ZPOperandType:  0xE4,
	ABSOperandType: 0xEC,
}

var cpyCommandMap = map[OperandType]byte{
	IMMOperandType: 0xC0,
	ZPOperandType:  0xC4,
	ABSOperandType: 0xCC,
}

var decCommandMap = map[OperandType]byte{
	ZPOperandType:   0x6C,
	ZPXOperandType:  0xDC,
	ABSOperandType:  0xCE,
	ABSXOperandType: 0xDE,
}

var dexCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xCA,
}

var deyCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x88,
}

var eorCommandMap = map[OperandType]byte{
	IMMOperandType:  0x49,
	ZPOperandType:   0x45,
	ZPXOperandType:  0x55,
	ABSOperandType:  0x4D,
	ABSXOperandType: 0x5D,
	ABSYOperandType: 0x59,
	INDXOperandType: 0x41,
	INDYOperandType: 0x51,
}

var incCommandMap = map[OperandType]byte{
	ZPOperandType:   0xE6,
	ZPXOperandType:  0xF6,
	ABSOperandType:  0xEE,
	ABSXOperandType: 0xFE,
}

var inxCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xE8,
}

var inyCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xC8,
}

var jmpCommandMap = map[OperandType]byte{
	ABSOperandType: 0x4C,
	INDOperandType: 0x6C,
}

var jsrCommandMap = map[OperandType]byte{
	ABSOperandType: 0x20,
}

var ldaCommandMap = map[OperandType]byte{
	IMMOperandType:  0xA9,
	ZPOperandType:   0xA5,
	ZPXOperandType:  0xB5,
	ABSOperandType:  0xAD,
	ABSXOperandType: 0xBD,
	ABSYOperandType: 0xB9,
	INDXOperandType: 0xA1,
	INDYOperandType: 0xB1,
}

var ldxCommandMap = map[OperandType]byte{
	IMMOperandType:  0xA2,
	ZPOperandType:   0xA6,
	ZPYOperandType:  0xB6,
	ABSOperandType:  0xAE,
	ABSYOperandType: 0xBE,
}

var ldyCommandMap = map[OperandType]byte{
	IMMOperandType:  0xA0,
	ZPOperandType:   0xA4,
	ZPXOperandType:  0xB4,
	ABSOperandType:  0xAC,
	ABSXOperandType: 0xBC,
}

var lsrCommandMap = map[OperandType]byte{
	ACCOperandType:  0x4A,
	ZPOperandType:   0x46,
	ZPXOperandType:  0x56,
	ABSOperandType:  0x4E,
	ABSXOperandType: 0x5E,
}

var nopCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xEA,
}

var oraCommandMap = map[OperandType]byte{
	IMMOperandType:  0x09,
	ZPOperandType:   0x05,
	ZPXOperandType:  0x15,
	ABSOperandType:  0x0D,
	ABSXOperandType: 0x1D,
	ABSYOperandType: 0x19,
	INDXOperandType: 0x01,
	INDYOperandType: 0x11,
}

var phaCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x48,
}

var phpCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x08,
}

var plaCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x68,
}

var plpCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x28,
}

var rolCommandMap = map[OperandType]byte{
	ACCOperandType:  0x2A,
	ZPOperandType:   0x26,
	ZPXOperandType:  0x36,
	ABSOperandType:  0x2E,
	ABSXOperandType: 0x3E,
}

var rorCommandMap = map[OperandType]byte{
	ACCOperandType:  0x6A,
	ZPOperandType:   0x66,
	ZPXOperandType:  0x76,
	ABSOperandType:  0x6E,
	ABSXOperandType: 0x7E,
}

var rtiCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x40,
}

var rtsCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x60,
}

var sbcCommandMap = map[OperandType]byte{
	IMMOperandType:  0xE9,
	ZPOperandType:   0xE5,
	ZPXOperandType:  0xF5,
	ABSOperandType:  0xED,
	ABSXOperandType: 0xFD,
	ABSYOperandType: 0xF9,
	INDXOperandType: 0xE1,
	INDYOperandType: 0xF1,
}

var secCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x38,
}

var sedCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xF8,
}

var seiCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x78,
}

var staCommandMap = map[OperandType]byte{
	ZPOperandType:   0x85,
	ZPXOperandType:  0x95,
	ABSOperandType:  0x8D,
	ABSXOperandType: 0x9D,
	ABSYOperandType: 0x99,
	INDXOperandType: 0x81,
	INDYOperandType: 0x91,
}

var stxCommandMap = map[OperandType]byte{
	ZPOperandType:  0x86,
	ZPYOperandType: 0x96,
	ABSOperandType: 0x8E,
}

var styCommandMap = map[OperandType]byte{
	ZPOperandType:  0x84,
	ZPXOperandType: 0x94,
	ABSOperandType: 0x8C,
}

var taxCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xAA,
}

var tayCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xA8,
}

var tsxCommandMap = map[OperandType]byte{
	IMPLOperandType: 0xBA,
}

var txaCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x8A,
}

var txsCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x9A,
}

var tyaCommandMap = map[OperandType]byte{
	IMPLOperandType: 0x98,
}

//Commands

func ADC(operand Operand) Instruction {
	return &simpleInstruction{
		values:  adcCommandMap,
		operand: operand,
	}
}

func AND(operand Operand) Instruction {
	return &simpleInstruction{
		values:  andCommandMap,
		operand: operand,
	}
}

func ASL(operand Operand) Instruction {
	return &simpleInstruction{
		values:  aslCommandMap,
		operand: operand,
	}
}

func BCC(operand Operand) Instruction {
	return &simpleInstruction{
		values:  bccCommandMap,
		operand: operand,
	}
}

func BCS(operand Operand) Instruction {
	return &simpleInstruction{
		values:  bcsCommandMap,
		operand: operand,
	}
}

func BEQ(operand Operand) Instruction {
	return &simpleInstruction{
		values:  beqCommandMap,
		operand: operand,
	}
}

func BIT(operand Operand) Instruction {
	return &simpleInstruction{
		values:  bitCommandMap,
		operand: operand,
	}
}

func BMI(operand Operand) Instruction {
	return &simpleInstruction{
		values:  bmiCommandMap,
		operand: operand,
	}
}

func BNE(operand Operand) Instruction {
	return &simpleInstruction{
		values:  bneCommandMap,
		operand: operand,
	}
}

func BPL(operand Operand) Instruction {
	return &simpleInstruction{
		values:  bplCommandMap,
		operand: operand,
	}
}

func BRK(operand Operand) Instruction {
	return &simpleInstruction{
		values:  brkCommandMap,
		operand: operand,
	}
}

func BVC(operand Operand) Instruction {
	return &simpleInstruction{
		values:  bvcCommandMap,
		operand: operand,
	}
}

func BVS(operand Operand) Instruction {
	return &simpleInstruction{
		values:  bvsCommandMap,
		operand: operand,
	}
}

func CLC(operand Operand) Instruction {
	return &simpleInstruction{
		values:  clcCommandMap,
		operand: operand,
	}
}

func CLD(operand Operand) Instruction {
	return &simpleInstruction{
		values:  cldCommandMap,
		operand: operand,
	}
}

func CLI(operand Operand) Instruction {
	return &simpleInstruction{
		values:  cliCommandMap,
		operand: operand,
	}
}

func CLV(operand Operand) Instruction {
	return &simpleInstruction{
		values:  clvCommandMap,
		operand: operand,
	}
}

func CMP(operand Operand) Instruction {
	return &simpleInstruction{
		values:  cmpCommandMap,
		operand: operand,
	}
}

func CPX(operand Operand) Instruction {
	return &simpleInstruction{
		values:  cpxCommandMap,
		operand: operand,
	}
}

func CPY(operand Operand) Instruction {
	return &simpleInstruction{
		values:  cpyCommandMap,
		operand: operand,
	}
}

func DEC(operand Operand) Instruction {
	return &simpleInstruction{
		values:  decCommandMap,
		operand: operand,
	}
}

func DEX(operand Operand) Instruction {
	return &simpleInstruction{
		values:  dexCommandMap,
		operand: operand,
	}
}

func DEY(operand Operand) Instruction {
	return &simpleInstruction{
		values:  deyCommandMap,
		operand: operand,
	}
}

func EOR(operand Operand) Instruction {
	return &simpleInstruction{
		values:  eorCommandMap,
		operand: operand,
	}
}

func INC(operand Operand) Instruction {
	return &simpleInstruction{
		values:  incCommandMap,
		operand: operand,
	}
}

func INX(operand Operand) Instruction {
	return &simpleInstruction{
		values:  inxCommandMap,
		operand: operand,
	}
}

func INY(operand Operand) Instruction {
	return &simpleInstruction{
		values:  inyCommandMap,
		operand: operand,
	}
}

func JMP(operand Operand) Instruction {
	return &simpleInstruction{
		values:  jmpCommandMap,
		operand: operand,
	}
}

func JSR(operand Operand) Instruction {
	return &simpleInstruction{
		values:  jsrCommandMap,
		operand: operand,
	}
}

func LDA(operand Operand) Instruction {
	return &simpleInstruction{
		values:  ldaCommandMap,
		operand: operand,
	}
}

func LDX(operand Operand) Instruction {
	return &simpleInstruction{
		values:  ldxCommandMap,
		operand: operand,
	}
}

func LDY(operand Operand) Instruction {
	return &simpleInstruction{
		values:  ldyCommandMap,
		operand: operand,
	}
}

func LSR(operand Operand) Instruction {
	return &simpleInstruction{
		values:  lsrCommandMap,
		operand: operand,
	}
}

func NOP(operand Operand) Instruction {
	return &simpleInstruction{
		values:  nopCommandMap,
		operand: operand,
	}
}

func ORA(operand Operand) Instruction {
	return &simpleInstruction{
		values:  oraCommandMap,
		operand: operand,
	}
}

func PHA(operand Operand) Instruction {
	return &simpleInstruction{
		values:  phaCommandMap,
		operand: operand,
	}
}

func PHP(operand Operand) Instruction {
	return &simpleInstruction{
		values:  phpCommandMap,
		operand: operand,
	}
}

func PLA(operand Operand) Instruction {
	return &simpleInstruction{
		values:  plaCommandMap,
		operand: operand,
	}
}

func PLP(operand Operand) Instruction {
	return &simpleInstruction{
		values:  plpCommandMap,
		operand: operand,
	}
}

func ROL(operand Operand) Instruction {
	return &simpleInstruction{
		values:  rolCommandMap,
		operand: operand,
	}
}

func ROR(operand Operand) Instruction {
	return &simpleInstruction{
		values:  rorCommandMap,
		operand: operand,
	}
}

func RTI(operand Operand) Instruction {
	return &simpleInstruction{
		values:  rtiCommandMap,
		operand: operand,
	}
}

func RTS(operand Operand) Instruction {
	return &simpleInstruction{
		values:  rtsCommandMap,
		operand: operand,
	}
}

func SBC(operand Operand) Instruction {
	return &simpleInstruction{
		values:  sbcCommandMap,
		operand: operand,
	}
}

func SEC(operand Operand) Instruction {
	return &simpleInstruction{
		values:  secCommandMap,
		operand: operand,
	}
}

func SED(operand Operand) Instruction {
	return &simpleInstruction{
		values:  sedCommandMap,
		operand: operand,
	}
}

func SEI(operand Operand) Instruction {
	return &simpleInstruction{
		values:  seiCommandMap,
		operand: operand,
	}
}

func STA(operand Operand) Instruction {
	return &simpleInstruction{
		values:  staCommandMap,
		operand: operand,
	}
}

func STX(operand Operand) Instruction {
	return &simpleInstruction{
		values:  stxCommandMap,
		operand: operand,
	}
}

func STY(operand Operand) Instruction {
	return &simpleInstruction{
		values:  styCommandMap,
		operand: operand,
	}
}

func TAX(operand Operand) Instruction {
	return &simpleInstruction{
		values:  taxCommandMap,
		operand: operand,
	}
}

func TAY(operand Operand) Instruction {
	return &simpleInstruction{
		values:  tayCommandMap,
		operand: operand,
	}
}

func TSX(operand Operand) Instruction {
	return &simpleInstruction{
		values:  tsxCommandMap,
		operand: operand,
	}
}

func TXA(operand Operand) Instruction {
	return &simpleInstruction{
		values:  txaCommandMap,
		operand: operand,
	}
}

func TXS(operand Operand) Instruction {
	return &simpleInstruction{
		values:  txsCommandMap,
		operand: operand,
	}
}

func TYA(operand Operand) Instruction {
	return &simpleInstruction{
		values:  tyaCommandMap,
		operand: operand,
	}
}
