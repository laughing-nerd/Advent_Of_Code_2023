package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var max_red int = 12
var max_green int = 13
var max_blue int = 14

func part1(games []string) int {
	for _, game := range games {
		cubes := strings.Split(strings.Trim(game, " "), ",")
		for _, cube := range cubes {
			count, _ := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Trim(cube, " "), " ")[0]))
			color := strings.TrimSpace(strings.Split(strings.Trim(cube, " "), " ")[1])

			if color == "red" && count > max_red {
				return 0
			}
			if color == "green" && count > max_green {
				return 0
			}
			if color == "blue" && count > max_blue {
				return 0
			}
		}
	}
	return 1
}

func part2(games []string) int {
	red := 0
	green := 0
	blue := 0
	for _, game := range games {
		cubes := strings.Split(strings.Trim(game, " "), ",")

		for _, cube := range cubes {
			count, _ := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Trim(cube, " "), " ")[0]))
			color := strings.TrimSpace(strings.Split(strings.Trim(cube, " "), " ")[1])

			if color == "red" && count > red {
				red = count
			}
			if color == "green" && count > green {
				green = count
			}
			if color == "blue" && count > blue {
				blue = count
			}
		}
	}
	return (red * green * blue)
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("Input file could not be opened")
	}
	lines := bufio.NewScanner(file)

	sum1 := 0
	sum2 := 0

	for lines.Scan() {
		line := strings.Trim(lines.Text(), " ")

		split := strings.Split(line, ":")
		gameId, _ := strconv.Atoi(strings.Split(split[0], " ")[1])

		games := strings.Split(strings.TrimSpace(strings.Trim(split[1], " ")), ";")

		// Part 1
		possible := part1(games)
		if possible == 1 {
			sum1 += gameId
		}

		//Part 2
		sum2 += part2(games)
	}
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}

