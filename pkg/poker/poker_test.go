package poker

import (
	"testing"
)

func TestThreeOfAKind(t *testing.T) {
	var vals = GetValues("2s Kh Kd Kc Js")
	var v = IsOfAKind(vals)
	if v != "Three Of A Kind" {
		t.Error("Expected Three Of A Kind, got ", v)
	}
}

func TestFourOfAKind(t *testing.T) {
	var vals = GetValues("10s 8s 8s 8s 8s")
	var v = IsOfAKind(vals)
	if v != "Four Of A Kind" {
		t.Error("Expected Four Of A Kind, got ", v)
	}
}

func TestOnePair(t *testing.T) {
	var vals = GetValues("10s 3s 4s 8s 8s")
	var v = IsOfAKind(vals)
	if v != "One Pair" {
		t.Error("Expected One Pair, got ", v)
	}
}

func TestFullHouse(t *testing.T) {
	var vals = GetValues("10s 10s 8s 8s 8s")
	var v = IsOfAKind(vals)
	if v != "Full House" {
		t.Error("Expected Full House, got ", v)
	}
}

func TestTwoPair(t *testing.T) {
	var vals = GetValues("10s 10s 7s 8s 8s")
	var v = IsOfAKind(vals)
	if v != "Two Pair" {
		t.Error("Expected Two Pair, got ", v)
	}
}

func TestFlush(t *testing.T) {
	var vals = GetValues("10s 10s 7s 8s 8s")
	var v = IsFlush(vals)
	if v != "Flush" {
		t.Error("Expected Flush, got ", v)
	}
}

func TestStraightFlush(t *testing.T) {
	var vals = GetValues("10s 9s 7s 8s 6s")
	var v = IsFlush(vals)
	if v != "Straight Flush" {
		t.Error("Expected Straight Flush, got ", v)
	}
}

func TestRoyalFlush(t *testing.T) {
	var vals = GetValues("10s Js Qs Ks As")
	var v = IsFlush(vals)
	if v != "Royal Flush" {
		t.Error("Expected Royal Flush, got ", v)
	}
}

func TestEdgeCaseWithLowStraight(t *testing.T) {
	var vals = GetValues("2d 3d 4d 5d 6d")
	var v = IsFlush(vals)
	if v != "Straight Flush" {
		t.Error("Expected Straight Flush, got ", v)
	}
}

func TestEdgeCaseWithLowAceStraight(t *testing.T) {
	var vals = GetValues("2d 3d 4d 5d Ad")
	var v = IsFlush(vals)
	if v != "Straight Flush" {
		t.Error("Expected Straight Flush, got ", v)
	}
}
