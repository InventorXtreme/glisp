package main

import "fmt"
import "os"
import "github.com/InventorXtreme/glisp/lexer"

func main() {
    args := os.Args
	fmt.Println(lexer.Lex(args[1]))
}
