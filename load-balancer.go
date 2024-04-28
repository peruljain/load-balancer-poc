package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

var serverHosts = []string{
	"localhost:8080",
	"localhost:8081",
	"localhost:8082",
	"localhost:8083",
}

var index = 0

func main() {

	 // Open a file for logging
	 logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	 if err != nil {
		 log.Fatal("Failed to open log file:", err)
	 }
	 defer logFile.Close()
 
	 // Set log output to the file
	 log.SetOutput(logFile)


    // Listen on local port 8080
    listener, err := net.Listen("tcp", ":9090")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    for {
        // Accept a connection
        conn, err := listener.Accept()
		clientIP := conn.RemoteAddr().String()
     	log.Printf("Connection accepted from %s", clientIP)
        if err != nil {
            log.Print(err)
            continue
        }

        // Handle the connection in a new goroutine
        go handleConnection(conn)
    }
}

func handleConnection(clientConn net.Conn) {
    defer clientConn.Close()
	var hostIndex = index%len(serverHosts)
	index = index + 1
	log.Printf("host %s", strconv.Itoa(hostIndex))
    serverConn, err := net.Dial("tcp", serverHosts[hostIndex])
    if err != nil {
        log.Print(err)
        return
    }
    defer serverConn.Close()

    var wg sync.WaitGroup
    wg.Add(2) // Add for two goroutines

    go func() {
        defer wg.Done()
        io.Copy(serverConn, clientConn) // client to server
    }()
    
    go func() {
        defer wg.Done()
        io.Copy(clientConn, serverConn) // server to client
    }()

    wg.Wait() // Wait for both copies to complete
}
