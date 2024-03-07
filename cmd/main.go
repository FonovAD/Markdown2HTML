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

	FileLines, _ := processing.Split(file)
	fmt.Println(FileLines)
}
