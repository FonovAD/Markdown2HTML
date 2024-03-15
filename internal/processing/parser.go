package processing

import (
	"fmt"
)

var EmptyToken Token = Token{TokenType{"", ""}, "", 0}

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
	if token := P.Match([]TokenType{TokenTypeList["HEADING"]}); token != EmptyToken {
		n := P.ParseText()
		return Node{operator: token, operand: []*Node{&n}}
	}
	if token := P.Match([]TokenType{TokenTypeList["LINE"]}); token != EmptyToken {
		return Node{operator: token, operand: nil}
	}
	if token := P.Match([]TokenType{TokenTypeList["NUMBEREDLIST"]}); token != EmptyToken {
		P.Pos -= 1
		n := P.ParseText()
		return Node{operator: token, operand: []*Node{&n}}
	}
	var node = Node{operator: Token{}, operand: []*Node{}}
	for P.Match([]TokenType{TokenTypeList["SEMICOLON"]}) != EmptyToken {
		n := P.ParseText()
		node.operand = append(node.operand, &n)
	}
	return node
}

func (P *Parser) ParseList() Node {
	listnode := Node{operator: Token{Type: TokenTypeList["GROUPNUMBEREDLIST"], Text: "GROUPNUMBEREDLIST", Pos: P.Pos}}
	token := P.Match([]TokenType{TokenTypeList["NUMBEREDLIST"]})
	for token != EmptyToken {
		node := Node{operator: token, operand: []*Node{}}
		for P.Match([]TokenType{TokenTypeList["SEMICOLON"]}) != EmptyToken {
			n := P.ParseText()
			node.operand = append(node.operand, &n)
		}
		listnode.operand = append(listnode.operand, &node)
	}
	return listnode
}

func (P *Parser) ParseText() Node {
	if word := P.Match([]TokenType{TokenTypeList["WORD"]}); word != EmptyToken {
		n := P.ParseText()
		return Node{operator: word, operand: []*Node{&n}}
	}
	if space := P.Match([]TokenType{TokenTypeList["SPACE"]}); space != EmptyToken {
		n := P.ParseText()
		return Node{operator: space, operand: []*Node{&n}}
	}
	if operator := P.Match([]TokenType{TokenTypeList["CODE"]}); operator != EmptyToken {
		CodeNodes := []*Node{}
		for P.Match([]TokenType{TokenTypeList["CODE"]}) == EmptyToken {
			if token := P.Match([]TokenType{TokenTypeList["WORD"]}); token != EmptyToken {
				CodeNodes = append(CodeNodes, &Node{operator: token, operand: nil})
			}
			if token := P.Match([]TokenType{TokenTypeList["SPACE"]}); token != EmptyToken {
				CodeNodes = append(CodeNodes, &Node{operator: token, operand: nil})
			}
		}
		P.Require([]TokenType{TokenTypeList["CODE"]})
		return Node{operator: operator, operand: CodeNodes}
	}
	return Node{}
}
