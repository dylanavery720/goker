package isflush

import types "poker/types"

func IsFlush(hand types.Hand) string {
	var rank = ""
	for i := 0; i < len(hand.Suits)-1; i++ {
		if hand.Suits[i] == hand.Suits[i+1] {
			continue
		} else {
			return "false"
		}
	}
	rank = "Flush"
	return rank
}
