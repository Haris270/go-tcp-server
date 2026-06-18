package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func confirmReception(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error encountered during reading input")
			return
		}
		fmt.Printf("Client: %s\n", clientMessage)
		clientMessage = strings.TrimSpace(clientMessage)

		confirmationMsg := fmt.Sprintf("Confirmed: You sent %s\n", clientMessage)

		_, err = conn.Write([]byte(confirmationMsg)) //cast the response string to []byte
		if err != nil {
			fmt.Println("Error encountered while Writing output")
			return
		}
	}

}
