package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kevindoubleu/advent-of-code-2023/lib"
)

var (
	numbers = strings.Split(lib.DIGITS, "")
)

func Part1(filename string) {
	lines := lib.Read(filename)
	sum := 0

	for _, line := range lines {
		sum += getCombinedDigit(line)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func getCombinedDigit(line string) int {
	tens := getFirstDigit(line) * 10
	ones := getLastDigit(line)

	return tens + ones
}

func getFirstDigit(line string) int {
	chars := strings.Split(line, "")

	for i := 0; i < len(chars); i++ {
		for _, number := range numbers {
			if chars[i] == number {
				firstDigit, _ := strconv.ParseInt(chars[i], 10, 0)
				return int(firstDigit)
			}
		}
	}
	return 0
}

func getLastDigit(line string) int {
	chars := strings.Split(line, "")

	for i := len(chars) - 1; i >= 0; i-- {
		for _, number := range numbers {
			if chars[i] == number {
				lastDigit, _ := strconv.ParseInt(chars[i], 10, 0)
				return int(lastDigit)
			}
		}
	}
	return 0
}
