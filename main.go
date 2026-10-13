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
	var prefix = "http://"
	var prefixLength = len(prefix)
	var strLength = len(str)
	var buffer = []byte(str)
	var sentenceEndings = map[byte]bool{
		'.': true,
		'!': true,
		'?': true,
	}

	for i := 0; i < strLength-prefixLength; i++ {
		if str[i:i+prefixLength] == prefix {
			i += prefixLength
			for j := i; j < strLength; j++ {
				if str[j] == ' ' {
					if sentenceEndings[str[j-1]] {
						buffer[j-1] = str[j-1]
					}
					break
				}
				if j == strLength-1 && sentenceEndings[str[j]] {
					break
				}
				buffer[j] = '*'
			}
		}
	}

	return string(buffer)
}
