package processing

type TokenType struct {
	name  string
	regex string
}

var TokenTypes = map[string]TokenType{
	// "PUNCTUATION":  TokenType{"PUNCTUATION", "\\D[.]{1}"},
	"SEMICOLON":    TokenType{"SEMICOLON", "[\\n|\\v]"},
	"HEADING":      TokenType{"HEADING", "[#]{1,6}"},
	"LINE":         TokenType{"LINE", "[=|-|*]{3,}"},
	"LIST":         TokenType{"LIST", "[*|+|-]{1}[ ]{1}"},
	"NUMBEREDLIST": TokenType{"NUMBEREDLIST", "\\d[.]"},
	"CODE":         TokenType{"CODE", "[`]"},
	"CODEBLOCK":    TokenType{"CODEBLOCK", "[`]{3}"},
	"WORD":         TokenType{"WORD", "\\w+[.|,|!|?]?"},
	"SPACE":        TokenType{"SPACE", "[ ]{1,}"},
}

var SecondTokenTypes = map[string]TokenType{
	"BLOCKNUMBEREDLIST": TokenType{"BLOCKNUMBEREDLIST", "[]"},
}
