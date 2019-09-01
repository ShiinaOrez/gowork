package server

import (
    "net"
    "net/rpc"
    "fmt"
    "log"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "Hello: " + request
    return nil
}

func StartServer() {
    rpc.RegisterName("HelloService", new(HelloService))

    fmt.Printf("Listening...")
    listener, err := net.Listen("tcp", ":2333")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }
    fmt.Println("OK")

    fmt.Printf("Accepting...")
    conn, err := listener.Accept()
    if err != nil {
        log.Fatal()
    }
    fmt.Println("OK")

    rpc.ServeConn(conn)
}
