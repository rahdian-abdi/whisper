package certgen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	CertDir  = "whispercert"
	CertFile = filepath.Join(CertDir, "c2.crt")
	KeyFile  = filepath.Join(CertDir, "c2.key")
)

func EnsureCertificate() error {
	if fileExists(CertFile) && fileExists(KeyFile) {
		fmt.Println("[*] TLS cert already generated...")
		return nil
	}

	fmt.Println("[*] TLS cert not found, generating one in ./whispercert...")
	if err := os.MkdirAll(CertDir, 0755); err != nil {
		return err
	}

	cmd := exec.Command("openssl", "req", "-new", "-x509",
		"-keyout", KeyFile,
		"-out", CertFile,
		"-days", "365",
		"-nodes",
		"-subj", "/CN=whisper-c2")

	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate TLS cert: %w", err)
	}
	fmt.Println("[*] TLS cert generated successfully.")
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
