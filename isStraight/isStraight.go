package isstraight

import (
	types "poker/types"
	"sort"
)

func IsStraight(hand types.Hand) string {
	var isStraightWithAce = Contains(hand.Values, 2)
	if isStraightWithAce == true {
		hand.Values = handleAce(hand)
	}
	for i := 0; i < len(hand.Values)-1; i++ {
		if hand.Values[i] == (hand.Values[i+1] - 1) {
			continue
		} else {
			return "next"
		}
	}
	return "Straight"
}

func handleAce(hand types.Hand) []int {
	var indexOfAce = sort.IntSlice(hand.Values).Search(14)
	hand.Values[indexOfAce] = 1
	sort.Slice(hand.Values, func(i, j int) bool { return hand.Values[i] < hand.Values[j] })
	return hand.Values
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
