package main

import (
	"aoc_2023/pkg/file"
	"aoc_2023/pkg/utils"
	"fmt"
	"strings"
)

func getParserEnum() []string {
	return []string{
		"seeds",
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}
}

type Seed struct {
	seedNum     int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type Seeds []Seed

func main() {
	var seeds Seeds
	f := file.Open("./cmd/day5/smol")
	i := 0
	for f.GetLine() {
		line := f.LineContent()
		if line == "" {
			i++
		}
		switch i {
		case 0:
			seeds = getSeeds(line)
		case 1:
			if strings.Contains(line, getParserEnum()[i]) || line == "" {
				continue
			}
			mapFields(&seeds, line, "soil", "seedNum")
		case 2:
			if strings.Contains(line, getParserEnum()[i]) || line == "" {
				fillEmptyField(&seeds, "soil", "seedNum")
				continue
			}
			mapFields(&seeds, line, "fertilizer", "soil")
		case 3:
			if strings.Contains(line, getParserEnum()[i]) || line == "" {
				fillEmptyField(&seeds, "fertilizer", "soil")
				continue
			}
			mapFields(&seeds, line, "water", "fertilizer")
		case 4:
			if strings.Contains(line, getParserEnum()[i]) || line == "" {
				fillEmptyField(&seeds, "water", "fertilizer")
				continue
			}
			mapFields(&seeds, line, "light", "water")
		case 5:
			if strings.Contains(line, getParserEnum()[i]) || line == "" {
				fillEmptyField(&seeds, "light", "water")
				continue
			}
			mapFields(&seeds, line, "temperature", "light")
		case 6:
			if strings.Contains(line, getParserEnum()[i]) || line == "" {
				fillEmptyField(&seeds, "temperature", "light")
				continue
			}
			mapFields(&seeds, line, "humidity", "temperature")
		case 7:
			if strings.Contains(line, getParserEnum()[i]) || line == "" {
				fillEmptyField(&seeds, "humidity", "temperature")
				continue
			}
			mapFields(&seeds, line, "location", "humidity")
		}
	}
	f.Close()
	fillEmptyField(&seeds, "location", "humidity")
	smolNum := seeds[0].location
	for _, seed := range seeds {
		if seed.location < smolNum {
			smolNum = seed.location
		}
	}
	fmt.Println("Part 1 ->", smolNum)
}

func getFields(seeds *Seeds, position int) map[string]*int {
	fields := map[string]*int{
		"seedNum":     &(*seeds)[position].seedNum,
		"soil":        &(*seeds)[position].soil,
		"fertilizer":  &(*seeds)[position].fertilizer,
		"water":       &(*seeds)[position].water,
		"light":       &(*seeds)[position].light,
		"temperature": &(*seeds)[position].temperature,
		"humidity":    &(*seeds)[position].humidity,
		"location":    &(*seeds)[position].location,
	}
	return fields
}

func mapFields(seeds *Seeds, line string, dstField, srcField string) {
	nums := utils.NumberLineToIntSlice(line)
	expandedRange := expandRange(nums)
	for j := 0; j < len(*seeds); j++ {
		fields := getFields(seeds, j)
		contain, position := contains(expandedRange[1], *fields[srcField])
		if contain {
			*fields[dstField] = expandedRange[0][position]
		}
	}
}

func fillEmptyField(seeds *Seeds, emptyField, dstField string) {
	for i := 0; i < len(*seeds); i++ {
		fields := getFields(seeds, i)
		if *fields[emptyField] == 0 {
			*fields[emptyField] = *fields[dstField]
		}
	}
}

func prettyPrint(seeds Seeds) {
	for i, seed := range seeds {
		fmt.Printf("Seed %d:", i+1)
		fmt.Println(seed)
	}
}

func contains(s []int, e int) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
}

func expandRange(nums []int) [][]int {
	var expandedRange [][]int
	rangeLen := nums[2]
	var dstExpanded []int
	for i := 0; i < rangeLen; i++ {
		dstExpanded = append(dstExpanded, nums[0]+i)
	}
	var srcExpanded []int
	for i := 0; i < rangeLen; i++ {
		srcExpanded = append(srcExpanded, nums[1]+i)
	}
	expandedRange = append(expandedRange, dstExpanded)
	expandedRange = append(expandedRange, srcExpanded)
	return expandedRange
}

func getSeeds(line string) Seeds {
	var seeds Seeds
	splitLine := strings.Split(line, ":")
	seedsString := strings.Trim(splitLine[1], " ")
	seedIds := utils.NumberLineToIntSlice(seedsString)
	for _, seedId := range seedIds {
		seeds = append(seeds, Seed{
			seedId, 0, 0, 0, 0, 0, 0, 0})
	}
	return seeds
}
