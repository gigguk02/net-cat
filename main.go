package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

var welcome string = `
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    '.       | '' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     '-'       '--'
`
var (
	join          = make(chan mes)
	channel       = make(chan mes)
	leaving       = make(chan mes)
	history       string
	numberOfUsers = 0

	mutex sync.Mutex
)

type mes struct {
	name,
	text,
	addr,
	ourTime string
}

func main() {
	args := os.Args[1:]
	var port string
	switch len(args) {
	case 0:
		port = ":9090"
	case 1:
		port = args[0]
		_, err := strconv.Atoi(port)
		if err != nil {
			log.Println(err)
			return
		}
		port = ":" + args[0]
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
		return

	}
	fmt.Println("Server is listening...")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
		return
	}
	go broadcast(clients)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		mutex.Lock()
		numberOfUsers++
		mutex.Unlock()
		if numberOfUsers <= 10 {
			if err != nil {
				fmt.Println(err)
				return
			}
			go Handle(conn)

		} else {
			conn.Write([]byte("Chat is full. Try again later!"))
			mutex.Lock()
			numberOfUsers--
			mutex.Unlock()

		}

	}
}
