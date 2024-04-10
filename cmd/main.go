package main

import (
	utils "Markdown_Processor/internal/Utils"
	MarkdownToHTML "Markdown_Processor/pkg/md2html"
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

const (
	HTMLsizeMultiplier = 5
	HTMLsizeDevisor    = 4
	HTMLPrefix         = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body style="margin-left: 3vw; margin-top: 2vh;">`
	HTMLPostfix = `
	</body>
	</html>`
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
	file, err := utils.ReadFile(FileName)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	FileWithoutByte0 := make([]byte, 0, len(file))
	for i := range file {
		if file[i] != byte(0) {
			FileWithoutByte0 = append(FileWithoutByte0, file[i])
		}
	}
	HTML, err := MarkdownToHTML.Convert(string(FileWithoutByte0))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(HTML)
}
