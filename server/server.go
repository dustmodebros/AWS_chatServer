package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// Deal with an error event.
	fmt.Println("Err:", err)
	return
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for {
		conn, errA := ln.Accept()
		//fmt.Println("got connection:", conn)
		if errA != nil {
			handleError(errA)
		}
		//fmt.Println("sending conn through channel:")
		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	reader := bufio.NewReader(client)
	for {
		rec, errR := reader.ReadString('\n')
		if errR != nil {
			handleError(errR)
		}
		//fmt.Println("sending", rec, "down the messages channel")
		if rec != "\n" {
			msgs <- Message{clientid, rec}
		}
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)
	//fmt.Println("started listening on port", *portPtr)
	//Create a channel for connections
	conns := make(chan net.Conn)
	//fmt.Println("made new channel for connections")
	//Create a channel for messages
	msgs := make(chan Message)
	//fmt.Println("made new channel for messages")

	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)
	//fmt.Println("made mapping")

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//fmt.Println("got new connection through conns channel", conn)
			clientID := len(clients)
			clients[clientID] = conn
			go handleClient(conn, clientID, msgs)
		case msg := <-msgs:
			if msg.message != "\n" {
				fmt.Print("new message received: ", msg.message)
				for id, conn := range clients {
					if id != msg.sender {
						fmt.Fprintln(conn, msg.message)
					}
				}
			}
		}
	}
}
