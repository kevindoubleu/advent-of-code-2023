package main

import (
	"os"

	"github.com/kevindoubleu/advent-of-code-2023/days/day02"
	"github.com/kevindoubleu/advent-of-code-2023/lib"
)

func main() {
	// os.Chdir("days/day01")
	// day01.Part1(lib.TESTFILE)
	// day01.Part1(lib.INPUTFILE)
	// day01.Part2(lib.INPUTFILE)

	os.Chdir("days/day02")
	// day02.Part1(lib.TESTFILE)
	// day02.Part1(lib.INPUTFILE)
	// day02.Part2(lib.TESTFILE)
	day02.Part2(lib.INPUTFILE)
}
