package input

import (
	"fmt"
	"os"
	"strings"
)

func ReadAocInput(aocDay int) ([]string, error) {
	filename := fmt.Sprintf("assets/inputs/day%d.txt", aocDay)
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}
