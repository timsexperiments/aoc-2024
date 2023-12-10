package day2

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/timsexperiments/aoc-2024/internal/util"
)

type Color int

const (
	Green Color = iota
	Red
	Blue
)

func Part1(contents []string) error {
	games := make([]game, 0)
	for _, line := range contents {
		game, err := createGame(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create game: %s\n", err)
			return err
		}
		games = append(games, *game)
	}
	ids := make([]int, 0)
	for _, game := range games {
		if game.couldMeetTarget(12, 13, 14) {
			ids = append(ids, game.id)
		}
	}
	answer := util.SumSlice(ids)
	fmt.Printf("The answer to part one is %d\n", answer)
	return nil
}

func Part2(contents []string) error {
	games := make([]game, 0)
	for _, line := range contents {
		game, err := createGame(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create game: %s\n", err)
			return err
		}
		games = append(games, *game)
	}
	answer := 0
	for _, game := range games {
		answer += game.power()
	}
	fmt.Printf("The answer to part two is %d\n", answer)
	return nil
}

type game struct {
	id     int
	rounds []round
}

type round struct {
	red   int
	green int
	blue  int
}

func (g game) couldMeetTarget(red, green, blue int) bool {
	return g.minRed() <= red && g.minGreen() <= green && g.minBlue() <= blue
}

func (g game) power() int {
	return g.minRed() * g.minGreen() * g.minBlue()
}

func (g game) minRed() int {
	max := -1
	for _, round := range g.rounds {
		max = int(math.Max(float64(max), float64(round.red)))
	}
	return max
}

func (g game) minGreen() int {
	max := -1
	for _, round := range g.rounds {
		max = int(math.Max(float64(max), float64(round.green)))
	}
	return max
}

func (g game) minBlue() int {
	max := -1
	for _, round := range g.rounds {
		max = int(math.Max(float64(max), float64(round.blue)))
	}
	return max
}

func (g game) String() string {
	return fmt.Sprintf("Game %d: %v", g.id, g.rounds)
}

func (r round) setRed(amount int) {
	r.red = amount
}

func (r round) setGreen(amount int) {
	r.green = amount
}

func (r round) setBlue(amount int) {
	r.blue = amount
}

func (r round) String() string {
	return fmt.Sprintf("{ red: %d, green: %d, blue: %d }", r.red, r.green, r.blue)
}

func createRound(roundStr string) (*round, error) {
	red, green, blue := 0, 0, 0
	numbers := make([]int, 0)
	wasInt := false
	var number strings.Builder
	colors := make([]Color, 0)
	for i, char := range roundStr {
		maybeRed := util.LettersFromStart(roundStr, i, 3)
		maybeGreen := util.LettersFromStart(roundStr, i, 5)
		maybeBlue := util.LettersFromStart(roundStr, i, 4)
		if maybeRed == "red" {
			colors = append(colors, Red)
		}
		if maybeGreen == "green" {
			colors = append(colors, Green)
		}
		if maybeBlue == "blue" {
			colors = append(colors, Blue)
		}
		if util.IsNumber(byte(char)) {
			number.WriteRune(char)
			wasInt = true
		} else {
			wasInt = false
		}
		if !wasInt && number.Len() > 0 {
			fullNumber := number.String()
			converted, err := strconv.Atoi(fullNumber)
			if err != nil {
				return nil, fmt.Errorf("Unable to convert [%s] to a number: %s\n", fullNumber, err)
			}
			numbers = append(numbers, converted)
			number.Reset()
		}
		if len(roundStr)-i < 4 {
			break
		}
	}

	for i, color := range colors {
		if color == Red {
			red = numbers[i]
		}
		if color == Green {
			green = numbers[i]
		}
		if color == Blue {
			blue = numbers[i]
		}
	}

	round := &round{red: red, green: green, blue: blue}
	return round, nil
}

func createGame(line string) (*game, error) {
	gameParts := strings.Split(line, ":")
	if len(gameParts) != 2 {
		return nil, errors.New("Could not identify game parts. A game should have a name and round details separated by a comma.")
	}
	gameString, roundString := gameParts[0], gameParts[1]
	id, err := strconv.Atoi(util.NumbersFromString(gameString))
	if err != nil {
		return nil, err
	}

	game := game{id: id}
	if game.id == 0 {
		return nil, errors.New("Unable to create a game without an ID.")
	}
	rounds := strings.Split(roundString, ";")
	for _, roundStr := range rounds {
		round, err := createRound(roundStr)
		if err != nil {
			return nil, err
		}
		game.rounds = append(game.rounds, *round)
	}
	return &game, nil
}

type score struct {
	color  Color
	amount int
}
