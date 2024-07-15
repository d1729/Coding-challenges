package internal

import (
	"os"
)

type Bytes struct {
}

func (bytes *Bytes) findSize(fileName string) int64 {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	return fileInfo.Size()
}
