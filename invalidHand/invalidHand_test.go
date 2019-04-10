package invalidhand

import (
	getvalues "poker/getValues"
	"testing"
)

// func TestInvalidHand(t *testing.T) {
// 	var vals = getvalues.GetValues("10s 4s Hs 9s")
// 	var v = InvalidHand(vals)
// 	if v != "Invalid Hand" {
// 		t.Error("Expected Invalid Hand, got ", v)
// 	}
// }

func TestInvalidCards(t *testing.T) {
	var vals = getvalues.GetValues("11s 45s Hs 9s 9d")
	var v = InvalidHand(vals)
	if v != "Invalid Hand" {
		t.Error("Expected Invalid Hand, got ", v)
	}
}
