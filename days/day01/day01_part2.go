package day01

import (
	"fmt"

	"github.com/kevindoubleu/advent-of-code-2023/lib"
)

var (
	digitsWithSpelled = map[string]int{
		"one":   1,
		"1":     1,
		"two":   2,
		"2":     2,
		"three": 3,
		"3":     3,
		"four":  4,
		"4":     4,
		"five":  5,
		"5":     5,
		"six":   6,
		"6":     6,
		"seven": 7,
		"7":     7,
		"eight": 8,
		"8":     8,
		"nine":  9,
		"9":     9,
		// "zero",
	}
)

func Part2(filename string) {
	lines := lib.Read(filename)
	sum := 0

	for _, line := range lines {
		sum += getCombinedDigitWithSpelled(line)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func getCombinedDigitWithSpelled(line string) int {
	tens := getFirstDigitWithSpelled(line) * 10
	ones := getLastDigitWithSpelled(line)

	return tens + ones
}

func getFirstDigitWithSpelled(line string) int {
	for i := range line {
		for key, val := range digitsWithSpelled {
			substringLen := min(len(key), len(line)-i)
			if line[i:i+substringLen] == key {
				return val
			}
		}
	}
	return 0
}

func getLastDigitWithSpelled(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		for key, val := range digitsWithSpelled {
			substringLen := min(len(key), len(line)-i)
			if line[i:i+substringLen] == key {
				return val
			}
		}
	}
	return 0
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}
