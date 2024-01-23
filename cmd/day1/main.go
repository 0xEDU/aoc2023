package main

import (
	"aoc_2023/pkg/file"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f := file.File{}
	f.Open("./cmd/day1/input")
	finalNumberPart1 := 0
	finalNumberPart2 := 0
	for f.GetLine() {
		finalNumberPart1 += getNumberFromLine([]rune(f.LineContent()), false)
		finalNumberPart2 += getNumberFromLine([]rune(f.LineContent()), true)
	}
	fmt.Println("Part 1 ->", finalNumberPart1)
	fmt.Println("Part 2 ->", finalNumberPart2)
	f.Close()
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

func findNumber(text string, c rune, i int, wordsAreNumbers bool) rune {
	if unicode.IsDigit(c) {
		return c
	}
	if wordsAreNumbers {
		textSlice := text[i:]
		result, prefix, substitute := hasAnyNumberPrefix(textSlice)
		if result {
			textSlice = strings.Replace(textSlice, prefix, substitute, 1)
			return rune(textSlice[0])
		}
	}
	return 0
}

func getNumberFromLine(text []rune, wordsAreNumbers bool) int {
	var first, last rune
	textString := string(text)
	for i, c := range textString {
		if first = findNumber(textString, c, i, wordsAreNumbers); first != 0 {
			break
		}
	}
	for i := len(textString) - 1; i >= 0; i-- {
		if last = findNumber(textString, text[i], i, wordsAreNumbers); last != 0 {
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
