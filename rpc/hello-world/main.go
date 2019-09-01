package main

import (
    "github.com/ShiinaOrez/gowork/rpc/hello-world/client"
    "github.com/ShiinaOrez/gowork/rpc/hello-world/server"
)

func main() {
    go server.StartServer()
    client.StartClient()
}
