package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(&buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("want '%s' got '%s'", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("Sleep was called incorrect amount of times, got '%d' want '%d'", spySleeper.Calls, 3)
	}
}
