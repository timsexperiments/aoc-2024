package util

func IsNumber(c byte) bool {
	return c >= '0' && c <= '9'
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
