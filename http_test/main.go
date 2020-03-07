package main

import (
    "net/http"
    "fmt"
)

func main() {
    client := http.Client{}
    request, _ := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login", nil)
    resp, err := client.Do(request)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp)
}
