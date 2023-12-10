package main

import (
	"log"

	"github.com/timsexperiments/aoc-2024/internal/day3"
	"github.com/timsexperiments/aoc-2024/internal/input"
)

func main() {
	contents, err := input.ReadAocInput(3)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	day3.Part1(contents)
	day3.Part2(contents)
}
