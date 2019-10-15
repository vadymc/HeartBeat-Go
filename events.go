package main

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	EventType string `json:"eventType"`
	Series Series `json:"series,omitempty"`
	Movie Movie `json:"movie,omitempty"`
	Episodes []Episode `json:"episodes,omitempty"`
}

type Series struct {
	Title string `json:"title"`
}

type Movie struct {
	Title string `json:"title"`
}

type Episode struct {
	EpisodeNumber int `json:"episodeNumber"`
	SeasonNumber int `json:"seasonNumber"`
}

func ParseEvent(b []byte) (string, string) {
	e := &Event{}
	err := json.Unmarshal(b, e)
	if err != nil {
		return "", string(b)
	}

	var body string
	//title
	if e.Series.Title != "" {
		body = fmt.Sprintf("%v", e.Series.Title)
	} else {
		body = fmt.Sprintf("%v", e.Movie.Title)
	}

	//series
	if len(e.Episodes) != 0 {
		body = fmt.Sprintf("%v S%vE%v", body, e.Episodes[0].SeasonNumber, e.Episodes[0].EpisodeNumber)
	}
	return e.EventType, body
}