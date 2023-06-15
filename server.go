package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var config = BuildFromEnv()

func search(w http.ResponseWriter, req *http.Request) {
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

func lyric(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	id := params.Get("id")
	log.Printf("downloading %s ", id)
	r, err := config.Get(id)
	if err != nil {
		log.Printf("get content for : %s, %s", id, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(r)
}

func main() {
	http.HandleFunc("/search", search)
	http.HandleFunc("/lyric", lyric)

	http.ListenAndServe(":8000", nil)
}
