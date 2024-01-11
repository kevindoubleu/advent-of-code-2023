package day02

import (
	"fmt"

	"github.com/kevindoubleu/advent-of-code-2023/lib"
)

func Part2(filename string) {
	lines := lib.Read(filename)
	games := parseGames(lines)
	sum := 0

	for _, game := range games {
		power := game.highRed * game.highGreen * game.highBlue
		sum += power
	}

	fmt.Printf("The answer is: %d\n", sum)
}
