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
	c.connMap[conn.RemoteAddr().String()] = Client{conn, clientName}
}

func (c *ClientRegistry) removeClient(clientAddress string) {
	c.mu.Lock()

	client, ok := c.connMap[clientAddress]
	if ok {
		delete(c.connMap, clientAddress)
	}

	c.mu.Unlock()

	if ok {
		client.Connection.Close()
	}
}

func (c *ClientRegistry) CloseAllClients() {
	allClients := c.getAllClients()

	for _, client := range allClients {
		//c.removeClient(client.Connection.RemoteAddr().String())
		client.Connection.Close()
	}
}

func (c *ClientRegistry) getAllClients() []Client {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var allClients []Client

	for _, client := range c.connMap {
		allClients = append(allClients, client)
	}

	return allClients
}
