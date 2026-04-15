package main

import (
    "unicode"
)

// --- Types ---
type TokenType int

const (
    TOKEN_EOF TokenType = iota
    TOKEN_ILLEGAL
    
    TOKEN_IDENT  // typeln, x
    TOKEN_NUMBER // 123
    TOKEN_STRING // "Hello"
    
    TOKEN_LPAREN // (
    TOKEN_RPAREN // )
    TOKEN_LBRACE // {
    TOKEN_RBRACE // }
    
    TOKEN_ASSIGN // :=
    TOKEN_PLUS   // +
)

func (t TokenType) String() string {
    switch t {
    case TOKEN_EOF: return "EOF"
    case TOKEN_ILLEGAL: return "ILLEGAL"
    case TOKEN_IDENT: return "IDENT"
    case TOKEN_NUMBER: return "NUMBER"
    case TOKEN_STRING: return "STRING"
    case TOKEN_LPAREN: return "("
    case TOKEN_RPAREN: return ")"
    case TOKEN_LBRACE: return "{"
    case TOKEN_RBRACE: return "}"
    case TOKEN_ASSIGN: return ":="
    case TOKEN_PLUS: return "+"
    default: return "UNKNOWN"
    }
}

type Token struct {
    Type  TokenType
    Value string
}

// --- Lexer ---
type Lexer struct {
    input   []rune
    pos     int
}

func NewLexer(input string) *Lexer {
    return &Lexer{input: []rune(input)}
}

func (l *Lexer) peek() rune {
    if l.pos >= len(l.input) { return 0 }
    return l.input[l.pos]
}

func (l *Lexer) next() {
    l.pos++
}

func (l *Lexer) skipWhitespace() {
    for unicode.IsSpace(l.peek()) { l.next() }
}

func (l *Lexer) NextToken() Token {
    l.skipWhitespace()
    current := l.peek()

    switch current {
    case '(': l.next(); return Token{TOKEN_LPAREN, "("}
    case ')': l.next(); return Token{TOKEN_RPAREN, ")"}
    case '{': l.next(); return Token{TOKEN_LBRACE, "{"}
    case '}': l.next(); return Token{TOKEN_RBRACE, "}"}
    case '"':
        l.next()
        start := l.pos
        for l.peek() != '"' && l.peek() != 0 { l.next() }
        value := string(l.input[start:l.pos])
        l.next()
        return Token{TOKEN_STRING, value}
    case ':':
        l.next()
        if l.peek() == '=' {
            l.next()
            return Token{TOKEN_ASSIGN, ":="}
        }
        return Token{TOKEN_ILLEGAL, ":"}
    case 0: return Token{TOKEN_EOF, ""}
    }

    if unicode.IsDigit(current) {
        start := l.pos
        for unicode.IsDigit(l.peek()) { l.next() }
        return Token{TOKEN_NUMBER, string(l.input[start:l.pos])}
    }

    if unicode.IsLetter(current) {
        start := l.pos
        for unicode.IsLetter(l.peek()) { l.next() }
        return Token{TOKEN_IDENT, string(l.input[start:l.pos])}
    }

    l.next()
    return Token{TOKEN_ILLEGAL, string(current)}
}