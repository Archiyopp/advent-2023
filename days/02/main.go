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
	part1, part2 := answer()
	println("Part 1: ", part1)
	println("Part 2: ", part2)
}

var colorMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var example = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func answer() (int, int) {
	result := 0
	power := 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		game := strings.Split(line, ":")
		gameIdStr := strings.Split(game[0], " ")[1]
		gameId, err := strconv.Atoi(gameIdStr)
		if err != nil {
			errorStr := fmt.Sprintf("Game id should be a number: %v; game: %v", gameIdStr, game)
			panic(errorStr)
		}

		sets := strings.Split(game[1], ";")
		isValid := true
		minNumColorMap := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				cubeArray := strings.Split(strings.TrimSpace(cube), " ")
				numStr, color := cubeArray[0], cubeArray[1]
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic("Num should be a number of cubes")
				}
				if num > minNumColorMap[color] {
					minNumColorMap[color] = num
				}
				if colorMap[color] < num {
					isValid = false
				}
			}
		}
		if isValid {
			result += gameId
		}
		power += minNumColorMap["red"] * minNumColorMap["blue"] * minNumColorMap["green"]
	}
	return result, power
}
