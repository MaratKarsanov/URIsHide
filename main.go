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
		fmt.Println(hideURIsFast(scanner.Text()))
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

func hideURIsFast(str string) string {
	var strLength = len(str)
	var result = []byte(str)
	var sentenceEndings = map[byte]bool{
		'.': true,
		'!': true,
		'?': true,
	}

	for i := 5; i < strLength-1; i++ {
		if str[i] == '/' && str[i-1] == ':' && str[i-2] == 'p' && str[i-3] == 't' && str[i-4] == 't' && str[i-5] == 'h' && str[i+1] == '/' {
			i += 2
			for j := i; j < strLength; j++ {
				if str[j] == ' ' {
					if sentenceEndings[str[j-1]] {
						result[j-1] = str[j-1]
					}
					break
				}
				if j == strLength-1 && sentenceEndings[str[j]] {
					break
				}
				result[j] = '*'
			}
		}
	}

	return string(result)
}
