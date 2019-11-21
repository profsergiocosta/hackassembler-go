package parser

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/profsergiocosta/hack-assembler/command"
)

type Parser struct {
	tokens    []string
	position  int
	currToken string
}

func New(fname string) *Parser {
	p := new(Parser)
	reComments, err := regexp.Compile(`//.*`)
	if err != nil {
		// tratar o erro aqui
		panic("Error")
	}

	reWhiteSpace, err := regexp.Compile(` *`)
	if err != nil {
		// tratar o erro aqui
		panic("Error")
	}

	reTokens, _ := regexp.Compile(".*")

	code, _ := ioutil.ReadFile(fname)
	codeProc := reComments.ReplaceAllString(string(code), "")
	codeProc = reWhiteSpace.ReplaceAllString(string(codeProc), "")

	p.tokens = reTokens.FindAllString(codeProc, -1)
	p.position = 0
	return p
}

func (p *Parser) Reset() {
	p.position = 0
}

func (p *Parser) HasMoreCommands() bool {
	return p.position < len(p.tokens)-1
}

func (p *Parser) Advance() {
	p.currToken = p.tokens[p.position]
	p.position++
}

func split(value string, begin int, end int) string {
	runes := []rune(value)
	return string(runes[begin:end])
}

func (p *Parser) NextCommand() command.Command {
	p.Advance()
	str := p.currToken
	switch p.currToken[0] {
	case '(':
		return command.LCommand{Label: split(str, 1, len(str)-1)}
	case '@':
		return command.ACommand{At: split(str, 1, len(str))}
	default:
		dest := strings.Split(str, "=")[0]
		l := strings.Split(str, ";")
		var jmp, cmp string
		if len(l) > 1 {
			jmp = l[1]
		} else {
			jmp = ""
		}
		l2 := strings.Split(l[0], "=")
		if len(l2) > 1 {
			cmp = l2[1]
		} else {
			cmp = ""
		}
		return command.CCommand{Dest: dest, Comp: cmp, Jump: jmp}
	}

}
