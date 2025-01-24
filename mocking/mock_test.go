package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("writes desired output", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Countdown(&buffer, &SpyCountdownOperations{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("want '%s' got '%s'", got, want)
		}
	})

	t.Run("sleeps after each call", func(t *testing.T) {
		spy := &SpyCountdownOperations{}
		Countdown(spy, spy)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		got := spy.Calls

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want calls '%q' got calls '%q'", want, got)
		}
	})

}
