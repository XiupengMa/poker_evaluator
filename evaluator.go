package evaluator

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

// GetValidRanks returns all the valid ranks
func GetValidRanks() []string {
	return []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"T",
		"J",
		"Q",
		"K",
		"A",
	}
}

// GetValidSuits returns all the valid suits
func GetValidSuits() []string {
	return []string{
		"S",
		"H",
		"C",
		"D",
	}
}

// Evaluator is used to evaluate a hand
type Evaluator struct {
	handRanksLookups []uint32
	handTypes        []string
	cardEncodes      map[string]uint8
}

// NewEvaluator returns a new EValuator
func NewEvaluator() (*Evaluator, error) {
	ranks, err := loadHandRank("./data/HandRanks.dat")
	if err != nil {
		return nil, err
	}
	encodes, _ := GenerateCardEncodes()
	return &Evaluator{
		handRanksLookups: ranks,
		handTypes:        GetHandTypes(),
		cardEncodes:      encodes,
	}, nil
}

// EvalHand evaluates a hand and returns its rank and type
func (evaluator *Evaluator) EvalHand(hand []string) (uint32, string) {
	index := uint32(53)

	for _, card := range hand {
		index = evaluator.handRanksLookups[index+uint32(evaluator.cardEncodes[card])]
	}

	if len(hand) < 7 {
		index = evaluator.handRanksLookups[index]
	}

	return index, evaluator.handTypes[index>>12]
}

func loadHandRank(filePath string) ([]uint32, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	bufr := bufio.NewReader(file)

	size := stats.Size()
	rankSize, remainder := size/4, size%4

	if remainder != 0 {
		return nil, errors.New("wrong rank size")
	}

	ranks := make([]uint32, rankSize)
	err = binary.Read(bufr, binary.LittleEndian, ranks)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return ranks, nil
}

// GetHandTypes return all the valid hand types
func GetHandTypes() []string {
	return []string{
		"BAD",
		"HIGH CARD",
		"PAIR",
		"TWO PAIRS",
		"THREE OF A KIND",
		"STRAIGHT",
		"FLUSH",
		"FULL HOUSE",
		"FOUR OF A KIND",
		"STRAIGHT FLUSH",
	}
}

// GenerateCardEncodes returns a map of card to its encode, and a slice of all valid cards
func GenerateCardEncodes() (map[string]uint8, []string) {
	ranks := GetValidRanks()
	suits := GetValidSuits()

	count := uint8(0)
	cards := make([]string, 52)
	encodes := make(map[string]uint8)
	for _, rank := range ranks {
		for _, suit := range suits {
			card := fmt.Sprintf("%s%s", suit, rank)
			cards[count] = card
			encodes[card] = count + 1
			count++
		}
	}
	return encodes, cards
}
