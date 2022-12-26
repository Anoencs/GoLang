package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var (
	conns   []net.Conn
	connCh  = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgCh   = make(chan string)
)

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')

		fmt.Print(msg)
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

func onMessageServer(conn net.Conn) {
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
func ConnectS(addr string) {
	connection, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", addr))
	if err != nil {

	}

	fmt.Print("your name:")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')

	nameInput = nameInput[:len(nameInput)-1]

	fmt.Println("********** MESSAGES **********")

	go onMessage(connection)

	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s: %s\n", nameInput,
			msg[:len(msg)-1])

		connection.Write([]byte(msg))
	}

	connection.Close()
}
func ListenS(addr string) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
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

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		ListenS("2901")
		wg.Done()
	}()
	go func() {
		ConnectS("2900")
		wg.Done()
	}()
	wg.Wait()
}
