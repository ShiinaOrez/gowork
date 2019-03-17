package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    GetTest()
}

func GetTest() {
    u, err := url.Parse("http://127.0.0.1:8080")
    if err != nil {
        fmt.Println("Parse URL: localhost:8080 FAILED!")
        return
    }
    response, err := http.Get(u.String())
    if err != nil {
        fmt.Println("Request localhost:8080 FAILED!")
        return
    }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Read Body FAILED!")
        return
    }
    if string(body) == "Hello, MUXI." {
        fmt.Println("Check Successful!")
    } else {
        fmt.Println("Please Check your Code!")
    }
    return
}
