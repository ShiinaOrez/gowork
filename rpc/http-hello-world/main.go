package main

import (
	_ "github.com/ShiinaOrez/gowork/rpc/http-hello-world/client"
	"github.com/ShiinaOrez/gowork/rpc/http-hello-world/server"
)

func main() {
	server.StartServer()
}
