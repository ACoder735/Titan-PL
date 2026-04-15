package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args
    if len(args) > 1 && args[1] == "version" {
        fmt.Println("Titan Programming Language v1.0.0")
        return
    }

    code := `typeln("Hello Titan!")`

    // 1. Lexing
    lexer := NewLexer(code)

    // 2. Parsing
    parser := NewParser(lexer)
    ast := parser.Parse()

    // 3. Print the Result
    if ast != nil {
        fmt.Printf("Found Call: %s\n", ast.CallName)
        fmt.Printf("Argument: %s\n", ast.Args[0].Value)
    } else {
        fmt.Println("Parsing Error!")
    }
}