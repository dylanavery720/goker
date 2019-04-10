package isflush

import (
	isstraight "poker/isStraight"
	types "poker/types"
)

func IsFlush(hand types.Hand) string {
	var rank = ""
	for i := 0; i < len(hand.Suits)-1; i++ {
		if hand.Suits[i] == hand.Suits[i+1] {
			continue
		} else {
			return "next"
		}
	}
	rank = "Flush"
	if isstraight.IsStraight(hand) == "Straight" {
		rank = "Straight Flush"
	}
	if hand.Values[0] >= 10 {
		rank = "Royal Flush"
	}
	return rank
}
