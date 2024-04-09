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

// The function converts markdown to html
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

// StringAnalysis parses the transmitted string,
// after which it transmits the resulting html
// fragment to the channel
//
//	type HTMLLine struct {
//		SequenceNumber int
//		content        string
//	}
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

// The part of the quick sort that performs
// comparison and rearrangement of array fragments
//
//	type HTMLLine struct {
//		SequenceNumber int
//		content        string
//	}
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

// A part of the quick sort that recursively
// divides the array and sorts both parts
//
//	type HTMLLine struct {
//		SequenceNumber int
//		content        string
//	}
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

func Split(file []byte) []string {
	buf := []string{}
	start := 0
	for i := 0; i < len(file); i++ {
		if file[i] == byte(10) || i == len(file)-1 {
			buf = append(buf, string(file[start:i]))
			start = i + 1
		}
	}
	return buf
}
