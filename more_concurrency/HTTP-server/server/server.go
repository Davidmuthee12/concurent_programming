package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

var r, _ = regexp.Compile("GET (.+) HTTP/1.1\r\n")

func handleHTTPRequest(conn net.Conn) {
	// Creates buffer to store HTTP request
	buff := make([]byte, 1024)
	// Reads from the connection into the buffer
	size, _ := conn.Read(buff)
	if r.Match(buff[:size]) {
		file, err := os.ReadFile(
			fmt.Sprintf("../resources/%s", r.FindSubmatch(buff[:size])[1]))

			// if the file exists, responds to the client with the HTTP header and file contents
		if err == nil {
			conn.Write([]byte(fmt.Sprintf(
				"HTTP/1.1 200 ok\r\nContent-length: %d\r\n\r\n", len(file))))
			conn.Write(file)
		} else {
			// if the file does not exist, responds with an error
			conn.Write([]byte(
				"HTTP/1.1 404 NOT Found\r\n\r\n<html>Not Found</html>"))
		}
	} else {
		// if the HTTP request is not valid, respond with an error
		conn.Write([]byte("HTTP/1.1 500 Internal server Error\r\n\r\n"))
	}
	conn.Close() // Closes the connection after handling the request
}

// The following function initializez all the goroutines in the worker pool.
// The function simply starts n goroutines, each reading from the input channel containing the client connections.
// when a new connection is received on the channel, the handleHTTPRequest() function is called to handle the  clients request

func startHTTPWorkers(n int, incomingConnections <-chan net.Conn) {
	for i := 0; i < n; i++ {
		// Starts n goroutines
		go func() {
			// Consumes connections from the work queue channel until the channel is closed
			for c := range incomingConnections {
				// Handles the HTTP request from the received connection
				handleHTTPRequest(c)  
			}
		}()
	}
}

// This main() goroutine listens for new connections on a port and pass on 
// any newly established connection on the work queue channel
// the main() function creates the work queue channel, starts up the worker pool
// and then binds a TCP listen connection on port 8080. 
// In an infinite loop, when a new connection is established, the Accept() function unblocks
// and returns the connection.
// This connection is then passed on the channel to be used by one of the goroutines in the worker pool

func main() {
	incomingConnections := make(chan net.Conn)

	startHTTPWorkers(3, incomingConnections)

	server, err := net.Listen("tcp", "localhost:8087")
	if err != nil {
		fmt.Println("Failed to listen on port 8080:", err)
		return
	}

	defer server.Close()

	fmt.Println("Server listening on http://localhost:8087")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		incomingConnections <- conn
	}
}