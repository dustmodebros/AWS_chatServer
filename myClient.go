package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func acknowledge(connection net.Conn) {
	reader := bufio.NewReader(connection)
	msg, _ := reader.ReadString('\n')
	fmt.Print(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	for {
		fmt.Printf("enter text->")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		acknowledge(conn)
	}
}
