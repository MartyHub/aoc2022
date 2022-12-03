package aoc2022

import (
	"bufio"
	"log"
	"os"
)

type LineReader struct {
	file    *os.File
	count   *int
	scanner *bufio.Scanner
}

func NewLineReader(fileName string) LineReader {
	log.Printf("Opening %v...", fileName)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return LineReader{
		file:    file,
		count:   new(int),
		scanner: bufio.NewScanner(file),
	}
}

func (lr LineReader) Close() error {
	return lr.file.Close()
}

func (lr LineReader) HasNext() bool {
	result := lr.scanner.Scan()

	if result {
		*lr.count += 1
	}

	return result
}

func (lr LineReader) Count() int {
	return *lr.count
}

func (lr LineReader) Text() string {
	return lr.scanner.Text()
}
