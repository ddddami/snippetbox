package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2024, 1, 1, 13, 30, 0, 0, time.UTC)
	hd := humanDate(tm)

	if hd != "01 Jan 2024 at 13:30" {
		t.Errorf("got %q; want %q", hd, "01 Jan 2024 at 13:30")
	}
}
