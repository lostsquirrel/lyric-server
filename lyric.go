package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const lyricSuffix = "lrc"

type FindResult struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Id     string `json:"id"`
}

type Lyric struct {
	Lyric string `json:"lyric"`
	Id    string `json:"id"`
}

func (cfg *Config) Find(artist, title string) (*FindResult, error) {
	fileName := fmt.Sprintf("%s-%s.%s", artist, title, lyricSuffix)
	filePath := filepath.Join(cfg.LyricsPath, fileName)
	_, err := os.Stat(filePath)

	if err != nil {
		log.Printf("%s not found", fileName)
		return nil, err
	}
	r := FindResult{
		Artist: artist,
		Title:  title,
		Id:     fileName,
	}
	return &r, nil
}

func (cfg *Config) Get(id string) ([]byte, error) {
	filePath := filepath.Join(cfg.LyricsPath, id)
	return ioutil.ReadFile(filePath)
}
