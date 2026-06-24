package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	userPort := flag.Int("port", 8080, "port to connect on")
	flag.Parse()

	address := fmt.Sprintf("localhost:%d", *userPort)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Error connecting to server at %s: %v", address, err)
	}

	defer conn.Close()
	fmt.Println("Successfully connected to the server. Type your message below: ")

	//background goroutine to continously read from the server
	go func() {
		serverReader := bufio.NewReader(conn)

		for {
			serverResponse, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println("Server Disconnected")
				os.Exit(0) //client exits if server disconnected
			}

			fmt.Println(serverResponse)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading user std input: %v", err)
			continue
		}

		_, err = conn.Write([]byte(userInput))
		if err != nil {
			log.Fatalf("Error writing to server: %v", err)
		}

	}

}
