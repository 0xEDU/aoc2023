package main

import (
	"aoc_2023/pkg/file"
	"fmt"
	"math"
	"strconv"
	"unicode"
)

type Point struct {
	x, y  int
	value rune
}

type SchemeLine []Point
type Scheme [][]Point

func main() {
	scheme := loadScheme()
	result1, result2 := evaluateSchemeNumbers(scheme)
	fmt.Println("Part 1 ->", result1)
	fmt.Println("Part 2 ->", result2)
}

// Load scheme into memory
func loadScheme() Scheme {
	var scheme Scheme
	f := file.Open("./cmd/day3/input")
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
func evaluateSchemeNumbers(scheme Scheme) (int, int) {
	sum := 0
	ratio := 0
	for _, schemeLine := range scheme {
		for _, point := range schemeLine {
			if point.value != '.' && !unicode.IsDigit(point.value) {
				var adjacentNumbers []int
				floodFillPoint(point.x, point.y, &scheme, &adjacentNumbers)
				for _, number := range adjacentNumbers {
					sum += number
				}
				if len(adjacentNumbers) == 2 {
					ratio += adjacentNumbers[0] * adjacentNumbers[1]
				}
			}
		}
	}
	return sum, ratio
}

func floodFillPoint(x, y int, scheme *Scheme, adjacentNumbers *[]int) {
	if x >= len(*scheme) || x < 0 || y >= len((*scheme)[0]) || y < 0 {
		return
	}
	point := (*scheme)[x][y]
	value := point.value
	if value == '.' {
		return
	}
	// Go right, then left
	// (value * (10 ^ exponent))
	if unicode.IsDigit(value) {
		var tempSlice []int
		k := 0
		for {
			if y+k >= len((*scheme)[0]) {
				break
			}
			if !unicode.IsDigit((*scheme)[x][y+k].value) {
				break
			}
			if unicode.IsDigit((*scheme)[x][y+k].value) {
				j := 0
			INNER:
				for {
					if y+k+j >= len((*scheme)[0]) {
						break INNER
					}
					if !unicode.IsDigit((*scheme)[x][y+k+j].value) {
						break INNER
					}
					j++
				}
				exponent := j - 1
				if j == 0 {
					exponent = j
				}
				farNumber, _ := strconv.Atoi(string((*scheme)[x][y+k].value))
				tempSlice = append(tempSlice, farNumber*int(math.Pow(10, float64(exponent))))
				(*scheme)[x][y+k].value = '.'
			}
			k++
		}
		l := 1
		for {
			if y-l < 0 {
				break
			}
			if !unicode.IsDigit((*scheme)[x][y-l].value) {
				break
			}
			if unicode.IsDigit((*scheme)[x][y-l].value) {
				farNumber, _ := strconv.Atoi(string((*scheme)[x][y-l].value))
				tempSlice = append(tempSlice, farNumber*int(math.Pow(10, float64(k+l-1))))
				(*scheme)[x][y-l].value = '.'
			}
			l++
		}
		adjacentNumber := 0
		for _, number := range tempSlice {
			adjacentNumber += number
		}
		*adjacentNumbers = append(*adjacentNumbers, adjacentNumber)
		return
	}
	floodFillPoint(x, y+1, scheme, adjacentNumbers)
	floodFillPoint(x+1, y+1, scheme, adjacentNumbers)
	floodFillPoint(x-1, y+1, scheme, adjacentNumbers)
	floodFillPoint(x, y-1, scheme, adjacentNumbers)
	floodFillPoint(x+1, y-1, scheme, adjacentNumbers)
	floodFillPoint(x-1, y-1, scheme, adjacentNumbers)
	floodFillPoint(x+1, y, scheme, adjacentNumbers)
	floodFillPoint(x+1, y+1, scheme, adjacentNumbers)
	floodFillPoint(x+1, y-1, scheme, adjacentNumbers)
	floodFillPoint(x-1, y, scheme, adjacentNumbers)
	floodFillPoint(x-1, y+1, scheme, adjacentNumbers)
	floodFillPoint(x-1, y-1, scheme, adjacentNumbers)
}
