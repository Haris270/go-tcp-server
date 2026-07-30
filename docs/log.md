# Project Development Log
---
**06/24/26**

## Server

   **Main.go**
- creates a TCP listener using Go’s Net.Listen() on address (localhost:8080 by default)
- Creates a map to store all connected clients
- In the loop, launches a goroutine to handle each client connection

	**ConfirmReception**
- confirmReception() initiates the Client Connection by getting their name and maintains a registry of active client connections (connMap). 
- It sanitizes the userName by removing the newline delimiter from the name.
- Within the for loop, the func calls the msg_receiver() and broadcast() funcs to receive client message and
  broadcasts the formatted message (display_msg) it to all clients (including the sender) respectively.

	**msg_receiver.go**
- msg_receiver func() creates a bufio Reader and reads the Client message using ReadString method. 
- It returns the read string and error (if any)

	**broadcast.go**
- broadcast func() iterates over the connMap and encodes the client message (byte[]) before writing to
  all clients (including the sender as well)
- If it fails to write, it returns the error, nil otherwise
