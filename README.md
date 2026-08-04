# Concurrent TCP Chat Server

A concurrent multi-client chat server built in Go to explore networking, concurrency, and server architecture from first principles.

This project uses Go's `net` package directly to implement TCP communication, connection management, synchronized shared state, and graceful shutdown. A lightweight Python asyncio client was also developed to compare Go's goroutine model with Python's asynchronous programming model.

---

## Features

* Multi-client TCP chat server
* Goroutine-per-client connection handling
* Thread-safe client registry using `sync.RWMutex`
* Broadcast messaging to all connected clients
* Graceful client connection and disconnection handling
* Graceful server shutdown using OS signal handling
* Go and Python client implementations
* Modular server architecture with separated networking responsibilities

---

## Architecture

```text
                TCP Connections

        Client A        Client B        Client C
            │               │               │
            └───────────────┼───────────────┘
                            │
                    Go TCP Server
                            │
                    handleClient()
                            │
                    ClientRegistry
          ┌─────────────────┴─────────────────┐
          │                                   │
     Message Receiver                  Broadcast Manager
```

Each client connection is serviced by its own goroutine while a shared `ClientRegistry` maintains active connections with synchronized access.

---

## Project Structure
```
.
├── cmd/
│   ├── server/
│   │   ├── main.go
│   │   ├── handleClient.go
│   │   ├── broadcast.go
│   │   ├── clientConnection.go
│   │   ├── msg_receiver.go
│   │   └── shutdownServer.go
│   │
│   ├── client/
│   │   └── main.go
│   │
│   └── client-python/
│       └── client.py
│
├── docs/
│   └── dev-logs.md
│
├── go.mod
└── README.md
```
---

## Design Highlights

### Concurrent Client Handling

Each accepted TCP connection is handled in its own goroutine, allowing multiple clients to communicate with the server simultaneously.

### Thread-Safe Shared State

Connected clients are managed through a `ClientRegistry` abstraction protected by `sync.RWMutex`, preventing concurrent access to shared state while encapsulating registry operations.

### Connection Lifecycle

Each client handler owns the lifecycle of its connection.

Cleanup includes:

* removing disconnected clients
* closing network connections
* preventing duplicate cleanup through idempotent removal logic

### Broadcasting

Broadcast operations send messages to every connected client while isolating failures so one disconnected client does not interrupt delivery to the remaining clients.

### Graceful Shutdown

The server listens for operating system interrupt signals.

Shutdown sequence:

1. Stop accepting new connections.
2. Close all active client connections.
3. Allow client goroutines to terminate naturally.
4. Exit cleanly.

---

## Python Client

A companion asyncio client was implemented to compare Go's concurrency model with Python's event loop.

Key concepts explored include:

* `asyncio.TaskGroup`
* `asyncio.StreamReader`
* `asyncio.StreamWriter`
* asynchronous send/receive loops
* `asyncio.to_thread()` for non-blocking console input

---

## Learning Objectives

This project was built to strengthen foundational knowledge in backend and systems programming before developing larger distributed applications.

Topics explored include:

* TCP networking
* Concurrent server design
* Goroutines and synchronization
* Mutexes (`sync.RWMutex`)
* Connection ownership
* Resource cleanup
* Graceful shutdown
* Client-server communication
* Python asyncio
* Modular software architecture

---

## Potential Improvements

* Message framing protocol instead of newline delimiters
* Client authentication
* Private messaging
* Chat rooms
* Structured logging
* Configuration via environment variables
* Unit and integration tests
* TLS-encrypted connections
* Persistent message storage

---

## Technologies

* Go
* Python
* TCP/IP
* Goroutines
* `sync.RWMutex`
* asyncio

---

# Getting Started

Follow these steps to run the TCP chat server locally.

## Prerequisites

Install the following:

* Go 1.XX or later
* Python 3.XX or later (only required for the Python asyncio client)

Verify installations:

```bash
go version
python --version
```

---

## Clone the Repository

```bash
git clone https://github.com/Haris270/go-tcp-server
cd go-tcp-server
```

Install Go dependencies:

```bash
go mod tidy
```

---

# Running the Server

From the project root:

```bash
go run ./cmd/server
```

The server starts listening for TCP connections on:

```text
localhost:8080
```

Expected output:

```text
Server listening on localhost:8080
```

---

# Running Clients

The project includes both a Go client and a Python asyncio client.

## Go Client

Open a new terminal window from the project root:

```bash
go run ./cmd/client
```

Enter a username when prompted.

---

## Python Asyncio Client

Open another terminal:

```bash
python ./cmd/client-python/client.py
```

The Python client uses only Python's standard `asyncio` library and does not require additional dependencies.

---

# Running Multiple Clients

To test concurrent communication:

1. Start the server.
2. Launch multiple client instances.
3. Connect using different usernames.
4. Send messages between clients.

Example:

```text
John: Hello everyone!
Ann: Hi John!
```

Messages are broadcast to all active connections.

---

# Graceful Shutdown

The server supports graceful shutdown through OS interrupt handling.

Stop the server using:

```text
CTRL + C
```

The shutdown process:

1. Stops accepting new client connections.
2. Closes active client connections.
3. Cleans up resources.
4. Terminates the server process.

---

# Development

Run the Go race detector to validate concurrent access:

```bash
go run -race ./cmd/server
```

The race detector helps identify unsafe concurrent access between goroutines.


