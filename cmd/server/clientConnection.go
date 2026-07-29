package main

import (
	"net"
	"sync"
)

type ClientRegistry struct {
	mu      sync.RWMutex
	connMap map[string]Client
}

func NewClientRegistry() *ClientRegistry {

	return &ClientRegistry{connMap: make(map[string]Client)}
}

func (c *ClientRegistry) addClient(conn net.Conn, clientName string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.connMap[conn.RemoteAddr().String()] = Client{conn, conn.RemoteAddr().String(), clientName}
}

func (c *ClientRegistry) removeClient(conn net.Conn, clientName string) {
	defer conn.Close()
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.connMap, clientName)

}

// func (c *ClientRegistry) getClient(clientID string) Client {
// 	return c.connMap[clientID]
// }

func (c *ClientRegistry) getAllClients() []Client {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var allClients []Client

	for _, client := range c.connMap {
		allClients = append(allClients, client)
	}

	return allClients
}
