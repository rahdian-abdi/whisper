# Whisper C2 Framework

![Status](https://img.shields.io/badge/status-proof%20of%20concept-brightgreen) [![Go Version](https://img.shields.io/badge/go-1.18+-00ADD8.svg)](https://go.dev/) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Whisper is a lightweight Command and Control (C2) framework built entirely in Go. It was developed for educational purposes to explore the fundamentals of C2 architecture and the challenges of evading modern security solutions.

This project serves as a practical tool for researchers, students, and cybersecurity professionals to understand offensive tooling in a controlled, ethical environment.

---

## ⚠️ Disclaimer

This project is intended **solely for educational and research purposes**. The goal is to study adversarial techniques to build better defensive strategies. The author is not responsible for any misuse or damage. **Use this tool ethically and only on systems you have explicit permission to test.**

---

## 🧠 Core Concepts: Bypassing Static vs. Dynamic Analysis

A key goal of this project is to understand how security products detect threats. The current version of Whisper can remain undetected by some basic antivirus (AV) solutions for two main reasons:

1.  **Unique Signature:** As a custom-coded tool, its file hash is not in any AV vendor's database of known malware. It bypasses **static analysis** because it has no recognized malicious signature.

2.  **Simple Behavior:** The agent's current actions are not inherently malicious enough to trigger alarms in AV products that are not performing aggressive **behavioral analysis**.

The real challenge, and the future goal of this project, is to evade **Endpoint Detection and Response (EDR)** systems. An EDR actively monitors system behavior and would likely flag Whisper for:

* An unknown, unsigned program making a persistent network connection.
* An unusual parent-child process relationship (e.g., `agent.exe` spawning `whoami.exe`).

---

## ✨ Current Features

* **Encrypted C2 Channels (TLS)**: Secure traffic using auto-generated certificates.
* **Malleable C2 Profiles**: Agent mimics web traffic (e.g. HTTPS endpoints like `/cdn/image`).
* **Sleep Obfuscation & Jitter**: Randomized beacon intervals for evasion.

---

## 🗺️ Supported Platforms

The following features are planned to evolve Whisper from a simple C2 into a framework capable of evading more sophisticated, behavior-based security solutions.

* [✅] **Windows**
* [⏳] **Linux** 
* [⏳] **MacOS**

---


## 🛠️ Installation & Usage

### Prerequisites

* [Go](https://go.dev/doc/install) (Version 1.18 or higher)
* A basic understanding of networking and command-line interfaces.

### Installation and Setup

First, clone the repository to your local machine:


    git clone https://github.com/rahdian-abdi/whisper.git
    cd whisper


### **Option 1: Using `go run` (For Development)**

You can use `go run` directly during testing or while modifying the framework.

#### **Server Setup (Raw TCP)**

1.  **Generate TLS Certificates:**
    ```bash
    go run ./cmd/whisper.go certificate
    ```

2.  **Generate an Agent:**
    ```bash
    go run ./cmd/whisper.go generate rawtcp [output.exe] [C2_Server:Port]
    ```
3.  **Start Listener:**
    ```bash
    go run ./cmd/whisper.go listen rawtcp [port]
    ```    

### **Option 2: Compile Whisper to a Single Binary**

To make Whisper portable and easier to deploy:

```bash
go build -o whisper.exe ./cmd/whisper.go
```
Then use it just like a standard CLI:
#### **Server Setup (HTTPS)**

1.  **Generate Certificate:**
    ```bash
    ./whisper.exe certificate
    ```

2.  **Generate Agent:**
    ```bash
    ./whisper.exe generate https agent.exe https://[C2_Server]
    ```

3.  **Start Listener:**
    ```bash
    ./whisper.exe listen https
    ```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page for this repository to contribute.

---

## 📜 License

Distributed under the MIT License. See `LICENSE` for more information.