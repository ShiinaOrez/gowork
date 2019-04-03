package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var host = flag.String("host", "", "host")     //host by flag
var port = flag.String("port", "3333", "port") //port by flag

func main () {
	flag.Parse() //parse flag
	l, err := net.Listen("tcp", *host+":"+*port)  //new listener, use tcp4&tcp6, on host:port
	checkError(err) // checkError, Create Listener Failed
	defer l.Close() // defer close(), if you create a connectionn, please close it in the end
	fmt.Println("Listening on " + *host + ":" + *port)

	for { // listen & listen & listen
		conn, err := l.Accept() // waiting for connect
		checkError(err)
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr()) // actually, just connect but not receive message

		go handleRequest(conn) // handle request by complicated
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
		io.Copy(conn, conn) // send it back
	}
}