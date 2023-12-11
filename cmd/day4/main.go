package main

import (
	"log"

	"github.com/timsexperiments/aoc-2024/internal/day4"
	"github.com/timsexperiments/aoc-2024/internal/input"
)

func main() {
	contents, err := input.ReadAocInput(4)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	day4.Part1(contents)
	day4.Part2(contents)
}
