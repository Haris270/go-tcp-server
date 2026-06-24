package main

import (
	"bufio"
	"net"
)

/*
	msg_receiver func() creates a bufio Reader and reads the Client message using ReadString method. It
	returns the read string and error (if any)
*/

func msg_receiver(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)

	client_message, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return client_message, nil
}
