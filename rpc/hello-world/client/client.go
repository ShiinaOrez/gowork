package client

import (
    "net/rpc"
    "fmt"
    "log"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "Hello: " + request
    return nil
}

func StartClient() {
    fmt.Printf("Dialing...")
    client, err := rpc.Dial("tcp", "localhost:2333")
    if err != nil {
        log.Fatal("Dialing:", err)
    }
    fmt.Println("OK")

    reply := ""
    err = client.Call("HelloService.Hello", "rpc", &reply)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(reply)
}
