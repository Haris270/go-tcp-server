package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	Connection net.Conn
	Name       string
}

func main() {

	userPort := flag.Int("port", 8080, "Port to connect server to: ") //user can call with -port {} to define the port to use. 8080 is the default port.
	flag.Parse()

	address := fmt.Sprintf("localhost:%d", *userPort)

	// create a listener for TCP connections on address (localhost:8080 by default)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error creating the Listener")
	}

	defer listener.Close() // close the listener when program exits main()

	fmt.Printf("Server successfully listening on: %s\n", address)

	allConn := make(map[string]Client) // map to store all connected clients

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in accepting Connection from Client.")
			panic(err)
		}

		//creates a goroutine calling the client handler for each client
		go confirmReception(conn, allConn)
	}
}
