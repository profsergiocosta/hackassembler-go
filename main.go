package main

/*

remove whitespace and comments
Translating 23 pre-defined symbols:

Initialization:
	Construct an empty symbol table
	Add the pre-defined symbols to the symbol table

*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/profsergiocosta/hack-assembler/code"
	"github.com/profsergiocosta/hack-assembler/command"
	"github.com/profsergiocosta/hack-assembler/parser"
	"github.com/profsergiocosta/hack-assembler/symboltable"
)

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		abs, _ := filepath.Abs(path)
		fmt.Printf("Could not find file or directory: %s \n", abs)
		os.Exit(1)
	}
	return fileInfo.IsDir()
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func firstPass(p *parser.Parser, st *symboltable.SymbolTable) {
	var curAddress symboltable.Address = 0

	for p.HasMoreCommands() {
		switch cmd := p.NextCommand().(type) {
		case command.CCommand, command.ACommand:
			curAddress++
		case command.LCommand:
			st.AddEntry(cmd.Label, curAddress)
		}
	}

}

func write(file *os.File, str string) {
	s := fmt.Sprintf("%s\n", str)
	file.WriteString(s)
}

func secondPass(p *parser.Parser, st *symboltable.SymbolTable, pathName string) {

	file, _ := os.Create(pathName)
	code := code.New()
	varAddress := 16
	fmt.Printf("Assembling to %s\n", pathName)
	for p.HasMoreCommands() {
		switch cmd := p.NextCommand().(type) {
		case command.CCommand:
			write(file, code.GenCCommand(cmd.Dest, cmd.Comp, cmd.Jump))
		case command.ACommand:
			address, hasAddress := st.GetAddress(cmd.At)
			if hasAddress {
				write(file, code.GenACommand(address))
			} else {
				address, err := strconv.Atoi(cmd.At)
				if err == nil {
					write(file, code.GenACommand(symboltable.Address(address)))
				} else {
					st.AddEntry(cmd.At, symboltable.Address(varAddress))
					//fmt.Println(cmd.At)
					write(file, code.GenACommand(symboltable.Address(varAddress)))
					varAddress++
				}
			}
		default:

		}
	}
	file.Close()
}

func main() {
	arg := os.Args[1:]

	if len(arg) == 1 {
		path := arg[0]

		if isDirectory(path) {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range files {

				if filepath.Ext(f.Name()) == ".vm" {

				}

			}

		} else {
			p := parser.New(path)

			st := symboltable.NewSymbolTable()

			firstPass(p, st)
			p.Reset()
			abs, _ := filepath.Abs(path)
			secondPass(p, st, filenameWithoutExtension(abs)+".hack")
		}

	}

}
