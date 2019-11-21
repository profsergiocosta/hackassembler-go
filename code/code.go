package code

import (
	"fmt"
	"strings"

	"github.com/profsergiocosta/hack-assembler/symboltable"
)

type Code struct {
	cmpTable  map[string]string
	destTable map[string]string
	jmpTable  map[string]string
}

func New() *Code {
	code := new(Code)
	code.cmpTable = map[string]string{
		"":  "0000000",
		"0": "0101010", "1": "0111111", "-1": "0111010", "D": "0001100",
		"A": "0110000", "!D": "0001101", "!A": "0110001", "-D": "0001111",
		"-A": "0110011", "D+1": "0011111", "A+1": "0110111", "D-1": "0001110",
		"A-1": "0110010", "D+A": "0000010", "D-A": "0010011", "A-D": "0000111",
		"D&A": "0000000", "D|A": "0010101",
		"M": "1110000", "!M": "1110001",
		"-M": "1110011", "M+1": "1110111",
		"M-1": "1110010", "D+M": "1000010", "D-M": "1010011", "M-D": "1000111",
		"D&M": "1000000", "D|M": "1010101",
	}

	code.destTable = map[string]string{
		"": "000", "M": "001", "D": "010",
		"MD": "011", "A": "100", "AM": "101", "AD": "110", "AMD": "111",
	}

	code.jmpTable = map[string]string{
		"": "000", "JGT": "001", "JEQ": "010",
		"JGE": "011", "JLT": "100", "JNE": "101", "JLE": "110", "JMP": "111",
	}
	return code
}

func (code *Code) comp(value string) string {
	return code.cmpTable[value]
}

func (code *Code) dest(value string) string {
	return code.destTable[value]
}

func (code *Code) jump(value string) string {
	return code.jmpTable[value]
}

func convertBinary(value symboltable.Address) string {
	s := fmt.Sprintf("%b", value)
	repeat := 15 - len(s)
	s = strings.Repeat("0", repeat) + s
	return s
}

func (code *Code) GenACommand(addr symboltable.Address) string {
	return "0" + convertBinary(addr)
}

func (code *Code) GenCCommand(dest string, comp string, jump string) string {
	return "111" + code.comp(comp) + code.dest(dest) + code.jump(jump)
}
