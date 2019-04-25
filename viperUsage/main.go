package main

import (
    "fmt"
    "github.com/spf13/viper"
)

func main() {
    viper.BindEnv("auth", "QQMailAuthCode")
    authCode := viper.Get("auth")
    fmt.Println(authCode)
}
