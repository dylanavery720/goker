package poker

import (
	"testing"
)

func TestThreeOfAKind(t *testing.T) {
	var vals = GetValues("10s Qs 8s 8s 8s")
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
