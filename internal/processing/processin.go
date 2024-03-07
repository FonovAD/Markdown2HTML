package processing

import (
	"errors"
	"io"
	"os"
)

func ReadFile(FileName string) ([]byte, error) {
	MarkdownFile, err := os.Open(FileName)
	if err != nil {
		return []byte{0}, errors.New("error opening file")
	}
	defer MarkdownFile.Close()
	FileData := make([]byte, 64)
	for {
		_, err := MarkdownFile.Read(FileData)
		if err == io.EOF {
			break
		}
	}
	return FileData, nil
}

func Split(file []byte) ([]string, error) {
	buf := []string{}
	start := 0
	for i := 0; i < len(file); i++ {
		if file[i] == byte(10) {
			buf = append(buf, string(file[start:i]))
			start = i + 1
		}
	}
	return buf, nil
}
