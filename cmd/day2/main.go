package main

import (
	"log"

	"github.com/timsexperiments/aoc-2024/internal/day2"
	"github.com/timsexperiments/aoc-2024/internal/input"
)

func main() {
	contents, err := input.ReadAocInput(2)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	day2.Part1(contents)
	day2.Part2(contents)
}
