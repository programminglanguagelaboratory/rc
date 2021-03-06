package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/programminglanguagelaboratory/rc/pkg/codegen"
	"github.com/programminglanguagelaboratory/rc/pkg/desugar"
	"github.com/programminglanguagelaboratory/rc/pkg/parser"
)

func main() {
	for {
		fmt.Printf("> ")

		reader := bufio.NewReader(os.Stdin)
		code, _ := reader.ReadString('\n')
		fmt.Printf("< code: %v\n", code)

		ast, err := parser.Parse(code)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("< ast: %v\n", ast)

		desugared := desugar.Desugar(ast)
		if err != nil {
			fmt.Println(err)
			continue
		}

		codegen := codegen.NewCodegen()
		ir, err := codegen.Gen(desugared)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("< ir: %v\n", ir)
	}
}
