package poker

import (
	"sort"
	"strings"
)

type Hand struct {
	Values []int
	Suits  []string
}

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

func IsStraight(hand Hand) string {
	var isStraightWithAce = Contains(hand.Values, 2)
	if isStraightWithAce == true {
		if Contains(hand.Values, 14) {
			hand.Values = handleAce(hand)
		}
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

func handleAce(hand Hand) []int {
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

func IsOfAKind(hand Hand) string {
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

func countCards(hand Hand) map[int]int {
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

func IsFlush(hand Hand) string {
	var rank = ""
	for i := 0; i < len(hand.Suits)-1; i++ {
		if hand.Suits[i] == hand.Suits[i+1] {
			continue
		} else {
			return "next"
		}
	}
	rank = "Flush"
	if IsStraight(hand) == "Straight" {
		rank = "Straight Flush"
	}
	if hand.Values[0] >= 10 {
		rank = "Royal Flush"
	}
	return rank
}

func InvalidHand(hand Hand) string {
	if len(hand.Values) != 5 {
		return "Invalid Hand"
	}
	if Contains(hand.Values, 0) {
		return "Invalid Hand"
	}
	return "Valid Hand"
}

func GetValues(hand string) Hand {
	var cards = strings.Split(hand, " ")
	var newHand Hand
	newHand.Values = []int{}
	newHand.Suits = []string{}
	for i, x := range cards {
		i = i
		newHand.Values = append(newHand.Values, deck[x[:1]])
		newHand.Suits = append(newHand.Suits, x[len(x)-1:])
	}
	sort.Slice(newHand.Values, func(i, j int) bool { return newHand.Values[i] < newHand.Values[j] })
	return Hand{newHand.Values, newHand.Suits}
}
