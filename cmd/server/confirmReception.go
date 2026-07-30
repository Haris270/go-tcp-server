package main

import (
	//"bufio"
	"fmt"
	"net"
	"strings"
)

func confirmReception(conn net.Conn, connMap *ClientRegistry) { //map[string]Client) {
	defer conn.Close()

	//---------- initiating the connection -------------\\

	// val, ok := connMap.getClient(conn.RemoteAddr().String())//connMap[conn.RemoteAddr().String()]
	// var welcome_msg string
	// if ok {
	// 	var name string = val.Name
	// 	welcome_msg = fmt.Sprintf("Server: Welcome Back %s\n!", name)
	// } else {

	welcome_msg := "Welcome! Enter Your Name: \n"
	// }
	_, err := conn.Write([]byte(welcome_msg))
	if err != nil {
		fmt.Println("Error while forming connection")
		return
	}

	clientName, err := msg_receiver(conn)
	if err != nil {
		fmt.Println("Error getting the name. Disconnecting...")
		return
	}
	clientName = strings.TrimSuffix(clientName, "\n")
	//connMap[conn.RemoteAddr().String()] = Client{conn, clientName}
	connMap.addClient(conn, clientName)

	join_notif := fmt.Sprintf("%s has joined the chat room!\n", clientName)
	connMap.broadcast(join_notif)

	defer func() {
		connMap.removeClient(conn.RemoteAddr().String())
	}()

	for {

		clientMessage, err := msg_receiver(conn)
		if err != nil {
			fmt.Println("Client Disconnected")
			//delete(connMap, conn.RemoteAddr().String())
			// connMap.removeClient(conn, conn.RemoteAddr().String())
			//connMap.broadcast(fmt.Sprintf("%s has left the chat\n", clientName))
			return
		}

		clientMessage = strings.TrimSpace(clientMessage)
		display_msg := fmt.Sprintf("%s: %s\n", clientName, clientMessage) // ClientName + ClientMessage to be displayed on Terminal

		// Broadcast the Message to all Clients (Including Sender)
		failed_clients := connMap.broadcast(display_msg)

		if len(failed_clients) != 0 {

			for _, client := range failed_clients {
				connMap.removeClient(client)
			}

		}

	}

}

/*
	confirmReception() initiates the Client Connection by getting their name and adding them to the connMap. It
	removes the newline delimiter from the name.

	Within the for loop, the func calls the msg_receiver() and broadcast() funcs to receive client message and
	broadcast the display_msg it to all clients (including the sender) respectively.
*/
