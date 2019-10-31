package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) { // recursive end condition
	case 0:
		return nil // if channels = []
	case 1:
		return channels[0] // if channels slice just have one channel, just return it
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]: // if first channel close, orDone will be shut down at next
			case <-channels[1]: // if second channel close, orDone will be shut down at next
			}
		default:
			select {
			case <-channels[0]: // the number case here, decide the channels tree's width.
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...): // recursive definition, to the next level
			}
		}
	}()
	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Println("Done after %v", time.Since(start))
}
