package rawtcp

import (
	"crypto/tls"
	"fmt"
	"log"
	"whisper/internal/certgen"
	"whisper/server/rawtcp/handler"
)

func Start(port string) {

	cert, err := tls.LoadX509KeyPair(certgen.CertFile, certgen.KeyFile)
	if err != nil {
		log.Fatalf("Failed to load cert/key: %s", err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	addr := ":" + port
	ln, err := tls.Listen("tcp", addr, config)
	if err != nil {
		log.Fatalf("Failed to start TLS listener: %s", err)
	}
	defer ln.Close()
	connetionMessage := fmt.Sprintf("[*] Whisper C2 listening on 0.0.0.0%s (TLS)", addr)
	fmt.Println(connetionMessage)

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
