package isofakind

import (
	types "poker/types"
)

func IsOfAKind(hand types.Hand) string {
	var rank = "High Card"
	var cards = countCards(hand)
	for k := range cards {
		if cards[k] == 4 {
			return "Four Of A Kind"
		}
		if cards[k] == 3 {
			rank = "Three Of A Kind"
			if len(cards) < 3 {
				return "Full House"
			}
		}
		if cards[k] == 2 {
			rank = "One Pair"
			if len(cards) < 3 {
				return "Full House"
			}
			if len(cards) < 4 {
				return "Two Pair"
			}
		}
	}
	return rank
}

func countCards(hand types.Hand) map[int]int {
	var handCount = make(map[int]int)
	for i := 0; i < len(hand.Values); i++ {
		if handCount[hand.Values[i]] < 1 {
			handCount[hand.Values[i]] = 1
		} else {
			handCount[hand.Values[i]] = handCount[hand.Values[i]] + 1
		}
	}
	return handCount
}
