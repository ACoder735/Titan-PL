package main

import (
    "fmt"
    "os"
    "unicode"
)

// 1. Define the Token Types (The "Fake Enum" using iota)
type TokenType int

const (
    TOKEN_EOF TokenType = iota
    TOKEN_ILLEGAL
    
    TOKEN_IDENT  // typeln, x, my_func
    TOKEN_NUMBER // 123
    TOKEN_STRING // "Hello"
    
    TOKEN_LPAREN // (
    TOKEN_RPAREN // )
    TOKEN_LBRACE // {
    TOKEN_RBRACE // }
    
    TOKEN_ASSIGN // :=
    TOKEN_PLUS   // +
)

// Add this right after the Token constants
func (t TokenType) String() string {
    switch t {
    case TOKEN_EOF:
        return "EOF"
    case TOKEN_ILLEGAL:
        return "ILLEGAL"
    case TOKEN_IDENT:
        return "IDENT"
    case TOKEN_NUMBER:
        return "NUMBER"
    case TOKEN_STRING:
        return "STRING"
    case TOKEN_LPAREN:
        return "("
    case TOKEN_RPAREN:
        return ")"
    case TOKEN_LBRACE:
        return "{"
    case TOKEN_RBRACE:
        return "}"
    case TOKEN_ASSIGN:
        return ":="
    case TOKEN_PLUS:
        return "+"
    default:
        return "UNKNOWN"
    }
}

// 2. The Token Structure
type Token struct {
    Type  TokenType
    Value string
}

// 3. The Lexer Structure
type Lexer struct {
    input   []rune // The source code as characters
    pos     int    // Current position
}

// Helper to look at current character
func (l *Lexer) peek() rune {
    if l.pos >= len(l.input) {
        return 0 // EOF
    }
    return l.input[l.pos]
}

// Helper to move to next character
func (l *Lexer) next() {
    l.pos++
}

// Helper to skip spaces
func (l *Lexer) skipWhitespace() {
    for unicode.IsSpace(l.peek()) {
        l.next()
    }
}

// 4. The Main Logic: Read one token at a time
func (l *Lexer) NextToken() Token {
    l.skipWhitespace()

    current := l.peek()

    // Handle Symbols
    switch current {
    case '(':
        l.next()
        return Token{TOKEN_LPAREN, "("}
    case ')':
        l.next()
        return Token{TOKEN_RPAREN, ")"}
    case '{':
        l.next()
        return Token{TOKEN_LBRACE, "{"}
    case '}':
        l.next()
        return Token{TOKEN_RBRACE, "}"}
    case '"': // Start of a string
        l.next()
        start := l.pos
        for l.peek() != '"' && l.peek() != 0 {
            l.next()
        }
        value := string(l.input[start:l.pos])
        l.next() // skip closing quote
        return Token{TOKEN_STRING, value}
    case ':': // Check for :=
        l.next()
        if l.peek() == '=' {
            l.next()
            return Token{TOKEN_ASSIGN, ":="}
        }
        return Token{TOKEN_ILLEGAL, ":"}
    case 0:
        return Token{TOKEN_EOF, ""}
    }

    // Handle Numbers
    if unicode.IsDigit(current) {
        start := l.pos
        for unicode.IsDigit(l.peek()) {
            l.next()
        }
        return Token{TOKEN_NUMBER, string(l.input[start:l.pos])}
    }

    // Handle Identifiers (typeln, x, etc.)
    if unicode.IsLetter(current) {
        start := l.pos
        for unicode.IsLetter(l.peek()) {
            l.next()
        }
        return Token{TOKEN_IDENT, string(l.input[start:l.pos])}
    }

    // Unknown character
    l.next()
    return Token{TOKEN_ILLEGAL, string(current)}
}

func runCode(code string){
	lexer := &Lexer{input: []rune(code)}
    fmt.Println("---")
    for {
        tok := lexer.NextToken()
        fmt.Printf("Type: %s, Value: %q\n", tok.Type, tok.Value)
        
        if tok.Type == TOKEN_EOF {
            break
        }
    }
}

func main() {
    args := os.Args
    if len(args) > 1 && args[1] == "version" {
        fmt.Println("Titan version 1.0.0")
        return
    }
	code := `typeln("Hello Titan!")`
	runCode(code)
}