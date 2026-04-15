package main

import (
    "fmt"
    "os"
)

func main() {
    // 1. Handle Arguments
    args := os.Args
    if len(args) > 1 && args[1] == "version" {
        fmt.Println("Titan Programming Language v1.0.0")
        return
    }

    // 2. The Code to Test (Later we will read from a file)
    code := `typeln("Hello Titan!")`

    // 3. Initialize Lexer
    lexer := NewLexer(code)

    // 4. Print Tokens (Temporary test)
    fmt.Println("--- TITAN LEXER OUTPUT ---")
    for {
        tok := lexer.NextToken()
        fmt.Printf("Type: %s, Value: %q\n", tok.Type, tok.Value)
        if tok.Type == TOKEN_EOF { break }
    }
}