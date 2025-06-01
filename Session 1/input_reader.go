package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Single word input using fmt.Scanln
	var wordInput string
	fmt.Print("Enter a word: ")
	fmt.Scanln(&wordInput)

	word := strings.Fields(wordInput)
	fmt.Println("Word (using Scanln):", word)

	// Full sentence input using bufio.NewReader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	sentenceInput, _ := reader.ReadString('\n')
	sentenceInput = strings.TrimSpace(sentenceInput)

	words := strings.Fields(sentenceInput)
	fmt.Println("Words (using bufio.Reader):", words)
}
