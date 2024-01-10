package lib

import (
	"bufio"
	"os"
)

func Read(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
