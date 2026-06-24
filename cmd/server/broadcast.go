package main

/*
	broadcast func() iterates over the connMap and writes the client message (converted to byte[]) to
	all clients (including the sender as well)

	If it fails to write, it returns the error, nil otherwise
*/

func broadcast(connMap map[string]Client, msg string) error {
	for _, val := range connMap {
		_, err := val.Connection.Write([]byte(msg)) //cast the response string to []byte
		if err != nil {
			return err
		}

	}

	return nil

}
