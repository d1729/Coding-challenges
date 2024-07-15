package internal

import (
	"bufio"
	"os"
)

type Lines struct {
}

func (l *Lines) findSize(filename string) int64 {
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := 0
	for fileScanner.Scan() {
		lines++
	}

	e := readFile.Close()
	if e != nil {
		panic(e)
	}
	return int64(lines)
}
