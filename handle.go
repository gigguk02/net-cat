package main

import (
	"bufio"
	"net"
	"strings"
	"time"
)

var history string
var clients = make(map[string]net.Conn)

func Handle(conn net.Conn) {
	conn.Write([]byte(welcome))
	conn.Write([]byte("[ENTER YOUR NAME]:"))
	userName := bufio.NewScanner(conn)
	var name string
	for userName.Scan() {
		name = strings.TrimSpace(userName.Text())
		if !Correct(name) {
			conn.Write([]byte("Please enter latin or kirill letters! :"))
			continue
		}

		break
	}
	clients[name] = conn
	Message(name, conn)

}

func Message(name string, conn net.Conn) {
	mes1 := bufio.NewScanner(conn)
	for {
		time := "[" + time.Now().Format("01-02-2006 15:04:05") + "]"
		conn.Write([]byte(time + "[" + name + "]: "))
		for mes1.Scan() {
			// conn.Write([]byte(mes1.Text() + "\n"))
			// if !Correct(message) {
			// 	conn.Write([]byte("Please enter latin or kirill letters! :"))
			// 	message = strings.TrimSpace(mes.Text())
			// 	fmt.Println(message)
			// 	continue
			// }

			history = time + "[" + name + "]:" + mes1.Text() + "\n"
			address := conn.RemoteAddr().String()
			anonymStructForChannel := mes{
				name:       name,
				historyMes: history,
				text:       mes1.Text(),
				addr:       address,
			}

			channel <- anonymStructForChannel
			break

		}

	}

}
func Correct(name string) bool {
	if name == "" {
		return false
	}
	for _, i := range name {
		if i < 32 {
			return false
		}
	}
	return true

}

func broadcast(map[string]net.Conn) {
	for {
		select {
		case mesUser := <-channel:
			for _, conn := range clients {

				if mesUser.addr != conn.RemoteAddr().String() {
					conn.Write([]byte(mesUser.name + " has joined our chat..." + "\n"))
					conn.Write([]byte(mesUser.historyMes))
				}
				if mesUser.addr == conn.RemoteAddr().String() {
					conn.Write([]byte(mesUser.historyMes))
					continue

				}

			}
		}
	}

}
