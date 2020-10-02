package evaluator

import (
	"testing"
)

func TestEvaluateHand(t *testing.T) {
	e, err := NewEvaluator()
	if err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		hand           []string
		expectHandType string
	}{
		{
			hand:           []string{"SA", "S2", "S3", "S4", "S5"},
			expectHandType: "STRAIGHT FLUSH",
		}, {
			hand:           []string{"HA", "S2", "S3", "S4", "S5"},
			expectHandType: "STRAIGHT",
		}, {
			hand:           []string{"S7", "S2", "S3", "S4", "S5"},
			expectHandType: "FLUSH",
		}, {
			hand:           []string{"H2", "S2", "S3", "S4", "S5"},
			expectHandType: "PAIR",
		}, {
			hand:           []string{"H2", "S2", "C2", "S4", "S5"},
			expectHandType: "THREE OF A KIND",
		}, {
			hand:           []string{"DK", "S2", "S3", "S4", "S5"},
			expectHandType: "HIGH CARD",
		},
	}

	for _, test := range tests {
		_, handType := e.EvalHand(test.hand)
		if handType != test.expectHandType {
			t.Errorf("Expected %s, got %s", test.expectHandType, handType)
		}
	}
}

func BenchmarkEvaluateAllHands(b *testing.B) {
	e, err := NewEvaluator()
	if err != nil {
		b.Error(e)
		return
	}

	handTypeFrequency := make(map[string]uint32)
	_, cards := GenerateCardEncodes()

	for c1 := 0; c1 < len(cards)-6; c1++ {
		for c2 := c1 + 1; c2 < len(cards)-5; c2++ {
			for c3 := c2 + 1; c3 < len(cards)-4; c3++ {
				for c4 := c3 + 1; c4 < len(cards)-3; c4++ {
					for c5 := c4 + 1; c5 < len(cards)-2; c5++ {
						for c6 := c5 + 1; c6 < len(cards)-1; c6++ {
							for c7 := c6 + 1; c7 < len(cards); c7++ {
								hand := []string{
									cards[c1],
									cards[c2],
									cards[c3],
									cards[c4],
									cards[c5],
									cards[c6],
									cards[c7],
								}
								_, handType := e.EvalHand(hand)
								handTypeFrequency[handType] = handTypeFrequency[handType] + 1
							}
						}
					}
				}
			}
		}
	}

	tests := []struct {
		handType     string
		expectCounts uint32
	}{
		{
			handType:     "HIGH CARD",
			expectCounts: 23294460,
		}, {
			handType:     "PAIR",
			expectCounts: 58627800,
		}, {
			handType:     "TWO PAIRS",
			expectCounts: 31433400,
		}, {
			handType:     "THREE OF A KIND",
			expectCounts: 6461620,
		}, {
			handType:     "STRAIGHT",
			expectCounts: 6180020,
		}, {
			handType:     "FLUSH",
			expectCounts: 4047644,
		}, {
			handType:     "FULL HOUSE",
			expectCounts: 3473184,
		}, {
			handType:     "FOUR OF A KIND",
			expectCounts: 224848,
		}, {
			handType:     "STRAIGHT FLUSH",
			expectCounts: 41584,
		},
	}

	for _, test := range tests {
		if frequency, ok := handTypeFrequency[test.handType]; !ok || frequency != test.expectCounts {
			b.Errorf("Incorrect frequence for hand type %s. Expect %d, got %d.", test.handType, test.expectCounts, frequency)
		}
	}
}
