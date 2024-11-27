package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"netcat/functions"
)

func main() {
	// define default port if user does not specify a port
	port := "8989"
	args := os.Args

	if len(args) == 2 {
		port = args[1]
	} else if len(args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	serverlistener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Error starting TCP server: %v", err)
		return
	}
	defer serverlistener.Close()
	fmt.Printf("Chat server started on port %s---\n", port)

	for {
		connection, err := serverlistener.Accept()
		if err != nil {
			log.Printf("Failed to accept connections :%v", err)
			continue
		}
		go functions.ClientHandler(connection)
	}
}
