package main

import (
	"log"

	"github.com/timsexperiments/aoc-2024/internal/day1"
	"github.com/timsexperiments/aoc-2024/internal/input"
)

func main() {
	contents, err := input.ReadAocInput(1)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	day1.Part1(contents)
	day1.Part2(contents)
}
