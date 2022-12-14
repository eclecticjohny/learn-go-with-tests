package main

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// make a channel to receive the result
	ch := make(chan struct{})
	// start a goroutine to make the HTTP request
	go func() {
		// make the HTTP request and close the channel
		http.Get(url)
		// signal that the request has been completed
		close(ch)
	}()
	// return the channel
	return ch
}
