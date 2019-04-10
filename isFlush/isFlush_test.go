package isflush

import (
	getvalues "poker/getValues"
	"testing"
)

func TestIsFlush(t *testing.T) {
	var vals = getvalues.GetValues("10s Qs 8s 9s 8s")
	var v = IsFlush(vals)
	if v != "Flush" {
		t.Error("Expected Flush, got ", v)
	}
}

func TestIsRoyalFlush(t *testing.T) {
	var vals = getvalues.GetValues("10s Qs Js Ks As")
	var v = IsFlush(vals)
	if v != "Royal Flush" {
		t.Error("Expected Royal Flush, got ", v)
	}
}

func TestIsStraightFlush(t *testing.T) {
	var vals = getvalues.GetValues("10s Qs 8s 9s Js")
	var v = IsFlush(vals)
	if v != "Straight Flush" {
		t.Error("Expected Straight Flush, got ", v)
	}
}
