package hand

// Hand is a slice of Cards
type Hand []Card

func (h *Hand) Len() int {
	return len(*h)
}
