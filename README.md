# Netcat

## **Description**

This project implements a simple TCP-based chat server in Go. The server allows multiple clients to connect, enabling them to communicate in real time. It demonstrates the use of Go's `net` package for TCP connections and employs concurrency to handle multiple clients simultaneously.

## **Features**

* **Customizable Port:** Specify the port to run the server on, or use the default port `8989`.  
* **Concurrent Clients:** Handles multiple client connections using goroutines.  
* **Modular Design:** Separation of concerns through the `functions` package, which manages client connections and other functionalities.

## **How to Run**

Build the project:  
bash  
Copy code  
`go build -o TCPChat`

1. 

Run the server:  
bash  
Copy code  
`./TCPChat [port]`

2.   
   * Replace `[port]` with your desired port number. If omitted, the server defaults to port `8989`.

Connect to the server using a TCP client, such as `netcat`:  
bash  
Copy code  
`nc localhost [port]`

3. Replace `[port]` with the server's port.

## **Directory Structure**

* `main.go`: The main entry point of the application, setting up the server and handling incoming connections.  
* `functions/ClientHandler.go`: Contains the logic for handling individual client connections.  
* `functions/ReadArt.go`: Auxiliary functions (likely related to ASCII art handling or other utilities).

## **Dependencies**

* Standard Go libraries (`net`, `os`, `fmt`, `log`).

## **Example**

Run the server:

bash  
Copy code  
`./TCPChat 9090`

Connect to the server:

bash  
Copy code  
`nc localhost 9090`

Start chatting\!
