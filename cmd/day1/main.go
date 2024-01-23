package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("./cmd/day1/input")
	if err != nil {
		panic(err)
	}
	finalNumber := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		finalNumber += getNumberFromLine([]rune(scanner.Text()))
	}
	fmt.Println(finalNumber)
}

func hasAnyNumberPrefix(text string) (bool, string, string) {
	prefixes := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	substitutes := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i, prefix := range prefixes {
		if strings.HasPrefix(text, prefix) {
			return true, prefix, substitutes[i]
		}
	}
	return false, "", ""
}

func findNumber(text string, c rune, i int) rune {
	textSlice := text[i:]
	if unicode.IsDigit(c) {
		return c
	}
	result, prefix, substitute := hasAnyNumberPrefix(textSlice)
	if result {
		textSlice = strings.Replace(textSlice, prefix, substitute, 1)
		return rune(textSlice[0])
	}
	return 0
}

func getNumberFromLine(text []rune) int {
	var first, last rune
	textString := string(text)
	for i, c := range textString {
		if first = findNumber(textString, c, i); first != 0 {
			break
		}
	}
	for i := len(textString) - 1; i >= 0; i-- {
		if last = findNumber(textString, text[i], i); last != 0 {
			break
		}
	}
	finalString := fmt.Sprintf("%c%c", first, last)
	result, err := strconv.Atoi(finalString)
	if err != nil {
		panic(err)
	}
	return result
}
