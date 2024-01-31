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
