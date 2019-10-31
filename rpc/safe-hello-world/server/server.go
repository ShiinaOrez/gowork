package server

import (
	"fmt"
	"github.com/ShiinaOrez/gowork/rpc/safe-hello-world/constvar"
	"log"
	"net"
	"net/rpc"
)

// new interface
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(service HelloServiceInterface) error {
	return rpc.RegisterName(constvar.HelloServiceName, service)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello: " + request
	return nil
}

func StartServer() *chan struct{} {
	ch := make(chan struct{})
	go func() {
		RegisterHelloService(new(HelloService))

		fmt.Printf("Listening...")
		listener, err := net.Listen("tcp", ":2333")
		if err != nil {
			log.Fatal("ListenTCP error:", err)
		}
		fmt.Println("OK")
		ch <- struct{}{}
		fmt.Printf("Accepting...")
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal()
			}
			fmt.Println("OK")

			go rpc.ServeConn(conn)
		}
	}()
	return &ch
}
