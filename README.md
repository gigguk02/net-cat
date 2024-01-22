# NET-CAT

## TCPChat - Group Chat Server and Client

TCPChat is a simple group chat implementation using a Server-Client architecture. It allows multiple clients to connect to a central server and exchange messages in a group chat format. The server handles incoming connections, manages client messages, and broadcasts messages to all connected clients.

## Features 
- TCP connection between the server and multiple clients (one-to-many).
- Clients are required to provide a name.
- Control the maximum number of allowed connections (10 connections).
- Clients can send messages to the chat.
- Empty messages are not broadcasted.
- Messages are timestamped with the sender's name and time of sending.
- New clients receive previous chat history upon joining.
- Clients are notified when another client joins or leaves.
- All clients receive messages sent by other clients.
- Clients can disconnect without affecting other clients.
- Default port is 8989, but can be specified as a command line argument.

## Usage
Compile the project:
```bash
go build -o Tcpchat main.go
```

Run the server:
```bash
./Tcpchat [port]
```
* `port` (optional): The port number to listen on. If not provided, the default port is 8989 is used.

Clients can connect to the server using `nc` command:
```bash
nc localhost [port]
```

If you want to connect to other computers USE flag (-localconn):
```bash
./Tcpchat --localconn 
```
Clients can connect to the server using `nc` command and host address [10.42.0.1] and port address [27960]:
```bash
nc 10.42.0.1 27960
```
### !!!NEED TURN ON HOTSPOT!!!
Each client is prompted to enter a name. Once connected, clients can start sending messages to the chat.


To exit, simply close the client's terminal.

## Example
```bash
./tcpchat 8888
```
```bash
nc localhost 8888
```

## Dependencies
Go version 1.20.1
Used only: 
- io
- log
- os
- fmt
- net
- sync
- time
- bufio
- errors
- strings
- reflect


## Author
[@gbekbola](https://01.alem.school/git/gbekbola)
