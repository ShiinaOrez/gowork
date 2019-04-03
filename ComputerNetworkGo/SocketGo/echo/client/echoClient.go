package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

var host = flag.String("host", "localhost", "host") // host by flag
var port = flag.String("port", "3333", "port")      // port by flag

func main() {
	flag.Parse() // flag parse
	conn, err := net.Dial("tcp", *host + ":" + *port) // connect to host:port by tcp4&tcp6
	checkError(err) // checkError
	defer conn.Close() // makesure close()
	fmt.Println("Connecting to " + *host + ":" + *port)

	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg) // write complicated
	go handleRead(conn, &wg)  // read complicated

	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for i:=10; i>0; i-- {
		_, err := conn.Write([]byte("Hello " + strconv.Itoa(i) + "\r\n"))
		fmt.Println("Hello", i)
		if err != nil {
			fmt.Println("Error to send message because of", err.Error())
			break
		}
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	reader := bufio.NewReader(conn)
	for i:=1; i<=10; i++ {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Println("Error to read message because of", err.Error())
			return
		}
		fmt.Println(line[:len(line)-1])
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}