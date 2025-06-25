package handler

import (
	"fmt"
	"io"
	"net/http"
)

func HandleResult(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, _ := io.ReadAll(r.Body)
	fmt.Printf("[Agent Output] %s\n", string(data))
}
