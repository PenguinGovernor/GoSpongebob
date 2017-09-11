package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

var (
	textToConvert string
	verboseMode   bool
)

func textToSponge(s string) string {
	letterIndex := 1
	stringArr := []rune(strings.ToUpper(s))

	for i, val := range stringArr {
		if unicode.IsLetter(val) {
			if letterIndex%2 == 0 {
				stringArr[i] = unicode.ToLower(val)
			}
			letterIndex++
		}
	}

	return string(stringArr)

}

func printSponge(s string) {
	fmt.Println(textToSponge(s))
}

func main() {
	preText, _ := ioutil.ReadAll(os.Stdin)
	text := string(preText)
	printSponge(text)
}
