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
	second := part2()
	println("Part 2:", second)
}

func part1() int {
	var result = 0
	for _, line := range strings.Split(input, "\n") {
		row := [2]int{}
		isFirstNum := true
		for _, char := range strings.TrimSpace(line) {
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

var numMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func part2() int {
	var result = 0
	for _, line := range strings.Split(
		input, "\n") {
		row := [2]int{}
		isFirstNum := true
		strNum := ""
		for _, char := range strings.TrimSpace(line) {
			strNum += strings.ToLower(string(char))
			num := 0
			for key, value := range numMap {
				if strings.Contains(strNum, key) {
					num = value
					strNum = ""
				}
			}

			value, err := strconv.Atoi(string(char))
			if err == nil {
				num = value
			}
			if num == 0 {
				continue
			}
			strNum = ""
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
