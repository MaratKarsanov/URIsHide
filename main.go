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

func hideURIs(str string) string {
	var strLength = len(str)
	var result = make([]byte, 0, strLength)
	var sentenceEndings = map[byte]bool{
		'.': true,
		'!': true,
		'?': true,
		')': true,
	}

	for i := 0; i < strLength; i++ {
		if str[i] == '/' && i+2 < strLength && i > 4 && str[i-1] == ':' && str[i-2] == 'p' && str[i-3] == 't' && str[i-4] == 't' && str[i-5] == 'h' && str[i+1] == '/' {
			result = append(result, '/', '/')
			var prevChar = str[i+2]
			var starsCount = 0
			for j := i + 2; j < strLength; j++ {
				if str[j] != ' ' {
					starsCount++
					prevChar = str[j]
					if j == strLength-1 {
						if sentenceEndings[prevChar] {
							starsCount--
							i = j - 1
							break
						}
						i = j
						break
					}
					continue
				}
				if sentenceEndings[prevChar] {
					starsCount--
					i = j - 2
					break
				}
				i = j - 1
				break
			}
			for j := 0; j < starsCount; j++ {
				result = append(result, '*')
			}
			continue
		}
		result = append(result, str[i])
	}

	return string(result)
}
