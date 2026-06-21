// Command tcpchat is a simple TCP chat server for learning networking.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	addr := ":9000"
	if v := os.Getenv("TCP_ADDR"); v != "" {
		addr = v
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "listen: %v\n", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Printf("TCP chat server listening on %s\n", addr)

	var mu sync.Mutex
	clients := make(map[net.Conn]bool)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		mu.Lock()
		clients[conn] = true
		mu.Unlock()

		go handleClient(conn, &mu, clients)
	}
}

func handleClient(conn net.Conn, mu *sync.Mutex, clients map[net.Conn]bool) {
	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		msg := strings.TrimSpace(line)
		broadcast(mu, clients, fmt.Sprintf("[%s] %s\n", conn.RemoteAddr(), msg))
	}
}

func broadcast(mu *sync.Mutex, clients map[net.Conn]bool, msg string) {
	mu.Lock()
	defer mu.Unlock()
	for c := range clients {
		_, _ = c.Write([]byte(msg))
	}
}
