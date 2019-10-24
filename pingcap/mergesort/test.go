package main

import (
    "fmt"
    "flag"
    "log"
    "os"
    "runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }
    testData := []int64{4, 6, 2, 7, 5}
    MergeSort(testData)
    fmt.Println(testData)
}
