package main

// 1. The AST Nodes (The Tree)
type NodeType int

const (
    NODE_CALL NodeType = iota
    NODE_STRING
    NODE_NUMBER
)

type Node struct {
    Type     NodeType
    Value    string         // For Strings/Numbers
    CallName string         // For Function Calls
    Args     []*Node        // Arguments for calls
}

// 2. The Parser Structure
type Parser struct {
    lexer  *Lexer
    current Token
}

func NewParser(lexer *Lexer) *Parser {
    p := &Parser{lexer: lexer}
    p.nextToken() // Load first token
    return p
}

func (p *Parser) nextToken() {
    p.current = p.lexer.NextToken()
}

// 3. The Logic: Parse the program
func (p *Parser) Parse() *Node {
    // We expect an Identifier (function name)
    if p.current.Type != TOKEN_IDENT {
        return nil // Error or empty
    }

    callNode := &Node{
        Type:     NODE_CALL,
        CallName: p.current.Value,
    }

    p.nextToken() // Skip function name

    // Expect (
    if p.current.Type != TOKEN_LPAREN {
        return nil
    }
    
    p.nextToken() // Skip (

    // Parse Arguments
    for p.current.Type != TOKEN_RPAREN && p.current.Type != TOKEN_EOF {
        if p.current.Type == TOKEN_STRING {
            arg := &Node{Type: NODE_STRING, Value: p.current.Value}
            callNode.Args = append(callNode.Args, arg)
            p.nextToken()
        } else if p.current.Type == TOKEN_NUMBER {
            arg := &Node{Type: NODE_NUMBER, Value: p.current.Value}
            callNode.Args = append(callNode.Args, arg)
            p.nextToken()
        } else {
            p.nextToken() // Skip unknown
        }
    }

    return callNode
}