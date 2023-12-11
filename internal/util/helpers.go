package util

import (
	"regexp"
	"strings"
)

func IsNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsNaN(c rune) bool {
	return c < '0' || c > '9'
}

func SumSlice(sliceToSum []int) int {
	sum := 0
	for i := 0; i < len(sliceToSum); i++ {
		sum += sliceToSum[i]
	}
	return sum
}

func LettersFromStart(word string, start, numLetters int) string {
	end := min(len(word), start+numLetters)
	return word[start:end]
}

func LettersFromEnd(word string, end, numLetters int) string {
	start := max(0, end-numLetters)
	return word[start+1 : end+1]
}

func NumbersFromString(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if !IsNaN(c) {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func Intesect[T comparable](first []T, second []T) []T {
	longerList, shorterList := longerThenShorter(&first, &second)
	longer, shorter := ToSet(*longerList), ToSet(*shorterList)
	intersection := make([]T, 0)
	for key := range shorter {
		if longer[key] {
			intersection = append(intersection, key)
		}
	}
	return intersection
}

func ToSet[T comparable](slice []T) map[T]bool {
	m := map[T]bool{}
	for _, val := range slice {
		m[val] = true
	}
	return m
}

func longerThenShorter[T comparable](first *[]T, second *[]T) (*[]T, *[]T) {
	if len(*first) > len(*second) {
		return first, second
	} else {
		return second, first
	}
}

func ArrMap[T any, U any](arr []T, f func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func StripDuplicateWhitespace(str string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(str, " ")
}
