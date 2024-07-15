package internal

import (
	"bufio"
	"os"
)

type CountCharacter struct{}

func (c *CountCharacter) findSize(fileName string) int64 {
	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	characters := 0
	for fileScanner.Scan() {
		characters++
	}
	e := readFile.Close()
	if e != nil {
		panic(e)
	}

	return int64(characters)
}
