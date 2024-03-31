package main

import (
	"Markdown_Processor/internal/processing"
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	var FileName string
	app := cli.NewApp()
	app.Usage = "Converts Markdown to HTML"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "The name of the Markdown file",
		},
	}
	app.Action = func(c *cli.Context) error {
		FileName = c.GlobalString("file")
		return nil
	}
	app.Run(os.Args)
	file, err := processing.ReadFile(FileName)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	FileWithoutByte0 := make([]byte, len(file))
	for i := range file {
		if file[i] != byte(0) {
			FileWithoutByte0 = append(FileWithoutByte0, file[i])
		}
	}
	lex := processing.Lexer{
		Code:      string(FileWithoutByte0),
		Pos:       0,
		TokenList: []processing.Token{},
	}
	if err := lex.LexAnalusis(); err != nil {
		fmt.Println("Beda")
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
	fmt.Println(HTML)
}
