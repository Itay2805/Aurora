package main

import (
	compiler "Aurora/Compiler"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("REPL Mode: AST")
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		fmt.Printf("%+v\n", compiler.BuildAST(text).Expressions)
	}
}
