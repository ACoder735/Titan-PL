package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    args := os.Args

    // 1. Handle Version
    if len(args) > 1 && args[1] == "version" {
        fmt.Println("Titan Programming Language v1.0.0")
        return
    }

    // 2. Check if file is provided
    if len(args) < 2 {
        fmt.Println("Error: No file specified.")
        fmt.Println("Usage: titan <filename.tn>")
        return
    }

    filename := args[1]

    // 3. Read the file
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error: Could not read file %s\n", filename)
        return
    }
    code := string(data)

    // 4. Lexing & Parsing
    lexer := NewLexer(code)
    parser := NewParser(lexer)
    ast := parser.Parse()

    // 5. Generate EXE
    outputName := filename + ".exe"
    err = GenerateEXE(ast, outputName)
    if err != nil {
        fmt.Printf("Error generating code: %v\n", err)
    }
}