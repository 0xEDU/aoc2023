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

func getGameId(gameIdString string) int {
	var builder strings.Builder

	for _, c := range gameIdString {
		if unicode.IsDigit(c) {
			builder.WriteRune(c)
		}
	}
	result, _ := strconv.Atoi(builder.String())
	return result
}

func makeGameFromLine(line string) *Game {
	game := new(Game)

	splitString := strings.Split(line, ":")
	gameIdString := splitString[0]
	gameSetsSequence := splitString[1]
	game.id = getGameId(gameIdString)
	game.reds, game.greens, game.blues = getGameSetsFromSequence(gameSetsSequence)
	return game
}

func getGameSetsFromSequence(sequence string) ([]int, []int, []int) {
	var reds, greens, blues []int

	splitSequence := strings.Split(sequence, ";")
	fmt.Println(splitSequence)
	return reds, greens, blues
}

func main() {
	f := file.Open("./cmd/day2/input")
	gameList := GameList{nil}
	gameList.head = &Game{0, []int{12}, []int{13}, []int{14}, nil}
	for f.GetLine() {
		makeGameFromLine(f.LineContent())
	}
	f.Close()
}
