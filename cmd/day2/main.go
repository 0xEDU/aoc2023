package main

import (
	"aoc_2023/pkg/file"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Game struct {
	id     int
	reds   []int
	greens []int
	blues  []int
	next   *Game
}

type GameList struct {
	head *Game
}

func extractNumberFromString(str string) int {
	var builder strings.Builder

	for _, c := range str {
		if unicode.IsDigit(c) {
			builder.WriteRune(c)
		}
	}
	result, _ := strconv.Atoi(builder.String())
	return result
}

func getGameSetsFromSequence(sequence string) ([]int, []int, []int) {
	var reds, greens, blues []int

	splitSequence := strings.Split(sequence, ";")
	fmt.Println(splitSequence)
	for _, sequence := range splitSequence {
		colors := strings.Split(sequence, ",")
		for _, color := range colors {
			if strings.Contains(color, "red") {
				reds = append(reds, extractNumberFromString(color))
			}
			if strings.Contains(color, "green") {
				greens = append(greens, extractNumberFromString(color))
			}
			if strings.Contains(color, "blue") {
				blues = append(blues, extractNumberFromString(color))
			}
		}
	}
	fmt.Println("Reds:", reds, "Greens:", greens, "Blues:", blues)
	return reds, greens, blues
}

func makeGameFromLine(line string) *Game {
	game := new(Game)

	splitString := strings.Split(line, ":")
	gameIdString := splitString[0]
	gameSetsSequence := splitString[1]
	game.id = extractNumberFromString(gameIdString)
	game.reds, game.greens, game.blues = getGameSetsFromSequence(gameSetsSequence)
	game.next = nil
	return game
}

func appendGameToList(head **Game, game **Game) {
	temp := *head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = *game
}

func main() {
	f := file.Open("./cmd/day2/input")
	gameList := GameList{nil}
	gameList.head = &Game{0, []int{12}, []int{13}, []int{14}, nil}
	for f.GetLine() {
		game := makeGameFromLine(f.LineContent())
		appendGameToList(&gameList.head, &game)
	}
	f.Close()
}
