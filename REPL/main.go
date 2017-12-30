package main

import (
	compiler "Aurora/Compiler"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("REPL Mode: AST")
	fmt.Println("To exit type 'exit'")
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		if strings.HasPrefix(strings.Trim(text, " 	"), "exit") {
			return
		}
		ast := compiler.BuildAST(text)
		program := compiler.SetupProgram(ast)
		if program == nil {
			fmt.Println("Could not compile!")
		} else {
			for _, mod := range program.Modules {
				fmt.Println(mod.Name + ":")
				fmt.Println("\tVariables:")
				for _, v := range mod.Variables {
					fmt.Printf("\t\t%+v\n", v)
				}
			}
		}
	}
}
