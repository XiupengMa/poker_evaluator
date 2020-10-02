package hand

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

// Card is a 2 char string that representing suit and rank respectively.
type Card string

// NewCard creates a Card from its string representation.
func NewCard(str string) Card {
	if isValidCard(str) {
		return Card(str)
	}
	return ""
}

func isValidCard(str string) bool {
	if len(str) != 2 {
		return false
	}
	suit, rank := str[0:1], str[1:]
	if !find(GetValidSuits(), suit) {
		return false
	}
	if !find(GetValidRanks(), rank) {
		return false
	}
	return true
}

func find(slice []string, target string) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}
