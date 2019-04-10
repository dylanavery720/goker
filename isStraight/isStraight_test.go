package isstraight

import (
	getvalues "poker/getValues"
	"testing"
)

func TestIsStraight(t *testing.T) {
	var vals = getvalues.GetValues("5s 4s 7s 6h 8d")
	var v = IsStraight(vals)
	if v != "Straight" {
		t.Error("Expected Straight, got ", v)
	}
}

func TestIsStraightHandleAce(t *testing.T) {
	var vals = getvalues.GetValues("As 2s 3s 4h 5d")
	var v = IsStraight(vals)
	if v != "Straight" {
		t.Error("Expected Straight, got ", v)
	}
}

func TestIsStraightHandleAce2(t *testing.T) {
	var vals = getvalues.GetValues("As 10s Js Qh Kd")
	var v = IsStraight(vals)
	if v != "Straight" {
		t.Error("Expected Straight, got ", v)
	}
}
