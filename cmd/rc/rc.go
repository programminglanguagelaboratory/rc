package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/programminglanguagelaboratory/rc/pkg/lexer"
	"github.com/programminglanguagelaboratory/rc/pkg/parser"
)

func main() {
	for {
		fmt.Printf("> ")

		reader := bufio.NewReader(os.Stdin)
		code, _ := reader.ReadString('\n')
		fmt.Printf("< code: %v\n", code)

		lexer := lexer.NewLexer(code)
		parser := parser.NewParser(lexer)

		ast, err := parser.Parse()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("< ast: %v\n", ast)
	}
}
