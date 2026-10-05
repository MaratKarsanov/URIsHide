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
	var strLength = len(input)
	var buffer = []byte(input)
	var sentenceEndings = map[byte]bool{
		'.': true,
		'!': true,
		'?': true,
	}

	for i := 0; i < strLength; i++ {
		if isPrefix(i, input, httpPrefix) {
			i = i + 7
			for i < strLength && buffer[i] != ' ' {
				buffer[i] = '*'
				i++
			}
			if sentenceEndings[input[i-1]] {
				buffer[i-1] = input[i-1]
			}
		}
	}

	return string(buffer)
}

func isPrefix(index int, str string, prefix string) bool {
	if len(str) < index+len(prefix) {
		return false
	}

	return str[index:index+len(prefix)] == prefix
}
