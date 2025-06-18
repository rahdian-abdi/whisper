# Whisper C2 Framework

![Status](https://img.shields.io/badge/status-proof%20of%20concept-brightgreen) [![Go Version](https://img.shields.io/badge/go-1.18+-00ADD8.svg)](https://go.dev/) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Whisper is a lightweight, TCP-based Command and Control (C2) framework built entirely in Go. It was developed for educational purposes to explore the fundamentals of C2 architecture and the challenges of evading modern security solutions.

This project serves as a practical tool for researchers, students, and cybersecurity professionals to understand offensive tooling in a controlled, ethical environment.

---

## ⚠️ Disclaimer

This project is intended **solely for educational and research purposes**. The goal is to study adversarial techniques to build better defensive strategies. The author is not responsible for any misuse or damage. **Use this tool ethically and only on systems you have explicit permission to test.**

---

## 🧠 Core Concepts: Bypassing Static vs. Dynamic Analysis

A key goal of this project is to understand how security products detect threats. The current version of Whisper can remain undetected by some basic antivirus (AV) solutions for two main reasons:

1.  **Unique Signature:** As a custom-coded tool, its file hash is not in any AV vendor's database of known malware. It bypasses **static analysis** because it has no recognized malicious signature.

2.  **Simple Behavior:** The agent's current actions (making a TCP connection, running a command) are not inherently malicious enough to trigger alarms in AV products that are not performing aggressive **behavioral analysis**.

The real challenge, and the future goal of this project, is to evade **Endpoint Detection and Response (EDR)** systems. An EDR actively monitors system behavior and would likely flag Whisper for:

* An unknown, unsigned program making a persistent network connection.
* An unusual parent-child process relationship (e.g., `agent.exe` spawning `whoami.exe`).

The roadmap below outlines the features needed to bypass this more advanced, behavior-based detection.

---

## ✨ Current Features

* **TCP-Based C2 Channel**: Simple and direct network communication between the agent and server.
* **Remote Command Execution**: Execute shell commands on a target machine via the agent.
* **Asynchronous Command Handling**: Built with Go's goroutines to manage multiple agents without blocking.
* **Cross-Platform Portability**: Easily cross-compile the agent for Windows, Linux, and macOS.

---

## 🗺️ The Road to True Stealth: Project Roadmap

The following features are planned to evolve Whisper from a simple C2 into a framework capable of evading more sophisticated, behavior-based security solutions.

* [X] **Encrypted C2 Channels (TLS):** Encrypt all traffic between the agent and server to prevent network inspection and hide commands.
* [ ] **Malleable C2 Profiles:** Disguise C2 traffic to look like legitimate web traffic (e.g., HTTP/S POST requests) to blend in with normal network activity.
* [ ] **Sleep Obfuscation & Jitter:** Introduce randomized delays (jitter) to the agent's check-ins to make its beaconing pattern less predictable and harder to detect.

---


## 🛠️ Getting Started

### Prerequisites

* [Go](https://go.dev/doc/install) (Version 1.18 or higher)
* A basic understanding of networking and command-line interfaces.

### Installation and Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/rahdian-abdi/whisper.git
    cd whisper
    ```

2.  **Configure the Agent:**
    Before compiling, open `agent/main.go` and set the `c2Host` and `c2Port` variables

3.  **Compile the binary:**
    For evading the AV, try to compile the server with this command
    ```bash
    GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o msbuild64.exe main.go
    ```

4.  **TLS Certificate:**
    This C2 uses TLS for encrypted communication.  
    You must generate your own certificate and private key and create `certs` folder in `server` folder

    ```bash
    openssl req -new -x509 -keyout certs/c2.key -out certs/c2.crt -days 365 -nodes
