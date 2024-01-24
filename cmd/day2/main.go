package main

import (
	"aoc_2023/pkg/file"
	"fmt"
	"sort"
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
	sort.Sort(sort.Reverse(sort.IntSlice(reds)))
	sort.Sort(sort.Reverse(sort.IntSlice(greens)))
	sort.Sort(sort.Reverse(sort.IntSlice(blues)))
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

func sliceHasLargerValue(colorSlice []int, value int) bool {
	for _, sliceValue := range colorSlice {
		if sliceValue > value {
			return true
		}
	}
	return false
}

func getSumOfValidGames(list *GameList) int {
	sumOfValidGames := 0
	temp := (*list).head
	validValues := []int{temp.reds[0], temp.greens[0], temp.blues[0]}
	temp = temp.next
	for temp != nil {
		if sliceHasLargerValue(temp.reds, validValues[0]) ||
			sliceHasLargerValue(temp.greens, validValues[1]) ||
			sliceHasLargerValue(temp.blues, validValues[2]) {
			temp = temp.next
			continue
		}
		sumOfValidGames += temp.id
		temp = temp.next
	}
	return sumOfValidGames
}

func getSumOfMinimumGamesPower(list *GameList) int {
	sumOfMinimumGamesPower := 0
	temp := (*list).head.next
	for temp != nil {
		power := temp.reds[0] * temp.greens[0] * temp.blues[0]
		sumOfMinimumGamesPower += power
		temp = temp.next
	}
	return sumOfMinimumGamesPower
}

func main() {
	f := file.Open("./cmd/day2/input")
	gameList := GameList{nil}
	gameList.head = &Game{0, []int{12}, []int{13}, []int{14}, nil}
	for f.GetLine() {
		game := makeGameFromLine(f.LineContent())
		appendGameToList(&gameList.head, &game)
	}
	sumOfValidGames := getSumOfValidGames(&gameList)
	fmt.Println("Part 1 ->", sumOfValidGames)
	sumOfMinimumGamesPower := getSumOfMinimumGamesPower(&gameList)
	fmt.Println("Part 2 ->", sumOfMinimumGamesPower)
	f.Close()
}
