package main

import (
	"aoc_2023/pkg/file"
	"fmt"
)

func main() {
	f := file.Open("./cmd/day5/smol")
	for f.GetLine() {
		fmt.Println(f.LineContent())
	}
	f.Close()
}
