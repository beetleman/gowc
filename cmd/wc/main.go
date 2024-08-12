package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	runeCount := 0
	for scanner.Scan() {
		wordCount++
		runeCount += utf8.RuneCount(scanner.Bytes())
	}
	fmt.Printf("Words: %d, Runes: %d\n", wordCount, runeCount)
}
