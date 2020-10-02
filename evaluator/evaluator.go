package evaluator

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/XiupengMa/2p2_poker_evaluator/hand"
)

type Evaluator struct {
	handRanksLookups []uint32
	handTypes        []string
	cardIndices      map[hand.Card]uint8
}

func NewEvaluator(handRanksFilePath string) (*Evaluator, error) {
	ranks, err := loadHandRank(handRanksFilePath)
	if err != nil {
		return nil, err
	}
	return &Evaluator{
		handRanksLookups: ranks,
		handTypes:        getHandTypes(),
		cardIndices:      generateCardIndices(),
	}, nil
}

func (evaluator *Evaluator) evalHand(hand hand.Hand) (uint32, string) {
	index := uint32(53)

	for _, card := range hand {
		index = evaluator.handRanksLookups[index+uint32(evaluator.cardIndices[card])]
	}

	if hand.Len() < 7 {
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

func getHandTypes() []string {
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

func generateCardIndices() map[hand.Card]uint8 {
	ranks := hand.GetValidRanks()
	suits := hand.GetValidSuits()

	count := uint8(1)
	indices := make(map[hand.Card]uint8)
	for _, rank := range ranks {
		for _, suit := range suits {
			card := hand.Card(fmt.Sprintf("%s%s", suit, rank))
			indices[card] = count
			count++
		}
	}
	return indices
}

// Debug print out debug info
func Debug() {
	e, err := NewEvaluator("./data/HandRanks.dat")

	if err != nil {
		fmt.Println(err)
	}
	hand1 := []hand.Card{
		hand.Card("S2"),
		hand.Card("S3"),
		hand.Card("S4"),
		hand.Card("S5"),
		hand.Card("S6"),
	}

	hand2 := []hand.Card{
		hand.Card("S2"),
		hand.Card("S3"),
		hand.Card("S4"),
		hand.Card("S5"),
		hand.Card("SA"),
	}
	rank, name := e.evalHand(hand1)
	fmt.Println(rank)
	fmt.Println(name)

	rank2, name2 := e.evalHand(hand2)
	fmt.Println(rank2)
	fmt.Println(name2)
}
