package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Connect to the server at localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send a message to the server
	fmt.Fprintf(conn, "Hello, Server!\n")

	// Listen for reply
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Message from server: ", message)
}
