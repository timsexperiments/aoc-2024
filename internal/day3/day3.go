package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/timsexperiments/aoc-2024/internal/util"
)

func Part1(contents []string) error {
	schematics, err := createSchematics(contents)
	if err != nil {
		return err
	}
	validParts := schematics.validParts()
	answer := 0
	for _, part := range validParts {
		answer += part.PartNo
	}
	fmt.Printf("The answer to part one is %d\n", answer)
	return nil
}

func Part2(contents []string) error {
	schematics, err := createSchematics(contents)
	if err != nil {
		return err
	}
	answer := 0
	for _, gear := range schematics.Gears {
		answer += gear.Ratio()
	}
	fmt.Printf("The answer to part two is %d\n", answer)
	return nil
}

type location struct {
	X, Y int
}

func (location location) String() string {
	return fmt.Sprintf("location { X: %d, Y: %d }", location.X, location.Y)
}

type schematics struct {
	Gears           []gear
	Parts           []part
	partLocationMap map[location]part
	raw             []string
}

func (schematics schematics) validParts() []part {
	validParts := make([]part, 0)
	for _, part := range schematics.Parts {
		if bySymbol(schematics.raw, part) {
			validParts = append(validParts, part)
		}
	}
	return validParts
}

func (schematics schematics) String() string {
	return fmt.Sprintf("{ parts: %v }", schematics.Parts)
}

func createSchematics(instructions []string) (*schematics, error) {
	parts := make([]part, 0)
	partLocationMap := make(map[location]part)
	for i, instruction := range instructions {
		instructionParts, err := findPartsInInstruction(instruction, i)
		if err != nil {
			return nil, err
		}
		parts = append(parts, instructionParts...)
		for _, part := range instructionParts {
			for x := part.start; x <= part.end; x++ {
				location := location{x, part.instruction}
				partLocationMap[location] = part
			}
		}
	}
	gears := findAllGears(instructions, partLocationMap)

	return &schematics{Parts: parts, Gears: gears, raw: instructions, partLocationMap: partLocationMap}, nil
}

func findAllGears(instructions []string, partLocationMap map[location]part) []gear {
	gears := make([]gear, 0)
	for y := 0; y < len(instructions); y++ {
		rowAbove := max(y-1, 0)
		rowBelow := min(y+1, len(instructions)-1)
		for x := 0; x < len(instructions[y]); x++ {
			positionBefore := max(x-1, 0)
			positionAfter := min(x+1, len(instructions[y])-1)
			if isGear(instructions[y][x]) {
				gear := gear{Location: location{x, y}, adjacentParts: make(map[part]bool)}
				topLeft := location{positionBefore, rowAbove}
				if part, ok := partLocationMap[topLeft]; ok {
					gear.adjacentParts[part] = true
				}
				topMiddle := location{x, rowAbove}
				if part, ok := partLocationMap[topMiddle]; ok {
					gear.adjacentParts[part] = true
				}
				topRight := location{positionAfter, rowAbove}
				if part, ok := partLocationMap[topRight]; ok {
					gear.adjacentParts[part] = true
				}
				left := location{positionBefore, y}
				if part, ok := partLocationMap[left]; ok {
					gear.adjacentParts[part] = true
				}
				right := location{positionAfter, y}
				if part, ok := partLocationMap[right]; ok {
					gear.adjacentParts[part] = true
				}
				bottomLeft := location{positionBefore, rowBelow}
				if part, ok := partLocationMap[bottomLeft]; ok {
					gear.adjacentParts[part] = true
				}
				bottomMiddle := location{x, rowBelow}
				if part, ok := partLocationMap[bottomMiddle]; ok {
					gear.adjacentParts[part] = true
				}
				bottomRight := location{positionAfter, rowBelow}
				if part, ok := partLocationMap[bottomRight]; ok {
					gear.adjacentParts[part] = true
				}

				if len(gear.adjacentParts) == 2 {
					gears = append(gears, gear)
				}
			}
		}
	}
	return gears
}

func findPartsInInstruction(instruction string, row int) ([]part, error) {
	parts := make([]part, 0)
	prevWasNumber := false
	for i, char := range instruction {
		isNumber := !util.IsNaN(char)
		if !prevWasNumber && isNumber {
			partNo, start, _ := numberAndPositionEnd(instruction, i)
			part, err := createPart(partNo, row, start)
			if err != nil {
				return nil, err
			}
			parts = append(parts, *part)
		}

		if isNumber {
			prevWasNumber = true
		} else {
			prevWasNumber = false
		}
	}
	return parts, nil
}

type gear struct {
	Location      location
	adjacentParts map[part]bool
}

func (g gear) Ratio() int {
	ratio := 1
	for part := range g.adjacentParts {
		ratio *= part.PartNo
	}
	return ratio
}

func (g gear) String() string {
	return fmt.Sprintf("gear { Location: %v, adjacentParts: %v }", g.Location, g.adjacentParts)
}

type part struct {
	PartNo      int
	start       int
	end         int
	instruction int
}

func (p part) String() string {
	return fmt.Sprintf("%d", p.PartNo)
}

func createPart(number string, row, start int) (*part, error) {
	converterd, err := strconv.Atoi(number)
	if err != nil {
		return nil, err
	}
	part := &part{PartNo: converterd, instruction: row, start: start, end: start + len(number) - 1}
	return part, nil
}

func numberAndPositionEnd(line string, start int) (string, int, int) {
	if util.IsNaN(rune(line[start])) {
		return "", -1, -1
	}
	var number strings.Builder
	end := start
	for util.IsNumber(line[end]) {
		number.WriteByte(line[end])
		if end == len(line)-1 {
			break
		}
		end++
	}
	if number.Len() == 0 {
		return "", -1, -1
	}
	return number.String(), start, end
}

func bySymbol(schematics []string, part part) bool {
	row := part.instruction
	rowAbove := max(0, row-1)
	rowBelow := min(len(schematics)-1, row+1)
	for i := part.start; i <= part.end; i++ {
		leftOfPosition := max(0, i-1)
		rightOfPosition := min(len(schematics[row])-1, i+1)
		if isSymbol(schematics[rowAbove][leftOfPosition]) {
			return true
		}
		if isSymbol(schematics[rowAbove][i]) {
			return true
		}
		if isSymbol(schematics[rowAbove][rightOfPosition]) {
			return true
		}
		if isSymbol(schematics[row][leftOfPosition]) {
			return true
		}
		if isSymbol(schematics[row][rightOfPosition]) {
			return true
		}
		if isSymbol(schematics[rowBelow][leftOfPosition]) {
			return true
		}
		if isSymbol(schematics[rowBelow][i]) {
			return true
		}
		if isSymbol(schematics[rowBelow][rightOfPosition]) {
			return true
		}
	}
	return false
}

func isSymbol(char byte) bool {
	return util.IsNaN(rune(char)) && char != '.'
}

func isGear(char byte) bool {
	return char == '*'
}
