package main

/*
	broadcast func() iterates over the connMap and writes the client message (converted to byte[]) to
	all clients (including the sender as well)

	If it fails to write, it returns the error, nil otherwise
*/

func (c *ClientRegistry) broadcast(msg string) []Client {

	allClients := c.getAllClients()
	var failedWrites []Client
	for _, client := range allClients {

		//client.Connection.SetWriteDeadline(time.Now().Add(5 * time.Second))
		_, err := client.Connection.Write([]byte(msg))
		if err != nil {
			failedWrites = append(failedWrites, client)
		}
	}

	return failedWrites

}

// c.mu.RLock()
// defer c.mu.RUnlock()
//var wg sync.WaitGroup

// 	for _, client := range c.connMap {

// 		wg.Add(1)

// 		go func(conn net.Conn) error{
// 			defer wg.Done()

// 			conn.SetWriteDeadline()
// 			_, err := conn.Write([]byte(msg))
// 			if err != nil {
// 				return err
// 			}

// 		}(client.Connection)
// 		//_, err := val.Connection.Write([]byte(msg)) //cast the response string to []byte

// 	}

// 	return nil

// }
