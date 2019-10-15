package main

import (
	"testing"
	"io/ioutil"
)

func TestEventParsing(t *testing.T) {
	expectedTitle := "Download"
	expectedSeriesBody := "Gravity Falls S2E14"
	b, _ := ioutil.ReadFile("test_data/sonarr_single_episode.json")

	title, body := ParseEvent(b)
	if title != expectedTitle || body != expectedSeriesBody {
		t.Errorf("failed parsing single episode json. Expected %v %v, actual %v %v", expectedTitle, expectedSeriesBody, title, body)
	}

	b, _ = ioutil.ReadFile("test_data/sonarr_two_episodes.json")
	title, body = ParseEvent(b)
	if title != expectedTitle || body != expectedSeriesBody  {
		t.Errorf("failed parsing two episodes json. Expected %v %v, actual %v %v", expectedTitle, expectedSeriesBody, title, body)
	}

	expectedMovieBody := "Finding Nemo"
	b, _ = ioutil.ReadFile("test_data/radarr_payload.json")
	title, body = ParseEvent(b)
	if title != expectedTitle || body != expectedMovieBody {
		t.Errorf("failed parsing movie json. Expected %v %v, actual %v %v", expectedTitle, expectedMovieBody, title, body)
	}
}