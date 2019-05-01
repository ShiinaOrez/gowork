package main

import (
    "os"
    "fmt"
    "time"
    "github.com/garyburd/redigo/redis"
)

func Dial() (redis.Conn, error) {
    conn, err := redis.Dial("tcp", os.Getenv("Elite_RedisServer"))
    if err != nil {
        return nil, err
    }
    if _, err := conn.Do("AUTH", os.Getenv("Elite_RedisPassword")); err != nil {
        conn.Close()
        return nil, err
    }
    return conn, nil
}

func main() {
    c, err := Dial()
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
