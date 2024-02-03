package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func ExtractNumberFromString(str string) int {
	var builder strings.Builder

	for _, c := range str {
		if unicode.IsDigit(c) {
			builder.WriteRune(c)
		}
	}
	result, _ := strconv.Atoi(builder.String())
	return result
}

func NumberLineToIntSlice(numberLine string) []int {
	var numbers []int

	numberLine = strings.Trim(numberLine, " ")
	splitNumberLine := strings.Split(numberLine, " ")
	for _, numberString := range splitNumberLine {
		if numberString == "" {
			continue
		}
		number, _ := strconv.Atoi(numberString)
		numbers = append(numbers, number)
	}
	return numbers
}
