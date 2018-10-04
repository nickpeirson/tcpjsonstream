package main

import (
	"bufio"
	"fmt"
	"github.com/bcicen/jstream"
	"github.com/pquerna/ffjson/ffjson"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "61222"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		//go printRequest(conn)
		//go parseRequest(conn)
		//go parseRequestFfjson(conn)
		go parseRequestBufferedFfjson(conn)
	}

}

func printRequest(conn net.Conn) {
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print(string(message))
	}
}

func parseRequest(conn net.Conn) {
	decoder := jstream.NewDecoder(conn, 0) // extract JSON values at a depth level of 1
	for mv := range decoder.Stream() {
		fmt.Printf("%v\n ", mv.Value.(map[string]interface{}))
	}
}

func parseRequestFfjson(conn net.Conn) {
	dec := ffjson.NewDecoder()

	var v map[string]interface{}

	for {
		dec.DecodeReader(conn, &v)
		fmt.Printf("%v\n ", v)
	}
}

func parseRequestBufferedFfjson(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	var v map[string]interface{}

	for scanner.Scan() {
		ffjson.Unmarshal(scanner.Bytes(), &v)
		fmt.Printf("%v\n ", v)
	}
}
