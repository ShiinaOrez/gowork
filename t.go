package main

import (
    "gowork/muxibook/models"
    "fmt"
)

func main(){
    var a models.User
    a.Username="Shiina"
    fmt.Println(a.Username)
}
