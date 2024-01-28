package main

import (
	"aoc_2023/pkg/file"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	f := file.Open("./cmd/day4/input")
	points := 0
	for f.GetLine() {
		points += getCardPoints(f.LineContent())
	}
	fmt.Println("Result ->", points)
	f.Close()
}

func getCardPoints(cardLine string) int {
	splitCardLine := strings.Split(cardLine, ":")
	numbers := splitCardLine[1]
	splitNumbers := strings.Split(numbers, "|")
	winningNumbers := numberLineToIntSlice(splitNumbers[0])
	foundNumbers := numberLineToIntSlice(splitNumbers[1])
	counter := 0
	for _, winningNumber := range winningNumbers {
		for _, foundNumber := range foundNumbers {
			if winningNumber == foundNumber {
				counter++
			}
		}
	}
	points := calculatePoints(counter)
	return points
}

func numberLineToIntSlice(numberLine string) []int {
	var numbers []int

	numberLine = strings.Trim(numberLine, " ")
	splitNumberLine := strings.Split(numberLine, " ")
	fmt.Println(numberLine)
	for _, numberString := range splitNumberLine {
		if numberString == "" {
			continue
		}
		number, _ := strconv.Atoi(numberString)
		numbers = append(numbers, number)
	}
	fmt.Println(numbers)
	return numbers
}

func calculatePoints(counter int) int {
	if counter == 0 {
		return 0
	}
	if counter == 1 {
		return 1
	}
	counter--
	return int(math.Pow(2, float64(counter)))
}
