package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {

	userPort := flag.Int("port", 8080, "Port to connect server to: ")
	flag.Parse()

	address := fmt.Sprintf("localhost:%d", *userPort)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error creating the Listener")
	}

	defer listener.Close() // close the listener when program exits main()

	fmt.Printf("Server successfully listening on: %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in accepting Connection from Client.")
			panic(err)
		}
		go confirmReception(conn)
	}
}
