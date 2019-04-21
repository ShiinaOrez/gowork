package main

import (
    "fmt"
    "time"
    "github.com/garyburd/redigo/redis"
)

func main() {
    c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("Connection failed.", err.Error())
    }
    defer c.Close()

    _, err = c.Do("SET", "second", "2", "EX", "6")
    if err != nil {
        fmt.Println("Set key-value failed.", err.Error())
    }

    time.Sleep(8 * time.Second)

    num, err := redis.String(c.Do("GET", "second"))
    if err != nil {
        fmt.Println("Get value failed.", err.Error())
    } else {
        fmt.Println("The first is:", num)
    }
}
