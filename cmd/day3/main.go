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
	// Go left and right
	// (value * (10 ^ number count))
	if unicode.IsDigit(value) {
		numberValue, _ := strconv.Atoi(string(value))
		k := 0
		for {
			if y+k > len((*scheme)[0]) {
				break
			}
			if unicode.IsDigit((*scheme)[x][y+k].value) {
				exponent := k - 1
				if k == 0 {
					exponent = k - 1
				}
				farNumber, _ := strconv.Atoi(string((*scheme)[x][y+k].value))
				*adjacentNumbers = append(*adjacentNumbers, farNumber*int(math.Pow(10, float64(exponent))))
			}
			if !unicode.IsDigit((*scheme)[x][y+k].value) {
				(*scheme)[x][y].value = '.'
				break
			}
			k++
		}
		l := 1
		for {
			if y-l < 0 {
				break
			}
			if unicode.IsDigit((*scheme)[x][y-l].value) {
				farNumber, _ := strconv.Atoi(string((*scheme)[x][y-l].value))
				*adjacentNumbers = append(*adjacentNumbers, farNumber*int(math.Pow(10, float64(k+l-1))))
				if !unicode.IsDigit((*scheme)[x][y-l].value) {
					break
				}
				(*scheme)[x][y-l].value = '.'
			}
			l++
		}
		*adjacentNumbers = append(*adjacentNumbers, numberValue*int(math.Pow(10, float64(k-1))))
		fmt.Println("POSITION", x, y)
		fmt.Println(adjacentNumbers)
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
