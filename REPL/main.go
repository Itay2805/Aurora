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
		fmt.Printf("%+v\n", compiler.BuildAST(text).Blocks)
	}
}
