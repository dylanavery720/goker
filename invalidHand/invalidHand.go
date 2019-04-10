package invalidhand

import (
	isstraight "poker/isStraight"
	types "poker/types"
)

func InvalidHand(hand types.Hand) string {
	if len(hand.Values) != 5 {
		return "Invalid Hand"
	}
	if isstraight.Contains(hand.Values, 0) {
		return "Invalid Hand"
	}
	return "Valid Hand"
}
