package main

import (
    "time"
)

func main() {
	alertCh := make(chan struct{})

	// sleep 5 mins, and trigger panic
	go func() {
		time.Sleep(5 *  time.Second)
		close(alertCh)
		panic("In GOROUTINE. For Alert Test: Made by songruyang@bytedance.com")
	}()

	select {
	case <-alertCh:
		panic("For Alert Test: Made by songruyang@bytedance.com")
	}
}
