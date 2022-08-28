package main

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	// ping a and b concurrently
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
