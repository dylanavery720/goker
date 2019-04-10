// if (hand.values.length != 5 || hand.values.includes(undefined)) return 'Invalid Hand'

package invalidhand

import (
	getvalues "poker/getValues"
	"testing"
)

func TestShouldHold5Cards(t *testing.T) {
	var vals = getvalues.GetValues("5s 4s 7s 6h")
	var v = InvalidHand(vals)
	if v != "Invalid Hand" {
		t.Error("Expected Invalid Hand, got ", v)
	}
}

func TestShouldNotExceptStrangeCards(t *testing.T) {
	var vals = getvalues.GetValues("15s Fs 7s 6h 6h")
	var v = InvalidHand(vals)
	if v != "Invalid Hand" {
		t.Error("Expected Invalid Hand, got ", v)
	}
}
