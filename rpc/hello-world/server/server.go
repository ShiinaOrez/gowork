package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello: " + request
	return nil
}

func StartServer() *chan struct{} {
	ch := make(chan struct{})
	go func() {
		rpc.RegisterName("HelloService", new(HelloService))

		fmt.Printf("Listening...")
		listener, err := net.Listen("tcp", ":2333")
		if err != nil {
			log.Fatal("ListenTCP error:", err)
		}
		fmt.Println("OK")
		ch <- struct{}{}
		fmt.Printf("Accepting...")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal()
		}
		fmt.Println("OK")

		rpc.ServeConn(conn)
	}()
	return &ch
}
