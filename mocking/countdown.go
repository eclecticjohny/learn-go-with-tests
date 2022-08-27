package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const (
	finalWord      = "Go!"
	countDownStart = 3
)

func Countdown(w io.Writer, s Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
