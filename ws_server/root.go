package wsserver

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	wsInit sync.Once
)

func NewServer() {
	wsInit.Do(func() {
		http.HandleFunc("/", handleBroadcast)
		fmt.Println("WebSocket server is listening on :8081")
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			fmt.Println("Error starting WebSocket server:", err)
		}
	})
}
