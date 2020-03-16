package main

import (
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(1)
    go func() {
        panic("already call")
    }()
    for {}
}
