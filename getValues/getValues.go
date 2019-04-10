package getvalues

import (
	types "poker/types"
	"sort"
	"strings"
)

var deck = map[string]int{
	"A": 14,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"1": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
}

func GetValues(hand string) types.Hand {
	var cards = strings.Split(hand, " ")
	var newHand types.Hand
	newHand.Values = []int{}
	newHand.Suits = []string{}
	for i, x := range cards {
		i = i
		newHand.Values = append(newHand.Values, deck[x[:1]])
		newHand.Suits = append(newHand.Suits, x[len(x)-1:])
	}
	sort.Slice(newHand.Values, func(i, j int) bool { return newHand.Values[i] < newHand.Values[j] })
	return types.Hand{newHand.Values, newHand.Suits}
}
