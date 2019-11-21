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
	"strings"

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
	fmt.Println(st)
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
			//code.CloseFile()

		} else {
			p := parser.New(path)

			fmt.Println(p)
			st := symboltable.NewSymbolTable()

			firstPass(p, st)
		}

	}

}
