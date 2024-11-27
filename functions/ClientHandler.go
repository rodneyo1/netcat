package functions

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func displayArt(conn net.Conn) {
	art := ReadArt("logo.txt")
	conn.Write([]byte(art))
}

type Client struct {
	conn net.Conn
	name string
}

var (
	clients      = make(map[net.Conn]*Client)
	msgHistory   []string
	historyMutex sync.Mutex
	clientsMutex sync.Mutex
)

func ClientHandler(conn net.Conn) {
	defer conn.Close()
	displayArt(conn)
	fmt.Fprint(conn, "ENTER YOUR NAME: ")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	name := scanner.Text()
	client := &Client{conn: conn, name: name}
	registerClient(client)
	defer unregisterClient(client)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "" {
			continue
		}
		// Check for the /name command
		if len(message) > 6 && message[:6] == "/name " {
			newName := message[6:]
			if newName != "" {
				clientsMutex.Lock()
				oldName := client.name
				client.name = newName
				clientsMutex.Unlock()
				notification := fmt.Sprintf("%s has changed their name to %s", oldName, newName)
				broadcast(notification, client)
				fmt.Fprintf(conn, "Your name has been changed to %s\n", newName)
			} else {
				fmt.Fprintln(conn, "Name change failed. New name cannot be empty.")
			}
			continue
		}
		// Standard message handling
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, client.name, message)
		logMessage(formattedMessage)
		broadcast(formattedMessage, client)
	}
	if scanner.Err() != nil {
		log.Printf("Error reading from client %s: %v", client.name, scanner.Err())
	}
}
func registerClient(client *Client) {
	clientsMutex.Lock()
	clients[client.conn] = client
	clientsMutex.Unlock()
	welcomeMessage := fmt.Sprintf("%s has joined the chat", client.name)
	broadcast(welcomeMessage, client)
	sendHistory(client)
}
