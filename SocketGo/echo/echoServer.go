package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "3333", "port")

func main () {
	flag.Parse()
	l, err := net.Listen("tcp", *host+":"+*port)
	checkError(err)
	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)

	for {
		conn, err := l.Accept()
		checkError(err)
		fmt.Println("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		io.Copy(conn, conn)
	}
}