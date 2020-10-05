package main

import (
	"runtime"
	"time"
	"sync"
)

var i = sync.Mutex{}

func a() {
	i.Lock()
	i.Unlock()
}

func main() {
// 	i := 0
	runtime.GOMAXPROCS(1)
	go func() { 
		for {
			a()
		}
	}()
	
	time.Sleep(time.Millisecond)
	println("OK")
}

