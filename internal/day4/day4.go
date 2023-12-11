package day4

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/timsexperiments/aoc-2024/internal/util"
)

func Part1(contents []string) error {
	cards, err := readCards(contents)
	if err != nil {
		return err
	}
	answer := 0
	for _, card := range cards {
		answer += card.Score()
	}
	fmt.Printf("The answer to part one is %d\n", answer)
	return nil
}

func Part2(contents []string) error {
	cards, err := readCardsToMap(contents)
	if err != nil {
		return err
	}
	answer := totalCardsAndNewCards(cards)
	fmt.Printf("The answer to part two is %d\n", answer)
	return nil
}

func readCards(contents []string) ([]card, error) {
	cards := make([]card, 0)
	for _, details := range contents {
		card, err := createCard(details)
		if err != nil {
			return nil, err
		}
		cards = append(cards, *card)
	}
	return cards, nil
}

func readCardsToMap(contents []string) (map[int]card, error) {
	cards := map[int]card{}
	cardList, err := readCards(contents)
	if err != nil {
		return nil, err
	}
	for _, card := range cardList {
		cards[card.ID] = card
	}
	return cards, nil
}

type card struct {
	ID           int
	Numbers      []string
	Winning      []string
	intersection []string
}

func (c card) Score() int {
	return int(math.Pow(2, float64(len(c.intersection)-1)))
}

func (c card) String() string {
	return fmt.Sprintf("Card %d { numbers: %v, winnings: %v, intersection: %v, score: %d }", c.ID, c.Numbers, c.Winning, c.intersection, c.Score())
}

func areEqual(a, b card) bool {
	return reflect.DeepEqual(a, b)
}

func createCard(details string) (*card, error) {
	details = util.StripDuplicateWhitespace(details)
	parts := strings.Split(details, ": ")
	if len(parts) != 2 {
		return nil, errors.New("Card details string must have a card number and a list of numbers separated by ': '.")
	}
	idString, numbersParts := parts[0], strings.Split(parts[1], " | ")
	rawId := util.NumbersFromString(idString)
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return nil, fmt.Errorf("Unable to get id from card details: %v", err)
	}
	if len(numbersParts) != 2 {
		return nil, errors.New("Card details must contain a list of winning numbers separated by the list of you card numbers. Each of the numbers must be separated by a ' ' character.")
	}
	winningNumbers := util.ArrMap(strings.Split(numbersParts[0], " "), util.NumbersFromString)
	numbers := util.ArrMap(strings.Split(numbersParts[1], " "), util.NumbersFromString)
	intersection := util.Intesect(winningNumbers, numbers)
	return &card{id, winningNumbers, numbers, intersection}, nil
}

func totalCardsAndNewCards(cards map[int]card) int {
	cardList := make([]card, 0)
	for _, card := range cards {
		cardList = append(cardList, card)
	}
	currentCard := 0
	for currentCard < len(cardList) {
		card := cardList[currentCard]
		totalToAdd := len(card.intersection)
		for i := card.ID + 1; i <= card.ID+totalToAdd; i++ {
			if val, ok := cards[i]; ok {
				cardList = append(cardList, val)
			}
		}
		currentCard++
	}
	return len(cardList)
}
