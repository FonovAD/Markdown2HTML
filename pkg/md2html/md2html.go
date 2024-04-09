package MarkdownToHTML

import (
	MarkdownToHTML "Markdown_Processor/pkg/md2html/processing"
	"bytes"
	"fmt"
	"strings"
	"sync"
)

type HTMLLine struct {
	SequenceNumber int
	content        string
}

func Convert(MdText string) (string, error) {
	Code := strings.Split(string(MdText), "\n")
	var HTML bytes.Buffer
	HTML.WriteString(HTMLPrefix)
	var wg sync.WaitGroup
	htmlLines := make([]HTMLLine, 0, len(Code))
	c := make(chan HTMLLine, 16)
	for i, j := range Code {
		wg.Add(1)
		go StringAnalysis(&wg, c, j, i+1) // Добавить обработку ошибок
		htmlLine := <-c
		htmlLines = append(htmlLines, htmlLine)
	}
	htmlLines = QuickSort(htmlLines)
	for _, j := range htmlLines {
		HTML.WriteString(j.content)
	}
	HTML.WriteString(HTMLPostfix)
	return HTML.String(), nil
}

func StringAnalysis(wg *sync.WaitGroup, c chan HTMLLine, line string, lineNumber int) {
	htmlLine := HTMLLine{SequenceNumber: lineNumber}

	lex := MarkdownToHTML.Lexer{
		Code:      string(line),
		Pos:       0,
		TokenList: []MarkdownToHTML.Token{},
	}
	if err := lex.LexAnalusis(); err != nil {
		fmt.Println(err)
	}
	parser := MarkdownToHTML.Parser{Tokens: lex.TokenList, Pos: 0}
	root := parser.NewParseCode()
	HTMLsize := (len(line) * HTMLsizeMultiplier) / HTMLsizeDevisor
	htmlLine.content = MarkdownToHTML.Run(root, HTMLsize)
	c <- htmlLine
	wg.Done()
}

func QSpart1(arr []HTMLLine, low, high int) ([]HTMLLine, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j].SequenceNumber < pivot.SequenceNumber {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QSpart2(arr []HTMLLine, low, high int) []HTMLLine {
	if low < high {
		var p int
		arr, p = QSpart1(arr, low, high)
		arr = QSpart2(arr, low, p-1)
		arr = QSpart2(arr, p+1, high)
	}
	return arr
}

func QuickSort(arr []HTMLLine) []HTMLLine {
	return QSpart2(arr, 0, len(arr)-1)
}
