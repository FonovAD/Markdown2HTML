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
	FileWithoutByte0 := []byte{}
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
	lex.LexAnalusis()
	HTML := `<!DOCTYPE html>
	<html lang="en">
	<head style="margin: 100vw">
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>`
	parser := processing.Parser{Tokens: lex.TokenList, Pos: 0}
	root := parser.NewParseCode()
	HTML += processing.Run(root)
	HTML += `
	</body>
	</html>`
	fmt.Println(HTML)
}
