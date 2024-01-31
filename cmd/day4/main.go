package main

import (
	"aoc_2023/pkg/file"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func prettyPrint(cards []string) {
	for _, card := range cards {
		fmt.Println(card)
	}
}

func main() {
	f := file.Open("./cmd/day4/smol")
	points := 0
	var cards []string
	for f.GetLine() {
		cards = append(cards, f.LineContent())
		points += getCardPoints(f.LineContent())
	}
	for i, card := range cards {
		matches := getCardMatches(card)
		for j := i+1; j < matches; j++ {
			fmt.Println(cards[j])
			// cards = insert(cards, j, cards[j])
		}
	}
	prettyPrint(cards)
	fmt.Println("Result Part 1 ->", points)
	f.Close()
}

func insert(cards []string, position int, newCard string) []string {
	if position == len(cards) {
		cards = append(cards, newCard)
		return cards
	}
	cards = append(cards[:position+1], cards[position:]...)
	cards[position] = newCard
	return cards
}

func getCardMatches(cardLine string) int {
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
	return counter
}

func getCardPoints(cardLine string) int {
	counter := getCardMatches(cardLine)
	points := calculatePoints(counter)
	return points
}

func numberLineToIntSlice(numberLine string) []int {
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
