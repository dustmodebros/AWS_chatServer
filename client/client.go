package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		in, _ := reader.ReadString('\n')
		if in != "\n" {
			fmt.Print("\nreceived message is: ", in)
			fmt.Print("Enter text:")
		}
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text:")
		in, _ := reader.ReadString('\n')
		fmt.Fprintln(conn, in)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	conn, err := net.Dial("tcp", *addrPtr)

	if err != nil {
		fmt.Println("wops:")
		log.Fatal(err)
	}
	go read(conn)

	go write(conn)

	select {}
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
