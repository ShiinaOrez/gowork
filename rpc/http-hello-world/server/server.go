package server

import (
    "net/rpc"
    "net/rpc/jsonrpc"
    "io"
    "net/http"
    "github.com/ShiinaOrez/gowork/rpc/safe-hello-world/constvar"
)

// new interface
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(service HelloServiceInterface) error {
	return rpc.RegisterName(constvar.HelloServiceName, service)
}

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "Hello: " + request
    return nil
}

func StartServer() {
        RegisterHelloService(new(HelloService))

        http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
            var conn io.ReadWriteCloser = struct {
                io.Writer
                io.ReadCloser
            }{
                ReadCloser: r.Body,
                Writer:     w,
            }
            rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
        })
        http.ListenAndServe(":2333", nil)
    return
}
