package main

import (
	"aoc_2023/pkg/file"
	"fmt"
	"unicode"
)

type Point struct {
	x, y  int
	value rune
}

type SchemeLine []Point
type Scheme [][]Point

func printScheme(scheme Scheme) {
	for _, schemeLine := range scheme {
		for _, point := range schemeLine {
			fmt.Printf("%c", point.value)
		}
		fmt.Println()
	}
}

func main() {
	scheme := loadScheme()
	result := evaluateSchemeNumbers(scheme)
	fmt.Println(result)
}

// Load scheme into memory
func loadScheme() Scheme {
	var scheme Scheme
	f := file.Open("./cmd/day3/smol")
	linePosition := 0

	for f.GetLine() {
		line := stringToSchemeLine(f.LineContent(), linePosition)
		scheme = append(scheme, line)
		linePosition++
	}
	f.Close()
	return scheme
}

func stringToSchemeLine(str string, linePosition int) SchemeLine {
	var schemeLine SchemeLine
	for i, c := range str {
		schemeLine = append(schemeLine, Point{linePosition, i, c})
	}
	return schemeLine
}

// Operate on scheme to get the final result
func evaluateSchemeNumbers(scheme Scheme) int {
	fmt.Println("BEFORE:")
	printScheme(scheme)
	fmt.Println()
	sum := 0
	for _, schemeLine := range scheme {
		for _, point := range schemeLine {
			if point.value != '.' && !unicode.IsDigit(point.value) {
				var adjacentNumbers []int
				floodFillPoint(point, &scheme, &adjacentNumbers)
				for _, number := range adjacentNumbers {
					sum += number
				}
			}
		}
	}
	fmt.Println("AFTER:")
	printScheme(scheme)
	return sum
}

func floodFillPoint(point Point, scheme *Scheme, adjacentNumbers *[]int) {
	x := point.x
	y := point.y
	value := point.value
	if value == '.' {
		return
	}
	if unicode.IsDigit(value) {
		(*scheme)[x][y].value = '.'
		return
	}
	floodFillPoint((*scheme)[x][y+1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x+1][y+1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x-1][y+1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x][y-1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x+1][y-1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x-1][y-1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x+1][y], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x+1][y+1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x+1][y-1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x-1][y], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x-1][y+1], scheme, adjacentNumbers)
	floodFillPoint((*scheme)[x-1][y-1], scheme, adjacentNumbers)
}
