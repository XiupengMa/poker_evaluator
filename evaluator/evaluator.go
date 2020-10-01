package evaluator

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

type Evaluator struct {
	handRanks []uint32
}

func NewEvaluator(handRanksFilePath string) (*Evaluator, error) {
	ranks, err := loadHandRank(handRanksFilePath)
	if err != nil {
		return nil, err
	}
	return &Evaluator{
		handRanks: ranks,
	}, nil
}

func (evaluator *Evaluator) evalHand() {

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

// Debug print out debug info
func Debug() {
	ranks, err := loadHandRank()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(ranks))
	fmt.Println(ranks[54])
	for i := 0; i < 100; i++ {
		fmt.Printf("%d:%d\n", i, ranks[i])
	}
}
