package main

import (
    "time"
)

func handler() error {
    time.Sleep(3 * time.Second)
    return nil
}

func main() {
    done := make(chan struct{}, 0)
    errChan := make(chan error, 1)
    go func() {
        select {
        case <-done: {
            return
        }
        case errChan<- handler(): {
            return
        }
        }
    }()
    time.Sleep(1 * time.Second)
    close(done)
    close(errChan)

    time.Sleep(5 * time.Second)
}
