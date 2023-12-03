package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	first := part1()
	println("First part: ", first)
}

var example = `467..114..
...*......
..35..633.
.........*
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

type PartNumber struct {
	x   int
	y   int
	num int
}

func part1() int {
	symbolsMap := map[string](bool){}
	parts := []PartNumber{}
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		strNum := ""
		numLastX, numLastY := 0, 0
		for x, char := range line {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				if num, err := strconv.Atoi(strNum); err == nil {
					parts = append(parts, PartNumber{numLastX, numLastY, num})
					strNum = ""
				}
				if char != '.' {
					symbolsMap[getCoordinate(x, y)] = true
				}
			} else {
				strNum += string(char)
				numLastX, numLastY = x, y
			}
		}
		// check if there is a num at the end of the line
		if num, err := strconv.Atoi(strNum); err == nil {
			parts = append(parts, PartNumber{numLastX, numLastY, num})
		}
	}
	result := 0
	for _, part := range parts {
		if checkCoordinates(part, symbolsMap) {
			result += part.num
		}
	}
	return result
}

func checkCoordinates(part PartNumber, symbolsMap map[string](bool)) bool {
	numLength := len(fmt.Sprint(part.num))
	for i := -1; i <= numLength; i++ {
		for j := -1; j < 2; j++ {
			coordinate := getCoordinate(part.x-i, part.y+j)
			if symbolsMap[coordinate] {
				return true
			}
		}
	}
	return false
}

func getCoordinate(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
