package main

import (
	"os"

	"github.com/kevindoubleu/advent-of-code-2023/days/day01"
	"github.com/kevindoubleu/advent-of-code-2023/lib"
)

func main() {
	os.Chdir("days/day01")
	// day01.Part1(lib.TESTFILE)
	day01.Part1(lib.INPUTFILE)
	day01.Part2(lib.INPUTFILE)
}
