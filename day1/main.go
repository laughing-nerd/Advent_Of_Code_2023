package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func extractDigit(v string) int {
	switch v {
	case "one", "1":
		return 1
	case "two", "2":
		return 2
	case "three", "3":
		return 3
	case "four", "4":
		return 4
	case "five", "5":
		return 5
	case "six", "6":
		return 6
	case "seven", "7":
		return 7
	case "eight", "8":
		return 8
	case "nine", "9":
		return 9
	default:
		return 0
	}
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("Input file could not be opened")
	}
	defer file.Close()
	lines := bufio.NewScanner(file)
	sum := 0
	toCheckFor := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	for lines.Scan() {
		line := lines.Text()

		//Check for the first index of a number
		left := 0
		right := 0
		left_index := len(line)
		right_index := -1

		for _, v := range toCheckFor {
			if strings.Index(line, v) < left_index && strings.Index(line, v) >= 0 {
				left_index = strings.Index(line, v)
				left = extractDigit(v)
			}
			if strings.LastIndex(line, v) > right_index && strings.LastIndex(line, v) >= 0 {
				right_index = strings.LastIndex(line, v)
				right = extractDigit(v)
			}
		}
        sum = sum + (left*10 + right)
	}
	fmt.Println(sum)
}
