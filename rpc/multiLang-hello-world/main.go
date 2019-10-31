package main

import (
	"github.com/ShiinaOrez/gowork/rpc/multiLang-hello-world/client"
	_ "github.com/ShiinaOrez/gowork/rpc/multiLang-hello-world/server"
)

func main() {
	// ch := server.StartServer()
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	client.StartClient(&ch)
}
