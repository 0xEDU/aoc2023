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
func evaluateSchemeNumbers(scheme Scheme) int {
	fmt.Println("BEFORE:")
	printScheme(scheme)
	fmt.Println()
	sum := 0
	for _, schemeLine := range scheme {
		for _, point := range schemeLine {
			if point.value != '.' && !unicode.IsDigit(point.value) {
				var adjacentNumbers []int
				floodFillPoint(point.x, point.y, &scheme, &adjacentNumbers)
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

func floodFillPoint(x, y int, scheme *Scheme, adjacentNumbers *[]int) {
	if x >= len(*scheme) || x < 0 || y >= len((*scheme)[0]) || y < 0 {
		return
	}
	point := (*scheme)[x][y]
	value := point.value
	if value == '.' {
		return
	}
	// Go left and right
	// (value * (10 ^ number count))
	if x == 120 && y == 35 {
		fmt.Println("dale")
	}
	if unicode.IsDigit(value) {
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
				*adjacentNumbers = append(*adjacentNumbers, farNumber*int(math.Pow(10, float64(exponent))))
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
				*adjacentNumbers = append(*adjacentNumbers, farNumber*int(math.Pow(10, float64(k+l-1))))
				(*scheme)[x][y-l].value = '.'
			}
			l++
		}
		fmt.Println(adjacentNumbers)
		return
	}
	fmt.Printf("Evaluating position (%d,%d) with value '%s'\n", x+1, y+1, string(value))
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
