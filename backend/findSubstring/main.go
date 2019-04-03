package main

import (
    "fmt"
    "strings"
    "errors"
)

type ss string

func Find(stra, strb string) (int, error) {
    position := strings.Index(stra, strb)
    if position < 0 {
        return -1, errors.New("Substring not found!")
    }
    return position, nil
}

func (stra *ss) Find(strb string) (int, error) {
    position := strings.Index(string(*stra), strb)
    if position < 0 {
        return -1, errors.New("Substring not found!")
    }
    return position, nil
}

func (stra *ss) FindAll(strb string) ([]int, error) {
    res :=  []int{}
    base := 0
    for strings.Index(string(*stra)[base:], strb) >= 0 {
        position := base + strings.Index(string(*stra)[base:], strb)
        res = append(res, position)
        base = position + len(strb)
    }
    if len(res) == 0 {
        return res, errors.New("Substring not found!")
    }
    return res, nil
}

func test1() {
    var stra, strb string
    stra = "abcabcacbabc"
    strb = "bc"
    res, err := Find(stra, strb)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%s in %s is: %d\n", strb, stra, res)
}

func test2() {
    var stra, strb ss
    stra = "abcabcacbabc"
    strb = "ca"
    res, err := stra.Find(string(strb))
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%s in %s is: %d\n", strb, stra, res)
}

func test3() {
    var stra, strb ss
    stra = "abcabcacbabc"
    strb = "abc"
    res, err := stra.FindAll(string(strb))
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%s in %s is: ", strb, stra)
    fmt.Println(res)
}

func main() {
    test1()
    test2()
    test3()
}
