package internal

import (
	"bufio"
	"os"
)

type CountWord struct{}

func (countWord *CountWord) findSize(fileName string) int64 {
	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanWords)

	words := 0
	for fileScanner.Scan() {
		words++
	}

	e := readFile.Close()
	if e != nil {
		panic(e)
	}

	return int64(words)
}
