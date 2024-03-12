package processing

type TokenType struct {
	name  string
	regex string
}

var TokenTypeList = map[string]TokenType{
	"SEMICOLON":    TokenType{"SEMICOLON", "[\\n]"},
	"HEADING":      TokenType{"HEADING", "[#]{1,6}"},
	"LINE":         TokenType{"LINE", "[=|-|*]{3,}"},
	"LIST":         TokenType{"LIST", "[*|-|+]{1}"},
	"NUMBEREDLIST": TokenType{"NUMBEREDLIST", "[\\d.]"},
	"CODE":         TokenType{"CODE", "[`]"},
	"CODEBLOCK":    TokenType{"CODEBLOCK", "[`]{3}"},
	"TEXT":         TokenType{"TEXT", "\\w+"},
	"SPACE":        TokenType{"SPACE", "[\\t|\\v|\\s]"},
}
