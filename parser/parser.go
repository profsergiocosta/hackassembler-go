package parser

import (
	"io/ioutil"
	"regexp"
	"strings"

	"hackassembler-go/command"
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

	tokens := reTokens.FindAllString(codeProc, -1)
	// solucao ineficiente TODO: remover o uso de expressoes regulares
	var s []string
	for _, tk := range tokens {
		tkt := strings.TrimSpace(tk)
		if tk != "" && tk[0] != 13 {
			s = append(s, tkt)
		}
	}
	p.tokens = s
	p.position = 0
	return p
}

func (p *Parser) Reset() {
	p.position = 0
}

func (p *Parser) HasMoreCommands() bool {
	return p.position < len(p.tokens)
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
		var dest, cmp, jmp string
		if strings.Index(str, "=") != -1 { // tem destino
			dest = strings.Split(str, "=")[0]

			rest := strings.Split(str, "=")[1]
			if strings.Index(rest, ";") != -1 { // tem jump
				jmp = strings.Split(rest, ";")[1]
				cmp = strings.Split(rest, ";")[0]
			} else {
				cmp = rest
			}
		} else { // nao tem destino
			if strings.Index(str, ";") != -1 { // tem jump
				jmp = strings.Split(str, ";")[1]
				cmp = strings.Split(str, ";")[0]
			} else {
				cmp = str
			}
		}
		return command.CCommand{Dest: dest, Comp: cmp, Jump: jmp}
	}

}
