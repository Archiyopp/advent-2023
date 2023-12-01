package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	first := part1()
	println("Part 1:", first)
}

func part1() int {
	var result = 0
	for _, line := range strings.Split(input, "\n") {
		row := [2]int{}
		isFirstNum := true
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			if isFirstNum {
				row[0] = num
				isFirstNum = false
			}
			row[1] = num

		}
		result += (row[0] * 10) + row[1]
	}
	return result
}
