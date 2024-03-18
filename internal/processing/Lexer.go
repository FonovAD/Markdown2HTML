package processing

import (
	"regexp"
)

type Lexer struct {
	Code      string
	Pos       int
	TokenList []Token
}

func (L *Lexer) LexAnalusis() []Token {
	count := 0
	for L.NextToken() {
		if count > len(L.Code)*5 {
			panic("Не могу разобрать слова")
		} else {
			count += 1
		}
	}
	return L.TokenList
}

func (L *Lexer) NextToken() bool {
	if L.Pos >= len(L.Code) {
		return false
	}
	var tokenTypesValues []string
	for key := range TokenTypes {
		tokenTypesValues = append(tokenTypesValues, key)
	}
	for i := 0; i < len(tokenTypesValues)-1; i++ {
		tokenType := TokenTypes[tokenTypesValues[i]]
		r, _ := regexp.Compile("^" + tokenType.regex)
		result := r.FindString(L.Code[L.Pos:])
		found := r.MatchString(L.Code[L.Pos:])
		resultIndex := r.FindStringIndex(L.Code[L.Pos:])
		if found {
			token := Token{tokenType, result, resultIndex[1]}
			L.Pos = L.Pos + len(result)
			L.TokenList = append(L.TokenList, token)
			return true
		}
	}
	return true
}
