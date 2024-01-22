package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

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
		_, checkName := clients[name]
		if checkName || len(name) == 0 {
			conn.Write([]byte("Enter the correct name or this Username already exists\n" + "[ENTER YOUR NAME]: "))

		} else {
			clients[name] = conn
			time := "[" + time.Now().Format("01-02-2006 15:04:05") + "]"
			address := conn.RemoteAddr().String()
			conn.Write([]byte(history))
			history += name + " has joined our chat...\n"
			structForChannel := mes{
				text:    "has joined our chat",
				ourTime: time,
				name:    name,
				addr:    address,
			}
			join <- structForChannel
			Message(name, conn)
			break

		}

	}

}

func Message(name string, conn net.Conn) {
	mes1 := bufio.NewScanner(conn)
	for mes1.Scan() {
		if Correct(mes1.Text()) {
			time := "[" + time.Now().Format("01-02-2006 15:04:05") + "]"
			address := conn.RemoteAddr().String()
			history += time + "[" + name + "]:" + mes1.Text() + "\n"
			anonymStructForChannel := mes{
				name:    name,
				text:    mes1.Text(),
				addr:    address,
				ourTime: time,
			}
			channel <- anonymStructForChannel

		} else {
			conn.Write([]byte("Please enter latin or kirill letters! :"))
		}

	}
	mutex.Lock()
	delete(clients, name)
	numberOfUsers--
	mutex.Unlock()
	address := conn.RemoteAddr().String()
	time := "[" + time.Now().Format("01-02-2006 15:04:05") + "]"
	history += name + " has left our chat...\n"
	structForChannel := mes{
		text:    "has left our chat...",
		name:    name,
		addr:    address,
		ourTime: time,
	}
	leaving <- structForChannel

}

func broadcast(map[string]net.Conn) {
	for {
		select {
		case mesUser := <-channel:
			for name, conn := range clients {
				if mesUser.addr == conn.RemoteAddr().String() {
					fmt.Fprint(conn, mesUser.ourTime+"["+mesUser.name+"]:")
				}
				if mesUser.addr != conn.RemoteAddr().String() {
					fmt.Fprint(conn, ClearLine(mesUser.text)+mesUser.ourTime+"["+name+"]:\n")
					fmt.Fprint(conn, ClearLine(mesUser.text)+mesUser.ourTime+"["+mesUser.name+"]:"+mesUser.text+"\n")
					fmt.Fprint(conn, mesUser.ourTime+"["+name+"]:")
				}

			}
		case join := <-join:
			for name, conn := range clients {
				if join.addr != conn.RemoteAddr().String() {
					fmt.Fprint(conn, "\n"+join.name+" has joined our chat...\n")
					clients[name].Write([]byte(join.ourTime + "[" + name + "]:"))

				}

				if join.addr == conn.RemoteAddr().String() {
					fmt.Fprint(conn, join.ourTime+"["+join.name+"]:")
				}

			}
		case left := <-leaving:
			for name, conn := range clients {
				if left.addr != conn.RemoteAddr().String() {
					fmt.Fprint(conn, "\n"+left.name+" has left our chat...\n")
					fmt.Fprint(conn, left.ourTime+"["+name+"]:")
				}

			}

		}
	}
}
func ClearLine(s string) string {
	return "\r" + strings.Repeat(" ", len(s)+len(s)/2) + "\r"
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
