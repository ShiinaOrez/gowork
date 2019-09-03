package client

import (
    "net/rpc"
    "fmt"
    "log"
    "github.com/ShiinaOrez/gowork/rpc/safe-hello-world/constvar"
)

type HelloServiceClient struct {
    *rpc.Client
}

func (client *HelloServiceClient) Hello(request string, reply *string) error {
    return client.Client.Call(constvar.HelloServiceName+".Hello", request, reply)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
    c, err := rpc.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return &HelloServiceClient{Client: c}, nil
}

func StartClient(ch *chan struct{}) {
    _ = <-(*ch)
    close(*ch)
    // fmt.Printf("Dialing...")
    client, err := DialHelloService("tcp", "localhost:2333")
    if err != nil {
        log.Fatal("Dialing:", err)
    }
    fmt.Println("OK")

    reply := ""
    err = client.Hello("", &reply)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(reply)
}
