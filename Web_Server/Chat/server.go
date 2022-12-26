package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)
1111111111111

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}

			conns = append(conns, conn)
			connCh <- conn
		}
	}()

	for {
		select {
		case conn := <-connCh:
			go onMessage(conn)

		case msg := <-msgCh:
			fmt.Print(msg)

		case conn := <-closeCh:
			fmt.Println("client exit")
			removeConn(conn)
		}
	}
}

func removeConn(conn net.Conn) {
	var i int
	for i = range conns {
		if conns[i] == conn {
			break
		}
	}
	conns = append(conns[i:], conns[:i+1]...)
}

func publishMsg(conn net.Conn, msg string) {
	for i := range conns {
		if conns[i] != conn {
			conns[i].Write([]byte(msg))
		}
	}
}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		msgCh <- msg
		publishMsg(conn, msg)

	}

	closeCh <- conn
}
