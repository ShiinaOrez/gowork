package main

import (
    "github.com/ShiinaOrez/gowork/rpc/hello-world/client"
    "github.com/ShiinaOrez/gowork/rpc/hello-world/server"
)

func main() {
    ch := server.StartServer()
    client.StartClient(ch)
}
