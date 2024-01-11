package day02

import (
	"fmt"
	"math"
	"strings"

	"github.com/kevindoubleu/advent-of-code-2023/lib"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14

	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

type Game struct {
	id int

	highRed   int
	highGreen int
	highBlue  int

	// part2
	lowRed   int
	lowGreen int
	lowBlue  int
}

func Part1(filename string) {
	lines := lib.Read(filename)
	games := parseGames(lines)
	sum := 0

	for _, game := range games {
		if gameIsPossible(game) {
			sum += game.id
		}
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func parseGames(lines []string) []Game {
	var games []Game
	for _, line := range lines {
		game := parseGame(line)
		games = append(games, game)
	}
	return games
}

func parseGame(line string) Game {
	game := newGame()
	var roundsString string

	parts := strings.Split(line, ": ")
	headerString, roundsString := parts[0], parts[1]

	fmt.Sscanf(headerString, "Game %d", &game.id)
	parseRounds(game, roundsString)

	return *game
}

func newGame() *Game {
	return &Game{
		lowRed:   math.MaxInt,
		lowGreen: math.MaxInt,
		lowBlue:  math.MaxInt,
	}
}

func parseRounds(game *Game, roundsString string) {
	rounds := strings.Split(roundsString, "; ")
	for _, round := range rounds {
		parseRound(game, round)
	}
}

func parseRound(game *Game, round string) {
	colorWithCounts := strings.Split(round, ", ")
	for _, colorWithCount := range colorWithCounts {
		if isColor(RED, colorWithCount) {
			updateHigh(&game.highRed, colorWithCount)
		}
		if isColor(GREEN, colorWithCount) {
			updateHigh(&game.highGreen, colorWithCount)
		}
		if isColor(BLUE, colorWithCount) {
			updateHigh(&game.highBlue, colorWithCount)
		}
	}
}

func isColor(color string, colorWithCount string) bool {
	return strings.Contains(colorWithCount, color)
}

func updateHigh(i *int, colorWithCount string) {
	var count int
	fmt.Sscanf(colorWithCount, "%d", &count)
	if count > *i {
		*i = count
	}
}

func gameIsPossible(game Game) bool {
	if game.highRed > maxRed ||
		game.highGreen > maxGreen ||
		game.highBlue > maxBlue {
		return false
	}
	return true
}
