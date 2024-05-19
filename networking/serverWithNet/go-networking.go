package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen on TCP port 8080 on all interfaces.
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		// Accept a connection.
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the connection in a new goroutine.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		fmt.Println("Handling message in server.")
		// Read the incoming connection into the buffer.
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Message Received:", message)
		// Send a response back to the client.
		conn.Write([]byte("Message received.\n"))
	}
}
