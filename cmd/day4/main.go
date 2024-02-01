package main

import (
	"aoc_2023/pkg/file"
	"aoc_2023/pkg/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func prettyPrint(cards CardDeck) {
	for i, cardLine := range cards {
		fmt.Printf("Line %d [", i)
		for _, card := range cardLine {
			fmt.Printf("Card %d: %d matches | ", card.id, card.matches)
		}
		fmt.Printf("]")
		fmt.Println()
	}
}

type Card struct {
	id      int
	matches int
}

type CardLine []Card
type CardDeck []CardLine

func main() {
	f := file.Open("./cmd/day4/smol")
	points := 0
	var cards CardDeck
	for f.GetLine() {
		var cardLine CardLine
		card := createCard(f.LineContent())
		cardLine = append(cardLine, card)
		cards = append(cards, cardLine)
		points += getCardPoints(f.LineContent())
	}
	for i, cardLine := range cards {
		for _, card := range cardLine {
			for j := 0; j < card.matches; j++ {
				cards[i+1] = append(cards[i+1], cards[j][0]) // Need to figure this out
			}
		}
	}
	prettyPrint(cards)
	fmt.Println("Result Part 1 ->", points)
	f.Close()
}

func createCard(cardString string) Card {
	card := Card{id: getCardId(cardString), matches: getCardMatches(cardString)}
	return card
}

func getCardId(cardLine string) int {
	splitCardLine := strings.Split(cardLine, ":")
	id := utils.ExtractNumberFromString(splitCardLine[0])
	return id
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
