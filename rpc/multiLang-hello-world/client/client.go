package client

import (
	"fmt"
	"github.com/ShiinaOrez/gowork/rpc/safe-hello-world/constvar"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

func (client *HelloServiceClient) Hello(request string, reply *string) error {
	return client.Client.Call(constvar.HelloServiceName+".Hello", request, reply)
}

func StartClient(ch *chan struct{}) {
	_ = <-(*ch)
	close(*ch)
	fmt.Printf("Dialing...")
	conn, err := net.Dial("tcp", "localhost:2333")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	fmt.Println("OK")

	reply := ""
	err = client.Call(constvar.HelloServiceName+".Hello", "multiLang-rpc", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
