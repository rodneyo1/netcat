package functions

import "fmt"

func unregisterClient(client *Client) {
	clientsMutex.Lock()
	delete(clients, client.conn)
	clientsMutex.Unlock()

	farewellMessage := fmt.Sprintf("%s has left the chat", client.name)
	broadcast(farewellMessage, client)
	client.conn.Close()
}

func sendHistory(client *Client) {
	historyMutex.Lock()
	for _, msg := range msgHistory {
		fmt.Fprintln(client.conn, msg)
	}
	historyMutex.Unlock()
}

func logMessage(message string) {
	historyMutex.Lock()
	msgHistory = append(msgHistory, message)
	historyMutex.Unlock()
}

func broadcast(message string, sender *Client) {
	clientsMutex.Lock()
	for _, client := range clients {
		if client != sender {
			fmt.Fprintln(client.conn, message)
		}
	}
	clientsMutex.Unlock()
}
