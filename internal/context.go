package internal

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Context struct {
	countStrategy CountInterface
}

func (c *Context) setCountStrategy(strategy CountInterface) {
	c.countStrategy = strategy
}

func (c *Context) findSize(fileName string) int64 {
	return c.countStrategy.findSize(fileName)
}

func getBytes(context *Context, fileName string) int64 {
	countBytes := &Bytes{}
	context.setCountStrategy(countBytes)
	bytes := context.findSize(fileName)
	return bytes
}

func getLines(context *Context, fileName string) int64 {
	countLines := &Lines{}
	context.setCountStrategy(countLines)
	lines := context.findSize(fileName)
	return lines
}

func getWords(context *Context, fileName string) int64 {
	countWords := &CountWord{}
	context.setCountStrategy(countWords)
	words := context.findSize(fileName)
	return words
}

func getCharacters(context *Context, fileName string) int64 {
	countCharacters := &CountCharacter{}
	context.setCountStrategy(countCharacters)
	characters := context.findSize(fileName)
	return characters
}

func GetFileDetails(cmd *cobra.Command, args ...string) {
	context := &Context{}

	byteFileName, _ := cmd.Flags().GetString("Bytes")
	lineFileName, _ := cmd.Flags().GetString("Lines")
	wordFileName, _ := cmd.Flags().GetString("Words")
	characterFileName, _ := cmd.Flags().GetString("Characters")

	if byteFileName != "" {
		bytes := getBytes(context, byteFileName)
		fmt.Printf("%d %s\n", bytes, byteFileName)
	} else if lineFileName != "" {
		lines := getLines(context, lineFileName)
		fmt.Printf("%d %s\n", lines, lineFileName)
	} else if wordFileName != "" {
		words := getWords(context, wordFileName)
		fmt.Printf("%d %s\n", words, wordFileName)
	} else if characterFileName != "" {
		characters := getCharacters(context, characterFileName)
		fmt.Printf("%d %s\n", characters, characterFileName)
	} else {
		if len(args) > 1 {
			panic("Too many arguments")
		}
		bytes := getBytes(context, args[0])
		lines := getLines(context, args[0])
		words := getWords(context, args[0])
		fmt.Printf("%d %d %d %s\n", lines, words, bytes, args[0])
	}
}
