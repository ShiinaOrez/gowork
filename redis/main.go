package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func main() {
    c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("Connection failed.", err.Error())
    }
    defer c.Close()

    _, err = c.Do("SET", "first", "1")
    if err != nil {
        fmt.Println("Set key-value failed.", err.Error())
    }

    num, err := redis.String(c.Do("GET", "first"))
    if err != nil {
        fmt.Println("Get value failed.", err.Error())
    } else {
        fmt.Println("The first is:", num)
    }
}
