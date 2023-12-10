package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/timsexperiments/aoc-2024/internal/util"
)

func Part1(input []string) {
	numbers, err := getNumbers(input, readNumbersFromLine)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		return
	}

	fmt.Printf("The answer to part one is %d\n", util.SumSlice(numbers))
}

func Part2(input []string) {
	numbers, err := getNumbers(input, readNumbersFromLineIncludingSpelled)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		return
	}

	fmt.Printf("The answer to part two is %d\n", util.SumSlice(numbers))
}

func getNumbers(input []string, readFn func(string) (int, error)) ([]int, error) {
	numbers := make([]int, len(input))
	for _, line := range input {
		fromLine, err := readFn(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, fromLine)
	}
	return numbers, nil
}

func readNumbersFromLine(line string) (int, error) {
	lineLen := len(line)
	first := ""
	last := ""
	for i := 0; i < len(line); i++ {
		if first == "" && util.IsNumber(line[i]) {
			first = string(line[i])
		}
		if last == "" && util.IsNumber(line[lineLen-i-1]) {
			last = string(line[lineLen-i-1])
		}
		if first != "" && last != "" {
			break
		}
	}
	if first == "" || last == "" {
		return 0, nil
	}
	return strconv.Atoi(fmt.Sprintf("%s%s", first, last))
}

func readNumbersFromLineIncludingSpelled(line string) (int, error) {
	lineLen := len(line)
	first := ""
	last := ""
	for i := 0; i < len(line); i++ {
		if first == "" {
			if util.IsNumber(line[i]) {
				first = string(line[i])
			} else if spelledStartNumber := isSpelledNumber(line, i); spelledStartNumber > -1 {
				first = fmt.Sprint(spelledStartNumber)
			}
		}
		if last == "" {
			if util.IsNumber(line[lineLen-i-1]) {
				last = string(line[lineLen-i-1])
			} else if spelledEndNumber := isSpelledNumber(line, lineLen-i-1); spelledEndNumber > -1 {
				last = fmt.Sprint(spelledEndNumber)
			}
		}
		if first != "" && last != "" {
			break
		}
	}
	if first == "" || last == "" {
		return 0, nil
	}
	return strconv.Atoi(fmt.Sprintf("%s%s", first, last))
}

// zero
// one
// two
// three
// four
// five
// six
// seven
// eight
// nine

// Starting letters o, t, f, s, e, n
// Ending letters e, o, r, x, n, t

func isSpelledNumber(line string, startOrEnd int) int {
	threeLetterFromStart := util.LettersFromStart(line, startOrEnd, 3)
	threeLetterFromEnd := util.LettersFromEnd(line, startOrEnd, 3)
	fourLetterFromStart := util.LettersFromStart(line, startOrEnd, 4)
	fourLetterFromEnd := util.LettersFromEnd(line, startOrEnd, 4)
	fiveLetterFromStart := util.LettersFromStart(line, startOrEnd, 5)
	fiveLetterFromEnd := util.LettersFromEnd(line, startOrEnd, 5)
	if isZero(fourLetterFromStart) || isZero(fourLetterFromEnd) {
		return 0
	}
	if isOne(threeLetterFromStart) || isOne(threeLetterFromEnd) {
		return 1
	}
	if isTwo(threeLetterFromStart) || isTwo(threeLetterFromEnd) {
		return 2
	}
	if isThree(fiveLetterFromStart) || isThree(fiveLetterFromEnd) {
		return 3
	}
	if isFour(fourLetterFromStart) || isFour(fourLetterFromEnd) {
		return 4
	}
	if isFive(fourLetterFromStart) || isFive(fourLetterFromEnd) {
		return 5
	}
	if isSix(threeLetterFromStart) || isSix(threeLetterFromEnd) {
		return 6
	}
	if isSeven(fiveLetterFromStart) || isSeven(fiveLetterFromEnd) {
		return 7
	}
	if isEight(fiveLetterFromStart) || isEight(fiveLetterFromEnd) {
		return 8
	}
	if isNine(fourLetterFromStart) || isNine(fourLetterFromEnd) {
		return 9
	}
	return -1
}

func isZero(str string) bool {
	return strings.ToLower(str) == "zero"
}

func isOne(str string) bool {
	return strings.ToLower(str) == "one"
}

func isTwo(str string) bool {
	return strings.ToLower(str) == "two"
}

func isThree(str string) bool {
	return strings.ToLower(str) == "three"
}

func isFour(str string) bool {
	return strings.ToLower(str) == "four"
}

func isFive(str string) bool {
	return strings.ToLower(str) == "five"
}

func isSix(str string) bool {
	return strings.ToLower(str) == "six"
}

func isSeven(str string) bool {
	return strings.ToLower(str) == "seven"
}

func isEight(str string) bool {
	return strings.ToLower(str) == "eight"
}

func isNine(str string) bool {
	return strings.ToLower(str) == "nine"
}
