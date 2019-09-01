package main

import (
	"fmt"
	"sync"
	"time"
)

func gen() <-chan int {
	out := make(chan int)
	go func() {
		// range array, or enum
		for n:=0; n<=1000000; n++ {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// range channel
		for n := range in {
			out <- n*n
		}
		close(out)
	}()
	// channel returned, but data will be producted...
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// funcname: output
	// receives: the channel
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

    for _, c := range cs {
    	wg.Add(1)
    	go output(c)
    }

    go func() {
    	wg.Wait()
    	close(out)
    }()
    return out
}

func main() {
	t1 := time.Now()
	// get a in channel
	in := gen()

    // fan-in: from one channel to many function(or stage)
    channel1 := sq(in)
    channel2 := sq(in)
    
    // fan-out: from many channel to one channel out
    for n := range merge(channel1, channel2) {
    	fmt.Println(n)
    }
    t2 := time.Now()
    fmt.Println(t2.Sub(t1))
}