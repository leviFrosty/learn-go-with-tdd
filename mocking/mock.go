package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var finalWord string = "Go!"
var countdownStart int = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		s.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
