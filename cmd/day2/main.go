package main

import (
	"aoc_2023/pkg/file"
	"fmt"
)

func main() {
	f := file.Open("./cmd/day2/input")
	for f.GetLine() {
		fmt.Println(f.LineContent())
	}
	f.Close()
}
