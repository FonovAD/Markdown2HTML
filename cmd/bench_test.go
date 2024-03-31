package main

import (
	"Markdown_Processor/internal/processing"
	"fmt"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	errors := 0
	TotalOperations := 0
	for i := 0; i < b.N; i++ {
		TotalOperations += 1
		FileWithoutByte0 := `# User interface
## The main page
#### The functionality of the main page
1. Book Recommendation 
2. Collections of books

#### What will be located on the main page
1. the first user will be greeted by a book recommendation. He can click on it and go to the page with the book
2. Scrolling through below, the user will see a selection of books
3. In the upper left corner there will be a circle with an avatar, when clicking on which the user will be able to go to the profile
4. In the center of the top there will be a search bar where you can find a specific book

===
+ Book Recomm **endation**
+ Collections of *books skoob*

Lorem ipsum dolor sit amet consectetur adipisicing elit. Officiis dolor assumenda ut consectetur cum maiores explicabo delectus soluta veritatis maxime, repellat earum autem! Libero ut, ad odio cupiditate porro pariatur?

@Pepsi_King %
`
		lex := processing.Lexer{
			Code:      string(FileWithoutByte0),
			Pos:       0,
			TokenList: []processing.Token{},
		}
		if err := lex.LexAnalusis(); err != nil {
			errors += 1
		}
		HTML := `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body style="margin-left: 3vw; margin-top: 2vh;">`
		parser := processing.Parser{Tokens: lex.TokenList, Pos: 0}
		root := parser.NewParseCode()
		HTML += processing.Run(root)
		HTML += `
		</body>
		</html>`
	}
	fmt.Printf("\nTotat op: %d, \t errors %d\n", TotalOperations, errors)
}
