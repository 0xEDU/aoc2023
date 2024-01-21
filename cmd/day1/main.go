package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func getNumberFromLine(text []rune) int {
	var first rune
	var last rune
	for _, c := range text {
		if unicode.IsDigit(c) {
			first = c
			break
		}
	}
	for i := len(text) - 1; i >= 0; i-- {
		if unicode.IsDigit(text[i]) {
			last = text[i]
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
