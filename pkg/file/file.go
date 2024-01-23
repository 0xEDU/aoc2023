package file

import (
	"bufio"
	"os"
)

type File struct {
	filePtr *os.File
	scanner *bufio.Scanner
}

func Open(filePath string) File {
	f := File{}
	filePtr, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	f.filePtr = filePtr
	f.scanner = bufio.NewScanner(filePtr)
	return f
}

func (f File) GetLine() bool {
	return f.scanner.Scan()
}

func (f File) LineContent() string {
	return f.scanner.Text()
}

func (f File) Close() {
	if err := f.filePtr.Close(); err != nil {
		panic(err)
	}
}
