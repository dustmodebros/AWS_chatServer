package main

import ("fmt"
	"flag"
	"bufio"
	"net"
	"strings")


type Message struct {
	sender	int
	message string
}

func handleError(err error){
// TODO: all
// Deal with an error event.
}

func acceptConns(ln net.Listener, conns *chan net.Conn){
// TODO: all
// Accept a network connection from the Listener 
// and add it to the channel for handling connections.
}

func handleClient(client net.Conn, clientid int, msgs *chan Message){
// TODO: all
// Read in new messages delimited by '\n's
// Tidy up each message and add it to the messages channel, 
// recording which client it came from.
}


func main(){
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.

	//Create a channel for connections
	conns	:= make(chan net.Conn)
	//Create a channel for messages
	msgs	:= make(chan Message)
	//Create a channel mapping IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, &conns)
	for {
		select {
			case conn:= <-conns:
				//TODO Deal with a new connection
			case msg := <-msgs:
				//TODO Deal with a new message
			}
	}
}