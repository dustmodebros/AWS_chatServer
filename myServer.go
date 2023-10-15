package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(connection net.Conn) {
	reader := bufio.NewReader(connection)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Print(msg)
		fmt.Fprintln(connection, "OK")
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8000")
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
