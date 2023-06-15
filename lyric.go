package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const lyricSuffix = "lrc"

type FindResult struct {
	Artist string `json:"artist"`
	Title  string `json:"song"`
	Id     string `json:"id"`
	Lyric  string `json:"lyrics"`
}

type Lyric struct {
	Lyric string `json:"lyric"`
	Id    string `json:"id"`
}

func (cfg *Config) Find(artist, title string) (*FindResult, error) {
	fileName := fmt.Sprintf("%s - %s.%s", artist, title, lyricSuffix)
	filePath := filepath.Join(cfg.LyricsPath, fileName)
	_, err := os.Stat(filePath)

	if err != nil {
		log.Printf("%s not found", fileName)
		return nil, err
	}
	bc, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("read error: %v", err)
		return nil, err
	}
	r := FindResult{
		Artist: artist,
		Title:  title,
		Id:     string(bc),
		Lyric:  string(bc),
	}
	return &r, nil
}

// func (cfg *Config) Get(id string) ([]byte, error) {
// 	filePath := filepath.Join(cfg.LyricsPath, id)
// 	return os.ReadFile(filePath)
// }
