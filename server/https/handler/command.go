package handler

import (
	"fmt"
	"net/http"
)

var CurrentCommand = make(chan string, 1)

func HandleTask(w http.ResponseWriter, r *http.Request) {
	select {
	case cmd := <-CurrentCommand:
		fmt.Println("[*] Sent to Agent:", cmd)
		w.Write([]byte(cmd))
	default:
		w.Write([]byte("none"))
	}

}
