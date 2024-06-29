package main

import (
	"fmt"
	"log"
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
	newT := lc.time
	lc.mu.Unlock()
	return newT
}

func (lc *LamportClock) UpdateClock(t int) int {
	lc.mu.Lock()
	newT := max(t, lc.time) + 1
	lc.time = newT
	lc.mu.Unlock()
	return newT
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
		log.Println("Error listening: ", err)
		return
	}
	defer ln.Close()

	var t int
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting: ", err)
			continue
		}

		_, err = fmt.Fscanf(conn, "%d\n", &t)
		if err != nil {
			continue
		}

		newT := lc.UpdateClock(t)
		fmt.Printf("\nMessage Received!!\nTime: %d\n\n", newT)
		prompt()
	}
}

func handleEvents(lc *LamportClock) {
	for {
		prompt()

		var event, port string
		switch fmt.Scan(&event); event {
		case "c":
			newT := lc.IncrementClock()
			fmt.Printf("Calculate Event Success!!\nTime: %d\n\n", newT)
		case "s":
			fmt.Print("Please type the destination port: ")
			fmt.Scan(&port)
			address := fmt.Sprintf("localhost:%s", port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				log.Println("Connection error: ", err)
				continue
			}
			newT := lc.IncrementClock()
			fmt.Fprintf(conn, "%d\n", newT)
			fmt.Printf("Sending Event Success!!\nTime: %d\n\n", newT)
		default:
			log.Print("Error: Please type 'c' or 's'")
		}
	}
}

func prompt() {
	fmt.Print("Please type event, c or s: ")
}
