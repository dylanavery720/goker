package main

import (
	"fmt"
	"log"
	"net/http"

	poker "poker/pkg/poker"
)

func handler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["hand"]
	if !ok || len(keys[0]) < 1 {
		return
	}
	key := keys[0]
	var rank = ""
	var vals = poker.GetValues(key)
	var validated = poker.InvalidHand(vals)
	if validated == "Invalid Hand" {
		fmt.Fprintf(w, "%q\n", validated)
		return
	}
	rank = poker.IsFlush(vals)
	if rank != "next" {
		fmt.Fprintf(w, "%q\n", rank)
		return
	}
	rank = poker.IsStraight(vals)
	if rank != "next" {
		fmt.Fprintf(w, "%q\n", rank)
		return
	}
	rank = poker.IsOfAKind(vals)
	fmt.Fprintf(w, "%q\n", rank)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}