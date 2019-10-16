package main

import (
	"testing"
	"io/ioutil"
)

func TestEventParsing(t *testing.T) {
	testData := []struct {
		T string
		B string
		P string
	}{
		{"Download", "Gravity Falls S2E14", "test_data/sonarr_single_episode.json"},
		{"Download", "Gravity Falls S2E14", "test_data/sonarr_two_episodes.json"},
		{"Download", "Finding Nemo", "test_data/radarr_payload.json"},
	}

	for _, td := range testData {
		assertParsing(t, td.T, td.B, td.P)
	}
}

func assertParsing(t *testing.T, expectedTitle string, expectedBody string, testDataPath string) {
	b, _ := ioutil.ReadFile(testDataPath)
	title, body := ParseEvent(b)
	if title != expectedTitle || body != expectedBody {
		t.Errorf("Failed. Expected %v %v, actual %v %v", expectedTitle, expectedBody, title, body)
	}
}