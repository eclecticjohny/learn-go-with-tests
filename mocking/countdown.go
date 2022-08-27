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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
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

func (d *ConfigurableSleeper) Sleep() {
	d.sleep(d.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
