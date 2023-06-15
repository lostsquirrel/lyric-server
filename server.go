package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var config = BuildFromEnv()

func check(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		log.Printf("Method %s not support", req.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	params := req.URL.Query()
	artist := params.Get("artist")
	title := params.Get("song")
	log.Printf("search for %s - %s", artist, title)
	r, err := config.Find(artist, title)
	if err != nil {
		log.Printf("search error: %s", err)
		w.WriteHeader(http.StatusNotFound)
	}
	rb, err := json.Marshal(r)
	if err != nil {
		log.Printf("json encode error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(rb)
}

func main() {
	http.HandleFunc("/", check)
	// http.HandleFunc("/lyric", lyric)

	http.ListenAndServe(":8000", nil)
}
