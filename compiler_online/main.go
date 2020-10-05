package main

import (
    "github.com/ShiinaOrez/gowork/compiler_online/api"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.POST("/post", api.RunCode)
    router.Run(":1666")
}