package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

type LamportClock struct {
	time int
	mu   sync.Mutex
}

func (lc *LamportClock) IncrementClock() int {
	lc.mu.Lock()
	lc.time++
	updatedT := lc.time
	lc.mu.Unlock()
	return updatedT
}

func (lc *LamportClock) UpdateClock(t int) int {
	lc.mu.Lock()
	updatedT := max(t, lc.time) + 1
	lc.time = updatedT
	lc.mu.Unlock()
	return updatedT
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <port>")
		return
	}
	port := os.Args[1]
	lc := &LamportClock{time: 0}

	go receiveMessage(lc, port)
	go handleEvents(lc)

	select {}
}

func receiveMessage(lc *LamportClock, port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening: ", err)
		return
	}
	defer ln.Close()

	var t int
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			continue
		}

		_, err = fmt.Fscanf(conn, "%d\n", &t)
		if err != nil {
			continue
		}

		sender := conn.RemoteAddr().String()
		updatedT := lc.UpdateClock(t)
		fmt.Printf("\nMessage Received from %s!!\n", sender)
		fmt.Printf("Time: %d\n\n", updatedT)
		prompt()
	}
}

func handleEvents(lc *LamportClock) {
	for {
		prompt()

		var event, port string
		switch fmt.Scan(&event); event {
		case "c":
			updatedT := lc.IncrementClock()
			fmt.Printf("Calculate Event Success!!\nTime: %d\n\n", updatedT)
		case "s":
			fmt.Print("Please type the destination port: ")
			fmt.Scan(&port)
			address := fmt.Sprintf("localhost:%s", port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Println("Connection error: ", err)
				continue
			}
			updatedT := lc.IncrementClock()
			fmt.Fprintf(conn, "%d\n", updatedT)
			fmt.Printf("Sending Event Success!!\nTime: %d\n\n", updatedT)
		default:
			fmt.Print("Error: Please type 'c' or 's'")
		}
	}
}

func prompt() {
	fmt.Print("Please type event, c or s: ")
}
