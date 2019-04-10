package isstraight

import (
	types "poker/types"
)

func IsStraight(hand types.Hand) string {
	// add handle ace, test for handle ace
	for i := 0; i < len(hand.Values)-1; i++ {
		if hand.Values[i] == (hand.Values[i+1] - 1) {
			continue
		} else {
			return "false"
		}
	}
	return "Straight"
}
