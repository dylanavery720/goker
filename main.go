package main

import (
	"fmt"
	"log"
	"net/http"
	isflush "poker/isFlush"

	getvalues "poker/getValues"
)

func handler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["hand"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'hand' is missing")
		return
	}
	key := keys[0]
	var vals = getvalues.GetValues(key)
	var isFlush = isflush.IsFlush(vals)
	fmt.Fprintf(w, "%q\n", isFlush)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
