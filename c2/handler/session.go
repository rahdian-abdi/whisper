package handler

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func HandleSession(conn net.Conn) {
	defer conn.Close()
	fmt.Println("[*] Handler stub activated.")

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("[-] Agent disconnected:", err)
				return
			}
			fmt.Printf("[Agent] %s\n", string(buf[:n]))
			fmt.Print("[whisper]> ")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("[whisper]> ")
	for {
		if !scanner.Scan() {
			fmt.Println("[-] Input error or EOF.")
			return
		}
		command := scanner.Text() + "\n"
		_, err := conn.Write([]byte(command))
		if err != nil {
			fmt.Println("[-] Failed to send command:", err)
			return
		}
	}
}
