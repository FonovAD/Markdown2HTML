package processing

import (
	"fmt"
)

var EmptyToken Token = Token{TokenType{"", ""}, "", 0}

var prefixHeadings = map[string]string{
	"#":      "<h1>",
	"##":     "<h2>",
	"###":    "<h3>",
	"####":   "<h4>",
	"#####":  "<h5>",
	"######": "<h6>",
}

var postfixHeadings = map[string]string{
	"#":      "</h1>",
	"##":     "</h2>",
	"###":    "</h3>",
	"####":   "</h4>",
	"#####":  "</h5>",
	"######": "</h6>",
}

type Parser struct {
	Tokens []Token
	Pos    int
}

func (P *Parser) Match(ExpectedTokenTypes []TokenType) Token {
	if P.Pos < len(P.Tokens) {
		currentToken := P.Tokens[P.Pos]
		for i := range ExpectedTokenTypes {
			if ExpectedTokenTypes[i].name == currentToken.Type.name {
				P.Pos += 1
				return currentToken
			}
		}
	}
	return Token{TokenType{"", ""}, "", 0}
}

func (P *Parser) Require(ExpectedTokenTypes []TokenType) Token {
	token := P.Match(ExpectedTokenTypes)
	emptyTokenType := TokenType{"", ""}
	if token.Type == emptyTokenType {
		errorStr := fmt.Sprintf("На позиции %d ожидается %s", P.Pos, ExpectedTokenTypes[0].name)
		panic(errorStr)
	}
	return token
}

func (P *Parser) NewParseCode() StatmentsNode { //StatmentsNode - корень дерева(все узлы это строки)
	root := StatmentsNode{CodeString: []Node{}}
	for P.Pos < len(P.Tokens) {
		line := P.ParseLine()
		root.AddNode(line)
	}
	return root
}

func (P *Parser) ParseLine() Node {
	if token := P.Match([]TokenType{TokenTypes["HEADING"]}); token != EmptyToken {
		n := P.ParseText()
		return Node{operator: token, operand: []*Node{&n}}
	}
	if token := P.Match([]TokenType{TokenTypes["LINE"]}); token != EmptyToken {
		return Node{operator: token, operand: nil}
	}
	if token := P.Match([]TokenType{TokenTypes["NUMBEREDLIST"]}); token != EmptyToken {
		P.Pos -= 1
		n := P.ParseText()
		return Node{operator: token, operand: []*Node{&n}}
	}
	if token := P.Match([]TokenType{TokenTypes["WORD"]}); token != EmptyToken {
		n := P.ParseText()
		return Node{operator: token, operand: []*Node{&n}}
	}
	var node = Node{operator: Token{}, operand: []*Node{}}
	for P.Match([]TokenType{TokenTypes["SEMICOLON"]}) == EmptyToken {
		n := P.ParseText()
		node.operand = append(node.operand, &n)
	}
	if token := P.Match([]TokenType{TokenTypes["SEMICOLON"]}); token != EmptyToken {
		return Node{operator: token, operand: []*Node{}}
	}
	return node
}

func (P *Parser) ParseList() Node {
	listnode := Node{operator: Token{Type: SecondTokenTypes["GROUPNUMBEREDLIST"], Text: "GROUPNUMBEREDLIST", Pos: P.Pos}}
	token := P.Match([]TokenType{TokenTypes["NUMBEREDLIST"]})
	for token != EmptyToken {
		node := Node{operator: token, operand: []*Node{}}
		for P.Match([]TokenType{TokenTypes["SEMICOLON"]}) != EmptyToken {
			n := P.ParseText()
			node.operand = append(node.operand, &n)
		}
		listnode.operand = append(listnode.operand, &node)
	}
	return listnode
}

func (P *Parser) ParseText() Node {
	if word := P.Match([]TokenType{TokenTypes["WORD"]}); word != EmptyToken {
		n := P.ParseText()
		return Node{operator: word, operand: []*Node{&n}}
	}
	if space := P.Match([]TokenType{TokenTypes["SPACE"]}); space != EmptyToken {
		n := P.ParseText()
		return Node{operator: space, operand: []*Node{&n}}
	}
	if operator := P.Match([]TokenType{TokenTypes["CODE"]}); operator != EmptyToken {
		CodeNodes := []*Node{}
		for P.Match([]TokenType{TokenTypes["CODE"]}) == EmptyToken {
			n := P.ParseText()
			CodeNodes = append(CodeNodes, &n)
		}
		return Node{operator: operator, operand: CodeNodes}
	}
	return Node{}
}
