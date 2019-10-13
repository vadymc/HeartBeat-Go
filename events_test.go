package main

import (
	"testing"
)

func TestEventParsing(t *testing.T) {
	header := "[header]"
	body := "body"
	h, b := parseEvent(header + body)
	if h != header {
		t.Errorf("Header parsed incorrectly. Expected %v got %v\n", header, h)
	}
	if b != body {
		t.Errorf("Body parsed incorrectly. Expected %v got %v\n", body, b)
	}

	header = ""
	body = "body"
	h, b = parseEvent(header + body)
	if h != header {
		t.Errorf("Header parsed incorrectly. Expected %v got %v\n", header, h)
	}
	if b != body {
		t.Errorf("Body parsed incorrectly. Expected %v got %v\n", body, b)
	}
}
