package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var config = BuildFromEnv()

func check(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	params := req.URL.Query()
	artist := params.Get("artist")
	title := params.Get("title")
	log.Printf("search for %s-%s", artist, title)
	r, err := config.Find(artist, title)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	rb, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(rb)
}

func lyric(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	log.Printf("get %s", id)
	r, err := config.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Write(r)
}

func main() {
	http.HandleFunc("/search", check)
	http.HandleFunc("/lyric", lyric)

	http.ListenAndServe(":8000", nil)
}
