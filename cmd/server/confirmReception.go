package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func confirmReception(conn net.Conn, connMap map[string]net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Println("Error encountered during reading input")
			fmt.Println("Client Disconnected")
			delete(connMap, conn.RemoteAddr().String())
			return
		}
		fmt.Printf("Client: %s\n", clientMessage)
		clientMessage = strings.TrimSpace(clientMessage)
		var confirmationMsg string
		if strings.EqualFold(clientMessage, "all clients") {
			var allKeys strings.Builder

			for key := range connMap {
				allKeys.WriteString(key)
			}

			result := allKeys.String()
			confirmationMsg = fmt.Sprintf("Result: %s\n", result)
		} else {
			confirmationMsg = fmt.Sprintf("COnfirmed, you sent %s\n", clientMessage)
		}

		_, err = conn.Write([]byte(confirmationMsg)) //cast the response string to []byte
		if err != nil {
			//fmt.Println("Error encountered while Writing output")
			fmt.Println("Client Disconnected")
			delete(connMap, conn.RemoteAddr().String())
			return
		}
	}

}
