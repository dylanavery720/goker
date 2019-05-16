package main

import (
	"fmt"
	"log"
	"net/http"

	poker "poker/pkg/poker"
)

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	keys, ok := r.URL.Query()["hand"]
	if !ok || len(keys[0]) < 1 {
		return
	}
	hand := keys[0]
	var rank = ""
	var vals = poker.GetValues(hand)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("server is listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
