package main

import (
	"fmt"
	"log"
	"net/http"

	isflush "poker/isFlush"
	isofakind "poker/isOfAKind"
	isstraight "poker/isStraight"

	getvalues "poker/getValues"
)

func handler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["hand"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'hand' is missing")
		return
	}
	key := keys[0]
	var rank = ""
	var vals = getvalues.GetValues(key)
	rank = isflush.IsFlush(vals)
	if rank != "next" {
		fmt.Fprintf(w, "%q\n", rank)
		return
	}
	rank = isstraight.IsStraight(vals)
	if rank != "next" {
		fmt.Fprintf(w, "%q\n", rank)
		return
	}
	rank = isofakind.IsOfAKind(vals)
	fmt.Fprintf(w, "%q\n", rank)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
