package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println(hideURIs(scanner.Text()))
	}
}

func hideURIs(input string) string {
	const httpPrefix = "http://"
	var inputRunes = []rune(input)
	var buffer = []rune(input)
	var strLength = len(buffer)
	var sentenceEndings = map[rune]bool{
		'.': true,
		'!': true,
		'?': true,
	}

	for i := 0; i < strLength; i++ {
		if isPrefix(i, inputRunes, httpPrefix) {
			i = i + 7
			for i < strLength && buffer[i] != ' ' {
				buffer[i] = '*'
				i++
			}
			if sentenceEndings[inputRunes[i-1]] {
				buffer[i-1] = inputRunes[i-1]
			}
		}
	}

	return string(buffer)
}

func isPrefix(index int, runes []rune, prefix string) bool {
	if len(runes) < index+len(prefix) {
		return false
	}

	return string(runes[index:index+len(prefix)]) == prefix
}
