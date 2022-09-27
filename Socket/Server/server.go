package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

var port = "0.0.0.0:9001"

func echoMessage(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		message, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Printf("Message Received From Client is %v at %v\n", string(message), time.Now())
			conn.Write([]byte("From Server :"))
			conn.Write(message)
		case io.EOF:
		default:
			fmt.Println("Error ", err)
		}

	}
}

func main() {
	fmt.Println("Start the server ")
	ln, err := net.Listen("tcp", port)
	var clientConnected int
	for {
		conn, _ := ln.Accept()
		fmt.Println("Connected Client : ", conn.RemoteAddr())
		clientConnected++
		fmt.Println("Clients Connected : ", clientConnected)
		if err != nil {
			fmt.Println("Error ", err)
			continue
		}
		go echoMessage(conn)
	}
}
